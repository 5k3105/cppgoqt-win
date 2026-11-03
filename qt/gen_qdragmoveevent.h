#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDragMoveEvent_base(void* self);
void QDragMoveEvent_chain(void* self, void** out);
void* QDragMoveEvent_new(void* pos, int actions, void* data, int buttons, int modifiers);
void* QDragMoveEvent_new2(void* pos, int actions, void* data, int buttons, int modifiers, int type_);
void* QDragMoveEvent_answerRect(void* self);
void QDragMoveEvent_accept(void* self);
void QDragMoveEvent_ignore(void* self);
void QDragMoveEvent_accept2(void* self, void* r);
void QDragMoveEvent_ignore2(void* self, void* r);
int QDragMoveEvent_registerEventType();
int QDragMoveEvent_registerEventType2(int hint);
void QDragMoveEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
