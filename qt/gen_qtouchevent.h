#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTouchEvent_base(void* self);
void QTouchEvent_chain(void* self, void** out);
void* QTouchEvent_new(int eventType);
void* QTouchEvent_new2(int eventType, void* device, int modifiers, int touchPointStates);
void* QTouchEvent_new3(int eventType, void* device);
void* QTouchEvent_new4(int eventType, void* device, int modifiers);
void* QTouchEvent_new5(int eventType, void* device, int modifiers, void** touchPoints, long long touchPoints_len);
void* QTouchEvent_new6(int eventType, void* device, int modifiers, int touchPointStates, void** touchPoints, long long touchPoints_len);
void* QTouchEvent_target(void* self);
int QTouchEvent_touchPointStates(void* self);
void** QTouchEvent_touchPoints(void* self, long long* __out_len);
int QTouchEvent_registerEventType();
int QTouchEvent_registerEventType2(int hint);
void QTouchEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
