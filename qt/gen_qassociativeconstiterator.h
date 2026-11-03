#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAssociativeConstIterator_new(void* arg1);
void* QAssociativeConstIterator_key(void* self);
void* QAssociativeConstIterator_value(void* self);
void* QAssociativeConstIterator_OperatorMultiply(void* self);
void* QAssociativeConstIterator_OperatorMinusGreater(void* self);
void* QAssociativeConstIterator_metaContainer(void* self);
void QAssociativeConstIterator_delete(void* self);

#ifdef __cplusplus
}
#endif
