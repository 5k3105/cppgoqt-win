#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QByteArrayMatcher_new();
void* QByteArrayMatcher_new2(void* pattern);
void* QByteArrayMatcher_new3(void* pattern);
void* QByteArrayMatcher_new4(const char* pattern);
void* QByteArrayMatcher_new5(void* other);
void* QByteArrayMatcher_new6(const char* pattern, long long length);
void* QByteArrayMatcher_OperatorAssign(void* self, void* other);
void QByteArrayMatcher_setPattern(void* self, void* pattern);
long long QByteArrayMatcher_indexIn(void* self, const char* str, long long len_);
long long QByteArrayMatcher_indexIn2(void* self, void* data);
void* QByteArrayMatcher_pattern(void* self);
long long QByteArrayMatcher_indexIn3(void* self, const char* str, long long len_, long long from);
long long QByteArrayMatcher_indexIn4(void* self, void* data, long long from);
void QByteArrayMatcher_delete(void* self);

#ifdef __cplusplus
}
#endif
