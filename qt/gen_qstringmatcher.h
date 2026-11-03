#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStringMatcher_new();
void* QStringMatcher_new2(const char* pattern, long long pattern_len);
void* QStringMatcher_new3(void* uc, long long len_);
void* QStringMatcher_new4(void* pattern);
void* QStringMatcher_new5(void* other);
void* QStringMatcher_new6(const char* pattern, long long pattern_len, int cs);
void* QStringMatcher_new7(void* uc, long long len_, int cs);
void* QStringMatcher_new8(void* pattern, int cs);
void* QStringMatcher_OperatorAssign(void* self, void* other);
void QStringMatcher_setPattern(void* self, const char* pattern, long long pattern_len);
void QStringMatcher_setCaseSensitivity(void* self, int cs);
long long QStringMatcher_indexIn(void* self, const char* str, long long str_len);
long long QStringMatcher_indexIn2(void* self, void* str, long long length);
long long QStringMatcher_indexIn3(void* self, void* str);
char* QStringMatcher_pattern(void* self, long long* __out_len);
void* QStringMatcher_patternView(void* self);
int QStringMatcher_caseSensitivity(void* self);
long long QStringMatcher_indexIn4(void* self, const char* str, long long str_len, long long from);
long long QStringMatcher_indexIn5(void* self, void* str, long long length, long long from);
long long QStringMatcher_indexIn6(void* self, void* str, long long from);
void QStringMatcher_delete(void* self);

#ifdef __cplusplus
}
#endif
