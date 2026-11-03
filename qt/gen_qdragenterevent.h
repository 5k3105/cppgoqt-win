#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDragEnterEvent_base(void* self);
void QDragEnterEvent_chain(void* self, void** out);
void* QDragEnterEvent_new(void* pos, int actions, void* data, int buttons, int modifiers);
int QDragEnterEvent_registerEventType();
int QDragEnterEvent_registerEventType2(int hint);
void QDragEnterEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
