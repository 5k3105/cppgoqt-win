#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneWheelEvent_base(void* self);
void QGraphicsSceneWheelEvent_chain(void* self, void** out);
void QGraphicsSceneWheelEvent__set_override(void* self, int idx);
void* QGraphicsSceneWheelEvent_new(long long __goid);
void* QGraphicsSceneWheelEvent_new2(long long __goid, int type_);
void* QGraphicsSceneWheelEvent_pos(void* self);
void QGraphicsSceneWheelEvent_setPos(void* self, void* pos);
void* QGraphicsSceneWheelEvent_scenePos(void* self);
void QGraphicsSceneWheelEvent_setScenePos(void* self, void* pos);
void* QGraphicsSceneWheelEvent_screenPos(void* self);
void QGraphicsSceneWheelEvent_setScreenPos(void* self, void* pos);
int QGraphicsSceneWheelEvent_buttons(void* self);
void QGraphicsSceneWheelEvent_setButtons(void* self, int buttons);
int QGraphicsSceneWheelEvent_modifiers(void* self);
void QGraphicsSceneWheelEvent_setModifiers(void* self, int modifiers);
int QGraphicsSceneWheelEvent_delta(void* self);
void QGraphicsSceneWheelEvent_setDelta(void* self, int delta);
int QGraphicsSceneWheelEvent_orientation(void* self);
void QGraphicsSceneWheelEvent_setOrientation(void* self, int orientation);
int QGraphicsSceneWheelEvent_phase(void* self);
void QGraphicsSceneWheelEvent_setPhase(void* self, int scrollPhase);
void* QGraphicsSceneWheelEvent_pixelDelta(void* self);
void QGraphicsSceneWheelEvent_setPixelDelta(void* self, void* delta);
int QGraphicsSceneWheelEvent_isInverted(void* self);
void QGraphicsSceneWheelEvent_setInverted(void* self, int inverted);
int QGraphicsSceneWheelEvent_registerEventType();
int QGraphicsSceneWheelEvent_registerEventType2(int hint);
void QGraphicsSceneWheelEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneWheelEvent_clone_default(void* self);
void* QGraphicsSceneWheelEvent__Go_new(long long id);
void QGraphicsSceneWheelEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
