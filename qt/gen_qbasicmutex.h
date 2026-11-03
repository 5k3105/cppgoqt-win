#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBasicMutex_new();
void QBasicMutex_lock(void* self);
void QBasicMutex_unlock(void* self);
int QBasicMutex_tryLock(void* self);
int QBasicMutex_try_lock(void* self);
void QBasicMutex_delete(void* self);

#ifdef __cplusplus
}
#endif
