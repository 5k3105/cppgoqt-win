#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSceneEvent_base(void* self);
void* QGraphicsSceneEvent_new(int type_);
void* QGraphicsSceneEvent_widget(void* self);
void QGraphicsSceneEvent_setWidget(void* self, void* widget);
unsigned long long QGraphicsSceneEvent_timestamp(void* self);
void QGraphicsSceneEvent_setTimestamp(void* self, unsigned long long ts);
int QGraphicsSceneEvent_registerEventType();
int QGraphicsSceneEvent_registerEventType2(int hint);
void QGraphicsSceneEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
