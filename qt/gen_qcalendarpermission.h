#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCalendarPermission_new();
void* QCalendarPermission_new2(void* other);
void QCalendarPermission_setAccessMode(void* self, int mode);
int QCalendarPermission_accessMode(void* self);
void* QCalendarPermission_OperatorAssign(void* self, void* other);
void QCalendarPermission_swap(void* self, void* other);
void QCalendarPermission_delete(void* self);

#ifdef __cplusplus
}
#endif
