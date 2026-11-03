#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothPermission_new();
void* QBluetoothPermission_new2(void* other);
void QBluetoothPermission_setCommunicationModes(void* self, int modes);
int QBluetoothPermission_communicationModes(void* self);
void* QBluetoothPermission_OperatorAssign(void* self, void* other);
void QBluetoothPermission_swap(void* self, void* other);
void QBluetoothPermission_delete(void* self);

#ifdef __cplusplus
}
#endif
