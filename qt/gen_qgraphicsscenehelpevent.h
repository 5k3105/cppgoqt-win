#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneHelpEvent_base(void* self);
void QGraphicsSceneHelpEvent_chain(void* self, void** out);
void QGraphicsSceneHelpEvent__set_override(void* self, int idx);
void* QGraphicsSceneHelpEvent_new(long long __goid);
void* QGraphicsSceneHelpEvent_new2(long long __goid, int type_);
void* QGraphicsSceneHelpEvent_scenePos(void* self);
void QGraphicsSceneHelpEvent_setScenePos(void* self, void* pos);
void* QGraphicsSceneHelpEvent_screenPos(void* self);
void QGraphicsSceneHelpEvent_setScreenPos(void* self, void* pos);
int QGraphicsSceneHelpEvent_registerEventType();
int QGraphicsSceneHelpEvent_registerEventType2(int hint);
void QGraphicsSceneHelpEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneHelpEvent_clone_default(void* self);
void* QGraphicsSceneHelpEvent__Go_new(long long id);
void QGraphicsSceneHelpEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
