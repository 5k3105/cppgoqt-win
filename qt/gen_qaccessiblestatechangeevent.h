#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleStateChangeEvent_base(void* self);
void* QAccessibleStateChangeEvent_new(void* obj, void* state);
void* QAccessibleStateChangeEvent_new2(void* iface, void* state);
void* QAccessibleStateChangeEvent_changedStates(void* self);
void QAccessibleStateChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
