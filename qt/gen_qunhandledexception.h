#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUnhandledException_base(void* self);
void* QUnhandledException_new();
void* QUnhandledException_new2(void* other);
void QUnhandledException_swap(void* self, void* other);
void* QUnhandledException_OperatorAssign(void* self, void* other);
void QUnhandledException_delete(void* self);

#ifdef __cplusplus
}
#endif
