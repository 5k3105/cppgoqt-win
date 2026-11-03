#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScreenOrientationChangeEvent_base(void* self);
void* QScreenOrientationChangeEvent_new(void* screen, int orientation);
void* QScreenOrientationChangeEvent_screen(void* self);
int QScreenOrientationChangeEvent_orientation(void* self);
int QScreenOrientationChangeEvent_registerEventType();
int QScreenOrientationChangeEvent_registerEventType2(int hint);
void QScreenOrientationChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
