#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QHideEvent_base(void* self);
void QHideEvent__set_override(void* self, int idx);
void* QHideEvent_new(long long __goid);
int QHideEvent_registerEventType();
int QHideEvent_registerEventType2(int hint);
void* QHideEvent_clone_default(void* self);
void QHideEvent_setAccepted_default(void* self, int accepted);
void* QHideEvent__Go_new(long long id);
void QHideEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
