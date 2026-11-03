package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QJsonValueConstRef_new uintptr
	addr_QJsonValueConstRef_ToQJsonValue uintptr
	addr_QJsonValueConstRef_toVariant uintptr
	addr_QJsonValueConstRef_type uintptr
	addr_QJsonValueConstRef_isNull uintptr
	addr_QJsonValueConstRef_isBool uintptr
	addr_QJsonValueConstRef_isDouble uintptr
	addr_QJsonValueConstRef_isString uintptr
	addr_QJsonValueConstRef_isArray uintptr
	addr_QJsonValueConstRef_isObject uintptr
	addr_QJsonValueConstRef_isUndefined uintptr
	addr_QJsonValueConstRef_toBool uintptr
	addr_QJsonValueConstRef_toInt uintptr
	addr_QJsonValueConstRef_toInteger uintptr
	pg_QJsonValueConstRef_toDouble func(uintptr) float64
	addr_QJsonValueConstRef_toString uintptr
	addr_QJsonValueConstRef_toStringView uintptr
	addr_QJsonValueConstRef_toArray uintptr
	addr_QJsonValueConstRef_toObject uintptr
	addr_QJsonValueConstRef_OperatorSubscript uintptr
	addr_QJsonValueConstRef_OperatorSubscript2 uintptr
	addr_QJsonValueConstRef_OperatorSubscript3 uintptr
	addr_QJsonValueConstRef_toBool2 uintptr
	addr_QJsonValueConstRef_toInt2 uintptr
	addr_QJsonValueConstRef_toInteger2 uintptr
	pg_QJsonValueConstRef_toDouble2 func(uintptr, float64) float64
	addr_QJsonValueConstRef_toString2 uintptr
	addr_QJsonValueConstRef_toStringView2 uintptr
	addr_QJsonValueConstRef_OperatorEqual uintptr
	addr_QJsonValueConstRef_OperatorNotEqual uintptr
	addr_QJsonValueConstRef_OperatorEqual2 uintptr
	addr_QJsonValueConstRef_OperatorNotEqual2 uintptr
	addr_QJsonValueConstRef_OperatorEqual3 uintptr
	addr_QJsonValueConstRef_OperatorNotEqual3 uintptr
	addr_QJsonValueConstRef_OperatorEqual4 uintptr
	addr_QJsonValueConstRef_OperatorNotEqual4 uintptr
	pg_QJsonValueConstRef_delete func(uintptr)
)

func init() {
	symaddr(&addr_QJsonValueConstRef_new, "QJsonValueConstRef_new")
	symaddr(&addr_QJsonValueConstRef_ToQJsonValue, "QJsonValueConstRef_ToQJsonValue")
	symaddr(&addr_QJsonValueConstRef_toVariant, "QJsonValueConstRef_toVariant")
	symaddr(&addr_QJsonValueConstRef_type, "QJsonValueConstRef_type")
	symaddr(&addr_QJsonValueConstRef_isNull, "QJsonValueConstRef_isNull")
	symaddr(&addr_QJsonValueConstRef_isBool, "QJsonValueConstRef_isBool")
	symaddr(&addr_QJsonValueConstRef_isDouble, "QJsonValueConstRef_isDouble")
	symaddr(&addr_QJsonValueConstRef_isString, "QJsonValueConstRef_isString")
	symaddr(&addr_QJsonValueConstRef_isArray, "QJsonValueConstRef_isArray")
	symaddr(&addr_QJsonValueConstRef_isObject, "QJsonValueConstRef_isObject")
	symaddr(&addr_QJsonValueConstRef_isUndefined, "QJsonValueConstRef_isUndefined")
	symaddr(&addr_QJsonValueConstRef_toBool, "QJsonValueConstRef_toBool")
	symaddr(&addr_QJsonValueConstRef_toInt, "QJsonValueConstRef_toInt")
	symaddr(&addr_QJsonValueConstRef_toInteger, "QJsonValueConstRef_toInteger")
	reg(&pg_QJsonValueConstRef_toDouble, "QJsonValueConstRef_toDouble")
	symaddr(&addr_QJsonValueConstRef_toString, "QJsonValueConstRef_toString")
	symaddr(&addr_QJsonValueConstRef_toStringView, "QJsonValueConstRef_toStringView")
	symaddr(&addr_QJsonValueConstRef_toArray, "QJsonValueConstRef_toArray")
	symaddr(&addr_QJsonValueConstRef_toObject, "QJsonValueConstRef_toObject")
	symaddr(&addr_QJsonValueConstRef_OperatorSubscript, "QJsonValueConstRef_OperatorSubscript")
	symaddr(&addr_QJsonValueConstRef_OperatorSubscript2, "QJsonValueConstRef_OperatorSubscript2")
	symaddr(&addr_QJsonValueConstRef_OperatorSubscript3, "QJsonValueConstRef_OperatorSubscript3")
	symaddr(&addr_QJsonValueConstRef_toBool2, "QJsonValueConstRef_toBool2")
	symaddr(&addr_QJsonValueConstRef_toInt2, "QJsonValueConstRef_toInt2")
	symaddr(&addr_QJsonValueConstRef_toInteger2, "QJsonValueConstRef_toInteger2")
	reg(&pg_QJsonValueConstRef_toDouble2, "QJsonValueConstRef_toDouble2")
	symaddr(&addr_QJsonValueConstRef_toString2, "QJsonValueConstRef_toString2")
	symaddr(&addr_QJsonValueConstRef_toStringView2, "QJsonValueConstRef_toStringView2")
	symaddr(&addr_QJsonValueConstRef_OperatorEqual, "QJsonValueConstRef_OperatorEqual")
	symaddr(&addr_QJsonValueConstRef_OperatorNotEqual, "QJsonValueConstRef_OperatorNotEqual")
	symaddr(&addr_QJsonValueConstRef_OperatorEqual2, "QJsonValueConstRef_OperatorEqual2")
	symaddr(&addr_QJsonValueConstRef_OperatorNotEqual2, "QJsonValueConstRef_OperatorNotEqual2")
	symaddr(&addr_QJsonValueConstRef_OperatorEqual3, "QJsonValueConstRef_OperatorEqual3")
	symaddr(&addr_QJsonValueConstRef_OperatorNotEqual3, "QJsonValueConstRef_OperatorNotEqual3")
	symaddr(&addr_QJsonValueConstRef_OperatorEqual4, "QJsonValueConstRef_OperatorEqual4")
	symaddr(&addr_QJsonValueConstRef_OperatorNotEqual4, "QJsonValueConstRef_OperatorNotEqual4")
	reg(&pg_QJsonValueConstRef_delete, "QJsonValueConstRef_delete")
}

type QJsonValueConstRef struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QJsonValueConstRef) UnsafeGoID() int64     { return this.vid }
func (this *QJsonValueConstRef) UnsafeSetGoID(id int64) { this.vid = id }

func bindQJsonValueConstRef(h uintptr) *QJsonValueConstRef {
	if h == 0 {
		return nil
	}
	return &QJsonValueConstRef{h: h}
}

func newownedQJsonValueConstRef(h uintptr) *QJsonValueConstRef {
	o := bindQJsonValueConstRef(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonValueConstRef).Delete)
	}
	return o
}

func (this *QJsonValueConstRef) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonValueConstRef) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonValueConstRef(h unsafe.Pointer) *QJsonValueConstRef { return bindQJsonValueConstRef(uintptr(h)) }
func UnsafeNewQJsonValueConstRefOwned(h unsafe.Pointer) *QJsonValueConstRef { return newownedQJsonValueConstRef(uintptr(h)) }

func (this *QJsonValueConstRef) UnsafeRelease() { this.h = 0 }

func NewQJsonValueConstRef(arg1 *QJsonValueConstRef) *QJsonValueConstRef {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQJsonValueConstRef(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QJsonValueConstRef) ToQJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_ToQJsonValue, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueConstRef) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QJsonValueConstRef) Type() QJsonValue__Type {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QJsonValue__Type(__r)
}

func (this *QJsonValueConstRef) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) IsObject() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isObject, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonValueConstRef) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueConstRef) ToDouble() float64 {
	__r := pg_QJsonValueConstRef_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueConstRef) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QJsonValueConstRef) ToStringView() *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toStringView, this.h)
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QJsonValueConstRef) ToArray() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toArray, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonArray(__r)
}

func (this *QJsonValueConstRef) ToObject() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toObject, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QJsonValueConstRef) OperatorSubscript(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorSubscript, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueConstRef) OperatorSubscriptWithKey(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueConstRef) OperatorSubscriptWith(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorSubscript3, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonValueConstRef) ToBoolWithDefaultValue(defaultValue bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toBool2, this.h, uintptr(b2i(defaultValue)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonValueConstRef) ToIntWithDefaultValue(defaultValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toInt2, this.h, uintptr(defaultValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonValueConstRef) ToIntegerWithDefaultValue(defaultValue int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toInteger2, this.h, uintptr(defaultValue))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueConstRef) ToDoubleWithDefaultValue(defaultValue float64) float64 {
	__r := pg_QJsonValueConstRef_toDouble2(this.h, defaultValue)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonValueConstRef) ToStringWithDefaultValue(defaultValue string) string {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toString2, this.h, defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QJsonValueConstRef) ToStringViewWithDefaultValue(defaultValue string) *QAnyStringView {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonValueConstRef_toStringView2, this.h, defaultValue_p, uintptr(len(defaultValue)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQAnyStringView(__r)
}

func (this *QJsonValueConstRef) OperatorEqual(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorNotEqual(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorEqualWithRhs(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorNotEqualWithRhs(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorEqual2(rhs *QJsonValueRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorNotEqual2(rhs *QJsonValueRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorEqual3(rhs *QJsonObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) OperatorNotEqual3(rhs *QJsonObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonValueConstRef_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonValueConstRef) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonValueConstRef_delete(this.h)
	}
	this.UnsafeRelease()
}
