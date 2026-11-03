package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPartialOrdering_new uintptr
	addr_QPartialOrdering_new2 uintptr
	addr_QPartialOrdering_new3 uintptr
	addr_QPartialOrdering_new4 uintptr
	addr_QPartialOrdering_ToQtPartialOrdering uintptr
	addr_QPartialOrdering_OperatorEqual uintptr
	addr_QPartialOrdering_OperatorNotEqual uintptr
	addr_QPartialOrdering_OperatorEqual2 uintptr
	addr_QPartialOrdering_OperatorNotEqual2 uintptr
	pg_QPartialOrdering_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPartialOrdering_new, "QPartialOrdering_new")
	symaddr(&addr_QPartialOrdering_new2, "QPartialOrdering_new2")
	symaddr(&addr_QPartialOrdering_new3, "QPartialOrdering_new3")
	symaddr(&addr_QPartialOrdering_new4, "QPartialOrdering_new4")
	symaddr(&addr_QPartialOrdering_ToQtPartialOrdering, "QPartialOrdering_ToQtPartialOrdering")
	symaddr(&addr_QPartialOrdering_OperatorEqual, "QPartialOrdering_OperatorEqual")
	symaddr(&addr_QPartialOrdering_OperatorNotEqual, "QPartialOrdering_OperatorNotEqual")
	symaddr(&addr_QPartialOrdering_OperatorEqual2, "QPartialOrdering_OperatorEqual2")
	symaddr(&addr_QPartialOrdering_OperatorNotEqual2, "QPartialOrdering_OperatorNotEqual2")
	reg(&pg_QPartialOrdering_delete, "QPartialOrdering_delete")
}

type QPartialOrdering struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPartialOrdering) UnsafeGoID() int64     { return this.vid }
func (this *QPartialOrdering) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPartialOrdering(h uintptr) *QPartialOrdering {
	if h == 0 {
		return nil
	}
	return &QPartialOrdering{h: h}
}

func newownedQPartialOrdering(h uintptr) *QPartialOrdering {
	o := bindQPartialOrdering(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPartialOrdering).Delete)
	}
	return o
}

func (this *QPartialOrdering) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPartialOrdering) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPartialOrdering(h unsafe.Pointer) *QPartialOrdering { return bindQPartialOrdering(uintptr(h)) }
func UnsafeNewQPartialOrderingOwned(h unsafe.Pointer) *QPartialOrdering { return newownedQPartialOrdering(uintptr(h)) }

func (this *QPartialOrdering) UnsafeRelease() { this.h = 0 }

func NewQPartialOrdering(order *Qt__partial_ordering) *QPartialOrdering {
	__r, _, _ := purego.SyscallN(addr_QPartialOrdering_new, uintptr(order.UnsafePointer()))
	__o := newownedQPartialOrdering(__r)
	runtime.KeepAlive(order)
	return __o
}

func NewQPartialOrdering2(stdorder *Qt__weak_ordering) *QPartialOrdering {
	__r, _, _ := purego.SyscallN(addr_QPartialOrdering_new2, uintptr(stdorder.UnsafePointer()))
	__o := newownedQPartialOrdering(__r)
	runtime.KeepAlive(stdorder)
	return __o
}

func NewQPartialOrdering3(stdorder *Qt__strong_ordering) *QPartialOrdering {
	__r, _, _ := purego.SyscallN(addr_QPartialOrdering_new3, uintptr(stdorder.UnsafePointer()))
	__o := newownedQPartialOrdering(__r)
	runtime.KeepAlive(stdorder)
	return __o
}

func NewQPartialOrdering4(arg1 *QPartialOrdering) *QPartialOrdering {
	__r, _, _ := purego.SyscallN(addr_QPartialOrdering_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQPartialOrdering(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPartialOrdering) ToQtPartialOrdering() *Qt__partial_ordering {
	__r, _, _ := purego.SyscallN(addr_QPartialOrdering_ToQtPartialOrdering, this.h)
	runtime.KeepAlive(this)
	return newownedQt__partial_ordering(__r)
}

func (this *QPartialOrdering) OperatorEqual(rhs *QPartialOrdering) bool {
	__r0, _, _ := purego.SyscallN(addr_QPartialOrdering_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPartialOrdering) OperatorNotEqual(rhs *QPartialOrdering) bool {
	__r0, _, _ := purego.SyscallN(addr_QPartialOrdering_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPartialOrdering) OperatorEqualWithRhs(rhs *Qt__partial_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_QPartialOrdering_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPartialOrdering) OperatorNotEqualWithRhs(rhs *Qt__partial_ordering) bool {
	__r0, _, _ := purego.SyscallN(addr_QPartialOrdering_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPartialOrdering) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPartialOrdering_delete(this.h)
	}
	this.UnsafeRelease()
}
