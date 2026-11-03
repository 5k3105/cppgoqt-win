#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMoveEvent_base(void* self);
void* QMoveEvent_new(void* pos, void* oldPos);
void* QMoveEvent_pos(void* self);
void* QMoveEvent_oldPos(void* self);
int QMoveEvent_registerEventType();
int QMoveEvent_registerEventType2(int hint);
void QMoveEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
