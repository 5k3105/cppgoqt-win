#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCloseEvent_base(void* self);
void QCloseEvent__set_override(void* self, int idx);
void* QCloseEvent_new(long long __goid);
int QCloseEvent_registerEventType();
int QCloseEvent_registerEventType2(int hint);
void* QCloseEvent_clone_default(void* self);
void QCloseEvent_setAccepted_default(void* self, int accepted);
void* QCloseEvent__Go_new(long long id);
void QCloseEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
