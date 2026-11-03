#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCalendar_new();
void* QCalendar_new2(int system);
void* QCalendar_new3(const char* name, long long name_len);
void* QCalendar_new4(void* id);
int QCalendar_isValid(void* self);
int QCalendar_daysInMonth(void* self, int month);
int QCalendar_daysInYear(void* self, int year);
int QCalendar_monthsInYear(void* self, int year);
int QCalendar_isDateValid(void* self, int year, int month, int day);
int QCalendar_isLeapYear(void* self, int year);
int QCalendar_isGregorian(void* self);
int QCalendar_isLunar(void* self);
int QCalendar_isLuniSolar(void* self);
int QCalendar_isSolar(void* self);
int QCalendar_isProleptic(void* self);
int QCalendar_hasYearZero(void* self);
int QCalendar_maximumDaysInMonth(void* self);
int QCalendar_minimumDaysInMonth(void* self);
int QCalendar_maximumMonthsInYear(void* self);
char* QCalendar_name(void* self, long long* __out_len);
void* QCalendar_dateFromParts(void* self, int year, int month, int day);
void* QCalendar_dateFromParts2(void* self, void* parts);
void* QCalendar_matchCenturyToWeekday(void* self, void* parts, int dow);
void* QCalendar_partsFromDate(void* self, void* date);
int QCalendar_dayOfWeek(void* self, void* date);
char* QCalendar_monthName(void* self, void* locale, int month, long long* __out_len);
char* QCalendar_standaloneMonthName(void* self, void* locale, int month, long long* __out_len);
char* QCalendar_weekDayName(void* self, void* locale, int day, long long* __out_len);
char* QCalendar_standaloneWeekDayName(void* self, void* locale, int day, long long* __out_len);
char* QCalendar_dateTimeToString(void* self, void* format, void* datetime, void* dateOnly, void* timeOnly, void* locale, long long* __out_len);
char** QCalendar_availableCalendars(long long* __out_len);
int QCalendar_daysInMonth2(void* self, int month, int year);
char* QCalendar_monthName2(void* self, void* locale, int month, int year, long long* __out_len);
char* QCalendar_monthName3(void* self, void* locale, int month, int year, int format, long long* __out_len);
char* QCalendar_standaloneMonthName2(void* self, void* locale, int month, int year, long long* __out_len);
char* QCalendar_standaloneMonthName3(void* self, void* locale, int month, int year, int format, long long* __out_len);
char* QCalendar_weekDayName2(void* self, void* locale, int day, int format, long long* __out_len);
char* QCalendar_standaloneWeekDayName2(void* self, void* locale, int day, int format, long long* __out_len);
void QCalendar_delete(void* self);

#ifdef __cplusplus
}
#endif
