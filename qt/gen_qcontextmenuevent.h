#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QContextMenuEvent_base(void* self);
void QContextMenuEvent_chain(void* self, void** out);
void* QContextMenuEvent_new(int reason, void* pos, void* globalPos);
void* QContextMenuEvent_new2(int reason, void* pos);
void* QContextMenuEvent_new3(int reason, void* pos, void* globalPos, int modifiers);
int QContextMenuEvent_x(void* self);
int QContextMenuEvent_y(void* self);
int QContextMenuEvent_globalX(void* self);
int QContextMenuEvent_globalY(void* self);
void* QContextMenuEvent_pos(void* self);
void* QContextMenuEvent_globalPos(void* self);
int QContextMenuEvent_reason(void* self);
int QContextMenuEvent_registerEventType();
int QContextMenuEvent_registerEventType2(int hint);
void QContextMenuEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
