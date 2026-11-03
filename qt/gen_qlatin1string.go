package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLatin1String_new uintptr
	addr_QLatin1String_new2 uintptr
	addr_QLatin1String_new3 uintptr
	addr_QLatin1String_new4 uintptr
	addr_QLatin1String_new5 uintptr
	addr_QLatin1String_new6 uintptr
	addr_QLatin1String_new7 uintptr
	addr_QLatin1String_toString uintptr
	addr_QLatin1String_toUtf8 uintptr
	addr_QLatin1String_latin1 uintptr
	addr_QLatin1String_size uintptr
	addr_QLatin1String_data uintptr
	addr_QLatin1String_constData uintptr
	addr_QLatin1String_constBegin uintptr
	addr_QLatin1String_constEnd uintptr
	addr_QLatin1String_first uintptr
	addr_QLatin1String_last uintptr
	addr_QLatin1String_length uintptr
	addr_QLatin1String_isNull uintptr
	addr_QLatin1String_isEmpty uintptr
	addr_QLatin1String_empty uintptr
	addr_QLatin1String_at uintptr
	addr_QLatin1String_OperatorSubscript uintptr
	addr_QLatin1String_front uintptr
	addr_QLatin1String_back uintptr
	addr_QLatin1String_compare uintptr
	addr_QLatin1String_compare2 uintptr
	addr_QLatin1String_compare3 uintptr
	addr_QLatin1String_compare4 uintptr
	addr_QLatin1String_startsWith uintptr
	addr_QLatin1String_startsWith2 uintptr
	addr_QLatin1String_startsWith3 uintptr
	addr_QLatin1String_startsWith4 uintptr
	addr_QLatin1String_endsWith uintptr
	addr_QLatin1String_endsWith2 uintptr
	addr_QLatin1String_endsWith3 uintptr
	addr_QLatin1String_endsWith4 uintptr
	addr_QLatin1String_indexOf uintptr
	addr_QLatin1String_indexOf2 uintptr
	addr_QLatin1String_indexOf3 uintptr
	addr_QLatin1String_indexOf4 uintptr
	addr_QLatin1String_contains uintptr
	addr_QLatin1String_contains2 uintptr
	addr_QLatin1String_contains3 uintptr
	addr_QLatin1String_lastIndexOf uintptr
	addr_QLatin1String_lastIndexOf2 uintptr
	addr_QLatin1String_lastIndexOf3 uintptr
	addr_QLatin1String_lastIndexOf4 uintptr
	addr_QLatin1String_lastIndexOf5 uintptr
	addr_QLatin1String_lastIndexOf6 uintptr
	addr_QLatin1String_lastIndexOf7 uintptr
	addr_QLatin1String_lastIndexOf8 uintptr
	addr_QLatin1String_count uintptr
	addr_QLatin1String_count2 uintptr
	addr_QLatin1String_count3 uintptr
	addr_QLatin1String_toShort uintptr
	addr_QLatin1String_toUShort uintptr
	addr_QLatin1String_toInt uintptr
	addr_QLatin1String_toUInt uintptr
	addr_QLatin1String_toLong uintptr
	addr_QLatin1String_toULong uintptr
	addr_QLatin1String_toLongLong uintptr
	addr_QLatin1String_toULongLong uintptr
	pg_QLatin1String_toFloat func(uintptr) float32
	pg_QLatin1String_toDouble func(uintptr) float64
	addr_QLatin1String_begin uintptr
	addr_QLatin1String_cbegin uintptr
	addr_QLatin1String_end uintptr
	addr_QLatin1String_cend uintptr
	addr_QLatin1String_max_size uintptr
	addr_QLatin1String_maxSize uintptr
	addr_QLatin1String_mid uintptr
	addr_QLatin1String_left uintptr
	addr_QLatin1String_right uintptr
	addr_QLatin1String_sliced uintptr
	addr_QLatin1String_sliced2 uintptr
	addr_QLatin1String_first2 uintptr
	addr_QLatin1String_last2 uintptr
	addr_QLatin1String_chopped uintptr
	addr_QLatin1String_chop uintptr
	addr_QLatin1String_truncate uintptr
	addr_QLatin1String_trimmed uintptr
	addr_QLatin1String_OperatorAssign uintptr
	addr_QLatin1String_compare5 uintptr
	addr_QLatin1String_compare6 uintptr
	addr_QLatin1String_startsWith5 uintptr
	addr_QLatin1String_startsWith6 uintptr
	addr_QLatin1String_endsWith5 uintptr
	addr_QLatin1String_endsWith6 uintptr
	addr_QLatin1String_indexOf5 uintptr
	addr_QLatin1String_indexOf6 uintptr
	addr_QLatin1String_indexOf7 uintptr
	addr_QLatin1String_indexOf8 uintptr
	addr_QLatin1String_indexOf9 uintptr
	addr_QLatin1String_contains4 uintptr
	addr_QLatin1String_contains5 uintptr
	addr_QLatin1String_contains6 uintptr
	addr_QLatin1String_lastIndexOf9 uintptr
	addr_QLatin1String_lastIndexOf10 uintptr
	addr_QLatin1String_lastIndexOf11 uintptr
	addr_QLatin1String_lastIndexOf12 uintptr
	addr_QLatin1String_count4 uintptr
	addr_QLatin1String_count5 uintptr
	addr_QLatin1String_count6 uintptr
	addr_QLatin1String_toShort2 uintptr
	addr_QLatin1String_toShort3 uintptr
	addr_QLatin1String_toUShort2 uintptr
	addr_QLatin1String_toUShort3 uintptr
	addr_QLatin1String_toInt2 uintptr
	addr_QLatin1String_toInt3 uintptr
	addr_QLatin1String_toUInt2 uintptr
	addr_QLatin1String_toUInt3 uintptr
	addr_QLatin1String_toLong2 uintptr
	addr_QLatin1String_toLong3 uintptr
	addr_QLatin1String_toULong2 uintptr
	addr_QLatin1String_toULong3 uintptr
	addr_QLatin1String_toLongLong2 uintptr
	addr_QLatin1String_toLongLong3 uintptr
	addr_QLatin1String_toULongLong2 uintptr
	addr_QLatin1String_toULongLong3 uintptr
	pg_QLatin1String_toFloat2 func(uintptr, uintptr) float32
	pg_QLatin1String_toDouble2 func(uintptr, uintptr) float64
	addr_QLatin1String_mid2 uintptr
	pg_QLatin1String_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLatin1String_new, "QLatin1String_new")
	symaddr(&addr_QLatin1String_new2, "QLatin1String_new2")
	symaddr(&addr_QLatin1String_new3, "QLatin1String_new3")
	symaddr(&addr_QLatin1String_new4, "QLatin1String_new4")
	symaddr(&addr_QLatin1String_new5, "QLatin1String_new5")
	symaddr(&addr_QLatin1String_new6, "QLatin1String_new6")
	symaddr(&addr_QLatin1String_new7, "QLatin1String_new7")
	symaddr(&addr_QLatin1String_toString, "QLatin1String_toString")
	symaddr(&addr_QLatin1String_toUtf8, "QLatin1String_toUtf8")
	symaddr(&addr_QLatin1String_latin1, "QLatin1String_latin1")
	symaddr(&addr_QLatin1String_size, "QLatin1String_size")
	symaddr(&addr_QLatin1String_data, "QLatin1String_data")
	symaddr(&addr_QLatin1String_constData, "QLatin1String_constData")
	symaddr(&addr_QLatin1String_constBegin, "QLatin1String_constBegin")
	symaddr(&addr_QLatin1String_constEnd, "QLatin1String_constEnd")
	symaddr(&addr_QLatin1String_first, "QLatin1String_first")
	symaddr(&addr_QLatin1String_last, "QLatin1String_last")
	symaddr(&addr_QLatin1String_length, "QLatin1String_length")
	symaddr(&addr_QLatin1String_isNull, "QLatin1String_isNull")
	symaddr(&addr_QLatin1String_isEmpty, "QLatin1String_isEmpty")
	symaddr(&addr_QLatin1String_empty, "QLatin1String_empty")
	symaddr(&addr_QLatin1String_at, "QLatin1String_at")
	symaddr(&addr_QLatin1String_OperatorSubscript, "QLatin1String_OperatorSubscript")
	symaddr(&addr_QLatin1String_front, "QLatin1String_front")
	symaddr(&addr_QLatin1String_back, "QLatin1String_back")
	symaddr(&addr_QLatin1String_compare, "QLatin1String_compare")
	symaddr(&addr_QLatin1String_compare2, "QLatin1String_compare2")
	symaddr(&addr_QLatin1String_compare3, "QLatin1String_compare3")
	symaddr(&addr_QLatin1String_compare4, "QLatin1String_compare4")
	symaddr(&addr_QLatin1String_startsWith, "QLatin1String_startsWith")
	symaddr(&addr_QLatin1String_startsWith2, "QLatin1String_startsWith2")
	symaddr(&addr_QLatin1String_startsWith3, "QLatin1String_startsWith3")
	symaddr(&addr_QLatin1String_startsWith4, "QLatin1String_startsWith4")
	symaddr(&addr_QLatin1String_endsWith, "QLatin1String_endsWith")
	symaddr(&addr_QLatin1String_endsWith2, "QLatin1String_endsWith2")
	symaddr(&addr_QLatin1String_endsWith3, "QLatin1String_endsWith3")
	symaddr(&addr_QLatin1String_endsWith4, "QLatin1String_endsWith4")
	symaddr(&addr_QLatin1String_indexOf, "QLatin1String_indexOf")
	symaddr(&addr_QLatin1String_indexOf2, "QLatin1String_indexOf2")
	symaddr(&addr_QLatin1String_indexOf3, "QLatin1String_indexOf3")
	symaddr(&addr_QLatin1String_indexOf4, "QLatin1String_indexOf4")
	symaddr(&addr_QLatin1String_contains, "QLatin1String_contains")
	symaddr(&addr_QLatin1String_contains2, "QLatin1String_contains2")
	symaddr(&addr_QLatin1String_contains3, "QLatin1String_contains3")
	symaddr(&addr_QLatin1String_lastIndexOf, "QLatin1String_lastIndexOf")
	symaddr(&addr_QLatin1String_lastIndexOf2, "QLatin1String_lastIndexOf2")
	symaddr(&addr_QLatin1String_lastIndexOf3, "QLatin1String_lastIndexOf3")
	symaddr(&addr_QLatin1String_lastIndexOf4, "QLatin1String_lastIndexOf4")
	symaddr(&addr_QLatin1String_lastIndexOf5, "QLatin1String_lastIndexOf5")
	symaddr(&addr_QLatin1String_lastIndexOf6, "QLatin1String_lastIndexOf6")
	symaddr(&addr_QLatin1String_lastIndexOf7, "QLatin1String_lastIndexOf7")
	symaddr(&addr_QLatin1String_lastIndexOf8, "QLatin1String_lastIndexOf8")
	symaddr(&addr_QLatin1String_count, "QLatin1String_count")
	symaddr(&addr_QLatin1String_count2, "QLatin1String_count2")
	symaddr(&addr_QLatin1String_count3, "QLatin1String_count3")
	symaddr(&addr_QLatin1String_toShort, "QLatin1String_toShort")
	symaddr(&addr_QLatin1String_toUShort, "QLatin1String_toUShort")
	symaddr(&addr_QLatin1String_toInt, "QLatin1String_toInt")
	symaddr(&addr_QLatin1String_toUInt, "QLatin1String_toUInt")
	symaddr(&addr_QLatin1String_toLong, "QLatin1String_toLong")
	symaddr(&addr_QLatin1String_toULong, "QLatin1String_toULong")
	symaddr(&addr_QLatin1String_toLongLong, "QLatin1String_toLongLong")
	symaddr(&addr_QLatin1String_toULongLong, "QLatin1String_toULongLong")
	reg(&pg_QLatin1String_toFloat, "QLatin1String_toFloat")
	reg(&pg_QLatin1String_toDouble, "QLatin1String_toDouble")
	symaddr(&addr_QLatin1String_begin, "QLatin1String_begin")
	symaddr(&addr_QLatin1String_cbegin, "QLatin1String_cbegin")
	symaddr(&addr_QLatin1String_end, "QLatin1String_end")
	symaddr(&addr_QLatin1String_cend, "QLatin1String_cend")
	symaddr(&addr_QLatin1String_max_size, "QLatin1String_max_size")
	symaddr(&addr_QLatin1String_maxSize, "QLatin1String_maxSize")
	symaddr(&addr_QLatin1String_mid, "QLatin1String_mid")
	symaddr(&addr_QLatin1String_left, "QLatin1String_left")
	symaddr(&addr_QLatin1String_right, "QLatin1String_right")
	symaddr(&addr_QLatin1String_sliced, "QLatin1String_sliced")
	symaddr(&addr_QLatin1String_sliced2, "QLatin1String_sliced2")
	symaddr(&addr_QLatin1String_first2, "QLatin1String_first2")
	symaddr(&addr_QLatin1String_last2, "QLatin1String_last2")
	symaddr(&addr_QLatin1String_chopped, "QLatin1String_chopped")
	symaddr(&addr_QLatin1String_chop, "QLatin1String_chop")
	symaddr(&addr_QLatin1String_truncate, "QLatin1String_truncate")
	symaddr(&addr_QLatin1String_trimmed, "QLatin1String_trimmed")
	symaddr(&addr_QLatin1String_OperatorAssign, "QLatin1String_OperatorAssign")
	symaddr(&addr_QLatin1String_compare5, "QLatin1String_compare5")
	symaddr(&addr_QLatin1String_compare6, "QLatin1String_compare6")
	symaddr(&addr_QLatin1String_startsWith5, "QLatin1String_startsWith5")
	symaddr(&addr_QLatin1String_startsWith6, "QLatin1String_startsWith6")
	symaddr(&addr_QLatin1String_endsWith5, "QLatin1String_endsWith5")
	symaddr(&addr_QLatin1String_endsWith6, "QLatin1String_endsWith6")
	symaddr(&addr_QLatin1String_indexOf5, "QLatin1String_indexOf5")
	symaddr(&addr_QLatin1String_indexOf6, "QLatin1String_indexOf6")
	symaddr(&addr_QLatin1String_indexOf7, "QLatin1String_indexOf7")
	symaddr(&addr_QLatin1String_indexOf8, "QLatin1String_indexOf8")
	symaddr(&addr_QLatin1String_indexOf9, "QLatin1String_indexOf9")
	symaddr(&addr_QLatin1String_contains4, "QLatin1String_contains4")
	symaddr(&addr_QLatin1String_contains5, "QLatin1String_contains5")
	symaddr(&addr_QLatin1String_contains6, "QLatin1String_contains6")
	symaddr(&addr_QLatin1String_lastIndexOf9, "QLatin1String_lastIndexOf9")
	symaddr(&addr_QLatin1String_lastIndexOf10, "QLatin1String_lastIndexOf10")
	symaddr(&addr_QLatin1String_lastIndexOf11, "QLatin1String_lastIndexOf11")
	symaddr(&addr_QLatin1String_lastIndexOf12, "QLatin1String_lastIndexOf12")
	symaddr(&addr_QLatin1String_count4, "QLatin1String_count4")
	symaddr(&addr_QLatin1String_count5, "QLatin1String_count5")
	symaddr(&addr_QLatin1String_count6, "QLatin1String_count6")
	symaddr(&addr_QLatin1String_toShort2, "QLatin1String_toShort2")
	symaddr(&addr_QLatin1String_toShort3, "QLatin1String_toShort3")
	symaddr(&addr_QLatin1String_toUShort2, "QLatin1String_toUShort2")
	symaddr(&addr_QLatin1String_toUShort3, "QLatin1String_toUShort3")
	symaddr(&addr_QLatin1String_toInt2, "QLatin1String_toInt2")
	symaddr(&addr_QLatin1String_toInt3, "QLatin1String_toInt3")
	symaddr(&addr_QLatin1String_toUInt2, "QLatin1String_toUInt2")
	symaddr(&addr_QLatin1String_toUInt3, "QLatin1String_toUInt3")
	symaddr(&addr_QLatin1String_toLong2, "QLatin1String_toLong2")
	symaddr(&addr_QLatin1String_toLong3, "QLatin1String_toLong3")
	symaddr(&addr_QLatin1String_toULong2, "QLatin1String_toULong2")
	symaddr(&addr_QLatin1String_toULong3, "QLatin1String_toULong3")
	symaddr(&addr_QLatin1String_toLongLong2, "QLatin1String_toLongLong2")
	symaddr(&addr_QLatin1String_toLongLong3, "QLatin1String_toLongLong3")
	symaddr(&addr_QLatin1String_toULongLong2, "QLatin1String_toULongLong2")
	symaddr(&addr_QLatin1String_toULongLong3, "QLatin1String_toULongLong3")
	reg(&pg_QLatin1String_toFloat2, "QLatin1String_toFloat2")
	reg(&pg_QLatin1String_toDouble2, "QLatin1String_toDouble2")
	symaddr(&addr_QLatin1String_mid2, "QLatin1String_mid2")
	reg(&pg_QLatin1String_delete, "QLatin1String_delete")
}

type QLatin1String struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLatin1String) UnsafeGoID() int64     { return this.vid }
func (this *QLatin1String) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLatin1String(h uintptr) *QLatin1String {
	if h == 0 {
		return nil
	}
	return &QLatin1String{h: h}
}

func newownedQLatin1String(h uintptr) *QLatin1String {
	o := bindQLatin1String(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLatin1String).Delete)
	}
	return o
}

func (this *QLatin1String) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLatin1String) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLatin1String(h unsafe.Pointer) *QLatin1String { return bindQLatin1String(uintptr(h)) }
func UnsafeNewQLatin1StringOwned(h unsafe.Pointer) *QLatin1String { return newownedQLatin1String(uintptr(h)) }

func (this *QLatin1String) UnsafeRelease() { this.h = 0 }

func NewQLatin1String() *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new)
	__o := newownedQLatin1String(__r)
	return __o
}

func NewQLatin1String2(s string) *QLatin1String {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new2, uintptr(unsafe.Pointer(&s_b[0])))
	__o := newownedQLatin1String(__r)
	runtime.KeepAlive(s_b)
	return __o
}

func NewQLatin1String3(f string, l string) *QLatin1String {
	f_b := append([]byte(f), 0)
	l_b := append([]byte(l), 0)
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new3, uintptr(unsafe.Pointer(&f_b[0])), uintptr(unsafe.Pointer(&l_b[0])))
	__o := newownedQLatin1String(__r)
	runtime.KeepAlive(f_b)
	runtime.KeepAlive(l_b)
	return __o
}

func NewQLatin1String4(s string, sz int64) *QLatin1String {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new4, uintptr(unsafe.Pointer(&s_b[0])), uintptr(sz))
	__o := newownedQLatin1String(__r)
	runtime.KeepAlive(s_b)
	return __o
}

func NewQLatin1String5(s *QByteArray) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new5, uintptr(s.UnsafePointer()))
	__o := newownedQLatin1String(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQLatin1String6(s *QByteArrayView) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new6, uintptr(s.UnsafePointer()))
	__o := newownedQLatin1String(__r)
	runtime.KeepAlive(s)
	return __o
}

func NewQLatin1String7(arg1 *QLatin1String) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_new7, uintptr(arg1.UnsafePointer()))
	__o := newownedQLatin1String(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QLatin1String) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) ToUtf8() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_toUtf8, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QLatin1String) Latin1() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_latin1, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) Data() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_data, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) ConstData() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_constData, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) ConstBegin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_constBegin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) ConstEnd() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_constEnd, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) First() *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_first, this.h)
	runtime.KeepAlive(this)
	return newownedQLatin1Char(__r)
}

func (this *QLatin1String) Last() *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_last, this.h)
	runtime.KeepAlive(this)
	return newownedQLatin1Char(__r)
}

func (this *QLatin1String) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLatin1String) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLatin1String) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLatin1String) At(i int64) *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQLatin1Char(__r)
}

func (this *QLatin1String) OperatorSubscript(i int64) *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQLatin1Char(__r)
}

func (this *QLatin1String) Front() *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_front, this.h)
	runtime.KeepAlive(this)
	return newownedQLatin1Char(__r)
}

func (this *QLatin1String) Back() *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_back, this.h)
	runtime.KeepAlive(this)
	return newownedQLatin1Char(__r)
}

func (this *QLatin1String) Compare(other *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_compare, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QLatin1String) CompareWithOther(other *QLatin1String) int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_compare2, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QLatin1String) CompareWithQChar(c *QChar) int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_compare3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return int(__r)
}

func (this *QLatin1String) Compare2(c *QChar, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_compare4, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return int(__r)
}

func (this *QLatin1String) StartsWith(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_startsWith, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) StartsWithWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_startsWith2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) StartsWithWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_startsWith3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QLatin1String) StartsWith2(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_startsWith4, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QLatin1String) EndsWith(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_endsWith, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) EndsWithWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_endsWith2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) EndsWithWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_endsWith3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QLatin1String) EndsWith2(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_endsWith4, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QLatin1String) IndexOf(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) IndexOfWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf2, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) IndexOfWithQChar(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf3, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) IndexOf2(c *QChar, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf4, this.h, uintptr(c.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) Contains(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_contains, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) ContainsWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_contains2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) ContainsWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_contains3, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QLatin1String) LastIndexOf(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOf2(s *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf2, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOfWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf3, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOf3(s *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf4, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOfWithQChar(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf5, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) LastIndexOf4(c *QChar, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf6, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) LastIndexOf5(c *QChar, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf7, this.h, uintptr(c.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) LastIndexOf6(c *QChar, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf8, this.h, uintptr(c.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) Count(str *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_count, this.h, uintptr(str.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QLatin1String) CountWithStr(str *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_count2, this.h, uintptr(str.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QLatin1String) CountWithCh(ch *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_count3, this.h, uintptr(ch.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
	return __r
}

func (this *QLatin1String) ToShort() int16 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) ToUShort() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toUShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QLatin1String) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLatin1String) ToUInt() uint {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toUInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QLatin1String) ToLong() int32 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toLong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) ToULong() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toULong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QLatin1String) ToLongLong() int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toLongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) ToULongLong() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toULongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QLatin1String) ToFloat() float32 {
	__r := pg_QLatin1String_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) ToDouble() float64 {
	__r := pg_QLatin1String_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1String) Begin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_begin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) Cbegin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_cbegin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) End() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_end, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) Cend() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLatin1String_cend, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLatin1String) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func QLatin1String_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QLatin1String) Mid(pos int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_mid, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Left(n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_left, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Right(n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_right, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Sliced(pos int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_sliced, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Sliced2(pos int64, n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_sliced2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) FirstWith(n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_first2, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) LastWith(n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_last2, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Chopped(n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_chopped, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Chop(n int64) {
	purego.SyscallN(addr_QLatin1String_chop, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QLatin1String) Truncate(n int64) {
	purego.SyscallN(addr_QLatin1String_truncate, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QLatin1String) Trimmed() *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_trimmed, this.h)
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) OperatorAssign(arg1 *QLatin1String) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQLatin1String(__r)
}

func (this *QLatin1String) Compare3(other *QStringView, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_compare5, this.h, uintptr(other.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QLatin1String) Compare4(other *QLatin1String, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_compare6, this.h, uintptr(other.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QLatin1String) StartsWith3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_startsWith5, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) StartsWith4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_startsWith6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) EndsWith3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_endsWith5, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) EndsWith4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_endsWith6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) IndexOf3(s *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf5, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) IndexOf4(s *QStringView, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf6, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) IndexOf5(s *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf7, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) IndexOf6(s *QLatin1String, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf8, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) IndexOf7(c *QChar, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_indexOf9, this.h, uintptr(c.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QLatin1String) Contains2(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_contains4, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) Contains3(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_contains5, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QLatin1String) Contains4(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_contains6, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QLatin1String) LastIndexOf7(s *QStringView, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf9, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOf8(s *QStringView, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf10, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOf9(s *QLatin1String, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf11, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) LastIndexOf10(s *QLatin1String, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_lastIndexOf12, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLatin1String) Count2(str *QStringView, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_count4, this.h, uintptr(str.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QLatin1String) Count3(str *QLatin1String, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_count5, this.h, uintptr(str.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QLatin1String) Count4(ch *QChar, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_count6, this.h, uintptr(ch.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
	return __r
}

func (this *QLatin1String) ToShortWithOk() (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToShort2(base int) (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToUShortWithOk() (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toUShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QLatin1String) ToUShort2(base int) (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toUShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QLatin1String) ToIntWithOk() (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QLatin1String) ToInt2(base int) (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QLatin1String) ToUIntWithOk() (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toUInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QLatin1String) ToUInt2(base int) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toUInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QLatin1String) ToLongWithOk() (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToLong2(base int) (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToULongWithOk() (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toULong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QLatin1String) ToULong2(base int) (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toULong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QLatin1String) ToLongLongWithOk() (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toLongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToLongLong2(base int) (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toLongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToULongLongWithOk() (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toULongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QLatin1String) ToULongLong2(base int) (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLatin1String_toULongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QLatin1String) ToFloatWithOk() (float32, bool) {
	var ok_out int32
	__r := pg_QLatin1String_toFloat2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) ToDoubleWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QLatin1String_toDouble2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLatin1String) Mid2(pos int64, n int64) *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1String_mid2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1String) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLatin1String_delete(this.h)
	}
	this.UnsafeRelease()
}
