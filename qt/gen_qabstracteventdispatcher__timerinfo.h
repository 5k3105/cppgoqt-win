#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractEventDispatcher__TimerInfo_new();
void* QAbstractEventDispatcher__TimerInfo_new2(int id, int i, int t);
void* QAbstractEventDispatcher__TimerInfo_new3(void* arg1);
int QAbstractEventDispatcher__TimerInfo_timerId_get(void* self);
void QAbstractEventDispatcher__TimerInfo_timerId_set(void* self, int timerId);
int QAbstractEventDispatcher__TimerInfo_interval_get(void* self);
void QAbstractEventDispatcher__TimerInfo_interval_set(void* self, int interval);
int QAbstractEventDispatcher__TimerInfo_timerType_get(void* self);
void QAbstractEventDispatcher__TimerInfo_timerType_set(void* self, int timerType);
void QAbstractEventDispatcher__TimerInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
