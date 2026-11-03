#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCameraPermission_new();
void* QCameraPermission_new2(void* other);
void* QCameraPermission_OperatorAssign(void* self, void* other);
void QCameraPermission_swap(void* self, void* other);
void QCameraPermission_delete(void* self);

#ifdef __cplusplus
}
#endif
