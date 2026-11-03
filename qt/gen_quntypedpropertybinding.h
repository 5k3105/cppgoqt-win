#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUntypedPropertyBinding_new();
void* QUntypedPropertyBinding_new2(void* other);
void* QUntypedPropertyBinding_OperatorAssign(void* self, void* other);
int QUntypedPropertyBinding_isNull(void* self);
void* QUntypedPropertyBinding_error(void* self);
void* QUntypedPropertyBinding_valueMetaType(void* self);
void QUntypedPropertyBinding_delete(void* self);

#ifdef __cplusplus
}
#endif
