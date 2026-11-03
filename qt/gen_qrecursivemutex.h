#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRecursiveMutex_new();
void QRecursiveMutex_lock(void* self);
int QRecursiveMutex_tryLock(void* self, int timeout);
int QRecursiveMutex_tryLock2(void* self);
void QRecursiveMutex_unlock(void* self);
int QRecursiveMutex_try_lock(void* self);
int QRecursiveMutex_tryLock3(void* self, void* timer);
void QRecursiveMutex_delete(void* self);

#ifdef __cplusplus
}
#endif
