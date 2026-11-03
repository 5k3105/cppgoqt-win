#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBindingStorage_new();
int QBindingStorage_isEmpty(void* self);
int QBindingStorage_isValid(void* self);
void QBindingStorage_registerDependency(void* self, void* data);
void QBindingStorage_delete(void* self);

#ifdef __cplusplus
}
#endif
