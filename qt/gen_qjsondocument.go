package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QJsonDocument_new uintptr
	addr_QJsonDocument_new2 uintptr
	addr_QJsonDocument_new3 uintptr
	addr_QJsonDocument_new4 uintptr
	addr_QJsonDocument_OperatorAssign uintptr
	addr_QJsonDocument_swap uintptr
	addr_QJsonDocument_fromVariant uintptr
	addr_QJsonDocument_toVariant uintptr
	addr_QJsonDocument_fromJson uintptr
	addr_QJsonDocument_toJson uintptr
	addr_QJsonDocument_isEmpty uintptr
	addr_QJsonDocument_isArray uintptr
	addr_QJsonDocument_isObject uintptr
	addr_QJsonDocument_object uintptr
	addr_QJsonDocument_array uintptr
	addr_QJsonDocument_setObject uintptr
	addr_QJsonDocument_setArray uintptr
	addr_QJsonDocument_OperatorSubscript uintptr
	addr_QJsonDocument_OperatorSubscript2 uintptr
	addr_QJsonDocument_OperatorSubscript3 uintptr
	addr_QJsonDocument_OperatorSubscript4 uintptr
	addr_QJsonDocument_isNull uintptr
	addr_QJsonDocument_fromJson2 uintptr
	addr_QJsonDocument_toJson2 uintptr
	addr_QJsonDocument_OperatorEqual uintptr
	addr_QJsonDocument_OperatorNotEqual uintptr
	pg_QJsonDocument_delete func(uintptr)
)

func init() {
	symaddr(&addr_QJsonDocument_new, "QJsonDocument_new")
	symaddr(&addr_QJsonDocument_new2, "QJsonDocument_new2")
	symaddr(&addr_QJsonDocument_new3, "QJsonDocument_new3")
	symaddr(&addr_QJsonDocument_new4, "QJsonDocument_new4")
	symaddr(&addr_QJsonDocument_OperatorAssign, "QJsonDocument_OperatorAssign")
	symaddr(&addr_QJsonDocument_swap, "QJsonDocument_swap")
	symaddr(&addr_QJsonDocument_fromVariant, "QJsonDocument_fromVariant")
	symaddr(&addr_QJsonDocument_toVariant, "QJsonDocument_toVariant")
	symaddr(&addr_QJsonDocument_fromJson, "QJsonDocument_fromJson")
	symaddr(&addr_QJsonDocument_toJson, "QJsonDocument_toJson")
	symaddr(&addr_QJsonDocument_isEmpty, "QJsonDocument_isEmpty")
	symaddr(&addr_QJsonDocument_isArray, "QJsonDocument_isArray")
	symaddr(&addr_QJsonDocument_isObject, "QJsonDocument_isObject")
	symaddr(&addr_QJsonDocument_object, "QJsonDocument_object")
	symaddr(&addr_QJsonDocument_array, "QJsonDocument_array")
	symaddr(&addr_QJsonDocument_setObject, "QJsonDocument_setObject")
	symaddr(&addr_QJsonDocument_setArray, "QJsonDocument_setArray")
	symaddr(&addr_QJsonDocument_OperatorSubscript, "QJsonDocument_OperatorSubscript")
	symaddr(&addr_QJsonDocument_OperatorSubscript2, "QJsonDocument_OperatorSubscript2")
	symaddr(&addr_QJsonDocument_OperatorSubscript3, "QJsonDocument_OperatorSubscript3")
	symaddr(&addr_QJsonDocument_OperatorSubscript4, "QJsonDocument_OperatorSubscript4")
	symaddr(&addr_QJsonDocument_isNull, "QJsonDocument_isNull")
	symaddr(&addr_QJsonDocument_fromJson2, "QJsonDocument_fromJson2")
	symaddr(&addr_QJsonDocument_toJson2, "QJsonDocument_toJson2")
	symaddr(&addr_QJsonDocument_OperatorEqual, "QJsonDocument_OperatorEqual")
	symaddr(&addr_QJsonDocument_OperatorNotEqual, "QJsonDocument_OperatorNotEqual")
	reg(&pg_QJsonDocument_delete, "QJsonDocument_delete")
}

type QJsonDocument struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QJsonDocument) UnsafeGoID() int64     { return this.vid }
func (this *QJsonDocument) UnsafeSetGoID(id int64) { this.vid = id }

func bindQJsonDocument(h uintptr) *QJsonDocument {
	if h == 0 {
		return nil
	}
	return &QJsonDocument{h: h}
}

func newownedQJsonDocument(h uintptr) *QJsonDocument {
	o := bindQJsonDocument(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonDocument).Delete)
	}
	return o
}

func (this *QJsonDocument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonDocument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonDocument(h unsafe.Pointer) *QJsonDocument { return bindQJsonDocument(uintptr(h)) }
func UnsafeNewQJsonDocumentOwned(h unsafe.Pointer) *QJsonDocument { return newownedQJsonDocument(uintptr(h)) }

func (this *QJsonDocument) UnsafeRelease() { this.h = 0 }

func NewQJsonDocument() *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_new)
	__o := newownedQJsonDocument(__r)
	return __o
}

func NewQJsonDocument2(object *QJsonObject) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_new2, uintptr(object.UnsafePointer()))
	__o := newownedQJsonDocument(__r)
	runtime.KeepAlive(object)
	return __o
}

func NewQJsonDocument3(array *QJsonArray) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_new3, uintptr(array.UnsafePointer()))
	__o := newownedQJsonDocument(__r)
	runtime.KeepAlive(array)
	return __o
}

func NewQJsonDocument4(other *QJsonDocument) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_new4, uintptr(other.UnsafePointer()))
	__o := newownedQJsonDocument(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QJsonDocument) OperatorAssign(other *QJsonDocument) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQJsonDocument(__r)
}

func (this *QJsonDocument) Swap(other *QJsonDocument) {
	purego.SyscallN(addr_QJsonDocument_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func QJsonDocument_FromVariant(variant *QVariant) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_fromVariant, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(variant)
	return newownedQJsonDocument(__r)
}

func (this *QJsonDocument) ToVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_toVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func QJsonDocument_FromJson(json *QByteArray) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_fromJson, uintptr(json.UnsafePointer()))
	runtime.KeepAlive(json)
	return newownedQJsonDocument(__r)
}

func (this *QJsonDocument) ToJson() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_toJson, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QJsonDocument) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonDocument_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonDocument) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonDocument_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonDocument) IsObject() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonDocument_isObject, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonDocument) Object() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_object, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QJsonDocument) Array() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_array, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonArray(__r)
}

func (this *QJsonDocument) SetObject(object *QJsonObject) {
	purego.SyscallN(addr_QJsonDocument_setObject, this.h, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
}

func (this *QJsonDocument) SetArray(array *QJsonArray) {
	purego.SyscallN(addr_QJsonDocument_setArray, this.h, uintptr(array.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(array)
}

func (this *QJsonDocument) OperatorSubscript(key string) *QJsonValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorSubscript, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonDocument) OperatorSubscriptWithKey(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonDocument) OperatorSubscript2(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorSubscript3, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonDocument) OperatorSubscriptWith(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorSubscript4, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonDocument) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonDocument_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QJsonDocument_FromJson2(json *QByteArray, error_ *QJsonParseError) *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_fromJson2, uintptr(json.UnsafePointer()), uintptr(error_.UnsafePointer()))
	runtime.KeepAlive(json)
	runtime.KeepAlive(error_)
	return newownedQJsonDocument(__r)
}

func (this *QJsonDocument) ToJsonWithFormat(format QJsonDocument__JsonFormat) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QJsonDocument_toJson2, this.h, uintptr(format))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QJsonDocument) OperatorEqual(rhs *QJsonDocument) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonDocument) OperatorNotEqual(rhs *QJsonDocument) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonDocument_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonDocument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonDocument_delete(this.h)
	}
	this.UnsafeRelease()
}
