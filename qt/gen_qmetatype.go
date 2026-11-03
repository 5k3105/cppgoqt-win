package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaType_new uintptr
	addr_QMetaType_new2 uintptr
	addr_QMetaType_new3 uintptr
	addr_QMetaType_registerNormalizedTypedef uintptr
	addr_QMetaType_type uintptr
	addr_QMetaType_type2 uintptr
	addr_QMetaType_typeName uintptr
	addr_QMetaType_sizeOf uintptr
	addr_QMetaType_typeFlags uintptr
	addr_QMetaType_metaObjectForType uintptr
	addr_QMetaType_isRegistered uintptr
	addr_QMetaType_isValid uintptr
	addr_QMetaType_isRegistered2 uintptr
	addr_QMetaType_registerType uintptr
	addr_QMetaType_id uintptr
	addr_QMetaType_sizeOf2 uintptr
	addr_QMetaType_alignOf uintptr
	addr_QMetaType_flags uintptr
	addr_QMetaType_metaObject uintptr
	addr_QMetaType_name uintptr
	addr_QMetaType_isDefaultConstructible uintptr
	addr_QMetaType_isCopyConstructible uintptr
	addr_QMetaType_isMoveConstructible uintptr
	addr_QMetaType_isDestructible uintptr
	addr_QMetaType_isEqualityComparable uintptr
	addr_QMetaType_isOrdered uintptr
	addr_QMetaType_hasRegisteredDataStreamOperators uintptr
	addr_QMetaType_underlyingType uintptr
	addr_QMetaType_fromName uintptr
	addr_QMetaType_hasRegisteredDebugStreamOperator uintptr
	addr_QMetaType_hasRegisteredDebugStreamOperator2 uintptr
	addr_QMetaType_canConvert uintptr
	addr_QMetaType_canView uintptr
	addr_QMetaType_hasRegisteredConverterFunction uintptr
	addr_QMetaType_hasRegisteredMutableViewFunction uintptr
	addr_QMetaType_unregisterConverterFunction uintptr
	addr_QMetaType_unregisterMutableViewFunction uintptr
	addr_QMetaType_unregisterMetaType uintptr
	addr_QMetaType_id2 uintptr
	addr_QMetaType_OperatorEqual uintptr
	addr_QMetaType_OperatorNotEqual uintptr
	pg_QMetaType_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaType_new, "QMetaType_new")
	symaddr(&addr_QMetaType_new2, "QMetaType_new2")
	symaddr(&addr_QMetaType_new3, "QMetaType_new3")
	symaddr(&addr_QMetaType_registerNormalizedTypedef, "QMetaType_registerNormalizedTypedef")
	symaddr(&addr_QMetaType_type, "QMetaType_type")
	symaddr(&addr_QMetaType_type2, "QMetaType_type2")
	symaddr(&addr_QMetaType_typeName, "QMetaType_typeName")
	symaddr(&addr_QMetaType_sizeOf, "QMetaType_sizeOf")
	symaddr(&addr_QMetaType_typeFlags, "QMetaType_typeFlags")
	symaddr(&addr_QMetaType_metaObjectForType, "QMetaType_metaObjectForType")
	symaddr(&addr_QMetaType_isRegistered, "QMetaType_isRegistered")
	symaddr(&addr_QMetaType_isValid, "QMetaType_isValid")
	symaddr(&addr_QMetaType_isRegistered2, "QMetaType_isRegistered2")
	symaddr(&addr_QMetaType_registerType, "QMetaType_registerType")
	symaddr(&addr_QMetaType_id, "QMetaType_id")
	symaddr(&addr_QMetaType_sizeOf2, "QMetaType_sizeOf2")
	symaddr(&addr_QMetaType_alignOf, "QMetaType_alignOf")
	symaddr(&addr_QMetaType_flags, "QMetaType_flags")
	symaddr(&addr_QMetaType_metaObject, "QMetaType_metaObject")
	symaddr(&addr_QMetaType_name, "QMetaType_name")
	symaddr(&addr_QMetaType_isDefaultConstructible, "QMetaType_isDefaultConstructible")
	symaddr(&addr_QMetaType_isCopyConstructible, "QMetaType_isCopyConstructible")
	symaddr(&addr_QMetaType_isMoveConstructible, "QMetaType_isMoveConstructible")
	symaddr(&addr_QMetaType_isDestructible, "QMetaType_isDestructible")
	symaddr(&addr_QMetaType_isEqualityComparable, "QMetaType_isEqualityComparable")
	symaddr(&addr_QMetaType_isOrdered, "QMetaType_isOrdered")
	symaddr(&addr_QMetaType_hasRegisteredDataStreamOperators, "QMetaType_hasRegisteredDataStreamOperators")
	symaddr(&addr_QMetaType_underlyingType, "QMetaType_underlyingType")
	symaddr(&addr_QMetaType_fromName, "QMetaType_fromName")
	symaddr(&addr_QMetaType_hasRegisteredDebugStreamOperator, "QMetaType_hasRegisteredDebugStreamOperator")
	symaddr(&addr_QMetaType_hasRegisteredDebugStreamOperator2, "QMetaType_hasRegisteredDebugStreamOperator2")
	symaddr(&addr_QMetaType_canConvert, "QMetaType_canConvert")
	symaddr(&addr_QMetaType_canView, "QMetaType_canView")
	symaddr(&addr_QMetaType_hasRegisteredConverterFunction, "QMetaType_hasRegisteredConverterFunction")
	symaddr(&addr_QMetaType_hasRegisteredMutableViewFunction, "QMetaType_hasRegisteredMutableViewFunction")
	symaddr(&addr_QMetaType_unregisterConverterFunction, "QMetaType_unregisterConverterFunction")
	symaddr(&addr_QMetaType_unregisterMutableViewFunction, "QMetaType_unregisterMutableViewFunction")
	symaddr(&addr_QMetaType_unregisterMetaType, "QMetaType_unregisterMetaType")
	symaddr(&addr_QMetaType_id2, "QMetaType_id2")
	symaddr(&addr_QMetaType_OperatorEqual, "QMetaType_OperatorEqual")
	symaddr(&addr_QMetaType_OperatorNotEqual, "QMetaType_OperatorNotEqual")
	reg(&pg_QMetaType_delete, "QMetaType_delete")
}

type QMetaType struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaType) UnsafeGoID() int64     { return this.vid }
func (this *QMetaType) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaType(h uintptr) *QMetaType {
	if h == 0 {
		return nil
	}
	return &QMetaType{h: h}
}

func newownedQMetaType(h uintptr) *QMetaType {
	o := bindQMetaType(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaType).Delete)
	}
	return o
}

func (this *QMetaType) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaType) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaType(h unsafe.Pointer) *QMetaType { return bindQMetaType(uintptr(h)) }
func UnsafeNewQMetaTypeOwned(h unsafe.Pointer) *QMetaType { return newownedQMetaType(uintptr(h)) }

func (this *QMetaType) UnsafeRelease() { this.h = 0 }

func NewQMetaType(type_ int) *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaType_new, uintptr(type_))
	__o := newownedQMetaType(__r)
	return __o
}

func NewQMetaType2() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaType_new2)
	__o := newownedQMetaType(__r)
	return __o
}

func NewQMetaType3(arg1 *QMetaType) *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaType_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaType(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QMetaType_RegisterNormalizedTypedef(normalizedTypeName *QByteArray, type_ *QMetaType) {
	purego.SyscallN(addr_QMetaType_registerNormalizedTypedef, uintptr(normalizedTypeName.UnsafePointer()), uintptr(type_.UnsafePointer()))
	runtime.KeepAlive(normalizedTypeName)
	runtime.KeepAlive(type_)
}

func QMetaType_Type(typeName string) int {
	typeName_b := append([]byte(typeName), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaType_type, uintptr(unsafe.Pointer(&typeName_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(typeName_b)
	return int(__r)
}

func QMetaType_TypeWithTypeName(typeName *QByteArray) int {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_type2, uintptr(typeName.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(typeName)
	return int(__r)
}

func QMetaType_TypeName(type_ int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaType_typeName, uintptr(type_), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QMetaType_SizeOf(type_ int) int {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_sizeOf, uintptr(type_))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QMetaType_TypeFlags(type_ int) QMetaType__TypeFlag {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_typeFlags, uintptr(type_))
	__r := int32(uint32(__r0))
	return QMetaType__TypeFlag(__r)
}

func QMetaType_MetaObjectForType(type_ int) *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaType_metaObjectForType, uintptr(type_))
	return bindQMetaObject(__r)
}

func QMetaType_IsRegistered(type_ int) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isRegistered, uintptr(type_))
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QMetaType) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) IsRegistered() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isRegistered2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) RegisterType() {
	purego.SyscallN(addr_QMetaType_registerType, this.h)
	runtime.KeepAlive(this)
}

func (this *QMetaType) Id() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaType) SizeOf() int64 {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_sizeOf2, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMetaType) AlignOf() int64 {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_alignOf, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMetaType) Flags() QMetaType__TypeFlag {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMetaType__TypeFlag(__r)
}

func (this *QMetaType) MetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaType_metaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaType) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaType_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaType) IsDefaultConstructible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isDefaultConstructible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) IsCopyConstructible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isCopyConstructible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) IsMoveConstructible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isMoveConstructible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) IsDestructible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isDestructible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) IsEqualityComparable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isEqualityComparable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) IsOrdered() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_isOrdered, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) HasRegisteredDataStreamOperators() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_hasRegisteredDataStreamOperators, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaType) UnderlyingType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaType_underlyingType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func QMetaType_FromName(name *QByteArrayView) *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaType_fromName, uintptr(name.UnsafePointer()))
	runtime.KeepAlive(name)
	return newownedQMetaType(__r)
}

func (this *QMetaType) HasRegisteredDebugStreamOperator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_hasRegisteredDebugStreamOperator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMetaType_HasRegisteredDebugStreamOperator(typeId int) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_hasRegisteredDebugStreamOperator2, uintptr(typeId))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QMetaType_CanConvert(fromType *QMetaType, toType *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_canConvert, uintptr(fromType.UnsafePointer()), uintptr(toType.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fromType)
	runtime.KeepAlive(toType)
	return __r != 0
}

func QMetaType_CanView(fromType *QMetaType, toType *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_canView, uintptr(fromType.UnsafePointer()), uintptr(toType.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fromType)
	runtime.KeepAlive(toType)
	return __r != 0
}

func QMetaType_HasRegisteredConverterFunction(fromType *QMetaType, toType *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_hasRegisteredConverterFunction, uintptr(fromType.UnsafePointer()), uintptr(toType.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fromType)
	runtime.KeepAlive(toType)
	return __r != 0
}

func QMetaType_HasRegisteredMutableViewFunction(fromType *QMetaType, toType *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_hasRegisteredMutableViewFunction, uintptr(fromType.UnsafePointer()), uintptr(toType.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fromType)
	runtime.KeepAlive(toType)
	return __r != 0
}

func QMetaType_UnregisterConverterFunction(from *QMetaType, to *QMetaType) {
	purego.SyscallN(addr_QMetaType_unregisterConverterFunction, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func QMetaType_UnregisterMutableViewFunction(from *QMetaType, to *QMetaType) {
	purego.SyscallN(addr_QMetaType_unregisterMutableViewFunction, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func QMetaType_UnregisterMetaType(type_ *QMetaType) {
	purego.SyscallN(addr_QMetaType_unregisterMetaType, uintptr(type_.UnsafePointer()))
	runtime.KeepAlive(type_)
}

func (this *QMetaType) IdWithInt(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_id2, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaType) OperatorEqual(rhs *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaType) OperatorNotEqual(rhs *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaType_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaType) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaType_delete(this.h)
	}
	this.UnsafeRelease()
}
