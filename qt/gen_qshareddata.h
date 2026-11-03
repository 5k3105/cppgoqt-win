#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSharedData_new();
void* QSharedData_new2(void* arg1);
void* QSharedData_ref_get(void* self);
void QSharedData_ref_set(void* self, void* ref);
void QSharedData_delete(void* self);

#ifdef __cplusplus
}
#endif
