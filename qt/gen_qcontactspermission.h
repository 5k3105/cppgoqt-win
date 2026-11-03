#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QContactsPermission_new();
void* QContactsPermission_new2(void* other);
void QContactsPermission_setAccessMode(void* self, int mode);
int QContactsPermission_accessMode(void* self);
void* QContactsPermission_OperatorAssign(void* self, void* other);
void QContactsPermission_swap(void* self, void* other);
void QContactsPermission_delete(void* self);

#ifdef __cplusplus
}
#endif
