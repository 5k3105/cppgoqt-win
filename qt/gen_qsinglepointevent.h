#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSinglePointEvent_base(void* self);
void QSinglePointEvent_chain(void* self, void** out);
int QSinglePointEvent_button(void* self);
int QSinglePointEvent_buttons(void* self);
void* QSinglePointEvent_position(void* self);
void* QSinglePointEvent_scenePosition(void* self);
void* QSinglePointEvent_globalPosition(void* self);
void* QSinglePointEvent_exclusivePointGrabber(void* self);
void QSinglePointEvent_setExclusivePointGrabber(void* self, void* exclusiveGrabber);
int QSinglePointEvent_registerEventType();
int QSinglePointEvent_registerEventType2(int hint);
void QSinglePointEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
