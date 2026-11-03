#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMouseEvent_base(void* self);
void QMouseEvent_chain(void* self, void** out);
void* QMouseEvent_new(int type_, void* localPos, int button, int buttons, int modifiers);
void* QMouseEvent_new2(int type_, void* localPos, void* globalPos, int button, int buttons, int modifiers);
void* QMouseEvent_new3(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers);
void* QMouseEvent_new4(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers, int source);
void* QMouseEvent_new5(int type_, void* localPos, int button, int buttons, int modifiers, void* device);
void* QMouseEvent_new6(int type_, void* localPos, void* globalPos, int button, int buttons, int modifiers, void* device);
void* QMouseEvent_new7(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers, void* device);
void* QMouseEvent_new8(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers, int source, void* device);
void* QMouseEvent_pos(void* self);
void* QMouseEvent_globalPos(void* self);
int QMouseEvent_x(void* self);
int QMouseEvent_y(void* self);
int QMouseEvent_globalX(void* self);
int QMouseEvent_globalY(void* self);
void* QMouseEvent_localPos(void* self);
void* QMouseEvent_windowPos(void* self);
void* QMouseEvent_screenPos(void* self);
int QMouseEvent_source(void* self);
int QMouseEvent_flags(void* self);
int QMouseEvent_registerEventType();
int QMouseEvent_registerEventType2(int hint);
void QMouseEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
