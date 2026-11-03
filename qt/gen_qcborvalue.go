package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborValue_new uintptr
	addr_QCborValue_new2 uintptr
	addr_QCborValue_new3 uintptr
	addr_QCborValue_new4 uintptr
	addr_QCborValue_new5 uintptr
	addr_QCborValue_new6 uintptr
	pg_QCborValue_new7 func(float64) uintptr
	addr_QCborValue_new8 uintptr
	addr_QCborValue_new9 uintptr
	addr_QCborValue_new10 uintptr
	addr_QCborValue_new11 uintptr
	addr_QCborValue_new12 uintptr
	addr_QCborValue_new13 uintptr
	addr_QCborValue_new14 uintptr
	addr_QCborValue_new15 uintptr
	addr_QCborValue_new16 uintptr
	addr_QCborValue_new17 uintptr
	addr_QCborValue_new18 uintptr
	addr_QCborValue_new19 uintptr
	addr_QCborValue_new20 uintptr
	addr_QCborValue_new21 uintptr
	addr_QCborValue_new22 uintptr
	addr_QCborValue_new23 uintptr
	addr_QCborValue_new24 uintptr
	addr_QCborValue_OperatorAssign uintptr
	addr_QCborValue_swap uintptr
	addr_QCborValue_type uintptr
	addr_QCborValue_isInteger uintptr
	addr_QCborValue_isByteArray uintptr
	addr_QCborValue_isString uintptr
	addr_QCborValue_isArray uintptr
	addr_QCborValue_isMap uintptr
	addr_QCborValue_isTag uintptr
	addr_QCborValue_isFalse uintptr
	addr_QCborValue_isTrue uintptr
	addr_QCborValue_isBool uintptr
	addr_QCborValue_isNull uintptr
	addr_QCborValue_isUndefined uintptr
	addr_QCborValue_isDouble uintptr
	addr_QCborValue_isDateTime uintptr
	addr_QCborValue_isUrl uintptr
	addr_QCborValue_isRegularExpression uintptr
	addr_QCborValue_isUuid uintptr
	addr_QCborValue_isInvalid uintptr
	addr_QCborValue_isContainer uintptr
	addr_QCborValue_isSimpleType uintptr
	addr_QCborValue_isSimpleType2 uintptr
	addr_QCborValue_toSimpleType uintptr
	addr_QCborValue_toInteger uintptr
	addr_QCborValue_toBool uintptr
	pg_QCborValue_toDouble func(uintptr) float64
	addr_QCborValue_tag uintptr
	addr_QCborValue_taggedValue uintptr
	addr_QCborValue_toByteArray uintptr
	addr_QCborValue_toString uintptr
	addr_QCborValue_toStringView uintptr
	addr_QCborValue_toDateTime uintptr
	addr_QCborValue_toUrl uintptr
	addr_QCborValue_toRegularExpression uintptr
	addr_QCborValue_toUuid uintptr
	addr_QCborValue_toArray uintptr
	addr_QCborValue_toArray2 uintptr
	addr_QCborValue_toMap uintptr
	addr_QCborValue_toMap2 uintptr
	addr_QCborValue_OperatorSubscript uintptr
	addr_QCborValue_OperatorSubscript2 uintptr
	addr_QCborValue_OperatorSubscript3 uintptr
	addr_QCborValue_OperatorSubscript4 uintptr
	addr_QCborValue_OperatorSubscript5 uintptr
	addr_QCborValue_OperatorSubscript6 uintptr
	addr_QCborValue_compare uintptr
	addr_QCborValue_fromVariant uintptr
	addr_QCborValue_toVariant uintptr
	addr_QCborValue_fromJsonValue uintptr
	addr_QCborValue_toJsonValue uintptr
	addr_QCborValue_fromCbor uintptr
	addr_QCborValue_fromCbor2 uintptr
	addr_QCborValue_fromCbor3 uintptr
	addr_QCborValue_fromCbor4 uintptr
	addr_QCborValue_toCbor uintptr
	addr_QCborValue_toCbor2 uintptr
	addr_QCborValue_toDiagnosticNotation uintptr
	addr_QCborValue_toSimpleType2 uintptr
	addr_QCborValue_toInteger2 uintptr
	addr_QCborValue_toBool2 uintptr
	pg_QCborValue_toDouble2 func(uintptr, float64) float64
	addr_QCborValue_tag2 uintptr
	addr_QCborValue_taggedValue2 uintptr
	addr_QCborValue_toByteArray2 uintptr
	addr_QCborValue_toString2 uintptr
	addr_QCborValue_toStringView2 uintptr
	addr_QCborValue_toDateTime2 uintptr
	addr_QCborValue_toUrl2 uintptr
	addr_QCborValue_toRegularExpression2 uintptr
	addr_QCborValue_toUuid2 uintptr
	addr_QCborValue_fromCbor5 uintptr
	addr_QCborValue_fromCbor6 uintptr
	addr_QCborValue_fromCbor7 uintptr
	addr_QCborValue_toCbor3 uintptr
	addr_QCborValue_toCbor4 uintptr
	addr_QCborValue_toDiagnosticNotation2 uintptr
	addr_QCborValue_OperatorEqual uintptr
	addr_QCborValue_OperatorNotEqual uintptr
	addr_QCborValue_OperatorLesser uintptr
	addr_QCborValue_OperatorGreater uintptr
	addr_QCborValue_OperatorLesserOrEqual uintptr
	addr_QCborValue_OperatorGreaterOrEqual uintptr
	addr_QCborValue_OperatorEqual2 uintptr
	addr_QCborValue_OperatorNotEqual2 uintptr
	addr_QCborValue_OperatorLesser2 uintptr
	addr_QCborValue_OperatorGreater2 uintptr
	addr_QCborValue_OperatorLesserOrEqual2 uintptr
	addr_QCborValue_OperatorGreaterOrEqual2 uintptr
	addr_QCborValue_OperatorEqual3 uintptr
	addr_QCborValue_OperatorNotEqual3 uintptr
	addr_QCborValue_OperatorLesser3 uintptr
	addr_QCborValue_OperatorGreater3 uintptr
	addr_QCborValue_OperatorLesserOrEqual3 uintptr
	addr_QCborValue_OperatorGreaterOrEqual3 uintptr
	addr_QCborValue_OperatorEqual4 uintptr
	addr_QCborValue_OperatorNotEqual4 uintptr
	addr_QCborValue_OperatorLesser4 uintptr
	addr_QCborValue_OperatorGreater4 uintptr
	addr_QCborValue_OperatorLesserOrEqual4 uintptr
	addr_QCborValue_OperatorGreaterOrEqual4 uintptr
	pg_QCborValue_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborValue_new, "QCborValue_new")
	symaddr(&addr_QCborValue_new2, "QCborValue_new2")
	symaddr(&addr_QCborValue_new3, "QCborValue_new3")
	symaddr(&addr_QCborValue_new4, "QCborValue_new4")
	symaddr(&addr_QCborValue_new5, "QCborValue_new5")
	symaddr(&addr_QCborValue_new6, "QCborValue_new6")
	reg(&pg_QCborValue_new7, "QCborValue_new7")
	symaddr(&addr_QCborValue_new8, "QCborValue_new8")
	symaddr(&addr_QCborValue_new9, "QCborValue_new9")
	symaddr(&addr_QCborValue_new10, "QCborValue_new10")
	symaddr(&addr_QCborValue_new11, "QCborValue_new11")
	symaddr(&addr_QCborValue_new12, "QCborValue_new12")
	symaddr(&addr_QCborValue_new13, "QCborValue_new13")
	symaddr(&addr_QCborValue_new14, "QCborValue_new14")
	symaddr(&addr_QCborValue_new15, "QCborValue_new15")
	symaddr(&addr_QCborValue_new16, "QCborValue_new16")
	symaddr(&addr_QCborValue_new17, "QCborValue_new17")
	symaddr(&addr_QCborValue_new18, "QCborValue_new18")
	symaddr(&addr_QCborValue_new19, "QCborValue_new19")
	symaddr(&addr_QCborValue_new20, "QCborValue_new20")
	symaddr(&addr_QCborValue_new21, "QCborValue_new21")
	symaddr(&addr_QCborValue_new22, "QCborValue_new22")
	symaddr(&addr_QCborValue_new23, "QCborValue_new23")
	symaddr(&addr_QCborValue_new24, "QCborValue_new24")
	symaddr(&addr_QCborValue_OperatorAssign, "QCborValue_OperatorAssign")
	symaddr(&addr_QCborValue_swap, "QCborValue_swap")
	symaddr(&addr_QCborValue_type, "QCborValue_type")
	symaddr(&addr_QCborValue_isInteger, "QCborValue_isInteger")
	symaddr(&addr_QCborValue_isByteArray, "QCborValue_isByteArray")
	symaddr(&addr_QCborValue_isString, "QCborValue_isString")
	symaddr(&addr_QCborValue_isArray, "QCborValue_isArray")
	symaddr(&addr_QCborValue_isMap, "QCborValue_isMap")
	symaddr(&addr_QCborValue_isTag, "QCborValue_isTag")
	symaddr(&addr_QCborValue_isFalse, "QCborValue_isFalse")
	symaddr(&addr_QCborValue_isTrue, "QCborValue_isTrue")
	symaddr(&addr_QCborValue_isBool, "QCborValue_isBool")
	symaddr(&addr_QCborValue_isNull, "QCborValue_isNull")
	symaddr(&addr_QCborValue_isUndefined, "QCborValue_isUndefined")
	symaddr(&addr_QCborValue_isDouble, "QCborValue_isDouble")
	symaddr(&addr_QCborValue_isDateTime, "QCborValue_isDateTime")
	symaddr(&addr_QCborValue_isUrl, "QCborValue_isUrl")
	symaddr(&addr_QCborValue_isRegularExpression, "QCborValue_isRegularExpression")
	symaddr(&addr_QCborValue_isUuid, "QCborValue_isUuid")
	symaddr(&addr_QCborValue_isInvalid, "QCborValue_isInvalid")
	symaddr(&addr_QCborValue_isContainer, "QCborValue_isContainer")
	symaddr(&addr_QCborValue_isSimpleType, "QCborValue_isSimpleType")
	symaddr(&addr_QCborValue_isSimpleType2, "QCborValue_isSimpleType2")
	symaddr(&addr_QCborValue_toSimpleType, "QCborValue_toSimpleType")
	symaddr(&addr_QCborValue_toInteger, "QCborValue_toInteger")
	symaddr(&addr_QCborValue_toBool, "QCborValue_toBool")
	reg(&pg_QCborValue_toDouble, "QCborValue_toDouble")
	symaddr(&addr_QCborValue_tag, "QCborValue_tag")
	symaddr(&addr_QCborValue_taggedValue, "QCborValue_taggedValue")
	symaddr(&addr_QCborValue_toByteArray, "QCborValue_toByteArray")
	symaddr(&addr_QCborValue_toString, "QCborValue_toString")
	symaddr(&addr_QCborValue_toStringView, "QCborValue_toStringView")
	symaddr(&addr_QCborValue_toDateTime, "QCborValue_toDateTime")
	symaddr(&addr_QCborValue_toUrl, "QCborValue_toUrl")
	symaddr(&addr_QCborValue_toRegularExpression, "QCborValue_toRegularExpression")
	symaddr(&addr_QCborValue_toUuid, "QCborValue_toUuid")
	symaddr(&addr_QCborValue_toArray, "QCborValue_toArray")
	symaddr(&addr_QCborValue_toArray2, "QCborValue_toArray2")
	symaddr(&addr_QCborValue_toMap, "QCborValue_toMap")
	symaddr(&addr_QCborValue_toMap2, "QCborValue_toMap2")
	symaddr(&addr_QCborValue_OperatorSubscript, "QCborValue_OperatorSubscript")
	symaddr(&addr_QCborValue_OperatorSubscript2, "QCborValue_OperatorSubscript2")
	symaddr(&addr_QCborValue_OperatorSubscript3, "QCborValue_OperatorSubscript3")
	symaddr(&addr_QCborValue_OperatorSubscript4, "QCborValue_OperatorSubscript4")
	symaddr(&addr_QCborValue_OperatorSubscript5, "QCborValue_OperatorSubscript5")
	symaddr(&addr_QCborValue_OperatorSubscript6, "QCborValue_OperatorSubscript6")
	symaddr(&addr_QCborValue_compare, "QCborValue_compare")
	symaddr(&addr_QCborValue_fromVariant, "QCborValue_fromVariant")
	symaddr(&addr_QCborValue_toVariant, "QCborValue_toVariant")
	symaddr(&addr_QCborValue_fromJsonValue, "QCborValue_fromJsonValue")
	symaddr(&addr_QCborValue_toJsonValue, "QCborValue_toJsonValue")
	symaddr(&addr_QCborValue_fromCbor, "QCborValue_fromCbor")
	symaddr(&addr_QCborValue_fromCbor2, "QCborValue_fromCbor2")
	symaddr(&addr_QCborValue_fromCbor3, "QCborValue_fromCbor3")
	symaddr(&addr_QCborValue_fromCbor4, "QCborValue_fromCbor4")
	symaddr(&addr_QCborValue_toCbor, "QCborValue_toCbor")
	symaddr(&addr_QCborValue_toCbor2, "QCborValue_toCbor2")
	symaddr(&addr_QCborValue_toDiagnosticNotation, "QCborValue_toDiagnosticNotation")
	symaddr(&addr_QCborValue_toSimpleType2, "QCborValue_toSimpleType2")
	symaddr(&addr_QCborValue_toInteger2, "QCborValue_toInteger2")
	symaddr(&addr_QCborValue_toBool2, "QCborValue_toBool2")
	reg(&pg_QCborValue_toDouble2, "QCborValue_toDouble2")
	symaddr(&addr_QCborValue_tag2, "QCborValue_tag2")
	symaddr(&addr_QCborValue_taggedValue2, "QCborValue_taggedValue2")
	symaddr(&addr_QCborValue_toByteArray2, "QCborValue_toByteArray2")
	symaddr(&addr_QCborValue_toString2, "QCborValue_toString2")
	symaddr(&addr_QCborValue_toStringView2, "QCborValue_toStringView2")
	symaddr(&addr_QCborValue_toDateTime2, "QCborValue_toDateTime2")
	symaddr(&addr_QCborValue_toUrl2, "QCborValue_toUrl2")
	symaddr(&addr_QCborValue_toRegularExpression2, "QCborValue_toRegularExpression2")
	symaddr(&addr_QCborValue_toUuid2, "QCborValue_toUuid2")
	symaddr(&addr_QCborValue_fromCbor5, "QCborValue_fromCbor5")
	symaddr(&addr_QCborValue_fromCbor6, "QCborValue_fromCbor6")
	symaddr(&addr_QCborValue_fromCbor7, "QCborValue_fromCbor7")
	symaddr(&addr_QCborValue_toCbor3, "QCborValue_toCbor3")
	symaddr(&addr_QCborValue_toCbor4, "QCborValue_toCbor4")
	symaddr(&addr_QCborValue_toDiagnosticNotation2, "QCborValue_toDiagnosticNotation2")
	symaddr(&addr_QCborValue_OperatorEqual, "QCborValue_OperatorEqual")
	symaddr(&addr_QCborValue_OperatorNotEqual, "QCborValue_OperatorNotEqual")
	symaddr(&addr_QCborValue_OperatorLesser, "QCborValue_OperatorLesser")
	symaddr(&addr_QCborValue_OperatorGreater, "QCborValue_OperatorGreater")
	symaddr(&addr_QCborValue_OperatorLesserOrEqual, "QCborValue_OperatorLesserOrEqual")
	symaddr(&addr_QCborValue_OperatorGreaterOrEqual, "QCborValue_OperatorGreaterOrEqual")
	symaddr(&addr_QCborValue_OperatorEqual2, "QCborValue_OperatorEqual2")
	symaddr(&addr_QCborValue_OperatorNotEqual2, "QCborValue_OperatorNotEqual2")
	symaddr(&addr_QCborValue_OperatorLesser2, "QCborValue_OperatorLesser2")
	symaddr(&addr_QCborValue_OperatorGreater2, "QCborValue_OperatorGreater2")
	symaddr(&addr_QCborValue_OperatorLesserOrEqual2, "QCborValue_OperatorLesserOrEqual2")
	symaddr(&addr_QCborValue_OperatorGreaterOrEqual2, "QCborValue_OperatorGreaterOrEqual2")
	symaddr(&addr_QCborValue_OperatorEqual3, "QCborValue_OperatorEqual3")
	symaddr(&addr_QCborValue_OperatorNotEqual3, "QCborValue_OperatorNotEqual3")
	symaddr(&addr_QCborValue_OperatorLesser3, "QCborValue_OperatorLesser3")
	symaddr(&addr_QCborValue_OperatorGreater3, "QCborValue_OperatorGreater3")
	symaddr(&addr_QCborValue_OperatorLesserOrEqual3, "QCborValue_OperatorLesserOrEqual3")
	symaddr(&addr_QCborValue_OperatorGreaterOrEqual3, "QCborValue_OperatorGreaterOrEqual3")
	symaddr(&addr_QCborValue_OperatorEqual4, "QCborValue_OperatorEqual4")
	symaddr(&addr_QCborValue_OperatorNotEqual4, "QCborValue_OperatorNotEqual4")
	symaddr(&addr_QCborValue_OperatorLesser4, "QCborValue_OperatorLesser4")
	symaddr(&addr_QCborValue_OperatorGreater4, "QCborValue_OperatorGreater4")
	symaddr(&addr_QCborValue_OperatorLesserOrEqual4, "QCborValue_OperatorLesserOrEqual4")
	symaddr(&addr_QCborValue_OperatorGreaterOrEqual4, "QCborValue_OperatorGreaterOrEqual4")
	reg(&pg_QCborValue_delete, "QCborValue_delete")
}

type QCborValue struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborValue) UnsafeGoID() int64     { return this.vid }
func (this *QCborValue) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborValue(h uintptr) *QCborValue {
	if h == 0 {
		return nil
	}
	return &QCborValue{h: h}
}

func newownedQCborValue(h uintptr) *QCborValue {
	o := bindQCborValue(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborValue).Delete)
	}
	return o
}

func (this *QCborValue) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborValue) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborValue(h unsafe.Pointer) *QCborValue { return bindQCborValue(uintptr(h)) }
func UnsafeNewQCborValueOwned(h unsafe.Pointer) *QCborValue { return newownedQCborValue(uintptr(h)) }

func (this *QCborValue) UnsafeRelease() { this.h = 0 }

func NewQCborValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new)
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue2(t_ QCborValue__Type) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new2, uintptr(t_))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue3(b_ bool) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new3, uintptr(b2i(b_)))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue4(i int) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new4, uintptr(i))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue5(u uint) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new5, uintptr(u))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue6(i int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new6, uintptr(i))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue7(v float64) *QCborValue {
	__o := newownedQCborValue(pg_QCborValue_new7(v))
	return __o
}

func NewQCborValue8(st QCborSimpleType) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new8, uintptr(st))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue9(ba *QByteArray) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new9, uintptr(ba.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(ba)
	return __o
}

func NewQCborValue10(s string) *QCborValue {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValue_new10, s_p, uintptr(len(s)))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQCborValue11(s *QStringView) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new11, uintptr(s.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQCborValue12(s *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new12, uintptr(s.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQCborValue13(s string) *QCborValue {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QCborValue_new13, uintptr(unsafe.Pointer(&s_b[0])))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(s_b)
	return __o
}

func NewQCborValue14(a *QCborArray) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new14, uintptr(a.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(a)
	return __o
}

func NewQCborValue15(m *QCborMap) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new15, uintptr(m.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(m)
	return __o
}

func NewQCborValue16(tag QCborTag) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new16, uintptr(tag))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue17(t_ QCborKnownTags) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new17, uintptr(t_))
	__o := newownedQCborValue(__r)
	return __o
}

func NewQCborValue18(dt *QDateTime) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new18, uintptr(dt.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(dt)
	return __o
}

func NewQCborValue19(url *QUrl) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new19, uintptr(url.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(url)
	return __o
}

func NewQCborValue20(rx *QRegularExpression) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new20, uintptr(rx.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(rx)
	return __o
}

func NewQCborValue21(uuid *QUuid) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new21, uintptr(uuid.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(uuid)
	return __o
}

func NewQCborValue22(other *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new22, uintptr(other.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQCborValue23(tag QCborTag, taggedValue *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new23, uintptr(tag), uintptr(taggedValue.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(taggedValue)
	return __o
}

func NewQCborValue24(t_ QCborKnownTags, tv *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_new24, uintptr(t_), uintptr(tv.UnsafePointer()))
	__o := newownedQCborValue(__r)
	runtime.KeepAlive(tv)
	return __o
}

func (this *QCborValue) OperatorAssign(other *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCborValue(__r)
}

func (this *QCborValue) Swap(other *QCborValue) {
	purego.SyscallN(addr_QCborValue_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCborValue) Type() QCborValue__Type {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborValue__Type(__r)
}

func (this *QCborValue) IsInteger() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsByteArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isByteArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsMap() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isMap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsTag() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isTag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsFalse() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isFalse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsTrue() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isTrue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsDateTime() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isDateTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsUrl() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isUrl, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsRegularExpression() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isRegularExpression, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsUuid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isUuid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsInvalid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isInvalid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsContainer() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isContainer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsSimpleType() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_isSimpleType2, this.h, uintptr(st))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) ToSimpleType() QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_toSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborValue) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValue) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) ToDouble() float64 {
	__r := pg_QCborValue_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValue) Tag() QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_tag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborValue) TaggedValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_taggedValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValue) ToByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValue) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValue_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValue) ToStringView() *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toStringView, this.h)
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QCborValue) ToDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toDateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QCborValue) ToUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toUrl, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QCborValue) ToRegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toRegularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QCborValue) ToUuid() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toUuid, this.h)
	runtime.KeepAlive(this)
	return newownedQUuid(__r)
}

func (this *QCborValue) ToArray() *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toArray, this.h)
	runtime.KeepAlive(this)
	return newownedQCborArray(__r)
}

func (this *QCborValue) ToArrayWithDefaultValue(defaultValue *QCborArray) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toArray2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQCborArray(__r)
}

func (this *QCborValue) ToMap() *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toMap, this.h)
	runtime.KeepAlive(this)
	return newownedQCborMap(__r)
}

func (this *QCborValue) ToMapWithDefaultValue(defaultValue *QCborMap) *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toMap2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQCborMap(__r)
}

func (this *QCborValue) OperatorSubscript(key string) *QCborValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorSubscript, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborValue) OperatorSubscriptWithKey(key *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborValue) OperatorSubscript2(key int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorSubscript3, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborValue) OperatorSubscript3(key int64) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorSubscript4, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValueRef(__r)
}

func (this *QCborValue) OperatorSubscript4(key *QLatin1String) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorSubscript5, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborValue) OperatorSubscript5(key string) *QCborValueRef {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValue_OperatorSubscript6, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborValue) Compare(other *QCborValue) int {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func QCborValue_FromVariant(variant *QVariant) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromVariant, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(variant)
	return newownedQCborValue(__r)
}

func (this *QCborValue) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func QCborValue_FromJsonValue(v *QJsonValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromJsonValue, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(v)
	return newownedQCborValue(__r)
}

func (this *QCborValue) ToJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toJsonValue, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func QCborValue_FromCbor(reader *QCborStreamReader) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor, uintptr(reader.UnsafePointer()))
	runtime.KeepAlive(reader)
	return newownedQCborValue(__r)
}

func QCborValue_FromCborWithBa(ba *QByteArray) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor2, uintptr(ba.UnsafePointer()))
	runtime.KeepAlive(ba)
	return newownedQCborValue(__r)
}

func QCborValue_FromCbor2(data string, len_ int64) *QCborValue {
	data_b := append([]byte(data), 0)
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor3, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	runtime.KeepAlive(data_b)
	return newownedQCborValue(__r)
}

func QCborValue_FromCbor3(data []byte, len_ int64) *QCborValue {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor4, data_p, uintptr(len_))
	runtime.KeepAlive(data)
	return newownedQCborValue(__r)
}

func (this *QCborValue) ToCbor() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toCbor, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValue) ToCborWithWriter(writer *QCborStreamWriter) {
	purego.SyscallN(addr_QCborValue_toCbor2, this.h, uintptr(writer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(writer)
}

func (this *QCborValue) ToDiagnosticNotation() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValue_toDiagnosticNotation, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValue) ToSimpleTypeWithDefaultValue(defaultValue QCborSimpleType) QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_toSimpleType2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborValue) ToIntegerWithDefaultValue(defaultValue int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_toInteger2, this.h, uintptr(defaultValue))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValue) ToBoolWithDefaultValue(defaultValue bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_toBool2, this.h, uintptr(b2i(defaultValue)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborValue) ToDoubleWithDefaultValue(defaultValue float64) float64 {
	__r := pg_QCborValue_toDouble2(this.h, defaultValue)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborValue) TagWithDefaultValue(defaultValue QCborTag) QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_tag2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborValue) TaggedValueWithDefaultValue(defaultValue *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_taggedValue2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQCborValue(__r)
}

func (this *QCborValue) ToByteArrayWithDefaultValue(defaultValue *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toByteArray2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQByteArray(__r)
}

func (this *QCborValue) ToStringWithDefaultValue(defaultValue string) string {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValue_toString2, this.h, defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QCborValue) ToStringViewWithDefaultValue(defaultValue string) *QAnyStringView {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValue_toStringView2, this.h, defaultValue_p, uintptr(len(defaultValue)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQAnyStringView(__r)
}

func (this *QCborValue) ToDateTimeWithDefaultValue(defaultValue *QDateTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toDateTime2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQDateTime(__r)
}

func (this *QCborValue) ToUrlWithDefaultValue(defaultValue *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toUrl2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQUrl(__r)
}

func (this *QCborValue) ToRegularExpressionWithDefaultValue(defaultValue *QRegularExpression) *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toRegularExpression2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQRegularExpression(__r)
}

func (this *QCborValue) ToUuidWithDefaultValue(defaultValue *QUuid) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toUuid2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQUuid(__r)
}

func QCborValue_FromCbor4(ba *QByteArray, error_ *QCborParserError) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor5, uintptr(ba.UnsafePointer()), uintptr(error_.UnsafePointer()))
	runtime.KeepAlive(ba)
	runtime.KeepAlive(error_)
	return newownedQCborValue(__r)
}

func QCborValue_FromCbor5(data string, len_ int64, error_ *QCborParserError) *QCborValue {
	data_b := append([]byte(data), 0)
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor6, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_), uintptr(error_.UnsafePointer()))
	runtime.KeepAlive(data_b)
	runtime.KeepAlive(error_)
	return newownedQCborValue(__r)
}

func QCborValue_FromCbor6(data []byte, len_ int64, error_ *QCborParserError) *QCborValue {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCborValue_fromCbor7, data_p, uintptr(len_), uintptr(error_.UnsafePointer()))
	runtime.KeepAlive(data)
	runtime.KeepAlive(error_)
	return newownedQCborValue(__r)
}

func (this *QCborValue) ToCborWithOpt(opt QCborValue__EncodingOption) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborValue_toCbor3, this.h, uintptr(opt))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborValue) ToCbor2(writer *QCborStreamWriter, opt QCborValue__EncodingOption) {
	purego.SyscallN(addr_QCborValue_toCbor4, this.h, uintptr(writer.UnsafePointer()), uintptr(opt))
	runtime.KeepAlive(this)
	runtime.KeepAlive(writer)
}

func (this *QCborValue) ToDiagnosticNotationWithOpts(opts QCborValue__DiagnosticNotationOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborValue_toDiagnosticNotation2, this.h, uintptr(opts), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborValue) OperatorEqual(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorNotEqual(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesser(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreater(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesserOrEqual(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreaterOrEqual(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorEqualWithRhs(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorNotEqualWithRhs(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesserWithRhs(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreaterWithRhs(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesserOrEqualWithRhs(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreaterOrEqualWithRhs(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorNotEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesser2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreater2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesserOrEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreaterOrEqual2(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorNotEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesser3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesser4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreater3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreater4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorLesserOrEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorLesserOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) OperatorGreaterOrEqual3(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborValue_OperatorGreaterOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborValue) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborValue_delete(this.h)
	}
	this.UnsafeRelease()
}
