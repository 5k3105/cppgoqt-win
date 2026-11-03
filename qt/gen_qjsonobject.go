package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): QJsonObject_FromVariantMap, QJsonObject_FromVariantHash

var (
	addr_QJsonObject_new uintptr
	addr_QJsonObject_new2 uintptr
	addr_QJsonObject_OperatorAssign uintptr
	addr_QJsonObject_swap uintptr
	addr_QJsonObject_keys uintptr
	addr_QJsonObject_size uintptr
	addr_QJsonObject_count uintptr
	addr_QJsonObject_length uintptr
	addr_QJsonObject_isEmpty uintptr
	addr_QJsonObject_value uintptr
	addr_QJsonObject_OperatorSubscript uintptr
	addr_QJsonObject_OperatorSubscript2 uintptr
	addr_QJsonObject_value2 uintptr
	addr_QJsonObject_value3 uintptr
	addr_QJsonObject_OperatorSubscript3 uintptr
	addr_QJsonObject_OperatorSubscript4 uintptr
	addr_QJsonObject_OperatorSubscript5 uintptr
	addr_QJsonObject_OperatorSubscript6 uintptr
	addr_QJsonObject_remove uintptr
	addr_QJsonObject_take uintptr
	addr_QJsonObject_contains uintptr
	addr_QJsonObject_remove2 uintptr
	addr_QJsonObject_remove3 uintptr
	addr_QJsonObject_take2 uintptr
	addr_QJsonObject_take3 uintptr
	addr_QJsonObject_contains2 uintptr
	addr_QJsonObject_contains3 uintptr
	addr_QJsonObject_empty uintptr
	addr_QJsonObject_OperatorEqual uintptr
	addr_QJsonObject_OperatorNotEqual uintptr
	addr_QJsonObject_OperatorEqual2 uintptr
	addr_QJsonObject_OperatorNotEqual2 uintptr
	addr_QJsonObject_OperatorEqual3 uintptr
	addr_QJsonObject_OperatorNotEqual3 uintptr
	pg_QJsonObject_delete func(uintptr)
)

func init() {
	symaddr(&addr_QJsonObject_new, "QJsonObject_new")
	symaddr(&addr_QJsonObject_new2, "QJsonObject_new2")
	symaddr(&addr_QJsonObject_OperatorAssign, "QJsonObject_OperatorAssign")
	symaddr(&addr_QJsonObject_swap, "QJsonObject_swap")
	symaddr(&addr_QJsonObject_keys, "QJsonObject_keys")
	symaddr(&addr_QJsonObject_size, "QJsonObject_size")
	symaddr(&addr_QJsonObject_count, "QJsonObject_count")
	symaddr(&addr_QJsonObject_length, "QJsonObject_length")
	symaddr(&addr_QJsonObject_isEmpty, "QJsonObject_isEmpty")
	symaddr(&addr_QJsonObject_value, "QJsonObject_value")
	symaddr(&addr_QJsonObject_OperatorSubscript, "QJsonObject_OperatorSubscript")
	symaddr(&addr_QJsonObject_OperatorSubscript2, "QJsonObject_OperatorSubscript2")
	symaddr(&addr_QJsonObject_value2, "QJsonObject_value2")
	symaddr(&addr_QJsonObject_value3, "QJsonObject_value3")
	symaddr(&addr_QJsonObject_OperatorSubscript3, "QJsonObject_OperatorSubscript3")
	symaddr(&addr_QJsonObject_OperatorSubscript4, "QJsonObject_OperatorSubscript4")
	symaddr(&addr_QJsonObject_OperatorSubscript5, "QJsonObject_OperatorSubscript5")
	symaddr(&addr_QJsonObject_OperatorSubscript6, "QJsonObject_OperatorSubscript6")
	symaddr(&addr_QJsonObject_remove, "QJsonObject_remove")
	symaddr(&addr_QJsonObject_take, "QJsonObject_take")
	symaddr(&addr_QJsonObject_contains, "QJsonObject_contains")
	symaddr(&addr_QJsonObject_remove2, "QJsonObject_remove2")
	symaddr(&addr_QJsonObject_remove3, "QJsonObject_remove3")
	symaddr(&addr_QJsonObject_take2, "QJsonObject_take2")
	symaddr(&addr_QJsonObject_take3, "QJsonObject_take3")
	symaddr(&addr_QJsonObject_contains2, "QJsonObject_contains2")
	symaddr(&addr_QJsonObject_contains3, "QJsonObject_contains3")
	symaddr(&addr_QJsonObject_empty, "QJsonObject_empty")
	symaddr(&addr_QJsonObject_OperatorEqual, "QJsonObject_OperatorEqual")
	symaddr(&addr_QJsonObject_OperatorNotEqual, "QJsonObject_OperatorNotEqual")
	symaddr(&addr_QJsonObject_OperatorEqual2, "QJsonObject_OperatorEqual2")
	symaddr(&addr_QJsonObject_OperatorNotEqual2, "QJsonObject_OperatorNotEqual2")
	symaddr(&addr_QJsonObject_OperatorEqual3, "QJsonObject_OperatorEqual3")
	symaddr(&addr_QJsonObject_OperatorNotEqual3, "QJsonObject_OperatorNotEqual3")
	reg(&pg_QJsonObject_delete, "QJsonObject_delete")
}

type QJsonObject struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QJsonObject) UnsafeGoID() int64     { return this.vid }
func (this *QJsonObject) UnsafeSetGoID(id int64) { this.vid = id }

func bindQJsonObject(h uintptr) *QJsonObject {
	if h == 0 {
		return nil
	}
	return &QJsonObject{h: h}
}

func newownedQJsonObject(h uintptr) *QJsonObject {
	o := bindQJsonObject(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonObject).Delete)
	}
	return o
}

func (this *QJsonObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonObject(h unsafe.Pointer) *QJsonObject { return bindQJsonObject(uintptr(h)) }
func UnsafeNewQJsonObjectOwned(h unsafe.Pointer) *QJsonObject { return newownedQJsonObject(uintptr(h)) }

func (this *QJsonObject) UnsafeRelease() { this.h = 0 }

func NewQJsonObject() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_new)
	__o := newownedQJsonObject(__r)
	return __o
}

func NewQJsonObject2(other *QJsonObject) *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_new2, uintptr(other.UnsafePointer()))
	__o := newownedQJsonObject(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QJsonObject) OperatorAssign(other *QJsonObject) *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQJsonObject(__r)
}

func (this *QJsonObject) Swap(other *QJsonObject) {
	purego.SyscallN(addr_QJsonObject_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QJsonObject) Keys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonObject_keys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QJsonObject) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonObject) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonObject) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonObject) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonObject) Value(key string) *QJsonValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonObject_value, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) OperatorSubscript(key string) *QJsonValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorSubscript, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) OperatorSubscriptWithKey(key string) *QJsonValueRef {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorSubscript2, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValueRef(__r)
}

func (this *QJsonObject) ValueWithKey(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_value2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) Value2(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_value3, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) OperatorSubscript2(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorSubscript3, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) OperatorSubscript3(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorSubscript4, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) OperatorSubscript4(key *QStringView) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorSubscript5, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValueRef(__r)
}

func (this *QJsonObject) OperatorSubscript5(key *QLatin1String) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_OperatorSubscript6, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValueRef(__r)
}

func (this *QJsonObject) Remove(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QJsonObject_remove, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QJsonObject) Take(key string) *QJsonValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonObject_take, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) Contains(key string) bool {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_contains, this.h, key_p, uintptr(len(key)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QJsonObject) RemoveWithKey(key *QStringView) {
	purego.SyscallN(addr_QJsonObject_remove2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QJsonObject) Remove2(key *QLatin1String) {
	purego.SyscallN(addr_QJsonObject_remove3, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QJsonObject) TakeWithKey(key *QStringView) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_take2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) Take2(key *QLatin1String) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonObject_take3, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQJsonValue(__r)
}

func (this *QJsonObject) ContainsWithKey(key *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_contains2, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QJsonObject) Contains2(key *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_contains3, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QJsonObject) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonObject) OperatorEqual(rhs *QJsonObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonObject) OperatorNotEqual(rhs *QJsonObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonObject) OperatorEqualWithRhs(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonObject) OperatorNotEqualWithRhs(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonObject) OperatorEqual2(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonObject) OperatorNotEqual2(rhs *QJsonValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonObject_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonObject_delete(this.h)
	}
	this.UnsafeRelease()
}
