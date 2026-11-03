#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QReadLocker_new(void* readWriteLock);
void QReadLocker_unlock(void* self);
void QReadLocker_relock(void* self);
void* QReadLocker_readWriteLock(void* self);
void QReadLocker_delete(void* self);

#ifdef __cplusplus
}
#endif
