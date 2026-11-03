#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAssociativeIterable_new();
void* QAssociativeIterable_new2(void* arg1);
int QAssociativeIterable_containsKey(void* self, void* key);
void QAssociativeIterable_insertKey(void* self, void* key);
void QAssociativeIterable_removeKey(void* self, void* key);
void* QAssociativeIterable_value(void* self, void* key);
void QAssociativeIterable_setValue(void* self, void* key, void* mapped);
int QAssociativeIterable_canRandomAccessIterate(void* self);
int QAssociativeIterable_canInputIterate(void* self);
int QAssociativeIterable_canForwardIterate(void* self);
int QAssociativeIterable_canReverseIterate(void* self);
long long QAssociativeIterable_size(void* self);
void QAssociativeIterable_clear(void* self);
void* QAssociativeIterable_metaContainer(void* self);
void QAssociativeIterable_delete(void* self);

#ifdef __cplusplus
}
#endif
