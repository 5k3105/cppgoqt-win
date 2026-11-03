#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDynamicPropertyChangeEvent_base(void* self);
void* QDynamicPropertyChangeEvent_new(void* name);
void* QDynamicPropertyChangeEvent_propertyName(void* self);
int QDynamicPropertyChangeEvent_registerEventType();
int QDynamicPropertyChangeEvent_registerEventType2(int hint);
void QDynamicPropertyChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
