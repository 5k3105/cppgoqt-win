#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPointerEvent_base(void* self);
void QPointerEvent_chain(void* self, void** out);
void* QPointerEvent_new(int type_, void* dev);
void* QPointerEvent_new2(int type_, void* dev, int modifiers);
void* QPointerEvent_new3(int type_, void* dev, int modifiers, void** points, long long points_len);
void* QPointerEvent_pointingDevice(void* self);
int QPointerEvent_pointerType(void* self);
long long QPointerEvent_pointCount(void* self);
void* QPointerEvent_point(void* self, long long i);
void** QPointerEvent_points(void* self, long long* __out_len);
void* QPointerEvent_pointById(void* self, int id);
int QPointerEvent_allPointsGrabbed(void* self);
int QPointerEvent_isBeginEvent(void* self);
int QPointerEvent_isUpdateEvent(void* self);
int QPointerEvent_isEndEvent(void* self);
int QPointerEvent_allPointsAccepted(void* self);
void* QPointerEvent_exclusiveGrabber(void* self, void* point);
void QPointerEvent_setExclusiveGrabber(void* self, void* point, void* exclusiveGrabber);
void QPointerEvent_clearPassiveGrabbers(void* self, void* point);
int QPointerEvent_addPassiveGrabber(void* self, void* point, void* grabber);
int QPointerEvent_removePassiveGrabber(void* self, void* point, void* grabber);
int QPointerEvent_registerEventType();
int QPointerEvent_registerEventType2(int hint);
void QPointerEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
