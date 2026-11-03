#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDragLeaveEvent_base(void* self);
void QDragLeaveEvent__set_override(void* self, int idx);
void* QDragLeaveEvent_new(long long __goid);
int QDragLeaveEvent_registerEventType();
int QDragLeaveEvent_registerEventType2(int hint);
void* QDragLeaveEvent_clone_default(void* self);
void QDragLeaveEvent_setAccepted_default(void* self, int accepted);
void* QDragLeaveEvent__Go_new(long long id);
void QDragLeaveEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
