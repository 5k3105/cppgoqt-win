#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractEventDispatcher__TimerInfoV2_new(void* arg1);
void* QAbstractEventDispatcher__TimerInfoV2_new2();
void* QAbstractEventDispatcher__TimerInfoV2_OperatorAssign(void* self, void* arg1);
int QAbstractEventDispatcher__TimerInfoV2_timerId_get(void* self);
void QAbstractEventDispatcher__TimerInfoV2_timerId_set(void* self, int timerId);
int QAbstractEventDispatcher__TimerInfoV2_timerType_get(void* self);
void QAbstractEventDispatcher__TimerInfoV2_timerType_set(void* self, int timerType);
void QAbstractEventDispatcher__TimerInfoV2_delete(void* self);

#ifdef __cplusplus
}
#endif
