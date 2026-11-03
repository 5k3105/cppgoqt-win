#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMutex_base(void* self);
void* QMutex_new();
int QMutex_try_lock(void* self);
int QMutex_tryLock(void* self, int timeout);
int QMutex_tryLock2(void* self, void* timeout);
void QMutex_delete(void* self);

#ifdef __cplusplus
}
#endif
