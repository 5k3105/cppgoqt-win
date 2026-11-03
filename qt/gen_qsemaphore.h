#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSemaphore_new();
void* QSemaphore_new2(int n);
void QSemaphore_acquire(void* self);
int QSemaphore_tryAcquire(void* self);
int QSemaphore_tryAcquire2(void* self, int n, int timeout);
int QSemaphore_tryAcquire3(void* self, int n, void* timeout);
void QSemaphore_release(void* self);
int QSemaphore_available(void* self);
int QSemaphore_try_acquire(void* self);
void QSemaphore_acquire2(void* self, int n);
int QSemaphore_tryAcquire4(void* self, int n);
void QSemaphore_release2(void* self, int n);
void QSemaphore_delete(void* self);

#ifdef __cplusplus
}
#endif
