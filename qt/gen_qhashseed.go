package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QHashSeed_new uintptr
	addr_QHashSeed_new2 uintptr
	addr_QHashSeed_ToSizeT uintptr
	addr_QHashSeed_globalSeed uintptr
	addr_QHashSeed_setDeterministicGlobalSeed uintptr
	addr_QHashSeed_resetRandomGlobalSeed uintptr
	pg_QHashSeed_delete func(uintptr)
)

func init() {
	symaddr(&addr_QHashSeed_new, "QHashSeed_new")
	symaddr(&addr_QHashSeed_new2, "QHashSeed_new2")
	symaddr(&addr_QHashSeed_ToSizeT, "QHashSeed_ToSizeT")
	symaddr(&addr_QHashSeed_globalSeed, "QHashSeed_globalSeed")
	symaddr(&addr_QHashSeed_setDeterministicGlobalSeed, "QHashSeed_setDeterministicGlobalSeed")
	symaddr(&addr_QHashSeed_resetRandomGlobalSeed, "QHashSeed_resetRandomGlobalSeed")
	reg(&pg_QHashSeed_delete, "QHashSeed_delete")
}

type QHashSeed struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QHashSeed) UnsafeGoID() int64     { return this.vid }
func (this *QHashSeed) UnsafeSetGoID(id int64) { this.vid = id }

func bindQHashSeed(h uintptr) *QHashSeed {
	if h == 0 {
		return nil
	}
	return &QHashSeed{h: h}
}

func newownedQHashSeed(h uintptr) *QHashSeed {
	o := bindQHashSeed(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QHashSeed).Delete)
	}
	return o
}

func (this *QHashSeed) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QHashSeed) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQHashSeed(h unsafe.Pointer) *QHashSeed { return bindQHashSeed(uintptr(h)) }
func UnsafeNewQHashSeedOwned(h unsafe.Pointer) *QHashSeed { return newownedQHashSeed(uintptr(h)) }

func (this *QHashSeed) UnsafeRelease() { this.h = 0 }

func NewQHashSeed() *QHashSeed {
	__r, _, _ := purego.SyscallN(addr_QHashSeed_new)
	__o := newownedQHashSeed(__r)
	return __o
}

func NewQHashSeed2(d uint64) *QHashSeed {
	__r, _, _ := purego.SyscallN(addr_QHashSeed_new2, uintptr(d))
	__o := newownedQHashSeed(__r)
	return __o
}

func (this *QHashSeed) ToSizeT() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QHashSeed_ToSizeT, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func QHashSeed_GlobalSeed() *QHashSeed {
	__r, _, _ := purego.SyscallN(addr_QHashSeed_globalSeed)
	return newownedQHashSeed(__r)
}

func QHashSeed_SetDeterministicGlobalSeed() {
	purego.SyscallN(addr_QHashSeed_setDeterministicGlobalSeed)
}

func QHashSeed_ResetRandomGlobalSeed() {
	purego.SyscallN(addr_QHashSeed_resetRandomGlobalSeed)
}

func (this *QHashSeed) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QHashSeed_delete(this.h)
	}
	this.UnsafeRelease()
}
