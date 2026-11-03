#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPageRanges__Range_new();
void* QPageRanges__Range_new2(void* arg1);
int QPageRanges__Range_contains(void* self, int pageNumber);
int QPageRanges__Range_OperatorEqual(void* self, void* rhs);
int QPageRanges__Range_OperatorNotEqual(void* self, void* rhs);
int QPageRanges__Range_OperatorLesser(void* self, void* rhs);
int QPageRanges__Range_from_get(void* self);
void QPageRanges__Range_from_set(void* self, int from);
int QPageRanges__Range_to_get(void* self);
void QPageRanges__Range_to_set(void* self, int to);
void QPageRanges__Range_delete(void* self);

#ifdef __cplusplus
}
#endif
