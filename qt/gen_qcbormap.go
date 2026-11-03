package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): QCborMap_FromVariantMap, QCborMap_FromVariantHash

var (
	addr_QCborMap_new uintptr
	addr_QCborMap_new2 uintptr
	addr_QCborMap_OperatorAssign uintptr
	addr_QCborMap_swap uintptr
	addr_QCborMap_toCborValue uintptr
	addr_QCborMap_size uintptr
	addr_QCborMap_isEmpty uintptr
	addr_QCborMap_clear uintptr
	addr_QCborMap_keys uintptr
	addr_QCborMap_value uintptr
	addr_QCborMap_value2 uintptr
	addr_QCborMap_value3 uintptr
	addr_QCborMap_value4 uintptr
	addr_QCborMap_OperatorSubscript uintptr
	addr_QCborMap_OperatorSubscript2 uintptr
	addr_QCborMap_OperatorSubscript3 uintptr
	addr_QCborMap_OperatorSubscript4 uintptr
	addr_QCborMap_OperatorSubscript5 uintptr
	addr_QCborMap_OperatorSubscript6 uintptr
	addr_QCborMap_OperatorSubscript7 uintptr
	addr_QCborMap_OperatorSubscript8 uintptr
	addr_QCborMap_take uintptr
	addr_QCborMap_take2 uintptr
	addr_QCborMap_take3 uintptr
	addr_QCborMap_take4 uintptr
	addr_QCborMap_remove uintptr
	addr_QCborMap_remove2 uintptr
	addr_QCborMap_remove3 uintptr
	addr_QCborMap_remove4 uintptr
	addr_QCborMap_contains uintptr
	addr_QCborMap_contains2 uintptr
	addr_QCborMap_contains3 uintptr
	addr_QCborMap_contains4 uintptr
	addr_QCborMap_compare uintptr
	addr_QCborMap_empty uintptr
	addr_QCborMap_fromJsonObject uintptr
	addr_QCborMap_toJsonObject uintptr
	addr_QCborMap_OperatorEqual uintptr
	addr_QCborMap_OperatorNotEqual uintptr
	addr_QCborMap_OperatorLesser uintptr
	addr_QCborMap_OperatorGreater uintptr
	addr_QCborMap_OperatorLesserOrEqual uintptr
	addr_QCborMap_OperatorGreaterOrEqual uintptr
	addr_QCborMap_OperatorEqual2 uintptr
	addr_QCborMap_OperatorNotEqual2 uintptr
	addr_QCborMap_OperatorLesser2 uintptr
	addr_QCborMap_OperatorGreater2 uintptr
	addr_QCborMap_OperatorLesserOrEqual2 uintptr
	addr_QCborMap_OperatorGreaterOrEqual2 uintptr
	addr_QCborMap_OperatorEqual3 uintptr
	addr_QCborMap_OperatorNotEqual3 uintptr
	addr_QCborMap_OperatorLesser3 uintptr
	addr_QCborMap_OperatorGreater3 uintptr
	addr_QCborMap_OperatorLesserOrEqual3 uintptr
	addr_QCborMap_OperatorGreaterOrEqual3 uintptr
	pg_QCborMap_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborMap_new, "QCborMap_new")
	symaddr(&addr_QCborMap_new2, "QCborMap_new2")
	symaddr(&addr_QCborMap_OperatorAssign, "QCborMap_OperatorAssign")
	symaddr(&addr_QCborMap_swap, "QCborMap_swap")
	symaddr(&addr_QCborMap_toCborValue, "QCborMap_toCborValue")
	symaddr(&addr_QCborMap_size, "QCborMap_size")
	symaddr(&addr_QCborMap_isEmpty, "QCborMap_isEmpty")
	symaddr(&addr_QCborMap_clear, "QCborMap_clear")
	symaddr(&addr_QCborMap_keys, "QCborMap_keys")
	symaddr(&addr_QCborMap_value, "QCborMap_value")
	symaddr(&addr_QCborMap_value2, "QCborMap_value2")
	symaddr(&addr_QCborMap_value3, "QCborMap_value3")
	symaddr(&addr_QCborMap_value4, "QCborMap_value4")
	symaddr(&addr_QCborMap_OperatorSubscript, "QCborMap_OperatorSubscript")
	symaddr(&addr_QCborMap_OperatorSubscript2, "QCborMap_OperatorSubscript2")
	symaddr(&addr_QCborMap_OperatorSubscript3, "QCborMap_OperatorSubscript3")
	symaddr(&addr_QCborMap_OperatorSubscript4, "QCborMap_OperatorSubscript4")
	symaddr(&addr_QCborMap_OperatorSubscript5, "QCborMap_OperatorSubscript5")
	symaddr(&addr_QCborMap_OperatorSubscript6, "QCborMap_OperatorSubscript6")
	symaddr(&addr_QCborMap_OperatorSubscript7, "QCborMap_OperatorSubscript7")
	symaddr(&addr_QCborMap_OperatorSubscript8, "QCborMap_OperatorSubscript8")
	symaddr(&addr_QCborMap_take, "QCborMap_take")
	symaddr(&addr_QCborMap_take2, "QCborMap_take2")
	symaddr(&addr_QCborMap_take3, "QCborMap_take3")
	symaddr(&addr_QCborMap_take4, "QCborMap_take4")
	symaddr(&addr_QCborMap_remove, "QCborMap_remove")
	symaddr(&addr_QCborMap_remove2, "QCborMap_remove2")
	symaddr(&addr_QCborMap_remove3, "QCborMap_remove3")
	symaddr(&addr_QCborMap_remove4, "QCborMap_remove4")
	symaddr(&addr_QCborMap_contains, "QCborMap_contains")
	symaddr(&addr_QCborMap_contains2, "QCborMap_contains2")
	symaddr(&addr_QCborMap_contains3, "QCborMap_contains3")
	symaddr(&addr_QCborMap_contains4, "QCborMap_contains4")
	symaddr(&addr_QCborMap_compare, "QCborMap_compare")
	symaddr(&addr_QCborMap_empty, "QCborMap_empty")
	symaddr(&addr_QCborMap_fromJsonObject, "QCborMap_fromJsonObject")
	symaddr(&addr_QCborMap_toJsonObject, "QCborMap_toJsonObject")
	symaddr(&addr_QCborMap_OperatorEqual, "QCborMap_OperatorEqual")
	symaddr(&addr_QCborMap_OperatorNotEqual, "QCborMap_OperatorNotEqual")
	symaddr(&addr_QCborMap_OperatorLesser, "QCborMap_OperatorLesser")
	symaddr(&addr_QCborMap_OperatorGreater, "QCborMap_OperatorGreater")
	symaddr(&addr_QCborMap_OperatorLesserOrEqual, "QCborMap_OperatorLesserOrEqual")
	symaddr(&addr_QCborMap_OperatorGreaterOrEqual, "QCborMap_OperatorGreaterOrEqual")
	symaddr(&addr_QCborMap_OperatorEqual2, "QCborMap_OperatorEqual2")
	symaddr(&addr_QCborMap_OperatorNotEqual2, "QCborMap_OperatorNotEqual2")
	symaddr(&addr_QCborMap_OperatorLesser2, "QCborMap_OperatorLesser2")
	symaddr(&addr_QCborMap_OperatorGreater2, "QCborMap_OperatorGreater2")
	symaddr(&addr_QCborMap_OperatorLesserOrEqual2, "QCborMap_OperatorLesserOrEqual2")
	symaddr(&addr_QCborMap_OperatorGreaterOrEqual2, "QCborMap_OperatorGreaterOrEqual2")
	symaddr(&addr_QCborMap_OperatorEqual3, "QCborMap_OperatorEqual3")
	symaddr(&addr_QCborMap_OperatorNotEqual3, "QCborMap_OperatorNotEqual3")
	symaddr(&addr_QCborMap_OperatorLesser3, "QCborMap_OperatorLesser3")
	symaddr(&addr_QCborMap_OperatorGreater3, "QCborMap_OperatorGreater3")
	symaddr(&addr_QCborMap_OperatorLesserOrEqual3, "QCborMap_OperatorLesserOrEqual3")
	symaddr(&addr_QCborMap_OperatorGreaterOrEqual3, "QCborMap_OperatorGreaterOrEqual3")
	reg(&pg_QCborMap_delete, "QCborMap_delete")
}

type QCborMap struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborMap) UnsafeGoID() int64     { return this.vid }
func (this *QCborMap) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborMap(h uintptr) *QCborMap {
	if h == 0 {
		return nil
	}
	return &QCborMap{h: h}
}

func newownedQCborMap(h uintptr) *QCborMap {
	o := bindQCborMap(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborMap).Delete)
	}
	return o
}

func (this *QCborMap) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborMap) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborMap(h unsafe.Pointer) *QCborMap { return bindQCborMap(uintptr(h)) }
func UnsafeNewQCborMapOwned(h unsafe.Pointer) *QCborMap { return newownedQCborMap(uintptr(h)) }

func (this *QCborMap) UnsafeRelease() { this.h = 0 }

func NewQCborMap() *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborMap_new)
	__o := newownedQCborMap(__r)
	return __o
}

func NewQCborMap2(other *QCborMap) *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborMap_new2, uintptr(other.UnsafePointer()))
	__o := newownedQCborMap(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QCborMap) OperatorAssign(other *QCborMap) *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCborMap(__r)
}

func (this *QCborMap) Swap(other *QCborMap) {
	purego.SyscallN(addr_QCborMap_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCborMap) ToCborValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_toCborValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborMap) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborMap) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborMap) Clear() {
	purego.SyscallN(addr_QCborMap_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborMap) Keys() []*QCborValue {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborMap_keys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QCborValue, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQCborValue(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCborMap) Value(key int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_value, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborMap) ValueWithKey(key *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_value2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) Value2(key string) *QCborValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborMap_value3, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) Value3(key *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_value4, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) OperatorSubscript(key int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborMap) OperatorSubscriptWithKey(key *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) OperatorSubscript2(key string) *QCborValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript3, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) OperatorSubscript3(key *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript4, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) OperatorSubscript4(key int64) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript5, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValueRef(__r)
}

func (this *QCborMap) OperatorSubscript5(key *QLatin1String) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript6, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborMap) OperatorSubscript6(key string) *QCborValueRef {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript7, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborMap) OperatorSubscript7(key *QCborValue) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborMap_OperatorSubscript8, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValueRef(__r)
}

func (this *QCborMap) Take(key int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_take, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborMap) TakeWithKey(key *QLatin1String) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_take2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) Take2(key string) *QCborValue {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QCborMap_take3, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) Take3(key *QCborValue) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborMap_take4, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQCborValue(__r)
}

func (this *QCborMap) Remove(key int64) {
	purego.SyscallN(addr_QCborMap_remove, this.h, uintptr(key))
	runtime.KeepAlive(this)
}

func (this *QCborMap) RemoveWithKey(key *QLatin1String) {
	purego.SyscallN(addr_QCborMap_remove2, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QCborMap) Remove2(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QCborMap_remove3, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QCborMap) Remove3(key *QCborValue) {
	purego.SyscallN(addr_QCborMap_remove4, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QCborMap) Contains(key int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_contains, this.h, uintptr(key))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborMap) ContainsWithKey(key *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_contains2, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QCborMap) Contains2(key string) bool {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r0, _, _ := purego.SyscallN(addr_QCborMap_contains3, this.h, key_p, uintptr(len(key)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QCborMap) Contains3(key *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_contains4, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QCborMap) Compare(other *QCborMap) int {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QCborMap) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QCborMap_FromJsonObject(o *QJsonObject) *QCborMap {
	__r, _, _ := purego.SyscallN(addr_QCborMap_fromJsonObject, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(o)
	return newownedQCborMap(__r)
}

func (this *QCborMap) ToJsonObject() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QCborMap_toJsonObject, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QCborMap) OperatorEqual(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorNotEqual(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorLesser(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorGreater(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorLesserOrEqual(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorGreaterOrEqual(rhs *QCborMap) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorNotEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorLesserWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorGreaterWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorLesserOrEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorGreaterOrEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorNotEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorLesser2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorGreater2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorLesserOrEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) OperatorGreaterOrEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborMap_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborMap) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborMap_delete(this.h)
	}
	this.UnsafeRelease()
}
