#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneHoverEvent_base(void* self);
void QGraphicsSceneHoverEvent_chain(void* self, void** out);
void QGraphicsSceneHoverEvent__set_override(void* self, int idx);
void* QGraphicsSceneHoverEvent_new(long long __goid);
void* QGraphicsSceneHoverEvent_new2(long long __goid, int type_);
void* QGraphicsSceneHoverEvent_pos(void* self);
void QGraphicsSceneHoverEvent_setPos(void* self, void* pos);
void* QGraphicsSceneHoverEvent_scenePos(void* self);
void QGraphicsSceneHoverEvent_setScenePos(void* self, void* pos);
void* QGraphicsSceneHoverEvent_screenPos(void* self);
void QGraphicsSceneHoverEvent_setScreenPos(void* self, void* pos);
void* QGraphicsSceneHoverEvent_lastPos(void* self);
void QGraphicsSceneHoverEvent_setLastPos(void* self, void* pos);
void* QGraphicsSceneHoverEvent_lastScenePos(void* self);
void QGraphicsSceneHoverEvent_setLastScenePos(void* self, void* pos);
void* QGraphicsSceneHoverEvent_lastScreenPos(void* self);
void QGraphicsSceneHoverEvent_setLastScreenPos(void* self, void* pos);
int QGraphicsSceneHoverEvent_modifiers(void* self);
void QGraphicsSceneHoverEvent_setModifiers(void* self, int modifiers);
int QGraphicsSceneHoverEvent_registerEventType();
int QGraphicsSceneHoverEvent_registerEventType2(int hint);
void QGraphicsSceneHoverEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneHoverEvent_clone_default(void* self);
void* QGraphicsSceneHoverEvent__Go_new(long long id);
void QGraphicsSceneHoverEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
