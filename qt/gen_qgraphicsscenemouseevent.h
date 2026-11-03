#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneMouseEvent_base(void* self);
void QGraphicsSceneMouseEvent_chain(void* self, void** out);
void QGraphicsSceneMouseEvent__set_override(void* self, int idx);
void* QGraphicsSceneMouseEvent_new(long long __goid);
void* QGraphicsSceneMouseEvent_new2(long long __goid, int type_);
void* QGraphicsSceneMouseEvent_pos(void* self);
void QGraphicsSceneMouseEvent_setPos(void* self, void* pos);
void* QGraphicsSceneMouseEvent_scenePos(void* self);
void QGraphicsSceneMouseEvent_setScenePos(void* self, void* pos);
void* QGraphicsSceneMouseEvent_screenPos(void* self);
void QGraphicsSceneMouseEvent_setScreenPos(void* self, void* pos);
void* QGraphicsSceneMouseEvent_buttonDownPos(void* self, int button);
void QGraphicsSceneMouseEvent_setButtonDownPos(void* self, int button, void* pos);
void* QGraphicsSceneMouseEvent_buttonDownScenePos(void* self, int button);
void QGraphicsSceneMouseEvent_setButtonDownScenePos(void* self, int button, void* pos);
void* QGraphicsSceneMouseEvent_buttonDownScreenPos(void* self, int button);
void QGraphicsSceneMouseEvent_setButtonDownScreenPos(void* self, int button, void* pos);
void* QGraphicsSceneMouseEvent_lastPos(void* self);
void QGraphicsSceneMouseEvent_setLastPos(void* self, void* pos);
void* QGraphicsSceneMouseEvent_lastScenePos(void* self);
void QGraphicsSceneMouseEvent_setLastScenePos(void* self, void* pos);
void* QGraphicsSceneMouseEvent_lastScreenPos(void* self);
void QGraphicsSceneMouseEvent_setLastScreenPos(void* self, void* pos);
int QGraphicsSceneMouseEvent_buttons(void* self);
void QGraphicsSceneMouseEvent_setButtons(void* self, int buttons);
int QGraphicsSceneMouseEvent_button(void* self);
void QGraphicsSceneMouseEvent_setButton(void* self, int button);
int QGraphicsSceneMouseEvent_modifiers(void* self);
void QGraphicsSceneMouseEvent_setModifiers(void* self, int modifiers);
int QGraphicsSceneMouseEvent_source(void* self);
void QGraphicsSceneMouseEvent_setSource(void* self, int source);
int QGraphicsSceneMouseEvent_flags(void* self);
void QGraphicsSceneMouseEvent_setFlags(void* self, int arg1);
int QGraphicsSceneMouseEvent_registerEventType();
int QGraphicsSceneMouseEvent_registerEventType2(int hint);
void QGraphicsSceneMouseEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneMouseEvent_clone_default(void* self);
void* QGraphicsSceneMouseEvent__Go_new(long long id);
void QGraphicsSceneMouseEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
