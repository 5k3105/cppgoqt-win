package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QCborValueRef_base func(uintptr) uintptr
	addr_QCborValueRef_new uintptr
	addr_QCborValueRef_OperatorAssign uintptr
	addr_QCborValueRef_OperatorAssign2 uintptr
	addr_QCborValueRef_OperatorSubscript uintptr
	addr_QCborValueRef_OperatorSubscript2 uintptr
	addr_QCborValueRef_OperatorSubscript3 uintptr
	addr_QCborValueRef_ToQCborValue uintptr
	addr_QCborValueRef_type uintptr
	addr_QCborValueRef_isInteger uintptr
	addr_QCborValueRef_isByteArray uintptr
	addr_QCborValueRef_isString uintptr
	addr_QCborValueRef_isArray uintptr
	addr_QCborValueRef_isMap uintptr
	addr_QCborValueRef_isTag uintptr
	addr_QCborValueRef_isFalse uintptr
	addr_QCborValueRef_isTrue uintptr
	addr_QCborValueRef_isBool uintptr
	addr_QCborValueRef_isNull uintptr
	addr_QCborValueRef_isUndefined uintptr
	addr_QCborValueRef_isDouble uintptr
	addr_QCborValueRef_isDateTime uintptr
	addr_QCborValueRef_isUrl uintptr
	addr_QCborValueRef_isRegularExpression uintptr
	addr_QCborValueRef_isUuid uintptr
	addr_QCborValueRef_isInvalid uintptr
	addr_QCborValueRef_isContainer uintptr
	addr_QCborValueRef_isSimpleType uintptr
	addr_QCborValueRef_isSimpleType2 uintptr
	addr_QCborValueRef_toSimpleType uintptr
	addr_QCborValueRef_tag uintptr
	addr_QCborValueRef_taggedValue uintptr
	addr_QCborValueRef_toInteger uintptr
	addr_QCborValueRef_toBool uintptr
	pg_QCborValueRef_toDouble func(uintptr) float64
	addr_QCborValueRef_toByteArray uintptr
	addr_QCborValueRef_toString uintptr
	addr_QCborValueRef_toDateTime uintptr
	addr_QCborValueRef_toUrl uintptr
	addr_QCborValueRef_toRegularExpression uintptr
	addr_QCborValueRef_toUuid uintptr
	addr_QCborValueRef_toArray uintptr
	addr_QCborValueRef_toArray2 uintptr
	addr_QCborValueRef_toMap uintptr
	addr_QCborValueRef_toMap2 uintptr
	addr_QCborValueRef_OperatorSubscript4 uintptr
	addr_QCborValueRef_OperatorSubscript5 uintptr
	addr_QCborValueRef_OperatorSubscript6 uintptr
	addr_QCborValueRef_compare uintptr
	addr_QCborValueRef_toVariant uintptr
	addr_QCborValueRef_toJsonValue uintptr
	addr_QCborValueRef_toCbor uintptr
	addr_QCborValueRef_toCbor2 uintptr
	addr_QCborValueRef_toDiagnosticNotation uintptr
	addr_QCborValueRef_toSimpleType2 uintptr
	addr_QCborValueRef_tag2 uintptr
	addr_QCborValueRef_taggedValue2 uintptr
	addr_QCborValueRef_toInteger2 uintptr
	addr_QCborValueRef_toBool2 uintptr
	pg_QCborValueRef_toDouble2 func(uintptr, float64) float64
	addr_QCborValueRef_toByteArray2 uintptr
	addr_QCborValueRef_toString2 uintptr
	addr_QCborValueRef_toDateTime2 uintptr
	addr_QCborValueRef_toUrl2 uintptr
	addr_QCborValueRef_toRegularExpression2 uintptr
	addr_QCborValueRef_toUuid2 uintptr
	addr_QCborValueRef_toCbor3 uintptr
	addr_QCborValueRef_toCbor4 uintptr
	addr_QCborValueRef_toDiagnosticNotation2 uintptr
	pg_QCborValueRef_delete func(uintptr)
)

func init() {
	reg(&pg_QCborValueRef_base, "QCborValueRef_base")
	symaddr(&addr_QCborValueRef_new, "QCborValueRef_new")
	symaddr(&addr_QCborValueRef_OperatorAssign, "QCborValueRef_OperatorAssign")
	symaddr(&addr_QCborValueRef_OperatorAssign2, "QCborValueRef_OperatorAssign2")
	symaddr(&addr_QCborValueRef_OperatorSubscript, "QCborValueRef_OperatorSubscript")
	symaddr(&addr_QCborValueRef_OperatorSubscript2, "QCborValueRef_OperatorSubscript2")
	symaddr(&addr_QCborValueRef_OperatorSubscript3, "QCborValueRef_OperatorSubscript3")
	symaddr(&addr_QCborValueRef_ToQCborValue, "QCborValueRef_ToQCborValue")
	symaddr(&addr_QCborValueRef_type, "QCborValueRef_type")
	symaddr(&addr_QCborValueRef_isInteger, "QCborValueRef_isInteger")
	symaddr(&addr_QCborValueRef_isByteArray, "QCborValueRef_isByteArray")
	symaddr(&addr_QCborValueRef_isString, "QCborValueRef_isString")
	symaddr(&addr_QCborValueRef_isArray, "QCborValueRef_isArray")
	symaddr(&addr_QCborValueRef_isMap, "QCborValueRef_isMap")
	symaddr(&addr_QCborValueRef_isTag, "QCborValueRef_isTag")
	symaddr(&addr_QCborValueRef_isFalse, "QCborValueRef_isFalse")
	symaddr(&addr_QCborValueRef_isTrue, "QCborValueRef_isTrue")
	symaddr(&addr_QCborValueRef_isBool, "QCborValueRef_isBool")
	symaddr(&addr_QCborValueRef_isNull, "QCborValueRef_isNull")
	symaddr(&addr_QCborValueRef_isUndefined, "QCborValueRef_isUndefined")
	symaddr(&addr_QCborValueRef_isDouble, "QCborValueRef_isDouble")
	symaddr(&addr_QCborValueRef_isDateTime, "QCborValueRef_isDateTime")
	symaddr(&addr_QCborValueRef_isUrl, "QCborValueRef_isUrl")
	symaddr(&addr_QCborValueRef_isRegularExpression, "QCborValueRef_isRegularExpression")
	symaddr(&addr_QCborValueRef_isUuid, "QCborValueRef_isUuid")
	symaddr(&addr_QCborValueRef_isInvalid, "QCborValueRef_isInvalid")
	symaddr(&addr_QCborValueRef_isContainer, "QCborValueRef_isContainer")
	symaddr(&addr_QCborValueRef_isSimpleType, "QCborValueRef_isSimpleType")
	symaddr(&addr_QCborValueRef_isSimpleType2, "QCborValueRef_isSimpleType2")
	symaddr(&addr_QCborValueRef_toSimpleType, "QCborValueRef_toSimpleType")
	symaddr(&addr_QCborValueRef_tag, "QCborValueRef_tag")
	symaddr(&addr_QCborValueRef_taggedValue, "QCborValueRef_taggedValue")
	symaddr(&addr_QCborValueRef_toInteger, "QCborValueRef_toInteger")
	symaddr(&addr_QCborValueRef_toBool, "QCborValueRef_toBool")
	reg(&pg_QCborValueRef_toDouble, "QCborValueRef_toDouble")
	symaddr(&addr_QCborValueRef_toByteArray, "QCborValueRef_toByteArray")
	symaddr(&addr_QCborValueRef_toString, "QCborValueRef_toString")
	symaddr(&addr_QCborValueRef_toDateTime, "QCborValueRef_toDateTime")
	symaddr(&addr_QCborValueRef_toUrl, "QCborValueRef_toUrl")
	symaddr(&addr_QCborValueRef_toRegularExpression, "QCborValueRef_toRegularExpression")
	symaddr(&addr_QCborValueRef_toUuid, "QCborValueRef_toUuid")
	symaddr(&addr_QCborValueRef_toArray, "QCborValueRef_toArray")
	symaddr(&addr_QCborValueRef_toArray2, "QCborValueRef_toArray2")
	symaddr(&addr_QCborValueRef_toMap, "QCborValueRef_toMap")
	symaddr(&addr_QCborValueRef_toMap2, "QCborValueRef_toMap2")
	symaddr(&addr_QCborValueRef_OperatorSubscript4, "QCborValueRef_OperatorSubscript4")
	symaddr(&addr_QCborValueRef_OperatorSubscript5, "QCborValueRef_OperatorSubscript5")
	symaddr(&addr_QCborValueRef_OperatorSubscript6, "QCborValueRef_OperatorSubscript6")
	symaddr(&addr_QCborValueRef_compare, "QCborValueRef_compare")
	symaddr(&addr_QCborValueRef_toVariant, "QCborValueRef_toVariant")
	symaddr(&addr_QCborValueRef_toJsonValue, "QCborValueRef_toJsonValue")
	symaddr(&addr_QCborValueRef_toCbor, "QCborValueRef_toCbor")
	symaddr(&addr_QCborValueRef_toCbor2, "QCborValueRef_toCbor2")
	symaddr(&addr_QCborValueRef_toDiagnosticNotation, "QCborValueRef_toDiagnosticNotation")
	symaddr(&addr_QCborValueRef_toSimpleType2, "QCborValueRef_toSimpleType2")
	symaddr(&addr_QCborValueRef_tag2, "QCborValueRef_tag2")
	symaddr(&addr_QCborValueRef_taggedValue2, "QCborValueRef_taggedValue2")
	symaddr(&addr_QCborValueRef_toInteger2, "QCborValueRef_toInteger2")
	symaddr(&addr_QCborValueRef_toBool2, "QCborValueRef_toBool2")
	reg(&pg_QCborValueRef_toDouble2, "QCborValueRef_toDouble2")
	symaddr(&addr_QCborValueRef_toByteArray2, "QCborValueRef_toByteArray2")
	symaddr(&addr_QCborValueRef_toString2, "QCborValueRef_toString2")
	symaddr(&addr_QCborValueRef_toDateTime2, "QCborValueRef_toDateTime2")
	symaddr(&addr_QCborValueRef_toUrl2, "QCborValueRef_toUrl2")
	symaddr(&addr_QCborValueRef_toRegularExpression2, "QCborValueRef_toRegularExpression2")
	symaddr(&addr_QCborValueRef_toUuid2, "QCborValueRef_toUuid2")
	symaddr(&addr_QCborValueRef_toCbor3, "QCborValueRef_toCbor3")
	symaddr(&addr_QCborValueRef_toCbor4, "QCborValueRef_toCbor4")
	symaddr(&addr_QCborValueRef_toDiagnosticNotation2, "QCborValueRef_toDiagnosticNotation2")
	reg(&pg_QCborValueRef_delete, "QCborValueRef_delete")
}

type QCborValueRef struct {
	h uintptr
	*QCborValueConstRef
}

func bindQCborValueRef(h uintptr) *QCborValueRef {
	if h == 0 {
		return nil
	}
	return &QCborValueRef{h: h, QCborValueConstRef: bindQCborValueConstRef(pg_QCborValueRef_base(h))}
}

func newownedQCborValueRef(h uintptr) *QCborValueRef {
	o := bindQCborValueRef(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborValueRef).Delete)
	}
	return o
}

func (this *QCborValueRef) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborValueRef) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborValueRef(h unsafe.Pointer) *QCborValueRef { return bindQCborValueRef(uintptr(h)) }
func UnsafeNewQCborValueRefOwned(h unsafe.Pointer) *QCborValueRef { return newownedQCborValueRef(uintptr(h)) }

func (this *QCborValueRef) UnsafeRelease() {
	this.h = 0
	if this.QCborValueConstRef != nil {
		this.QCborValueConstRef.UnsafeRelease()
	}
}

func NewQCborValueRef(arg1 *QCborValueRef) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQCborValueRef(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QCborValueRef) OperatorAssign(other *QCborValue) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCborValueRef(__r)
}

func (this *QCborValueRef) OperatorAssignWithOther(other *QCborValueRef) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorAssign2, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCborValueRef(__r)
}

func (this *QCborValueRef) OperatorSubscript(key int64) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorSubscript, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValueRef(__r)
}

func (this *QCborValueRef) OperatorSubscriptWithKey(key *QLatin1String) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborValueRef) OperatorSubscript2(key string) *QCborValueRef {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorSubscript3, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborValueRef) ToQCborValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_ToQCborValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValueRef) Type() QCborValue__Type {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborValue__Type(__r)
}

func (this *QCborValueRef) IsInteger() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsByteArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isByteArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsMap() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isMap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsTag() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isTag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsFalse() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isFalse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsTrue() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isTrue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsDateTime() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isDateTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsUrl() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isUrl, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsRegularExpression() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isRegularExpression, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsUuid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isUuid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsInvalid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isInvalid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsContainer() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isContainer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsSimpleType() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_isSimpleType2, this.h, uintptr(st))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) ToSimpleType() QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_toSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborValueRef) Tag() QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_tag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborValueRef) TaggedValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_taggedValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValueRef) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueRef) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) ToDouble() float64 {
	__r := pg_QCborValueRef_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueRef) ToByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValueRef) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValueRef) ToDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toDateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QCborValueRef) ToUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toUrl, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QCborValueRef) ToRegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toRegularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QCborValueRef) ToUuid() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toUuid, this.h)
	runtime.KeepAlive(this)
	return newownedQUuid(__r)
}

func (this *QCborValueRef) ToArray() *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toArray, this.h)
	runtime.KeepAlive(this)
	return newownedQCborArray(__r)
}

func (this *QCborValueRef) ToArrayWithQCborArray(a *QCborArray) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toArray2, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return newownedQCborArray(__r)
}

func (this *QCborValueRef) ToMap() *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toMap, this.h)
	runtime.KeepAlive(this)
	return newownedQCborMap(__r)
}

func (this *QCborValueRef) ToMapWithQCborMap(m *QCborMap) *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toMap2, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQCborMap(__r)
}

func (this *QCborValueRef) OperatorSubscript3(key string) *QCborValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorSubscript4, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborValueRef) OperatorSubscript4(key *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorSubscript5, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborValueRef) OperatorSubscript5(key int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_OperatorSubscript6, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValueRef) Compare(other *QCborValue) int {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QCborValueRef) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCborValueRef) ToJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toJsonValue, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QCborValueRef) ToCbor() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toCbor, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValueRef) ToCborWithWriter(writer *QCborStreamWriter) {
	purego.SyscallN(addr_QCborValueRef_toCbor2, this.h, uintptr(writer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(writer)
}

func (this *QCborValueRef) ToDiagnosticNotation() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toDiagnosticNotation, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValueRef) ToSimpleTypeWithDefaultValue(defaultValue QCborSimpleType) QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_toSimpleType2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborValueRef) TagWithDefaultValue(defaultValue QCborTag) QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_tag2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborValueRef) TaggedValueWithDefaultValue(defaultValue *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_taggedValue2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQCborValue(__r)
}

func (this *QCborValueRef) ToIntegerWithDefaultValue(defaultValue int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_toInteger2, this.h, uintptr(defaultValue))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueRef) ToBoolWithDefaultValue(defaultValue bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueRef_toBool2, this.h, uintptr(b2i(defaultValue)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueRef) ToDoubleWithDefaultValue(defaultValue float64) float64 {
	__r := pg_QCborValueRef_toDouble2(this.h, defaultValue)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueRef) ToByteArrayWithDefaultValue(defaultValue *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toByteArray2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQByteArray(__r)
}

func (this *QCborValueRef) ToStringWithDefaultValue(defaultValue string) string {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toString2, this.h, defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QCborValueRef) ToDateTimeWithDefaultValue(defaultValue *QDateTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toDateTime2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQDateTime(__r)
}

func (this *QCborValueRef) ToUrlWithDefaultValue(defaultValue *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toUrl2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQUrl(__r)
}

func (this *QCborValueRef) ToRegularExpressionWithDefaultValue(defaultValue *QRegularExpression) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toRegularExpression2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQRegularExpression(__r)
}

func (this *QCborValueRef) ToUuidWithDefaultValue(defaultValue *QUuid) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toUuid2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQUuid(__r)
}

func (this *QCborValueRef) ToCborWithOpt(opt QCborValue__EncodingOption) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toCbor3, this.h, uintptr(opt))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValueRef) ToCbor2(writer *QCborStreamWriter, opt QCborValue__EncodingOption) {
	purego.SyscallN(addr_QCborValueRef_toCbor4, this.h, uintptr(writer.UnsafePointer()), uintptr(opt))
	runtime.KeepAlive(this)
	runtime.KeepAlive(writer)
}

func (this *QCborValueRef) ToDiagnosticNotationWithOpt(opt QCborValue__DiagnosticNotationOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueRef_toDiagnosticNotation2, this.h, uintptr(opt), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValueRef) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborValueRef_delete(this.h)
	}
	this.UnsafeRelease()
}
