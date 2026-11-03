package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextFragment_new uintptr
	addr_QTextFragment_new2 uintptr
	addr_QTextFragment_OperatorAssign uintptr
	addr_QTextFragment_isValid uintptr
	addr_QTextFragment_OperatorEqual uintptr
	addr_QTextFragment_OperatorNotEqual uintptr
	addr_QTextFragment_OperatorLesser uintptr
	addr_QTextFragment_position uintptr
	addr_QTextFragment_length uintptr
	addr_QTextFragment_contains uintptr
	addr_QTextFragment_charFormat uintptr
	addr_QTextFragment_charFormatIndex uintptr
	addr_QTextFragment_text uintptr
	addr_QTextFragment_glyphRuns uintptr
	addr_QTextFragment_glyphRuns2 uintptr
	addr_QTextFragment_glyphRuns3 uintptr
	pg_QTextFragment_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextFragment_new, "QTextFragment_new")
	symaddr(&addr_QTextFragment_new2, "QTextFragment_new2")
	symaddr(&addr_QTextFragment_OperatorAssign, "QTextFragment_OperatorAssign")
	symaddr(&addr_QTextFragment_isValid, "QTextFragment_isValid")
	symaddr(&addr_QTextFragment_OperatorEqual, "QTextFragment_OperatorEqual")
	symaddr(&addr_QTextFragment_OperatorNotEqual, "QTextFragment_OperatorNotEqual")
	symaddr(&addr_QTextFragment_OperatorLesser, "QTextFragment_OperatorLesser")
	symaddr(&addr_QTextFragment_position, "QTextFragment_position")
	symaddr(&addr_QTextFragment_length, "QTextFragment_length")
	symaddr(&addr_QTextFragment_contains, "QTextFragment_contains")
	symaddr(&addr_QTextFragment_charFormat, "QTextFragment_charFormat")
	symaddr(&addr_QTextFragment_charFormatIndex, "QTextFragment_charFormatIndex")
	symaddr(&addr_QTextFragment_text, "QTextFragment_text")
	symaddr(&addr_QTextFragment_glyphRuns, "QTextFragment_glyphRuns")
	symaddr(&addr_QTextFragment_glyphRuns2, "QTextFragment_glyphRuns2")
	symaddr(&addr_QTextFragment_glyphRuns3, "QTextFragment_glyphRuns3")
	reg(&pg_QTextFragment_delete, "QTextFragment_delete")
}

type QTextFragment struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextFragment) UnsafeGoID() int64     { return this.vid }
func (this *QTextFragment) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextFragment(h uintptr) *QTextFragment {
	if h == 0 {
		return nil
	}
	return &QTextFragment{h: h}
}

func newownedQTextFragment(h uintptr) *QTextFragment {
	o := bindQTextFragment(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextFragment).Delete)
	}
	return o
}

func (this *QTextFragment) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextFragment) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextFragment(h unsafe.Pointer) *QTextFragment { return bindQTextFragment(uintptr(h)) }
func UnsafeNewQTextFragmentOwned(h unsafe.Pointer) *QTextFragment { return newownedQTextFragment(uintptr(h)) }

func (this *QTextFragment) UnsafeRelease() { this.h = 0 }

func NewQTextFragment() *QTextFragment {
	__r, _, _ := purego.SyscallN(addr_QTextFragment_new)
	__o := newownedQTextFragment(__r)
	return __o
}

func NewQTextFragment2(o *QTextFragment) *QTextFragment {
	__r, _, _ := purego.SyscallN(addr_QTextFragment_new2, uintptr(o.UnsafePointer()))
	__o := newownedQTextFragment(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QTextFragment) OperatorAssign(o *QTextFragment) *QTextFragment {
	__r, _, _ := purego.SyscallN(addr_QTextFragment_OperatorAssign, this.h, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return bindQTextFragment(__r)
}

func (this *QTextFragment) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFragment) OperatorEqual(o *QTextFragment) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_OperatorEqual, this.h, uintptr(o.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return __r != 0
}

func (this *QTextFragment) OperatorNotEqual(o *QTextFragment) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_OperatorNotEqual, this.h, uintptr(o.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return __r != 0
}

func (this *QTextFragment) OperatorLesser(o *QTextFragment) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_OperatorLesser, this.h, uintptr(o.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return __r != 0
}

func (this *QTextFragment) Position() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_position, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFragment) Length() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_length, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFragment) Contains(position int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_contains, this.h, uintptr(position))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFragment) CharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFragment_charFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextFragment) CharFormatIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFragment_charFormatIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFragment) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFragment_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextFragment) GlyphRuns() []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFragment_glyphRuns, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextFragment) GlyphRunsWithFrom(from int) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFragment_glyphRuns2, this.h, uintptr(from), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextFragment) GlyphRuns2(from int, length int) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFragment_glyphRuns3, this.h, uintptr(from), uintptr(length), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextFragment) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextFragment_delete(this.h)
	}
	this.UnsafeRelease()
}
