#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLocationPermission_new();
void* QLocationPermission_new2(void* other);
void QLocationPermission_setAccuracy(void* self, int accuracy);
int QLocationPermission_accuracy(void* self);
void QLocationPermission_setAvailability(void* self, int availability);
int QLocationPermission_availability(void* self);
void* QLocationPermission_OperatorAssign(void* self, void* other);
void QLocationPermission_swap(void* self, void* other);
void QLocationPermission_delete(void* self);

#ifdef __cplusplus
}
#endif
