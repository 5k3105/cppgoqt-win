#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIconDragEvent_base(void* self);
void QIconDragEvent__set_override(void* self, int idx);
void* QIconDragEvent_new(long long __goid);
int QIconDragEvent_registerEventType();
int QIconDragEvent_registerEventType2(int hint);
void* QIconDragEvent_clone_default(void* self);
void QIconDragEvent_setAccepted_default(void* self, int accepted);
void* QIconDragEvent__Go_new(long long id);
void QIconDragEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
