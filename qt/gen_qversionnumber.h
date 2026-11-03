#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVersionNumber_new();
void* QVersionNumber_new2(int maj);
void* QVersionNumber_new3(int maj, int min);
void* QVersionNumber_new4(int maj, int min, int mic);
void* QVersionNumber_new5(void* arg1);
int QVersionNumber_isNull(void* self);
int QVersionNumber_isNormalized(void* self);
int QVersionNumber_majorVersion(void* self);
int QVersionNumber_minorVersion(void* self);
int QVersionNumber_microVersion(void* self);
void* QVersionNumber_normalized(void* self);
int* QVersionNumber_segments(void* self, long long* __out_len);
int QVersionNumber_segmentAt(void* self, long long index);
long long QVersionNumber_segmentCount(void* self);
int QVersionNumber_isPrefixOf(void* self, void* other);
int QVersionNumber_compare(void* v1, void* v2);
void* QVersionNumber_commonPrefix(void* v1, void* v2);
char* QVersionNumber_toString(void* self, long long* __out_len);
void* QVersionNumber_fromString(const char* string_, long long string__len);
void* QVersionNumber_OperatorAssign(void* self, void* arg1);
void* QVersionNumber_fromString2(const char* string_, long long string__len, long long* suffixIndex);
int QVersionNumber_OperatorEqual(void* self, void* rhs);
int QVersionNumber_OperatorNotEqual(void* self, void* rhs);
int QVersionNumber_OperatorLesser(void* self, void* rhs);
int QVersionNumber_OperatorGreater(void* self, void* rhs);
int QVersionNumber_OperatorLesserOrEqual(void* self, void* rhs);
int QVersionNumber_OperatorGreaterOrEqual(void* self, void* rhs);
void QVersionNumber_delete(void* self);

#ifdef __cplusplus
}
#endif
