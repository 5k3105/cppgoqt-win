#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBasicTimer_new();
void QBasicTimer_swap(void* self, void* other);
int QBasicTimer_isActive(void* self);
int QBasicTimer_timerId(void* self);
int QBasicTimer_id(void* self);
void QBasicTimer_start(void* self, int msec, void* obj);
void QBasicTimer_start2(void* self, int msec, int timerType, void* obj);
void QBasicTimer_stop(void* self);
void QBasicTimer_delete(void* self);

#ifdef __cplusplus
}
#endif
