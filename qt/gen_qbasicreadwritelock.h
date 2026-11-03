#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBasicReadWriteLock_new();
void QBasicReadWriteLock_lockForRead(void* self);
int QBasicReadWriteLock_tryLockForRead(void* self);
int QBasicReadWriteLock_tryLockForRead2(void* self, void* timeout);
void QBasicReadWriteLock_lockForWrite(void* self);
int QBasicReadWriteLock_tryLockForWrite(void* self);
int QBasicReadWriteLock_tryLockForWrite2(void* self, void* timeout);
void QBasicReadWriteLock_unlock(void* self);
void QBasicReadWriteLock_lock(void* self);
void QBasicReadWriteLock_lock_shared(void* self);
int QBasicReadWriteLock_try_lock(void* self);
int QBasicReadWriteLock_try_lock_shared(void* self);
void QBasicReadWriteLock_unlock_shared(void* self);
void QBasicReadWriteLock_delete(void* self);

#ifdef __cplusplus
}
#endif
