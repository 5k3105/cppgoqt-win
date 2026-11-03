package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringView_new uintptr
	addr_QStringView_new2 uintptr
	addr_QStringView_toString uintptr
	addr_QStringView_size uintptr
	addr_QStringView_data uintptr
	addr_QStringView_constData uintptr
	addr_QStringView_OperatorSubscript uintptr
	addr_QStringView_toLatin1 uintptr
	addr_QStringView_toUtf8 uintptr
	addr_QStringView_toLocal8Bit uintptr
	addr_QStringView_toUcs4 uintptr
	addr_QStringView_at uintptr
	addr_QStringView_mid uintptr
	addr_QStringView_left uintptr
	addr_QStringView_right uintptr
	addr_QStringView_first uintptr
	addr_QStringView_last uintptr
	addr_QStringView_sliced uintptr
	addr_QStringView_sliced2 uintptr
	addr_QStringView_chopped uintptr
	addr_QStringView_truncate uintptr
	addr_QStringView_chop uintptr
	addr_QStringView_trimmed uintptr
	addr_QStringView_slice uintptr
	addr_QStringView_slice2 uintptr
	addr_QStringView_compare uintptr
	addr_QStringView_compare2 uintptr
	addr_QStringView_compare3 uintptr
	addr_QStringView_compare4 uintptr
	addr_QStringView_localeAwareCompare uintptr
	addr_QStringView_startsWith uintptr
	addr_QStringView_startsWith2 uintptr
	addr_QStringView_startsWith3 uintptr
	addr_QStringView_startsWith4 uintptr
	addr_QStringView_endsWith uintptr
	addr_QStringView_endsWith2 uintptr
	addr_QStringView_endsWith3 uintptr
	addr_QStringView_endsWith4 uintptr
	addr_QStringView_indexOf uintptr
	addr_QStringView_indexOf2 uintptr
	addr_QStringView_indexOf3 uintptr
	addr_QStringView_contains uintptr
	addr_QStringView_contains2 uintptr
	addr_QStringView_contains3 uintptr
	addr_QStringView_count uintptr
	addr_QStringView_count2 uintptr
	addr_QStringView_count3 uintptr
	addr_QStringView_lastIndexOf uintptr
	addr_QStringView_lastIndexOf2 uintptr
	addr_QStringView_lastIndexOf3 uintptr
	addr_QStringView_lastIndexOf4 uintptr
	addr_QStringView_lastIndexOf5 uintptr
	addr_QStringView_lastIndexOf6 uintptr
	addr_QStringView_indexOf4 uintptr
	addr_QStringView_lastIndexOf7 uintptr
	addr_QStringView_contains4 uintptr
	addr_QStringView_count4 uintptr
	addr_QStringView_isRightToLeft uintptr
	addr_QStringView_isValidUtf16 uintptr
	addr_QStringView_isUpper uintptr
	addr_QStringView_isLower uintptr
	addr_QStringView_toShort uintptr
	addr_QStringView_toUShort uintptr
	addr_QStringView_toInt uintptr
	addr_QStringView_toUInt uintptr
	addr_QStringView_toLong uintptr
	addr_QStringView_toULong uintptr
	addr_QStringView_toLongLong uintptr
	addr_QStringView_toULongLong uintptr
	pg_QStringView_toFloat func(uintptr) float32
	pg_QStringView_toDouble func(uintptr) float64
	addr_QStringView_split uintptr
	addr_QStringView_split2 uintptr
	addr_QStringView_split3 uintptr
	addr_QStringView_begin uintptr
	addr_QStringView_end uintptr
	addr_QStringView_cbegin uintptr
	addr_QStringView_cend uintptr
	addr_QStringView_empty uintptr
	addr_QStringView_front uintptr
	addr_QStringView_back uintptr
	addr_QStringView_max_size uintptr
	addr_QStringView_constBegin uintptr
	addr_QStringView_constEnd uintptr
	addr_QStringView_isNull uintptr
	addr_QStringView_isEmpty uintptr
	addr_QStringView_length uintptr
	addr_QStringView_first2 uintptr
	addr_QStringView_last2 uintptr
	addr_QStringView_maxSize uintptr
	addr_QStringView_OperatorAssign uintptr
	addr_QStringView_mid2 uintptr
	addr_QStringView_compare5 uintptr
	addr_QStringView_compare6 uintptr
	addr_QStringView_startsWith5 uintptr
	addr_QStringView_startsWith6 uintptr
	addr_QStringView_endsWith5 uintptr
	addr_QStringView_endsWith6 uintptr
	addr_QStringView_indexOf5 uintptr
	addr_QStringView_indexOf6 uintptr
	addr_QStringView_indexOf7 uintptr
	addr_QStringView_indexOf8 uintptr
	addr_QStringView_indexOf9 uintptr
	addr_QStringView_indexOf10 uintptr
	addr_QStringView_contains5 uintptr
	addr_QStringView_contains6 uintptr
	addr_QStringView_contains7 uintptr
	addr_QStringView_count5 uintptr
	addr_QStringView_count6 uintptr
	addr_QStringView_count7 uintptr
	addr_QStringView_lastIndexOf8 uintptr
	addr_QStringView_lastIndexOf9 uintptr
	addr_QStringView_lastIndexOf10 uintptr
	addr_QStringView_lastIndexOf11 uintptr
	addr_QStringView_lastIndexOf12 uintptr
	addr_QStringView_lastIndexOf13 uintptr
	addr_QStringView_indexOf11 uintptr
	addr_QStringView_indexOf12 uintptr
	addr_QStringView_lastIndexOf14 uintptr
	addr_QStringView_contains8 uintptr
	addr_QStringView_toShort2 uintptr
	addr_QStringView_toShort3 uintptr
	addr_QStringView_toUShort2 uintptr
	addr_QStringView_toUShort3 uintptr
	addr_QStringView_toInt2 uintptr
	addr_QStringView_toInt3 uintptr
	addr_QStringView_toUInt2 uintptr
	addr_QStringView_toUInt3 uintptr
	addr_QStringView_toLong2 uintptr
	addr_QStringView_toLong3 uintptr
	addr_QStringView_toULong2 uintptr
	addr_QStringView_toULong3 uintptr
	addr_QStringView_toLongLong2 uintptr
	addr_QStringView_toLongLong3 uintptr
	addr_QStringView_toULongLong2 uintptr
	addr_QStringView_toULongLong3 uintptr
	pg_QStringView_toFloat2 func(uintptr, uintptr) float32
	pg_QStringView_toDouble2 func(uintptr, uintptr) float64
	addr_QStringView_split4 uintptr
	addr_QStringView_split5 uintptr
	addr_QStringView_split6 uintptr
	addr_QStringView_split7 uintptr
	addr_QStringView_split8 uintptr
	addr_QStringView_OperatorEqual uintptr
	addr_QStringView_OperatorNotEqual uintptr
	addr_QStringView_OperatorLesser uintptr
	addr_QStringView_OperatorGreater uintptr
	addr_QStringView_OperatorLesserOrEqual uintptr
	addr_QStringView_OperatorGreaterOrEqual uintptr
	addr_QStringView_OperatorEqual2 uintptr
	addr_QStringView_OperatorNotEqual2 uintptr
	addr_QStringView_OperatorLesser2 uintptr
	addr_QStringView_OperatorGreater2 uintptr
	addr_QStringView_OperatorLesserOrEqual2 uintptr
	addr_QStringView_OperatorGreaterOrEqual2 uintptr
	addr_QStringView_OperatorEqual3 uintptr
	addr_QStringView_OperatorNotEqual3 uintptr
	addr_QStringView_OperatorLesser3 uintptr
	addr_QStringView_OperatorGreater3 uintptr
	addr_QStringView_OperatorLesserOrEqual3 uintptr
	addr_QStringView_OperatorGreaterOrEqual3 uintptr
	addr_QStringView_OperatorEqual4 uintptr
	addr_QStringView_OperatorNotEqual4 uintptr
	addr_QStringView_OperatorLesser4 uintptr
	addr_QStringView_OperatorGreater4 uintptr
	addr_QStringView_OperatorLesserOrEqual4 uintptr
	addr_QStringView_OperatorGreaterOrEqual4 uintptr
	addr_QStringView_OperatorPlus uintptr
	pg_QStringView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStringView_new, "QStringView_new")
	symaddr(&addr_QStringView_new2, "QStringView_new2")
	symaddr(&addr_QStringView_toString, "QStringView_toString")
	symaddr(&addr_QStringView_size, "QStringView_size")
	symaddr(&addr_QStringView_data, "QStringView_data")
	symaddr(&addr_QStringView_constData, "QStringView_constData")
	symaddr(&addr_QStringView_OperatorSubscript, "QStringView_OperatorSubscript")
	symaddr(&addr_QStringView_toLatin1, "QStringView_toLatin1")
	symaddr(&addr_QStringView_toUtf8, "QStringView_toUtf8")
	symaddr(&addr_QStringView_toLocal8Bit, "QStringView_toLocal8Bit")
	symaddr(&addr_QStringView_toUcs4, "QStringView_toUcs4")
	symaddr(&addr_QStringView_at, "QStringView_at")
	symaddr(&addr_QStringView_mid, "QStringView_mid")
	symaddr(&addr_QStringView_left, "QStringView_left")
	symaddr(&addr_QStringView_right, "QStringView_right")
	symaddr(&addr_QStringView_first, "QStringView_first")
	symaddr(&addr_QStringView_last, "QStringView_last")
	symaddr(&addr_QStringView_sliced, "QStringView_sliced")
	symaddr(&addr_QStringView_sliced2, "QStringView_sliced2")
	symaddr(&addr_QStringView_chopped, "QStringView_chopped")
	symaddr(&addr_QStringView_truncate, "QStringView_truncate")
	symaddr(&addr_QStringView_chop, "QStringView_chop")
	symaddr(&addr_QStringView_trimmed, "QStringView_trimmed")
	symaddr(&addr_QStringView_slice, "QStringView_slice")
	symaddr(&addr_QStringView_slice2, "QStringView_slice2")
	symaddr(&addr_QStringView_compare, "QStringView_compare")
	symaddr(&addr_QStringView_compare2, "QStringView_compare2")
	symaddr(&addr_QStringView_compare3, "QStringView_compare3")
	symaddr(&addr_QStringView_compare4, "QStringView_compare4")
	symaddr(&addr_QStringView_localeAwareCompare, "QStringView_localeAwareCompare")
	symaddr(&addr_QStringView_startsWith, "QStringView_startsWith")
	symaddr(&addr_QStringView_startsWith2, "QStringView_startsWith2")
	symaddr(&addr_QStringView_startsWith3, "QStringView_startsWith3")
	symaddr(&addr_QStringView_startsWith4, "QStringView_startsWith4")
	symaddr(&addr_QStringView_endsWith, "QStringView_endsWith")
	symaddr(&addr_QStringView_endsWith2, "QStringView_endsWith2")
	symaddr(&addr_QStringView_endsWith3, "QStringView_endsWith3")
	symaddr(&addr_QStringView_endsWith4, "QStringView_endsWith4")
	symaddr(&addr_QStringView_indexOf, "QStringView_indexOf")
	symaddr(&addr_QStringView_indexOf2, "QStringView_indexOf2")
	symaddr(&addr_QStringView_indexOf3, "QStringView_indexOf3")
	symaddr(&addr_QStringView_contains, "QStringView_contains")
	symaddr(&addr_QStringView_contains2, "QStringView_contains2")
	symaddr(&addr_QStringView_contains3, "QStringView_contains3")
	symaddr(&addr_QStringView_count, "QStringView_count")
	symaddr(&addr_QStringView_count2, "QStringView_count2")
	symaddr(&addr_QStringView_count3, "QStringView_count3")
	symaddr(&addr_QStringView_lastIndexOf, "QStringView_lastIndexOf")
	symaddr(&addr_QStringView_lastIndexOf2, "QStringView_lastIndexOf2")
	symaddr(&addr_QStringView_lastIndexOf3, "QStringView_lastIndexOf3")
	symaddr(&addr_QStringView_lastIndexOf4, "QStringView_lastIndexOf4")
	symaddr(&addr_QStringView_lastIndexOf5, "QStringView_lastIndexOf5")
	symaddr(&addr_QStringView_lastIndexOf6, "QStringView_lastIndexOf6")
	symaddr(&addr_QStringView_indexOf4, "QStringView_indexOf4")
	symaddr(&addr_QStringView_lastIndexOf7, "QStringView_lastIndexOf7")
	symaddr(&addr_QStringView_contains4, "QStringView_contains4")
	symaddr(&addr_QStringView_count4, "QStringView_count4")
	symaddr(&addr_QStringView_isRightToLeft, "QStringView_isRightToLeft")
	symaddr(&addr_QStringView_isValidUtf16, "QStringView_isValidUtf16")
	symaddr(&addr_QStringView_isUpper, "QStringView_isUpper")
	symaddr(&addr_QStringView_isLower, "QStringView_isLower")
	symaddr(&addr_QStringView_toShort, "QStringView_toShort")
	symaddr(&addr_QStringView_toUShort, "QStringView_toUShort")
	symaddr(&addr_QStringView_toInt, "QStringView_toInt")
	symaddr(&addr_QStringView_toUInt, "QStringView_toUInt")
	symaddr(&addr_QStringView_toLong, "QStringView_toLong")
	symaddr(&addr_QStringView_toULong, "QStringView_toULong")
	symaddr(&addr_QStringView_toLongLong, "QStringView_toLongLong")
	symaddr(&addr_QStringView_toULongLong, "QStringView_toULongLong")
	reg(&pg_QStringView_toFloat, "QStringView_toFloat")
	reg(&pg_QStringView_toDouble, "QStringView_toDouble")
	symaddr(&addr_QStringView_split, "QStringView_split")
	symaddr(&addr_QStringView_split2, "QStringView_split2")
	symaddr(&addr_QStringView_split3, "QStringView_split3")
	symaddr(&addr_QStringView_begin, "QStringView_begin")
	symaddr(&addr_QStringView_end, "QStringView_end")
	symaddr(&addr_QStringView_cbegin, "QStringView_cbegin")
	symaddr(&addr_QStringView_cend, "QStringView_cend")
	symaddr(&addr_QStringView_empty, "QStringView_empty")
	symaddr(&addr_QStringView_front, "QStringView_front")
	symaddr(&addr_QStringView_back, "QStringView_back")
	symaddr(&addr_QStringView_max_size, "QStringView_max_size")
	symaddr(&addr_QStringView_constBegin, "QStringView_constBegin")
	symaddr(&addr_QStringView_constEnd, "QStringView_constEnd")
	symaddr(&addr_QStringView_isNull, "QStringView_isNull")
	symaddr(&addr_QStringView_isEmpty, "QStringView_isEmpty")
	symaddr(&addr_QStringView_length, "QStringView_length")
	symaddr(&addr_QStringView_first2, "QStringView_first2")
	symaddr(&addr_QStringView_last2, "QStringView_last2")
	symaddr(&addr_QStringView_maxSize, "QStringView_maxSize")
	symaddr(&addr_QStringView_OperatorAssign, "QStringView_OperatorAssign")
	symaddr(&addr_QStringView_mid2, "QStringView_mid2")
	symaddr(&addr_QStringView_compare5, "QStringView_compare5")
	symaddr(&addr_QStringView_compare6, "QStringView_compare6")
	symaddr(&addr_QStringView_startsWith5, "QStringView_startsWith5")
	symaddr(&addr_QStringView_startsWith6, "QStringView_startsWith6")
	symaddr(&addr_QStringView_endsWith5, "QStringView_endsWith5")
	symaddr(&addr_QStringView_endsWith6, "QStringView_endsWith6")
	symaddr(&addr_QStringView_indexOf5, "QStringView_indexOf5")
	symaddr(&addr_QStringView_indexOf6, "QStringView_indexOf6")
	symaddr(&addr_QStringView_indexOf7, "QStringView_indexOf7")
	symaddr(&addr_QStringView_indexOf8, "QStringView_indexOf8")
	symaddr(&addr_QStringView_indexOf9, "QStringView_indexOf9")
	symaddr(&addr_QStringView_indexOf10, "QStringView_indexOf10")
	symaddr(&addr_QStringView_contains5, "QStringView_contains5")
	symaddr(&addr_QStringView_contains6, "QStringView_contains6")
	symaddr(&addr_QStringView_contains7, "QStringView_contains7")
	symaddr(&addr_QStringView_count5, "QStringView_count5")
	symaddr(&addr_QStringView_count6, "QStringView_count6")
	symaddr(&addr_QStringView_count7, "QStringView_count7")
	symaddr(&addr_QStringView_lastIndexOf8, "QStringView_lastIndexOf8")
	symaddr(&addr_QStringView_lastIndexOf9, "QStringView_lastIndexOf9")
	symaddr(&addr_QStringView_lastIndexOf10, "QStringView_lastIndexOf10")
	symaddr(&addr_QStringView_lastIndexOf11, "QStringView_lastIndexOf11")
	symaddr(&addr_QStringView_lastIndexOf12, "QStringView_lastIndexOf12")
	symaddr(&addr_QStringView_lastIndexOf13, "QStringView_lastIndexOf13")
	symaddr(&addr_QStringView_indexOf11, "QStringView_indexOf11")
	symaddr(&addr_QStringView_indexOf12, "QStringView_indexOf12")
	symaddr(&addr_QStringView_lastIndexOf14, "QStringView_lastIndexOf14")
	symaddr(&addr_QStringView_contains8, "QStringView_contains8")
	symaddr(&addr_QStringView_toShort2, "QStringView_toShort2")
	symaddr(&addr_QStringView_toShort3, "QStringView_toShort3")
	symaddr(&addr_QStringView_toUShort2, "QStringView_toUShort2")
	symaddr(&addr_QStringView_toUShort3, "QStringView_toUShort3")
	symaddr(&addr_QStringView_toInt2, "QStringView_toInt2")
	symaddr(&addr_QStringView_toInt3, "QStringView_toInt3")
	symaddr(&addr_QStringView_toUInt2, "QStringView_toUInt2")
	symaddr(&addr_QStringView_toUInt3, "QStringView_toUInt3")
	symaddr(&addr_QStringView_toLong2, "QStringView_toLong2")
	symaddr(&addr_QStringView_toLong3, "QStringView_toLong3")
	symaddr(&addr_QStringView_toULong2, "QStringView_toULong2")
	symaddr(&addr_QStringView_toULong3, "QStringView_toULong3")
	symaddr(&addr_QStringView_toLongLong2, "QStringView_toLongLong2")
	symaddr(&addr_QStringView_toLongLong3, "QStringView_toLongLong3")
	symaddr(&addr_QStringView_toULongLong2, "QStringView_toULongLong2")
	symaddr(&addr_QStringView_toULongLong3, "QStringView_toULongLong3")
	reg(&pg_QStringView_toFloat2, "QStringView_toFloat2")
	reg(&pg_QStringView_toDouble2, "QStringView_toDouble2")
	symaddr(&addr_QStringView_split4, "QStringView_split4")
	symaddr(&addr_QStringView_split5, "QStringView_split5")
	symaddr(&addr_QStringView_split6, "QStringView_split6")
	symaddr(&addr_QStringView_split7, "QStringView_split7")
	symaddr(&addr_QStringView_split8, "QStringView_split8")
	symaddr(&addr_QStringView_OperatorEqual, "QStringView_OperatorEqual")
	symaddr(&addr_QStringView_OperatorNotEqual, "QStringView_OperatorNotEqual")
	symaddr(&addr_QStringView_OperatorLesser, "QStringView_OperatorLesser")
	symaddr(&addr_QStringView_OperatorGreater, "QStringView_OperatorGreater")
	symaddr(&addr_QStringView_OperatorLesserOrEqual, "QStringView_OperatorLesserOrEqual")
	symaddr(&addr_QStringView_OperatorGreaterOrEqual, "QStringView_OperatorGreaterOrEqual")
	symaddr(&addr_QStringView_OperatorEqual2, "QStringView_OperatorEqual2")
	symaddr(&addr_QStringView_OperatorNotEqual2, "QStringView_OperatorNotEqual2")
	symaddr(&addr_QStringView_OperatorLesser2, "QStringView_OperatorLesser2")
	symaddr(&addr_QStringView_OperatorGreater2, "QStringView_OperatorGreater2")
	symaddr(&addr_QStringView_OperatorLesserOrEqual2, "QStringView_OperatorLesserOrEqual2")
	symaddr(&addr_QStringView_OperatorGreaterOrEqual2, "QStringView_OperatorGreaterOrEqual2")
	symaddr(&addr_QStringView_OperatorEqual3, "QStringView_OperatorEqual3")
	symaddr(&addr_QStringView_OperatorNotEqual3, "QStringView_OperatorNotEqual3")
	symaddr(&addr_QStringView_OperatorLesser3, "QStringView_OperatorLesser3")
	symaddr(&addr_QStringView_OperatorGreater3, "QStringView_OperatorGreater3")
	symaddr(&addr_QStringView_OperatorLesserOrEqual3, "QStringView_OperatorLesserOrEqual3")
	symaddr(&addr_QStringView_OperatorGreaterOrEqual3, "QStringView_OperatorGreaterOrEqual3")
	symaddr(&addr_QStringView_OperatorEqual4, "QStringView_OperatorEqual4")
	symaddr(&addr_QStringView_OperatorNotEqual4, "QStringView_OperatorNotEqual4")
	symaddr(&addr_QStringView_OperatorLesser4, "QStringView_OperatorLesser4")
	symaddr(&addr_QStringView_OperatorGreater4, "QStringView_OperatorGreater4")
	symaddr(&addr_QStringView_OperatorLesserOrEqual4, "QStringView_OperatorLesserOrEqual4")
	symaddr(&addr_QStringView_OperatorGreaterOrEqual4, "QStringView_OperatorGreaterOrEqual4")
	symaddr(&addr_QStringView_OperatorPlus, "QStringView_OperatorPlus")
	reg(&pg_QStringView_delete, "QStringView_delete")
}

type QStringView struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStringView) UnsafeGoID() int64     { return this.vid }
func (this *QStringView) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStringView(h uintptr) *QStringView {
	if h == 0 {
		return nil
	}
	return &QStringView{h: h}
}

func newownedQStringView(h uintptr) *QStringView {
	o := bindQStringView(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStringView).Delete)
	}
	return o
}

func (this *QStringView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringView(h unsafe.Pointer) *QStringView { return bindQStringView(uintptr(h)) }
func UnsafeNewQStringViewOwned(h unsafe.Pointer) *QStringView { return newownedQStringView(uintptr(h)) }

func (this *QStringView) UnsafeRelease() { this.h = 0 }

func NewQStringView() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_new)
	__o := newownedQStringView(__r)
	return __o
}

func NewQStringView2(arg1 *QStringView) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQStringView(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QStringView) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStringView) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) Data() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_data, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) ConstData() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_constData, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) OperatorSubscript(n int64) *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_OperatorSubscript, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QStringView) ToLatin1() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QStringView_toLatin1, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QStringView) ToUtf8() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QStringView_toUtf8, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QStringView) ToLocal8Bit() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QStringView_toLocal8Bit, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QStringView) ToUcs4() []uint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_toUcs4, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]uint, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = uint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) At(n int64) *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_at, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QStringView) Mid(pos int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_mid, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Left(n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_left, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Right(n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_right, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) First(n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_first, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Last(n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_last, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Sliced(pos int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_sliced, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Sliced2(pos int64, n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_sliced2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Chopped(n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_chopped, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Truncate(n int64) {
	purego.SyscallN(addr_QStringView_truncate, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QStringView) Chop(n int64) {
	purego.SyscallN(addr_QStringView_chop, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QStringView) Trimmed() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_trimmed, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Slice(pos int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_slice, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQStringView(__r)
}

func (this *QStringView) Slice2(pos int64, n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_slice2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return bindQStringView(__r)
}

func (this *QStringView) Compare(other *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QStringView) CompareWithOther(other *QLatin1String) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_compare2, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QStringView) CompareWithQChar(c *QChar) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_compare3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return int(__r)
}

func (this *QStringView) Compare2(c *QChar, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_compare4, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return int(__r)
}

func (this *QStringView) LocaleAwareCompare(other *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_localeAwareCompare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QStringView) StartsWith(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_startsWith, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) StartsWithWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_startsWith2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) StartsWithWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_startsWith3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QStringView) StartsWith2(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_startsWith4, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QStringView) EndsWith(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_endsWith, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) EndsWithWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_endsWith2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) EndsWithWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_endsWith3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QStringView) EndsWith2(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_endsWith4, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QStringView) IndexOf(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) IndexOfWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf2, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) IndexOfWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf3, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) Contains(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QStringView) ContainsWithQStringView(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) ContainsWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains3, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) Count(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) CountWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count2, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) CountWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count3, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) LastIndexOf2(c *QChar, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf2, this.h, uintptr(c.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) LastIndexOfWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf3, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf3(s *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf4, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOfWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf5, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf4(s *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf6, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) IndexOfWithRe(re *QRegularExpression) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf4, this.h, uintptr(re.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QStringView) LastIndexOf5(re *QRegularExpression, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf7, this.h, uintptr(re.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QStringView) ContainsWithRe(re *QRegularExpression) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains4, this.h, uintptr(re.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r != 0
}

func (this *QStringView) CountWithRe(re *QRegularExpression) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count4, this.h, uintptr(re.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QStringView) IsRightToLeft() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_isRightToLeft, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) IsValidUtf16() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_isValidUtf16, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) IsUpper() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_isUpper, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) IsLower() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_isLower, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) ToShort() int16 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) ToUShort() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toUShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QStringView) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStringView) ToUInt() uint {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toUInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QStringView) ToLong() int32 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toLong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) ToULong() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toULong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QStringView) ToLongLong() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toLongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) ToULongLong() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_toULongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QStringView) ToFloat() float32 {
	__r := pg_QStringView_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) ToDouble() float64 {
	__r := pg_QStringView_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) Split(sep *QStringView) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split, this.h, uintptr(sep.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) SplitWithSep(sep *QChar) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split2, this.h, uintptr(sep.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) Split2(sep *QRegularExpression) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split3, this.h, uintptr(sep.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) Begin() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_begin, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) End() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_end, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) Cbegin() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_cbegin, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) Cend() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_cend, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) Front() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_front, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QStringView) Back() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_back, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QStringView) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) ConstBegin() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_constBegin, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) ConstEnd() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_constEnd, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QStringView) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringView) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringView) First2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_first2, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QStringView) Last2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringView_last2, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func QStringView_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QStringView) OperatorAssign(arg1 *QStringView) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStringView(__r)
}

func (this *QStringView) Mid2(pos int64, n int64) *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringView_mid2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringView) Compare3(other *QStringView, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_compare5, this.h, uintptr(other.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QStringView) Compare4(other *QLatin1String, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QStringView_compare6, this.h, uintptr(other.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QStringView) StartsWith3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_startsWith5, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) StartsWith4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_startsWith6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) EndsWith3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_endsWith5, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) EndsWith4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_endsWith6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) IndexOf2(c *QChar, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf5, this.h, uintptr(c.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) IndexOf3(c *QChar, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf6, this.h, uintptr(c.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) IndexOf4(s *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf7, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) IndexOf5(s *QStringView, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf8, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) IndexOf6(s *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf9, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) IndexOf7(s *QLatin1String, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf10, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) Contains2(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains5, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QStringView) Contains3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) Contains4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains7, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QStringView) Count2(c *QChar, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count5, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) Count3(s *QStringView, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) Count4(s *QLatin1String, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_count7, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf6(c *QChar, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf8, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) LastIndexOf7(c *QChar, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf9, this.h, uintptr(c.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QStringView) LastIndexOf8(s *QStringView, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf10, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf9(s *QStringView, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf11, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf10(s *QLatin1String, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf12, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) LastIndexOf11(s *QLatin1String, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf13, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QStringView) IndexOf8(re *QRegularExpression, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf11, this.h, uintptr(re.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QStringView) IndexOf9(re *QRegularExpression, from int64, rmatch *QRegularExpressionMatch) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_indexOf12, this.h, uintptr(re.UnsafePointer()), uintptr(from), uintptr(rmatch.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(rmatch)
	return __r
}

func (this *QStringView) LastIndexOf12(re *QRegularExpression, from int64, rmatch *QRegularExpressionMatch) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringView_lastIndexOf14, this.h, uintptr(re.UnsafePointer()), uintptr(from), uintptr(rmatch.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(rmatch)
	return __r
}

func (this *QStringView) Contains5(re *QRegularExpression, rmatch *QRegularExpressionMatch) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_contains8, this.h, uintptr(re.UnsafePointer()), uintptr(rmatch.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(rmatch)
	return __r != 0
}

func (this *QStringView) ToShortWithOk() (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToShort2(base int) (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToUShortWithOk() (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toUShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QStringView) ToUShort2(base int) (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toUShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QStringView) ToIntWithOk() (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QStringView) ToInt2(base int) (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QStringView) ToUIntWithOk() (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toUInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QStringView) ToUInt2(base int) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toUInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QStringView) ToLongWithOk() (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToLong2(base int) (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToULongWithOk() (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toULong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QStringView) ToULong2(base int) (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toULong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QStringView) ToLongLongWithOk() (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toLongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToLongLong2(base int) (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toLongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToULongLongWithOk() (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toULongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QStringView) ToULongLong2(base int) (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QStringView_toULongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QStringView) ToFloatWithOk() (float32, bool) {
	var ok_out int32
	__r := pg_QStringView_toFloat2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) ToDoubleWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QStringView_toDouble2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QStringView) Split3(sep *QStringView, behavior SplitBehaviorFlags) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split4, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) Split4(sep *QStringView, behavior SplitBehaviorFlags, cs CaseSensitivity) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split5, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(cs), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) Split5(sep *QChar, behavior SplitBehaviorFlags) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split6, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) Split6(sep *QChar, behavior SplitBehaviorFlags, cs CaseSensitivity) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split7, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(cs), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) Split7(sep *QRegularExpression, behavior SplitBehaviorFlags) []*QStringView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_split8, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]*QStringView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStringView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStringView) OperatorEqual(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorNotEqual(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesser(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreater(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesserOrEqual(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreaterOrEqual(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorNotEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesserWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreaterWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesserOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreaterOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorNotEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesser2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreater2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesserOrEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreaterOrEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorNotEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesser3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesser4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreater3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreater4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorLesserOrEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorLesserOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorGreaterOrEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringView_OperatorGreaterOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStringView) OperatorPlus(rhs string) string {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringView_OperatorPlus, this.h, rhs_p, uintptr(len(rhs)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return pgStr(__r, __n)
}

func (this *QStringView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStringView_delete(this.h)
	}
	this.UnsafeRelease()
}
