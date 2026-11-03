package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPageRanges__Range_new uintptr
	addr_QPageRanges__Range_new2 uintptr
	addr_QPageRanges__Range_contains uintptr
	addr_QPageRanges__Range_OperatorEqual uintptr
	addr_QPageRanges__Range_OperatorNotEqual uintptr
	addr_QPageRanges__Range_OperatorLesser uintptr
	addr_QPageRanges__Range_from_get uintptr
	addr_QPageRanges__Range_from_set uintptr
	addr_QPageRanges__Range_to_get uintptr
	addr_QPageRanges__Range_to_set uintptr
	pg_QPageRanges__Range_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPageRanges__Range_new, "QPageRanges__Range_new")
	symaddr(&addr_QPageRanges__Range_new2, "QPageRanges__Range_new2")
	symaddr(&addr_QPageRanges__Range_contains, "QPageRanges__Range_contains")
	symaddr(&addr_QPageRanges__Range_OperatorEqual, "QPageRanges__Range_OperatorEqual")
	symaddr(&addr_QPageRanges__Range_OperatorNotEqual, "QPageRanges__Range_OperatorNotEqual")
	symaddr(&addr_QPageRanges__Range_OperatorLesser, "QPageRanges__Range_OperatorLesser")
	symaddr(&addr_QPageRanges__Range_from_get, "QPageRanges__Range_from_get")
	symaddr(&addr_QPageRanges__Range_from_set, "QPageRanges__Range_from_set")
	symaddr(&addr_QPageRanges__Range_to_get, "QPageRanges__Range_to_get")
	symaddr(&addr_QPageRanges__Range_to_set, "QPageRanges__Range_to_set")
	reg(&pg_QPageRanges__Range_delete, "QPageRanges__Range_delete")
}

type QPageRanges__Range struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPageRanges__Range) UnsafeGoID() int64     { return this.vid }
func (this *QPageRanges__Range) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPageRanges__Range(h uintptr) *QPageRanges__Range {
	if h == 0 {
		return nil
	}
	return &QPageRanges__Range{h: h}
}

func newownedQPageRanges__Range(h uintptr) *QPageRanges__Range {
	o := bindQPageRanges__Range(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPageRanges__Range).Delete)
	}
	return o
}

func (this *QPageRanges__Range) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPageRanges__Range) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPageRanges__Range(h unsafe.Pointer) *QPageRanges__Range { return bindQPageRanges__Range(uintptr(h)) }
func UnsafeNewQPageRanges__RangeOwned(h unsafe.Pointer) *QPageRanges__Range { return newownedQPageRanges__Range(uintptr(h)) }

func (this *QPageRanges__Range) UnsafeRelease() { this.h = 0 }

func NewQPageRanges__Range() *QPageRanges__Range {
	__r, _, _ := purego.SyscallN(addr_QPageRanges__Range_new)
	__o := newownedQPageRanges__Range(__r)
	return __o
}

func NewQPageRanges__Range2(arg1 *QPageRanges__Range) *QPageRanges__Range {
	__r, _, _ := purego.SyscallN(addr_QPageRanges__Range_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPageRanges__Range(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPageRanges__Range) Contains(pageNumber int) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges__Range_contains, this.h, uintptr(pageNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageRanges__Range) OperatorEqual(rhs *QPageRanges__Range) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges__Range_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageRanges__Range) OperatorNotEqual(rhs *QPageRanges__Range) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges__Range_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageRanges__Range) OperatorLesser(rhs *QPageRanges__Range) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges__Range_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageRanges__Range) From() int {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges__Range_from_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPageRanges__Range) SetFrom(from int) {
	purego.SyscallN(addr_QPageRanges__Range_from_set, this.h, uintptr(from))
	runtime.KeepAlive(this)
}

func (this *QPageRanges__Range) To() int {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges__Range_to_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPageRanges__Range) SetTo(to int) {
	purego.SyscallN(addr_QPageRanges__Range_to_set, this.h, uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QPageRanges__Range) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPageRanges__Range_delete(this.h)
	}
	this.UnsafeRelease()
}
