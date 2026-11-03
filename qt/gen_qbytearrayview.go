package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QByteArrayView_new uintptr
	addr_QByteArrayView_new2 uintptr
	addr_QByteArrayView_toByteArray uintptr
	addr_QByteArrayView_size uintptr
	addr_QByteArrayView_data uintptr
	addr_QByteArrayView_constData uintptr
	addr_QByteArrayView_OperatorSubscript uintptr
	addr_QByteArrayView_at uintptr
	addr_QByteArrayView_first uintptr
	addr_QByteArrayView_last uintptr
	addr_QByteArrayView_sliced uintptr
	addr_QByteArrayView_sliced2 uintptr
	addr_QByteArrayView_slice uintptr
	addr_QByteArrayView_slice2 uintptr
	addr_QByteArrayView_chopped uintptr
	addr_QByteArrayView_left uintptr
	addr_QByteArrayView_right uintptr
	addr_QByteArrayView_mid uintptr
	addr_QByteArrayView_truncate uintptr
	addr_QByteArrayView_chop uintptr
	addr_QByteArrayView_trimmed uintptr
	addr_QByteArrayView_toShort uintptr
	addr_QByteArrayView_toUShort uintptr
	addr_QByteArrayView_toInt uintptr
	addr_QByteArrayView_toUInt uintptr
	addr_QByteArrayView_toLong uintptr
	addr_QByteArrayView_toULong uintptr
	addr_QByteArrayView_toLongLong uintptr
	addr_QByteArrayView_toULongLong uintptr
	pg_QByteArrayView_toFloat func(uintptr) float32
	pg_QByteArrayView_toDouble func(uintptr) float64
	addr_QByteArrayView_startsWith uintptr
	addr_QByteArrayView_startsWith2 uintptr
	addr_QByteArrayView_endsWith uintptr
	addr_QByteArrayView_endsWith2 uintptr
	addr_QByteArrayView_indexOf uintptr
	addr_QByteArrayView_indexOf2 uintptr
	addr_QByteArrayView_contains uintptr
	addr_QByteArrayView_contains2 uintptr
	addr_QByteArrayView_lastIndexOf uintptr
	addr_QByteArrayView_lastIndexOf2 uintptr
	addr_QByteArrayView_lastIndexOf3 uintptr
	addr_QByteArrayView_count uintptr
	addr_QByteArrayView_count2 uintptr
	addr_QByteArrayView_compare uintptr
	addr_QByteArrayView_isValidUtf8 uintptr
	addr_QByteArrayView_begin uintptr
	addr_QByteArrayView_end uintptr
	addr_QByteArrayView_cbegin uintptr
	addr_QByteArrayView_cend uintptr
	addr_QByteArrayView_empty uintptr
	addr_QByteArrayView_front uintptr
	addr_QByteArrayView_back uintptr
	addr_QByteArrayView_max_size uintptr
	addr_QByteArrayView_isNull uintptr
	addr_QByteArrayView_isEmpty uintptr
	addr_QByteArrayView_length uintptr
	addr_QByteArrayView_first2 uintptr
	addr_QByteArrayView_last2 uintptr
	addr_QByteArrayView_maxSize uintptr
	addr_QByteArrayView_OperatorAssign uintptr
	addr_QByteArrayView_mid2 uintptr
	addr_QByteArrayView_toShort2 uintptr
	addr_QByteArrayView_toShort3 uintptr
	addr_QByteArrayView_toUShort2 uintptr
	addr_QByteArrayView_toUShort3 uintptr
	addr_QByteArrayView_toInt2 uintptr
	addr_QByteArrayView_toInt3 uintptr
	addr_QByteArrayView_toUInt2 uintptr
	addr_QByteArrayView_toUInt3 uintptr
	addr_QByteArrayView_toLong2 uintptr
	addr_QByteArrayView_toLong3 uintptr
	addr_QByteArrayView_toULong2 uintptr
	addr_QByteArrayView_toULong3 uintptr
	addr_QByteArrayView_toLongLong2 uintptr
	addr_QByteArrayView_toLongLong3 uintptr
	addr_QByteArrayView_toULongLong2 uintptr
	addr_QByteArrayView_toULongLong3 uintptr
	pg_QByteArrayView_toFloat2 func(uintptr, uintptr) float32
	pg_QByteArrayView_toDouble2 func(uintptr, uintptr) float64
	addr_QByteArrayView_indexOf3 uintptr
	addr_QByteArrayView_indexOf4 uintptr
	addr_QByteArrayView_lastIndexOf4 uintptr
	addr_QByteArrayView_compare2 uintptr
	addr_QByteArrayView_OperatorEqual uintptr
	addr_QByteArrayView_OperatorNotEqual uintptr
	addr_QByteArrayView_OperatorLesser uintptr
	addr_QByteArrayView_OperatorGreater uintptr
	addr_QByteArrayView_OperatorLesserOrEqual uintptr
	addr_QByteArrayView_OperatorGreaterOrEqual uintptr
	addr_QByteArrayView_OperatorEqual2 uintptr
	addr_QByteArrayView_OperatorNotEqual2 uintptr
	addr_QByteArrayView_OperatorLesser2 uintptr
	addr_QByteArrayView_OperatorGreater2 uintptr
	addr_QByteArrayView_OperatorLesserOrEqual2 uintptr
	addr_QByteArrayView_OperatorGreaterOrEqual2 uintptr
	addr_QByteArrayView_OperatorEqual3 uintptr
	addr_QByteArrayView_OperatorNotEqual3 uintptr
	addr_QByteArrayView_OperatorLesser3 uintptr
	addr_QByteArrayView_OperatorGreater3 uintptr
	addr_QByteArrayView_OperatorLesserOrEqual3 uintptr
	addr_QByteArrayView_OperatorGreaterOrEqual3 uintptr
	addr_QByteArrayView_OperatorEqual4 uintptr
	addr_QByteArrayView_OperatorNotEqual4 uintptr
	addr_QByteArrayView_OperatorLesser4 uintptr
	addr_QByteArrayView_OperatorGreater4 uintptr
	addr_QByteArrayView_OperatorLesserOrEqual4 uintptr
	addr_QByteArrayView_OperatorGreaterOrEqual4 uintptr
	pg_QByteArrayView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QByteArrayView_new, "QByteArrayView_new")
	symaddr(&addr_QByteArrayView_new2, "QByteArrayView_new2")
	symaddr(&addr_QByteArrayView_toByteArray, "QByteArrayView_toByteArray")
	symaddr(&addr_QByteArrayView_size, "QByteArrayView_size")
	symaddr(&addr_QByteArrayView_data, "QByteArrayView_data")
	symaddr(&addr_QByteArrayView_constData, "QByteArrayView_constData")
	symaddr(&addr_QByteArrayView_OperatorSubscript, "QByteArrayView_OperatorSubscript")
	symaddr(&addr_QByteArrayView_at, "QByteArrayView_at")
	symaddr(&addr_QByteArrayView_first, "QByteArrayView_first")
	symaddr(&addr_QByteArrayView_last, "QByteArrayView_last")
	symaddr(&addr_QByteArrayView_sliced, "QByteArrayView_sliced")
	symaddr(&addr_QByteArrayView_sliced2, "QByteArrayView_sliced2")
	symaddr(&addr_QByteArrayView_slice, "QByteArrayView_slice")
	symaddr(&addr_QByteArrayView_slice2, "QByteArrayView_slice2")
	symaddr(&addr_QByteArrayView_chopped, "QByteArrayView_chopped")
	symaddr(&addr_QByteArrayView_left, "QByteArrayView_left")
	symaddr(&addr_QByteArrayView_right, "QByteArrayView_right")
	symaddr(&addr_QByteArrayView_mid, "QByteArrayView_mid")
	symaddr(&addr_QByteArrayView_truncate, "QByteArrayView_truncate")
	symaddr(&addr_QByteArrayView_chop, "QByteArrayView_chop")
	symaddr(&addr_QByteArrayView_trimmed, "QByteArrayView_trimmed")
	symaddr(&addr_QByteArrayView_toShort, "QByteArrayView_toShort")
	symaddr(&addr_QByteArrayView_toUShort, "QByteArrayView_toUShort")
	symaddr(&addr_QByteArrayView_toInt, "QByteArrayView_toInt")
	symaddr(&addr_QByteArrayView_toUInt, "QByteArrayView_toUInt")
	symaddr(&addr_QByteArrayView_toLong, "QByteArrayView_toLong")
	symaddr(&addr_QByteArrayView_toULong, "QByteArrayView_toULong")
	symaddr(&addr_QByteArrayView_toLongLong, "QByteArrayView_toLongLong")
	symaddr(&addr_QByteArrayView_toULongLong, "QByteArrayView_toULongLong")
	reg(&pg_QByteArrayView_toFloat, "QByteArrayView_toFloat")
	reg(&pg_QByteArrayView_toDouble, "QByteArrayView_toDouble")
	symaddr(&addr_QByteArrayView_startsWith, "QByteArrayView_startsWith")
	symaddr(&addr_QByteArrayView_startsWith2, "QByteArrayView_startsWith2")
	symaddr(&addr_QByteArrayView_endsWith, "QByteArrayView_endsWith")
	symaddr(&addr_QByteArrayView_endsWith2, "QByteArrayView_endsWith2")
	symaddr(&addr_QByteArrayView_indexOf, "QByteArrayView_indexOf")
	symaddr(&addr_QByteArrayView_indexOf2, "QByteArrayView_indexOf2")
	symaddr(&addr_QByteArrayView_contains, "QByteArrayView_contains")
	symaddr(&addr_QByteArrayView_contains2, "QByteArrayView_contains2")
	symaddr(&addr_QByteArrayView_lastIndexOf, "QByteArrayView_lastIndexOf")
	symaddr(&addr_QByteArrayView_lastIndexOf2, "QByteArrayView_lastIndexOf2")
	symaddr(&addr_QByteArrayView_lastIndexOf3, "QByteArrayView_lastIndexOf3")
	symaddr(&addr_QByteArrayView_count, "QByteArrayView_count")
	symaddr(&addr_QByteArrayView_count2, "QByteArrayView_count2")
	symaddr(&addr_QByteArrayView_compare, "QByteArrayView_compare")
	symaddr(&addr_QByteArrayView_isValidUtf8, "QByteArrayView_isValidUtf8")
	symaddr(&addr_QByteArrayView_begin, "QByteArrayView_begin")
	symaddr(&addr_QByteArrayView_end, "QByteArrayView_end")
	symaddr(&addr_QByteArrayView_cbegin, "QByteArrayView_cbegin")
	symaddr(&addr_QByteArrayView_cend, "QByteArrayView_cend")
	symaddr(&addr_QByteArrayView_empty, "QByteArrayView_empty")
	symaddr(&addr_QByteArrayView_front, "QByteArrayView_front")
	symaddr(&addr_QByteArrayView_back, "QByteArrayView_back")
	symaddr(&addr_QByteArrayView_max_size, "QByteArrayView_max_size")
	symaddr(&addr_QByteArrayView_isNull, "QByteArrayView_isNull")
	symaddr(&addr_QByteArrayView_isEmpty, "QByteArrayView_isEmpty")
	symaddr(&addr_QByteArrayView_length, "QByteArrayView_length")
	symaddr(&addr_QByteArrayView_first2, "QByteArrayView_first2")
	symaddr(&addr_QByteArrayView_last2, "QByteArrayView_last2")
	symaddr(&addr_QByteArrayView_maxSize, "QByteArrayView_maxSize")
	symaddr(&addr_QByteArrayView_OperatorAssign, "QByteArrayView_OperatorAssign")
	symaddr(&addr_QByteArrayView_mid2, "QByteArrayView_mid2")
	symaddr(&addr_QByteArrayView_toShort2, "QByteArrayView_toShort2")
	symaddr(&addr_QByteArrayView_toShort3, "QByteArrayView_toShort3")
	symaddr(&addr_QByteArrayView_toUShort2, "QByteArrayView_toUShort2")
	symaddr(&addr_QByteArrayView_toUShort3, "QByteArrayView_toUShort3")
	symaddr(&addr_QByteArrayView_toInt2, "QByteArrayView_toInt2")
	symaddr(&addr_QByteArrayView_toInt3, "QByteArrayView_toInt3")
	symaddr(&addr_QByteArrayView_toUInt2, "QByteArrayView_toUInt2")
	symaddr(&addr_QByteArrayView_toUInt3, "QByteArrayView_toUInt3")
	symaddr(&addr_QByteArrayView_toLong2, "QByteArrayView_toLong2")
	symaddr(&addr_QByteArrayView_toLong3, "QByteArrayView_toLong3")
	symaddr(&addr_QByteArrayView_toULong2, "QByteArrayView_toULong2")
	symaddr(&addr_QByteArrayView_toULong3, "QByteArrayView_toULong3")
	symaddr(&addr_QByteArrayView_toLongLong2, "QByteArrayView_toLongLong2")
	symaddr(&addr_QByteArrayView_toLongLong3, "QByteArrayView_toLongLong3")
	symaddr(&addr_QByteArrayView_toULongLong2, "QByteArrayView_toULongLong2")
	symaddr(&addr_QByteArrayView_toULongLong3, "QByteArrayView_toULongLong3")
	reg(&pg_QByteArrayView_toFloat2, "QByteArrayView_toFloat2")
	reg(&pg_QByteArrayView_toDouble2, "QByteArrayView_toDouble2")
	symaddr(&addr_QByteArrayView_indexOf3, "QByteArrayView_indexOf3")
	symaddr(&addr_QByteArrayView_indexOf4, "QByteArrayView_indexOf4")
	symaddr(&addr_QByteArrayView_lastIndexOf4, "QByteArrayView_lastIndexOf4")
	symaddr(&addr_QByteArrayView_compare2, "QByteArrayView_compare2")
	symaddr(&addr_QByteArrayView_OperatorEqual, "QByteArrayView_OperatorEqual")
	symaddr(&addr_QByteArrayView_OperatorNotEqual, "QByteArrayView_OperatorNotEqual")
	symaddr(&addr_QByteArrayView_OperatorLesser, "QByteArrayView_OperatorLesser")
	symaddr(&addr_QByteArrayView_OperatorGreater, "QByteArrayView_OperatorGreater")
	symaddr(&addr_QByteArrayView_OperatorLesserOrEqual, "QByteArrayView_OperatorLesserOrEqual")
	symaddr(&addr_QByteArrayView_OperatorGreaterOrEqual, "QByteArrayView_OperatorGreaterOrEqual")
	symaddr(&addr_QByteArrayView_OperatorEqual2, "QByteArrayView_OperatorEqual2")
	symaddr(&addr_QByteArrayView_OperatorNotEqual2, "QByteArrayView_OperatorNotEqual2")
	symaddr(&addr_QByteArrayView_OperatorLesser2, "QByteArrayView_OperatorLesser2")
	symaddr(&addr_QByteArrayView_OperatorGreater2, "QByteArrayView_OperatorGreater2")
	symaddr(&addr_QByteArrayView_OperatorLesserOrEqual2, "QByteArrayView_OperatorLesserOrEqual2")
	symaddr(&addr_QByteArrayView_OperatorGreaterOrEqual2, "QByteArrayView_OperatorGreaterOrEqual2")
	symaddr(&addr_QByteArrayView_OperatorEqual3, "QByteArrayView_OperatorEqual3")
	symaddr(&addr_QByteArrayView_OperatorNotEqual3, "QByteArrayView_OperatorNotEqual3")
	symaddr(&addr_QByteArrayView_OperatorLesser3, "QByteArrayView_OperatorLesser3")
	symaddr(&addr_QByteArrayView_OperatorGreater3, "QByteArrayView_OperatorGreater3")
	symaddr(&addr_QByteArrayView_OperatorLesserOrEqual3, "QByteArrayView_OperatorLesserOrEqual3")
	symaddr(&addr_QByteArrayView_OperatorGreaterOrEqual3, "QByteArrayView_OperatorGreaterOrEqual3")
	symaddr(&addr_QByteArrayView_OperatorEqual4, "QByteArrayView_OperatorEqual4")
	symaddr(&addr_QByteArrayView_OperatorNotEqual4, "QByteArrayView_OperatorNotEqual4")
	symaddr(&addr_QByteArrayView_OperatorLesser4, "QByteArrayView_OperatorLesser4")
	symaddr(&addr_QByteArrayView_OperatorGreater4, "QByteArrayView_OperatorGreater4")
	symaddr(&addr_QByteArrayView_OperatorLesserOrEqual4, "QByteArrayView_OperatorLesserOrEqual4")
	symaddr(&addr_QByteArrayView_OperatorGreaterOrEqual4, "QByteArrayView_OperatorGreaterOrEqual4")
	reg(&pg_QByteArrayView_delete, "QByteArrayView_delete")
}

type QByteArrayView struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QByteArrayView) UnsafeGoID() int64     { return this.vid }
func (this *QByteArrayView) UnsafeSetGoID(id int64) { this.vid = id }

func bindQByteArrayView(h uintptr) *QByteArrayView {
	if h == 0 {
		return nil
	}
	return &QByteArrayView{h: h}
}

func newownedQByteArrayView(h uintptr) *QByteArrayView {
	o := bindQByteArrayView(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QByteArrayView).Delete)
	}
	return o
}

func (this *QByteArrayView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QByteArrayView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQByteArrayView(h unsafe.Pointer) *QByteArrayView { return bindQByteArrayView(uintptr(h)) }
func UnsafeNewQByteArrayViewOwned(h unsafe.Pointer) *QByteArrayView { return newownedQByteArrayView(uintptr(h)) }

func (this *QByteArrayView) UnsafeRelease() { this.h = 0 }

func NewQByteArrayView() *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_new)
	__o := newownedQByteArrayView(__r)
	return __o
}

func NewQByteArrayView2(arg1 *QByteArrayView) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQByteArrayView(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QByteArrayView) ToByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_toByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArrayView) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Data() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_data, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArrayView) ConstData() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_constData, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArrayView) OperatorSubscript(n int64) int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorSubscript, this.h, uintptr(n))
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) At(n int64) int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_at, this.h, uintptr(n))
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) First(n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_first, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Last(n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_last, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Sliced(pos int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_sliced, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Sliced2(pos int64, n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_sliced2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Slice(pos int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_slice, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQByteArrayView(__r)
}

func (this *QByteArrayView) Slice2(pos int64, n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_slice2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return bindQByteArrayView(__r)
}

func (this *QByteArrayView) Chopped(len_ int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_chopped, this.h, uintptr(len_))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Left(n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_left, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Right(n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_right, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Mid(pos int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_mid, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) Truncate(n int64) {
	purego.SyscallN(addr_QByteArrayView_truncate, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QByteArrayView) Chop(n int64) {
	purego.SyscallN(addr_QByteArrayView_chop, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QByteArrayView) Trimmed() *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_trimmed, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) ToShort() int16 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) ToUShort() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toUShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QByteArrayView) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QByteArrayView) ToUInt() uint {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toUInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QByteArrayView) ToLong() int32 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toLong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) ToULong() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toULong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QByteArrayView) ToLongLong() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toLongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) ToULongLong() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toULongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QByteArrayView) ToFloat() float32 {
	__r := pg_QByteArrayView_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) ToDouble() float64 {
	__r := pg_QByteArrayView_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) StartsWith(other *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_startsWith, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QByteArrayView) StartsWithWithChar(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_startsWith2, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) EndsWith(other *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_endsWith, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QByteArrayView) EndsWithWithChar(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_endsWith2, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) IndexOf(a *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_indexOf, this.h, uintptr(a.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return __r
}

func (this *QByteArrayView) IndexOfWithCh(ch int8) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_indexOf2, this.h, uintptr(ch))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Contains(a *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_contains, this.h, uintptr(a.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return __r != 0
}

func (this *QByteArrayView) ContainsWithChar(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_contains2, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) LastIndexOf(a *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_lastIndexOf, this.h, uintptr(a.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return __r
}

func (this *QByteArrayView) LastIndexOf2(a *QByteArrayView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_lastIndexOf2, this.h, uintptr(a.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return __r
}

func (this *QByteArrayView) LastIndexOfWithCh(ch int8) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_lastIndexOf3, this.h, uintptr(ch))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Count(a *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_count, this.h, uintptr(a.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return __r
}

func (this *QByteArrayView) CountWithCh(ch int8) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_count2, this.h, uintptr(ch))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Compare(a *QByteArrayView) int {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_compare, this.h, uintptr(a.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return int(__r)
}

func (this *QByteArrayView) IsValidUtf8() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_isValidUtf8, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) Begin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_begin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArrayView) End() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_end, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArrayView) Cbegin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_cbegin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArrayView) Cend() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_cend, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArrayView) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) Front() int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_front, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Back() int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_back, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArrayView) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) First2() int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_first2, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Last2() int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_last2, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func QByteArrayView_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QByteArrayView) OperatorAssign(arg1 *QByteArrayView) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQByteArrayView(__r)
}

func (this *QByteArrayView) Mid2(pos int64, n int64) *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QByteArrayView_mid2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QByteArrayView) ToShortWithOk() (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToShort2(base int) (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToUShortWithOk() (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toUShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QByteArrayView) ToUShort2(base int) (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toUShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QByteArrayView) ToIntWithOk() (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QByteArrayView) ToInt2(base int) (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QByteArrayView) ToUIntWithOk() (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toUInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QByteArrayView) ToUInt2(base int) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toUInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QByteArrayView) ToLongWithOk() (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToLong2(base int) (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToULongWithOk() (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toULong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QByteArrayView) ToULong2(base int) (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toULong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QByteArrayView) ToLongLongWithOk() (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toLongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToLongLong2(base int) (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toLongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToULongLongWithOk() (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toULongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QByteArrayView) ToULongLong2(base int) (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_toULongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QByteArrayView) ToFloatWithOk() (float32, bool) {
	var ok_out int32
	__r := pg_QByteArrayView_toFloat2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) ToDoubleWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QByteArrayView_toDouble2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArrayView) IndexOf2(a *QByteArrayView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_indexOf3, this.h, uintptr(a.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return __r
}

func (this *QByteArrayView) IndexOf3(ch int8, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_indexOf4, this.h, uintptr(ch), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) LastIndexOf3(ch int8, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_lastIndexOf4, this.h, uintptr(ch), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArrayView) Compare2(a *QByteArrayView, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_compare2, this.h, uintptr(a.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return int(__r)
}

func (this *QByteArrayView) OperatorEqual(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorNotEqual(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesser(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreater(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesserOrEqual(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreaterOrEqual(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorNotEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesserWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreaterWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesserOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreaterOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorNotEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesser2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreater2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesserOrEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreaterOrEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorNotEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorNotEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesser3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesser4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreater3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreater4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorLesserOrEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorLesserOrEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) OperatorGreaterOrEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArrayView_OperatorGreaterOrEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArrayView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QByteArrayView_delete(this.h)
	}
	this.UnsafeRelease()
}
