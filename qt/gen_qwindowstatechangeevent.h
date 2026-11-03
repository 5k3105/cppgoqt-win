#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWindowStateChangeEvent_base(void* self);
void* QWindowStateChangeEvent_new(int oldState);
void* QWindowStateChangeEvent_new2(int oldState, int isOverride);
int QWindowStateChangeEvent_oldState(void* self);
int QWindowStateChangeEvent_isOverride(void* self);
int QWindowStateChangeEvent_registerEventType();
int QWindowStateChangeEvent_registerEventType2(int hint);
void QWindowStateChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
