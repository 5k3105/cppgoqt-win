#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QEnterEvent_base(void* self);
void QEnterEvent_chain(void* self, void** out);
void* QEnterEvent_new(void* localPos, void* scenePos, void* globalPos);
void* QEnterEvent_new2(void* localPos, void* scenePos, void* globalPos, void* device);
void* QEnterEvent_pos(void* self);
void* QEnterEvent_globalPos(void* self);
int QEnterEvent_x(void* self);
int QEnterEvent_y(void* self);
int QEnterEvent_globalX(void* self);
int QEnterEvent_globalY(void* self);
void* QEnterEvent_localPos(void* self);
void* QEnterEvent_windowPos(void* self);
void* QEnterEvent_screenPos(void* self);
int QEnterEvent_registerEventType();
int QEnterEvent_registerEventType2(int hint);
void QEnterEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
