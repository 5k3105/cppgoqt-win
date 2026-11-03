#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDropEvent_base(void* self);
void* QDropEvent_new(void* pos, int actions, void* data, int buttons, int modifiers);
void* QDropEvent_new2(void* pos, int actions, void* data, int buttons, int modifiers, int type_);
void* QDropEvent_pos(void* self);
void* QDropEvent_posF(void* self);
int QDropEvent_mouseButtons(void* self);
int QDropEvent_keyboardModifiers(void* self);
void* QDropEvent_position(void* self);
int QDropEvent_buttons(void* self);
int QDropEvent_modifiers(void* self);
int QDropEvent_possibleActions(void* self);
int QDropEvent_proposedAction(void* self);
void QDropEvent_acceptProposedAction(void* self);
int QDropEvent_dropAction(void* self);
void QDropEvent_setDropAction(void* self, int action);
void* QDropEvent_source(void* self);
void* QDropEvent_mimeData(void* self);
int QDropEvent_registerEventType();
int QDropEvent_registerEventType2(int hint);
void QDropEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
