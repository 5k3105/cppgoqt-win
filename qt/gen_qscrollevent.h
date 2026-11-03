#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScrollEvent_base(void* self);
void* QScrollEvent_new(void* contentPos, void* overshoot, int scrollState);
void* QScrollEvent_contentPos(void* self);
void* QScrollEvent_overshootDistance(void* self);
int QScrollEvent_scrollState(void* self);
int QScrollEvent_registerEventType();
int QScrollEvent_registerEventType2(int hint);
void QScrollEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
