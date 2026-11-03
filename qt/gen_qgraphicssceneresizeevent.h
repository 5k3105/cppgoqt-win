#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneResizeEvent_base(void* self);
void QGraphicsSceneResizeEvent_chain(void* self, void** out);
void QGraphicsSceneResizeEvent__set_override(void* self, int idx);
void* QGraphicsSceneResizeEvent_new(long long __goid);
void* QGraphicsSceneResizeEvent_oldSize(void* self);
void QGraphicsSceneResizeEvent_setOldSize(void* self, void* size);
void* QGraphicsSceneResizeEvent_newSize(void* self);
void QGraphicsSceneResizeEvent_setNewSize(void* self, void* size);
int QGraphicsSceneResizeEvent_registerEventType();
int QGraphicsSceneResizeEvent_registerEventType2(int hint);
void QGraphicsSceneResizeEvent_setAccepted_default(void* self, int accepted);
void* QGraphicsSceneResizeEvent_clone_default(void* self);
void* QGraphicsSceneResizeEvent__Go_new(long long id);
void QGraphicsSceneResizeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
