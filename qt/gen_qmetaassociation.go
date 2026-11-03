package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QMetaAssociation_base func(uintptr) uintptr
	addr_QMetaAssociation_new uintptr
	addr_QMetaAssociation_new2 uintptr
	addr_QMetaAssociation_keyMetaType uintptr
	addr_QMetaAssociation_mappedMetaType uintptr
	addr_QMetaAssociation_canInsertKey uintptr
	addr_QMetaAssociation_canRemoveKey uintptr
	addr_QMetaAssociation_canContainsKey uintptr
	addr_QMetaAssociation_canGetMappedAtKey uintptr
	addr_QMetaAssociation_canSetMappedAtKey uintptr
	addr_QMetaAssociation_canGetKeyAtIterator uintptr
	addr_QMetaAssociation_canGetKeyAtConstIterator uintptr
	addr_QMetaAssociation_canGetMappedAtIterator uintptr
	addr_QMetaAssociation_canGetMappedAtConstIterator uintptr
	addr_QMetaAssociation_canSetMappedAtIterator uintptr
	addr_QMetaAssociation_canCreateIteratorAtKey uintptr
	addr_QMetaAssociation_canCreateConstIteratorAtKey uintptr
	addr_QMetaAssociation_OperatorEqual uintptr
	addr_QMetaAssociation_OperatorNotEqual uintptr
	pg_QMetaAssociation_delete func(uintptr)
)

func init() {
	reg(&pg_QMetaAssociation_base, "QMetaAssociation_base")
	symaddr(&addr_QMetaAssociation_new, "QMetaAssociation_new")
	symaddr(&addr_QMetaAssociation_new2, "QMetaAssociation_new2")
	symaddr(&addr_QMetaAssociation_keyMetaType, "QMetaAssociation_keyMetaType")
	symaddr(&addr_QMetaAssociation_mappedMetaType, "QMetaAssociation_mappedMetaType")
	symaddr(&addr_QMetaAssociation_canInsertKey, "QMetaAssociation_canInsertKey")
	symaddr(&addr_QMetaAssociation_canRemoveKey, "QMetaAssociation_canRemoveKey")
	symaddr(&addr_QMetaAssociation_canContainsKey, "QMetaAssociation_canContainsKey")
	symaddr(&addr_QMetaAssociation_canGetMappedAtKey, "QMetaAssociation_canGetMappedAtKey")
	symaddr(&addr_QMetaAssociation_canSetMappedAtKey, "QMetaAssociation_canSetMappedAtKey")
	symaddr(&addr_QMetaAssociation_canGetKeyAtIterator, "QMetaAssociation_canGetKeyAtIterator")
	symaddr(&addr_QMetaAssociation_canGetKeyAtConstIterator, "QMetaAssociation_canGetKeyAtConstIterator")
	symaddr(&addr_QMetaAssociation_canGetMappedAtIterator, "QMetaAssociation_canGetMappedAtIterator")
	symaddr(&addr_QMetaAssociation_canGetMappedAtConstIterator, "QMetaAssociation_canGetMappedAtConstIterator")
	symaddr(&addr_QMetaAssociation_canSetMappedAtIterator, "QMetaAssociation_canSetMappedAtIterator")
	symaddr(&addr_QMetaAssociation_canCreateIteratorAtKey, "QMetaAssociation_canCreateIteratorAtKey")
	symaddr(&addr_QMetaAssociation_canCreateConstIteratorAtKey, "QMetaAssociation_canCreateConstIteratorAtKey")
	symaddr(&addr_QMetaAssociation_OperatorEqual, "QMetaAssociation_OperatorEqual")
	symaddr(&addr_QMetaAssociation_OperatorNotEqual, "QMetaAssociation_OperatorNotEqual")
	reg(&pg_QMetaAssociation_delete, "QMetaAssociation_delete")
}

type QMetaAssociation struct {
	h uintptr
	*QMetaContainer
}

func bindQMetaAssociation(h uintptr) *QMetaAssociation {
	if h == 0 {
		return nil
	}
	return &QMetaAssociation{h: h, QMetaContainer: bindQMetaContainer(pg_QMetaAssociation_base(h))}
}

func newownedQMetaAssociation(h uintptr) *QMetaAssociation {
	o := bindQMetaAssociation(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaAssociation).Delete)
	}
	return o
}

func (this *QMetaAssociation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaAssociation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaAssociation(h unsafe.Pointer) *QMetaAssociation { return bindQMetaAssociation(uintptr(h)) }
func UnsafeNewQMetaAssociationOwned(h unsafe.Pointer) *QMetaAssociation { return newownedQMetaAssociation(uintptr(h)) }

func (this *QMetaAssociation) UnsafeRelease() {
	this.h = 0
	if this.QMetaContainer != nil {
		this.QMetaContainer.UnsafeRelease()
	}
}

func NewQMetaAssociation() *QMetaAssociation {
	__r, _, _ := purego.SyscallN(addr_QMetaAssociation_new)
	__o := newownedQMetaAssociation(__r)
	return __o
}

func NewQMetaAssociation2(arg1 *QMetaAssociation) *QMetaAssociation {
	__r, _, _ := purego.SyscallN(addr_QMetaAssociation_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaAssociation(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaAssociation) KeyMetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaAssociation_keyMetaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaAssociation) MappedMetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaAssociation_mappedMetaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaAssociation) CanInsertKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canInsertKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanRemoveKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canRemoveKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanContainsKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canContainsKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanGetMappedAtKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canGetMappedAtKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanSetMappedAtKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canSetMappedAtKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanGetKeyAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canGetKeyAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanGetKeyAtConstIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canGetKeyAtConstIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanGetMappedAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canGetMappedAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanGetMappedAtConstIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canGetMappedAtConstIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanSetMappedAtIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canSetMappedAtIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanCreateIteratorAtKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canCreateIteratorAtKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) CanCreateConstIteratorAtKey() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_canCreateConstIteratorAtKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaAssociation) OperatorEqual(rhs *QMetaAssociation) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaAssociation) OperatorNotEqual(rhs *QMetaAssociation) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaAssociation_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaAssociation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaAssociation_delete(this.h)
	}
	this.UnsafeRelease()
}
