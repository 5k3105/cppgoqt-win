#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDate_new();
void* QDate_new2(int y, int m, int d);
void* QDate_new3(int y, int m, int d, void* cal);
void* QDate_new4(void* arg1);
int QDate_isNull(void* self);
int QDate_isValid(void* self);
int QDate_year(void* self);
int QDate_month(void* self);
int QDate_day(void* self);
int QDate_dayOfWeek(void* self);
int QDate_dayOfYear(void* self);
int QDate_daysInMonth(void* self);
int QDate_daysInYear(void* self);
int QDate_weekNumber(void* self);
int QDate_year2(void* self, void* cal);
int QDate_month2(void* self, void* cal);
int QDate_day2(void* self, void* cal);
int QDate_dayOfWeek2(void* self, void* cal);
int QDate_dayOfYear2(void* self, void* cal);
int QDate_daysInMonth2(void* self, void* cal);
int QDate_daysInYear2(void* self, void* cal);
void* QDate_startOfDay(void* self, int spec);
void* QDate_endOfDay(void* self, int spec);
void* QDate_startOfDay2(void* self, void* zone);
void* QDate_endOfDay2(void* self, void* zone);
void* QDate_startOfDay3(void* self);
void* QDate_endOfDay3(void* self);
char* QDate_toString(void* self, long long* __out_len);
char* QDate_toString2(void* self, const char* format, long long format_len, long long* __out_len);
char* QDate_toString3(void* self, const char* format, long long format_len, void* cal, long long* __out_len);
char* QDate_toString4(void* self, void* format, long long* __out_len);
char* QDate_toString5(void* self, void* format, void* cal, long long* __out_len);
int QDate_setDate(void* self, int year, int month, int day);
int QDate_setDate2(void* self, int year, int month, int day, void* cal);
void QDate_getDate(void* self, int* year, int* month, int* day);
void* QDate_addDays(void* self, long long days);
void* QDate_addMonths(void* self, int months);
void* QDate_addYears(void* self, int years);
void* QDate_addMonths2(void* self, int months, void* cal);
void* QDate_addYears2(void* self, int years, void* cal);
long long QDate_daysTo(void* self, void* d);
void* QDate_currentDate();
void* QDate_fromString(void* string_);
void* QDate_fromString2(const char* string_, long long string__len);
void* QDate_fromString3(void* string_, void* format, void* cal);
void* QDate_fromString4(const char* string_, long long string__len, void* format, void* cal);
void* QDate_fromString5(const char* string_, long long string__len, const char* format, long long format_len, void* cal);
void* QDate_fromString6(void* string_, void* format);
void* QDate_fromString7(void* string_, void* format, int baseYear, void* cal);
void* QDate_fromString8(const char* string_, long long string__len, void* format);
void* QDate_fromString9(const char* string_, long long string__len, void* format, int baseYear, void* cal);
void* QDate_fromString10(const char* string_, long long string__len, const char* format, long long format_len);
void* QDate_fromString11(const char* string_, long long string__len, const char* format, long long format_len, int baseYear, void* cal);
int QDate_isValid2(int y, int m, int d);
int QDate_isLeapYear(int year);
void* QDate_fromJulianDay(long long jd_);
long long QDate_toJulianDay(void* self);
int QDate_weekNumber2(void* self, int* yearNum);
void* QDate_startOfDay4(void* self, int spec, int offsetSeconds);
void* QDate_endOfDay4(void* self, int spec, int offsetSeconds);
char* QDate_toString6(void* self, int format, long long* __out_len);
void* QDate_fromString12(void* string_, int format);
void* QDate_fromString13(const char* string_, long long string__len, int format);
void* QDate_fromString14(void* string_, void* format, int baseYear);
void* QDate_fromString15(const char* string_, long long string__len, void* format, int baseYear);
void* QDate_fromString16(const char* string_, long long string__len, const char* format, long long format_len, int baseYear);
int QDate_OperatorEqual(void* self, void* rhs);
int QDate_OperatorNotEqual(void* self, void* rhs);
int QDate_OperatorLesser(void* self, void* rhs);
int QDate_OperatorGreater(void* self, void* rhs);
int QDate_OperatorLesserOrEqual(void* self, void* rhs);
int QDate_OperatorGreaterOrEqual(void* self, void* rhs);
void* QDate_OperatorPlusPlus(void* self);
void* QDate_OperatorMinusMinus(void* self);
void* QDate_OperatorPlusPlus2(void* self, int arg1);
void* QDate_OperatorMinusMinus2(void* self, int arg1);
void QDate_delete(void* self);

#ifdef __cplusplus
}
#endif
