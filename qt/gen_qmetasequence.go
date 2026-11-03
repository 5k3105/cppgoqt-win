package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QMetaSequence_base func(uintptr) uintptr
	addr_QMetaSequence_new uintptr
	addr_QMetaSequence_new2 uintptr
	addr_QMetaSequence_valueMetaType uintptr
	addr_QMetaSequence_isSortable uintptr
	addr_QMetaSequence_canAddValueAtBegin uintptr
	addr_QMetaSequence_canAddValueAtEnd uintptr
	addr_QMetaSequence_canRemoveValueAtBegin uintptr
	addr_QMetaSequence_canRemoveValueAtEnd uintptr
	addr_QMetaSequence_canGetValueAtIndex uintptr
	addr_QMetaSequence_canSetValueAtIndex uintptr
	addr_QMetaSequence_canAddValue uintptr
	addr_QMetaSequence_canRemoveValue uintptr
	addr_QMetaSequence_canGetValueAtIterator uintptr
	addr_QMetaSequence_canSetValueAtIterator uintptr
	addr_QMetaSequence_canInsertValueAtIterator uintptr
	addr_QMetaSequence_canEraseValueAtIterator uintptr
	addr_QMetaSequence_canEraseRangeAtIterator uintptr
	addr_QMetaSequence_canGetValueAtConstIterator uintptr
	addr_QMetaSequence_OperatorEqual uintptr
	addr_QMetaSequence_OperatorNotEqual uintptr
	pg_QMetaSequence_delete func(uintptr)
)

func init() {
	reg(&pg_QMetaSequence_base, "QMetaSequence_base")
	symaddr(&addr_QMetaSequence_new, "QMetaSequence_new")
	symaddr(&addr_QMetaSequence_new2, "QMetaSequence_new2")
	symaddr(&addr_QMetaSequence_valueMetaType, "QMetaSequence_valueMetaType")
	symaddr(&addr_QMetaSequence_isSortable, "QMetaSequence_isSortable")
	symaddr(&addr_QMetaSequence_canAddValueAtBegin, "QMetaSequence_canAddValueAtBegin")
	symaddr(&addr_QMetaSequence_canAddValueAtEnd, "QMetaSequence_canAddValueAtEnd")
	symaddr(&addr_QMetaSequence_canRemoveValueAtBegin, "QMetaSequence_canRemoveValueAtBegin")
	symaddr(&addr_QMetaSequence_canRemoveValueAtEnd, "QMetaSequence_canRemoveValueAtEnd")
	symaddr(&addr_QMetaSequence_canGetValueAtIndex, "QMetaSequence_canGetValueAtIndex")
	symaddr(&addr_QMetaSequence_canSetValueAtIndex, "QMetaSequence_canSetValueAtIndex")
	symaddr(&addr_QMetaSequence_canAddValue, "QMetaSequence_canAddValue")
	symaddr(&addr_QMetaSequence_canRemoveValue, "QMetaSequence_canRemoveValue")
	symaddr(&addr_QMetaSequence_canGetValueAtIterator, "QMetaSequence_canGetValueAtIterator")
	symaddr(&addr_QMetaSequence_canSetValueAtIterator, "QMetaSequence_canSetValueAtIterator")
	symaddr(&addr_QMetaSequence_canInsertValueAtIterator, "QMetaSequence_canInsertValueAtIterator")
	symaddr(&addr_QMetaSequence_canEraseValueAtIterator, "QMetaSequence_canEraseValueAtIterator")
	symaddr(&addr_QMetaSequence_canEraseRangeAtIterator, "QMetaSequence_canEraseRangeAtIterator")
	symaddr(&addr_QMetaSequence_canGetValueAtConstIterator, "QMetaSequence_canGetValueAtConstIterator")
	symaddr(&addr_QMetaSequence_OperatorEqual, "QMetaSequence_OperatorEqual")
	symaddr(&addr_QMetaSequence_OperatorNotEqual, "QMetaSequence_OperatorNotEqual")
	reg(&pg_QMetaSequence_delete, "QMetaSequence_delete")
}

type QMetaSequence struct {
	h uintptr
	*QMetaContainer
}

func bindQMetaSequence(h uintptr) *QMetaSequence {
	if h == 0 {
		return nil
	}
	return &QMetaSequence{h: h, QMetaContainer: bindQMetaContainer(pg_QMetaSequence_base(h))}
}

func newownedQMetaSequence(h uintptr) *QMetaSequence {
	o := bindQMetaSequence(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaSequence).Delete)
	}
	return o
}

func (this *QMetaSequence) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaSequence) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaSequence(h unsafe.Pointer) *QMetaSequence { return bindQMetaSequence(uintptr(h)) }
func UnsafeNewQMetaSequenceOwned(h unsafe.Pointer) *QMetaSequence { return newownedQMetaSequence(uintptr(h)) }

func (this *QMetaSequence) UnsafeRelease() {
	this.h = 0
	if this.QMetaContainer != nil {
		this.QMetaContainer.UnsafeRelease()
	}
}

func NewQMetaSequence() *QMetaSequence {
	__r, _, _ := purego.SyscallN(addr_QMetaSequence_new)
	__o := newownedQMetaSequence(__r)
	return __o
}

func NewQMetaSequence2(arg1 *QMetaSequence) *QMetaSequence {
	__r, _, _ := purego.SyscallN(addr_QMetaSequence_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaSequence(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaSequence) ValueMetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaSequence_valueMetaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaSequence) IsSortable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_isSortable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanAddValueAtBegin() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canAddValueAtBegin, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanAddValueAtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canAddValueAtEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanRemoveValueAtBegin() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canRemoveValueAtBegin, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanRemoveValueAtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canRemoveValueAtEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanGetValueAtIndex() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canGetValueAtIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanSetValueAtIndex() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canSetValueAtIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanAddValue() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canAddValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanRemoveValue() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canRemoveValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanGetValueAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canGetValueAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanSetValueAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canSetValueAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanInsertValueAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canInsertValueAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanEraseValueAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canEraseValueAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanEraseRangeAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canEraseRangeAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) CanGetValueAtConstIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_canGetValueAtConstIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaSequence) OperatorEqual(rhs *QMetaSequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaSequence) OperatorNotEqual(rhs *QMetaSequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaSequence_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaSequence) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaSequence_delete(this.h)
	}
	this.UnsafeRelease()
}
