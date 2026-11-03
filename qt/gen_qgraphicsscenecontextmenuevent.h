#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneContextMenuEvent_base(void* self);
void QGraphicsSceneContextMenuEvent_chain(void* self, void** out);
void QGraphicsSceneContextMenuEvent__set_override(void* self, int idx);
void* QGraphicsSceneContextMenuEvent_new(long long __goid);
void* QGraphicsSceneContextMenuEvent_new2(long long __goid, int type_);
void* QGraphicsSceneContextMenuEvent_pos(void* self);
void QGraphicsSceneContextMenuEvent_setPos(void* self, void* pos);
void* QGraphicsSceneContextMenuEvent_scenePos(void* self);
void QGraphicsSceneContextMenuEvent_setScenePos(void* self, void* pos);
void* QGraphicsSceneContextMenuEvent_screenPos(void* self);
void QGraphicsSceneContextMenuEvent_setScreenPos(void* self, void* pos);
int QGraphicsSceneContextMenuEvent_modifiers(void* self);
void QGraphicsSceneContextMenuEvent_setModifiers(void* self, int modifiers);
int QGraphicsSceneContextMenuEvent_reason(void* self);
void QGraphicsSceneContextMenuEvent_setReason(void* self, int reason);
int QGraphicsSceneContextMenuEvent_registerEventType();
int QGraphicsSceneContextMenuEvent_registerEventType2(int hint);
void QGraphicsSceneContextMenuEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneContextMenuEvent_clone_default(void* self);
void* QGraphicsSceneContextMenuEvent__Go_new(long long id);
void QGraphicsSceneContextMenuEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
