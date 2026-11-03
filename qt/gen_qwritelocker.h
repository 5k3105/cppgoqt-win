#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWriteLocker_new(void* readWriteLock);
void QWriteLocker_unlock(void* self);
void QWriteLocker_relock(void* self);
void* QWriteLocker_readWriteLock(void* self);
void QWriteLocker_delete(void* self);

#ifdef __cplusplus
}
#endif
