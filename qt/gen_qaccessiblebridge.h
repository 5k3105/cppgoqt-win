#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleBridge__set_override(void* self, int idx);
void QAccessibleBridge_setRootObject(void* self, void* arg1);
void QAccessibleBridge_notifyAccessibilityUpdate(void* self, void* event);
void* QAccessibleBridge_OperatorAssign(void* self, void* arg1);
void* QAccessibleBridge__Go_new(long long id);
void QAccessibleBridge_delete(void* self);

#ifdef __cplusplus
}
#endif
