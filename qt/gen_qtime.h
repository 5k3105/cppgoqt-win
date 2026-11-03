#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTime_new();
void* QTime_new2(int h, int m);
void* QTime_new3(void* arg1);
void* QTime_new4(int h, int m, int s);
void* QTime_new5(int h, int m, int s, int ms);
int QTime_isNull(void* self);
int QTime_isValid(void* self);
int QTime_hour(void* self);
int QTime_minute(void* self);
int QTime_second(void* self);
int QTime_msec(void* self);
char* QTime_toString(void* self, long long* __out_len);
char* QTime_toString2(void* self, const char* format, long long format_len, long long* __out_len);
char* QTime_toString3(void* self, void* format, long long* __out_len);
int QTime_setHMS(void* self, int h, int m, int s);
void* QTime_addSecs(void* self, int secs);
int QTime_secsTo(void* self, void* t);
void* QTime_addMSecs(void* self, int ms);
int QTime_msecsTo(void* self, void* t);
void* QTime_fromMSecsSinceStartOfDay(int msecs);
int QTime_msecsSinceStartOfDay(void* self);
void* QTime_currentTime();
void* QTime_fromString(void* string_);
void* QTime_fromString2(void* string_, void* format);
void* QTime_fromString3(const char* string_, long long string__len, void* format);
void* QTime_fromString4(const char* string_, long long string__len);
void* QTime_fromString5(const char* string_, long long string__len, const char* format, long long format_len);
int QTime_isValid2(int h, int m, int s);
char* QTime_toString4(void* self, int f, long long* __out_len);
int QTime_setHMS2(void* self, int h, int m, int s, int ms);
void* QTime_fromString6(void* string_, int format);
void* QTime_fromString7(const char* string_, long long string__len, int format);
int QTime_isValid3(int h, int m, int s, int ms);
int QTime_OperatorEqual(void* self, void* rhs);
int QTime_OperatorNotEqual(void* self, void* rhs);
int QTime_OperatorLesser(void* self, void* rhs);
int QTime_OperatorGreater(void* self, void* rhs);
int QTime_OperatorLesserOrEqual(void* self, void* rhs);
int QTime_OperatorGreaterOrEqual(void* self, void* rhs);
void QTime_delete(void* self);

#ifdef __cplusplus
}
#endif
