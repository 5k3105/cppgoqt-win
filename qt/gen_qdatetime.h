#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDateTime_new();
void* QDateTime_new2(void* date, void* time, int spec);
void* QDateTime_new3(void* date, void* time, void* timeZone);
void* QDateTime_new4(void* date, void* time);
void* QDateTime_new5(void* other);
void* QDateTime_new6(void* date, void* time, int spec, int offsetSeconds);
void* QDateTime_new7(void* date, void* time, void* timeZone, int resolve);
void* QDateTime_new8(void* date, void* time, int resolve);
void* QDateTime_OperatorAssign(void* self, void* other);
void QDateTime_swap(void* self, void* other);
int QDateTime_isNull(void* self);
int QDateTime_isValid(void* self);
void* QDateTime_date(void* self);
void* QDateTime_time(void* self);
int QDateTime_timeSpec(void* self);
int QDateTime_offsetFromUtc(void* self);
void* QDateTime_timeRepresentation(void* self);
void* QDateTime_timeZone(void* self);
char* QDateTime_timeZoneAbbreviation(void* self, long long* __out_len);
int QDateTime_isDaylightTime(void* self);
long long QDateTime_toMSecsSinceEpoch(void* self);
long long QDateTime_toSecsSinceEpoch(void* self);
void QDateTime_setDate(void* self, void* date);
void QDateTime_setTime(void* self, void* time);
void QDateTime_setTimeSpec(void* self, int spec);
void QDateTime_setOffsetFromUtc(void* self, int offsetSeconds);
void QDateTime_setTimeZone(void* self, void* toZone);
void QDateTime_setMSecsSinceEpoch(void* self, long long msecs);
void QDateTime_setSecsSinceEpoch(void* self, long long secs);
char* QDateTime_toString(void* self, long long* __out_len);
char* QDateTime_toString2(void* self, const char* format, long long format_len, long long* __out_len);
char* QDateTime_toString3(void* self, const char* format, long long format_len, void* cal, long long* __out_len);
char* QDateTime_toString4(void* self, void* format, long long* __out_len);
char* QDateTime_toString5(void* self, void* format, void* cal, long long* __out_len);
void* QDateTime_addDays(void* self, long long days);
void* QDateTime_addMonths(void* self, int months);
void* QDateTime_addYears(void* self, int years);
void* QDateTime_addSecs(void* self, long long secs);
void* QDateTime_addMSecs(void* self, long long msecs);
void* QDateTime_addDuration(void* self, long long msecs);
void* QDateTime_toTimeSpec(void* self, int spec);
void* QDateTime_toLocalTime(void* self);
void* QDateTime_toUTC(void* self);
void* QDateTime_toOffsetFromUtc(void* self, int offsetSeconds);
void* QDateTime_toTimeZone(void* self, void* toZone);
long long QDateTime_daysTo(void* self, void* arg1);
long long QDateTime_secsTo(void* self, void* arg1);
long long QDateTime_msecsTo(void* self, void* arg1);
void* QDateTime_currentDateTime(void* zone);
void* QDateTime_currentDateTime2();
void* QDateTime_currentDateTimeUtc();
void* QDateTime_fromString(void* string_);
void* QDateTime_fromString2(const char* string_, long long string__len);
void* QDateTime_fromString3(void* string_, void* format, void* cal);
void* QDateTime_fromString4(const char* string_, long long string__len, void* format, void* cal);
void* QDateTime_fromString5(const char* string_, long long string__len, const char* format, long long format_len, void* cal);
void* QDateTime_fromString6(void* string_, void* format);
void* QDateTime_fromString7(void* string_, void* format, int baseYear, void* cal);
void* QDateTime_fromString8(const char* string_, long long string__len, void* format);
void* QDateTime_fromString9(const char* string_, long long string__len, void* format, int baseYear, void* cal);
void* QDateTime_fromString10(const char* string_, long long string__len, const char* format, long long format_len);
void* QDateTime_fromString11(const char* string_, long long string__len, const char* format, long long format_len, int baseYear, void* cal);
void* QDateTime_fromMSecsSinceEpoch(long long msecs, int spec);
void* QDateTime_fromSecsSinceEpoch(long long secs, int spec);
void* QDateTime_fromMSecsSinceEpoch2(long long msecs, void* timeZone);
void* QDateTime_fromSecsSinceEpoch2(long long secs, void* timeZone);
void* QDateTime_fromMSecsSinceEpoch3(long long msecs);
void* QDateTime_fromSecsSinceEpoch3(long long secs);
long long QDateTime_currentMSecsSinceEpoch();
long long QDateTime_currentSecsSinceEpoch();
void* QDateTime_OperatorPlusAssign(void* self, long long duration);
void* QDateTime_OperatorMinusAssign(void* self, long long duration);
void QDateTime_setDate2(void* self, void* date, int resolve);
void QDateTime_setTime2(void* self, void* time, int resolve);
void QDateTime_setTimeZone2(void* self, void* toZone, int resolve);
char* QDateTime_toString6(void* self, int format, long long* __out_len);
void* QDateTime_fromString12(void* string_, int format);
void* QDateTime_fromString13(const char* string_, long long string__len, int format);
void* QDateTime_fromString14(void* string_, void* format, int baseYear);
void* QDateTime_fromString15(const char* string_, long long string__len, void* format, int baseYear);
void* QDateTime_fromString16(const char* string_, long long string__len, const char* format, long long format_len, int baseYear);
void* QDateTime_fromMSecsSinceEpoch4(long long msecs, int spec, int offsetFromUtc);
void* QDateTime_fromSecsSinceEpoch4(long long secs, int spec, int offsetFromUtc);
long long QDateTime_OperatorMinus(void* self, void* rhs);
void* QDateTime_OperatorPlus(void* self, long long duration);
void* QDateTime_OperatorMinus2(void* self, long long duration);
int QDateTime_OperatorEqual(void* self, void* rhs);
int QDateTime_OperatorNotEqual(void* self, void* rhs);
int QDateTime_OperatorLesser(void* self, void* rhs);
int QDateTime_OperatorGreater(void* self, void* rhs);
int QDateTime_OperatorLesserOrEqual(void* self, void* rhs);
int QDateTime_OperatorGreaterOrEqual(void* self, void* rhs);
void QDateTime_delete(void* self);

#ifdef __cplusplus
}
#endif
