package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QByteArray_new uintptr
	addr_QByteArray_new2 uintptr
	addr_QByteArray_new3 uintptr
	addr_QByteArray_new4 uintptr
	addr_QByteArray_new5 uintptr
	addr_QByteArray_new6 uintptr
	addr_QByteArray_new7 uintptr
	addr_QByteArray_OperatorAssign uintptr
	addr_QByteArray_OperatorAssign2 uintptr
	addr_QByteArray_swap uintptr
	addr_QByteArray_isEmpty uintptr
	addr_QByteArray_resize uintptr
	addr_QByteArray_resize2 uintptr
	addr_QByteArray_resizeForOverwrite uintptr
	addr_QByteArray_fill uintptr
	addr_QByteArray_capacity uintptr
	addr_QByteArray_reserve uintptr
	addr_QByteArray_squeeze uintptr
	addr_QByteArray_data uintptr
	addr_QByteArray_constData uintptr
	addr_QByteArray_detach uintptr
	addr_QByteArray_isDetached uintptr
	addr_QByteArray_isSharedWith uintptr
	addr_QByteArray_clear uintptr
	addr_QByteArray_at uintptr
	addr_QByteArray_OperatorSubscript uintptr
	addr_QByteArray_front uintptr
	addr_QByteArray_back uintptr
	addr_QByteArray_indexOf uintptr
	addr_QByteArray_indexOf2 uintptr
	addr_QByteArray_lastIndexOf uintptr
	addr_QByteArray_lastIndexOf2 uintptr
	addr_QByteArray_lastIndexOf3 uintptr
	addr_QByteArray_contains uintptr
	addr_QByteArray_contains2 uintptr
	addr_QByteArray_count uintptr
	addr_QByteArray_count2 uintptr
	addr_QByteArray_compare uintptr
	addr_QByteArray_left uintptr
	addr_QByteArray_left2 uintptr
	addr_QByteArray_right uintptr
	addr_QByteArray_right2 uintptr
	addr_QByteArray_mid uintptr
	addr_QByteArray_mid2 uintptr
	addr_QByteArray_first uintptr
	addr_QByteArray_last uintptr
	addr_QByteArray_sliced uintptr
	addr_QByteArray_sliced2 uintptr
	addr_QByteArray_chopped uintptr
	addr_QByteArray_first2 uintptr
	addr_QByteArray_last2 uintptr
	addr_QByteArray_sliced3 uintptr
	addr_QByteArray_sliced4 uintptr
	addr_QByteArray_chopped2 uintptr
	addr_QByteArray_startsWith uintptr
	addr_QByteArray_startsWith2 uintptr
	addr_QByteArray_endsWith uintptr
	addr_QByteArray_endsWith2 uintptr
	addr_QByteArray_isUpper uintptr
	addr_QByteArray_isLower uintptr
	addr_QByteArray_isValidUtf8 uintptr
	addr_QByteArray_truncate uintptr
	addr_QByteArray_chop uintptr
	addr_QByteArray_slice uintptr
	addr_QByteArray_slice2 uintptr
	addr_QByteArray_toLower uintptr
	addr_QByteArray_toLower2 uintptr
	addr_QByteArray_toUpper uintptr
	addr_QByteArray_toUpper2 uintptr
	addr_QByteArray_trimmed uintptr
	addr_QByteArray_trimmed2 uintptr
	addr_QByteArray_simplified uintptr
	addr_QByteArray_simplified2 uintptr
	addr_QByteArray_leftJustified uintptr
	addr_QByteArray_rightJustified uintptr
	addr_QByteArray_prepend uintptr
	addr_QByteArray_prepend2 uintptr
	addr_QByteArray_prepend3 uintptr
	addr_QByteArray_prepend4 uintptr
	addr_QByteArray_prepend5 uintptr
	addr_QByteArray_prepend6 uintptr
	addr_QByteArray_append uintptr
	addr_QByteArray_append2 uintptr
	addr_QByteArray_append3 uintptr
	addr_QByteArray_append4 uintptr
	addr_QByteArray_append5 uintptr
	addr_QByteArray_append6 uintptr
	addr_QByteArray_assign uintptr
	addr_QByteArray_assign2 uintptr
	addr_QByteArray_insert uintptr
	addr_QByteArray_insert2 uintptr
	addr_QByteArray_insert3 uintptr
	addr_QByteArray_insert4 uintptr
	addr_QByteArray_insert5 uintptr
	addr_QByteArray_insert6 uintptr
	addr_QByteArray_remove uintptr
	addr_QByteArray_removeAt uintptr
	addr_QByteArray_removeFirst uintptr
	addr_QByteArray_removeLast uintptr
	addr_QByteArray_replace uintptr
	addr_QByteArray_replace2 uintptr
	addr_QByteArray_replace3 uintptr
	addr_QByteArray_replace4 uintptr
	addr_QByteArray_replace5 uintptr
	addr_QByteArray_replace6 uintptr
	addr_QByteArray_OperatorPlusAssign uintptr
	addr_QByteArray_OperatorPlusAssign2 uintptr
	addr_QByteArray_OperatorPlusAssign3 uintptr
	addr_QByteArray_OperatorPlusAssign4 uintptr
	addr_QByteArray_split uintptr
	addr_QByteArray_repeated uintptr
	addr_QByteArray_toShort uintptr
	addr_QByteArray_toUShort uintptr
	addr_QByteArray_toInt uintptr
	addr_QByteArray_toUInt uintptr
	addr_QByteArray_toLong uintptr
	addr_QByteArray_toULong uintptr
	addr_QByteArray_toLongLong uintptr
	addr_QByteArray_toULongLong uintptr
	pg_QByteArray_toFloat func(uintptr) float32
	pg_QByteArray_toDouble func(uintptr) float64
	addr_QByteArray_toBase64 uintptr
	addr_QByteArray_toHex uintptr
	addr_QByteArray_toPercentEncoding uintptr
	addr_QByteArray_percentDecoded uintptr
	addr_QByteArray_percentDecoded2 uintptr
	addr_QByteArray_setNum uintptr
	addr_QByteArray_setNum2 uintptr
	addr_QByteArray_setNum3 uintptr
	addr_QByteArray_setNum4 uintptr
	addr_QByteArray_setNum5 uintptr
	addr_QByteArray_setNum6 uintptr
	addr_QByteArray_setNum7 uintptr
	addr_QByteArray_setNum8 uintptr
	pg_QByteArray_setNum9 func(uintptr, float32) uintptr
	pg_QByteArray_setNum10 func(uintptr, float64) uintptr
	addr_QByteArray_setRawData uintptr
	addr_QByteArray_number uintptr
	addr_QByteArray_number2 uintptr
	addr_QByteArray_number3 uintptr
	addr_QByteArray_number4 uintptr
	addr_QByteArray_number5 uintptr
	addr_QByteArray_number6 uintptr
	pg_QByteArray_number7 func(float64) uintptr
	addr_QByteArray_fromRawData uintptr
	addr_QByteArray_fromBase64Encoding uintptr
	addr_QByteArray_fromBase64 uintptr
	addr_QByteArray_fromHex uintptr
	addr_QByteArray_fromPercentEncoding uintptr
	addr_QByteArray_begin uintptr
	addr_QByteArray_cbegin uintptr
	addr_QByteArray_constBegin uintptr
	addr_QByteArray_end uintptr
	addr_QByteArray_cend uintptr
	addr_QByteArray_constEnd uintptr
	addr_QByteArray_push_back uintptr
	addr_QByteArray_push_back2 uintptr
	addr_QByteArray_push_back3 uintptr
	addr_QByteArray_push_back4 uintptr
	addr_QByteArray_push_front uintptr
	addr_QByteArray_push_front2 uintptr
	addr_QByteArray_push_front3 uintptr
	addr_QByteArray_push_front4 uintptr
	addr_QByteArray_shrink_to_fit uintptr
	addr_QByteArray_max_size uintptr
	addr_QByteArray_maxSize uintptr
	addr_QByteArray_size uintptr
	addr_QByteArray_count3 uintptr
	addr_QByteArray_length uintptr
	addr_QByteArray_isNull uintptr
	addr_QByteArray_nullTerminated uintptr
	addr_QByteArray_nullTerminated2 uintptr
	addr_QByteArray_nullTerminate uintptr
	addr_QByteArray_fill2 uintptr
	addr_QByteArray_indexOf3 uintptr
	addr_QByteArray_indexOf4 uintptr
	addr_QByteArray_lastIndexOf4 uintptr
	addr_QByteArray_compare2 uintptr
	addr_QByteArray_mid3 uintptr
	addr_QByteArray_mid4 uintptr
	addr_QByteArray_leftJustified2 uintptr
	addr_QByteArray_leftJustified3 uintptr
	addr_QByteArray_rightJustified2 uintptr
	addr_QByteArray_rightJustified3 uintptr
	addr_QByteArray_toShort2 uintptr
	addr_QByteArray_toShort3 uintptr
	addr_QByteArray_toUShort2 uintptr
	addr_QByteArray_toUShort3 uintptr
	addr_QByteArray_toInt2 uintptr
	addr_QByteArray_toInt3 uintptr
	addr_QByteArray_toUInt2 uintptr
	addr_QByteArray_toUInt3 uintptr
	addr_QByteArray_toLong2 uintptr
	addr_QByteArray_toLong3 uintptr
	addr_QByteArray_toULong2 uintptr
	addr_QByteArray_toULong3 uintptr
	addr_QByteArray_toLongLong2 uintptr
	addr_QByteArray_toLongLong3 uintptr
	addr_QByteArray_toULongLong2 uintptr
	addr_QByteArray_toULongLong3 uintptr
	pg_QByteArray_toFloat2 func(uintptr, uintptr) float32
	pg_QByteArray_toDouble2 func(uintptr, uintptr) float64
	addr_QByteArray_toBase642 uintptr
	addr_QByteArray_toHex2 uintptr
	addr_QByteArray_toPercentEncoding2 uintptr
	addr_QByteArray_toPercentEncoding3 uintptr
	addr_QByteArray_toPercentEncoding4 uintptr
	addr_QByteArray_percentDecoded3 uintptr
	addr_QByteArray_percentDecoded4 uintptr
	addr_QByteArray_setNum11 uintptr
	addr_QByteArray_setNum12 uintptr
	addr_QByteArray_setNum13 uintptr
	addr_QByteArray_setNum14 uintptr
	addr_QByteArray_setNum15 uintptr
	addr_QByteArray_setNum16 uintptr
	addr_QByteArray_setNum17 uintptr
	addr_QByteArray_setNum18 uintptr
	pg_QByteArray_setNum19 func(uintptr, float32, int8) uintptr
	pg_QByteArray_setNum20 func(uintptr, float32, int8, int32) uintptr
	pg_QByteArray_setNum21 func(uintptr, float64, int8) uintptr
	pg_QByteArray_setNum22 func(uintptr, float64, int8, int32) uintptr
	addr_QByteArray_number8 uintptr
	addr_QByteArray_number9 uintptr
	addr_QByteArray_number10 uintptr
	addr_QByteArray_number11 uintptr
	addr_QByteArray_number12 uintptr
	addr_QByteArray_number13 uintptr
	pg_QByteArray_number14 func(float64, int8) uintptr
	pg_QByteArray_number15 func(float64, int8, int32) uintptr
	addr_QByteArray_fromBase64Encoding2 uintptr
	addr_QByteArray_fromBase642 uintptr
	addr_QByteArray_fromPercentEncoding2 uintptr
	addr_QByteArray_OperatorEqual uintptr
	addr_QByteArray_OperatorNotEqual uintptr
	addr_QByteArray_OperatorLesser uintptr
	addr_QByteArray_OperatorGreater uintptr
	addr_QByteArray_OperatorLesserOrEqual uintptr
	addr_QByteArray_OperatorGreaterOrEqual uintptr
	addr_QByteArray_OperatorEqual2 uintptr
	addr_QByteArray_OperatorNotEqual2 uintptr
	addr_QByteArray_OperatorLesser2 uintptr
	addr_QByteArray_OperatorGreater2 uintptr
	addr_QByteArray_OperatorLesserOrEqual2 uintptr
	addr_QByteArray_OperatorGreaterOrEqual2 uintptr
	addr_QByteArray_OperatorPlus uintptr
	addr_QByteArray_OperatorPlus2 uintptr
	addr_QByteArray_OperatorPlus3 uintptr
	addr_QByteArray_OperatorEqual3 uintptr
	addr_QByteArray_OperatorNotEqual3 uintptr
	addr_QByteArray_OperatorLesser3 uintptr
	addr_QByteArray_OperatorGreater3 uintptr
	addr_QByteArray_OperatorLesserOrEqual3 uintptr
	addr_QByteArray_OperatorGreaterOrEqual3 uintptr
	addr_QByteArray_OperatorEqual4 uintptr
	addr_QByteArray_OperatorNotEqual4 uintptr
	addr_QByteArray_OperatorLesser4 uintptr
	addr_QByteArray_OperatorGreater4 uintptr
	addr_QByteArray_OperatorLesserOrEqual4 uintptr
	addr_QByteArray_OperatorGreaterOrEqual4 uintptr
	addr_QByteArray_OperatorPlus4 uintptr
	pg_QByteArray_delete func(uintptr)
)

func init() {
	symaddr(&addr_QByteArray_new, "QByteArray_new")
	symaddr(&addr_QByteArray_new2, "QByteArray_new2")
	symaddr(&addr_QByteArray_new3, "QByteArray_new3")
	symaddr(&addr_QByteArray_new4, "QByteArray_new4")
	symaddr(&addr_QByteArray_new5, "QByteArray_new5")
	symaddr(&addr_QByteArray_new6, "QByteArray_new6")
	symaddr(&addr_QByteArray_new7, "QByteArray_new7")
	symaddr(&addr_QByteArray_OperatorAssign, "QByteArray_OperatorAssign")
	symaddr(&addr_QByteArray_OperatorAssign2, "QByteArray_OperatorAssign2")
	symaddr(&addr_QByteArray_swap, "QByteArray_swap")
	symaddr(&addr_QByteArray_isEmpty, "QByteArray_isEmpty")
	symaddr(&addr_QByteArray_resize, "QByteArray_resize")
	symaddr(&addr_QByteArray_resize2, "QByteArray_resize2")
	symaddr(&addr_QByteArray_resizeForOverwrite, "QByteArray_resizeForOverwrite")
	symaddr(&addr_QByteArray_fill, "QByteArray_fill")
	symaddr(&addr_QByteArray_capacity, "QByteArray_capacity")
	symaddr(&addr_QByteArray_reserve, "QByteArray_reserve")
	symaddr(&addr_QByteArray_squeeze, "QByteArray_squeeze")
	symaddr(&addr_QByteArray_data, "QByteArray_data")
	symaddr(&addr_QByteArray_constData, "QByteArray_constData")
	symaddr(&addr_QByteArray_detach, "QByteArray_detach")
	symaddr(&addr_QByteArray_isDetached, "QByteArray_isDetached")
	symaddr(&addr_QByteArray_isSharedWith, "QByteArray_isSharedWith")
	symaddr(&addr_QByteArray_clear, "QByteArray_clear")
	symaddr(&addr_QByteArray_at, "QByteArray_at")
	symaddr(&addr_QByteArray_OperatorSubscript, "QByteArray_OperatorSubscript")
	symaddr(&addr_QByteArray_front, "QByteArray_front")
	symaddr(&addr_QByteArray_back, "QByteArray_back")
	symaddr(&addr_QByteArray_indexOf, "QByteArray_indexOf")
	symaddr(&addr_QByteArray_indexOf2, "QByteArray_indexOf2")
	symaddr(&addr_QByteArray_lastIndexOf, "QByteArray_lastIndexOf")
	symaddr(&addr_QByteArray_lastIndexOf2, "QByteArray_lastIndexOf2")
	symaddr(&addr_QByteArray_lastIndexOf3, "QByteArray_lastIndexOf3")
	symaddr(&addr_QByteArray_contains, "QByteArray_contains")
	symaddr(&addr_QByteArray_contains2, "QByteArray_contains2")
	symaddr(&addr_QByteArray_count, "QByteArray_count")
	symaddr(&addr_QByteArray_count2, "QByteArray_count2")
	symaddr(&addr_QByteArray_compare, "QByteArray_compare")
	symaddr(&addr_QByteArray_left, "QByteArray_left")
	symaddr(&addr_QByteArray_left2, "QByteArray_left2")
	symaddr(&addr_QByteArray_right, "QByteArray_right")
	symaddr(&addr_QByteArray_right2, "QByteArray_right2")
	symaddr(&addr_QByteArray_mid, "QByteArray_mid")
	symaddr(&addr_QByteArray_mid2, "QByteArray_mid2")
	symaddr(&addr_QByteArray_first, "QByteArray_first")
	symaddr(&addr_QByteArray_last, "QByteArray_last")
	symaddr(&addr_QByteArray_sliced, "QByteArray_sliced")
	symaddr(&addr_QByteArray_sliced2, "QByteArray_sliced2")
	symaddr(&addr_QByteArray_chopped, "QByteArray_chopped")
	symaddr(&addr_QByteArray_first2, "QByteArray_first2")
	symaddr(&addr_QByteArray_last2, "QByteArray_last2")
	symaddr(&addr_QByteArray_sliced3, "QByteArray_sliced3")
	symaddr(&addr_QByteArray_sliced4, "QByteArray_sliced4")
	symaddr(&addr_QByteArray_chopped2, "QByteArray_chopped2")
	symaddr(&addr_QByteArray_startsWith, "QByteArray_startsWith")
	symaddr(&addr_QByteArray_startsWith2, "QByteArray_startsWith2")
	symaddr(&addr_QByteArray_endsWith, "QByteArray_endsWith")
	symaddr(&addr_QByteArray_endsWith2, "QByteArray_endsWith2")
	symaddr(&addr_QByteArray_isUpper, "QByteArray_isUpper")
	symaddr(&addr_QByteArray_isLower, "QByteArray_isLower")
	symaddr(&addr_QByteArray_isValidUtf8, "QByteArray_isValidUtf8")
	symaddr(&addr_QByteArray_truncate, "QByteArray_truncate")
	symaddr(&addr_QByteArray_chop, "QByteArray_chop")
	symaddr(&addr_QByteArray_slice, "QByteArray_slice")
	symaddr(&addr_QByteArray_slice2, "QByteArray_slice2")
	symaddr(&addr_QByteArray_toLower, "QByteArray_toLower")
	symaddr(&addr_QByteArray_toLower2, "QByteArray_toLower2")
	symaddr(&addr_QByteArray_toUpper, "QByteArray_toUpper")
	symaddr(&addr_QByteArray_toUpper2, "QByteArray_toUpper2")
	symaddr(&addr_QByteArray_trimmed, "QByteArray_trimmed")
	symaddr(&addr_QByteArray_trimmed2, "QByteArray_trimmed2")
	symaddr(&addr_QByteArray_simplified, "QByteArray_simplified")
	symaddr(&addr_QByteArray_simplified2, "QByteArray_simplified2")
	symaddr(&addr_QByteArray_leftJustified, "QByteArray_leftJustified")
	symaddr(&addr_QByteArray_rightJustified, "QByteArray_rightJustified")
	symaddr(&addr_QByteArray_prepend, "QByteArray_prepend")
	symaddr(&addr_QByteArray_prepend2, "QByteArray_prepend2")
	symaddr(&addr_QByteArray_prepend3, "QByteArray_prepend3")
	symaddr(&addr_QByteArray_prepend4, "QByteArray_prepend4")
	symaddr(&addr_QByteArray_prepend5, "QByteArray_prepend5")
	symaddr(&addr_QByteArray_prepend6, "QByteArray_prepend6")
	symaddr(&addr_QByteArray_append, "QByteArray_append")
	symaddr(&addr_QByteArray_append2, "QByteArray_append2")
	symaddr(&addr_QByteArray_append3, "QByteArray_append3")
	symaddr(&addr_QByteArray_append4, "QByteArray_append4")
	symaddr(&addr_QByteArray_append5, "QByteArray_append5")
	symaddr(&addr_QByteArray_append6, "QByteArray_append6")
	symaddr(&addr_QByteArray_assign, "QByteArray_assign")
	symaddr(&addr_QByteArray_assign2, "QByteArray_assign2")
	symaddr(&addr_QByteArray_insert, "QByteArray_insert")
	symaddr(&addr_QByteArray_insert2, "QByteArray_insert2")
	symaddr(&addr_QByteArray_insert3, "QByteArray_insert3")
	symaddr(&addr_QByteArray_insert4, "QByteArray_insert4")
	symaddr(&addr_QByteArray_insert5, "QByteArray_insert5")
	symaddr(&addr_QByteArray_insert6, "QByteArray_insert6")
	symaddr(&addr_QByteArray_remove, "QByteArray_remove")
	symaddr(&addr_QByteArray_removeAt, "QByteArray_removeAt")
	symaddr(&addr_QByteArray_removeFirst, "QByteArray_removeFirst")
	symaddr(&addr_QByteArray_removeLast, "QByteArray_removeLast")
	symaddr(&addr_QByteArray_replace, "QByteArray_replace")
	symaddr(&addr_QByteArray_replace2, "QByteArray_replace2")
	symaddr(&addr_QByteArray_replace3, "QByteArray_replace3")
	symaddr(&addr_QByteArray_replace4, "QByteArray_replace4")
	symaddr(&addr_QByteArray_replace5, "QByteArray_replace5")
	symaddr(&addr_QByteArray_replace6, "QByteArray_replace6")
	symaddr(&addr_QByteArray_OperatorPlusAssign, "QByteArray_OperatorPlusAssign")
	symaddr(&addr_QByteArray_OperatorPlusAssign2, "QByteArray_OperatorPlusAssign2")
	symaddr(&addr_QByteArray_OperatorPlusAssign3, "QByteArray_OperatorPlusAssign3")
	symaddr(&addr_QByteArray_OperatorPlusAssign4, "QByteArray_OperatorPlusAssign4")
	symaddr(&addr_QByteArray_split, "QByteArray_split")
	symaddr(&addr_QByteArray_repeated, "QByteArray_repeated")
	symaddr(&addr_QByteArray_toShort, "QByteArray_toShort")
	symaddr(&addr_QByteArray_toUShort, "QByteArray_toUShort")
	symaddr(&addr_QByteArray_toInt, "QByteArray_toInt")
	symaddr(&addr_QByteArray_toUInt, "QByteArray_toUInt")
	symaddr(&addr_QByteArray_toLong, "QByteArray_toLong")
	symaddr(&addr_QByteArray_toULong, "QByteArray_toULong")
	symaddr(&addr_QByteArray_toLongLong, "QByteArray_toLongLong")
	symaddr(&addr_QByteArray_toULongLong, "QByteArray_toULongLong")
	reg(&pg_QByteArray_toFloat, "QByteArray_toFloat")
	reg(&pg_QByteArray_toDouble, "QByteArray_toDouble")
	symaddr(&addr_QByteArray_toBase64, "QByteArray_toBase64")
	symaddr(&addr_QByteArray_toHex, "QByteArray_toHex")
	symaddr(&addr_QByteArray_toPercentEncoding, "QByteArray_toPercentEncoding")
	symaddr(&addr_QByteArray_percentDecoded, "QByteArray_percentDecoded")
	symaddr(&addr_QByteArray_percentDecoded2, "QByteArray_percentDecoded2")
	symaddr(&addr_QByteArray_setNum, "QByteArray_setNum")
	symaddr(&addr_QByteArray_setNum2, "QByteArray_setNum2")
	symaddr(&addr_QByteArray_setNum3, "QByteArray_setNum3")
	symaddr(&addr_QByteArray_setNum4, "QByteArray_setNum4")
	symaddr(&addr_QByteArray_setNum5, "QByteArray_setNum5")
	symaddr(&addr_QByteArray_setNum6, "QByteArray_setNum6")
	symaddr(&addr_QByteArray_setNum7, "QByteArray_setNum7")
	symaddr(&addr_QByteArray_setNum8, "QByteArray_setNum8")
	reg(&pg_QByteArray_setNum9, "QByteArray_setNum9")
	reg(&pg_QByteArray_setNum10, "QByteArray_setNum10")
	symaddr(&addr_QByteArray_setRawData, "QByteArray_setRawData")
	symaddr(&addr_QByteArray_number, "QByteArray_number")
	symaddr(&addr_QByteArray_number2, "QByteArray_number2")
	symaddr(&addr_QByteArray_number3, "QByteArray_number3")
	symaddr(&addr_QByteArray_number4, "QByteArray_number4")
	symaddr(&addr_QByteArray_number5, "QByteArray_number5")
	symaddr(&addr_QByteArray_number6, "QByteArray_number6")
	reg(&pg_QByteArray_number7, "QByteArray_number7")
	symaddr(&addr_QByteArray_fromRawData, "QByteArray_fromRawData")
	symaddr(&addr_QByteArray_fromBase64Encoding, "QByteArray_fromBase64Encoding")
	symaddr(&addr_QByteArray_fromBase64, "QByteArray_fromBase64")
	symaddr(&addr_QByteArray_fromHex, "QByteArray_fromHex")
	symaddr(&addr_QByteArray_fromPercentEncoding, "QByteArray_fromPercentEncoding")
	symaddr(&addr_QByteArray_begin, "QByteArray_begin")
	symaddr(&addr_QByteArray_cbegin, "QByteArray_cbegin")
	symaddr(&addr_QByteArray_constBegin, "QByteArray_constBegin")
	symaddr(&addr_QByteArray_end, "QByteArray_end")
	symaddr(&addr_QByteArray_cend, "QByteArray_cend")
	symaddr(&addr_QByteArray_constEnd, "QByteArray_constEnd")
	symaddr(&addr_QByteArray_push_back, "QByteArray_push_back")
	symaddr(&addr_QByteArray_push_back2, "QByteArray_push_back2")
	symaddr(&addr_QByteArray_push_back3, "QByteArray_push_back3")
	symaddr(&addr_QByteArray_push_back4, "QByteArray_push_back4")
	symaddr(&addr_QByteArray_push_front, "QByteArray_push_front")
	symaddr(&addr_QByteArray_push_front2, "QByteArray_push_front2")
	symaddr(&addr_QByteArray_push_front3, "QByteArray_push_front3")
	symaddr(&addr_QByteArray_push_front4, "QByteArray_push_front4")
	symaddr(&addr_QByteArray_shrink_to_fit, "QByteArray_shrink_to_fit")
	symaddr(&addr_QByteArray_max_size, "QByteArray_max_size")
	symaddr(&addr_QByteArray_maxSize, "QByteArray_maxSize")
	symaddr(&addr_QByteArray_size, "QByteArray_size")
	symaddr(&addr_QByteArray_count3, "QByteArray_count3")
	symaddr(&addr_QByteArray_length, "QByteArray_length")
	symaddr(&addr_QByteArray_isNull, "QByteArray_isNull")
	symaddr(&addr_QByteArray_nullTerminated, "QByteArray_nullTerminated")
	symaddr(&addr_QByteArray_nullTerminated2, "QByteArray_nullTerminated2")
	symaddr(&addr_QByteArray_nullTerminate, "QByteArray_nullTerminate")
	symaddr(&addr_QByteArray_fill2, "QByteArray_fill2")
	symaddr(&addr_QByteArray_indexOf3, "QByteArray_indexOf3")
	symaddr(&addr_QByteArray_indexOf4, "QByteArray_indexOf4")
	symaddr(&addr_QByteArray_lastIndexOf4, "QByteArray_lastIndexOf4")
	symaddr(&addr_QByteArray_compare2, "QByteArray_compare2")
	symaddr(&addr_QByteArray_mid3, "QByteArray_mid3")
	symaddr(&addr_QByteArray_mid4, "QByteArray_mid4")
	symaddr(&addr_QByteArray_leftJustified2, "QByteArray_leftJustified2")
	symaddr(&addr_QByteArray_leftJustified3, "QByteArray_leftJustified3")
	symaddr(&addr_QByteArray_rightJustified2, "QByteArray_rightJustified2")
	symaddr(&addr_QByteArray_rightJustified3, "QByteArray_rightJustified3")
	symaddr(&addr_QByteArray_toShort2, "QByteArray_toShort2")
	symaddr(&addr_QByteArray_toShort3, "QByteArray_toShort3")
	symaddr(&addr_QByteArray_toUShort2, "QByteArray_toUShort2")
	symaddr(&addr_QByteArray_toUShort3, "QByteArray_toUShort3")
	symaddr(&addr_QByteArray_toInt2, "QByteArray_toInt2")
	symaddr(&addr_QByteArray_toInt3, "QByteArray_toInt3")
	symaddr(&addr_QByteArray_toUInt2, "QByteArray_toUInt2")
	symaddr(&addr_QByteArray_toUInt3, "QByteArray_toUInt3")
	symaddr(&addr_QByteArray_toLong2, "QByteArray_toLong2")
	symaddr(&addr_QByteArray_toLong3, "QByteArray_toLong3")
	symaddr(&addr_QByteArray_toULong2, "QByteArray_toULong2")
	symaddr(&addr_QByteArray_toULong3, "QByteArray_toULong3")
	symaddr(&addr_QByteArray_toLongLong2, "QByteArray_toLongLong2")
	symaddr(&addr_QByteArray_toLongLong3, "QByteArray_toLongLong3")
	symaddr(&addr_QByteArray_toULongLong2, "QByteArray_toULongLong2")
	symaddr(&addr_QByteArray_toULongLong3, "QByteArray_toULongLong3")
	reg(&pg_QByteArray_toFloat2, "QByteArray_toFloat2")
	reg(&pg_QByteArray_toDouble2, "QByteArray_toDouble2")
	symaddr(&addr_QByteArray_toBase642, "QByteArray_toBase642")
	symaddr(&addr_QByteArray_toHex2, "QByteArray_toHex2")
	symaddr(&addr_QByteArray_toPercentEncoding2, "QByteArray_toPercentEncoding2")
	symaddr(&addr_QByteArray_toPercentEncoding3, "QByteArray_toPercentEncoding3")
	symaddr(&addr_QByteArray_toPercentEncoding4, "QByteArray_toPercentEncoding4")
	symaddr(&addr_QByteArray_percentDecoded3, "QByteArray_percentDecoded3")
	symaddr(&addr_QByteArray_percentDecoded4, "QByteArray_percentDecoded4")
	symaddr(&addr_QByteArray_setNum11, "QByteArray_setNum11")
	symaddr(&addr_QByteArray_setNum12, "QByteArray_setNum12")
	symaddr(&addr_QByteArray_setNum13, "QByteArray_setNum13")
	symaddr(&addr_QByteArray_setNum14, "QByteArray_setNum14")
	symaddr(&addr_QByteArray_setNum15, "QByteArray_setNum15")
	symaddr(&addr_QByteArray_setNum16, "QByteArray_setNum16")
	symaddr(&addr_QByteArray_setNum17, "QByteArray_setNum17")
	symaddr(&addr_QByteArray_setNum18, "QByteArray_setNum18")
	reg(&pg_QByteArray_setNum19, "QByteArray_setNum19")
	reg(&pg_QByteArray_setNum20, "QByteArray_setNum20")
	reg(&pg_QByteArray_setNum21, "QByteArray_setNum21")
	reg(&pg_QByteArray_setNum22, "QByteArray_setNum22")
	symaddr(&addr_QByteArray_number8, "QByteArray_number8")
	symaddr(&addr_QByteArray_number9, "QByteArray_number9")
	symaddr(&addr_QByteArray_number10, "QByteArray_number10")
	symaddr(&addr_QByteArray_number11, "QByteArray_number11")
	symaddr(&addr_QByteArray_number12, "QByteArray_number12")
	symaddr(&addr_QByteArray_number13, "QByteArray_number13")
	reg(&pg_QByteArray_number14, "QByteArray_number14")
	reg(&pg_QByteArray_number15, "QByteArray_number15")
	symaddr(&addr_QByteArray_fromBase64Encoding2, "QByteArray_fromBase64Encoding2")
	symaddr(&addr_QByteArray_fromBase642, "QByteArray_fromBase642")
	symaddr(&addr_QByteArray_fromPercentEncoding2, "QByteArray_fromPercentEncoding2")
	symaddr(&addr_QByteArray_OperatorEqual, "QByteArray_OperatorEqual")
	symaddr(&addr_QByteArray_OperatorNotEqual, "QByteArray_OperatorNotEqual")
	symaddr(&addr_QByteArray_OperatorLesser, "QByteArray_OperatorLesser")
	symaddr(&addr_QByteArray_OperatorGreater, "QByteArray_OperatorGreater")
	symaddr(&addr_QByteArray_OperatorLesserOrEqual, "QByteArray_OperatorLesserOrEqual")
	symaddr(&addr_QByteArray_OperatorGreaterOrEqual, "QByteArray_OperatorGreaterOrEqual")
	symaddr(&addr_QByteArray_OperatorEqual2, "QByteArray_OperatorEqual2")
	symaddr(&addr_QByteArray_OperatorNotEqual2, "QByteArray_OperatorNotEqual2")
	symaddr(&addr_QByteArray_OperatorLesser2, "QByteArray_OperatorLesser2")
	symaddr(&addr_QByteArray_OperatorGreater2, "QByteArray_OperatorGreater2")
	symaddr(&addr_QByteArray_OperatorLesserOrEqual2, "QByteArray_OperatorLesserOrEqual2")
	symaddr(&addr_QByteArray_OperatorGreaterOrEqual2, "QByteArray_OperatorGreaterOrEqual2")
	symaddr(&addr_QByteArray_OperatorPlus, "QByteArray_OperatorPlus")
	symaddr(&addr_QByteArray_OperatorPlus2, "QByteArray_OperatorPlus2")
	symaddr(&addr_QByteArray_OperatorPlus3, "QByteArray_OperatorPlus3")
	symaddr(&addr_QByteArray_OperatorEqual3, "QByteArray_OperatorEqual3")
	symaddr(&addr_QByteArray_OperatorNotEqual3, "QByteArray_OperatorNotEqual3")
	symaddr(&addr_QByteArray_OperatorLesser3, "QByteArray_OperatorLesser3")
	symaddr(&addr_QByteArray_OperatorGreater3, "QByteArray_OperatorGreater3")
	symaddr(&addr_QByteArray_OperatorLesserOrEqual3, "QByteArray_OperatorLesserOrEqual3")
	symaddr(&addr_QByteArray_OperatorGreaterOrEqual3, "QByteArray_OperatorGreaterOrEqual3")
	symaddr(&addr_QByteArray_OperatorEqual4, "QByteArray_OperatorEqual4")
	symaddr(&addr_QByteArray_OperatorNotEqual4, "QByteArray_OperatorNotEqual4")
	symaddr(&addr_QByteArray_OperatorLesser4, "QByteArray_OperatorLesser4")
	symaddr(&addr_QByteArray_OperatorGreater4, "QByteArray_OperatorGreater4")
	symaddr(&addr_QByteArray_OperatorLesserOrEqual4, "QByteArray_OperatorLesserOrEqual4")
	symaddr(&addr_QByteArray_OperatorGreaterOrEqual4, "QByteArray_OperatorGreaterOrEqual4")
	symaddr(&addr_QByteArray_OperatorPlus4, "QByteArray_OperatorPlus4")
	reg(&pg_QByteArray_delete, "QByteArray_delete")
}

type QByteArray struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QByteArray) UnsafeGoID() int64     { return this.vid }
func (this *QByteArray) UnsafeSetGoID(id int64) { this.vid = id }

func bindQByteArray(h uintptr) *QByteArray {
	if h == 0 {
		return nil
	}
	return &QByteArray{h: h}
}

func newownedQByteArray(h uintptr) *QByteArray {
	o := bindQByteArray(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QByteArray).Delete)
	}
	return o
}

func (this *QByteArray) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QByteArray) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQByteArray(h unsafe.Pointer) *QByteArray { return bindQByteArray(uintptr(h)) }
func UnsafeNewQByteArrayOwned(h unsafe.Pointer) *QByteArray { return newownedQByteArray(uintptr(h)) }

func (this *QByteArray) UnsafeRelease() { this.h = 0 }

func NewQByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_new)
	__o := newownedQByteArray(__r)
	return __o
}

func NewQByteArray2(arg1 string) *QByteArray {
	arg1_b := append([]byte(arg1), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_new2, uintptr(unsafe.Pointer(&arg1_b[0])))
	__o := newownedQByteArray(__r)
	runtime.KeepAlive(arg1_b)
	return __o
}

func NewQByteArray3(size int64, c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_new3, uintptr(size), uintptr(c))
	__o := newownedQByteArray(__r)
	return __o
}

func NewQByteArray4(size int64, arg2 Initialization) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_new4, uintptr(size), uintptr(arg2))
	__o := newownedQByteArray(__r)
	return __o
}

func NewQByteArray5(v *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_new5, uintptr(v.UnsafePointer()))
	__o := newownedQByteArray(__r)
	runtime.KeepAlive(v)
	return __o
}

func NewQByteArray6(arg1 *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_new6, uintptr(arg1.UnsafePointer()))
	__o := newownedQByteArray(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQByteArray7(arg1 string, size int64) *QByteArray {
	arg1_b := append([]byte(arg1), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_new7, uintptr(unsafe.Pointer(&arg1_b[0])), uintptr(size))
	__o := newownedQByteArray(__r)
	runtime.KeepAlive(arg1_b)
	return __o
}

func (this *QByteArray) OperatorAssign(arg1 *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQByteArray(__r)
}

func (this *QByteArray) OperatorAssignWithStr(str string) *QByteArray {
	str_b := append([]byte(str), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorAssign2, this.h, uintptr(unsafe.Pointer(&str_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Swap(other *QByteArray) {
	purego.SyscallN(addr_QByteArray_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QByteArray) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) Resize(size int64) {
	purego.SyscallN(addr_QByteArray_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Resize2(size int64, c int8) {
	purego.SyscallN(addr_QByteArray_resize2, this.h, uintptr(size), uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QByteArray) ResizeForOverwrite(size int64) {
	purego.SyscallN(addr_QByteArray_resizeForOverwrite, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Fill(c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fill, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) Reserve(size int64) {
	purego.SyscallN(addr_QByteArray_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Squeeze() {
	purego.SyscallN(addr_QByteArray_squeeze, this.h)
	runtime.KeepAlive(this)
}

func (this *QByteArray) Data() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_data, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) ConstData() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_constData, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) Detach() {
	purego.SyscallN(addr_QByteArray_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QByteArray) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) IsSharedWith(other *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isSharedWith, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QByteArray) Clear() {
	purego.SyscallN(addr_QByteArray_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QByteArray) At(i int64) int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_at, this.h, uintptr(i))
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) OperatorSubscript(i int64) int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorSubscript, this.h, uintptr(i))
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) Front() int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_front, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) Back() int8 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_back, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) IndexOf(c int8) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_indexOf, this.h, uintptr(c))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) IndexOfWithBv(bv *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_indexOf2, this.h, uintptr(bv.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r
}

func (this *QByteArray) LastIndexOf(c int8) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_lastIndexOf, this.h, uintptr(c))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) LastIndexOfWithBv(bv *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_lastIndexOf2, this.h, uintptr(bv.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r
}

func (this *QByteArray) LastIndexOf2(bv *QByteArrayView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_lastIndexOf3, this.h, uintptr(bv.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r
}

func (this *QByteArray) Contains(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_contains, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) ContainsWithBv(bv *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_contains2, this.h, uintptr(bv.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r != 0
}

func (this *QByteArray) Count(c int8) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_count, this.h, uintptr(c))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) CountWithBv(bv *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_count2, this.h, uintptr(bv.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r
}

func (this *QByteArray) Compare(a *QByteArrayView) int {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_compare, this.h, uintptr(a.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return int(__r)
}

func (this *QByteArray) Left(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_left, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) LeftWith(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_left2, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Right(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_right, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) RightWith(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_right2, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Mid(index int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_mid, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) MidWithIndex(index int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_mid2, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) First(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_first, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Last(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_last, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Sliced(pos int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_sliced, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Sliced2(pos int64, n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_sliced2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Chopped(len_ int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_chopped, this.h, uintptr(len_))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) FirstWith(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_first2, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) LastWith(n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_last2, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) SlicedWithPos(pos int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_sliced3, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Sliced3(pos int64, n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_sliced4, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ChoppedWithLen(len_ int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_chopped2, this.h, uintptr(len_))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) StartsWith(bv *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_startsWith, this.h, uintptr(bv.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r != 0
}

func (this *QByteArray) StartsWithWithChar(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_startsWith2, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) EndsWith(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_endsWith, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) EndsWithWithBv(bv *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_endsWith2, this.h, uintptr(bv.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r != 0
}

func (this *QByteArray) IsUpper() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isUpper, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) IsLower() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isLower, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) IsValidUtf8() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isValidUtf8, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) Truncate(pos int64) {
	purego.SyscallN(addr_QByteArray_truncate, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Chop(n int64) {
	purego.SyscallN(addr_QByteArray_chop, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Slice(pos int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_slice, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Slice2(pos int64, n int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_slice2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) ToLower() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toLower, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToLower2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toLower2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToUpper() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toUpper, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToUpper2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toUpper2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Trimmed() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_trimmed, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Trimmed2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_trimmed2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Simplified() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_simplified, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Simplified2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_simplified2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) LeftJustified(width int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_leftJustified, this.h, uintptr(width))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) RightJustified(width int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_rightJustified, this.h, uintptr(width))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Prepend(c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_prepend, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Prepend2(count int64, c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_prepend2, this.h, uintptr(count), uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) PrependWithChar(s string) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_prepend3, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Prepend3(s string, len_ int64) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_prepend4, this.h, uintptr(unsafe.Pointer(&s_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) PrependWithQByteArray(a *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_prepend5, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQByteArray(__r)
}

func (this *QByteArray) PrependWithQByteArrayView(a *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_prepend6, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQByteArray(__r)
}

func (this *QByteArray) Append(c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_append, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Append2(count int64, c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_append2, this.h, uintptr(count), uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) AppendWithChar(s string) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_append3, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Append3(s string, len_ int64) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_append4, this.h, uintptr(unsafe.Pointer(&s_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) AppendWithQByteArray(a *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_append5, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQByteArray(__r)
}

func (this *QByteArray) AppendWithQByteArrayView(a *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_append6, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQByteArray(__r)
}

func (this *QByteArray) Assign(v *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_assign, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQByteArray(__r)
}

func (this *QByteArray) Assign2(n int64, c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_assign2, this.h, uintptr(n), uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Insert(i int64, data *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_insert, this.h, uintptr(i), uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return bindQByteArray(__r)
}

func (this *QByteArray) Insert2(i int64, s string) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_insert2, this.h, uintptr(i), uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Insert3(i int64, data *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_insert3, this.h, uintptr(i), uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return bindQByteArray(__r)
}

func (this *QByteArray) Insert4(i int64, count int64, c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_insert4, this.h, uintptr(i), uintptr(count), uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Insert5(i int64, c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_insert5, this.h, uintptr(i), uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Insert6(i int64, s string, len_ int64) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_insert6, this.h, uintptr(i), uintptr(unsafe.Pointer(&s_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Remove(index int64, len_ int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_remove, this.h, uintptr(index), uintptr(len_))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) RemoveAt(pos int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_removeAt, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) RemoveFirst() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_removeFirst, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) RemoveLast() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_removeLast, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Replace(index int64, len_ int64, s string, alen int64) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_replace, this.h, uintptr(index), uintptr(len_), uintptr(unsafe.Pointer(&s_b[0])), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Replace2(index int64, len_ int64, s *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_replace2, this.h, uintptr(index), uintptr(len_), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQByteArray(__r)
}

func (this *QByteArray) Replace3(before int8, after *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_replace3, this.h, uintptr(before), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(after)
	return bindQByteArray(__r)
}

func (this *QByteArray) Replace4(before string, bsize int64, after string, asize int64) *QByteArray {
	before_b := append([]byte(before), 0)
	after_b := append([]byte(after), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_replace4, this.h, uintptr(unsafe.Pointer(&before_b[0])), uintptr(bsize), uintptr(unsafe.Pointer(&after_b[0])), uintptr(asize))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before_b)
	runtime.KeepAlive(after_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) Replace5(before *QByteArrayView, after *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_replace5, this.h, uintptr(before.UnsafePointer()), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQByteArray(__r)
}

func (this *QByteArray) Replace6(before int8, after int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_replace6, this.h, uintptr(before), uintptr(after))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) OperatorPlusAssign(c int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlusAssign, this.h, uintptr(c))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) OperatorPlusAssignWithChar(s string) *QByteArray {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlusAssign2, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQByteArray(__r)
}

func (this *QByteArray) OperatorPlusAssignWithQByteArray(a *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlusAssign3, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQByteArray(__r)
}

func (this *QByteArray) OperatorPlusAssignWithQByteArrayView(a *QByteArrayView) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlusAssign4, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQByteArray(__r)
}

func (this *QByteArray) Split(sep int8) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_split, this.h, uintptr(sep), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QByteArray) Repeated(times int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_repeated, this.h, uintptr(times))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToShort() int16 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) ToUShort() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toUShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QByteArray) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QByteArray) ToUInt() uint {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toUInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QByteArray) ToLong() int32 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toLong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) ToULong() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toULong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QByteArray) ToLongLong() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toLongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) ToULongLong() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toULongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QByteArray) ToFloat() float32 {
	__r := pg_QByteArray_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) ToDouble() float64 {
	__r := pg_QByteArray_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) ToBase64() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toBase64, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToHex() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toHex, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToPercentEncoding() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toPercentEncoding, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) PercentDecoded() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_percentDecoded, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) PercentDecoded2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_percentDecoded2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) SetNum(arg1 int16) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNumWith(arg1 uint16) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNumWithInt(arg1 int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum3, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum2(arg1 uint) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum4, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNumWithLong(arg1 int32) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum5, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum3(arg1 uint32) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum6, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum4(arg1 int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum7, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum5(arg1 uint64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum8, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNumWithFloat(arg1 float32) *QByteArray {
	__r := pg_QByteArray_setNum9(this.h, arg1)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNumWithDouble(arg1 float64) *QByteArray {
	__r := pg_QByteArray_setNum10(this.h, arg1)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetRawData(a string, n int64) *QByteArray {
	a_b := append([]byte(a), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_setRawData, this.h, uintptr(unsafe.Pointer(&a_b[0])), uintptr(n))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a_b)
	return bindQByteArray(__r)
}

func QByteArray_Number(arg1 int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number, uintptr(arg1))
	return newownedQByteArray(__r)
}

func QByteArray_NumberWith(arg1 uint) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number2, uintptr(arg1))
	return newownedQByteArray(__r)
}

func QByteArray_NumberWithLong(arg1 int32) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number3, uintptr(arg1))
	return newownedQByteArray(__r)
}

func QByteArray_Number2(arg1 uint32) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number4, uintptr(arg1))
	return newownedQByteArray(__r)
}

func QByteArray_Number3(arg1 int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number5, uintptr(arg1))
	return newownedQByteArray(__r)
}

func QByteArray_Number4(arg1 uint64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number6, uintptr(arg1))
	return newownedQByteArray(__r)
}

func QByteArray_NumberWithDouble(arg1 float64) *QByteArray {
	__r := pg_QByteArray_number7(arg1)
	return newownedQByteArray(__r)
}

func QByteArray_FromRawData(data string, size int64) *QByteArray {
	data_b := append([]byte(data), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromRawData, uintptr(unsafe.Pointer(&data_b[0])), uintptr(size))
	runtime.KeepAlive(data_b)
	return newownedQByteArray(__r)
}

func QByteArray_FromBase64Encoding(base64 *QByteArray) *QByteArray__FromBase64Result {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromBase64Encoding, uintptr(base64.UnsafePointer()))
	runtime.KeepAlive(base64)
	return newownedQByteArray__FromBase64Result(__r)
}

func QByteArray_FromBase64(base64 *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromBase64, uintptr(base64.UnsafePointer()))
	runtime.KeepAlive(base64)
	return newownedQByteArray(__r)
}

func QByteArray_FromHex(hexEncoded *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromHex, uintptr(hexEncoded.UnsafePointer()))
	runtime.KeepAlive(hexEncoded)
	return newownedQByteArray(__r)
}

func QByteArray_FromPercentEncoding(pctEncoded *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromPercentEncoding, uintptr(pctEncoded.UnsafePointer()))
	runtime.KeepAlive(pctEncoded)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Begin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_begin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) Cbegin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_cbegin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) ConstBegin() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_constBegin, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) End() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_end, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) Cend() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_cend, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) ConstEnd() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_constEnd, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QByteArray) Push_back(c int8) {
	purego.SyscallN(addr_QByteArray_push_back, this.h, uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Push_backWithChar(s string) {
	s_b := append([]byte(s), 0)
	purego.SyscallN(addr_QByteArray_push_back2, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
}

func (this *QByteArray) Push_backWithQByteArray(a *QByteArray) {
	purego.SyscallN(addr_QByteArray_push_back3, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
}

func (this *QByteArray) Push_backWithQByteArrayView(a *QByteArrayView) {
	purego.SyscallN(addr_QByteArray_push_back4, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
}

func (this *QByteArray) Push_front(c int8) {
	purego.SyscallN(addr_QByteArray_push_front, this.h, uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QByteArray) Push_frontWithChar(c string) {
	c_b := append([]byte(c), 0)
	purego.SyscallN(addr_QByteArray_push_front2, this.h, uintptr(unsafe.Pointer(&c_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c_b)
}

func (this *QByteArray) Push_frontWithQByteArray(a *QByteArray) {
	purego.SyscallN(addr_QByteArray_push_front3, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
}

func (this *QByteArray) Push_frontWithQByteArrayView(a *QByteArrayView) {
	purego.SyscallN(addr_QByteArray_push_front4, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
}

func (this *QByteArray) Shrink_to_fit() {
	purego.SyscallN(addr_QByteArray_shrink_to_fit, this.h)
	runtime.KeepAlive(this)
}

func (this *QByteArray) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func QByteArray_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QByteArray) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) Count2() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_count3, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray) NullTerminated() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_nullTerminated, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) NullTerminated2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_nullTerminated2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) NullTerminate() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_nullTerminate, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) Fill2(c int8, size int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fill2, this.h, uintptr(c), uintptr(size))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) IndexOf2(c int8, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_indexOf3, this.h, uintptr(c), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) IndexOf3(bv *QByteArrayView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_indexOf4, this.h, uintptr(bv.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(bv)
	return __r
}

func (this *QByteArray) LastIndexOf3(c int8, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_lastIndexOf4, this.h, uintptr(c), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QByteArray) Compare2(a *QByteArrayView, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_compare2, this.h, uintptr(a.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return int(__r)
}

func (this *QByteArray) Mid2(index int64, len_ int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_mid3, this.h, uintptr(index), uintptr(len_))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) Mid3(index int64, len_ int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_mid4, this.h, uintptr(index), uintptr(len_))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) LeftJustified2(width int64, fill int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_leftJustified2, this.h, uintptr(width), uintptr(fill))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) LeftJustified3(width int64, fill int8, truncate bool) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_leftJustified3, this.h, uintptr(width), uintptr(fill), uintptr(b2i(truncate)))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) RightJustified2(width int64, fill int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_rightJustified2, this.h, uintptr(width), uintptr(fill))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) RightJustified3(width int64, fill int8, truncate bool) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_rightJustified3, this.h, uintptr(width), uintptr(fill), uintptr(b2i(truncate)))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToShortWithOk() (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToShort2(base int) (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToUShortWithOk() (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toUShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QByteArray) ToUShort2(base int) (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toUShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QByteArray) ToIntWithOk() (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QByteArray) ToInt2(base int) (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QByteArray) ToUIntWithOk() (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toUInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QByteArray) ToUInt2(base int) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toUInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QByteArray) ToLongWithOk() (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToLong2(base int) (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToULongWithOk() (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toULong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QByteArray) ToULong2(base int) (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toULong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QByteArray) ToLongLongWithOk() (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toLongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToLongLong2(base int) (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toLongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToULongLongWithOk() (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toULongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QByteArray) ToULongLong2(base int) (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QByteArray_toULongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QByteArray) ToFloatWithOk() (float32, bool) {
	var ok_out int32
	__r := pg_QByteArray_toFloat2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToDoubleWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QByteArray_toDouble2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QByteArray) ToBase64WithOptions(options QByteArray__Base64Option) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toBase642, this.h, uintptr(options))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToHexWithSeparator(separator int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toHex2, this.h, uintptr(separator))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToPercentEncodingWithExclude(exclude *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toPercentEncoding2, this.h, uintptr(exclude.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exclude)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToPercentEncoding2(exclude *QByteArray, include *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toPercentEncoding3, this.h, uintptr(exclude.UnsafePointer()), uintptr(include.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exclude)
	runtime.KeepAlive(include)
	return newownedQByteArray(__r)
}

func (this *QByteArray) ToPercentEncoding3(exclude *QByteArray, include *QByteArray, percent int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_toPercentEncoding4, this.h, uintptr(exclude.UnsafePointer()), uintptr(include.UnsafePointer()), uintptr(percent))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exclude)
	runtime.KeepAlive(include)
	return newownedQByteArray(__r)
}

func (this *QByteArray) PercentDecodedWithPercent(percent int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_percentDecoded3, this.h, uintptr(percent))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) PercentDecoded3(percent int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_percentDecoded4, this.h, uintptr(percent))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) SetNum6(arg1 int16, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum11, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum7(arg1 uint16, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum12, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum8(arg1 int, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum13, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum9(arg1 uint, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum14, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum10(arg1 int32, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum15, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum11(arg1 uint32, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum16, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum12(arg1 int64, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum17, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum13(arg1 uint64, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_setNum18, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum14(arg1 float32, format int8) *QByteArray {
	__r := pg_QByteArray_setNum19(this.h, arg1, format)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum15(arg1 float32, format int8, precision int) *QByteArray {
	__r := pg_QByteArray_setNum20(this.h, arg1, format, int32(precision))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum16(arg1 float64, format int8) *QByteArray {
	__r := pg_QByteArray_setNum21(this.h, arg1, format)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray) SetNum17(arg1 float64, format int8, precision int) *QByteArray {
	__r := pg_QByteArray_setNum22(this.h, arg1, format, int32(precision))
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func QByteArray_Number5(arg1 int, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number8, uintptr(arg1), uintptr(base))
	return newownedQByteArray(__r)
}

func QByteArray_Number6(arg1 uint, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number9, uintptr(arg1), uintptr(base))
	return newownedQByteArray(__r)
}

func QByteArray_Number7(arg1 int32, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number10, uintptr(arg1), uintptr(base))
	return newownedQByteArray(__r)
}

func QByteArray_Number8(arg1 uint32, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number11, uintptr(arg1), uintptr(base))
	return newownedQByteArray(__r)
}

func QByteArray_Number9(arg1 int64, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number12, uintptr(arg1), uintptr(base))
	return newownedQByteArray(__r)
}

func QByteArray_Number10(arg1 uint64, base int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_number13, uintptr(arg1), uintptr(base))
	return newownedQByteArray(__r)
}

func QByteArray_Number11(arg1 float64, format int8) *QByteArray {
	__r := pg_QByteArray_number14(arg1, format)
	return newownedQByteArray(__r)
}

func QByteArray_Number12(arg1 float64, format int8, precision int) *QByteArray {
	__r := pg_QByteArray_number15(arg1, format, int32(precision))
	return newownedQByteArray(__r)
}

func QByteArray_FromBase64Encoding2(base64 *QByteArray, options QByteArray__Base64Option) *QByteArray__FromBase64Result {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromBase64Encoding2, uintptr(base64.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(base64)
	return newownedQByteArray__FromBase64Result(__r)
}

func QByteArray_FromBase642(base64 *QByteArray, options QByteArray__Base64Option) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromBase642, uintptr(base64.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(base64)
	return newownedQByteArray(__r)
}

func QByteArray_FromPercentEncoding2(pctEncoded *QByteArray, percent int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_fromPercentEncoding2, uintptr(pctEncoded.UnsafePointer()), uintptr(percent))
	runtime.KeepAlive(pctEncoded)
	return newownedQByteArray(__r)
}

func (this *QByteArray) OperatorEqual(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorNotEqual(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesser(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreater(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesserOrEqual(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreaterOrEqual(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorNotEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesserWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreaterWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesserOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreaterOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorPlus(a2 *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlus, this.h, uintptr(a2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a2)
	return newownedQByteArray(__r)
}

func (this *QByteArray) OperatorPlusWithA2(a2 string) *QByteArray {
	a2_b := append([]byte(a2), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlus2, this.h, uintptr(unsafe.Pointer(&a2_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a2_b)
	return newownedQByteArray(__r)
}

func (this *QByteArray) OperatorPlus2(a2 int8) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlus3, this.h, uintptr(a2))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray) OperatorEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorNotEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesser2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreater2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesserOrEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreaterOrEqual2(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorNotEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorNotEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesser3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesser4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreater3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreater4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorLesserOrEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorLesserOrEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorGreaterOrEqual3(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QByteArray_OperatorGreaterOrEqual4, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray) OperatorPlusWithQString(s string) string {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QByteArray_OperatorPlus4, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return pgStr(__r, __n)
}

func (this *QByteArray) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QByteArray_delete(this.h)
	}
	this.UnsafeRelease()
}
