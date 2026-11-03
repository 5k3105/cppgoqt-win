#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QResizeEvent_base(void* self);
void* QResizeEvent_new(void* size, void* oldSize);
void* QResizeEvent_size(void* self);
void* QResizeEvent_oldSize(void* self);
int QResizeEvent_registerEventType();
int QResizeEvent_registerEventType2(int hint);
void QResizeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
