#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QChildEvent_base(void* self);
void* QChildEvent_new(int type_, void* child);
void* QChildEvent_child(void* self);
int QChildEvent_added(void* self);
int QChildEvent_polished(void* self);
int QChildEvent_removed(void* self);
int QChildEvent_registerEventType();
int QChildEvent_registerEventType2(int hint);
void QChildEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
