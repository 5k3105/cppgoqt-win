#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTimerEvent_base(void* self);
void* QTimerEvent_new(int timerId);
void* QTimerEvent_new2(int timerId);
int QTimerEvent_timerId(void* self);
int QTimerEvent_id(void* self);
int QTimerEvent_matches(void* self, void* timer);
int QTimerEvent_registerEventType();
int QTimerEvent_registerEventType2(int hint);
void QTimerEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
