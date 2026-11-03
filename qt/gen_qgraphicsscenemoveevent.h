#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneMoveEvent_base(void* self);
void QGraphicsSceneMoveEvent_chain(void* self, void** out);
void QGraphicsSceneMoveEvent__set_override(void* self, int idx);
void* QGraphicsSceneMoveEvent_new(long long __goid);
void* QGraphicsSceneMoveEvent_oldPos(void* self);
void QGraphicsSceneMoveEvent_setOldPos(void* self, void* pos);
void* QGraphicsSceneMoveEvent_newPos(void* self);
void QGraphicsSceneMoveEvent_setNewPos(void* self, void* pos);
int QGraphicsSceneMoveEvent_registerEventType();
int QGraphicsSceneMoveEvent_registerEventType2(int hint);
void QGraphicsSceneMoveEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneMoveEvent_clone_default(void* self);
void* QGraphicsSceneMoveEvent__Go_new(long long id);
void QGraphicsSceneMoveEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
