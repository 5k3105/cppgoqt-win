#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleValueInterface__set_override(void* self, int idx);
void* QAccessibleValueInterface_currentValue(void* self);
void QAccessibleValueInterface_setCurrentValue(void* self, void* value);
void* QAccessibleValueInterface_maximumValue(void* self);
void* QAccessibleValueInterface_minimumValue(void* self);
void* QAccessibleValueInterface_minimumStepSize(void* self);
void* QAccessibleValueInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleValueInterface__Go_new(long long id);
void QAccessibleValueInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
