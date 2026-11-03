package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaProperty_new uintptr
	addr_QMetaProperty_name uintptr
	addr_QMetaProperty_typeName uintptr
	addr_QMetaProperty_type uintptr
	addr_QMetaProperty_userType uintptr
	addr_QMetaProperty_typeId uintptr
	addr_QMetaProperty_metaType uintptr
	addr_QMetaProperty_propertyIndex uintptr
	addr_QMetaProperty_relativePropertyIndex uintptr
	addr_QMetaProperty_isReadable uintptr
	addr_QMetaProperty_isWritable uintptr
	addr_QMetaProperty_isResettable uintptr
	addr_QMetaProperty_isDesignable uintptr
	addr_QMetaProperty_isScriptable uintptr
	addr_QMetaProperty_isStored uintptr
	addr_QMetaProperty_isUser uintptr
	addr_QMetaProperty_isConstant uintptr
	addr_QMetaProperty_isFinal uintptr
	addr_QMetaProperty_isVirtual uintptr
	addr_QMetaProperty_isOverride uintptr
	addr_QMetaProperty_isRequired uintptr
	addr_QMetaProperty_isBindable uintptr
	addr_QMetaProperty_isFlagType uintptr
	addr_QMetaProperty_isEnumType uintptr
	addr_QMetaProperty_enumerator uintptr
	addr_QMetaProperty_hasNotifySignal uintptr
	addr_QMetaProperty_notifySignal uintptr
	addr_QMetaProperty_notifySignalIndex uintptr
	addr_QMetaProperty_revision uintptr
	addr_QMetaProperty_read uintptr
	addr_QMetaProperty_write uintptr
	addr_QMetaProperty_reset uintptr
	addr_QMetaProperty_bindable uintptr
	addr_QMetaProperty_hasStdCppSet uintptr
	addr_QMetaProperty_isAlias uintptr
	addr_QMetaProperty_isValid uintptr
	addr_QMetaProperty_enclosingMetaObject uintptr
	pg_QMetaProperty_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaProperty_new, "QMetaProperty_new")
	symaddr(&addr_QMetaProperty_name, "QMetaProperty_name")
	symaddr(&addr_QMetaProperty_typeName, "QMetaProperty_typeName")
	symaddr(&addr_QMetaProperty_type, "QMetaProperty_type")
	symaddr(&addr_QMetaProperty_userType, "QMetaProperty_userType")
	symaddr(&addr_QMetaProperty_typeId, "QMetaProperty_typeId")
	symaddr(&addr_QMetaProperty_metaType, "QMetaProperty_metaType")
	symaddr(&addr_QMetaProperty_propertyIndex, "QMetaProperty_propertyIndex")
	symaddr(&addr_QMetaProperty_relativePropertyIndex, "QMetaProperty_relativePropertyIndex")
	symaddr(&addr_QMetaProperty_isReadable, "QMetaProperty_isReadable")
	symaddr(&addr_QMetaProperty_isWritable, "QMetaProperty_isWritable")
	symaddr(&addr_QMetaProperty_isResettable, "QMetaProperty_isResettable")
	symaddr(&addr_QMetaProperty_isDesignable, "QMetaProperty_isDesignable")
	symaddr(&addr_QMetaProperty_isScriptable, "QMetaProperty_isScriptable")
	symaddr(&addr_QMetaProperty_isStored, "QMetaProperty_isStored")
	symaddr(&addr_QMetaProperty_isUser, "QMetaProperty_isUser")
	symaddr(&addr_QMetaProperty_isConstant, "QMetaProperty_isConstant")
	symaddr(&addr_QMetaProperty_isFinal, "QMetaProperty_isFinal")
	symaddr(&addr_QMetaProperty_isVirtual, "QMetaProperty_isVirtual")
	symaddr(&addr_QMetaProperty_isOverride, "QMetaProperty_isOverride")
	symaddr(&addr_QMetaProperty_isRequired, "QMetaProperty_isRequired")
	symaddr(&addr_QMetaProperty_isBindable, "QMetaProperty_isBindable")
	symaddr(&addr_QMetaProperty_isFlagType, "QMetaProperty_isFlagType")
	symaddr(&addr_QMetaProperty_isEnumType, "QMetaProperty_isEnumType")
	symaddr(&addr_QMetaProperty_enumerator, "QMetaProperty_enumerator")
	symaddr(&addr_QMetaProperty_hasNotifySignal, "QMetaProperty_hasNotifySignal")
	symaddr(&addr_QMetaProperty_notifySignal, "QMetaProperty_notifySignal")
	symaddr(&addr_QMetaProperty_notifySignalIndex, "QMetaProperty_notifySignalIndex")
	symaddr(&addr_QMetaProperty_revision, "QMetaProperty_revision")
	symaddr(&addr_QMetaProperty_read, "QMetaProperty_read")
	symaddr(&addr_QMetaProperty_write, "QMetaProperty_write")
	symaddr(&addr_QMetaProperty_reset, "QMetaProperty_reset")
	symaddr(&addr_QMetaProperty_bindable, "QMetaProperty_bindable")
	symaddr(&addr_QMetaProperty_hasStdCppSet, "QMetaProperty_hasStdCppSet")
	symaddr(&addr_QMetaProperty_isAlias, "QMetaProperty_isAlias")
	symaddr(&addr_QMetaProperty_isValid, "QMetaProperty_isValid")
	symaddr(&addr_QMetaProperty_enclosingMetaObject, "QMetaProperty_enclosingMetaObject")
	reg(&pg_QMetaProperty_delete, "QMetaProperty_delete")
}

type QMetaProperty struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaProperty) UnsafeGoID() int64     { return this.vid }
func (this *QMetaProperty) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaProperty(h uintptr) *QMetaProperty {
	if h == 0 {
		return nil
	}
	return &QMetaProperty{h: h}
}

func newownedQMetaProperty(h uintptr) *QMetaProperty {
	o := bindQMetaProperty(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaProperty).Delete)
	}
	return o
}

func (this *QMetaProperty) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaProperty) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaProperty(h unsafe.Pointer) *QMetaProperty { return bindQMetaProperty(uintptr(h)) }
func UnsafeNewQMetaPropertyOwned(h unsafe.Pointer) *QMetaProperty { return newownedQMetaProperty(uintptr(h)) }

func (this *QMetaProperty) UnsafeRelease() { this.h = 0 }

func NewQMetaProperty() *QMetaProperty {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_new)
	__o := newownedQMetaProperty(__r)
	return __o
}

func (this *QMetaProperty) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaProperty) TypeName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_typeName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaProperty) Type() QVariant__Type {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QVariant__Type(__r)
}

func (this *QMetaProperty) UserType() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_userType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaProperty) TypeId() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_typeId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaProperty) MetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_metaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaProperty) PropertyIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_propertyIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaProperty) RelativePropertyIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_relativePropertyIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaProperty) IsReadable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isReadable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsWritable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isWritable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsResettable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isResettable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsDesignable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isDesignable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsScriptable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isScriptable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsStored() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isStored, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsUser() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isUser, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsConstant() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isConstant, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsFinal() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isFinal, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsVirtual() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isVirtual, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsOverride() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isOverride, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsRequired() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isRequired, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsBindable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isBindable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsFlagType() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isFlagType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsEnumType() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isEnumType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) Enumerator() *QMetaEnum {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_enumerator, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaEnum(__r)
}

func (this *QMetaProperty) HasNotifySignal() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_hasNotifySignal, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) NotifySignal() *QMetaMethod {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_notifySignal, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaMethod(__r)
}

func (this *QMetaProperty) NotifySignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_notifySignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaProperty) Revision() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_revision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaProperty) Read(obj *QObject) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_read, this.h, uintptr(obj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
	return newownedQVariant(__r)
}

func (this *QMetaProperty) Write(obj *QObject, value *QVariant) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_write, this.h, uintptr(obj.UnsafePointer()), uintptr(value.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QMetaProperty) Reset(obj *QObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_reset, this.h, uintptr(obj.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
	return __r != 0
}

func (this *QMetaProperty) Bindable(object *QObject) *QUntypedBindable {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_bindable, this.h, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	return newownedQUntypedBindable(__r)
}

func (this *QMetaProperty) HasStdCppSet() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_hasStdCppSet, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsAlias() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isAlias, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaProperty_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaProperty) EnclosingMetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaProperty_enclosingMetaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaProperty) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaProperty_delete(this.h)
	}
	this.UnsafeRelease()
}
