#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QNativeGestureEvent_base(void* self);
void QNativeGestureEvent_chain(void* self, void** out);
void* QNativeGestureEvent_new(int type_, void* dev, void* localPos, void* scenePos, void* globalPos, double value, unsigned long long sequenceId, unsigned long long intArgument);
void* QNativeGestureEvent_new2(int type_, void* dev, int fingerCount, void* localPos, void* scenePos, void* globalPos, double value, void* delta);
void* QNativeGestureEvent_new3(int type_, void* dev, int fingerCount, void* localPos, void* scenePos, void* globalPos, double value, void* delta, unsigned long long sequenceId);
int QNativeGestureEvent_gestureType(void* self);
int QNativeGestureEvent_fingerCount(void* self);
double QNativeGestureEvent_value(void* self);
void* QNativeGestureEvent_delta(void* self);
void* QNativeGestureEvent_pos(void* self);
void* QNativeGestureEvent_globalPos(void* self);
void* QNativeGestureEvent_localPos(void* self);
void* QNativeGestureEvent_windowPos(void* self);
void* QNativeGestureEvent_screenPos(void* self);
int QNativeGestureEvent_registerEventType();
int QNativeGestureEvent_registerEventType2(int hint);
void QNativeGestureEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
