#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessible_installActivationObserver(void* arg1);
void QAccessible_removeActivationObserver(void* arg1);
void* QAccessible_queryAccessibleInterface(void* arg1);
unsigned int QAccessible_uniqueId(void* iface);
void* QAccessible_accessibleInterface(unsigned int uniqueId);
unsigned int QAccessible_registerAccessibleInterface(void* iface);
void QAccessible_deleteAccessibleInterface(unsigned int uniqueId);
void QAccessible_updateAccessibility(void* event);
int QAccessible_isActive();
void QAccessible_setActive(int active);
void QAccessible_setRootObject(void* object);
void QAccessible_cleanup();
int QAccessible_qAccessibleTextBoundaryHelper(void* cursor, int boundaryType, int* __out_b);
void QAccessible_delete(void* self);

#ifdef __cplusplus
}
#endif
