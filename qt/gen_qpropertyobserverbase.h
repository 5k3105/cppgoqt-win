#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyObserverBase_new();
void* QPropertyObserverBase_new2(void* arg1);
void* QPropertyObserverBase_OperatorAssign(void* self, void* arg1);
void QPropertyObserverBase_delete(void* self);

#ifdef __cplusplus
}
#endif
