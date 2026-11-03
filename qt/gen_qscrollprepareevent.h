#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScrollPrepareEvent_base(void* self);
void* QScrollPrepareEvent_new(void* startPos);
void* QScrollPrepareEvent_startPos(void* self);
void* QScrollPrepareEvent_viewportSize(void* self);
void* QScrollPrepareEvent_contentPosRange(void* self);
void* QScrollPrepareEvent_contentPos(void* self);
void QScrollPrepareEvent_setViewportSize(void* self, void* size);
void QScrollPrepareEvent_setContentPosRange(void* self, void* rect);
void QScrollPrepareEvent_setContentPos(void* self, void* pos);
int QScrollPrepareEvent_registerEventType();
int QScrollPrepareEvent_registerEventType2(int hint);
void QScrollPrepareEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
