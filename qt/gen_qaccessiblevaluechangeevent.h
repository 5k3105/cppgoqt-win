#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleValueChangeEvent_base(void* self);
void* QAccessibleValueChangeEvent_new(void* obj, void* val);
void* QAccessibleValueChangeEvent_new2(void* iface, void* val);
void QAccessibleValueChangeEvent_setValue(void* self, void* val);
void* QAccessibleValueChangeEvent_value(void* self);
void QAccessibleValueChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
