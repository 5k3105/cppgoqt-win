#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTimeZone__OffsetData_new(void* arg1);
void* QTimeZone__OffsetData_new2();
void* QTimeZone__OffsetData_OperatorAssign(void* self, void* arg1);
char* QTimeZone__OffsetData_abbreviation_get(void* self, long long* __out_len);
void QTimeZone__OffsetData_abbreviation_set(void* self, const char* abbreviation, long long abbreviation_len);
void* QTimeZone__OffsetData_atUtc_get(void* self);
void QTimeZone__OffsetData_atUtc_set(void* self, void* atUtc);
int QTimeZone__OffsetData_offsetFromUtc_get(void* self);
void QTimeZone__OffsetData_offsetFromUtc_set(void* self, int offsetFromUtc);
int QTimeZone__OffsetData_standardTimeOffset_get(void* self);
void QTimeZone__OffsetData_standardTimeOffset_set(void* self, int standardTimeOffset);
int QTimeZone__OffsetData_daylightTimeOffset_get(void* self);
void QTimeZone__OffsetData_daylightTimeOffset_set(void* self, int daylightTimeOffset);
void QTimeZone__OffsetData_delete(void* self);

#ifdef __cplusplus
}
#endif
