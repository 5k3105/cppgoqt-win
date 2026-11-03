#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPointingDeviceUniqueId_new();
void* QPointingDeviceUniqueId_new2(void* arg1);
void* QPointingDeviceUniqueId_fromNumericId(long long id);
int QPointingDeviceUniqueId_isValid(void* self);
long long QPointingDeviceUniqueId_numericId(void* self);
int QPointingDeviceUniqueId_OperatorEqual(void* self, void* rhs);
int QPointingDeviceUniqueId_OperatorNotEqual(void* self, void* rhs);
void QPointingDeviceUniqueId_delete(void* self);

#ifdef __cplusplus
}
#endif
