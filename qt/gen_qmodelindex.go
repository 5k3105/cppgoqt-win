package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QModelIndex_new uintptr
	addr_QModelIndex_new2 uintptr
	addr_QModelIndex_row uintptr
	addr_QModelIndex_column uintptr
	addr_QModelIndex_internalId uintptr
	addr_QModelIndex_parent uintptr
	addr_QModelIndex_sibling uintptr
	addr_QModelIndex_siblingAtColumn uintptr
	addr_QModelIndex_siblingAtRow uintptr
	addr_QModelIndex_data uintptr
	addr_QModelIndex_multiData uintptr
	addr_QModelIndex_flags uintptr
	addr_QModelIndex_model uintptr
	addr_QModelIndex_isValid uintptr
	addr_QModelIndex_OperatorAssign uintptr
	addr_QModelIndex_data2 uintptr
	addr_QModelIndex_OperatorEqual uintptr
	addr_QModelIndex_OperatorNotEqual uintptr
	addr_QModelIndex_OperatorLesser uintptr
	addr_QModelIndex_OperatorGreater uintptr
	addr_QModelIndex_OperatorLesserOrEqual uintptr
	addr_QModelIndex_OperatorGreaterOrEqual uintptr
	addr_QModelIndex_OperatorEqual2 uintptr
	addr_QModelIndex_OperatorNotEqual2 uintptr
	addr_QModelIndex_OperatorLesser2 uintptr
	addr_QModelIndex_OperatorGreater2 uintptr
	addr_QModelIndex_OperatorLesserOrEqual2 uintptr
	addr_QModelIndex_OperatorGreaterOrEqual2 uintptr
	pg_QModelIndex_delete func(uintptr)
)

func init() {
	symaddr(&addr_QModelIndex_new, "QModelIndex_new")
	symaddr(&addr_QModelIndex_new2, "QModelIndex_new2")
	symaddr(&addr_QModelIndex_row, "QModelIndex_row")
	symaddr(&addr_QModelIndex_column, "QModelIndex_column")
	symaddr(&addr_QModelIndex_internalId, "QModelIndex_internalId")
	symaddr(&addr_QModelIndex_parent, "QModelIndex_parent")
	symaddr(&addr_QModelIndex_sibling, "QModelIndex_sibling")
	symaddr(&addr_QModelIndex_siblingAtColumn, "QModelIndex_siblingAtColumn")
	symaddr(&addr_QModelIndex_siblingAtRow, "QModelIndex_siblingAtRow")
	symaddr(&addr_QModelIndex_data, "QModelIndex_data")
	symaddr(&addr_QModelIndex_multiData, "QModelIndex_multiData")
	symaddr(&addr_QModelIndex_flags, "QModelIndex_flags")
	symaddr(&addr_QModelIndex_model, "QModelIndex_model")
	symaddr(&addr_QModelIndex_isValid, "QModelIndex_isValid")
	symaddr(&addr_QModelIndex_OperatorAssign, "QModelIndex_OperatorAssign")
	symaddr(&addr_QModelIndex_data2, "QModelIndex_data2")
	symaddr(&addr_QModelIndex_OperatorEqual, "QModelIndex_OperatorEqual")
	symaddr(&addr_QModelIndex_OperatorNotEqual, "QModelIndex_OperatorNotEqual")
	symaddr(&addr_QModelIndex_OperatorLesser, "QModelIndex_OperatorLesser")
	symaddr(&addr_QModelIndex_OperatorGreater, "QModelIndex_OperatorGreater")
	symaddr(&addr_QModelIndex_OperatorLesserOrEqual, "QModelIndex_OperatorLesserOrEqual")
	symaddr(&addr_QModelIndex_OperatorGreaterOrEqual, "QModelIndex_OperatorGreaterOrEqual")
	symaddr(&addr_QModelIndex_OperatorEqual2, "QModelIndex_OperatorEqual2")
	symaddr(&addr_QModelIndex_OperatorNotEqual2, "QModelIndex_OperatorNotEqual2")
	symaddr(&addr_QModelIndex_OperatorLesser2, "QModelIndex_OperatorLesser2")
	symaddr(&addr_QModelIndex_OperatorGreater2, "QModelIndex_OperatorGreater2")
	symaddr(&addr_QModelIndex_OperatorLesserOrEqual2, "QModelIndex_OperatorLesserOrEqual2")
	symaddr(&addr_QModelIndex_OperatorGreaterOrEqual2, "QModelIndex_OperatorGreaterOrEqual2")
	reg(&pg_QModelIndex_delete, "QModelIndex_delete")
}

type QModelIndex struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QModelIndex) UnsafeGoID() int64     { return this.vid }
func (this *QModelIndex) UnsafeSetGoID(id int64) { this.vid = id }

func bindQModelIndex(h uintptr) *QModelIndex {
	if h == 0 {
		return nil
	}
	return &QModelIndex{h: h}
}

func newownedQModelIndex(h uintptr) *QModelIndex {
	o := bindQModelIndex(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QModelIndex).Delete)
	}
	return o
}

func (this *QModelIndex) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QModelIndex) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQModelIndex(h unsafe.Pointer) *QModelIndex { return bindQModelIndex(uintptr(h)) }
func UnsafeNewQModelIndexOwned(h unsafe.Pointer) *QModelIndex { return newownedQModelIndex(uintptr(h)) }

func (this *QModelIndex) UnsafeRelease() { this.h = 0 }

func NewQModelIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_new)
	__o := newownedQModelIndex(__r)
	return __o
}

func NewQModelIndex2(arg1 *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQModelIndex(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QModelIndex) Row() int {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_row, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QModelIndex) Column() int {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_column, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QModelIndex) InternalId() uintptr {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_internalId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uintptr(__r)
}

func (this *QModelIndex) Parent() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_parent, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QModelIndex) Sibling(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_sibling, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QModelIndex) SiblingAtColumn(column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_siblingAtColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QModelIndex) SiblingAtRow(row int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_siblingAtRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QModelIndex) Data() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_data, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QModelIndex) MultiData(roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QModelIndex_multiData, this.h, uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QModelIndex) Flags() ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemFlag(__r)
}

func (this *QModelIndex) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QModelIndex) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QModelIndex) OperatorAssign(arg1 *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQModelIndex(__r)
}

func (this *QModelIndex) DataWithRole(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QModelIndex_data2, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QModelIndex) OperatorEqual(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorNotEqual(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorLesser(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorGreater(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorLesserOrEqual(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorGreaterOrEqual(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorEqualWithRhs(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorNotEqualWithRhs(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorLesserWithRhs(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorGreaterWithRhs(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorLesserOrEqualWithRhs(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) OperatorGreaterOrEqualWithRhs(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QModelIndex_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QModelIndex) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QModelIndex_delete(this.h)
	}
	this.UnsafeRelease()
}
