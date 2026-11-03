package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QString_new uintptr
	addr_QString_new2 uintptr
	addr_QString_new3 uintptr
	addr_QString_new4 uintptr
	addr_QString_new5 uintptr
	addr_QString_new6 uintptr
	addr_QString_new7 uintptr
	addr_QString_new8 uintptr
	addr_QString_new9 uintptr
	addr_QString_new10 uintptr
	addr_QString_new11 uintptr
	addr_QString_OperatorAssign uintptr
	addr_QString_OperatorAssign2 uintptr
	addr_QString_OperatorAssign3 uintptr
	addr_QString_swap uintptr
	addr_QString_maxSize uintptr
	addr_QString_size uintptr
	addr_QString_count uintptr
	addr_QString_length uintptr
	addr_QString_isEmpty uintptr
	addr_QString_resize uintptr
	addr_QString_resize2 uintptr
	addr_QString_resizeForOverwrite uintptr
	addr_QString_fill uintptr
	addr_QString_truncate uintptr
	addr_QString_chop uintptr
	addr_QString_slice uintptr
	addr_QString_slice2 uintptr
	addr_QString_capacity uintptr
	addr_QString_reserve uintptr
	addr_QString_squeeze uintptr
	addr_QString_unicode uintptr
	addr_QString_data uintptr
	addr_QString_data2 uintptr
	addr_QString_constData uintptr
	addr_QString_detach uintptr
	addr_QString_isDetached uintptr
	addr_QString_isSharedWith uintptr
	addr_QString_clear uintptr
	addr_QString_at uintptr
	addr_QString_OperatorSubscript uintptr
	addr_QString_OperatorSubscript2 uintptr
	addr_QString_front uintptr
	addr_QString_front2 uintptr
	addr_QString_back uintptr
	addr_QString_back2 uintptr
	addr_QString_asprintf uintptr
	addr_QString_indexOf uintptr
	addr_QString_indexOf2 uintptr
	addr_QString_indexOf3 uintptr
	addr_QString_indexOf4 uintptr
	addr_QString_lastIndexOf uintptr
	addr_QString_lastIndexOf2 uintptr
	addr_QString_lastIndexOf3 uintptr
	addr_QString_lastIndexOf4 uintptr
	addr_QString_lastIndexOf5 uintptr
	addr_QString_lastIndexOf6 uintptr
	addr_QString_lastIndexOf7 uintptr
	addr_QString_lastIndexOf8 uintptr
	addr_QString_contains uintptr
	addr_QString_contains2 uintptr
	addr_QString_contains3 uintptr
	addr_QString_contains4 uintptr
	addr_QString_count2 uintptr
	addr_QString_count3 uintptr
	addr_QString_count4 uintptr
	addr_QString_indexOf5 uintptr
	addr_QString_lastIndexOf9 uintptr
	addr_QString_contains5 uintptr
	addr_QString_count5 uintptr
	addr_QString_section uintptr
	addr_QString_section2 uintptr
	addr_QString_section3 uintptr
	addr_QString_left uintptr
	addr_QString_left2 uintptr
	addr_QString_right uintptr
	addr_QString_right2 uintptr
	addr_QString_mid uintptr
	addr_QString_mid2 uintptr
	addr_QString_first uintptr
	addr_QString_last uintptr
	addr_QString_sliced uintptr
	addr_QString_sliced2 uintptr
	addr_QString_chopped uintptr
	addr_QString_first2 uintptr
	addr_QString_last2 uintptr
	addr_QString_sliced3 uintptr
	addr_QString_sliced4 uintptr
	addr_QString_chopped2 uintptr
	addr_QString_startsWith uintptr
	addr_QString_startsWith2 uintptr
	addr_QString_startsWith3 uintptr
	addr_QString_startsWith4 uintptr
	addr_QString_endsWith uintptr
	addr_QString_endsWith2 uintptr
	addr_QString_endsWith3 uintptr
	addr_QString_endsWith4 uintptr
	addr_QString_isUpper uintptr
	addr_QString_isLower uintptr
	addr_QString_leftJustified uintptr
	addr_QString_rightJustified uintptr
	addr_QString_toLower uintptr
	addr_QString_toLower2 uintptr
	addr_QString_toUpper uintptr
	addr_QString_toUpper2 uintptr
	addr_QString_toCaseFolded uintptr
	addr_QString_toCaseFolded2 uintptr
	addr_QString_trimmed uintptr
	addr_QString_trimmed2 uintptr
	addr_QString_simplified uintptr
	addr_QString_simplified2 uintptr
	addr_QString_toHtmlEscaped uintptr
	addr_QString_insert uintptr
	addr_QString_insert2 uintptr
	addr_QString_insert3 uintptr
	addr_QString_insert4 uintptr
	addr_QString_insert5 uintptr
	addr_QString_append uintptr
	addr_QString_append2 uintptr
	addr_QString_append3 uintptr
	addr_QString_append4 uintptr
	addr_QString_append5 uintptr
	addr_QString_prepend uintptr
	addr_QString_prepend2 uintptr
	addr_QString_prepend3 uintptr
	addr_QString_prepend4 uintptr
	addr_QString_prepend5 uintptr
	addr_QString_assign uintptr
	addr_QString_assign2 uintptr
	addr_QString_OperatorPlusAssign uintptr
	addr_QString_OperatorPlusAssign2 uintptr
	addr_QString_OperatorPlusAssign3 uintptr
	addr_QString_OperatorPlusAssign4 uintptr
	addr_QString_remove uintptr
	addr_QString_remove2 uintptr
	addr_QString_remove3 uintptr
	addr_QString_remove4 uintptr
	addr_QString_removeAt uintptr
	addr_QString_removeFirst uintptr
	addr_QString_removeLast uintptr
	addr_QString_replace uintptr
	addr_QString_replace2 uintptr
	addr_QString_replace3 uintptr
	addr_QString_replace4 uintptr
	addr_QString_replace5 uintptr
	addr_QString_replace6 uintptr
	addr_QString_replace7 uintptr
	addr_QString_replace8 uintptr
	addr_QString_replace9 uintptr
	addr_QString_replace10 uintptr
	addr_QString_replace11 uintptr
	addr_QString_replace12 uintptr
	addr_QString_remove5 uintptr
	addr_QString_split uintptr
	addr_QString_split2 uintptr
	addr_QString_split3 uintptr
	addr_QString_normalized uintptr
	addr_QString_repeated uintptr
	addr_QString_nullTerminated uintptr
	addr_QString_nullTerminated2 uintptr
	addr_QString_nullTerminate uintptr
	addr_QString_toLatin1 uintptr
	addr_QString_toLatin12 uintptr
	addr_QString_toUtf8 uintptr
	addr_QString_toUtf82 uintptr
	addr_QString_toLocal8Bit uintptr
	addr_QString_toLocal8Bit2 uintptr
	addr_QString_toUcs4 uintptr
	addr_QString_fromLatin1 uintptr
	addr_QString_fromLatin12 uintptr
	addr_QString_fromUtf8 uintptr
	addr_QString_fromUtf82 uintptr
	addr_QString_fromLocal8Bit uintptr
	addr_QString_fromLocal8Bit2 uintptr
	addr_QString_fromRawData uintptr
	addr_QString_setRawData uintptr
	addr_QString_setUnicode uintptr
	addr_QString_compare uintptr
	addr_QString_compare2 uintptr
	addr_QString_compare3 uintptr
	addr_QString_compare4 uintptr
	addr_QString_compare5 uintptr
	addr_QString_compare6 uintptr
	addr_QString_compare7 uintptr
	addr_QString_compare8 uintptr
	addr_QString_compare9 uintptr
	addr_QString_localeAwareCompare uintptr
	addr_QString_localeAwareCompare2 uintptr
	addr_QString_localeAwareCompare3 uintptr
	addr_QString_localeAwareCompare4 uintptr
	addr_QString_toShort uintptr
	addr_QString_toUShort uintptr
	addr_QString_toInt uintptr
	addr_QString_toUInt uintptr
	addr_QString_toLong uintptr
	addr_QString_toULong uintptr
	addr_QString_toLongLong uintptr
	addr_QString_toULongLong uintptr
	pg_QString_toFloat func(uintptr) float32
	pg_QString_toDouble func(uintptr) float64
	addr_QString_setNum uintptr
	addr_QString_setNum2 uintptr
	addr_QString_setNum3 uintptr
	addr_QString_setNum4 uintptr
	addr_QString_setNum5 uintptr
	addr_QString_setNum6 uintptr
	addr_QString_setNum7 uintptr
	addr_QString_setNum8 uintptr
	pg_QString_setNum9 func(uintptr, float32) uintptr
	pg_QString_setNum10 func(uintptr, float64) uintptr
	addr_QString_number uintptr
	addr_QString_number2 uintptr
	addr_QString_number3 uintptr
	addr_QString_number4 uintptr
	addr_QString_number5 uintptr
	addr_QString_number6 uintptr
	pg_QString_number7 func(float64, *int64) uintptr
	addr_QString_OperatorAssign4 uintptr
	addr_QString_OperatorAssign5 uintptr
	addr_QString_prepend6 uintptr
	addr_QString_prepend7 uintptr
	addr_QString_append6 uintptr
	addr_QString_append7 uintptr
	addr_QString_insert6 uintptr
	addr_QString_insert7 uintptr
	addr_QString_OperatorPlusAssign5 uintptr
	addr_QString_OperatorPlusAssign6 uintptr
	addr_QString_begin uintptr
	addr_QString_begin2 uintptr
	addr_QString_cbegin uintptr
	addr_QString_constBegin uintptr
	addr_QString_end uintptr
	addr_QString_end2 uintptr
	addr_QString_cend uintptr
	addr_QString_constEnd uintptr
	addr_QString_push_back uintptr
	addr_QString_push_back2 uintptr
	addr_QString_push_front uintptr
	addr_QString_push_front2 uintptr
	addr_QString_shrink_to_fit uintptr
	addr_QString_erase uintptr
	addr_QString_erase2 uintptr
	addr_QString_max_size uintptr
	addr_QString_isNull uintptr
	addr_QString_isRightToLeft uintptr
	addr_QString_isValidUtf16 uintptr
	addr_QString_fill2 uintptr
	addr_QString_indexOf6 uintptr
	addr_QString_indexOf7 uintptr
	addr_QString_indexOf8 uintptr
	addr_QString_indexOf9 uintptr
	addr_QString_indexOf10 uintptr
	addr_QString_indexOf11 uintptr
	addr_QString_indexOf12 uintptr
	addr_QString_indexOf13 uintptr
	addr_QString_lastIndexOf10 uintptr
	addr_QString_lastIndexOf11 uintptr
	addr_QString_lastIndexOf12 uintptr
	addr_QString_lastIndexOf13 uintptr
	addr_QString_lastIndexOf14 uintptr
	addr_QString_lastIndexOf15 uintptr
	addr_QString_lastIndexOf16 uintptr
	addr_QString_lastIndexOf17 uintptr
	addr_QString_contains6 uintptr
	addr_QString_contains7 uintptr
	addr_QString_contains8 uintptr
	addr_QString_contains9 uintptr
	addr_QString_count6 uintptr
	addr_QString_count7 uintptr
	addr_QString_count8 uintptr
	addr_QString_indexOf14 uintptr
	addr_QString_indexOf15 uintptr
	addr_QString_lastIndexOf18 uintptr
	addr_QString_contains10 uintptr
	addr_QString_section4 uintptr
	addr_QString_section5 uintptr
	addr_QString_section6 uintptr
	addr_QString_section7 uintptr
	addr_QString_section8 uintptr
	addr_QString_section9 uintptr
	addr_QString_mid3 uintptr
	addr_QString_mid4 uintptr
	addr_QString_startsWith5 uintptr
	addr_QString_startsWith6 uintptr
	addr_QString_startsWith7 uintptr
	addr_QString_startsWith8 uintptr
	addr_QString_endsWith5 uintptr
	addr_QString_endsWith6 uintptr
	addr_QString_endsWith7 uintptr
	addr_QString_endsWith8 uintptr
	addr_QString_leftJustified2 uintptr
	addr_QString_leftJustified3 uintptr
	addr_QString_rightJustified2 uintptr
	addr_QString_rightJustified3 uintptr
	addr_QString_remove6 uintptr
	addr_QString_remove7 uintptr
	addr_QString_remove8 uintptr
	addr_QString_replace13 uintptr
	addr_QString_replace14 uintptr
	addr_QString_replace15 uintptr
	addr_QString_replace16 uintptr
	addr_QString_replace17 uintptr
	addr_QString_replace18 uintptr
	addr_QString_replace19 uintptr
	addr_QString_replace20 uintptr
	addr_QString_split4 uintptr
	addr_QString_split5 uintptr
	addr_QString_split6 uintptr
	addr_QString_split7 uintptr
	addr_QString_split8 uintptr
	addr_QString_normalized2 uintptr
	addr_QString_compare10 uintptr
	addr_QString_compare11 uintptr
	addr_QString_compare12 uintptr
	addr_QString_compare13 uintptr
	addr_QString_compare14 uintptr
	addr_QString_compare15 uintptr
	addr_QString_compare16 uintptr
	addr_QString_compare17 uintptr
	addr_QString_compare18 uintptr
	addr_QString_toShort2 uintptr
	addr_QString_toShort3 uintptr
	addr_QString_toUShort2 uintptr
	addr_QString_toUShort3 uintptr
	addr_QString_toInt2 uintptr
	addr_QString_toInt3 uintptr
	addr_QString_toUInt2 uintptr
	addr_QString_toUInt3 uintptr
	addr_QString_toLong2 uintptr
	addr_QString_toLong3 uintptr
	addr_QString_toULong2 uintptr
	addr_QString_toULong3 uintptr
	addr_QString_toLongLong2 uintptr
	addr_QString_toLongLong3 uintptr
	addr_QString_toULongLong2 uintptr
	addr_QString_toULongLong3 uintptr
	pg_QString_toFloat2 func(uintptr, uintptr) float32
	pg_QString_toDouble2 func(uintptr, uintptr) float64
	addr_QString_setNum11 uintptr
	addr_QString_setNum12 uintptr
	addr_QString_setNum13 uintptr
	addr_QString_setNum14 uintptr
	addr_QString_setNum15 uintptr
	addr_QString_setNum16 uintptr
	addr_QString_setNum17 uintptr
	addr_QString_setNum18 uintptr
	pg_QString_setNum19 func(uintptr, float32, int8) uintptr
	pg_QString_setNum20 func(uintptr, float32, int8, int32) uintptr
	pg_QString_setNum21 func(uintptr, float64, int8) uintptr
	pg_QString_setNum22 func(uintptr, float64, int8, int32) uintptr
	addr_QString_number8 uintptr
	addr_QString_number9 uintptr
	addr_QString_number10 uintptr
	addr_QString_number11 uintptr
	addr_QString_number12 uintptr
	addr_QString_number13 uintptr
	pg_QString_number14 func(float64, int8, *int64) uintptr
	pg_QString_number15 func(float64, int8, int32, *int64) uintptr
	addr_QString_OperatorEqual uintptr
	addr_QString_OperatorNotEqual uintptr
	addr_QString_OperatorLesser uintptr
	addr_QString_OperatorGreater uintptr
	addr_QString_OperatorLesserOrEqual uintptr
	addr_QString_OperatorGreaterOrEqual uintptr
	addr_QString_OperatorEqual2 uintptr
	addr_QString_OperatorNotEqual2 uintptr
	addr_QString_OperatorLesser2 uintptr
	addr_QString_OperatorGreater2 uintptr
	addr_QString_OperatorLesserOrEqual2 uintptr
	addr_QString_OperatorGreaterOrEqual2 uintptr
	addr_QString_OperatorEqual3 uintptr
	addr_QString_OperatorNotEqual3 uintptr
	addr_QString_OperatorLesser3 uintptr
	addr_QString_OperatorGreater3 uintptr
	addr_QString_OperatorLesserOrEqual3 uintptr
	addr_QString_OperatorGreaterOrEqual3 uintptr
	addr_QString_OperatorEqual4 uintptr
	addr_QString_OperatorNotEqual4 uintptr
	addr_QString_OperatorLesser4 uintptr
	addr_QString_OperatorGreater4 uintptr
	addr_QString_OperatorLesserOrEqual4 uintptr
	addr_QString_OperatorGreaterOrEqual4 uintptr
	addr_QString_OperatorPlus uintptr
	addr_QString_OperatorPlus2 uintptr
	addr_QString_OperatorPlus3 uintptr
	addr_QString_OperatorPlus4 uintptr
	addr_QString_OperatorPlus5 uintptr
	pg_QString_delete func(uintptr)
)

func init() {
	symaddr(&addr_QString_new, "QString_new")
	symaddr(&addr_QString_new2, "QString_new2")
	symaddr(&addr_QString_new3, "QString_new3")
	symaddr(&addr_QString_new4, "QString_new4")
	symaddr(&addr_QString_new5, "QString_new5")
	symaddr(&addr_QString_new6, "QString_new6")
	symaddr(&addr_QString_new7, "QString_new7")
	symaddr(&addr_QString_new8, "QString_new8")
	symaddr(&addr_QString_new9, "QString_new9")
	symaddr(&addr_QString_new10, "QString_new10")
	symaddr(&addr_QString_new11, "QString_new11")
	symaddr(&addr_QString_OperatorAssign, "QString_OperatorAssign")
	symaddr(&addr_QString_OperatorAssign2, "QString_OperatorAssign2")
	symaddr(&addr_QString_OperatorAssign3, "QString_OperatorAssign3")
	symaddr(&addr_QString_swap, "QString_swap")
	symaddr(&addr_QString_maxSize, "QString_maxSize")
	symaddr(&addr_QString_size, "QString_size")
	symaddr(&addr_QString_count, "QString_count")
	symaddr(&addr_QString_length, "QString_length")
	symaddr(&addr_QString_isEmpty, "QString_isEmpty")
	symaddr(&addr_QString_resize, "QString_resize")
	symaddr(&addr_QString_resize2, "QString_resize2")
	symaddr(&addr_QString_resizeForOverwrite, "QString_resizeForOverwrite")
	symaddr(&addr_QString_fill, "QString_fill")
	symaddr(&addr_QString_truncate, "QString_truncate")
	symaddr(&addr_QString_chop, "QString_chop")
	symaddr(&addr_QString_slice, "QString_slice")
	symaddr(&addr_QString_slice2, "QString_slice2")
	symaddr(&addr_QString_capacity, "QString_capacity")
	symaddr(&addr_QString_reserve, "QString_reserve")
	symaddr(&addr_QString_squeeze, "QString_squeeze")
	symaddr(&addr_QString_unicode, "QString_unicode")
	symaddr(&addr_QString_data, "QString_data")
	symaddr(&addr_QString_data2, "QString_data2")
	symaddr(&addr_QString_constData, "QString_constData")
	symaddr(&addr_QString_detach, "QString_detach")
	symaddr(&addr_QString_isDetached, "QString_isDetached")
	symaddr(&addr_QString_isSharedWith, "QString_isSharedWith")
	symaddr(&addr_QString_clear, "QString_clear")
	symaddr(&addr_QString_at, "QString_at")
	symaddr(&addr_QString_OperatorSubscript, "QString_OperatorSubscript")
	symaddr(&addr_QString_OperatorSubscript2, "QString_OperatorSubscript2")
	symaddr(&addr_QString_front, "QString_front")
	symaddr(&addr_QString_front2, "QString_front2")
	symaddr(&addr_QString_back, "QString_back")
	symaddr(&addr_QString_back2, "QString_back2")
	symaddr(&addr_QString_asprintf, "QString_asprintf")
	symaddr(&addr_QString_indexOf, "QString_indexOf")
	symaddr(&addr_QString_indexOf2, "QString_indexOf2")
	symaddr(&addr_QString_indexOf3, "QString_indexOf3")
	symaddr(&addr_QString_indexOf4, "QString_indexOf4")
	symaddr(&addr_QString_lastIndexOf, "QString_lastIndexOf")
	symaddr(&addr_QString_lastIndexOf2, "QString_lastIndexOf2")
	symaddr(&addr_QString_lastIndexOf3, "QString_lastIndexOf3")
	symaddr(&addr_QString_lastIndexOf4, "QString_lastIndexOf4")
	symaddr(&addr_QString_lastIndexOf5, "QString_lastIndexOf5")
	symaddr(&addr_QString_lastIndexOf6, "QString_lastIndexOf6")
	symaddr(&addr_QString_lastIndexOf7, "QString_lastIndexOf7")
	symaddr(&addr_QString_lastIndexOf8, "QString_lastIndexOf8")
	symaddr(&addr_QString_contains, "QString_contains")
	symaddr(&addr_QString_contains2, "QString_contains2")
	symaddr(&addr_QString_contains3, "QString_contains3")
	symaddr(&addr_QString_contains4, "QString_contains4")
	symaddr(&addr_QString_count2, "QString_count2")
	symaddr(&addr_QString_count3, "QString_count3")
	symaddr(&addr_QString_count4, "QString_count4")
	symaddr(&addr_QString_indexOf5, "QString_indexOf5")
	symaddr(&addr_QString_lastIndexOf9, "QString_lastIndexOf9")
	symaddr(&addr_QString_contains5, "QString_contains5")
	symaddr(&addr_QString_count5, "QString_count5")
	symaddr(&addr_QString_section, "QString_section")
	symaddr(&addr_QString_section2, "QString_section2")
	symaddr(&addr_QString_section3, "QString_section3")
	symaddr(&addr_QString_left, "QString_left")
	symaddr(&addr_QString_left2, "QString_left2")
	symaddr(&addr_QString_right, "QString_right")
	symaddr(&addr_QString_right2, "QString_right2")
	symaddr(&addr_QString_mid, "QString_mid")
	symaddr(&addr_QString_mid2, "QString_mid2")
	symaddr(&addr_QString_first, "QString_first")
	symaddr(&addr_QString_last, "QString_last")
	symaddr(&addr_QString_sliced, "QString_sliced")
	symaddr(&addr_QString_sliced2, "QString_sliced2")
	symaddr(&addr_QString_chopped, "QString_chopped")
	symaddr(&addr_QString_first2, "QString_first2")
	symaddr(&addr_QString_last2, "QString_last2")
	symaddr(&addr_QString_sliced3, "QString_sliced3")
	symaddr(&addr_QString_sliced4, "QString_sliced4")
	symaddr(&addr_QString_chopped2, "QString_chopped2")
	symaddr(&addr_QString_startsWith, "QString_startsWith")
	symaddr(&addr_QString_startsWith2, "QString_startsWith2")
	symaddr(&addr_QString_startsWith3, "QString_startsWith3")
	symaddr(&addr_QString_startsWith4, "QString_startsWith4")
	symaddr(&addr_QString_endsWith, "QString_endsWith")
	symaddr(&addr_QString_endsWith2, "QString_endsWith2")
	symaddr(&addr_QString_endsWith3, "QString_endsWith3")
	symaddr(&addr_QString_endsWith4, "QString_endsWith4")
	symaddr(&addr_QString_isUpper, "QString_isUpper")
	symaddr(&addr_QString_isLower, "QString_isLower")
	symaddr(&addr_QString_leftJustified, "QString_leftJustified")
	symaddr(&addr_QString_rightJustified, "QString_rightJustified")
	symaddr(&addr_QString_toLower, "QString_toLower")
	symaddr(&addr_QString_toLower2, "QString_toLower2")
	symaddr(&addr_QString_toUpper, "QString_toUpper")
	symaddr(&addr_QString_toUpper2, "QString_toUpper2")
	symaddr(&addr_QString_toCaseFolded, "QString_toCaseFolded")
	symaddr(&addr_QString_toCaseFolded2, "QString_toCaseFolded2")
	symaddr(&addr_QString_trimmed, "QString_trimmed")
	symaddr(&addr_QString_trimmed2, "QString_trimmed2")
	symaddr(&addr_QString_simplified, "QString_simplified")
	symaddr(&addr_QString_simplified2, "QString_simplified2")
	symaddr(&addr_QString_toHtmlEscaped, "QString_toHtmlEscaped")
	symaddr(&addr_QString_insert, "QString_insert")
	symaddr(&addr_QString_insert2, "QString_insert2")
	symaddr(&addr_QString_insert3, "QString_insert3")
	symaddr(&addr_QString_insert4, "QString_insert4")
	symaddr(&addr_QString_insert5, "QString_insert5")
	symaddr(&addr_QString_append, "QString_append")
	symaddr(&addr_QString_append2, "QString_append2")
	symaddr(&addr_QString_append3, "QString_append3")
	symaddr(&addr_QString_append4, "QString_append4")
	symaddr(&addr_QString_append5, "QString_append5")
	symaddr(&addr_QString_prepend, "QString_prepend")
	symaddr(&addr_QString_prepend2, "QString_prepend2")
	symaddr(&addr_QString_prepend3, "QString_prepend3")
	symaddr(&addr_QString_prepend4, "QString_prepend4")
	symaddr(&addr_QString_prepend5, "QString_prepend5")
	symaddr(&addr_QString_assign, "QString_assign")
	symaddr(&addr_QString_assign2, "QString_assign2")
	symaddr(&addr_QString_OperatorPlusAssign, "QString_OperatorPlusAssign")
	symaddr(&addr_QString_OperatorPlusAssign2, "QString_OperatorPlusAssign2")
	symaddr(&addr_QString_OperatorPlusAssign3, "QString_OperatorPlusAssign3")
	symaddr(&addr_QString_OperatorPlusAssign4, "QString_OperatorPlusAssign4")
	symaddr(&addr_QString_remove, "QString_remove")
	symaddr(&addr_QString_remove2, "QString_remove2")
	symaddr(&addr_QString_remove3, "QString_remove3")
	symaddr(&addr_QString_remove4, "QString_remove4")
	symaddr(&addr_QString_removeAt, "QString_removeAt")
	symaddr(&addr_QString_removeFirst, "QString_removeFirst")
	symaddr(&addr_QString_removeLast, "QString_removeLast")
	symaddr(&addr_QString_replace, "QString_replace")
	symaddr(&addr_QString_replace2, "QString_replace2")
	symaddr(&addr_QString_replace3, "QString_replace3")
	symaddr(&addr_QString_replace4, "QString_replace4")
	symaddr(&addr_QString_replace5, "QString_replace5")
	symaddr(&addr_QString_replace6, "QString_replace6")
	symaddr(&addr_QString_replace7, "QString_replace7")
	symaddr(&addr_QString_replace8, "QString_replace8")
	symaddr(&addr_QString_replace9, "QString_replace9")
	symaddr(&addr_QString_replace10, "QString_replace10")
	symaddr(&addr_QString_replace11, "QString_replace11")
	symaddr(&addr_QString_replace12, "QString_replace12")
	symaddr(&addr_QString_remove5, "QString_remove5")
	symaddr(&addr_QString_split, "QString_split")
	symaddr(&addr_QString_split2, "QString_split2")
	symaddr(&addr_QString_split3, "QString_split3")
	symaddr(&addr_QString_normalized, "QString_normalized")
	symaddr(&addr_QString_repeated, "QString_repeated")
	symaddr(&addr_QString_nullTerminated, "QString_nullTerminated")
	symaddr(&addr_QString_nullTerminated2, "QString_nullTerminated2")
	symaddr(&addr_QString_nullTerminate, "QString_nullTerminate")
	symaddr(&addr_QString_toLatin1, "QString_toLatin1")
	symaddr(&addr_QString_toLatin12, "QString_toLatin12")
	symaddr(&addr_QString_toUtf8, "QString_toUtf8")
	symaddr(&addr_QString_toUtf82, "QString_toUtf82")
	symaddr(&addr_QString_toLocal8Bit, "QString_toLocal8Bit")
	symaddr(&addr_QString_toLocal8Bit2, "QString_toLocal8Bit2")
	symaddr(&addr_QString_toUcs4, "QString_toUcs4")
	symaddr(&addr_QString_fromLatin1, "QString_fromLatin1")
	symaddr(&addr_QString_fromLatin12, "QString_fromLatin12")
	symaddr(&addr_QString_fromUtf8, "QString_fromUtf8")
	symaddr(&addr_QString_fromUtf82, "QString_fromUtf82")
	symaddr(&addr_QString_fromLocal8Bit, "QString_fromLocal8Bit")
	symaddr(&addr_QString_fromLocal8Bit2, "QString_fromLocal8Bit2")
	symaddr(&addr_QString_fromRawData, "QString_fromRawData")
	symaddr(&addr_QString_setRawData, "QString_setRawData")
	symaddr(&addr_QString_setUnicode, "QString_setUnicode")
	symaddr(&addr_QString_compare, "QString_compare")
	symaddr(&addr_QString_compare2, "QString_compare2")
	symaddr(&addr_QString_compare3, "QString_compare3")
	symaddr(&addr_QString_compare4, "QString_compare4")
	symaddr(&addr_QString_compare5, "QString_compare5")
	symaddr(&addr_QString_compare6, "QString_compare6")
	symaddr(&addr_QString_compare7, "QString_compare7")
	symaddr(&addr_QString_compare8, "QString_compare8")
	symaddr(&addr_QString_compare9, "QString_compare9")
	symaddr(&addr_QString_localeAwareCompare, "QString_localeAwareCompare")
	symaddr(&addr_QString_localeAwareCompare2, "QString_localeAwareCompare2")
	symaddr(&addr_QString_localeAwareCompare3, "QString_localeAwareCompare3")
	symaddr(&addr_QString_localeAwareCompare4, "QString_localeAwareCompare4")
	symaddr(&addr_QString_toShort, "QString_toShort")
	symaddr(&addr_QString_toUShort, "QString_toUShort")
	symaddr(&addr_QString_toInt, "QString_toInt")
	symaddr(&addr_QString_toUInt, "QString_toUInt")
	symaddr(&addr_QString_toLong, "QString_toLong")
	symaddr(&addr_QString_toULong, "QString_toULong")
	symaddr(&addr_QString_toLongLong, "QString_toLongLong")
	symaddr(&addr_QString_toULongLong, "QString_toULongLong")
	reg(&pg_QString_toFloat, "QString_toFloat")
	reg(&pg_QString_toDouble, "QString_toDouble")
	symaddr(&addr_QString_setNum, "QString_setNum")
	symaddr(&addr_QString_setNum2, "QString_setNum2")
	symaddr(&addr_QString_setNum3, "QString_setNum3")
	symaddr(&addr_QString_setNum4, "QString_setNum4")
	symaddr(&addr_QString_setNum5, "QString_setNum5")
	symaddr(&addr_QString_setNum6, "QString_setNum6")
	symaddr(&addr_QString_setNum7, "QString_setNum7")
	symaddr(&addr_QString_setNum8, "QString_setNum8")
	reg(&pg_QString_setNum9, "QString_setNum9")
	reg(&pg_QString_setNum10, "QString_setNum10")
	symaddr(&addr_QString_number, "QString_number")
	symaddr(&addr_QString_number2, "QString_number2")
	symaddr(&addr_QString_number3, "QString_number3")
	symaddr(&addr_QString_number4, "QString_number4")
	symaddr(&addr_QString_number5, "QString_number5")
	symaddr(&addr_QString_number6, "QString_number6")
	reg(&pg_QString_number7, "QString_number7")
	symaddr(&addr_QString_OperatorAssign4, "QString_OperatorAssign4")
	symaddr(&addr_QString_OperatorAssign5, "QString_OperatorAssign5")
	symaddr(&addr_QString_prepend6, "QString_prepend6")
	symaddr(&addr_QString_prepend7, "QString_prepend7")
	symaddr(&addr_QString_append6, "QString_append6")
	symaddr(&addr_QString_append7, "QString_append7")
	symaddr(&addr_QString_insert6, "QString_insert6")
	symaddr(&addr_QString_insert7, "QString_insert7")
	symaddr(&addr_QString_OperatorPlusAssign5, "QString_OperatorPlusAssign5")
	symaddr(&addr_QString_OperatorPlusAssign6, "QString_OperatorPlusAssign6")
	symaddr(&addr_QString_begin, "QString_begin")
	symaddr(&addr_QString_begin2, "QString_begin2")
	symaddr(&addr_QString_cbegin, "QString_cbegin")
	symaddr(&addr_QString_constBegin, "QString_constBegin")
	symaddr(&addr_QString_end, "QString_end")
	symaddr(&addr_QString_end2, "QString_end2")
	symaddr(&addr_QString_cend, "QString_cend")
	symaddr(&addr_QString_constEnd, "QString_constEnd")
	symaddr(&addr_QString_push_back, "QString_push_back")
	symaddr(&addr_QString_push_back2, "QString_push_back2")
	symaddr(&addr_QString_push_front, "QString_push_front")
	symaddr(&addr_QString_push_front2, "QString_push_front2")
	symaddr(&addr_QString_shrink_to_fit, "QString_shrink_to_fit")
	symaddr(&addr_QString_erase, "QString_erase")
	symaddr(&addr_QString_erase2, "QString_erase2")
	symaddr(&addr_QString_max_size, "QString_max_size")
	symaddr(&addr_QString_isNull, "QString_isNull")
	symaddr(&addr_QString_isRightToLeft, "QString_isRightToLeft")
	symaddr(&addr_QString_isValidUtf16, "QString_isValidUtf16")
	symaddr(&addr_QString_fill2, "QString_fill2")
	symaddr(&addr_QString_indexOf6, "QString_indexOf6")
	symaddr(&addr_QString_indexOf7, "QString_indexOf7")
	symaddr(&addr_QString_indexOf8, "QString_indexOf8")
	symaddr(&addr_QString_indexOf9, "QString_indexOf9")
	symaddr(&addr_QString_indexOf10, "QString_indexOf10")
	symaddr(&addr_QString_indexOf11, "QString_indexOf11")
	symaddr(&addr_QString_indexOf12, "QString_indexOf12")
	symaddr(&addr_QString_indexOf13, "QString_indexOf13")
	symaddr(&addr_QString_lastIndexOf10, "QString_lastIndexOf10")
	symaddr(&addr_QString_lastIndexOf11, "QString_lastIndexOf11")
	symaddr(&addr_QString_lastIndexOf12, "QString_lastIndexOf12")
	symaddr(&addr_QString_lastIndexOf13, "QString_lastIndexOf13")
	symaddr(&addr_QString_lastIndexOf14, "QString_lastIndexOf14")
	symaddr(&addr_QString_lastIndexOf15, "QString_lastIndexOf15")
	symaddr(&addr_QString_lastIndexOf16, "QString_lastIndexOf16")
	symaddr(&addr_QString_lastIndexOf17, "QString_lastIndexOf17")
	symaddr(&addr_QString_contains6, "QString_contains6")
	symaddr(&addr_QString_contains7, "QString_contains7")
	symaddr(&addr_QString_contains8, "QString_contains8")
	symaddr(&addr_QString_contains9, "QString_contains9")
	symaddr(&addr_QString_count6, "QString_count6")
	symaddr(&addr_QString_count7, "QString_count7")
	symaddr(&addr_QString_count8, "QString_count8")
	symaddr(&addr_QString_indexOf14, "QString_indexOf14")
	symaddr(&addr_QString_indexOf15, "QString_indexOf15")
	symaddr(&addr_QString_lastIndexOf18, "QString_lastIndexOf18")
	symaddr(&addr_QString_contains10, "QString_contains10")
	symaddr(&addr_QString_section4, "QString_section4")
	symaddr(&addr_QString_section5, "QString_section5")
	symaddr(&addr_QString_section6, "QString_section6")
	symaddr(&addr_QString_section7, "QString_section7")
	symaddr(&addr_QString_section8, "QString_section8")
	symaddr(&addr_QString_section9, "QString_section9")
	symaddr(&addr_QString_mid3, "QString_mid3")
	symaddr(&addr_QString_mid4, "QString_mid4")
	symaddr(&addr_QString_startsWith5, "QString_startsWith5")
	symaddr(&addr_QString_startsWith6, "QString_startsWith6")
	symaddr(&addr_QString_startsWith7, "QString_startsWith7")
	symaddr(&addr_QString_startsWith8, "QString_startsWith8")
	symaddr(&addr_QString_endsWith5, "QString_endsWith5")
	symaddr(&addr_QString_endsWith6, "QString_endsWith6")
	symaddr(&addr_QString_endsWith7, "QString_endsWith7")
	symaddr(&addr_QString_endsWith8, "QString_endsWith8")
	symaddr(&addr_QString_leftJustified2, "QString_leftJustified2")
	symaddr(&addr_QString_leftJustified3, "QString_leftJustified3")
	symaddr(&addr_QString_rightJustified2, "QString_rightJustified2")
	symaddr(&addr_QString_rightJustified3, "QString_rightJustified3")
	symaddr(&addr_QString_remove6, "QString_remove6")
	symaddr(&addr_QString_remove7, "QString_remove7")
	symaddr(&addr_QString_remove8, "QString_remove8")
	symaddr(&addr_QString_replace13, "QString_replace13")
	symaddr(&addr_QString_replace14, "QString_replace14")
	symaddr(&addr_QString_replace15, "QString_replace15")
	symaddr(&addr_QString_replace16, "QString_replace16")
	symaddr(&addr_QString_replace17, "QString_replace17")
	symaddr(&addr_QString_replace18, "QString_replace18")
	symaddr(&addr_QString_replace19, "QString_replace19")
	symaddr(&addr_QString_replace20, "QString_replace20")
	symaddr(&addr_QString_split4, "QString_split4")
	symaddr(&addr_QString_split5, "QString_split5")
	symaddr(&addr_QString_split6, "QString_split6")
	symaddr(&addr_QString_split7, "QString_split7")
	symaddr(&addr_QString_split8, "QString_split8")
	symaddr(&addr_QString_normalized2, "QString_normalized2")
	symaddr(&addr_QString_compare10, "QString_compare10")
	symaddr(&addr_QString_compare11, "QString_compare11")
	symaddr(&addr_QString_compare12, "QString_compare12")
	symaddr(&addr_QString_compare13, "QString_compare13")
	symaddr(&addr_QString_compare14, "QString_compare14")
	symaddr(&addr_QString_compare15, "QString_compare15")
	symaddr(&addr_QString_compare16, "QString_compare16")
	symaddr(&addr_QString_compare17, "QString_compare17")
	symaddr(&addr_QString_compare18, "QString_compare18")
	symaddr(&addr_QString_toShort2, "QString_toShort2")
	symaddr(&addr_QString_toShort3, "QString_toShort3")
	symaddr(&addr_QString_toUShort2, "QString_toUShort2")
	symaddr(&addr_QString_toUShort3, "QString_toUShort3")
	symaddr(&addr_QString_toInt2, "QString_toInt2")
	symaddr(&addr_QString_toInt3, "QString_toInt3")
	symaddr(&addr_QString_toUInt2, "QString_toUInt2")
	symaddr(&addr_QString_toUInt3, "QString_toUInt3")
	symaddr(&addr_QString_toLong2, "QString_toLong2")
	symaddr(&addr_QString_toLong3, "QString_toLong3")
	symaddr(&addr_QString_toULong2, "QString_toULong2")
	symaddr(&addr_QString_toULong3, "QString_toULong3")
	symaddr(&addr_QString_toLongLong2, "QString_toLongLong2")
	symaddr(&addr_QString_toLongLong3, "QString_toLongLong3")
	symaddr(&addr_QString_toULongLong2, "QString_toULongLong2")
	symaddr(&addr_QString_toULongLong3, "QString_toULongLong3")
	reg(&pg_QString_toFloat2, "QString_toFloat2")
	reg(&pg_QString_toDouble2, "QString_toDouble2")
	symaddr(&addr_QString_setNum11, "QString_setNum11")
	symaddr(&addr_QString_setNum12, "QString_setNum12")
	symaddr(&addr_QString_setNum13, "QString_setNum13")
	symaddr(&addr_QString_setNum14, "QString_setNum14")
	symaddr(&addr_QString_setNum15, "QString_setNum15")
	symaddr(&addr_QString_setNum16, "QString_setNum16")
	symaddr(&addr_QString_setNum17, "QString_setNum17")
	symaddr(&addr_QString_setNum18, "QString_setNum18")
	reg(&pg_QString_setNum19, "QString_setNum19")
	reg(&pg_QString_setNum20, "QString_setNum20")
	reg(&pg_QString_setNum21, "QString_setNum21")
	reg(&pg_QString_setNum22, "QString_setNum22")
	symaddr(&addr_QString_number8, "QString_number8")
	symaddr(&addr_QString_number9, "QString_number9")
	symaddr(&addr_QString_number10, "QString_number10")
	symaddr(&addr_QString_number11, "QString_number11")
	symaddr(&addr_QString_number12, "QString_number12")
	symaddr(&addr_QString_number13, "QString_number13")
	reg(&pg_QString_number14, "QString_number14")
	reg(&pg_QString_number15, "QString_number15")
	symaddr(&addr_QString_OperatorEqual, "QString_OperatorEqual")
	symaddr(&addr_QString_OperatorNotEqual, "QString_OperatorNotEqual")
	symaddr(&addr_QString_OperatorLesser, "QString_OperatorLesser")
	symaddr(&addr_QString_OperatorGreater, "QString_OperatorGreater")
	symaddr(&addr_QString_OperatorLesserOrEqual, "QString_OperatorLesserOrEqual")
	symaddr(&addr_QString_OperatorGreaterOrEqual, "QString_OperatorGreaterOrEqual")
	symaddr(&addr_QString_OperatorEqual2, "QString_OperatorEqual2")
	symaddr(&addr_QString_OperatorNotEqual2, "QString_OperatorNotEqual2")
	symaddr(&addr_QString_OperatorLesser2, "QString_OperatorLesser2")
	symaddr(&addr_QString_OperatorGreater2, "QString_OperatorGreater2")
	symaddr(&addr_QString_OperatorLesserOrEqual2, "QString_OperatorLesserOrEqual2")
	symaddr(&addr_QString_OperatorGreaterOrEqual2, "QString_OperatorGreaterOrEqual2")
	symaddr(&addr_QString_OperatorEqual3, "QString_OperatorEqual3")
	symaddr(&addr_QString_OperatorNotEqual3, "QString_OperatorNotEqual3")
	symaddr(&addr_QString_OperatorLesser3, "QString_OperatorLesser3")
	symaddr(&addr_QString_OperatorGreater3, "QString_OperatorGreater3")
	symaddr(&addr_QString_OperatorLesserOrEqual3, "QString_OperatorLesserOrEqual3")
	symaddr(&addr_QString_OperatorGreaterOrEqual3, "QString_OperatorGreaterOrEqual3")
	symaddr(&addr_QString_OperatorEqual4, "QString_OperatorEqual4")
	symaddr(&addr_QString_OperatorNotEqual4, "QString_OperatorNotEqual4")
	symaddr(&addr_QString_OperatorLesser4, "QString_OperatorLesser4")
	symaddr(&addr_QString_OperatorGreater4, "QString_OperatorGreater4")
	symaddr(&addr_QString_OperatorLesserOrEqual4, "QString_OperatorLesserOrEqual4")
	symaddr(&addr_QString_OperatorGreaterOrEqual4, "QString_OperatorGreaterOrEqual4")
	symaddr(&addr_QString_OperatorPlus, "QString_OperatorPlus")
	symaddr(&addr_QString_OperatorPlus2, "QString_OperatorPlus2")
	symaddr(&addr_QString_OperatorPlus3, "QString_OperatorPlus3")
	symaddr(&addr_QString_OperatorPlus4, "QString_OperatorPlus4")
	symaddr(&addr_QString_OperatorPlus5, "QString_OperatorPlus5")
	reg(&pg_QString_delete, "QString_delete")
}

type QString struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QString) UnsafeGoID() int64     { return this.vid }
func (this *QString) UnsafeSetGoID(id int64) { this.vid = id }

func bindQString(h uintptr) *QString {
	if h == 0 {
		return nil
	}
	return &QString{h: h}
}

func newownedQString(h uintptr) *QString {
	o := bindQString(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QString).Delete)
	}
	return o
}

func (this *QString) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QString) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQString(h unsafe.Pointer) *QString { return bindQString(uintptr(h)) }
func UnsafeNewQStringOwned(h unsafe.Pointer) *QString { return newownedQString(uintptr(h)) }

func (this *QString) UnsafeRelease() { this.h = 0 }

func NewQString() *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new)
	__o := newownedQString(__r)
	return __o
}

func NewQString2(unicode *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new2, uintptr(unicode.UnsafePointer()))
	__o := newownedQString(__r)
	runtime.KeepAlive(unicode)
	return __o
}

func NewQString3(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new3, uintptr(c.UnsafePointer()))
	__o := newownedQString(__r)
	runtime.KeepAlive(c)
	return __o
}

func NewQString4(size int64, c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new4, uintptr(size), uintptr(c.UnsafePointer()))
	__o := newownedQString(__r)
	runtime.KeepAlive(c)
	return __o
}

func NewQString5(latin1 *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new5, uintptr(latin1.UnsafePointer()))
	__o := newownedQString(__r)
	runtime.KeepAlive(latin1)
	return __o
}

func NewQString6(sv *QStringView) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new6, uintptr(sv.UnsafePointer()))
	__o := newownedQString(__r)
	runtime.KeepAlive(sv)
	return __o
}

func NewQString7(arg1 string) *QString {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_new7, arg1_p, uintptr(len(arg1)))
	__o := newownedQString(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQString8(ch string) *QString {
	ch_b := append([]byte(ch), 0)
	__r, _, _ := purego.SyscallN(addr_QString_new8, uintptr(unsafe.Pointer(&ch_b[0])))
	__o := newownedQString(__r)
	runtime.KeepAlive(ch_b)
	return __o
}

func NewQString9(a *QByteArray) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new9, uintptr(a.UnsafePointer()))
	__o := newownedQString(__r)
	runtime.KeepAlive(a)
	return __o
}

func NewQString10(size int64, arg2 Initialization) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new10, uintptr(size), uintptr(arg2))
	__o := newownedQString(__r)
	return __o
}

func NewQString11(unicode *QChar, size int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_new11, uintptr(unicode.UnsafePointer()), uintptr(size))
	__o := newownedQString(__r)
	runtime.KeepAlive(unicode)
	return __o
}

func (this *QString) OperatorAssign(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorAssign, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) OperatorAssignWithQString(arg1 string) *QString {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_OperatorAssign2, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQString(__r)
}

func (this *QString) OperatorAssignWithLatin1(latin1 *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorAssign3, this.h, uintptr(latin1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(latin1)
	return bindQString(__r)
}

func (this *QString) Swap(other *QString) {
	purego.SyscallN(addr_QString_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func QString_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QString) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) Resize(size int64) {
	purego.SyscallN(addr_QString_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QString) Resize2(size int64, fillChar *QChar) {
	purego.SyscallN(addr_QString_resize2, this.h, uintptr(size), uintptr(fillChar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fillChar)
}

func (this *QString) ResizeForOverwrite(size int64) {
	purego.SyscallN(addr_QString_resizeForOverwrite, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QString) Fill(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_fill, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) Truncate(pos int64) {
	purego.SyscallN(addr_QString_truncate, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QString) Chop(n int64) {
	purego.SyscallN(addr_QString_chop, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QString) Slice(pos int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_slice, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) Slice2(pos int64, n int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_slice2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) Reserve(size int64) {
	purego.SyscallN(addr_QString_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QString) Squeeze() {
	purego.SyscallN(addr_QString_squeeze, this.h)
	runtime.KeepAlive(this)
}

func (this *QString) Unicode() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_unicode, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Data() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_data, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Data2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_data2, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) ConstData() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_constData, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Detach() {
	purego.SyscallN(addr_QString_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QString) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) IsSharedWith(other string) bool {
	var other_p uintptr
	if len(other) > 0 {
		other_p = uintptr(unsafe.Pointer(unsafe.StringData(other)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_isSharedWith, this.h, other_p, uintptr(len(other)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QString) Clear() {
	purego.SyscallN(addr_QString_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QString) At(i int64) *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QString) OperatorSubscript(i int64) *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QString) OperatorSubscriptWith(i int64) *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Front() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_front, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QString) Front2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_front2, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Back() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_back, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QString) Back2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_back2, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func QString_Asprintf(format string) string {
	format_b := append([]byte(format), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_asprintf, uintptr(unsafe.Pointer(&format_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(format_b)
	return pgStr(__r, __n)
}

func (this *QString) IndexOf(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) IndexOfWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf2, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOfWithQString(s string) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf3, this.h, s_p, uintptr(len(s)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOfWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf4, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) LastIndexOf2(c *QChar, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf2, this.h, uintptr(c.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) LastIndexOfWith(s *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf3, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf3(s *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf4, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOfWithQString(s string) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf5, this.h, s_p, uintptr(len(s)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf4(s string, from int64) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf6, this.h, s_p, uintptr(len(s)), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOfWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf7, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf5(s *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf8, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) Contains(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QString) ContainsWithQString(s string) bool {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_contains2, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) ContainsWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains3, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) ContainsWithQStringView(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains4, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) CountWithQChar(c *QChar) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_count2, this.h, uintptr(c.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) CountWithQString(s string) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_count3, this.h, s_p, uintptr(len(s)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) CountWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_count4, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOfWithRe(re *QRegularExpression) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf5, this.h, uintptr(re.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QString) LastIndexOf6(re *QRegularExpression, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf9, this.h, uintptr(re.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QString) ContainsWithRe(re *QRegularExpression) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains5, this.h, uintptr(re.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r != 0
}

func (this *QString) CountWithRe(re *QRegularExpression) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_count5, this.h, uintptr(re.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QString) Section(sep *QChar, start int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section, this.h, uintptr(sep.UnsafePointer()), uintptr(start), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	return pgStr(__r, __n)
}

func (this *QString) Section2(in_sep string, start int64) string {
	var in_sep_p uintptr
	if len(in_sep) > 0 {
		in_sep_p = uintptr(unsafe.Pointer(unsafe.StringData(in_sep)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section2, this.h, in_sep_p, uintptr(len(in_sep)), uintptr(start), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in_sep)
	return pgStr(__r, __n)
}

func (this *QString) Section3(re *QRegularExpression, start int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section3, this.h, uintptr(re.UnsafePointer()), uintptr(start), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return pgStr(__r, __n)
}

func (this *QString) Left(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_left, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) LeftWith(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_left2, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Right(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_right, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) RightWith(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_right2, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Mid(position int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_mid, this.h, uintptr(position), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) MidWithPosition(position int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_mid2, this.h, uintptr(position), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) First(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_first, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Last(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_last, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Sliced(pos int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_sliced, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Sliced2(pos int64, n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_sliced2, this.h, uintptr(pos), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Chopped(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_chopped, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) FirstWith(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_first2, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) LastWith(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_last2, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) SlicedWithPos(pos int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_sliced3, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Sliced3(pos int64, n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_sliced4, this.h, uintptr(pos), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ChoppedWith(n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_chopped2, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) StartsWith(s string) bool {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) StartsWithWithQStringView(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) StartsWithWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith3, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) StartsWithWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith4, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QString) EndsWith(s string) bool {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) EndsWithWithQStringView(s *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) EndsWithWith(s *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith3, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) EndsWithWithQChar(c *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith4, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QString) IsUpper() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isUpper, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) IsLower() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isLower, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) LeftJustified(width int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_leftJustified, this.h, uintptr(width), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) RightJustified(width int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_rightJustified, this.h, uintptr(width), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToLower() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toLower, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToLower2() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toLower2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToUpper() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toUpper, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToUpper2() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toUpper2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToCaseFolded() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toCaseFolded, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToCaseFolded2() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toCaseFolded2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Trimmed() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_trimmed, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Trimmed2() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_trimmed2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Simplified() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_simplified, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Simplified2() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_simplified2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) ToHtmlEscaped() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toHtmlEscaped, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Insert(i int64, c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_insert, this.h, uintptr(i), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) Insert2(i int64, uc *QChar, len_ int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_insert2, this.h, uintptr(i), uintptr(uc.UnsafePointer()), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(uc)
	return bindQString(__r)
}

func (this *QString) Insert3(i int64, s string) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_insert3, this.h, uintptr(i), s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Insert4(i int64, v *QStringView) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_insert4, this.h, uintptr(i), uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQString(__r)
}

func (this *QString) Insert5(i int64, s *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_insert5, this.h, uintptr(i), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Append(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_append, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) Append2(uc *QChar, len_ int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_append2, this.h, uintptr(uc.UnsafePointer()), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(uc)
	return bindQString(__r)
}

func (this *QString) AppendWithQString(s string) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_append3, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) AppendWithQStringView(v *QStringView) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_append4, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQString(__r)
}

func (this *QString) AppendWith(s *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_append5, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Prepend(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_prepend, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) Prepend2(uc *QChar, len_ int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_prepend2, this.h, uintptr(uc.UnsafePointer()), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(uc)
	return bindQString(__r)
}

func (this *QString) PrependWithQString(s string) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_prepend3, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) PrependWithQStringView(v *QStringView) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_prepend4, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQString(__r)
}

func (this *QString) PrependWith(s *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_prepend5, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Assign(s string) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_assign, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Assign2(n int64, c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_assign2, this.h, uintptr(n), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) OperatorPlusAssign(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlusAssign, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) OperatorPlusAssignWithQString(s string) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlusAssign2, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) OperatorPlusAssignWithQStringView(v *QStringView) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlusAssign3, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
	return bindQString(__r)
}

func (this *QString) OperatorPlusAssignWith(s *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlusAssign4, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Remove(i int64, len_ int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_remove, this.h, uintptr(i), uintptr(len_))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) RemoveWithQChar(c *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_remove2, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) RemoveWith(s *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_remove3, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) RemoveWithQString(s string) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_remove4, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) RemoveAt(pos int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_removeAt, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) RemoveFirst() *QString {
	__r, _, _ := purego.SyscallN(addr_QString_removeFirst, this.h)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) RemoveLast() *QString {
	__r, _, _ := purego.SyscallN(addr_QString_removeLast, this.h)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) Replace(i int64, len_ int64, after *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace, this.h, uintptr(i), uintptr(len_), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace2(i int64, len_ int64, s *QChar, slen int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace2, this.h, uintptr(i), uintptr(len_), uintptr(s.UnsafePointer()), uintptr(slen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Replace3(i int64, len_ int64, after string) *QString {
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace3, this.h, uintptr(i), uintptr(len_), after_p, uintptr(len(after)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace4(before *QChar, after *QChar) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace4, this.h, uintptr(before.UnsafePointer()), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace5(before *QChar, blen int64, after *QChar, alen int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace5, this.h, uintptr(before.UnsafePointer()), uintptr(blen), uintptr(after.UnsafePointer()), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace6(before *QLatin1String, after *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace6, this.h, uintptr(before.UnsafePointer()), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace7(before *QLatin1String, after string) *QString {
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace7, this.h, uintptr(before.UnsafePointer()), after_p, uintptr(len(after)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace8(before string, after *QLatin1String) *QString {
	var before_p uintptr
	if len(before) > 0 {
		before_p = uintptr(unsafe.Pointer(unsafe.StringData(before)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace8, this.h, before_p, uintptr(len(before)), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace9(before string, after string) *QString {
	var before_p uintptr
	if len(before) > 0 {
		before_p = uintptr(unsafe.Pointer(unsafe.StringData(before)))
	}
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace9, this.h, before_p, uintptr(len(before)), after_p, uintptr(len(after)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace10(c *QChar, after string) *QString {
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace10, this.h, uintptr(c.UnsafePointer()), after_p, uintptr(len(after)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace11(c *QChar, after *QLatin1String) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace11, this.h, uintptr(c.UnsafePointer()), uintptr(after.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace12(re *QRegularExpression, after string) *QString {
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace12, this.h, uintptr(re.UnsafePointer()), after_p, uintptr(len(after)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) RemoveWithRe(re *QRegularExpression) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_remove5, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return bindQString(__r)
}

func (this *QString) Split(sep string) []string {
	var sep_p uintptr
	if len(sep) > 0 {
		sep_p = uintptr(unsafe.Pointer(unsafe.StringData(sep)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split, this.h, sep_p, uintptr(len(sep)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) SplitWithSep(sep *QChar) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split2, this.h, uintptr(sep.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Split2(sep *QRegularExpression) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split3, this.h, uintptr(sep.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Normalized(mode QString__NormalizationForm) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_normalized, this.h, uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Repeated(times int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_repeated, this.h, uintptr(times), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) NullTerminated() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_nullTerminated, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) NullTerminated2() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_nullTerminated2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) NullTerminate() *QString {
	__r, _, _ := purego.SyscallN(addr_QString_nullTerminate, this.h)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) ToLatin1() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QString_toLatin1, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QString) ToLatin12() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QString_toLatin12, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QString) ToUtf8() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QString_toUtf8, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QString) ToUtf82() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QString_toUtf82, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QString) ToLocal8Bit() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QString_toLocal8Bit, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QString) ToLocal8Bit2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QString_toLocal8Bit2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QString) ToUcs4() []uint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_toUcs4, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]uint, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = uint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QString_FromLatin1(ba *QByteArrayView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromLatin1, uintptr(ba.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(ba)
	return pgStr(__r, __n)
}

func QString_FromLatin12(str string, size int64) string {
	str_b := append([]byte(str), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromLatin12, uintptr(unsafe.Pointer(&str_b[0])), uintptr(size), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str_b)
	return pgStr(__r, __n)
}

func QString_FromUtf8(utf8 *QByteArrayView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromUtf8, uintptr(utf8.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(utf8)
	return pgStr(__r, __n)
}

func QString_FromUtf82(utf8 string, size int64) string {
	utf8_b := append([]byte(utf8), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromUtf82, uintptr(unsafe.Pointer(&utf8_b[0])), uintptr(size), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(utf8_b)
	return pgStr(__r, __n)
}

func QString_FromLocal8Bit(ba *QByteArrayView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromLocal8Bit, uintptr(ba.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(ba)
	return pgStr(__r, __n)
}

func QString_FromLocal8Bit2(str string, size int64) string {
	str_b := append([]byte(str), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromLocal8Bit2, uintptr(unsafe.Pointer(&str_b[0])), uintptr(size), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str_b)
	return pgStr(__r, __n)
}

func QString_FromRawData(arg1 *QChar, size int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_fromRawData, uintptr(arg1.UnsafePointer()), uintptr(size), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func (this *QString) SetRawData(unicode *QChar, size int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setRawData, this.h, uintptr(unicode.UnsafePointer()), uintptr(size))
	runtime.KeepAlive(this)
	runtime.KeepAlive(unicode)
	return bindQString(__r)
}

func (this *QString) SetUnicode(unicode *QChar, size int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setUnicode, this.h, uintptr(unicode.UnsafePointer()), uintptr(size))
	runtime.KeepAlive(this)
	runtime.KeepAlive(unicode)
	return bindQString(__r)
}

func (this *QString) Compare(s string) int {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QString) CompareWithOther(other *QLatin1String) int {
	__r0, _, _ := purego.SyscallN(addr_QString_compare2, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QString) CompareWithQStringView(s *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QString_compare3, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QString) CompareWithCh(ch *QChar) int {
	__r0, _, _ := purego.SyscallN(addr_QString_compare4, this.h, uintptr(ch.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
	return int(__r)
}

func QString_Compare(s1 string, s2 string) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare5, s1_p, uintptr(len(s1)), s2_p, uintptr(len(s2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare2(s1 string, s2 *QLatin1String) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare6, s1_p, uintptr(len(s1)), uintptr(s2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare3(s1 *QLatin1String, s2 string) int {
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare7, uintptr(s1.UnsafePointer()), s2_p, uintptr(len(s2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare4(s1 string, s2 *QStringView) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare8, s1_p, uintptr(len(s1)), uintptr(s2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare5(s1 *QStringView, s2 string) int {
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare9, uintptr(s1.UnsafePointer()), s2_p, uintptr(len(s2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func (this *QString) LocaleAwareCompare(s string) int {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_localeAwareCompare, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QString) LocaleAwareCompareWithQStringView(s *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QString_localeAwareCompare2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func QString_LocaleAwareCompare(s1 string, s2 string) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_localeAwareCompare3, s1_p, uintptr(len(s1)), s2_p, uintptr(len(s2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_LocaleAwareCompare2(s1 *QStringView, s2 *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QString_localeAwareCompare4, uintptr(s1.UnsafePointer()), uintptr(s2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func (this *QString) ToShort() int16 {
	__r0, _, _ := purego.SyscallN(addr_QString_toShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) ToUShort() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QString_toUShort, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QString) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QString_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QString) ToUInt() uint {
	__r0, _, _ := purego.SyscallN(addr_QString_toUInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QString) ToLong() int32 {
	__r0, _, _ := purego.SyscallN(addr_QString_toLong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) ToULong() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QString_toULong, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QString) ToLongLong() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_toLongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) ToULongLong() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QString_toULongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QString) ToFloat() float32 {
	__r := pg_QString_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) ToDouble() float64 {
	__r := pg_QString_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) SetNum(arg1 int16) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNumWith(arg1 uint16) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNumWithInt(arg1 int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum3, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum2(arg1 uint) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum4, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNumWithLong(arg1 int32) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum5, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum3(arg1 uint32) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum6, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum4(arg1 int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum7, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum5(arg1 uint64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum8, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNumWithFloat(arg1 float32) *QString {
	__r := pg_QString_setNum9(this.h, arg1)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNumWithDouble(arg1 float64) *QString {
	__r := pg_QString_setNum10(this.h, arg1)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func QString_Number(arg1 int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_NumberWith(arg1 uint) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number2, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_NumberWithLong(arg1 int32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number3, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number2(arg1 uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number4, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number3(arg1 int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number5, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number4(arg1 uint64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number6, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_NumberWithDouble(arg1 float64) string {
	var __n int64
	__r := pg_QString_number7(arg1, &__n)
	return pgStr(__r, __n)
}

func (this *QString) OperatorAssignWithCh(ch string) *QString {
	ch_b := append([]byte(ch), 0)
	__r, _, _ := purego.SyscallN(addr_QString_OperatorAssign4, this.h, uintptr(unsafe.Pointer(&ch_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch_b)
	return bindQString(__r)
}

func (this *QString) OperatorAssignWithQByteArray(a *QByteArray) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorAssign5, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQString(__r)
}

func (this *QString) PrependWithChar(s string) *QString {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QString_prepend6, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQString(__r)
}

func (this *QString) PrependWithQByteArray(s *QByteArray) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_prepend7, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) AppendWithChar(s string) *QString {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QString_append6, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQString(__r)
}

func (this *QString) AppendWithQByteArray(s *QByteArray) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_append7, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Insert6(i int64, s string) *QString {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QString_insert6, this.h, uintptr(i), uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQString(__r)
}

func (this *QString) Insert7(i int64, s *QByteArray) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_insert7, this.h, uintptr(i), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) OperatorPlusAssignWithChar(s string) *QString {
	s_b := append([]byte(s), 0)
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlusAssign5, this.h, uintptr(unsafe.Pointer(&s_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	return bindQString(__r)
}

func (this *QString) OperatorPlusAssignWithQByteArray(s *QByteArray) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlusAssign6, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Begin() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_begin, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Begin2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_begin2, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Cbegin() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_cbegin, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) ConstBegin() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_constBegin, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) End() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_end, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) End2() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_end2, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Cend() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_cend, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) ConstEnd() *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_constEnd, this.h)
	runtime.KeepAlive(this)
	return bindQChar(__r)
}

func (this *QString) Push_back(c *QChar) {
	purego.SyscallN(addr_QString_push_back, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QString) Push_backWithQString(s string) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	purego.SyscallN(addr_QString_push_back2, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QString) Push_front(c *QChar) {
	purego.SyscallN(addr_QString_push_front, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QString) Push_frontWithQString(s string) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	purego.SyscallN(addr_QString_push_front2, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QString) Shrink_to_fit() {
	purego.SyscallN(addr_QString_shrink_to_fit, this.h)
	runtime.KeepAlive(this)
}

func (this *QString) Erase(first *QChar, last *QChar) *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_erase, this.h, uintptr(first.UnsafePointer()), uintptr(last.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(first)
	runtime.KeepAlive(last)
	return bindQChar(__r)
}

func (this *QString) EraseWithIt(it *QChar) *QChar {
	__r, _, _ := purego.SyscallN(addr_QString_erase2, this.h, uintptr(it.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(it)
	return bindQChar(__r)
}

func (this *QString) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QString) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) IsRightToLeft() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isRightToLeft, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) IsValidUtf16() bool {
	__r0, _, _ := purego.SyscallN(addr_QString_isValidUtf16, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QString) Fill2(c *QChar, size int64) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_fill2, this.h, uintptr(c.UnsafePointer()), uintptr(size))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) IndexOf2(c *QChar, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf6, this.h, uintptr(c.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) IndexOf3(c *QChar, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf7, this.h, uintptr(c.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) IndexOf4(s *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf8, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOf5(s *QLatin1String, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf9, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOf6(s string, from int64) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf10, this.h, s_p, uintptr(len(s)), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOf7(s string, from int64, cs CaseSensitivity) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf11, this.h, s_p, uintptr(len(s)), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOf8(s *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf12, this.h, uintptr(s.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOf9(s *QStringView, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf13, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf7(c *QChar, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf10, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) LastIndexOf8(c *QChar, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf11, this.h, uintptr(c.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) LastIndexOf9(s *QLatin1String, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf12, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf10(s *QLatin1String, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf13, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf11(s string, cs CaseSensitivity) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf14, this.h, s_p, uintptr(len(s)), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf12(s string, from int64, cs CaseSensitivity) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf15, this.h, s_p, uintptr(len(s)), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf13(s *QStringView, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf16, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) LastIndexOf14(s *QStringView, from int64, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf17, this.h, uintptr(s.UnsafePointer()), uintptr(from), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) Contains2(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains6, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QString) Contains3(s string, cs CaseSensitivity) bool {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_contains7, this.h, s_p, uintptr(len(s)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) Contains4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains8, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) Contains5(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains9, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) Count2(c *QChar, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_count6, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r
}

func (this *QString) Count3(s string, cs CaseSensitivity) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_count7, this.h, s_p, uintptr(len(s)), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) Count4(s *QStringView, cs CaseSensitivity) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_count8, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QString) IndexOf10(re *QRegularExpression, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf14, this.h, uintptr(re.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return __r
}

func (this *QString) IndexOf11(re *QRegularExpression, from int64, rmatch *QRegularExpressionMatch) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_indexOf15, this.h, uintptr(re.UnsafePointer()), uintptr(from), uintptr(rmatch.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(rmatch)
	return __r
}

func (this *QString) LastIndexOf15(re *QRegularExpression, from int64, rmatch *QRegularExpressionMatch) int64 {
	__r0, _, _ := purego.SyscallN(addr_QString_lastIndexOf18, this.h, uintptr(re.UnsafePointer()), uintptr(from), uintptr(rmatch.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(rmatch)
	return __r
}

func (this *QString) Contains6(re *QRegularExpression, rmatch *QRegularExpressionMatch) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_contains10, this.h, uintptr(re.UnsafePointer()), uintptr(rmatch.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	runtime.KeepAlive(rmatch)
	return __r != 0
}

func (this *QString) Section4(sep *QChar, start int64, end int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section4, this.h, uintptr(sep.UnsafePointer()), uintptr(start), uintptr(end), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	return pgStr(__r, __n)
}

func (this *QString) Section5(sep *QChar, start int64, end int64, flags QString__SectionFlag) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section5, this.h, uintptr(sep.UnsafePointer()), uintptr(start), uintptr(end), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	return pgStr(__r, __n)
}

func (this *QString) Section6(in_sep string, start int64, end int64) string {
	var in_sep_p uintptr
	if len(in_sep) > 0 {
		in_sep_p = uintptr(unsafe.Pointer(unsafe.StringData(in_sep)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section6, this.h, in_sep_p, uintptr(len(in_sep)), uintptr(start), uintptr(end), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in_sep)
	return pgStr(__r, __n)
}

func (this *QString) Section7(in_sep string, start int64, end int64, flags QString__SectionFlag) string {
	var in_sep_p uintptr
	if len(in_sep) > 0 {
		in_sep_p = uintptr(unsafe.Pointer(unsafe.StringData(in_sep)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section7, this.h, in_sep_p, uintptr(len(in_sep)), uintptr(start), uintptr(end), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in_sep)
	return pgStr(__r, __n)
}

func (this *QString) Section8(re *QRegularExpression, start int64, end int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section8, this.h, uintptr(re.UnsafePointer()), uintptr(start), uintptr(end), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return pgStr(__r, __n)
}

func (this *QString) Section9(re *QRegularExpression, start int64, end int64, flags QString__SectionFlag) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_section9, this.h, uintptr(re.UnsafePointer()), uintptr(start), uintptr(end), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
	return pgStr(__r, __n)
}

func (this *QString) Mid2(position int64, n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_mid3, this.h, uintptr(position), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Mid3(position int64, n int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_mid4, this.h, uintptr(position), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) StartsWith2(s string, cs CaseSensitivity) bool {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith5, this.h, s_p, uintptr(len(s)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) StartsWith3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) StartsWith4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith7, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) StartsWith5(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_startsWith8, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QString) EndsWith2(s string, cs CaseSensitivity) bool {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith5, this.h, s_p, uintptr(len(s)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) EndsWith3(s *QStringView, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith6, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) EndsWith4(s *QLatin1String, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith7, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QString) EndsWith5(c *QChar, cs CaseSensitivity) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_endsWith8, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QString) LeftJustified2(width int64, fill *QChar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_leftJustified2, this.h, uintptr(width), uintptr(fill.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fill)
	return pgStr(__r, __n)
}

func (this *QString) LeftJustified3(width int64, fill *QChar, trunc bool) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_leftJustified3, this.h, uintptr(width), uintptr(fill.UnsafePointer()), uintptr(b2i(trunc)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fill)
	return pgStr(__r, __n)
}

func (this *QString) RightJustified2(width int64, fill *QChar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_rightJustified2, this.h, uintptr(width), uintptr(fill.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fill)
	return pgStr(__r, __n)
}

func (this *QString) RightJustified3(width int64, fill *QChar, trunc bool) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_rightJustified3, this.h, uintptr(width), uintptr(fill.UnsafePointer()), uintptr(b2i(trunc)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fill)
	return pgStr(__r, __n)
}

func (this *QString) Remove2(c *QChar, cs CaseSensitivity) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_remove6, this.h, uintptr(c.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQString(__r)
}

func (this *QString) Remove3(s *QLatin1String, cs CaseSensitivity) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_remove7, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Remove4(s string, cs CaseSensitivity) *QString {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_remove8, this.h, s_p, uintptr(len(s)), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQString(__r)
}

func (this *QString) Replace13(before *QChar, after *QChar, cs CaseSensitivity) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace13, this.h, uintptr(before.UnsafePointer()), uintptr(after.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace14(before *QChar, blen int64, after *QChar, alen int64, cs CaseSensitivity) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace14, this.h, uintptr(before.UnsafePointer()), uintptr(blen), uintptr(after.UnsafePointer()), uintptr(alen), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace15(before *QLatin1String, after *QLatin1String, cs CaseSensitivity) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace15, this.h, uintptr(before.UnsafePointer()), uintptr(after.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace16(before *QLatin1String, after string, cs CaseSensitivity) *QString {
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace16, this.h, uintptr(before.UnsafePointer()), after_p, uintptr(len(after)), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace17(before string, after *QLatin1String, cs CaseSensitivity) *QString {
	var before_p uintptr
	if len(before) > 0 {
		before_p = uintptr(unsafe.Pointer(unsafe.StringData(before)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace17, this.h, before_p, uintptr(len(before)), uintptr(after.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace18(before string, after string, cs CaseSensitivity) *QString {
	var before_p uintptr
	if len(before) > 0 {
		before_p = uintptr(unsafe.Pointer(unsafe.StringData(before)))
	}
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace18, this.h, before_p, uintptr(len(before)), after_p, uintptr(len(after)), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace19(c *QChar, after string, cs CaseSensitivity) *QString {
	var after_p uintptr
	if len(after) > 0 {
		after_p = uintptr(unsafe.Pointer(unsafe.StringData(after)))
	}
	__r, _, _ := purego.SyscallN(addr_QString_replace19, this.h, uintptr(c.UnsafePointer()), after_p, uintptr(len(after)), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Replace20(c *QChar, after *QLatin1String, cs CaseSensitivity) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_replace20, this.h, uintptr(c.UnsafePointer()), uintptr(after.UnsafePointer()), uintptr(cs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	runtime.KeepAlive(after)
	return bindQString(__r)
}

func (this *QString) Split3(sep string, behavior SplitBehaviorFlags) []string {
	var sep_p uintptr
	if len(sep) > 0 {
		sep_p = uintptr(unsafe.Pointer(unsafe.StringData(sep)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split4, this.h, sep_p, uintptr(len(sep)), uintptr(behavior), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Split4(sep string, behavior SplitBehaviorFlags, cs CaseSensitivity) []string {
	var sep_p uintptr
	if len(sep) > 0 {
		sep_p = uintptr(unsafe.Pointer(unsafe.StringData(sep)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split5, this.h, sep_p, uintptr(len(sep)), uintptr(behavior), uintptr(cs), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Split5(sep *QChar, behavior SplitBehaviorFlags) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split6, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Split6(sep *QChar, behavior SplitBehaviorFlags, cs CaseSensitivity) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split7, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(cs), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Split7(sep *QRegularExpression, behavior SplitBehaviorFlags) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_split8, this.h, uintptr(sep.UnsafePointer()), uintptr(behavior), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sep)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QString) Normalized2(mode QString__NormalizationForm, version QChar__UnicodeVersion) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_normalized2, this.h, uintptr(mode), uintptr(version), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QString) Compare2(s string, cs CaseSensitivity) int {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare10, this.h, s_p, uintptr(len(s)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QString) Compare3(other *QLatin1String, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QString_compare11, this.h, uintptr(other.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return int(__r)
}

func (this *QString) Compare4(s *QStringView, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QString_compare12, this.h, uintptr(s.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QString) Compare5(ch *QChar, cs CaseSensitivity) int {
	__r0, _, _ := purego.SyscallN(addr_QString_compare13, this.h, uintptr(ch.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
	return int(__r)
}

func QString_Compare6(s1 string, s2 string, cs CaseSensitivity) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare14, s1_p, uintptr(len(s1)), s2_p, uintptr(len(s2)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare7(s1 string, s2 *QLatin1String, cs CaseSensitivity) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare15, s1_p, uintptr(len(s1)), uintptr(s2.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare8(s1 *QLatin1String, s2 string, cs CaseSensitivity) int {
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare16, uintptr(s1.UnsafePointer()), s2_p, uintptr(len(s2)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare9(s1 string, s2 *QStringView, cs CaseSensitivity) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare17, s1_p, uintptr(len(s1)), uintptr(s2.UnsafePointer()), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QString_Compare10(s1 *QStringView, s2 string, cs CaseSensitivity) int {
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_compare18, uintptr(s1.UnsafePointer()), s2_p, uintptr(len(s2)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func (this *QString) ToShortWithOk() (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToShort2(base int) (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToUShortWithOk() (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toUShort2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QString) ToUShort2(base int) (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toUShort3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r), ok_out != 0
}

func (this *QString) ToIntWithOk() (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QString) ToInt2(base int) (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QString) ToUIntWithOk() (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toUInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QString) ToUInt2(base int) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toUInt3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QString) ToLongWithOk() (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToLong2(base int) (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToULongWithOk() (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toULong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QString) ToULong2(base int) (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toULong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r), ok_out != 0
}

func (this *QString) ToLongLongWithOk() (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toLongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToLongLong2(base int) (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toLongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToULongLongWithOk() (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toULongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QString) ToULongLong2(base int) (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QString_toULongLong3, this.h, uintptr(unsafe.Pointer(&ok_out)), uintptr(base))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QString) ToFloatWithOk() (float32, bool) {
	var ok_out int32
	__r := pg_QString_toFloat2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) ToDoubleWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QString_toDouble2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QString) SetNum6(arg1 int16, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum11, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum7(arg1 uint16, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum12, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum8(arg1 int, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum13, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum9(arg1 uint, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum14, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum10(arg1 int32, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum15, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum11(arg1 uint32, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum16, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum12(arg1 int64, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum17, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum13(arg1 uint64, base int) *QString {
	__r, _, _ := purego.SyscallN(addr_QString_setNum18, this.h, uintptr(arg1), uintptr(base))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum14(arg1 float32, format int8) *QString {
	__r := pg_QString_setNum19(this.h, arg1, format)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum15(arg1 float32, format int8, precision int) *QString {
	__r := pg_QString_setNum20(this.h, arg1, format, int32(precision))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum16(arg1 float64, format int8) *QString {
	__r := pg_QString_setNum21(this.h, arg1, format)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QString) SetNum17(arg1 float64, format int8, precision int) *QString {
	__r := pg_QString_setNum22(this.h, arg1, format, int32(precision))
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func QString_Number5(arg1 int, base int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number8, uintptr(arg1), uintptr(base), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number6(arg1 uint, base int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number9, uintptr(arg1), uintptr(base), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number7(arg1 int32, base int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number10, uintptr(arg1), uintptr(base), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number8(arg1 uint32, base int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number11, uintptr(arg1), uintptr(base), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number9(arg1 int64, base int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number12, uintptr(arg1), uintptr(base), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number10(arg1 uint64, base int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_number13, uintptr(arg1), uintptr(base), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QString_Number11(arg1 float64, format int8) string {
	var __n int64
	__r := pg_QString_number14(arg1, format, &__n)
	return pgStr(__r, __n)
}

func QString_Number12(arg1 float64, format int8, precision int) string {
	var __n int64
	__r := pg_QString_number15(arg1, format, int32(precision), &__n)
	return pgStr(__r, __n)
}

func (this *QString) OperatorEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorNotEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorNotEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesser(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesser, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreater(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreater, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesserOrEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesserOrEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreaterOrEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreaterOrEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorNotEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesserWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreaterWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesserOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreaterOrEqualWithRhs(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorNotEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesser2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreater2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesserOrEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreaterOrEqual2(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorNotEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesser3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesser4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreater3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreater4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorLesserOrEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorLesserOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorGreaterOrEqual3(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QString_OperatorGreaterOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QString) OperatorPlus(s2 string) string {
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlus, this.h, s2_p, uintptr(len(s2)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return pgStr(__r, __n)
}

func (this *QString) OperatorPlusWithS2(s2 *QChar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlus2, this.h, uintptr(s2.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return pgStr(__r, __n)
}

func (this *QString) OperatorPlusWithRhs(rhs *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlus3, this.h, uintptr(rhs.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return pgStr(__r, __n)
}

func (this *QString) OperatorPlus2(s2 string) string {
	s2_b := append([]byte(s2), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlus4, this.h, uintptr(unsafe.Pointer(&s2_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2_b)
	return pgStr(__r, __n)
}

func (this *QString) OperatorPlusWithBa(ba *QByteArray) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QString_OperatorPlus5, this.h, uintptr(ba.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ba)
	return pgStr(__r, __n)
}

func (this *QString) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QString_delete(this.h)
	}
	this.UnsafeRelease()
}
