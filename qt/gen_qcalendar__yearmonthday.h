#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCalendar__YearMonthDay_new();
void* QCalendar__YearMonthDay_new2(int y);
void* QCalendar__YearMonthDay_new3(int y, int m);
void* QCalendar__YearMonthDay_new4(int y, int m, int d);
int QCalendar__YearMonthDay_isValid(void* self);
int QCalendar__YearMonthDay_year_get(void* self);
void QCalendar__YearMonthDay_year_set(void* self, int year);
int QCalendar__YearMonthDay_month_get(void* self);
void QCalendar__YearMonthDay_month_set(void* self, int month);
int QCalendar__YearMonthDay_day_get(void* self);
void QCalendar__YearMonthDay_day_set(void* self, int day);
void QCalendar__YearMonthDay_delete(void* self);

#ifdef __cplusplus
}
#endif
