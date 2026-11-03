#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSequentialConstIterator_new(void* arg1);
void* QSequentialConstIterator_OperatorMultiply(void* self);
void* QSequentialConstIterator_OperatorMinusGreater(void* self);
void* QSequentialConstIterator_metaContainer(void* self);
void QSequentialConstIterator_delete(void* self);

#ifdef __cplusplus
}
#endif
