package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QJsonValueRef_base func(uintptr) uintptr
	addr_QJsonValueRef_new uintptr
	addr_QJsonValueRef_new2 uintptr
	addr_QJsonValueRef_new3 uintptr
	addr_QJsonValueRef_OperatorAssign uintptr
	addr_QJsonValueRef_OperatorAssign2 uintptr
	addr_QJsonValueRef_ToQJsonValue uintptr
	addr_QJsonValueRef_toVariant uintptr
	addr_QJsonValueRef_type uintptr
	addr_QJsonValueRef_isNull uintptr
	addr_QJsonValueRef_isBool uintptr
	addr_QJsonValueRef_isDouble uintptr
	addr_QJsonValueRef_isString uintptr
	addr_QJsonValueRef_isArray uintptr
	addr_QJsonValueRef_isObject uintptr
	addr_QJsonValueRef_isUndefined uintptr
	addr_QJsonValueRef_toBool uintptr
	addr_QJsonValueRef_toInt uintptr
	addr_QJsonValueRef_toInteger uintptr
	pg_QJsonValueRef_toDouble func(uintptr) float64
	addr_QJsonValueRef_toString uintptr
	addr_QJsonValueRef_toStringView uintptr
	addr_QJsonValueRef_toArray uintptr
	addr_QJsonValueRef_toObject uintptr
	addr_QJsonValueRef_OperatorSubscript uintptr
	addr_QJsonValueRef_OperatorSubscript2 uintptr
	addr_QJsonValueRef_OperatorSubscript3 uintptr
	addr_QJsonValueRef_toBool2 uintptr
	addr_QJsonValueRef_toInt2 uintptr
	addr_QJsonValueRef_toInteger2 uintptr
	pg_QJsonValueRef_toDouble2 func(uintptr, float64) float64
	addr_QJsonValueRef_toString2 uintptr
	addr_QJsonValueRef_toStringView2 uintptr
	addr_QJsonValueRef_OperatorEqual uintptr
	addr_QJsonValueRef_OperatorNotEqual uintptr
	addr_QJsonValueRef_OperatorEqual2 uintptr
	addr_QJsonValueRef_OperatorNotEqual2 uintptr
	pg_QJsonValueRef_delete func(uintptr)
)

func init() {
	reg(&pg_QJsonValueRef_base, "QJsonValueRef_base")
	symaddr(&addr_QJsonValueRef_new, "QJsonValueRef_new")
	symaddr(&addr_QJsonValueRef_new2, "QJsonValueRef_new2")
	symaddr(&addr_QJsonValueRef_new3, "QJsonValueRef_new3")
	symaddr(&addr_QJsonValueRef_OperatorAssign, "QJsonValueRef_OperatorAssign")
	symaddr(&addr_QJsonValueRef_OperatorAssign2, "QJsonValueRef_OperatorAssign2")
	symaddr(&addr_QJsonValueRef_ToQJsonValue, "QJsonValueRef_ToQJsonValue")
	symaddr(&addr_QJsonValueRef_toVariant, "QJsonValueRef_toVariant")
	symaddr(&addr_QJsonValueRef_type, "QJsonValueRef_type")
	symaddr(&addr_QJsonValueRef_isNull, "QJsonValueRef_isNull")
	symaddr(&addr_QJsonValueRef_isBool, "QJsonValueRef_isBool")
	symaddr(&addr_QJsonValueRef_isDouble, "QJsonValueRef_isDouble")
	symaddr(&addr_QJsonValueRef_isString, "QJsonValueRef_isString")
	symaddr(&addr_QJsonValueRef_isArray, "QJsonValueRef_isArray")
	symaddr(&addr_QJsonValueRef_isObject, "QJsonValueRef_isObject")
	symaddr(&addr_QJsonValueRef_isUndefined, "QJsonValueRef_isUndefined")
	symaddr(&addr_QJsonValueRef_toBool, "QJsonValueRef_toBool")
	symaddr(&addr_QJsonValueRef_toInt, "QJsonValueRef_toInt")
	symaddr(&addr_QJsonValueRef_toInteger, "QJsonValueRef_toInteger")
	reg(&pg_QJsonValueRef_toDouble, "QJsonValueRef_toDouble")
	symaddr(&addr_QJsonValueRef_toString, "QJsonValueRef_toString")
	symaddr(&addr_QJsonValueRef_toStringView, "QJsonValueRef_toStringView")
	symaddr(&addr_QJsonValueRef_toArray, "QJsonValueRef_toArray")
	symaddr(&addr_QJsonValueRef_toObject, "QJsonValueRef_toObject")
	symaddr(&addr_QJsonValueRef_OperatorSubscript, "QJsonValueRef_OperatorSubscript")
	symaddr(&addr_QJsonValueRef_OperatorSubscript2, "QJsonValueRef_OperatorSubscript2")
	symaddr(&addr_QJsonValueRef_OperatorSubscript3, "QJsonValueRef_OperatorSubscript3")
	symaddr(&addr_QJsonValueRef_toBool2, "QJsonValueRef_toBool2")
	symaddr(&addr_QJsonValueRef_toInt2, "QJsonValueRef_toInt2")
	symaddr(&addr_QJsonValueRef_toInteger2, "QJsonValueRef_toInteger2")
	reg(&pg_QJsonValueRef_toDouble2, "QJsonValueRef_toDouble2")
	symaddr(&addr_QJsonValueRef_toString2, "QJsonValueRef_toString2")
	symaddr(&addr_QJsonValueRef_toStringView2, "QJsonValueRef_toStringView2")
	symaddr(&addr_QJsonValueRef_OperatorEqual, "QJsonValueRef_OperatorEqual")
	symaddr(&addr_QJsonValueRef_OperatorNotEqual, "QJsonValueRef_OperatorNotEqual")
	symaddr(&addr_QJsonValueRef_OperatorEqual2, "QJsonValueRef_OperatorEqual2")
	symaddr(&addr_QJsonValueRef_OperatorNotEqual2, "QJsonValueRef_OperatorNotEqual2")
	reg(&pg_QJsonValueRef_delete, "QJsonValueRef_delete")
}

type QJsonValueRef struct {
	h uintptr
	*QJsonValueConstRef
}

func bindQJsonValueRef(h uintptr) *QJsonValueRef {
	if h == 0 {
		return nil
	}
	return &QJsonValueRef{h: h, QJsonValueConstRef: bindQJsonValueConstRef(pg_QJsonValueRef_base(h))}
}

func newownedQJsonValueRef(h uintptr) *QJsonValueRef {
	o := bindQJsonValueRef(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonValueRef).Delete)
	}
	return o
}

func (this *QJsonValueRef) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonValueRef) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonValueRef(h unsafe.Pointer) *QJsonValueRef { return bindQJsonValueRef(uintptr(h)) }
func UnsafeNewQJsonValueRefOwned(h unsafe.Pointer) *QJsonValueRef { return newownedQJsonValueRef(uintptr(h)) }

func (this *QJsonValueRef) UnsafeRelease() {
	this.h = 0
	if this.QJsonValueConstRef != nil {
		this.QJsonValueConstRef.UnsafeRelease()
	}
}

func NewQJsonValueRef(arg1 *QJsonValueRef) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQJsonValueRef(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQJsonValueRef2(array *QJsonArray, idx int64) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_new2, uintptr(array.UnsafePointer()), uintptr(idx))
	__o := newownedQJsonValueRef(__r)
	runtime.KeepAlive(array)
	return __o
}

func NewQJsonValueRef3(object *QJsonObject, idx int64) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_new3, uintptr(object.UnsafePointer()), uintptr(idx))
	__o := newownedQJsonValueRef(__r)
	runtime.KeepAlive(object)
	return __o
}

func (this *QJsonValueRef) OperatorAssign(val *QJsonValue) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorAssign, this.h, uintptr(val.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val)
	return bindQJsonValueRef(__r)
}

func (this *QJsonValueRef) OperatorAssignWithVal(val *QJsonValueRef) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorAssign2, this.h, uintptr(val.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val)
	return bindQJsonValueRef(__r)
}

func (this *QJsonValueRef) ToQJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_ToQJsonValue, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueRef) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QJsonValueRef) Type() QJsonValue__Type {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QJsonValue__Type(__r)
}

func (this *QJsonValueRef) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) IsObject() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isObject, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonValueRef) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueRef) ToDouble() float64 {
	__r := pg_QJsonValueRef_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueRef) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QJsonValueRef) ToStringView() *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toStringView, this.h)
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QJsonValueRef) ToArray() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toArray, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonArray(__r)
}

func (this *QJsonValueRef) ToObject() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toObject, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QJsonValueRef) OperatorSubscript(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorSubscript, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueRef) OperatorSubscriptWithKey(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueRef) OperatorSubscriptWith(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorSubscript3, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueRef) ToBoolWithDefaultValue(defaultValue bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_toBool2, this.h, uintptr(b2i(defaultValue)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueRef) ToIntWithDefaultValue(defaultValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_toInt2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonValueRef) ToIntegerWithDefaultValue(defaultValue int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_toInteger2, this.h, uintptr(defaultValue))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueRef) ToDoubleWithDefaultValue(defaultValue float64) float64 {
	__r := pg_QJsonValueRef_toDouble2(this.h, defaultValue)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueRef) ToStringWithDefaultValue(defaultValue string) string {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toString2, this.h, defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QJsonValueRef) ToStringViewWithDefaultValue(defaultValue string) *QAnyStringView {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonValueRef_toStringView2, this.h, defaultValue_p, uintptr(len(defaultValue)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQAnyStringView(__r)
}

func (this *QJsonValueRef) OperatorEqual(rhs *QJsonValueRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueRef) OperatorNotEqual(rhs *QJsonValueRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueRef) OperatorEqualWithRhs(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueRef) OperatorNotEqualWithRhs(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueRef_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueRef) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonValueRef_delete(this.h)
	}
	this.UnsafeRelease()
}
