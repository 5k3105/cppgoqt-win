#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QToolBarChangeEvent_base(void* self);
void* QToolBarChangeEvent_new(int t);
int QToolBarChangeEvent_toggle(void* self);
int QToolBarChangeEvent_registerEventType();
int QToolBarChangeEvent_registerEventType2(int hint);
void QToolBarChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
