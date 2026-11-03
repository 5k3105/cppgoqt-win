#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIODeviceBase_new(void* arg1);
void* QIODeviceBase_OperatorAssign(void* self, void* arg1);
void QIODeviceBase_delete(void* self);

#ifdef __cplusplus
}
#endif
