package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QItemSelection_new uintptr
	addr_QItemSelection_new2 uintptr
	addr_QItemSelection_new3 uintptr
	addr_QItemSelection_select uintptr
	addr_QItemSelection_contains uintptr
	addr_QItemSelection_indexes uintptr
	addr_QItemSelection_merge uintptr
	addr_QItemSelection_split uintptr
	addr_QItemSelection_maxSize uintptr
	addr_QItemSelection_size uintptr
	addr_QItemSelection_count uintptr
	addr_QItemSelection_length uintptr
	addr_QItemSelection_isEmpty uintptr
	addr_QItemSelection_resize uintptr
	addr_QItemSelection_resize2 uintptr
	addr_QItemSelection_resizeForOverwrite uintptr
	addr_QItemSelection_capacity uintptr
	addr_QItemSelection_reserve uintptr
	addr_QItemSelection_squeeze uintptr
	addr_QItemSelection_detach uintptr
	addr_QItemSelection_isDetached uintptr
	addr_QItemSelection_isSharedWith uintptr
	addr_QItemSelection_data uintptr
	addr_QItemSelection_data2 uintptr
	addr_QItemSelection_constData uintptr
	addr_QItemSelection_clear uintptr
	addr_QItemSelection_at uintptr
	addr_QItemSelection_OperatorSubscript uintptr
	addr_QItemSelection_OperatorSubscript2 uintptr
	addr_QItemSelection_append uintptr
	addr_QItemSelection_append2 uintptr
	addr_QItemSelection_prepend uintptr
	addr_QItemSelection_replace uintptr
	addr_QItemSelection_remove uintptr
	addr_QItemSelection_removeFirst uintptr
	addr_QItemSelection_removeLast uintptr
	addr_QItemSelection_takeFirst uintptr
	addr_QItemSelection_takeLast uintptr
	addr_QItemSelection_removeAt uintptr
	addr_QItemSelection_takeAt uintptr
	addr_QItemSelection_move uintptr
	addr_QItemSelection_first uintptr
	addr_QItemSelection_first2 uintptr
	addr_QItemSelection_constFirst uintptr
	addr_QItemSelection_last uintptr
	addr_QItemSelection_last2 uintptr
	addr_QItemSelection_constLast uintptr
	addr_QItemSelection_startsWith uintptr
	addr_QItemSelection_endsWith uintptr
	addr_QItemSelection_mid uintptr
	addr_QItemSelection_first3 uintptr
	addr_QItemSelection_last3 uintptr
	addr_QItemSelection_sliced uintptr
	addr_QItemSelection_sliced2 uintptr
	addr_QItemSelection_value uintptr
	addr_QItemSelection_value2 uintptr
	addr_QItemSelection_swapItemsAt uintptr
	addr_QItemSelection_push_back uintptr
	addr_QItemSelection_push_front uintptr
	addr_QItemSelection_pop_back uintptr
	addr_QItemSelection_pop_front uintptr
	addr_QItemSelection_empty uintptr
	addr_QItemSelection_front uintptr
	addr_QItemSelection_front2 uintptr
	addr_QItemSelection_back uintptr
	addr_QItemSelection_back2 uintptr
	addr_QItemSelection_shrink_to_fit uintptr
	addr_QItemSelection_max_size uintptr
	addr_QItemSelection_OperatorPlus uintptr
	addr_QItemSelection_OperatorPlus2 uintptr
	addr_QItemSelection_fromList uintptr
	addr_QItemSelection_toList uintptr
	addr_QItemSelection_fromVector uintptr
	addr_QItemSelection_toVector uintptr
	addr_QItemSelection_remove2 uintptr
	addr_QItemSelection_mid2 uintptr
	pg_QItemSelection_delete func(uintptr)
)

func init() {
	symaddr(&addr_QItemSelection_new, "QItemSelection_new")
	symaddr(&addr_QItemSelection_new2, "QItemSelection_new2")
	symaddr(&addr_QItemSelection_new3, "QItemSelection_new3")
	symaddr(&addr_QItemSelection_select, "QItemSelection_select")
	symaddr(&addr_QItemSelection_contains, "QItemSelection_contains")
	symaddr(&addr_QItemSelection_indexes, "QItemSelection_indexes")
	symaddr(&addr_QItemSelection_merge, "QItemSelection_merge")
	symaddr(&addr_QItemSelection_split, "QItemSelection_split")
	symaddr(&addr_QItemSelection_maxSize, "QItemSelection_maxSize")
	symaddr(&addr_QItemSelection_size, "QItemSelection_size")
	symaddr(&addr_QItemSelection_count, "QItemSelection_count")
	symaddr(&addr_QItemSelection_length, "QItemSelection_length")
	symaddr(&addr_QItemSelection_isEmpty, "QItemSelection_isEmpty")
	symaddr(&addr_QItemSelection_resize, "QItemSelection_resize")
	symaddr(&addr_QItemSelection_resize2, "QItemSelection_resize2")
	symaddr(&addr_QItemSelection_resizeForOverwrite, "QItemSelection_resizeForOverwrite")
	symaddr(&addr_QItemSelection_capacity, "QItemSelection_capacity")
	symaddr(&addr_QItemSelection_reserve, "QItemSelection_reserve")
	symaddr(&addr_QItemSelection_squeeze, "QItemSelection_squeeze")
	symaddr(&addr_QItemSelection_detach, "QItemSelection_detach")
	symaddr(&addr_QItemSelection_isDetached, "QItemSelection_isDetached")
	symaddr(&addr_QItemSelection_isSharedWith, "QItemSelection_isSharedWith")
	symaddr(&addr_QItemSelection_data, "QItemSelection_data")
	symaddr(&addr_QItemSelection_data2, "QItemSelection_data2")
	symaddr(&addr_QItemSelection_constData, "QItemSelection_constData")
	symaddr(&addr_QItemSelection_clear, "QItemSelection_clear")
	symaddr(&addr_QItemSelection_at, "QItemSelection_at")
	symaddr(&addr_QItemSelection_OperatorSubscript, "QItemSelection_OperatorSubscript")
	symaddr(&addr_QItemSelection_OperatorSubscript2, "QItemSelection_OperatorSubscript2")
	symaddr(&addr_QItemSelection_append, "QItemSelection_append")
	symaddr(&addr_QItemSelection_append2, "QItemSelection_append2")
	symaddr(&addr_QItemSelection_prepend, "QItemSelection_prepend")
	symaddr(&addr_QItemSelection_replace, "QItemSelection_replace")
	symaddr(&addr_QItemSelection_remove, "QItemSelection_remove")
	symaddr(&addr_QItemSelection_removeFirst, "QItemSelection_removeFirst")
	symaddr(&addr_QItemSelection_removeLast, "QItemSelection_removeLast")
	symaddr(&addr_QItemSelection_takeFirst, "QItemSelection_takeFirst")
	symaddr(&addr_QItemSelection_takeLast, "QItemSelection_takeLast")
	symaddr(&addr_QItemSelection_removeAt, "QItemSelection_removeAt")
	symaddr(&addr_QItemSelection_takeAt, "QItemSelection_takeAt")
	symaddr(&addr_QItemSelection_move, "QItemSelection_move")
	symaddr(&addr_QItemSelection_first, "QItemSelection_first")
	symaddr(&addr_QItemSelection_first2, "QItemSelection_first2")
	symaddr(&addr_QItemSelection_constFirst, "QItemSelection_constFirst")
	symaddr(&addr_QItemSelection_last, "QItemSelection_last")
	symaddr(&addr_QItemSelection_last2, "QItemSelection_last2")
	symaddr(&addr_QItemSelection_constLast, "QItemSelection_constLast")
	symaddr(&addr_QItemSelection_startsWith, "QItemSelection_startsWith")
	symaddr(&addr_QItemSelection_endsWith, "QItemSelection_endsWith")
	symaddr(&addr_QItemSelection_mid, "QItemSelection_mid")
	symaddr(&addr_QItemSelection_first3, "QItemSelection_first3")
	symaddr(&addr_QItemSelection_last3, "QItemSelection_last3")
	symaddr(&addr_QItemSelection_sliced, "QItemSelection_sliced")
	symaddr(&addr_QItemSelection_sliced2, "QItemSelection_sliced2")
	symaddr(&addr_QItemSelection_value, "QItemSelection_value")
	symaddr(&addr_QItemSelection_value2, "QItemSelection_value2")
	symaddr(&addr_QItemSelection_swapItemsAt, "QItemSelection_swapItemsAt")
	symaddr(&addr_QItemSelection_push_back, "QItemSelection_push_back")
	symaddr(&addr_QItemSelection_push_front, "QItemSelection_push_front")
	symaddr(&addr_QItemSelection_pop_back, "QItemSelection_pop_back")
	symaddr(&addr_QItemSelection_pop_front, "QItemSelection_pop_front")
	symaddr(&addr_QItemSelection_empty, "QItemSelection_empty")
	symaddr(&addr_QItemSelection_front, "QItemSelection_front")
	symaddr(&addr_QItemSelection_front2, "QItemSelection_front2")
	symaddr(&addr_QItemSelection_back, "QItemSelection_back")
	symaddr(&addr_QItemSelection_back2, "QItemSelection_back2")
	symaddr(&addr_QItemSelection_shrink_to_fit, "QItemSelection_shrink_to_fit")
	symaddr(&addr_QItemSelection_max_size, "QItemSelection_max_size")
	symaddr(&addr_QItemSelection_OperatorPlus, "QItemSelection_OperatorPlus")
	symaddr(&addr_QItemSelection_OperatorPlus2, "QItemSelection_OperatorPlus2")
	symaddr(&addr_QItemSelection_fromList, "QItemSelection_fromList")
	symaddr(&addr_QItemSelection_toList, "QItemSelection_toList")
	symaddr(&addr_QItemSelection_fromVector, "QItemSelection_fromVector")
	symaddr(&addr_QItemSelection_toVector, "QItemSelection_toVector")
	symaddr(&addr_QItemSelection_remove2, "QItemSelection_remove2")
	symaddr(&addr_QItemSelection_mid2, "QItemSelection_mid2")
	reg(&pg_QItemSelection_delete, "QItemSelection_delete")
}

type QItemSelection struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QItemSelection) UnsafeGoID() int64     { return this.vid }
func (this *QItemSelection) UnsafeSetGoID(id int64) { this.vid = id }

func bindQItemSelection(h uintptr) *QItemSelection {
	if h == 0 {
		return nil
	}
	return &QItemSelection{h: h}
}

func newownedQItemSelection(h uintptr) *QItemSelection {
	o := bindQItemSelection(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QItemSelection).Delete)
	}
	return o
}

func (this *QItemSelection) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QItemSelection) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQItemSelection(h unsafe.Pointer) *QItemSelection { return bindQItemSelection(uintptr(h)) }
func UnsafeNewQItemSelectionOwned(h unsafe.Pointer) *QItemSelection { return newownedQItemSelection(uintptr(h)) }

func (this *QItemSelection) UnsafeRelease() { this.h = 0 }

func NewQItemSelection(topLeft *QModelIndex, bottomRight *QModelIndex) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_new, uintptr(topLeft.UnsafePointer()), uintptr(bottomRight.UnsafePointer()))
	__o := newownedQItemSelection(__r)
	runtime.KeepAlive(topLeft)
	runtime.KeepAlive(bottomRight)
	return __o
}

func NewQItemSelection2() *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_new2)
	__o := newownedQItemSelection(__r)
	return __o
}

func NewQItemSelection3(arg1 *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQItemSelection(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QItemSelection) Select(topLeft *QModelIndex, bottomRight *QModelIndex) {
	purego.SyscallN(addr_QItemSelection_select, this.h, uintptr(topLeft.UnsafePointer()), uintptr(bottomRight.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(topLeft)
	runtime.KeepAlive(bottomRight)
}

func (this *QItemSelection) Contains(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_contains, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QItemSelection) Indexes() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_indexes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) Merge(other *QItemSelection, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelection_merge, this.h, uintptr(other.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func QItemSelection_Split(range_ *QItemSelectionRange, other *QItemSelectionRange, result *QItemSelection) {
	purego.SyscallN(addr_QItemSelection_split, uintptr(range_.UnsafePointer()), uintptr(other.UnsafePointer()), uintptr(result.UnsafePointer()))
	runtime.KeepAlive(range_)
	runtime.KeepAlive(other)
	runtime.KeepAlive(result)
}

func QItemSelection_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QItemSelection) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QItemSelection) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QItemSelection) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QItemSelection) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelection) Resize(size int64) {
	purego.SyscallN(addr_QItemSelection_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Resize2(size int64, c *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelection_resize2, this.h, uintptr(size), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QItemSelection) ResizeForOverwrite(size int64) {
	purego.SyscallN(addr_QItemSelection_resizeForOverwrite, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QItemSelection) Reserve(size int64) {
	purego.SyscallN(addr_QItemSelection_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Squeeze() {
	purego.SyscallN(addr_QItemSelection_squeeze, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Detach() {
	purego.SyscallN(addr_QItemSelection_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelection) IsSharedWith(other []*QItemSelectionRange) bool {
	other_arr := make([]uintptr, len(other))
	for __i, __e := range other {
		other_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var other_p uintptr
	if len(other_arr) > 0 {
		other_p = uintptr(unsafe.Pointer(&other_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_isSharedWith, this.h, other_p, uintptr(len(other)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	runtime.KeepAlive(other_arr)
	return __r != 0
}

func (this *QItemSelection) Data() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_data, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Data2() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_data2, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) ConstData() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_constData, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Clear() {
	purego.SyscallN(addr_QItemSelection_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) At(i int64) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) OperatorSubscript(i int64) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) OperatorSubscriptWith(i int64) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Append(t *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelection_append, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QItemSelection) AppendWith(l []*QItemSelectionRange) {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	purego.SyscallN(addr_QItemSelection_append2, this.h, l_p, uintptr(len(l)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
}

func (this *QItemSelection) Prepend(t *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelection_prepend, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QItemSelection) Replace(i int64, t *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelection_replace, this.h, uintptr(i), uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QItemSelection) Remove(i int64) {
	purego.SyscallN(addr_QItemSelection_remove, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) RemoveFirst() {
	purego.SyscallN(addr_QItemSelection_removeFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) RemoveLast() {
	purego.SyscallN(addr_QItemSelection_removeLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) TakeFirst() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_takeFirst, this.h)
	runtime.KeepAlive(this)
	return newownedQItemSelectionRange(__r)
}

func (this *QItemSelection) TakeLast() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_takeLast, this.h)
	runtime.KeepAlive(this)
	return newownedQItemSelectionRange(__r)
}

func (this *QItemSelection) RemoveAt(i int64) {
	purego.SyscallN(addr_QItemSelection_removeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) TakeAt(i int64) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_takeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQItemSelectionRange(__r)
}

func (this *QItemSelection) Move(from int64, to int64) {
	purego.SyscallN(addr_QItemSelection_move, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) First() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_first, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) First2() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_first2, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) ConstFirst() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_constFirst, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Last() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_last, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Last2() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_last2, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) ConstLast() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_constLast, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) StartsWith(t *QItemSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_startsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QItemSelection) EndsWith(t *QItemSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_endsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QItemSelection) Mid(pos int64) []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_mid, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) FirstWith(n int64) []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_first3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) LastWith(n int64) []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_last3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) Sliced(pos int64) []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_sliced, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) Sliced2(pos int64, n int64) []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_sliced2, this.h, uintptr(pos), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) Value(i int64) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_value, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQItemSelectionRange(__r)
}

func (this *QItemSelection) Value2(i int64, defaultValue *QItemSelectionRange) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_value2, this.h, uintptr(i), uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQItemSelectionRange(__r)
}

func (this *QItemSelection) SwapItemsAt(i int64, j int64) {
	purego.SyscallN(addr_QItemSelection_swapItemsAt, this.h, uintptr(i), uintptr(j))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Push_back(t *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelection_push_back, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QItemSelection) Push_front(t *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelection_push_front, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QItemSelection) Pop_back() {
	purego.SyscallN(addr_QItemSelection_pop_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Pop_front() {
	purego.SyscallN(addr_QItemSelection_pop_front, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelection) Front() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_front, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Front2() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_front2, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Back() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_back, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Back2() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelection_back2, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelection) Shrink_to_fit() {
	purego.SyscallN(addr_QItemSelection_shrink_to_fit, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QItemSelection_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QItemSelection) OperatorPlus(l []*QItemSelectionRange) []*QItemSelectionRange {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_OperatorPlus, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) OperatorPlusWith(l []*QItemSelectionRange) []*QItemSelectionRange {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_OperatorPlus2, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QItemSelection_FromList(list []*QItemSelectionRange) []*QItemSelectionRange {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_fromList, list_p, uintptr(len(list)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) ToList() []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_toList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QItemSelection_FromVector(vector []*QItemSelectionRange) []*QItemSelectionRange {
	vector_arr := make([]uintptr, len(vector))
	for __i, __e := range vector {
		vector_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var vector_p uintptr
	if len(vector_arr) > 0 {
		vector_p = uintptr(unsafe.Pointer(&vector_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_fromVector, vector_p, uintptr(len(vector)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(vector)
	runtime.KeepAlive(vector_arr)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) ToVector() []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_toVector, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) Remove2(i int64, n int64) {
	purego.SyscallN(addr_QItemSelection_remove2, this.h, uintptr(i), uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QItemSelection) Mid2(pos int64, len_ int64) []*QItemSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelection_mid2, this.h, uintptr(pos), uintptr(len_), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QItemSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQItemSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelection) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QItemSelection_delete(this.h)
	}
	this.UnsafeRelease()
}
