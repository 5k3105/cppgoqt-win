#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVariantConstPointer_new(void* variant);
void* QVariantConstPointer_new2(void* arg1);
void* QVariantConstPointer_OperatorMultiply(void* self);
void* QVariantConstPointer_OperatorMinusGreater(void* self);
void* QVariantConstPointer_OperatorAssign(void* self, void* arg1);
void QVariantConstPointer_delete(void* self);

#ifdef __cplusplus
}
#endif
