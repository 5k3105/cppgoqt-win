#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAssociativeIterator_new(void* arg1);
void* QAssociativeIterator_key(void* self);
void* QAssociativeIterator_metaContainer(void* self);
void QAssociativeIterator_delete(void* self);

#ifdef __cplusplus
}
#endif
