#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWaitCondition_new();
int QWaitCondition_wait(void* self, void* lockedMutex);
int QWaitCondition_wait2(void* self, void* lockedMutex, unsigned long time);
int QWaitCondition_wait3(void* self, void* lockedReadWriteLock);
int QWaitCondition_wait4(void* self, void* lockedReadWriteLock, unsigned long time);
void QWaitCondition_wakeOne(void* self);
void QWaitCondition_wakeAll(void* self);
void QWaitCondition_notify_one(void* self);
void QWaitCondition_notify_all(void* self);
int QWaitCondition_wait5(void* self, void* lockedMutex, void* deadline);
int QWaitCondition_wait6(void* self, void* lockedReadWriteLock, void* deadline);
void QWaitCondition_delete(void* self);

#ifdef __cplusplus
}
#endif
