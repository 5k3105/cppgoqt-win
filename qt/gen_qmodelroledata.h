#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QModelRoleData_new(int role);
void* QModelRoleData_new2(void* arg1);
int QModelRoleData_role(void* self);
void* QModelRoleData_data(void* self);
void* QModelRoleData_data2(void* self);
void QModelRoleData_clearData(void* self);
void* QModelRoleData_OperatorAssign(void* self, void* arg1);
void QModelRoleData_delete(void* self);

#ifdef __cplusplus
}
#endif
