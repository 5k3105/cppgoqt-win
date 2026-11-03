#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLatin1StringMatcher_new();
void* QLatin1StringMatcher_new2(void* pattern);
void* QLatin1StringMatcher_new3(void* pattern, int cs);
void QLatin1StringMatcher_setPattern(void* self, void* pattern);
void* QLatin1StringMatcher_pattern(void* self);
void QLatin1StringMatcher_setCaseSensitivity(void* self, int cs);
int QLatin1StringMatcher_caseSensitivity(void* self);
long long QLatin1StringMatcher_indexIn(void* self, void* haystack);
long long QLatin1StringMatcher_indexIn2(void* self, void* haystack);
long long QLatin1StringMatcher_indexIn3(void* self, void* haystack, long long from);
long long QLatin1StringMatcher_indexIn4(void* self, void* haystack, long long from);
void QLatin1StringMatcher_delete(void* self);

#ifdef __cplusplus
}
#endif
