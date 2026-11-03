#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneDragDropEvent_base(void* self);
void QGraphicsSceneDragDropEvent_chain(void* self, void** out);
void QGraphicsSceneDragDropEvent__set_override(void* self, int idx);
void* QGraphicsSceneDragDropEvent_new(long long __goid);
void* QGraphicsSceneDragDropEvent_new2(long long __goid, int type_);
void* QGraphicsSceneDragDropEvent_pos(void* self);
void QGraphicsSceneDragDropEvent_setPos(void* self, void* pos);
void* QGraphicsSceneDragDropEvent_scenePos(void* self);
void QGraphicsSceneDragDropEvent_setScenePos(void* self, void* pos);
void* QGraphicsSceneDragDropEvent_screenPos(void* self);
void QGraphicsSceneDragDropEvent_setScreenPos(void* self, void* pos);
int QGraphicsSceneDragDropEvent_buttons(void* self);
void QGraphicsSceneDragDropEvent_setButtons(void* self, int buttons);
int QGraphicsSceneDragDropEvent_modifiers(void* self);
void QGraphicsSceneDragDropEvent_setModifiers(void* self, int modifiers);
int QGraphicsSceneDragDropEvent_possibleActions(void* self);
void QGraphicsSceneDragDropEvent_setPossibleActions(void* self, int actions);
int QGraphicsSceneDragDropEvent_proposedAction(void* self);
void QGraphicsSceneDragDropEvent_setProposedAction(void* self, int action);
void QGraphicsSceneDragDropEvent_acceptProposedAction(void* self);
int QGraphicsSceneDragDropEvent_dropAction(void* self);
void QGraphicsSceneDragDropEvent_setDropAction(void* self, int action);
void* QGraphicsSceneDragDropEvent_source(void* self);
void QGraphicsSceneDragDropEvent_setSource(void* self, void* source);
void* QGraphicsSceneDragDropEvent_mimeData(void* self);
void QGraphicsSceneDragDropEvent_setMimeData(void* self, void* data);
int QGraphicsSceneDragDropEvent_registerEventType();
int QGraphicsSceneDragDropEvent_registerEventType2(int hint);
void QGraphicsSceneDragDropEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneDragDropEvent_clone_default(void* self);
void* QGraphicsSceneDragDropEvent__Go_new(long long id);
void QGraphicsSceneDragDropEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
