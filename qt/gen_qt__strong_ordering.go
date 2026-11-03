package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_Qt__strong_ordering_new uintptr
	addr_Qt__strong_ordering_ToQtPartialOrdering uintptr
	addr_Qt__strong_ordering_ToQtWeakOrdering uintptr
	addr_Qt__strong_ordering_OperatorEqual uintptr
	addr_Qt__strong_ordering_OperatorNotEqual uintptr
	addr_Qt__strong_ordering_OperatorEqual2 uintptr
	addr_Qt__strong_ordering_OperatorNotEqual2 uintptr
	addr_Qt__strong_ordering_OperatorEqual3 uintptr
	addr_Qt__strong_ordering_OperatorNotEqual3 uintptr
	pg_Qt__strong_ordering_delete func(uintptr)
)

func init() {
	symaddr(&addr_Qt__strong_ordering_new, "Qt__strong_ordering_new")
	symaddr(&addr_Qt__strong_ordering_ToQtPartialOrdering, "Qt__strong_ordering_ToQtPartialOrdering")
	symaddr(&addr_Qt__strong_ordering_ToQtWeakOrdering, "Qt__strong_ordering_ToQtWeakOrdering")
	symaddr(&addr_Qt__strong_ordering_OperatorEqual, "Qt__strong_ordering_OperatorEqual")
	symaddr(&addr_Qt__strong_ordering_OperatorNotEqual, "Qt__strong_ordering_OperatorNotEqual")
	symaddr(&addr_Qt__strong_ordering_OperatorEqual2, "Qt__strong_ordering_OperatorEqual2")
	symaddr(&addr_Qt__strong_ordering_OperatorNotEqual2, "Qt__strong_ordering_OperatorNotEqual2")
	symaddr(&addr_Qt__strong_ordering_OperatorEqual3, "Qt__strong_ordering_OperatorEqual3")
	symaddr(&addr_Qt__strong_ordering_OperatorNotEqual3, "Qt__strong_ordering_OperatorNotEqual3")
	reg(&pg_Qt__strong_ordering_delete, "Qt__strong_ordering_delete")
}

type Qt__strong_ordering struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *Qt__strong_ordering) UnsafeGoID() int64     { return this.vid }
func (this *Qt__strong_ordering) UnsafeSetGoID(id int64) { this.vid = id }

func bindQt__strong_ordering(h uintptr) *Qt__strong_ordering {
	if h == 0 {
		return nil
	}
	return &Qt__strong_ordering{h: h}
}

func newownedQt__strong_ordering(h uintptr) *Qt__strong_ordering {
	o := bindQt__strong_ordering(h)
	if o != nil {
		runtime.SetFinalizer(o, (*Qt__strong_ordering).Delete)
	}
	return o
}

func (this *Qt__strong_ordering) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *Qt__strong_ordering) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQt__strong_ordering(h unsafe.Pointer) *Qt__strong_ordering { return bindQt__strong_ordering(uintptr(h)) }
func UnsafeNewQt__strong_orderingOwned(h unsafe.Pointer) *Qt__strong_ordering { return newownedQt__strong_ordering(uintptr(h)) }

func (this *Qt__strong_ordering) UnsafeRelease() { this.h = 0 }

func NewQt__strong_ordering(arg1 *Qt__strong_ordering) *Qt__strong_ordering {
	__r, _, _ := purego.SyscallN(addr_Qt__strong_ordering_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQt__strong_ordering(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *Qt__strong_ordering) ToQtPartialOrdering() *Qt__partial_ordering {
	__r, _, _ := purego.SyscallN(addr_Qt__strong_ordering_ToQtPartialOrdering, this.h)
	runtime.KeepAlive(this)
	return newownedQt__partial_ordering(__r)
}

func (this *Qt__strong_ordering) ToQtWeakOrdering() *Qt__weak_ordering {
	__r, _, _ := purego.SyscallN(addr_Qt__strong_ordering_ToQtWeakOrdering, this.h)
	runtime.KeepAlive(this)
	return newownedQt__weak_ordering(__r)
}

func (this *Qt__strong_ordering) OperatorEqual(rhs *Qt__strong_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_Qt__strong_ordering_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *Qt__strong_ordering) OperatorNotEqual(rhs *Qt__strong_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_Qt__strong_ordering_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *Qt__strong_ordering) OperatorEqualWithRhs(rhs *Qt__partial_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_Qt__strong_ordering_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *Qt__strong_ordering) OperatorNotEqualWithRhs(rhs *Qt__partial_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_Qt__strong_ordering_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *Qt__strong_ordering) OperatorEqual2(rhs *Qt__weak_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_Qt__strong_ordering_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *Qt__strong_ordering) OperatorNotEqual2(rhs *Qt__weak_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_Qt__strong_ordering_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *Qt__strong_ordering) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_Qt__strong_ordering_delete(this.h)
	}
	this.UnsafeRelease()
}
