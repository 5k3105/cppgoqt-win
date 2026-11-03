#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGestureEvent_base(void* self);
void* QGestureEvent_new(void** gestures, long long gestures_len);
void* QGestureEvent_new2(void* arg1);
void** QGestureEvent_gestures(void* self, long long* __out_len);
void* QGestureEvent_gesture(void* self, int type_);
void** QGestureEvent_activeGestures(void* self, long long* __out_len);
void** QGestureEvent_canceledGestures(void* self, long long* __out_len);
void QGestureEvent_setAccepted(void* self, void* arg1, int arg2);
void QGestureEvent_accept(void* self, void* arg1);
void QGestureEvent_ignore(void* self, void* arg1);
int QGestureEvent_isAccepted(void* self, void* arg1);
void QGestureEvent_setAccepted2(void* self, int arg1, int arg2);
void QGestureEvent_accept2(void* self, int arg1);
void QGestureEvent_ignore2(void* self, int arg1);
int QGestureEvent_isAccepted2(void* self, int arg1);
void QGestureEvent_setWidget(void* self, void* widget);
void* QGestureEvent_widget(void* self);
void* QGestureEvent_mapToGraphicsScene(void* self, void* gesturePoint);
void* QGestureEvent_OperatorAssign(void* self, void* arg1);
int QGestureEvent_registerEventType();
int QGestureEvent_registerEventType2(int hint);
void QGestureEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
