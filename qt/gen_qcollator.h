#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCollator_new();
void* QCollator_new2(void* locale);
void* QCollator_new3(void* arg1);
void* QCollator_OperatorAssign(void* self, void* arg1);
void QCollator_swap(void* self, void* other);
void QCollator_setLocale(void* self, void* locale);
void* QCollator_locale(void* self);
int QCollator_caseSensitivity(void* self);
void QCollator_setCaseSensitivity(void* self, int cs);
void QCollator_setNumericMode(void* self, int on);
int QCollator_numericMode(void* self);
void QCollator_setIgnorePunctuation(void* self, int on);
int QCollator_ignorePunctuation(void* self);
int QCollator_compare(void* self, const char* s1, long long s1_len, const char* s2, long long s2_len);
int QCollator_compare2(void* self, void* s1, long long len1, void* s2, long long len2);
int QCollator_OperatorCall(void* self, const char* s1, long long s1_len, const char* s2, long long s2_len);
int QCollator_compare3(void* self, void* s1, void* s2);
int QCollator_OperatorCall2(void* self, void* s1, void* s2);
void* QCollator_sortKey(void* self, const char* string_, long long string__len);
int QCollator_defaultCompare(void* s1, void* s2);
void* QCollator_defaultSortKey(void* key);
void QCollator_delete(void* self);

#ifdef __cplusplus
}
#endif
