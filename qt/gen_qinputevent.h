#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QInputEvent_base(void* self);
void* QInputEvent_new(int type_, void* m_dev);
void* QInputEvent_new2(int type_, void* m_dev, int modifiers);
void* QInputEvent_device(void* self);
int QInputEvent_deviceType(void* self);
int QInputEvent_modifiers(void* self);
void QInputEvent_setModifiers(void* self, int modifiers);
unsigned long long QInputEvent_timestamp(void* self);
void QInputEvent_setTimestamp(void* self, unsigned long long timestamp);
int QInputEvent_registerEventType();
int QInputEvent_registerEventType2(int hint);
void QInputEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
