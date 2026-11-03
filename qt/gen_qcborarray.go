package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborArray_new uintptr
	addr_QCborArray_new2 uintptr
	addr_QCborArray_OperatorAssign uintptr
	addr_QCborArray_swap uintptr
	addr_QCborArray_toCborValue uintptr
	addr_QCborArray_size uintptr
	addr_QCborArray_isEmpty uintptr
	addr_QCborArray_clear uintptr
	addr_QCborArray_at uintptr
	addr_QCborArray_first uintptr
	addr_QCborArray_last uintptr
	addr_QCborArray_OperatorSubscript uintptr
	addr_QCborArray_first2 uintptr
	addr_QCborArray_last2 uintptr
	addr_QCborArray_OperatorSubscript2 uintptr
	addr_QCborArray_insert uintptr
	addr_QCborArray_prepend uintptr
	addr_QCborArray_append uintptr
	addr_QCborArray_removeAt uintptr
	addr_QCborArray_takeAt uintptr
	addr_QCborArray_removeFirst uintptr
	addr_QCborArray_removeLast uintptr
	addr_QCborArray_takeFirst uintptr
	addr_QCborArray_takeLast uintptr
	addr_QCborArray_contains uintptr
	addr_QCborArray_compare uintptr
	addr_QCborArray_push_back uintptr
	addr_QCborArray_push_front uintptr
	addr_QCborArray_pop_front uintptr
	addr_QCborArray_pop_back uintptr
	addr_QCborArray_empty uintptr
	addr_QCborArray_OperatorPlus uintptr
	addr_QCborArray_OperatorPlusAssign uintptr
	addr_QCborArray_OperatorShiftLeft uintptr
	addr_QCborArray_fromStringList uintptr
	addr_QCborArray_fromVariantList uintptr
	addr_QCborArray_fromJsonArray uintptr
	addr_QCborArray_toVariantList uintptr
	addr_QCborArray_toJsonArray uintptr
	addr_QCborArray_OperatorEqual uintptr
	addr_QCborArray_OperatorNotEqual uintptr
	addr_QCborArray_OperatorLesser uintptr
	addr_QCborArray_OperatorGreater uintptr
	addr_QCborArray_OperatorLesserOrEqual uintptr
	addr_QCborArray_OperatorGreaterOrEqual uintptr
	addr_QCborArray_OperatorEqual2 uintptr
	addr_QCborArray_OperatorNotEqual2 uintptr
	addr_QCborArray_OperatorLesser2 uintptr
	addr_QCborArray_OperatorGreater2 uintptr
	addr_QCborArray_OperatorLesserOrEqual2 uintptr
	addr_QCborArray_OperatorGreaterOrEqual2 uintptr
	addr_QCborArray_OperatorEqual3 uintptr
	addr_QCborArray_OperatorNotEqual3 uintptr
	addr_QCborArray_OperatorLesser3 uintptr
	addr_QCborArray_OperatorGreater3 uintptr
	addr_QCborArray_OperatorLesserOrEqual3 uintptr
	addr_QCborArray_OperatorGreaterOrEqual3 uintptr
	pg_QCborArray_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborArray_new, "QCborArray_new")
	symaddr(&addr_QCborArray_new2, "QCborArray_new2")
	symaddr(&addr_QCborArray_OperatorAssign, "QCborArray_OperatorAssign")
	symaddr(&addr_QCborArray_swap, "QCborArray_swap")
	symaddr(&addr_QCborArray_toCborValue, "QCborArray_toCborValue")
	symaddr(&addr_QCborArray_size, "QCborArray_size")
	symaddr(&addr_QCborArray_isEmpty, "QCborArray_isEmpty")
	symaddr(&addr_QCborArray_clear, "QCborArray_clear")
	symaddr(&addr_QCborArray_at, "QCborArray_at")
	symaddr(&addr_QCborArray_first, "QCborArray_first")
	symaddr(&addr_QCborArray_last, "QCborArray_last")
	symaddr(&addr_QCborArray_OperatorSubscript, "QCborArray_OperatorSubscript")
	symaddr(&addr_QCborArray_first2, "QCborArray_first2")
	symaddr(&addr_QCborArray_last2, "QCborArray_last2")
	symaddr(&addr_QCborArray_OperatorSubscript2, "QCborArray_OperatorSubscript2")
	symaddr(&addr_QCborArray_insert, "QCborArray_insert")
	symaddr(&addr_QCborArray_prepend, "QCborArray_prepend")
	symaddr(&addr_QCborArray_append, "QCborArray_append")
	symaddr(&addr_QCborArray_removeAt, "QCborArray_removeAt")
	symaddr(&addr_QCborArray_takeAt, "QCborArray_takeAt")
	symaddr(&addr_QCborArray_removeFirst, "QCborArray_removeFirst")
	symaddr(&addr_QCborArray_removeLast, "QCborArray_removeLast")
	symaddr(&addr_QCborArray_takeFirst, "QCborArray_takeFirst")
	symaddr(&addr_QCborArray_takeLast, "QCborArray_takeLast")
	symaddr(&addr_QCborArray_contains, "QCborArray_contains")
	symaddr(&addr_QCborArray_compare, "QCborArray_compare")
	symaddr(&addr_QCborArray_push_back, "QCborArray_push_back")
	symaddr(&addr_QCborArray_push_front, "QCborArray_push_front")
	symaddr(&addr_QCborArray_pop_front, "QCborArray_pop_front")
	symaddr(&addr_QCborArray_pop_back, "QCborArray_pop_back")
	symaddr(&addr_QCborArray_empty, "QCborArray_empty")
	symaddr(&addr_QCborArray_OperatorPlus, "QCborArray_OperatorPlus")
	symaddr(&addr_QCborArray_OperatorPlusAssign, "QCborArray_OperatorPlusAssign")
	symaddr(&addr_QCborArray_OperatorShiftLeft, "QCborArray_OperatorShiftLeft")
	symaddr(&addr_QCborArray_fromStringList, "QCborArray_fromStringList")
	symaddr(&addr_QCborArray_fromVariantList, "QCborArray_fromVariantList")
	symaddr(&addr_QCborArray_fromJsonArray, "QCborArray_fromJsonArray")
	symaddr(&addr_QCborArray_toVariantList, "QCborArray_toVariantList")
	symaddr(&addr_QCborArray_toJsonArray, "QCborArray_toJsonArray")
	symaddr(&addr_QCborArray_OperatorEqual, "QCborArray_OperatorEqual")
	symaddr(&addr_QCborArray_OperatorNotEqual, "QCborArray_OperatorNotEqual")
	symaddr(&addr_QCborArray_OperatorLesser, "QCborArray_OperatorLesser")
	symaddr(&addr_QCborArray_OperatorGreater, "QCborArray_OperatorGreater")
	symaddr(&addr_QCborArray_OperatorLesserOrEqual, "QCborArray_OperatorLesserOrEqual")
	symaddr(&addr_QCborArray_OperatorGreaterOrEqual, "QCborArray_OperatorGreaterOrEqual")
	symaddr(&addr_QCborArray_OperatorEqual2, "QCborArray_OperatorEqual2")
	symaddr(&addr_QCborArray_OperatorNotEqual2, "QCborArray_OperatorNotEqual2")
	symaddr(&addr_QCborArray_OperatorLesser2, "QCborArray_OperatorLesser2")
	symaddr(&addr_QCborArray_OperatorGreater2, "QCborArray_OperatorGreater2")
	symaddr(&addr_QCborArray_OperatorLesserOrEqual2, "QCborArray_OperatorLesserOrEqual2")
	symaddr(&addr_QCborArray_OperatorGreaterOrEqual2, "QCborArray_OperatorGreaterOrEqual2")
	symaddr(&addr_QCborArray_OperatorEqual3, "QCborArray_OperatorEqual3")
	symaddr(&addr_QCborArray_OperatorNotEqual3, "QCborArray_OperatorNotEqual3")
	symaddr(&addr_QCborArray_OperatorLesser3, "QCborArray_OperatorLesser3")
	symaddr(&addr_QCborArray_OperatorGreater3, "QCborArray_OperatorGreater3")
	symaddr(&addr_QCborArray_OperatorLesserOrEqual3, "QCborArray_OperatorLesserOrEqual3")
	symaddr(&addr_QCborArray_OperatorGreaterOrEqual3, "QCborArray_OperatorGreaterOrEqual3")
	reg(&pg_QCborArray_delete, "QCborArray_delete")
}

type QCborArray struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborArray) UnsafeGoID() int64     { return this.vid }
func (this *QCborArray) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborArray(h uintptr) *QCborArray {
	if h == 0 {
		return nil
	}
	return &QCborArray{h: h}
}

func newownedQCborArray(h uintptr) *QCborArray {
	o := bindQCborArray(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborArray).Delete)
	}
	return o
}

func (this *QCborArray) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborArray) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborArray(h unsafe.Pointer) *QCborArray { return bindQCborArray(uintptr(h)) }
func UnsafeNewQCborArrayOwned(h unsafe.Pointer) *QCborArray { return newownedQCborArray(uintptr(h)) }

func (this *QCborArray) UnsafeRelease() { this.h = 0 }

func NewQCborArray() *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_new)
	__o := newownedQCborArray(__r)
	return __o
}

func NewQCborArray2(other *QCborArray) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_new2, uintptr(other.UnsafePointer()))
	__o := newownedQCborArray(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QCborArray) OperatorAssign(other *QCborArray) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCborArray(__r)
}

func (this *QCborArray) Swap(other *QCborArray) {
	purego.SyscallN(addr_QCborArray_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCborArray) ToCborValue() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_toCborValue, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborArray) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborArray) Clear() {
	purego.SyscallN(addr_QCborArray_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborArray) At(i int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) First() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_first, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) Last() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_last, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) OperatorSubscript(i int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) First2() *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborArray_first2, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValueRef(__r)
}

func (this *QCborArray) Last2() *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborArray_last2, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValueRef(__r)
}

func (this *QCborArray) OperatorSubscriptWith(i int64) *QCborValueRef {
	__r, _, _ := purego.SyscallN(addr_QCborArray_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQCborValueRef(__r)
}

func (this *QCborArray) Insert(i int64, value *QCborValue) {
	purego.SyscallN(addr_QCborArray_insert, this.h, uintptr(i), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QCborArray) Prepend(value *QCborValue) {
	purego.SyscallN(addr_QCborArray_prepend, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QCborArray) Append(value *QCborValue) {
	purego.SyscallN(addr_QCborArray_append, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QCborArray) RemoveAt(i int64) {
	purego.SyscallN(addr_QCborArray_removeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QCborArray) TakeAt(i int64) *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_takeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) RemoveFirst() {
	purego.SyscallN(addr_QCborArray_removeFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborArray) RemoveLast() {
	purego.SyscallN(addr_QCborArray_removeLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborArray) TakeFirst() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_takeFirst, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) TakeLast() *QCborValue {
	__r, _, _ := purego.SyscallN(addr_QCborArray_takeLast, this.h)
	runtime.KeepAlive(this)
	return newownedQCborValue(__r)
}

func (this *QCborArray) Contains(value *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_contains, this.h, uintptr(value.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QCborArray) Compare(other *QCborArray) int {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QCborArray) Push_back(t *QCborValue) {
	purego.SyscallN(addr_QCborArray_push_back, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QCborArray) Push_front(t *QCborValue) {
	purego.SyscallN(addr_QCborArray_push_front, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QCborArray) Pop_front() {
	purego.SyscallN(addr_QCborArray_pop_front, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborArray) Pop_back() {
	purego.SyscallN(addr_QCborArray_pop_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborArray) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborArray) OperatorPlus(v *QCborValue) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_OperatorPlus, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return newownedQCborArray(__r)
}

func (this *QCborArray) OperatorPlusAssign(v *QCborValue) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_OperatorPlusAssign, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQCborArray(__r)
}

func (this *QCborArray) OperatorShiftLeft(v *QCborValue) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_OperatorShiftLeft, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQCborArray(__r)
}

func QCborArray_FromStringList(list []string) *QCborArray {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = pgCString(__e)
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCborArray_fromStringList, list_p, uintptr(len(list)))
	for __i := range list_arr {
		pgFree(list_arr[__i])
	}
	runtime.KeepAlive(list_arr)
	return newownedQCborArray(__r)
}

func QCborArray_FromVariantList(list []*QVariant) *QCborArray {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCborArray_fromVariantList, list_p, uintptr(len(list)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	return newownedQCborArray(__r)
}

func QCborArray_FromJsonArray(array *QJsonArray) *QCborArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_fromJsonArray, uintptr(array.UnsafePointer()))
	runtime.KeepAlive(array)
	return newownedQCborArray(__r)
}

func (this *QCborArray) ToVariantList() []*QVariant {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborArray_toVariantList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QVariant, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQVariant(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCborArray) ToJsonArray() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QCborArray_toJsonArray, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonArray(__r)
}

func (this *QCborArray) OperatorEqual(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorNotEqual(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorLesser(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorGreater(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorLesserOrEqual(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorGreaterOrEqual(rhs *QCborArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorNotEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorLesserWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorGreaterWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorLesserOrEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorGreaterOrEqualWithRhs(rhs *QCborValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorNotEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorLesser2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorGreater2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorLesserOrEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) OperatorGreaterOrEqual2(rhs *QCborValueConstRef) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborArray_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCborArray) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborArray_delete(this.h)
	}
	this.UnsafeRelease()
}
