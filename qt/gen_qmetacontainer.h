#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaContainer_new();
void* QMetaContainer_new2(void* arg1);
int QMetaContainer_hasInputIterator(void* self);
int QMetaContainer_hasForwardIterator(void* self);
int QMetaContainer_hasBidirectionalIterator(void* self);
int QMetaContainer_hasRandomAccessIterator(void* self);
int QMetaContainer_hasSize(void* self);
int QMetaContainer_canClear(void* self);
int QMetaContainer_hasIterator(void* self);
int QMetaContainer_hasConstIterator(void* self);
void QMetaContainer_delete(void* self);

#ifdef __cplusplus
}
#endif
