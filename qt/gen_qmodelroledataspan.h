#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QModelRoleDataSpan_new();
void* QModelRoleDataSpan_new2(void* modelRoleData);
void* QModelRoleDataSpan_new3(void* modelRoleData, long long len_);
void* QModelRoleDataSpan_new4(void* arg1);
long long QModelRoleDataSpan_size(void* self);
long long QModelRoleDataSpan_length(void* self);
void* QModelRoleDataSpan_data(void* self);
void* QModelRoleDataSpan_begin(void* self);
void* QModelRoleDataSpan_end(void* self);
void* QModelRoleDataSpan_OperatorSubscript(void* self, long long index);
void* QModelRoleDataSpan_dataForRole(void* self, int role);
void QModelRoleDataSpan_delete(void* self);

#ifdef __cplusplus
}
#endif
