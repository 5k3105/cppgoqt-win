#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QApplicationStateChangeEvent_base(void* self);
void* QApplicationStateChangeEvent_new(int state);
int QApplicationStateChangeEvent_applicationState(void* self);
int QApplicationStateChangeEvent_registerEventType();
int QApplicationStateChangeEvent_registerEventType2(int hint);
void QApplicationStateChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
