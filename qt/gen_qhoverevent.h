#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QHoverEvent_base(void* self);
void QHoverEvent_chain(void* self, void** out);
void* QHoverEvent_new(int type_, void* scenePos, void* globalPos, void* oldPos);
void* QHoverEvent_new2(int type_, void* pos, void* oldPos);
void* QHoverEvent_new3(int type_, void* scenePos, void* globalPos, void* oldPos, int modifiers);
void* QHoverEvent_new4(int type_, void* scenePos, void* globalPos, void* oldPos, int modifiers, void* device);
void* QHoverEvent_new5(int type_, void* pos, void* oldPos, int modifiers);
void* QHoverEvent_new6(int type_, void* pos, void* oldPos, int modifiers, void* device);
void* QHoverEvent_pos(void* self);
void* QHoverEvent_posF(void* self);
void* QHoverEvent_oldPos(void* self);
void* QHoverEvent_oldPosF(void* self);
int QHoverEvent_registerEventType();
int QHoverEvent_registerEventType2(int hint);
void QHoverEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
