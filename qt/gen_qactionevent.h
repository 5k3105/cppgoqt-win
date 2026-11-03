#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QActionEvent_base(void* self);
void* QActionEvent_new(int type_, void* action);
void* QActionEvent_new2(int type_, void* action, void* before);
void* QActionEvent_action(void* self);
void* QActionEvent_before(void* self);
int QActionEvent_registerEventType();
int QActionEvent_registerEventType2(int hint);
void QActionEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
