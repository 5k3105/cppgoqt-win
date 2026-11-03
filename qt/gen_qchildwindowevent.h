#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QChildWindowEvent_base(void* self);
void* QChildWindowEvent_new(int type_, void* childWindow);
void* QChildWindowEvent_child(void* self);
int QChildWindowEvent_registerEventType();
int QChildWindowEvent_registerEventType2(int hint);
void QChildWindowEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
