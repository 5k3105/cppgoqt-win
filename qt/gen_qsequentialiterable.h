#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSequentialIterable_new();
void* QSequentialIterable_new2(void* arg1);
void* QSequentialIterable_at(void* self, long long idx);
void QSequentialIterable_set(void* self, long long idx, void* value);
void QSequentialIterable_addValue(void* self, void* value);
void QSequentialIterable_removeValue(void* self);
void* QSequentialIterable_valueMetaType(void* self);
int QSequentialIterable_canRandomAccessIterate(void* self);
int QSequentialIterable_canInputIterate(void* self);
int QSequentialIterable_canForwardIterate(void* self);
int QSequentialIterable_canReverseIterate(void* self);
long long QSequentialIterable_size(void* self);
void QSequentialIterable_clear(void* self);
void* QSequentialIterable_metaContainer(void* self);
void QSequentialIterable_addValue2(void* self, void* value, int position);
void QSequentialIterable_removeValue2(void* self, int position);
void QSequentialIterable_delete(void* self);

#ifdef __cplusplus
}
#endif
