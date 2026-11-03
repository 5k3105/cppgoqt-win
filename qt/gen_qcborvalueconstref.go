package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborValueConstRef_new uintptr
	addr_QCborValueConstRef_ToQCborValue uintptr
	addr_QCborValueConstRef_type uintptr
	addr_QCborValueConstRef_isInteger uintptr
	addr_QCborValueConstRef_isByteArray uintptr
	addr_QCborValueConstRef_isString uintptr
	addr_QCborValueConstRef_isArray uintptr
	addr_QCborValueConstRef_isMap uintptr
	addr_QCborValueConstRef_isTag uintptr
	addr_QCborValueConstRef_isFalse uintptr
	addr_QCborValueConstRef_isTrue uintptr
	addr_QCborValueConstRef_isBool uintptr
	addr_QCborValueConstRef_isNull uintptr
	addr_QCborValueConstRef_isUndefined uintptr
	addr_QCborValueConstRef_isDouble uintptr
	addr_QCborValueConstRef_isDateTime uintptr
	addr_QCborValueConstRef_isUrl uintptr
	addr_QCborValueConstRef_isRegularExpression uintptr
	addr_QCborValueConstRef_isUuid uintptr
	addr_QCborValueConstRef_isInvalid uintptr
	addr_QCborValueConstRef_isContainer uintptr
	addr_QCborValueConstRef_isSimpleType uintptr
	addr_QCborValueConstRef_isSimpleType2 uintptr
	addr_QCborValueConstRef_toSimpleType uintptr
	addr_QCborValueConstRef_tag uintptr
	addr_QCborValueConstRef_taggedValue uintptr
	addr_QCborValueConstRef_toInteger uintptr
	addr_QCborValueConstRef_toBool uintptr
	pg_QCborValueConstRef_toDouble func(uintptr) float64
	addr_QCborValueConstRef_toByteArray uintptr
	addr_QCborValueConstRef_toString uintptr
	addr_QCborValueConstRef_toStringView uintptr
	addr_QCborValueConstRef_toDateTime uintptr
	addr_QCborValueConstRef_toUrl uintptr
	addr_QCborValueConstRef_toRegularExpression uintptr
	addr_QCborValueConstRef_toUuid uintptr
	addr_QCborValueConstRef_toArray uintptr
	addr_QCborValueConstRef_toArray2 uintptr
	addr_QCborValueConstRef_toMap uintptr
	addr_QCborValueConstRef_toMap2 uintptr
	addr_QCborValueConstRef_OperatorSubscript uintptr
	addr_QCborValueConstRef_OperatorSubscript2 uintptr
	addr_QCborValueConstRef_OperatorSubscript3 uintptr
	addr_QCborValueConstRef_compare uintptr
	addr_QCborValueConstRef_toVariant uintptr
	addr_QCborValueConstRef_toJsonValue uintptr
	addr_QCborValueConstRef_toCbor uintptr
	addr_QCborValueConstRef_toCbor2 uintptr
	addr_QCborValueConstRef_toDiagnosticNotation uintptr
	addr_QCborValueConstRef_toSimpleType2 uintptr
	addr_QCborValueConstRef_tag2 uintptr
	addr_QCborValueConstRef_taggedValue2 uintptr
	addr_QCborValueConstRef_toInteger2 uintptr
	addr_QCborValueConstRef_toBool2 uintptr
	pg_QCborValueConstRef_toDouble2 func(uintptr, float64) float64
	addr_QCborValueConstRef_toByteArray2 uintptr
	addr_QCborValueConstRef_toString2 uintptr
	addr_QCborValueConstRef_toStringView2 uintptr
	addr_QCborValueConstRef_toDateTime2 uintptr
	addr_QCborValueConstRef_toUrl2 uintptr
	addr_QCborValueConstRef_toRegularExpression2 uintptr
	addr_QCborValueConstRef_toUuid2 uintptr
	addr_QCborValueConstRef_toCbor3 uintptr
	addr_QCborValueConstRef_toCbor4 uintptr
	addr_QCborValueConstRef_toDiagnosticNotation2 uintptr
	addr_QCborValueConstRef_OperatorEqual uintptr
	addr_QCborValueConstRef_OperatorNotEqual uintptr
	addr_QCborValueConstRef_OperatorLesser uintptr
	addr_QCborValueConstRef_OperatorGreater uintptr
	addr_QCborValueConstRef_OperatorLesserOrEqual uintptr
	addr_QCborValueConstRef_OperatorGreaterOrEqual uintptr
	addr_QCborValueConstRef_OperatorEqual2 uintptr
	addr_QCborValueConstRef_OperatorNotEqual2 uintptr
	addr_QCborValueConstRef_OperatorLesser2 uintptr
	addr_QCborValueConstRef_OperatorGreater2 uintptr
	addr_QCborValueConstRef_OperatorLesserOrEqual2 uintptr
	addr_QCborValueConstRef_OperatorGreaterOrEqual2 uintptr
	addr_QCborValueConstRef_OperatorEqual3 uintptr
	addr_QCborValueConstRef_OperatorNotEqual3 uintptr
	addr_QCborValueConstRef_OperatorLesser3 uintptr
	addr_QCborValueConstRef_OperatorGreater3 uintptr
	addr_QCborValueConstRef_OperatorLesserOrEqual3 uintptr
	addr_QCborValueConstRef_OperatorGreaterOrEqual3 uintptr
	addr_QCborValueConstRef_OperatorEqual4 uintptr
	addr_QCborValueConstRef_OperatorNotEqual4 uintptr
	addr_QCborValueConstRef_OperatorLesser4 uintptr
	addr_QCborValueConstRef_OperatorGreater4 uintptr
	addr_QCborValueConstRef_OperatorLesserOrEqual4 uintptr
	addr_QCborValueConstRef_OperatorGreaterOrEqual4 uintptr
	pg_QCborValueConstRef_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborValueConstRef_new, "QCborValueConstRef_new")
	symaddr(&addr_QCborValueConstRef_ToQCborValue, "QCborValueConstRef_ToQCborValue")
	symaddr(&addr_QCborValueConstRef_type, "QCborValueConstRef_type")
	symaddr(&addr_QCborValueConstRef_isInteger, "QCborValueConstRef_isInteger")
	symaddr(&addr_QCborValueConstRef_isByteArray, "QCborValueConstRef_isByteArray")
	symaddr(&addr_QCborValueConstRef_isString, "QCborValueConstRef_isString")
	symaddr(&addr_QCborValueConstRef_isArray, "QCborValueConstRef_isArray")
	symaddr(&addr_QCborValueConstRef_isMap, "QCborValueConstRef_isMap")
	symaddr(&addr_QCborValueConstRef_isTag, "QCborValueConstRef_isTag")
	symaddr(&addr_QCborValueConstRef_isFalse, "QCborValueConstRef_isFalse")
	symaddr(&addr_QCborValueConstRef_isTrue, "QCborValueConstRef_isTrue")
	symaddr(&addr_QCborValueConstRef_isBool, "QCborValueConstRef_isBool")
	symaddr(&addr_QCborValueConstRef_isNull, "QCborValueConstRef_isNull")
	symaddr(&addr_QCborValueConstRef_isUndefined, "QCborValueConstRef_isUndefined")
	symaddr(&addr_QCborValueConstRef_isDouble, "QCborValueConstRef_isDouble")
	symaddr(&addr_QCborValueConstRef_isDateTime, "QCborValueConstRef_isDateTime")
	symaddr(&addr_QCborValueConstRef_isUrl, "QCborValueConstRef_isUrl")
	symaddr(&addr_QCborValueConstRef_isRegularExpression, "QCborValueConstRef_isRegularExpression")
	symaddr(&addr_QCborValueConstRef_isUuid, "QCborValueConstRef_isUuid")
	symaddr(&addr_QCborValueConstRef_isInvalid, "QCborValueConstRef_isInvalid")
	symaddr(&addr_QCborValueConstRef_isContainer, "QCborValueConstRef_isContainer")
	symaddr(&addr_QCborValueConstRef_isSimpleType, "QCborValueConstRef_isSimpleType")
	symaddr(&addr_QCborValueConstRef_isSimpleType2, "QCborValueConstRef_isSimpleType2")
	symaddr(&addr_QCborValueConstRef_toSimpleType, "QCborValueConstRef_toSimpleType")
	symaddr(&addr_QCborValueConstRef_tag, "QCborValueConstRef_tag")
	symaddr(&addr_QCborValueConstRef_taggedValue, "QCborValueConstRef_taggedValue")
	symaddr(&addr_QCborValueConstRef_toInteger, "QCborValueConstRef_toInteger")
	symaddr(&addr_QCborValueConstRef_toBool, "QCborValueConstRef_toBool")
	reg(&pg_QCborValueConstRef_toDouble, "QCborValueConstRef_toDouble")
	symaddr(&addr_QCborValueConstRef_toByteArray, "QCborValueConstRef_toByteArray")
	symaddr(&addr_QCborValueConstRef_toString, "QCborValueConstRef_toString")
	symaddr(&addr_QCborValueConstRef_toStringView, "QCborValueConstRef_toStringView")
	symaddr(&addr_QCborValueConstRef_toDateTime, "QCborValueConstRef_toDateTime")
	symaddr(&addr_QCborValueConstRef_toUrl, "QCborValueConstRef_toUrl")
	symaddr(&addr_QCborValueConstRef_toRegularExpression, "QCborValueConstRef_toRegularExpression")
	symaddr(&addr_QCborValueConstRef_toUuid, "QCborValueConstRef_toUuid")
	symaddr(&addr_QCborValueConstRef_toArray, "QCborValueConstRef_toArray")
	symaddr(&addr_QCborValueConstRef_toArray2, "QCborValueConstRef_toArray2")
	symaddr(&addr_QCborValueConstRef_toMap, "QCborValueConstRef_toMap")
	symaddr(&addr_QCborValueConstRef_toMap2, "QCborValueConstRef_toMap2")
	symaddr(&addr_QCborValueConstRef_OperatorSubscript, "QCborValueConstRef_OperatorSubscript")
	symaddr(&addr_QCborValueConstRef_OperatorSubscript2, "QCborValueConstRef_OperatorSubscript2")
	symaddr(&addr_QCborValueConstRef_OperatorSubscript3, "QCborValueConstRef_OperatorSubscript3")
	symaddr(&addr_QCborValueConstRef_compare, "QCborValueConstRef_compare")
	symaddr(&addr_QCborValueConstRef_toVariant, "QCborValueConstRef_toVariant")
	symaddr(&addr_QCborValueConstRef_toJsonValue, "QCborValueConstRef_toJsonValue")
	symaddr(&addr_QCborValueConstRef_toCbor, "QCborValueConstRef_toCbor")
	symaddr(&addr_QCborValueConstRef_toCbor2, "QCborValueConstRef_toCbor2")
	symaddr(&addr_QCborValueConstRef_toDiagnosticNotation, "QCborValueConstRef_toDiagnosticNotation")
	symaddr(&addr_QCborValueConstRef_toSimpleType2, "QCborValueConstRef_toSimpleType2")
	symaddr(&addr_QCborValueConstRef_tag2, "QCborValueConstRef_tag2")
	symaddr(&addr_QCborValueConstRef_taggedValue2, "QCborValueConstRef_taggedValue2")
	symaddr(&addr_QCborValueConstRef_toInteger2, "QCborValueConstRef_toInteger2")
	symaddr(&addr_QCborValueConstRef_toBool2, "QCborValueConstRef_toBool2")
	reg(&pg_QCborValueConstRef_toDouble2, "QCborValueConstRef_toDouble2")
	symaddr(&addr_QCborValueConstRef_toByteArray2, "QCborValueConstRef_toByteArray2")
	symaddr(&addr_QCborValueConstRef_toString2, "QCborValueConstRef_toString2")
	symaddr(&addr_QCborValueConstRef_toStringView2, "QCborValueConstRef_toStringView2")
	symaddr(&addr_QCborValueConstRef_toDateTime2, "QCborValueConstRef_toDateTime2")
	symaddr(&addr_QCborValueConstRef_toUrl2, "QCborValueConstRef_toUrl2")
	symaddr(&addr_QCborValueConstRef_toRegularExpression2, "QCborValueConstRef_toRegularExpression2")
	symaddr(&addr_QCborValueConstRef_toUuid2, "QCborValueConstRef_toUuid2")
	symaddr(&addr_QCborValueConstRef_toCbor3, "QCborValueConstRef_toCbor3")
	symaddr(&addr_QCborValueConstRef_toCbor4, "QCborValueConstRef_toCbor4")
	symaddr(&addr_QCborValueConstRef_toDiagnosticNotation2, "QCborValueConstRef_toDiagnosticNotation2")
	symaddr(&addr_QCborValueConstRef_OperatorEqual, "QCborValueConstRef_OperatorEqual")
	symaddr(&addr_QCborValueConstRef_OperatorNotEqual, "QCborValueConstRef_OperatorNotEqual")
	symaddr(&addr_QCborValueConstRef_OperatorLesser, "QCborValueConstRef_OperatorLesser")
	symaddr(&addr_QCborValueConstRef_OperatorGreater, "QCborValueConstRef_OperatorGreater")
	symaddr(&addr_QCborValueConstRef_OperatorLesserOrEqual, "QCborValueConstRef_OperatorLesserOrEqual")
	symaddr(&addr_QCborValueConstRef_OperatorGreaterOrEqual, "QCborValueConstRef_OperatorGreaterOrEqual")
	symaddr(&addr_QCborValueConstRef_OperatorEqual2, "QCborValueConstRef_OperatorEqual2")
	symaddr(&addr_QCborValueConstRef_OperatorNotEqual2, "QCborValueConstRef_OperatorNotEqual2")
	symaddr(&addr_QCborValueConstRef_OperatorLesser2, "QCborValueConstRef_OperatorLesser2")
	symaddr(&addr_QCborValueConstRef_OperatorGreater2, "QCborValueConstRef_OperatorGreater2")
	symaddr(&addr_QCborValueConstRef_OperatorLesserOrEqual2, "QCborValueConstRef_OperatorLesserOrEqual2")
	symaddr(&addr_QCborValueConstRef_OperatorGreaterOrEqual2, "QCborValueConstRef_OperatorGreaterOrEqual2")
	symaddr(&addr_QCborValueConstRef_OperatorEqual3, "QCborValueConstRef_OperatorEqual3")
	symaddr(&addr_QCborValueConstRef_OperatorNotEqual3, "QCborValueConstRef_OperatorNotEqual3")
	symaddr(&addr_QCborValueConstRef_OperatorLesser3, "QCborValueConstRef_OperatorLesser3")
	symaddr(&addr_QCborValueConstRef_OperatorGreater3, "QCborValueConstRef_OperatorGreater3")
	symaddr(&addr_QCborValueConstRef_OperatorLesserOrEqual3, "QCborValueConstRef_OperatorLesserOrEqual3")
	symaddr(&addr_QCborValueConstRef_OperatorGreaterOrEqual3, "QCborValueConstRef_OperatorGreaterOrEqual3")
	symaddr(&addr_QCborValueConstRef_OperatorEqual4, "QCborValueConstRef_OperatorEqual4")
	symaddr(&addr_QCborValueConstRef_OperatorNotEqual4, "QCborValueConstRef_OperatorNotEqual4")
	symaddr(&addr_QCborValueConstRef_OperatorLesser4, "QCborValueConstRef_OperatorLesser4")
	symaddr(&addr_QCborValueConstRef_OperatorGreater4, "QCborValueConstRef_OperatorGreater4")
	symaddr(&addr_QCborValueConstRef_OperatorLesserOrEqual4, "QCborValueConstRef_OperatorLesserOrEqual4")
	symaddr(&addr_QCborValueConstRef_OperatorGreaterOrEqual4, "QCborValueConstRef_OperatorGreaterOrEqual4")
	reg(&pg_QCborValueConstRef_delete, "QCborValueConstRef_delete")
}

type QCborValueConstRef struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborValueConstRef) UnsafeGoID() int64     { return this.vid }
func (this *QCborValueConstRef) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborValueConstRef(h uintptr) *QCborValueConstRef {
	if h == 0 {
		return nil
	}
	return &QCborValueConstRef{h: h}
}

func newownedQCborValueConstRef(h uintptr) *QCborValueConstRef {
	o := bindQCborValueConstRef(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborValueConstRef).Delete)
	}
	return o
}

func (this *QCborValueConstRef) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborValueConstRef) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborValueConstRef(h unsafe.Pointer) *QCborValueConstRef { return bindQCborValueConstRef(uintptr(h)) }
func UnsafeNewQCborValueConstRefOwned(h unsafe.Pointer) *QCborValueConstRef { return newownedQCborValueConstRef(uintptr(h)) }

func (this *QCborValueConstRef) UnsafeRelease() { this.h = 0 }

func NewQCborValueConstRef(arg1 *QCborValueConstRef) *QCborValueConstRef {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQCborValueConstRef(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QCborValueConstRef) ToQCborValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_ToQCborValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValueConstRef) Type() QCborValue__Type {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborValue__Type(__r)
}

func (this *QCborValueConstRef) IsInteger() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsByteArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isByteArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsMap() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isMap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsTag() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isTag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsFalse() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isFalse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsTrue() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isTrue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsDateTime() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isDateTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsUrl() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isUrl, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsRegularExpression() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isRegularExpression, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsUuid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isUuid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsInvalid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isInvalid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsContainer() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isContainer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsSimpleType() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_isSimpleType2, this.h, uintptr(st))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) ToSimpleType() QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_toSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborValueConstRef) Tag() QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_tag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborValueConstRef) TaggedValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_taggedValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValueConstRef) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueConstRef) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) ToDouble() float64 {
	__r := pg_QCborValueConstRef_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueConstRef) ToByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValueConstRef) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValueConstRef) ToStringView() *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toStringView, this.h)
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QCborValueConstRef) ToDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toDateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QCborValueConstRef) ToUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toUrl, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QCborValueConstRef) ToRegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toRegularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QCborValueConstRef) ToUuid() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toUuid, this.h)
	runtime.KeepAlive(this)
	return newownedQUuid(__r)
}

func (this *QCborValueConstRef) ToArray() *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toArray, this.h)
	runtime.KeepAlive(this)
	return newownedQCborArray(__r)
}

func (this *QCborValueConstRef) ToArrayWithQCborArray(a *QCborArray) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toArray2, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return newownedQCborArray(__r)
}

func (this *QCborValueConstRef) ToMap() *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toMap, this.h)
	runtime.KeepAlive(this)
	return newownedQCborMap(__r)
}

func (this *QCborValueConstRef) ToMapWithQCborMap(m *QCborMap) *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toMap2, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQCborMap(__r)
}

func (this *QCborValueConstRef) OperatorSubscript(key string) *QCborValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorSubscript, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborValueConstRef) OperatorSubscriptWithKey(key *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborValueConstRef) OperatorSubscript2(key int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorSubscript3, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValueConstRef) Compare(other *QCborValue) int {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QCborValueConstRef) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCborValueConstRef) ToJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toJsonValue, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QCborValueConstRef) ToCbor() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toCbor, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValueConstRef) ToCborWithWriter(writer *QCborStreamWriter) {
	purego.SyscallN(addr_QCborValueConstRef_toCbor2, this.h, uintptr(writer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(writer)
}

func (this *QCborValueConstRef) ToDiagnosticNotation() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toDiagnosticNotation, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValueConstRef) ToSimpleTypeWithDefaultValue(defaultValue QCborSimpleType) QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_toSimpleType2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborValueConstRef) TagWithDefaultValue(defaultValue QCborTag) QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_tag2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborValueConstRef) TaggedValueWithDefaultValue(defaultValue *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_taggedValue2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQCborValue(__r)
}

func (this *QCborValueConstRef) ToIntegerWithDefaultValue(defaultValue int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_toInteger2, this.h, uintptr(defaultValue))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueConstRef) ToBoolWithDefaultValue(defaultValue bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_toBool2, this.h, uintptr(b2i(defaultValue)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValueConstRef) ToDoubleWithDefaultValue(defaultValue float64) float64 {
	__r := pg_QCborValueConstRef_toDouble2(this.h, defaultValue)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValueConstRef) ToByteArrayWithDefaultValue(defaultValue *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toByteArray2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQByteArray(__r)
}

func (this *QCborValueConstRef) ToStringWithDefaultValue(defaultValue string) string {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toString2, this.h, defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QCborValueConstRef) ToStringViewWithDefaultValue(defaultValue string) *QAnyStringView {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toStringView2, this.h, defaultValue_p, uintptr(len(defaultValue)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQAnyStringView(__r)
}

func (this *QCborValueConstRef) ToDateTimeWithDefaultValue(defaultValue *QDateTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toDateTime2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQDateTime(__r)
}

func (this *QCborValueConstRef) ToUrlWithDefaultValue(defaultValue *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toUrl2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQUrl(__r)
}

func (this *QCborValueConstRef) ToRegularExpressionWithDefaultValue(defaultValue *QRegularExpression) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toRegularExpression2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQRegularExpression(__r)
}

func (this *QCborValueConstRef) ToUuidWithDefaultValue(defaultValue *QUuid) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toUuid2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQUuid(__r)
}

func (this *QCborValueConstRef) ToCborWithOpt(opt QCborValue__EncodingOption) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toCbor3, this.h, uintptr(opt))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValueConstRef) ToCbor2(writer *QCborStreamWriter, opt QCborValue__EncodingOption) {
	purego.SyscallN(addr_QCborValueConstRef_toCbor4, this.h, uintptr(writer.UnsafePointer()), uintptr(opt))
	runtime.KeepAlive(this)
	runtime.KeepAlive(writer)
}

func (this *QCborValueConstRef) ToDiagnosticNotationWithOpt(opt QCborValue__DiagnosticNotationOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValueConstRef_toDiagnosticNotation2, this.h, uintptr(opt), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValueConstRef) OperatorEqual(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorNotEqual(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesser(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreater(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesserOrEqual(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreaterOrEqual(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorNotEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesserWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreaterWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesserOrEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreaterOrEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorNotEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesser2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreater2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesserOrEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreaterOrEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorNotEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesser3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesser4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreater3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreater4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorLesserOrEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorLesserOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) OperatorGreaterOrEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValueConstRef_OperatorGreaterOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValueConstRef) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborValueConstRef_delete(this.h)
	}
	this.UnsafeRelease()
}
