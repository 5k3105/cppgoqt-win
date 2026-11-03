#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QShowEvent_base(void* self);
void QShowEvent__set_override(void* self, int idx);
void* QShowEvent_new(long long __goid);
int QShowEvent_registerEventType();
int QShowEvent_registerEventType2(int hint);
void* QShowEvent_clone_default(void* self);
void QShowEvent_setAccepted_default(void* self, int accepted);
void* QShowEvent__Go_new(long long id);
void QShowEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
