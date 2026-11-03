#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPermission_new();
void* QPermission_new2(void* arg1);
int QPermission_status(void* self);
void* QPermission_type(void* self);
void* QPermission_OperatorAssign(void* self, void* arg1);
void QPermission_delete(void* self);

#ifdef __cplusplus
}
#endif
