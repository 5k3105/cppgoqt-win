package qt

import (
	"os"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	_libOnce sync.Once
	_lib     uintptr
	pgFree   func(uintptr)
	pgMalloc func(uintptr) uintptr
)

// _load opens the shim library once (openShimLib / shimLibName are per-GOOS — see
// gen_purego_unix.go / gen_purego_windows.go). Override the path with $CPPGO_LIB_QT.
// The allocator pair is bound from the SHIM LIBRARY ITSELF (cppgo_alloc.cpp): Go frees
// exactly what the shims malloc'd — same C runtime by construction, no libc.so.6
// assumption, and it ports to Windows unchanged.
func _load() {
	_libOnce.Do(func() {
		path := os.Getenv("CPPGO_LIB_QT")
		if path == "" {
			path = shimLibName
		}
		h, err := openShimLib(path)
		if err != nil || h == 0 {
			msg := "cppgo/purego: cannot load " + path
			if err != nil {
				msg += ": " + err.Error()
			}
			panic(msg)
		}
		_lib = h
		purego.RegisterLibFunc(&pgFree, _lib, "cppgo_free")
		purego.RegisterLibFunc(&pgMalloc, _lib, "cppgo_malloc")
	})
}

func reg(p any, sym string) {
	_load()
	purego.RegisterLibFunc(p, _lib, sym)
}

// symaddr resolves a shim symbol's address once at init. Wrappers with register-only signatures call
// it directly via purego.SyscallN — no per-call reflection marshalling (several times cheaper than a
// RegisterLibFunc func var). Float-carrying signatures still go through reg() (SyscallN is
// integer-register only).
func symaddr(p *uintptr, sym string) {
	_load()
	a, err := shimSym(sym)
	if err != nil || a == 0 {
		panic("cppgo/purego: cannot resolve symbol " + sym)
	}
	*p = a
}

// b2i bridges Go bool <-> the C ABI int (0/1) representation of C++ bool.
func b2i(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// pgCString allocates a NUL-terminated C copy of s (caller frees via pgFree).
func pgCString(s string) uintptr {
	p := pgMalloc(uintptr(len(s) + 1))
	d := unsafe.Slice((*byte)(unsafe.Pointer(p)), len(s)+1)
	copy(d, s)
	d[len(s)] = 0
	return p
}

// pgStr copies a malloc'd UTF-8 C string of n bytes returned by a shim, then frees it.
func pgStr(p uintptr, n int64) string {
	if p == 0 || n <= 0 {
		return ""
	}
	s := string(unsafe.Slice((*byte)(unsafe.Pointer(p)), int(n)))
	if pgFree != nil {
		pgFree(p)
	}
	return s
}

// pgStrNoFree copies n UTF-8 bytes of a C string the caller does NOT own (a signal/virtual
// arg the C++ side keeps) — copy, never free.
func pgStrNoFree(p uintptr, n int64) string {
	if p == 0 || n <= 0 {
		return ""
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(p)), int(n)))
}

// pgGoStr reads a NUL-terminated C string (e.g. a virtual's slot-key literal) without freeing.
func pgGoStr(p uintptr) string {
	if p == 0 {
		return ""
	}
	var n int
	for *(*byte)(unsafe.Pointer(p + uintptr(n))) != 0 {
		n++
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(p)), n))
}

// Signal-callback registry: a Go func can't be passed to C, so it's stored by a sequential int64 id
// and the id crosses instead; the (per-shape) NewCallback dispatch looks it up. Slice-indexed (ids
// are dense, allocated here); a Disconnect nils the slot so the closure is collectable. RWMutex:
// deliveries only read.
var (
	cbMu  sync.RWMutex
	cbReg []any
)

func registerCallback(fn any) int64 {
	cbMu.Lock()
	defer cbMu.Unlock()
	cbReg = append(cbReg, fn)
	return int64(len(cbReg))
}

func lookupCallback(id int64) any {
	cbMu.RLock()
	defer cbMu.RUnlock()
	if id < 1 || int(id) > len(cbReg) {
		return nil
	}
	return cbReg[id-1]
}

// unregisterCallback releases the closure after Disconnect — without it, connect/disconnect-heavy
// long-running apps leak every callback ever registered.
func unregisterCallback(id int64) {
	cbMu.Lock()
	defer cbMu.Unlock()
	if id >= 1 && int(id) <= len(cbReg) {
		cbReg[id-1] = nil
	}
}

// Virtual-override registry: per __Go instance (goid) -> (slot id -> Go impl). Slot ids are
// generator-computed FNV hashes of the virtual's slot key — identical across classes and packages,
// so a registration through a promoted base wrapper still matches the derived subclass's dispatch.
// The inner list is a tiny linear scan (an instance typically overrides a handful of virtuals): no
// per-delivery hashing, no allocation.
type virtEnt struct {
	slot int64
	fn   any
}

// enableEnt flips one bit of the instance's C++-side override mask (__goOn) at registration time, so
// an UNREGISTERED virtual never crosses into Go — the C++ override falls straight through to the base.
type enableEnt struct {
	h      uintptr
	idxOf  func(slot int64) int32
	enable func(h uintptr, idx int32)
}

var (
	vMu     sync.RWMutex
	vReg    [][]virtEnt // index = goid-1; ids are never reused (an empty slot per dead instance, not a leak class)
	vEnable []enableEnt // index = goid-1
)

func nextVirtID() int64 {
	vMu.Lock()
	defer vMu.Unlock()
	vReg = append(vReg, nil)
	vEnable = append(vEnable, enableEnt{})
	return int64(len(vReg))
}

// registerInstance records how to enable a freshly constructed __Go instance's override-mask bits.
func registerInstance(id int64, h uintptr, idxOf func(int64) int32, enable func(uintptr, int32)) {
	vMu.Lock()
	defer vMu.Unlock()
	if id >= 1 && int(id) <= len(vEnable) {
		vEnable[id-1] = enableEnt{h: h, idxOf: idxOf, enable: enable}
	}
}

func registerVirt(id int64, slot int64, fn any) {
	if id < 1 { // not a Go-constructed instance (borrowed wrapper): nothing to override
		return
	}
	vMu.Lock()
	if int(id) > len(vReg) {
		vMu.Unlock()
		return
	}
	ents := vReg[id-1]
	replaced := false
	for i := range ents {
		if ents[i].slot == slot {
			ents[i].fn = fn
			replaced = true
			break
		}
	}
	if !replaced {
		vReg[id-1] = append(ents, virtEnt{slot: slot, fn: fn})
	}
	e := vEnable[id-1]
	vMu.Unlock()
	if e.enable != nil && e.idxOf != nil {
		if idx := e.idxOf(slot); idx >= 0 {
			e.enable(e.h, idx) // flip the C++ mask bit: this virtual now crosses into Go
		}
	}
}

func lookupVirt(id int64, slot int64) any {
	vMu.RLock()
	defer vMu.RUnlock()
	if id < 1 || int(id) > len(vReg) {
		return nil
	}
	for _, e := range vReg[id-1] {
		if e.slot == slot {
			return e.fn
		}
	}
	return nil
}

// unregisterInstance drops an instance's overrides + enable entry (called from Delete).
func unregisterInstance(id int64) {
	vMu.Lock()
	defer vMu.Unlock()
	if id >= 1 && int(id) <= len(vReg) {
		vReg[id-1] = nil
		vEnable[id-1] = enableEnt{}
	}
}
