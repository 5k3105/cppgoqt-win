package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QJsonArray_new uintptr
	addr_QJsonArray_new2 uintptr
	addr_QJsonArray_OperatorAssign uintptr
	addr_QJsonArray_fromStringList uintptr
	addr_QJsonArray_fromVariantList uintptr
	addr_QJsonArray_toVariantList uintptr
	addr_QJsonArray_size uintptr
	addr_QJsonArray_count uintptr
	addr_QJsonArray_isEmpty uintptr
	addr_QJsonArray_at uintptr
	addr_QJsonArray_first uintptr
	addr_QJsonArray_last uintptr
	addr_QJsonArray_prepend uintptr
	addr_QJsonArray_append uintptr
	addr_QJsonArray_removeAt uintptr
	addr_QJsonArray_takeAt uintptr
	addr_QJsonArray_removeFirst uintptr
	addr_QJsonArray_removeLast uintptr
	addr_QJsonArray_insert uintptr
	addr_QJsonArray_replace uintptr
	addr_QJsonArray_contains uintptr
	addr_QJsonArray_OperatorSubscript uintptr
	addr_QJsonArray_OperatorSubscript2 uintptr
	addr_QJsonArray_swap uintptr
	addr_QJsonArray_OperatorPlus uintptr
	addr_QJsonArray_OperatorPlusAssign uintptr
	addr_QJsonArray_OperatorShiftLeft uintptr
	addr_QJsonArray_push_back uintptr
	addr_QJsonArray_push_front uintptr
	addr_QJsonArray_pop_front uintptr
	addr_QJsonArray_pop_back uintptr
	addr_QJsonArray_empty uintptr
	addr_QJsonArray_OperatorEqual uintptr
	addr_QJsonArray_OperatorNotEqual uintptr
	addr_QJsonArray_OperatorEqual2 uintptr
	addr_QJsonArray_OperatorNotEqual2 uintptr
	pg_QJsonArray_delete func(uintptr)
)

func init() {
	symaddr(&addr_QJsonArray_new, "QJsonArray_new")
	symaddr(&addr_QJsonArray_new2, "QJsonArray_new2")
	symaddr(&addr_QJsonArray_OperatorAssign, "QJsonArray_OperatorAssign")
	symaddr(&addr_QJsonArray_fromStringList, "QJsonArray_fromStringList")
	symaddr(&addr_QJsonArray_fromVariantList, "QJsonArray_fromVariantList")
	symaddr(&addr_QJsonArray_toVariantList, "QJsonArray_toVariantList")
	symaddr(&addr_QJsonArray_size, "QJsonArray_size")
	symaddr(&addr_QJsonArray_count, "QJsonArray_count")
	symaddr(&addr_QJsonArray_isEmpty, "QJsonArray_isEmpty")
	symaddr(&addr_QJsonArray_at, "QJsonArray_at")
	symaddr(&addr_QJsonArray_first, "QJsonArray_first")
	symaddr(&addr_QJsonArray_last, "QJsonArray_last")
	symaddr(&addr_QJsonArray_prepend, "QJsonArray_prepend")
	symaddr(&addr_QJsonArray_append, "QJsonArray_append")
	symaddr(&addr_QJsonArray_removeAt, "QJsonArray_removeAt")
	symaddr(&addr_QJsonArray_takeAt, "QJsonArray_takeAt")
	symaddr(&addr_QJsonArray_removeFirst, "QJsonArray_removeFirst")
	symaddr(&addr_QJsonArray_removeLast, "QJsonArray_removeLast")
	symaddr(&addr_QJsonArray_insert, "QJsonArray_insert")
	symaddr(&addr_QJsonArray_replace, "QJsonArray_replace")
	symaddr(&addr_QJsonArray_contains, "QJsonArray_contains")
	symaddr(&addr_QJsonArray_OperatorSubscript, "QJsonArray_OperatorSubscript")
	symaddr(&addr_QJsonArray_OperatorSubscript2, "QJsonArray_OperatorSubscript2")
	symaddr(&addr_QJsonArray_swap, "QJsonArray_swap")
	symaddr(&addr_QJsonArray_OperatorPlus, "QJsonArray_OperatorPlus")
	symaddr(&addr_QJsonArray_OperatorPlusAssign, "QJsonArray_OperatorPlusAssign")
	symaddr(&addr_QJsonArray_OperatorShiftLeft, "QJsonArray_OperatorShiftLeft")
	symaddr(&addr_QJsonArray_push_back, "QJsonArray_push_back")
	symaddr(&addr_QJsonArray_push_front, "QJsonArray_push_front")
	symaddr(&addr_QJsonArray_pop_front, "QJsonArray_pop_front")
	symaddr(&addr_QJsonArray_pop_back, "QJsonArray_pop_back")
	symaddr(&addr_QJsonArray_empty, "QJsonArray_empty")
	symaddr(&addr_QJsonArray_OperatorEqual, "QJsonArray_OperatorEqual")
	symaddr(&addr_QJsonArray_OperatorNotEqual, "QJsonArray_OperatorNotEqual")
	symaddr(&addr_QJsonArray_OperatorEqual2, "QJsonArray_OperatorEqual2")
	symaddr(&addr_QJsonArray_OperatorNotEqual2, "QJsonArray_OperatorNotEqual2")
	reg(&pg_QJsonArray_delete, "QJsonArray_delete")
}

type QJsonArray struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QJsonArray) UnsafeGoID() int64     { return this.vid }
func (this *QJsonArray) UnsafeSetGoID(id int64) { this.vid = id }

func bindQJsonArray(h uintptr) *QJsonArray {
	if h == 0 {
		return nil
	}
	return &QJsonArray{h: h}
}

func newownedQJsonArray(h uintptr) *QJsonArray {
	o := bindQJsonArray(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonArray).Delete)
	}
	return o
}

func (this *QJsonArray) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonArray) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonArray(h unsafe.Pointer) *QJsonArray { return bindQJsonArray(uintptr(h)) }
func UnsafeNewQJsonArrayOwned(h unsafe.Pointer) *QJsonArray { return newownedQJsonArray(uintptr(h)) }

func (this *QJsonArray) UnsafeRelease() { this.h = 0 }

func NewQJsonArray() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_new)
	__o := newownedQJsonArray(__r)
	return __o
}

func NewQJsonArray2(other *QJsonArray) *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_new2, uintptr(other.UnsafePointer()))
	__o := newownedQJsonArray(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QJsonArray) OperatorAssign(other *QJsonArray) *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQJsonArray(__r)
}

func QJsonArray_FromStringList(list []string) *QJsonArray {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = pgCString(__e)
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonArray_fromStringList, list_p, uintptr(len(list)))
	for __i := range list_arr {
		pgFree(list_arr[__i])
	}
	runtime.KeepAlive(list_arr)
	return newownedQJsonArray(__r)
}

func QJsonArray_FromVariantList(list []*QVariant) *QJsonArray {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QJsonArray_fromVariantList, list_p, uintptr(len(list)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	return newownedQJsonArray(__r)
}

func (this *QJsonArray) ToVariantList() []*QVariant {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonArray_toVariantList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QVariant, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQVariant(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QJsonArray) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonArray) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QJsonArray) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonArray) At(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonArray) First() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_first, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonArray) Last() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_last, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonArray) Prepend(value *QJsonValue) {
	purego.SyscallN(addr_QJsonArray_prepend, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QJsonArray) Append(value *QJsonValue) {
	purego.SyscallN(addr_QJsonArray_append, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QJsonArray) RemoveAt(i int64) {
	purego.SyscallN(addr_QJsonArray_removeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QJsonArray) TakeAt(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_takeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonArray) RemoveFirst() {
	purego.SyscallN(addr_QJsonArray_removeFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QJsonArray) RemoveLast() {
	purego.SyscallN(addr_QJsonArray_removeLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QJsonArray) Insert(i int64, value *QJsonValue) {
	purego.SyscallN(addr_QJsonArray_insert, this.h, uintptr(i), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QJsonArray) Replace(i int64, value *QJsonValue) {
	purego.SyscallN(addr_QJsonArray_replace, this.h, uintptr(i), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QJsonArray) Contains(element *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_contains, this.h, uintptr(element.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(element)
	return __r != 0
}

func (this *QJsonArray) OperatorSubscript(i int64) *QJsonValueRef {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValueRef(__r)
}

func (this *QJsonArray) OperatorSubscriptWith(i int64) *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QJsonArray) Swap(other *QJsonArray) {
	purego.SyscallN(addr_QJsonArray_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QJsonArray) OperatorPlus(v *QJsonValue) *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_OperatorPlus, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return newownedQJsonArray(__r)
}

func (this *QJsonArray) OperatorPlusAssign(v *QJsonValue) *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_OperatorPlusAssign, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQJsonArray(__r)
}

func (this *QJsonArray) OperatorShiftLeft(v *QJsonValue) *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QJsonArray_OperatorShiftLeft, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQJsonArray(__r)
}

func (this *QJsonArray) Push_back(t *QJsonValue) {
	purego.SyscallN(addr_QJsonArray_push_back, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QJsonArray) Push_front(t *QJsonValue) {
	purego.SyscallN(addr_QJsonArray_push_front, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QJsonArray) Pop_front() {
	purego.SyscallN(addr_QJsonArray_pop_front, this.h)
	runtime.KeepAlive(this)
}

func (this *QJsonArray) Pop_back() {
	purego.SyscallN(addr_QJsonArray_pop_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QJsonArray) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QJsonArray) OperatorEqual(rhs *QJsonArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonArray) OperatorNotEqual(rhs *QJsonArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonArray) OperatorEqualWithRhs(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonArray) OperatorNotEqualWithRhs(rhs *QJsonValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QJsonArray_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QJsonArray) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonArray_delete(this.h)
	}
	this.UnsafeRelease()
}
