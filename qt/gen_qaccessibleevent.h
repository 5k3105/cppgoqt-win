#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleEvent_new(void* obj, int typ);
void* QAccessibleEvent_new2(void* iface, int typ);
int QAccessibleEvent_type(void* self);
void* QAccessibleEvent_object(void* self);
unsigned int QAccessibleEvent_uniqueId(void* self);
void QAccessibleEvent_setChild(void* self, int chld);
int QAccessibleEvent_child(void* self);
void* QAccessibleEvent_accessibleInterface(void* self);
void QAccessibleEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
