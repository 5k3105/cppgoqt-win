#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPageRanges_new();
void* QPageRanges_new2(void* other);
void* QPageRanges_OperatorAssign(void* self, void* other);
void QPageRanges_swap(void* self, void* other);
void QPageRanges_addPage(void* self, int pageNumber);
void QPageRanges_addRange(void* self, int from, int to);
void** QPageRanges_toRangeList(void* self, long long* __out_len);
void QPageRanges_clear(void* self);
char* QPageRanges_toString(void* self, long long* __out_len);
void* QPageRanges_fromString(const char* ranges, long long ranges_len);
int QPageRanges_contains(void* self, int pageNumber);
int QPageRanges_isEmpty(void* self);
int QPageRanges_firstPage(void* self);
int QPageRanges_lastPage(void* self);
void QPageRanges_detach(void* self);
int QPageRanges_OperatorEqual(void* self, void* rhs);
int QPageRanges_OperatorNotEqual(void* self, void* rhs);
void QPageRanges_delete(void* self);

#ifdef __cplusplus
}
#endif
