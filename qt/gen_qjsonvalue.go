package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QJsonValue_new uintptr
	addr_QJsonValue_new2 uintptr
	pg_QJsonValue_new3 func(float64) uintptr
	addr_QJsonValue_new4 uintptr
	addr_QJsonValue_new5 uintptr
	addr_QJsonValue_new6 uintptr
	addr_QJsonValue_new7 uintptr
	addr_QJsonValue_new8 uintptr
	addr_QJsonValue_new9 uintptr
	addr_QJsonValue_new10 uintptr
	addr_QJsonValue_new11 uintptr
	addr_QJsonValue_new12 uintptr
	addr_QJsonValue_OperatorAssign uintptr
	addr_QJsonValue_swap uintptr
	addr_QJsonValue_fromVariant uintptr
	addr_QJsonValue_toVariant uintptr
	addr_QJsonValue_fromJson uintptr
	addr_QJsonValue_toJson uintptr
	addr_QJsonValue_type uintptr
	addr_QJsonValue_isNull uintptr
	addr_QJsonValue_isBool uintptr
	addr_QJsonValue_isDouble uintptr
	addr_QJsonValue_isString uintptr
	addr_QJsonValue_isArray uintptr
	addr_QJsonValue_isObject uintptr
	addr_QJsonValue_isUndefined uintptr
	addr_QJsonValue_toBool uintptr
	addr_QJsonValue_toInt uintptr
	addr_QJsonValue_toInteger uintptr
	pg_QJsonValue_toDouble func(uintptr) float64
	addr_QJsonValue_toString uintptr
	addr_QJsonValue_toString2 uintptr
	addr_QJsonValue_toStringView uintptr
	addr_QJsonValue_toArray uintptr
	addr_QJsonValue_toArray2 uintptr
	addr_QJsonValue_toObject uintptr
	addr_QJsonValue_toObject2 uintptr
	addr_QJsonValue_OperatorSubscript uintptr
	addr_QJsonValue_OperatorSubscript2 uintptr
	addr_QJsonValue_OperatorSubscript3 uintptr
	addr_QJsonValue_OperatorSubscript4 uintptr
	addr_QJsonValue_fromJson2 uintptr
	addr_QJsonValue_toJson2 uintptr
	addr_QJsonValue_toBool2 uintptr
	addr_QJsonValue_toInt2 uintptr
	addr_QJsonValue_toInteger2 uintptr
	pg_QJsonValue_toDouble2 func(uintptr, float64) float64
	addr_QJsonValue_toStringView2 uintptr
	addr_QJsonValue_OperatorEqual uintptr
	addr_QJsonValue_OperatorNotEqual uintptr
	addr_QJsonValue_OperatorEqual2 uintptr
	addr_QJsonValue_OperatorNotEqual2 uintptr
	addr_QJsonValue_OperatorEqual3 uintptr
	addr_QJsonValue_OperatorNotEqual3 uintptr
	addr_QJsonValue_OperatorEqual4 uintptr
	addr_QJsonValue_OperatorNotEqual4 uintptr
	pg_QJsonValue_delete func(uintptr)
)

func init() {
	symaddr(&addr_QJsonValue_new, "QJsonValue_new")
	symaddr(&addr_QJsonValue_new2, "QJsonValue_new2")
	reg(&pg_QJsonValue_new3, "QJsonValue_new3")
	symaddr(&addr_QJsonValue_new4, "QJsonValue_new4")
	symaddr(&addr_QJsonValue_new5, "QJsonValue_new5")
	symaddr(&addr_QJsonValue_new6, "QJsonValue_new6")
	symaddr(&addr_QJsonValue_new7, "QJsonValue_new7")
	symaddr(&addr_QJsonValue_new8, "QJsonValue_new8")
	symaddr(&addr_QJsonValue_new9, "QJsonValue_new9")
	symaddr(&addr_QJsonValue_new10, "QJsonValue_new10")
	symaddr(&addr_QJsonValue_new11, "QJsonValue_new11")
	symaddr(&addr_QJsonValue_new12, "QJsonValue_new12")
	symaddr(&addr_QJsonValue_OperatorAssign, "QJsonValue_OperatorAssign")
	symaddr(&addr_QJsonValue_swap, "QJsonValue_swap")
	symaddr(&addr_QJsonValue_fromVariant, "QJsonValue_fromVariant")
	symaddr(&addr_QJsonValue_toVariant, "QJsonValue_toVariant")
	symaddr(&addr_QJsonValue_fromJson, "QJsonValue_fromJson")
	symaddr(&addr_QJsonValue_toJson, "QJsonValue_toJson")
	symaddr(&addr_QJsonValue_type, "QJsonValue_type")
	symaddr(&addr_QJsonValue_isNull, "QJsonValue_isNull")
	symaddr(&addr_QJsonValue_isBool, "QJsonValue_isBool")
	symaddr(&addr_QJsonValue_isDouble, "QJsonValue_isDouble")
	symaddr(&addr_QJsonValue_isString, "QJsonValue_isString")
	symaddr(&addr_QJsonValue_isArray, "QJsonValue_isArray")
	symaddr(&addr_QJsonValue_isObject, "QJsonValue_isObject")
	symaddr(&addr_QJsonValue_isUndefined, "QJsonValue_isUndefined")
	symaddr(&addr_QJsonValue_toBool, "QJsonValue_toBool")
	symaddr(&addr_QJsonValue_toInt, "QJsonValue_toInt")
	symaddr(&addr_QJsonValue_toInteger, "QJsonValue_toInteger")
	reg(&pg_QJsonValue_toDouble, "QJsonValue_toDouble")
	symaddr(&addr_QJsonValue_toString, "QJsonValue_toString")
	symaddr(&addr_QJsonValue_toString2, "QJsonValue_toString2")
	symaddr(&addr_QJsonValue_toStringView, "QJsonValue_toStringView")
	symaddr(&addr_QJsonValue_toArray, "QJsonValue_toArray")
	symaddr(&addr_QJsonValue_toArray2, "QJsonValue_toArray2")
	symaddr(&addr_QJsonValue_toObject, "QJsonValue_toObject")
	symaddr(&addr_QJsonValue_toObject2, "QJsonValue_toObject2")
	symaddr(&addr_QJsonValue_OperatorSubscript, "QJsonValue_OperatorSubscript")
	symaddr(&addr_QJsonValue_OperatorSubscript2, "QJsonValue_OperatorSubscript2")
	symaddr(&addr_QJsonValue_OperatorSubscript3, "QJsonValue_OperatorSubscript3")
	symaddr(&addr_QJsonValue_OperatorSubscript4, "QJsonValue_OperatorSubscript4")
	symaddr(&addr_QJsonValue_fromJson2, "QJsonValue_fromJson2")
	symaddr(&addr_QJsonValue_toJson2, "QJsonValue_toJson2")
	symaddr(&addr_QJsonValue_toBool2, "QJsonValue_toBool2")
	symaddr(&addr_QJsonValue_toInt2, "QJsonValue_toInt2")
	symaddr(&addr_QJsonValue_toInteger2, "QJsonValue_toInteger2")
	reg(&pg_QJsonValue_toDouble2, "QJsonValue_toDouble2")
	symaddr(&addr_QJsonValue_toStringView2, "QJsonValue_toStringView2")
	symaddr(&addr_QJsonValue_OperatorEqual, "QJsonValue_OperatorEqual")
	symaddr(&addr_QJsonValue_OperatorNotEqual, "QJsonValue_OperatorNotEqual")
	symaddr(&addr_QJsonValue_OperatorEqual2, "QJsonValue_OperatorEqual2")
	symaddr(&addr_QJsonValue_OperatorNotEqual2, "QJsonValue_OperatorNotEqual2")
	symaddr(&addr_QJsonValue_OperatorEqual3, "QJsonValue_OperatorEqual3")
	symaddr(&addr_QJsonValue_OperatorNotEqual3, "QJsonValue_OperatorNotEqual3")
	symaddr(&addr_QJsonValue_OperatorEqual4, "QJsonValue_OperatorEqual4")
	symaddr(&addr_QJsonValue_OperatorNotEqual4, "QJsonValue_OperatorNotEqual4")
	reg(&pg_QJsonValue_delete, "QJsonValue_delete")
}

type QJsonValue struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QJsonValue) UnsafeGoID() int64     { return this.vid }
func (this *QJsonValue) UnsafeSetGoID(id int64) { this.vid = id }

func bindQJsonValue(h uintptr) *QJsonValue {
	if h == 0 {
		return nil
	}
	return &QJsonValue{h: h}
}

func newownedQJsonValue(h uintptr) *QJsonValue {
	o := bindQJsonValue(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonValue).Delete)
	}
	return o
}

func (this *QJsonValue) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonValue) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonValue(h unsafe.Pointer) *QJsonValue { return bindQJsonValue(uintptr(h)) }
func UnsafeNewQJsonValueOwned(h unsafe.Pointer) *QJsonValue { return newownedQJsonValue(uintptr(h)) }

func (this *QJsonValue) UnsafeRelease() { this.h = 0 }

func NewQJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new)
	__o := newownedQJsonValue(__r)
	return __o
}

func NewQJsonValue2(b bool) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new2, uintptr(b2i(b)))
	__o := newownedQJsonValue(__r)
	return __o
}

func NewQJsonValue3(n float64) *QJsonValue {
	__o := newownedQJsonValue(pg_QJsonValue_new3(n))
	return __o
}

func NewQJsonValue4(n int) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new4, uintptr(n))
	__o := newownedQJsonValue(__r)
	return __o
}

func NewQJsonValue5(v int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new5, uintptr(v))
	__o := newownedQJsonValue(__r)
	return __o
}

func NewQJsonValue6(s string) *QJsonValue {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new6, s_p, uintptr(len(s)))
	__o := newownedQJsonValue(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQJsonValue7(s *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new7, uintptr(s.UnsafePointer()))
	__o := newownedQJsonValue(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQJsonValue8(s string) *QJsonValue {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new8, uintptr(unsafe.Pointer(&s_b[0])))
	__o := newownedQJsonValue(__r)
	runtime.KeepAlive(s_b)
	return __o
}

func NewQJsonValue9(a *QJsonArray) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new9, uintptr(a.UnsafePointer()))
	__o := newownedQJsonValue(__r)
	runtime.KeepAlive(a)
	return __o
}

func NewQJsonValue10(o *QJsonObject) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new10, uintptr(o.UnsafePointer()))
	__o := newownedQJsonValue(__r)
	runtime.KeepAlive(o)
	return __o
}

func NewQJsonValue11(other *QJsonValue) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new11, uintptr(other.UnsafePointer()))
	__o := newownedQJsonValue(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQJsonValue12(arg1 QJsonValue__Type) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_new12, uintptr(arg1))
	__o := newownedQJsonValue(__r)
	return __o
}

func (this *QJsonValue) OperatorAssign(other *QJsonValue) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQJsonValue(__r)
}

func (this *QJsonValue) Swap(other *QJsonValue) {
	purego.SyscallN(addr_QJsonValue_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func QJsonValue_FromVariant(variant *QVariant) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_fromVariant, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(variant)
	return newownedQJsonValue(__r)
}

func (this *QJsonValue) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func QJsonValue_FromJson(json *QByteArrayView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_fromJson, uintptr(json.UnsafePointer()))
	runtime.KeepAlive(json)
	return newownedQJsonValue(__r)
}

func (this *QJsonValue) ToJson() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toJson, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QJsonValue) Type() QJsonValue__Type {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QJsonValue__Type(__r)
}

func (this *QJsonValue) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) IsObject() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isObject, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonValue) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValue) ToDouble() float64 {
	__r := pg_QJsonValue_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValue) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QJsonValue) ToStringWithDefaultValue(defaultValue string) string {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toString2, this.h, defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QJsonValue) ToStringView() *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toStringView, this.h)
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QJsonValue) ToArray() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toArray, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonArray(__r)
}

func (this *QJsonValue) ToArrayWithDefaultValue(defaultValue *QJsonArray) *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toArray2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQJsonArray(__r)
}

func (this *QJsonValue) ToObject() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toObject, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QJsonValue) ToObjectWithDefaultValue(defaultValue *QJsonObject) *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toObject2, this.h, uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQJsonObject(__r)
}

func (this *QJsonValue) OperatorSubscript(key string) *QJsonValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonValue_OperatorSubscript, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValue) OperatorSubscriptWithKey(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValue) OperatorSubscript2(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_OperatorSubscript3, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValue) OperatorSubscriptWith(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_OperatorSubscript4, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func QJsonValue_FromJson2(json *QByteArrayView, error_ *QJsonParseError) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_fromJson2, uintptr(json.UnsafePointer()), uintptr(error_.UnsafePointer()))
	runtime.KeepAlive(json)
	runtime.KeepAlive(error_)
	return newownedQJsonValue(__r)
}

func (this *QJsonValue) ToJsonWithFormat(format QJsonDocument__JsonFormat) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toJson2, this.h, uintptr(format))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QJsonValue) ToBoolWithDefaultValue(defaultValue bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_toBool2, this.h, uintptr(b2i(defaultValue)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValue) ToIntWithDefaultValue(defaultValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_toInt2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonValue) ToIntegerWithDefaultValue(defaultValue int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_toInteger2, this.h, uintptr(defaultValue))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValue) ToDoubleWithDefaultValue(defaultValue float64) float64 {
	__r := pg_QJsonValue_toDouble2(this.h, defaultValue)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValue) ToStringViewWithDefaultValue(defaultValue string) *QAnyStringView {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonValue_toStringView2, this.h, defaultValue_p, uintptr(len(defaultValue)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQAnyStringView(__r)
}

func (this *QJsonValue) OperatorEqual(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorNotEqual(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorEqualWithRhs(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorNotEqualWithRhs(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorEqual2(rhs *QJsonArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorNotEqual2(rhs *QJsonArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorEqual3(rhs *QJsonObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) OperatorNotEqual3(rhs *QJsonObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValue_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValue) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonValue_delete(this.h)
	}
	this.UnsafeRelease()
}
