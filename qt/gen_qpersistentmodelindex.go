package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPersistentModelIndex_new uintptr
	addr_QPersistentModelIndex_new2 uintptr
	addr_QPersistentModelIndex_new3 uintptr
	addr_QPersistentModelIndex_OperatorAssign uintptr
	addr_QPersistentModelIndex_swap uintptr
	addr_QPersistentModelIndex_OperatorAssign2 uintptr
	addr_QPersistentModelIndex_ToQModelIndex uintptr
	addr_QPersistentModelIndex_row uintptr
	addr_QPersistentModelIndex_column uintptr
	addr_QPersistentModelIndex_internalId uintptr
	addr_QPersistentModelIndex_parent uintptr
	addr_QPersistentModelIndex_sibling uintptr
	addr_QPersistentModelIndex_data uintptr
	addr_QPersistentModelIndex_multiData uintptr
	addr_QPersistentModelIndex_flags uintptr
	addr_QPersistentModelIndex_model uintptr
	addr_QPersistentModelIndex_isValid uintptr
	addr_QPersistentModelIndex_data2 uintptr
	addr_QPersistentModelIndex_OperatorEqual uintptr
	addr_QPersistentModelIndex_OperatorNotEqual uintptr
	addr_QPersistentModelIndex_OperatorLesser uintptr
	addr_QPersistentModelIndex_OperatorGreater uintptr
	addr_QPersistentModelIndex_OperatorLesserOrEqual uintptr
	addr_QPersistentModelIndex_OperatorGreaterOrEqual uintptr
	addr_QPersistentModelIndex_OperatorEqual2 uintptr
	addr_QPersistentModelIndex_OperatorNotEqual2 uintptr
	addr_QPersistentModelIndex_OperatorLesser2 uintptr
	addr_QPersistentModelIndex_OperatorGreater2 uintptr
	addr_QPersistentModelIndex_OperatorLesserOrEqual2 uintptr
	addr_QPersistentModelIndex_OperatorGreaterOrEqual2 uintptr
	pg_QPersistentModelIndex_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPersistentModelIndex_new, "QPersistentModelIndex_new")
	symaddr(&addr_QPersistentModelIndex_new2, "QPersistentModelIndex_new2")
	symaddr(&addr_QPersistentModelIndex_new3, "QPersistentModelIndex_new3")
	symaddr(&addr_QPersistentModelIndex_OperatorAssign, "QPersistentModelIndex_OperatorAssign")
	symaddr(&addr_QPersistentModelIndex_swap, "QPersistentModelIndex_swap")
	symaddr(&addr_QPersistentModelIndex_OperatorAssign2, "QPersistentModelIndex_OperatorAssign2")
	symaddr(&addr_QPersistentModelIndex_ToQModelIndex, "QPersistentModelIndex_ToQModelIndex")
	symaddr(&addr_QPersistentModelIndex_row, "QPersistentModelIndex_row")
	symaddr(&addr_QPersistentModelIndex_column, "QPersistentModelIndex_column")
	symaddr(&addr_QPersistentModelIndex_internalId, "QPersistentModelIndex_internalId")
	symaddr(&addr_QPersistentModelIndex_parent, "QPersistentModelIndex_parent")
	symaddr(&addr_QPersistentModelIndex_sibling, "QPersistentModelIndex_sibling")
	symaddr(&addr_QPersistentModelIndex_data, "QPersistentModelIndex_data")
	symaddr(&addr_QPersistentModelIndex_multiData, "QPersistentModelIndex_multiData")
	symaddr(&addr_QPersistentModelIndex_flags, "QPersistentModelIndex_flags")
	symaddr(&addr_QPersistentModelIndex_model, "QPersistentModelIndex_model")
	symaddr(&addr_QPersistentModelIndex_isValid, "QPersistentModelIndex_isValid")
	symaddr(&addr_QPersistentModelIndex_data2, "QPersistentModelIndex_data2")
	symaddr(&addr_QPersistentModelIndex_OperatorEqual, "QPersistentModelIndex_OperatorEqual")
	symaddr(&addr_QPersistentModelIndex_OperatorNotEqual, "QPersistentModelIndex_OperatorNotEqual")
	symaddr(&addr_QPersistentModelIndex_OperatorLesser, "QPersistentModelIndex_OperatorLesser")
	symaddr(&addr_QPersistentModelIndex_OperatorGreater, "QPersistentModelIndex_OperatorGreater")
	symaddr(&addr_QPersistentModelIndex_OperatorLesserOrEqual, "QPersistentModelIndex_OperatorLesserOrEqual")
	symaddr(&addr_QPersistentModelIndex_OperatorGreaterOrEqual, "QPersistentModelIndex_OperatorGreaterOrEqual")
	symaddr(&addr_QPersistentModelIndex_OperatorEqual2, "QPersistentModelIndex_OperatorEqual2")
	symaddr(&addr_QPersistentModelIndex_OperatorNotEqual2, "QPersistentModelIndex_OperatorNotEqual2")
	symaddr(&addr_QPersistentModelIndex_OperatorLesser2, "QPersistentModelIndex_OperatorLesser2")
	symaddr(&addr_QPersistentModelIndex_OperatorGreater2, "QPersistentModelIndex_OperatorGreater2")
	symaddr(&addr_QPersistentModelIndex_OperatorLesserOrEqual2, "QPersistentModelIndex_OperatorLesserOrEqual2")
	symaddr(&addr_QPersistentModelIndex_OperatorGreaterOrEqual2, "QPersistentModelIndex_OperatorGreaterOrEqual2")
	reg(&pg_QPersistentModelIndex_delete, "QPersistentModelIndex_delete")
}

type QPersistentModelIndex struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPersistentModelIndex) UnsafeGoID() int64     { return this.vid }
func (this *QPersistentModelIndex) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPersistentModelIndex(h uintptr) *QPersistentModelIndex {
	if h == 0 {
		return nil
	}
	return &QPersistentModelIndex{h: h}
}

func newownedQPersistentModelIndex(h uintptr) *QPersistentModelIndex {
	o := bindQPersistentModelIndex(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPersistentModelIndex).Delete)
	}
	return o
}

func (this *QPersistentModelIndex) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPersistentModelIndex) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPersistentModelIndex(h unsafe.Pointer) *QPersistentModelIndex { return bindQPersistentModelIndex(uintptr(h)) }
func UnsafeNewQPersistentModelIndexOwned(h unsafe.Pointer) *QPersistentModelIndex { return newownedQPersistentModelIndex(uintptr(h)) }

func (this *QPersistentModelIndex) UnsafeRelease() { this.h = 0 }

func NewQPersistentModelIndex() *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_new)
	__o := newownedQPersistentModelIndex(__r)
	return __o
}

func NewQPersistentModelIndex2(index *QModelIndex) *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_new2, uintptr(index.UnsafePointer()))
	__o := newownedQPersistentModelIndex(__r)
	runtime.KeepAlive(index)
	return __o
}

func NewQPersistentModelIndex3(other *QPersistentModelIndex) *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_new3, uintptr(other.UnsafePointer()))
	__o := newownedQPersistentModelIndex(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QPersistentModelIndex) OperatorAssign(other *QPersistentModelIndex) *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPersistentModelIndex(__r)
}

func (this *QPersistentModelIndex) Swap(other *QPersistentModelIndex) {
	purego.SyscallN(addr_QPersistentModelIndex_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPersistentModelIndex) OperatorAssignWithOther(other *QModelIndex) *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorAssign2, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPersistentModelIndex(__r)
}

func (this *QPersistentModelIndex) ToQModelIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_ToQModelIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QPersistentModelIndex) Row() int {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_row, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPersistentModelIndex) Column() int {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_column, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPersistentModelIndex) InternalId() uintptr {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_internalId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uintptr(__r)
}

func (this *QPersistentModelIndex) Parent() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_parent, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_sibling, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QPersistentModelIndex) Data() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_data, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPersistentModelIndex) MultiData(roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QPersistentModelIndex_multiData, this.h, uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QPersistentModelIndex) Flags() ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemFlag(__r)
}

func (this *QPersistentModelIndex) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QPersistentModelIndex) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPersistentModelIndex) DataWithRole(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPersistentModelIndex_data2, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPersistentModelIndex) OperatorEqual(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorNotEqual(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorLesser(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorGreater(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorLesserOrEqual(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorGreaterOrEqual(rhs *QPersistentModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorEqualWithRhs(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorNotEqualWithRhs(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorLesserWithRhs(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorGreaterWithRhs(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorLesserOrEqualWithRhs(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) OperatorGreaterOrEqualWithRhs(rhs *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QPersistentModelIndex_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPersistentModelIndex) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPersistentModelIndex_delete(this.h)
	}
	this.UnsafeRelease()
}
