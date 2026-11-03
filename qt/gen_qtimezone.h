#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTimeZone_new();
void* QTimeZone_new2(int spec);
void* QTimeZone_new3(int offsetSeconds);
void* QTimeZone_new4(void* ianaId);
void* QTimeZone_new5(void* zoneId, int offsetSeconds, const char* name, long long name_len, const char* abbreviation, long long abbreviation_len);
void* QTimeZone_new6(void* other);
void* QTimeZone_new7(void* zoneId, int offsetSeconds, const char* name, long long name_len, const char* abbreviation, long long abbreviation_len, int territory);
void* QTimeZone_new8(void* zoneId, int offsetSeconds, const char* name, long long name_len, const char* abbreviation, long long abbreviation_len, int territory, const char* comment, long long comment_len);
void* QTimeZone_OperatorAssign(void* self, void* other);
void QTimeZone_swap(void* self, void* other);
int QTimeZone_isValid(void* self);
void* QTimeZone_fromDurationAheadOfUtc(long long offset);
void* QTimeZone_fromSecondsAheadOfUtc(int offset);
int QTimeZone_timeSpec(void* self);
int QTimeZone_fixedSecondsAheadOfUtc(void* self);
int QTimeZone_isUtcOrFixedOffset(int spec);
int QTimeZone_isUtcOrFixedOffset2(void* self);
void* QTimeZone_asBackendZone(void* self);
int QTimeZone_hasAlternativeName(void* self, void* alias);
void* QTimeZone_id(void* self);
int QTimeZone_territory(void* self);
int QTimeZone_country(void* self);
char* QTimeZone_comment(void* self, long long* __out_len);
char* QTimeZone_displayName(void* self, void* atDateTime, long long* __out_len);
char* QTimeZone_displayName2(void* self, int timeType, long long* __out_len);
char* QTimeZone_abbreviation(void* self, void* atDateTime, long long* __out_len);
int QTimeZone_offsetFromUtc(void* self, void* atDateTime);
int QTimeZone_standardTimeOffset(void* self, void* atDateTime);
int QTimeZone_daylightTimeOffset(void* self, void* atDateTime);
int QTimeZone_hasDaylightTime(void* self);
int QTimeZone_isDaylightTime(void* self, void* atDateTime);
void* QTimeZone_offsetData(void* self, void* forDateTime);
int QTimeZone_hasTransitions(void* self);
void* QTimeZone_nextTransition(void* self, void* afterDateTime);
void* QTimeZone_previousTransition(void* self, void* beforeDateTime);
void** QTimeZone_transitions(void* self, void* fromDateTime, void* toDateTime, long long* __out_len);
void* QTimeZone_systemTimeZoneId();
void* QTimeZone_systemTimeZone();
void* QTimeZone_utc();
int QTimeZone_isTimeZoneIdAvailable(void* ianaId);
void** QTimeZone_availableTimeZoneIds(long long* __out_len);
void** QTimeZone_availableTimeZoneIds2(int territory, long long* __out_len);
void** QTimeZone_availableTimeZoneIds3(int offsetSeconds, long long* __out_len);
void* QTimeZone_ianaIdToWindowsId(void* ianaId);
void* QTimeZone_windowsIdToDefaultIanaId(void* windowsId);
void* QTimeZone_windowsIdToDefaultIanaId2(void* windowsId, int territory);
void** QTimeZone_windowsIdToIanaIds(void* windowsId, long long* __out_len);
void** QTimeZone_windowsIdToIanaIds2(void* windowsId, int territory, long long* __out_len);
char* QTimeZone_displayName3(void* self, void* atDateTime, int nameType, long long* __out_len);
char* QTimeZone_displayName4(void* self, void* atDateTime, int nameType, void* locale, long long* __out_len);
char* QTimeZone_displayName5(void* self, int timeType, int nameType, long long* __out_len);
char* QTimeZone_displayName6(void* self, int timeType, int nameType, void* locale, long long* __out_len);
int QTimeZone_OperatorEqual(void* self, void* rhs);
int QTimeZone_OperatorNotEqual(void* self, void* rhs);
void QTimeZone_delete(void* self);

#ifdef __cplusplus
}
#endif
