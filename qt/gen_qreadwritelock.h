#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QReadWriteLock_base(void* self);
void* QReadWriteLock_new();
void* QReadWriteLock_new2(int recursionMode);
int QReadWriteLock_tryLockForRead(void* self, int timeout);
int QReadWriteLock_tryLockForWrite(void* self, int timeout);
void QReadWriteLock_delete(void* self);

#ifdef __cplusplus
}
#endif
