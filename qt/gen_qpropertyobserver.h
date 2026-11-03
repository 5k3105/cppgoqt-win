#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyObserver_base(void* self);
void* QPropertyObserver_new();
void QPropertyObserver_delete(void* self);

#ifdef __cplusplus
}
#endif
