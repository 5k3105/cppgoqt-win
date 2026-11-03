#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QHelpEvent_base(void* self);
void* QHelpEvent_new(int type_, void* pos, void* globalPos);
int QHelpEvent_x(void* self);
int QHelpEvent_y(void* self);
int QHelpEvent_globalX(void* self);
int QHelpEvent_globalY(void* self);
void* QHelpEvent_pos(void* self);
void* QHelpEvent_globalPos(void* self);
int QHelpEvent_registerEventType();
int QHelpEvent_registerEventType2(int hint);
void QHelpEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
