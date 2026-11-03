package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLocale_new uintptr
	addr_QLocale_new2 uintptr
	addr_QLocale_new3 uintptr
	addr_QLocale_new4 uintptr
	addr_QLocale_new5 uintptr
	addr_QLocale_new6 uintptr
	addr_QLocale_new7 uintptr
	addr_QLocale_new8 uintptr
	addr_QLocale_OperatorAssign uintptr
	addr_QLocale_swap uintptr
	addr_QLocale_language uintptr
	addr_QLocale_script uintptr
	addr_QLocale_territory uintptr
	addr_QLocale_country uintptr
	addr_QLocale_name uintptr
	addr_QLocale_bcp47Name uintptr
	addr_QLocale_nativeLanguageName uintptr
	addr_QLocale_nativeTerritoryName uintptr
	addr_QLocale_nativeCountryName uintptr
	addr_QLocale_toShort uintptr
	addr_QLocale_toUShort uintptr
	addr_QLocale_toInt uintptr
	addr_QLocale_toUInt uintptr
	addr_QLocale_toLong uintptr
	addr_QLocale_toULong uintptr
	addr_QLocale_toLongLong uintptr
	addr_QLocale_toULongLong uintptr
	pg_QLocale_toFloat func(uintptr, string, int64) float32
	pg_QLocale_toDouble func(uintptr, string, int64) float64
	addr_QLocale_toShort2 uintptr
	addr_QLocale_toUShort2 uintptr
	addr_QLocale_toInt2 uintptr
	addr_QLocale_toUInt2 uintptr
	addr_QLocale_toLong2 uintptr
	addr_QLocale_toULong2 uintptr
	addr_QLocale_toLongLong2 uintptr
	addr_QLocale_toULongLong2 uintptr
	pg_QLocale_toFloat2 func(uintptr, uintptr) float32
	pg_QLocale_toDouble2 func(uintptr, uintptr) float64
	addr_QLocale_toString uintptr
	addr_QLocale_toString2 uintptr
	addr_QLocale_toString3 uintptr
	addr_QLocale_toString4 uintptr
	addr_QLocale_toString5 uintptr
	addr_QLocale_toString6 uintptr
	addr_QLocale_toString7 uintptr
	addr_QLocale_toString8 uintptr
	addr_QLocale_toString9 uintptr
	addr_QLocale_toString10 uintptr
	addr_QLocale_toString11 uintptr
	addr_QLocale_toString12 uintptr
	addr_QLocale_toString13 uintptr
	addr_QLocale_toString14 uintptr
	addr_QLocale_toString15 uintptr
	addr_QLocale_toString16 uintptr
	pg_QLocale_toString17 func(uintptr, float64, *int64) uintptr
	pg_QLocale_toString18 func(uintptr, float32, *int64) uintptr
	addr_QLocale_toString19 uintptr
	addr_QLocale_toString20 uintptr
	addr_QLocale_toString21 uintptr
	addr_QLocale_toString22 uintptr
	addr_QLocale_toString23 uintptr
	addr_QLocale_toString24 uintptr
	addr_QLocale_toString25 uintptr
	addr_QLocale_toString26 uintptr
	addr_QLocale_toString27 uintptr
	addr_QLocale_toString28 uintptr
	addr_QLocale_toString29 uintptr
	addr_QLocale_toString30 uintptr
	addr_QLocale_toString31 uintptr
	addr_QLocale_dateFormat uintptr
	addr_QLocale_timeFormat uintptr
	addr_QLocale_dateTimeFormat uintptr
	addr_QLocale_toTime uintptr
	addr_QLocale_toTime2 uintptr
	addr_QLocale_toDate uintptr
	addr_QLocale_toDate2 uintptr
	addr_QLocale_toDateTime uintptr
	addr_QLocale_toDateTime2 uintptr
	addr_QLocale_toDate3 uintptr
	addr_QLocale_toDate4 uintptr
	addr_QLocale_toDateTime3 uintptr
	addr_QLocale_toDateTime4 uintptr
	addr_QLocale_decimalPoint uintptr
	addr_QLocale_groupSeparator uintptr
	addr_QLocale_percent uintptr
	addr_QLocale_zeroDigit uintptr
	addr_QLocale_negativeSign uintptr
	addr_QLocale_positiveSign uintptr
	addr_QLocale_exponential uintptr
	addr_QLocale_monthName uintptr
	addr_QLocale_standaloneMonthName uintptr
	addr_QLocale_dayName uintptr
	addr_QLocale_standaloneDayName uintptr
	addr_QLocale_firstDayOfWeek uintptr
	addr_QLocale_weekdays uintptr
	addr_QLocale_amText uintptr
	addr_QLocale_pmText uintptr
	addr_QLocale_measurementSystem uintptr
	addr_QLocale_collation uintptr
	addr_QLocale_textDirection uintptr
	addr_QLocale_toUpper uintptr
	addr_QLocale_toLower uintptr
	addr_QLocale_currencySymbol uintptr
	addr_QLocale_toCurrencyString uintptr
	addr_QLocale_toCurrencyString2 uintptr
	addr_QLocale_toCurrencyString3 uintptr
	addr_QLocale_toCurrencyString4 uintptr
	addr_QLocale_toCurrencyString5 uintptr
	addr_QLocale_toCurrencyString6 uintptr
	pg_QLocale_toCurrencyString7 func(uintptr, float64, *int64) uintptr
	pg_QLocale_toCurrencyString8 func(uintptr, float32, *int64) uintptr
	addr_QLocale_formattedDataSize uintptr
	addr_QLocale_uiLanguages uintptr
	addr_QLocale_languageToCode uintptr
	addr_QLocale_codeToLanguage uintptr
	addr_QLocale_territoryToCode uintptr
	addr_QLocale_codeToTerritory uintptr
	addr_QLocale_countryToCode uintptr
	addr_QLocale_codeToCountry uintptr
	addr_QLocale_scriptToCode uintptr
	addr_QLocale_codeToScript uintptr
	addr_QLocale_languageToString uintptr
	addr_QLocale_territoryToString uintptr
	addr_QLocale_countryToString uintptr
	addr_QLocale_scriptToString uintptr
	addr_QLocale_setDefault uintptr
	addr_QLocale_c uintptr
	addr_QLocale_system uintptr
	addr_QLocale_matchingLocales uintptr
	addr_QLocale_countriesForLanguage uintptr
	addr_QLocale_setNumberOptions uintptr
	addr_QLocale_numberOptions uintptr
	addr_QLocale_quoteString uintptr
	addr_QLocale_quoteString2 uintptr
	addr_QLocale_createSeparatedList uintptr
	addr_QLocale_name2 uintptr
	addr_QLocale_bcp47Name2 uintptr
	addr_QLocale_toShort3 uintptr
	addr_QLocale_toUShort3 uintptr
	addr_QLocale_toInt3 uintptr
	addr_QLocale_toUInt3 uintptr
	addr_QLocale_toLong3 uintptr
	addr_QLocale_toULong3 uintptr
	addr_QLocale_toLongLong3 uintptr
	addr_QLocale_toULongLong3 uintptr
	pg_QLocale_toFloat3 func(uintptr, string, int64, uintptr) float32
	pg_QLocale_toDouble3 func(uintptr, string, int64, uintptr) float64
	addr_QLocale_toShort4 uintptr
	addr_QLocale_toUShort4 uintptr
	addr_QLocale_toInt4 uintptr
	addr_QLocale_toUInt4 uintptr
	addr_QLocale_toLong4 uintptr
	addr_QLocale_toULong4 uintptr
	addr_QLocale_toLongLong4 uintptr
	addr_QLocale_toULongLong4 uintptr
	pg_QLocale_toFloat4 func(uintptr, uintptr, uintptr) float32
	pg_QLocale_toDouble4 func(uintptr, uintptr, uintptr) float64
	pg_QLocale_toString32 func(uintptr, float64, int8, *int64) uintptr
	pg_QLocale_toString33 func(uintptr, float64, int8, int32, *int64) uintptr
	pg_QLocale_toString34 func(uintptr, float32, int8, *int64) uintptr
	pg_QLocale_toString35 func(uintptr, float32, int8, int32, *int64) uintptr
	addr_QLocale_toString36 uintptr
	addr_QLocale_toString37 uintptr
	addr_QLocale_toString38 uintptr
	addr_QLocale_dateFormat2 uintptr
	addr_QLocale_timeFormat2 uintptr
	addr_QLocale_dateTimeFormat2 uintptr
	addr_QLocale_toTime3 uintptr
	addr_QLocale_toDate5 uintptr
	addr_QLocale_toDate6 uintptr
	addr_QLocale_toDate7 uintptr
	addr_QLocale_toDateTime5 uintptr
	addr_QLocale_toDateTime6 uintptr
	addr_QLocale_toDateTime7 uintptr
	addr_QLocale_toDate8 uintptr
	addr_QLocale_toDate9 uintptr
	addr_QLocale_toDateTime8 uintptr
	addr_QLocale_toDateTime9 uintptr
	addr_QLocale_monthName2 uintptr
	addr_QLocale_standaloneMonthName2 uintptr
	addr_QLocale_dayName2 uintptr
	addr_QLocale_standaloneDayName2 uintptr
	addr_QLocale_currencySymbol2 uintptr
	addr_QLocale_toCurrencyString9 uintptr
	addr_QLocale_toCurrencyString10 uintptr
	addr_QLocale_toCurrencyString11 uintptr
	addr_QLocale_toCurrencyString12 uintptr
	addr_QLocale_toCurrencyString13 uintptr
	addr_QLocale_toCurrencyString14 uintptr
	pg_QLocale_toCurrencyString15 func(uintptr, float64, string, int64, *int64) uintptr
	pg_QLocale_toCurrencyString16 func(uintptr, float64, string, int64, int32, *int64) uintptr
	pg_QLocale_toCurrencyString17 func(uintptr, float32, string, int64, *int64) uintptr
	pg_QLocale_toCurrencyString18 func(uintptr, float32, string, int64, int32, *int64) uintptr
	addr_QLocale_formattedDataSize2 uintptr
	addr_QLocale_formattedDataSize3 uintptr
	addr_QLocale_uiLanguages2 uintptr
	addr_QLocale_languageToCode2 uintptr
	addr_QLocale_codeToLanguage2 uintptr
	addr_QLocale_quoteString3 uintptr
	addr_QLocale_quoteString4 uintptr
	addr_QLocale_OperatorEqual uintptr
	addr_QLocale_OperatorNotEqual uintptr
	pg_QLocale_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLocale_new, "QLocale_new")
	symaddr(&addr_QLocale_new2, "QLocale_new2")
	symaddr(&addr_QLocale_new3, "QLocale_new3")
	symaddr(&addr_QLocale_new4, "QLocale_new4")
	symaddr(&addr_QLocale_new5, "QLocale_new5")
	symaddr(&addr_QLocale_new6, "QLocale_new6")
	symaddr(&addr_QLocale_new7, "QLocale_new7")
	symaddr(&addr_QLocale_new8, "QLocale_new8")
	symaddr(&addr_QLocale_OperatorAssign, "QLocale_OperatorAssign")
	symaddr(&addr_QLocale_swap, "QLocale_swap")
	symaddr(&addr_QLocale_language, "QLocale_language")
	symaddr(&addr_QLocale_script, "QLocale_script")
	symaddr(&addr_QLocale_territory, "QLocale_territory")
	symaddr(&addr_QLocale_country, "QLocale_country")
	symaddr(&addr_QLocale_name, "QLocale_name")
	symaddr(&addr_QLocale_bcp47Name, "QLocale_bcp47Name")
	symaddr(&addr_QLocale_nativeLanguageName, "QLocale_nativeLanguageName")
	symaddr(&addr_QLocale_nativeTerritoryName, "QLocale_nativeTerritoryName")
	symaddr(&addr_QLocale_nativeCountryName, "QLocale_nativeCountryName")
	symaddr(&addr_QLocale_toShort, "QLocale_toShort")
	symaddr(&addr_QLocale_toUShort, "QLocale_toUShort")
	symaddr(&addr_QLocale_toInt, "QLocale_toInt")
	symaddr(&addr_QLocale_toUInt, "QLocale_toUInt")
	symaddr(&addr_QLocale_toLong, "QLocale_toLong")
	symaddr(&addr_QLocale_toULong, "QLocale_toULong")
	symaddr(&addr_QLocale_toLongLong, "QLocale_toLongLong")
	symaddr(&addr_QLocale_toULongLong, "QLocale_toULongLong")
	reg(&pg_QLocale_toFloat, "QLocale_toFloat")
	reg(&pg_QLocale_toDouble, "QLocale_toDouble")
	symaddr(&addr_QLocale_toShort2, "QLocale_toShort2")
	symaddr(&addr_QLocale_toUShort2, "QLocale_toUShort2")
	symaddr(&addr_QLocale_toInt2, "QLocale_toInt2")
	symaddr(&addr_QLocale_toUInt2, "QLocale_toUInt2")
	symaddr(&addr_QLocale_toLong2, "QLocale_toLong2")
	symaddr(&addr_QLocale_toULong2, "QLocale_toULong2")
	symaddr(&addr_QLocale_toLongLong2, "QLocale_toLongLong2")
	symaddr(&addr_QLocale_toULongLong2, "QLocale_toULongLong2")
	reg(&pg_QLocale_toFloat2, "QLocale_toFloat2")
	reg(&pg_QLocale_toDouble2, "QLocale_toDouble2")
	symaddr(&addr_QLocale_toString, "QLocale_toString")
	symaddr(&addr_QLocale_toString2, "QLocale_toString2")
	symaddr(&addr_QLocale_toString3, "QLocale_toString3")
	symaddr(&addr_QLocale_toString4, "QLocale_toString4")
	symaddr(&addr_QLocale_toString5, "QLocale_toString5")
	symaddr(&addr_QLocale_toString6, "QLocale_toString6")
	symaddr(&addr_QLocale_toString7, "QLocale_toString7")
	symaddr(&addr_QLocale_toString8, "QLocale_toString8")
	symaddr(&addr_QLocale_toString9, "QLocale_toString9")
	symaddr(&addr_QLocale_toString10, "QLocale_toString10")
	symaddr(&addr_QLocale_toString11, "QLocale_toString11")
	symaddr(&addr_QLocale_toString12, "QLocale_toString12")
	symaddr(&addr_QLocale_toString13, "QLocale_toString13")
	symaddr(&addr_QLocale_toString14, "QLocale_toString14")
	symaddr(&addr_QLocale_toString15, "QLocale_toString15")
	symaddr(&addr_QLocale_toString16, "QLocale_toString16")
	reg(&pg_QLocale_toString17, "QLocale_toString17")
	reg(&pg_QLocale_toString18, "QLocale_toString18")
	symaddr(&addr_QLocale_toString19, "QLocale_toString19")
	symaddr(&addr_QLocale_toString20, "QLocale_toString20")
	symaddr(&addr_QLocale_toString21, "QLocale_toString21")
	symaddr(&addr_QLocale_toString22, "QLocale_toString22")
	symaddr(&addr_QLocale_toString23, "QLocale_toString23")
	symaddr(&addr_QLocale_toString24, "QLocale_toString24")
	symaddr(&addr_QLocale_toString25, "QLocale_toString25")
	symaddr(&addr_QLocale_toString26, "QLocale_toString26")
	symaddr(&addr_QLocale_toString27, "QLocale_toString27")
	symaddr(&addr_QLocale_toString28, "QLocale_toString28")
	symaddr(&addr_QLocale_toString29, "QLocale_toString29")
	symaddr(&addr_QLocale_toString30, "QLocale_toString30")
	symaddr(&addr_QLocale_toString31, "QLocale_toString31")
	symaddr(&addr_QLocale_dateFormat, "QLocale_dateFormat")
	symaddr(&addr_QLocale_timeFormat, "QLocale_timeFormat")
	symaddr(&addr_QLocale_dateTimeFormat, "QLocale_dateTimeFormat")
	symaddr(&addr_QLocale_toTime, "QLocale_toTime")
	symaddr(&addr_QLocale_toTime2, "QLocale_toTime2")
	symaddr(&addr_QLocale_toDate, "QLocale_toDate")
	symaddr(&addr_QLocale_toDate2, "QLocale_toDate2")
	symaddr(&addr_QLocale_toDateTime, "QLocale_toDateTime")
	symaddr(&addr_QLocale_toDateTime2, "QLocale_toDateTime2")
	symaddr(&addr_QLocale_toDate3, "QLocale_toDate3")
	symaddr(&addr_QLocale_toDate4, "QLocale_toDate4")
	symaddr(&addr_QLocale_toDateTime3, "QLocale_toDateTime3")
	symaddr(&addr_QLocale_toDateTime4, "QLocale_toDateTime4")
	symaddr(&addr_QLocale_decimalPoint, "QLocale_decimalPoint")
	symaddr(&addr_QLocale_groupSeparator, "QLocale_groupSeparator")
	symaddr(&addr_QLocale_percent, "QLocale_percent")
	symaddr(&addr_QLocale_zeroDigit, "QLocale_zeroDigit")
	symaddr(&addr_QLocale_negativeSign, "QLocale_negativeSign")
	symaddr(&addr_QLocale_positiveSign, "QLocale_positiveSign")
	symaddr(&addr_QLocale_exponential, "QLocale_exponential")
	symaddr(&addr_QLocale_monthName, "QLocale_monthName")
	symaddr(&addr_QLocale_standaloneMonthName, "QLocale_standaloneMonthName")
	symaddr(&addr_QLocale_dayName, "QLocale_dayName")
	symaddr(&addr_QLocale_standaloneDayName, "QLocale_standaloneDayName")
	symaddr(&addr_QLocale_firstDayOfWeek, "QLocale_firstDayOfWeek")
	symaddr(&addr_QLocale_weekdays, "QLocale_weekdays")
	symaddr(&addr_QLocale_amText, "QLocale_amText")
	symaddr(&addr_QLocale_pmText, "QLocale_pmText")
	symaddr(&addr_QLocale_measurementSystem, "QLocale_measurementSystem")
	symaddr(&addr_QLocale_collation, "QLocale_collation")
	symaddr(&addr_QLocale_textDirection, "QLocale_textDirection")
	symaddr(&addr_QLocale_toUpper, "QLocale_toUpper")
	symaddr(&addr_QLocale_toLower, "QLocale_toLower")
	symaddr(&addr_QLocale_currencySymbol, "QLocale_currencySymbol")
	symaddr(&addr_QLocale_toCurrencyString, "QLocale_toCurrencyString")
	symaddr(&addr_QLocale_toCurrencyString2, "QLocale_toCurrencyString2")
	symaddr(&addr_QLocale_toCurrencyString3, "QLocale_toCurrencyString3")
	symaddr(&addr_QLocale_toCurrencyString4, "QLocale_toCurrencyString4")
	symaddr(&addr_QLocale_toCurrencyString5, "QLocale_toCurrencyString5")
	symaddr(&addr_QLocale_toCurrencyString6, "QLocale_toCurrencyString6")
	reg(&pg_QLocale_toCurrencyString7, "QLocale_toCurrencyString7")
	reg(&pg_QLocale_toCurrencyString8, "QLocale_toCurrencyString8")
	symaddr(&addr_QLocale_formattedDataSize, "QLocale_formattedDataSize")
	symaddr(&addr_QLocale_uiLanguages, "QLocale_uiLanguages")
	symaddr(&addr_QLocale_languageToCode, "QLocale_languageToCode")
	symaddr(&addr_QLocale_codeToLanguage, "QLocale_codeToLanguage")
	symaddr(&addr_QLocale_territoryToCode, "QLocale_territoryToCode")
	symaddr(&addr_QLocale_codeToTerritory, "QLocale_codeToTerritory")
	symaddr(&addr_QLocale_countryToCode, "QLocale_countryToCode")
	symaddr(&addr_QLocale_codeToCountry, "QLocale_codeToCountry")
	symaddr(&addr_QLocale_scriptToCode, "QLocale_scriptToCode")
	symaddr(&addr_QLocale_codeToScript, "QLocale_codeToScript")
	symaddr(&addr_QLocale_languageToString, "QLocale_languageToString")
	symaddr(&addr_QLocale_territoryToString, "QLocale_territoryToString")
	symaddr(&addr_QLocale_countryToString, "QLocale_countryToString")
	symaddr(&addr_QLocale_scriptToString, "QLocale_scriptToString")
	symaddr(&addr_QLocale_setDefault, "QLocale_setDefault")
	symaddr(&addr_QLocale_c, "QLocale_c")
	symaddr(&addr_QLocale_system, "QLocale_system")
	symaddr(&addr_QLocale_matchingLocales, "QLocale_matchingLocales")
	symaddr(&addr_QLocale_countriesForLanguage, "QLocale_countriesForLanguage")
	symaddr(&addr_QLocale_setNumberOptions, "QLocale_setNumberOptions")
	symaddr(&addr_QLocale_numberOptions, "QLocale_numberOptions")
	symaddr(&addr_QLocale_quoteString, "QLocale_quoteString")
	symaddr(&addr_QLocale_quoteString2, "QLocale_quoteString2")
	symaddr(&addr_QLocale_createSeparatedList, "QLocale_createSeparatedList")
	symaddr(&addr_QLocale_name2, "QLocale_name2")
	symaddr(&addr_QLocale_bcp47Name2, "QLocale_bcp47Name2")
	symaddr(&addr_QLocale_toShort3, "QLocale_toShort3")
	symaddr(&addr_QLocale_toUShort3, "QLocale_toUShort3")
	symaddr(&addr_QLocale_toInt3, "QLocale_toInt3")
	symaddr(&addr_QLocale_toUInt3, "QLocale_toUInt3")
	symaddr(&addr_QLocale_toLong3, "QLocale_toLong3")
	symaddr(&addr_QLocale_toULong3, "QLocale_toULong3")
	symaddr(&addr_QLocale_toLongLong3, "QLocale_toLongLong3")
	symaddr(&addr_QLocale_toULongLong3, "QLocale_toULongLong3")
	reg(&pg_QLocale_toFloat3, "QLocale_toFloat3")
	reg(&pg_QLocale_toDouble3, "QLocale_toDouble3")
	symaddr(&addr_QLocale_toShort4, "QLocale_toShort4")
	symaddr(&addr_QLocale_toUShort4, "QLocale_toUShort4")
	symaddr(&addr_QLocale_toInt4, "QLocale_toInt4")
	symaddr(&addr_QLocale_toUInt4, "QLocale_toUInt4")
	symaddr(&addr_QLocale_toLong4, "QLocale_toLong4")
	symaddr(&addr_QLocale_toULong4, "QLocale_toULong4")
	symaddr(&addr_QLocale_toLongLong4, "QLocale_toLongLong4")
	symaddr(&addr_QLocale_toULongLong4, "QLocale_toULongLong4")
	reg(&pg_QLocale_toFloat4, "QLocale_toFloat4")
	reg(&pg_QLocale_toDouble4, "QLocale_toDouble4")
	reg(&pg_QLocale_toString32, "QLocale_toString32")
	reg(&pg_QLocale_toString33, "QLocale_toString33")
	reg(&pg_QLocale_toString34, "QLocale_toString34")
	reg(&pg_QLocale_toString35, "QLocale_toString35")
	symaddr(&addr_QLocale_toString36, "QLocale_toString36")
	symaddr(&addr_QLocale_toString37, "QLocale_toString37")
	symaddr(&addr_QLocale_toString38, "QLocale_toString38")
	symaddr(&addr_QLocale_dateFormat2, "QLocale_dateFormat2")
	symaddr(&addr_QLocale_timeFormat2, "QLocale_timeFormat2")
	symaddr(&addr_QLocale_dateTimeFormat2, "QLocale_dateTimeFormat2")
	symaddr(&addr_QLocale_toTime3, "QLocale_toTime3")
	symaddr(&addr_QLocale_toDate5, "QLocale_toDate5")
	symaddr(&addr_QLocale_toDate6, "QLocale_toDate6")
	symaddr(&addr_QLocale_toDate7, "QLocale_toDate7")
	symaddr(&addr_QLocale_toDateTime5, "QLocale_toDateTime5")
	symaddr(&addr_QLocale_toDateTime6, "QLocale_toDateTime6")
	symaddr(&addr_QLocale_toDateTime7, "QLocale_toDateTime7")
	symaddr(&addr_QLocale_toDate8, "QLocale_toDate8")
	symaddr(&addr_QLocale_toDate9, "QLocale_toDate9")
	symaddr(&addr_QLocale_toDateTime8, "QLocale_toDateTime8")
	symaddr(&addr_QLocale_toDateTime9, "QLocale_toDateTime9")
	symaddr(&addr_QLocale_monthName2, "QLocale_monthName2")
	symaddr(&addr_QLocale_standaloneMonthName2, "QLocale_standaloneMonthName2")
	symaddr(&addr_QLocale_dayName2, "QLocale_dayName2")
	symaddr(&addr_QLocale_standaloneDayName2, "QLocale_standaloneDayName2")
	symaddr(&addr_QLocale_currencySymbol2, "QLocale_currencySymbol2")
	symaddr(&addr_QLocale_toCurrencyString9, "QLocale_toCurrencyString9")
	symaddr(&addr_QLocale_toCurrencyString10, "QLocale_toCurrencyString10")
	symaddr(&addr_QLocale_toCurrencyString11, "QLocale_toCurrencyString11")
	symaddr(&addr_QLocale_toCurrencyString12, "QLocale_toCurrencyString12")
	symaddr(&addr_QLocale_toCurrencyString13, "QLocale_toCurrencyString13")
	symaddr(&addr_QLocale_toCurrencyString14, "QLocale_toCurrencyString14")
	reg(&pg_QLocale_toCurrencyString15, "QLocale_toCurrencyString15")
	reg(&pg_QLocale_toCurrencyString16, "QLocale_toCurrencyString16")
	reg(&pg_QLocale_toCurrencyString17, "QLocale_toCurrencyString17")
	reg(&pg_QLocale_toCurrencyString18, "QLocale_toCurrencyString18")
	symaddr(&addr_QLocale_formattedDataSize2, "QLocale_formattedDataSize2")
	symaddr(&addr_QLocale_formattedDataSize3, "QLocale_formattedDataSize3")
	symaddr(&addr_QLocale_uiLanguages2, "QLocale_uiLanguages2")
	symaddr(&addr_QLocale_languageToCode2, "QLocale_languageToCode2")
	symaddr(&addr_QLocale_codeToLanguage2, "QLocale_codeToLanguage2")
	symaddr(&addr_QLocale_quoteString3, "QLocale_quoteString3")
	symaddr(&addr_QLocale_quoteString4, "QLocale_quoteString4")
	symaddr(&addr_QLocale_OperatorEqual, "QLocale_OperatorEqual")
	symaddr(&addr_QLocale_OperatorNotEqual, "QLocale_OperatorNotEqual")
	reg(&pg_QLocale_delete, "QLocale_delete")
}

type QLocale struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLocale) UnsafeGoID() int64     { return this.vid }
func (this *QLocale) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLocale(h uintptr) *QLocale {
	if h == 0 {
		return nil
	}
	return &QLocale{h: h}
}

func newownedQLocale(h uintptr) *QLocale {
	o := bindQLocale(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLocale).Delete)
	}
	return o
}

func (this *QLocale) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLocale) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLocale(h unsafe.Pointer) *QLocale { return bindQLocale(uintptr(h)) }
func UnsafeNewQLocaleOwned(h unsafe.Pointer) *QLocale { return newownedQLocale(uintptr(h)) }

func (this *QLocale) UnsafeRelease() { this.h = 0 }

func NewQLocale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new)
	__o := newownedQLocale(__r)
	return __o
}

func NewQLocale2(name string) *QLocale {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_new2, name_p, uintptr(len(name)))
	__o := newownedQLocale(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQLocale3(name *QStringView) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new3, uintptr(name.UnsafePointer()))
	__o := newownedQLocale(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQLocale4(language QLocale__Language, territory QLocale__Country) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new4, uintptr(language), uintptr(territory))
	__o := newownedQLocale(__r)
	return __o
}

func NewQLocale5(language QLocale__Language) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new5, uintptr(language))
	__o := newownedQLocale(__r)
	return __o
}

func NewQLocale6(other *QLocale) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new6, uintptr(other.UnsafePointer()))
	__o := newownedQLocale(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQLocale7(language QLocale__Language, script QLocale__Script) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new7, uintptr(language), uintptr(script))
	__o := newownedQLocale(__r)
	return __o
}

func NewQLocale8(language QLocale__Language, script QLocale__Script, territory QLocale__Country) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_new8, uintptr(language), uintptr(script), uintptr(territory))
	__o := newownedQLocale(__r)
	return __o
}

func (this *QLocale) OperatorAssign(other *QLocale) *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQLocale(__r)
}

func (this *QLocale) Swap(other *QLocale) {
	purego.SyscallN(addr_QLocale_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QLocale) Language() QLocale__Language {
	__r0, _, _ := purego.SyscallN(addr_QLocale_language, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__Language(__r)
}

func (this *QLocale) Script() QLocale__Script {
	__r0, _, _ := purego.SyscallN(addr_QLocale_script, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__Script(__r)
}

func (this *QLocale) Territory() QLocale__Country {
	__r0, _, _ := purego.SyscallN(addr_QLocale_territory, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__Country(__r)
}

func (this *QLocale) Country() QLocale__Country {
	__r0, _, _ := purego.SyscallN(addr_QLocale_country, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__Country(__r)
}

func (this *QLocale) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) Bcp47Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_bcp47Name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) NativeLanguageName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_nativeLanguageName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) NativeTerritoryName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_nativeTerritoryName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) NativeCountryName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_nativeCountryName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToShort(s string) int16 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toShort, this.h, s_p, uintptr(len(s)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToUShort(s string) uint16 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUShort, this.h, s_p, uintptr(len(s)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint16(__r)
}

func (this *QLocale) ToInt(s string) int {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toInt, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QLocale) ToUInt(s string) uint {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUInt, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint(__r)
}

func (this *QLocale) ToLong(s string) int32 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLong, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToULong(s string) uint32 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULong, this.h, s_p, uintptr(len(s)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint32(__r)
}

func (this *QLocale) ToLongLong(s string) int64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLongLong, this.h, s_p, uintptr(len(s)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToULongLong(s string) uint64 {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULongLong, this.h, s_p, uintptr(len(s)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint64(__r)
}

func (this *QLocale) ToFloat(s string) float32 {
	__r := pg_QLocale_toFloat(this.h, s, int64(len(s)))
	runtime.KeepAlive(this)
	return __r
}

func (this *QLocale) ToDouble(s string) float64 {
	__r := pg_QLocale_toDouble(this.h, s, int64(len(s)))
	runtime.KeepAlive(this)
	return __r
}

func (this *QLocale) ToShortWithQStringView(s *QStringView) int16 {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toShort2, this.h, uintptr(s.UnsafePointer()))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToUShortWithQStringView(s *QStringView) uint16 {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUShort2, this.h, uintptr(s.UnsafePointer()))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint16(__r)
}

func (this *QLocale) ToIntWithQStringView(s *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toInt2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r)
}

func (this *QLocale) ToUIntWithQStringView(s *QStringView) uint {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUInt2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint(__r)
}

func (this *QLocale) ToLongWithQStringView(s *QStringView) int32 {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLong2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToULongWithQStringView(s *QStringView) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULong2, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint32(__r)
}

func (this *QLocale) ToLongLongWithQStringView(s *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLongLong2, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToULongLongWithQStringView(s *QStringView) uint64 {
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULongLong2, this.h, uintptr(s.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint64(__r)
}

func (this *QLocale) ToFloatWithQStringView(s *QStringView) float32 {
	__r := pg_QLocale_toFloat2(this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToDoubleWithQStringView(s *QStringView) float64 {
	__r := pg_QLocale_toDouble2(this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r
}

func (this *QLocale) ToString(i int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWith(i uint64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString2, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithLong(i int32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString3, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString2(i uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString4, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithShort(i int16) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString5, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString3(i uint16) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString6, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithInt(i int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString7, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString4(i uint) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString8, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString5(number int64, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString9, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString6(number uint64, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString10, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString7(number int, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString11, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString8(number uint, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString12, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString9(number int32, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString13, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString10(number uint32, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString14, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString11(number int16, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString15, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString12(number uint16, fieldWidth int, fillChar uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString16, this.h, uintptr(number), uintptr(fieldWidth), uintptr(fillChar), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithDouble(f float64) string {
	var __n int64
	__r := pg_QLocale_toString17(this.h, f, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithFloat(f float32) string {
	var __n int64
	__r := pg_QLocale_toString18(this.h, f, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString13(date *QDate, format string) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString19, this.h, uintptr(date.UnsafePointer()), format_p, uintptr(len(format)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString14(time *QTime, format string) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString20, this.h, uintptr(time.UnsafePointer()), format_p, uintptr(len(format)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString15(dateTime *QDateTime, format string) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString21, this.h, uintptr(dateTime.UnsafePointer()), format_p, uintptr(len(format)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString16(date *QDate, format *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString22, this.h, uintptr(date.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString17(time *QTime, format *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString23, this.h, uintptr(time.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString18(dateTime *QDateTime, format *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString24, this.h, uintptr(dateTime.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithDate(date *QDate) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString25, this.h, uintptr(date.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithTime(time *QTime) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString26, this.h, uintptr(time.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
	return pgStr(__r, __n)
}

func (this *QLocale) ToStringWithDateTime(dateTime *QDateTime) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString27, this.h, uintptr(dateTime.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString19(date *QDate, format *QStringView, cal *QCalendar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString28, this.h, uintptr(date.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString20(date *QDate, format QLocale__FormatType, cal *QCalendar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString29, this.h, uintptr(date.UnsafePointer()), uintptr(format), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString21(dateTime *QDateTime, format QLocale__FormatType, cal *QCalendar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString30, this.h, uintptr(dateTime.UnsafePointer()), uintptr(format), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString22(dateTime *QDateTime, format *QStringView, cal *QCalendar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString31, this.h, uintptr(dateTime.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QLocale) DateFormat() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_dateFormat, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) TimeFormat() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_timeFormat, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) DateTimeFormat() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_dateTimeFormat, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToTime(string_ string) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toTime, this.h, string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQTime(__r)
}

func (this *QLocale) ToTime2(string_ string, format string) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toTime2, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQTime(__r)
}

func (this *QLocale) ToDate(string_ string) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate, this.h, string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func (this *QLocale) ToDate2(string_ string, format string) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate2, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func (this *QLocale) ToDateTime(string_ string) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime, this.h, string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDateTime2(string_ string, format string) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime2, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDate3(string_ string, format QLocale__FormatType, cal *QCalendar) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate3, this.h, string__p, uintptr(len(string_)), uintptr(format), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func (this *QLocale) ToDate4(string_ string, format string, cal *QCalendar) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate4, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func (this *QLocale) ToDateTime3(string_ string, format QLocale__FormatType, cal *QCalendar) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime3, this.h, string__p, uintptr(len(string_)), uintptr(format), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDateTime4(string_ string, format string, cal *QCalendar) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime4, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func (this *QLocale) DecimalPoint() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_decimalPoint, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) GroupSeparator() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_groupSeparator, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) Percent() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_percent, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ZeroDigit() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_zeroDigit, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) NegativeSign() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_negativeSign, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) PositiveSign() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_positiveSign, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) Exponential() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_exponential, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) MonthName(arg1 int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_monthName, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) StandaloneMonthName(arg1 int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_standaloneMonthName, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) DayName(arg1 int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_dayName, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) StandaloneDayName(arg1 int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_standaloneDayName, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) FirstDayOfWeek() DayOfWeek {
	__r0, _, _ := purego.SyscallN(addr_QLocale_firstDayOfWeek, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DayOfWeek(__r)
}

func (this *QLocale) Weekdays() []DayOfWeek {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_weekdays, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]DayOfWeek, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = DayOfWeek(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QLocale) AmText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_amText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) PmText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_pmText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) MeasurementSystem() QLocale__MeasurementSystem {
	__r0, _, _ := purego.SyscallN(addr_QLocale_measurementSystem, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__MeasurementSystem(__r)
}

func (this *QLocale) Collation() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_collation, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QLocale) TextDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QLocale_textDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QLocale) ToUpper(str string) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toUpper, this.h, str_p, uintptr(len(str)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func (this *QLocale) ToLower(str string) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toLower, this.h, str_p, uintptr(len(str)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func (this *QLocale) CurrencySymbol() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_currencySymbol, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString(arg1 int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyStringWith(arg1 uint64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString2, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyStringWithShort(i int16) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString3, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString2(i uint16) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString4, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyStringWithInt(i int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString5, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString3(i uint) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString6, this.h, uintptr(i), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyStringWithDouble(arg1 float64) string {
	var __n int64
	__r := pg_QLocale_toCurrencyString7(this.h, arg1, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyStringWithFloat(i float32) string {
	var __n int64
	__r := pg_QLocale_toCurrencyString8(this.h, i, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) FormattedDataSize(bytes int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_formattedDataSize, this.h, uintptr(bytes), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) UiLanguages() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_uiLanguages, this.h, uintptr(unsafe.Pointer(&__n)))
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

func QLocale_LanguageToCode(language QLocale__Language) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_languageToCode, uintptr(language), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_CodeToLanguage(languageCode *QStringView) QLocale__Language {
	__r0, _, _ := purego.SyscallN(addr_QLocale_codeToLanguage, uintptr(languageCode.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(languageCode)
	return QLocale__Language(__r)
}

func QLocale_TerritoryToCode(territory QLocale__Country) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_territoryToCode, uintptr(territory), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_CodeToTerritory(territoryCode *QStringView) QLocale__Country {
	__r0, _, _ := purego.SyscallN(addr_QLocale_codeToTerritory, uintptr(territoryCode.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(territoryCode)
	return QLocale__Country(__r)
}

func QLocale_CountryToCode(country QLocale__Country) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_countryToCode, uintptr(country), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_CodeToCountry(countryCode *QStringView) QLocale__Country {
	__r0, _, _ := purego.SyscallN(addr_QLocale_codeToCountry, uintptr(countryCode.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(countryCode)
	return QLocale__Country(__r)
}

func QLocale_ScriptToCode(script QLocale__Script) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_scriptToCode, uintptr(script), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_CodeToScript(scriptCode *QStringView) QLocale__Script {
	__r0, _, _ := purego.SyscallN(addr_QLocale_codeToScript, uintptr(scriptCode.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(scriptCode)
	return QLocale__Script(__r)
}

func QLocale_LanguageToString(language QLocale__Language) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_languageToString, uintptr(language), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_TerritoryToString(territory QLocale__Country) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_territoryToString, uintptr(territory), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_CountryToString(country QLocale__Country) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_countryToString, uintptr(country), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_ScriptToString(script QLocale__Script) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_scriptToString, uintptr(script), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_SetDefault(locale *QLocale) {
	purego.SyscallN(addr_QLocale_setDefault, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(locale)
}

func QLocale_C() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_c)
	return newownedQLocale(__r)
}

func QLocale_System() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QLocale_system)
	return newownedQLocale(__r)
}

func QLocale_MatchingLocales(language QLocale__Language, script QLocale__Script, territory QLocale__Country) []*QLocale {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_matchingLocales, uintptr(language), uintptr(script), uintptr(territory), uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QLocale, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQLocale(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QLocale_CountriesForLanguage(lang QLocale__Language) []QLocale__Country {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_countriesForLanguage, uintptr(lang), uintptr(unsafe.Pointer(&__n)))
	__out := make([]QLocale__Country, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = QLocale__Country(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QLocale) SetNumberOptions(options QLocale__NumberOption) {
	purego.SyscallN(addr_QLocale_setNumberOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QLocale) NumberOptions() QLocale__NumberOption {
	__r0, _, _ := purego.SyscallN(addr_QLocale_numberOptions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__NumberOption(__r)
}

func (this *QLocale) QuoteString(str string) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_quoteString, this.h, str_p, uintptr(len(str)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func (this *QLocale) QuoteStringWithStr(str *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_quoteString2, this.h, uintptr(str.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func (this *QLocale) CreateSeparatedList(strl []string) string {
	strl_arr := make([]uintptr, len(strl))
	for __i, __e := range strl {
		strl_arr[__i] = pgCString(__e)
	}
	var strl_p uintptr
	if len(strl_arr) > 0 {
		strl_p = uintptr(unsafe.Pointer(&strl_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_createSeparatedList, this.h, strl_p, uintptr(len(strl)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range strl_arr {
		pgFree(strl_arr[__i])
	}
	runtime.KeepAlive(strl_arr)
	return pgStr(__r, __n)
}

func (this *QLocale) NameWithSeparator(separator QLocale__TagSeparator) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_name2, this.h, uintptr(separator), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) Bcp47NameWithSeparator(separator QLocale__TagSeparator) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_bcp47Name2, this.h, uintptr(separator), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToShort2(s string) (int16, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toShort3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToUShort2(s string) (uint16, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUShort3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint16(__r), ok_out != 0
}

func (this *QLocale) ToInt2(s string) (int, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toInt3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r), ok_out != 0
}

func (this *QLocale) ToUInt2(s string) (uint, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUInt3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint(__r), ok_out != 0
}

func (this *QLocale) ToLong2(s string) (int32, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLong3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToULong2(s string) (uint32, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULong3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint32(__r), ok_out != 0
}

func (this *QLocale) ToLongLong2(s string) (int64, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLongLong3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToULongLong2(s string) (uint64, bool) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULongLong3, this.h, s_p, uintptr(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint64(__r), ok_out != 0
}

func (this *QLocale) ToFloat2(s string) (float32, bool) {
	var ok_out int32
	__r := pg_QLocale_toFloat3(this.h, s, int64(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLocale) ToDouble2(s string) (float64, bool) {
	var ok_out int32
	__r := pg_QLocale_toDouble3(this.h, s, int64(len(s)), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QLocale) ToShort3(s *QStringView) (int16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toShort4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToUShort3(s *QStringView) (uint16, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUShort4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint16(__r), ok_out != 0
}

func (this *QLocale) ToInt3(s *QStringView) (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toInt4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return int(__r), ok_out != 0
}

func (this *QLocale) ToUInt3(s *QStringView) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toUInt4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint(__r), ok_out != 0
}

func (this *QLocale) ToLong3(s *QStringView) (int32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLong4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToULong3(s *QStringView) (uint32, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULong4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint32(__r), ok_out != 0
}

func (this *QLocale) ToLongLong3(s *QStringView) (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toLongLong4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToULongLong3(s *QStringView) (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QLocale_toULongLong4, this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return uint64(__r), ok_out != 0
}

func (this *QLocale) ToFloat3(s *QStringView) (float32, bool) {
	var ok_out int32
	__r := pg_QLocale_toFloat4(this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToDouble3(s *QStringView) (float64, bool) {
	var ok_out int32
	__r := pg_QLocale_toDouble4(this.h, uintptr(s.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r, ok_out != 0
}

func (this *QLocale) ToString23(f float64, format int8) string {
	var __n int64
	__r := pg_QLocale_toString32(this.h, f, format, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString24(f float64, format int8, precision int) string {
	var __n int64
	__r := pg_QLocale_toString33(this.h, f, format, int32(precision), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString25(f float32, format int8) string {
	var __n int64
	__r := pg_QLocale_toString34(this.h, f, format, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString26(f float32, format int8, precision int) string {
	var __n int64
	__r := pg_QLocale_toString35(this.h, f, format, int32(precision), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString27(date *QDate, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString36, this.h, uintptr(date.UnsafePointer()), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString28(time *QTime, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString37, this.h, uintptr(time.UnsafePointer()), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
	return pgStr(__r, __n)
}

func (this *QLocale) ToString29(dateTime *QDateTime, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toString38, this.h, uintptr(dateTime.UnsafePointer()), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
	return pgStr(__r, __n)
}

func (this *QLocale) DateFormatWithFormat(format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_dateFormat2, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) TimeFormatWithFormat(format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_timeFormat2, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) DateTimeFormatWithFormat(format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_dateTimeFormat2, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToTime3(string_ string, arg2 QLocale__FormatType) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toTime3, this.h, string__p, uintptr(len(string_)), uintptr(arg2))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQTime(__r)
}

func (this *QLocale) ToDate5(string_ string, arg2 QLocale__FormatType) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate5, this.h, string__p, uintptr(len(string_)), uintptr(arg2))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func (this *QLocale) ToDate6(string_ string, arg2 QLocale__FormatType, baseYear int) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate6, this.h, string__p, uintptr(len(string_)), uintptr(arg2), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func (this *QLocale) ToDate7(string_ string, format string, baseYear int) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate7, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func (this *QLocale) ToDateTime5(string_ string, format QLocale__FormatType) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime5, this.h, string__p, uintptr(len(string_)), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDateTime6(string_ string, format QLocale__FormatType, baseYear int) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime6, this.h, string__p, uintptr(len(string_)), uintptr(format), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDateTime7(string_ string, format string, baseYear int) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime7, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDate8(string_ string, format QLocale__FormatType, cal *QCalendar, baseYear int) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate8, this.h, string__p, uintptr(len(string_)), uintptr(format), uintptr(cal.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func (this *QLocale) ToDate9(string_ string, format string, cal *QCalendar, baseYear int) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDate9, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func (this *QLocale) ToDateTime8(string_ string, format QLocale__FormatType, cal *QCalendar, baseYear int) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime8, this.h, string__p, uintptr(len(string_)), uintptr(format), uintptr(cal.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func (this *QLocale) ToDateTime9(string_ string, format string, cal *QCalendar, baseYear int) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QLocale_toDateTime9, this.h, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func (this *QLocale) MonthName2(arg1 int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_monthName2, this.h, uintptr(arg1), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) StandaloneMonthName2(arg1 int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_standaloneMonthName2, this.h, uintptr(arg1), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) DayName2(arg1 int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_dayName2, this.h, uintptr(arg1), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) StandaloneDayName2(arg1 int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_standaloneDayName2, this.h, uintptr(arg1), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) CurrencySymbolWithQLocaleCurrencySymbolFormat(arg1 QLocale__CurrencySymbolFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_currencySymbol2, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString4(arg1 int64, symbol string) string {
	var symbol_p uintptr
	if len(symbol) > 0 {
		symbol_p = uintptr(unsafe.Pointer(unsafe.StringData(symbol)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString9, this.h, uintptr(arg1), symbol_p, uintptr(len(symbol)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(symbol)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString5(arg1 uint64, symbol string) string {
	var symbol_p uintptr
	if len(symbol) > 0 {
		symbol_p = uintptr(unsafe.Pointer(unsafe.StringData(symbol)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString10, this.h, uintptr(arg1), symbol_p, uintptr(len(symbol)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(symbol)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString6(i int16, symbol string) string {
	var symbol_p uintptr
	if len(symbol) > 0 {
		symbol_p = uintptr(unsafe.Pointer(unsafe.StringData(symbol)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString11, this.h, uintptr(i), symbol_p, uintptr(len(symbol)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(symbol)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString7(i uint16, symbol string) string {
	var symbol_p uintptr
	if len(symbol) > 0 {
		symbol_p = uintptr(unsafe.Pointer(unsafe.StringData(symbol)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString12, this.h, uintptr(i), symbol_p, uintptr(len(symbol)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(symbol)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString8(i int, symbol string) string {
	var symbol_p uintptr
	if len(symbol) > 0 {
		symbol_p = uintptr(unsafe.Pointer(unsafe.StringData(symbol)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString13, this.h, uintptr(i), symbol_p, uintptr(len(symbol)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(symbol)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString9(i uint, symbol string) string {
	var symbol_p uintptr
	if len(symbol) > 0 {
		symbol_p = uintptr(unsafe.Pointer(unsafe.StringData(symbol)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_toCurrencyString14, this.h, uintptr(i), symbol_p, uintptr(len(symbol)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(symbol)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString10(arg1 float64, symbol string) string {
	var __n int64
	__r := pg_QLocale_toCurrencyString15(this.h, arg1, symbol, int64(len(symbol)), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString11(arg1 float64, symbol string, precision int) string {
	var __n int64
	__r := pg_QLocale_toCurrencyString16(this.h, arg1, symbol, int64(len(symbol)), int32(precision), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString12(i float32, symbol string) string {
	var __n int64
	__r := pg_QLocale_toCurrencyString17(this.h, i, symbol, int64(len(symbol)), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) ToCurrencyString13(i float32, symbol string, precision int) string {
	var __n int64
	__r := pg_QLocale_toCurrencyString18(this.h, i, symbol, int64(len(symbol)), int32(precision), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) FormattedDataSize2(bytes int64, precision int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_formattedDataSize2, this.h, uintptr(bytes), uintptr(precision), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) FormattedDataSize3(bytes int64, precision int, format QLocale__DataSizeFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_formattedDataSize3, this.h, uintptr(bytes), uintptr(precision), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLocale) UiLanguagesWithSeparator(separator QLocale__TagSeparator) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_uiLanguages2, this.h, uintptr(separator), uintptr(unsafe.Pointer(&__n)))
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

func QLocale_LanguageToCode2(language QLocale__Language, codeTypes QLocale__LanguageCodeType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_languageToCode2, uintptr(language), uintptr(codeTypes), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLocale_CodeToLanguage2(languageCode *QStringView, codeTypes QLocale__LanguageCodeType) QLocale__Language {
	__r0, _, _ := purego.SyscallN(addr_QLocale_codeToLanguage2, uintptr(languageCode.UnsafePointer()), uintptr(codeTypes))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(languageCode)
	return QLocale__Language(__r)
}

func (this *QLocale) QuoteString2(str string, style QLocale__QuotationStyle) string {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_quoteString3, this.h, str_p, uintptr(len(str)), uintptr(style), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func (this *QLocale) QuoteString3(str *QStringView, style QLocale__QuotationStyle) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLocale_quoteString4, this.h, uintptr(str.UnsafePointer()), uintptr(style), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return pgStr(__r, __n)
}

func (this *QLocale) OperatorEqual(rhs *QLocale) bool {
	__r0, _, _ := purego.SyscallN(addr_QLocale_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLocale) OperatorNotEqual(rhs *QLocale) bool {
	__r0, _, _ := purego.SyscallN(addr_QLocale_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLocale) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLocale_delete(this.h)
	}
	this.UnsafeRelease()
}
