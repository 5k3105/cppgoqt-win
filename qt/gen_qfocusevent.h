#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFocusEvent_base(void* self);
void* QFocusEvent_new(int type_);
void* QFocusEvent_new2(int type_, int reason);
int QFocusEvent_gotFocus(void* self);
int QFocusEvent_lostFocus(void* self);
int QFocusEvent_reason(void* self);
int QFocusEvent_registerEventType();
int QFocusEvent_registerEventType2(int hint);
void QFocusEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
