#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QEvent_new(int type_);
int QEvent_type(void* self);
int QEvent_spontaneous(void* self);
void QEvent_setAccepted(void* self, int accepted);
int QEvent_isAccepted(void* self);
void QEvent_accept(void* self);
void QEvent_ignore(void* self);
int QEvent_isInputEvent(void* self);
int QEvent_isPointerEvent(void* self);
int QEvent_isSinglePointEvent(void* self);
int QEvent_registerEventType();
void* QEvent_clone(void* self);
int QEvent_registerEventType2(int hint);
void QEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
