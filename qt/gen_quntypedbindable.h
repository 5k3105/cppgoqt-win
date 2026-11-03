#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUntypedBindable_new();
void* QUntypedBindable_new2(void* arg1);
int QUntypedBindable_isValid(void* self);
int QUntypedBindable_isBindable(void* self);
int QUntypedBindable_isReadOnly(void* self);
void* QUntypedBindable_makeBinding(void* self);
void* QUntypedBindable_takeBinding(void* self);
void QUntypedBindable_observe(void* self, void* observer);
void* QUntypedBindable_binding(void* self);
int QUntypedBindable_setBinding(void* self, void* binding);
int QUntypedBindable_hasBinding(void* self);
void* QUntypedBindable_metaType(void* self);
void* QUntypedBindable_makeBinding2(void* self, void* location);
void QUntypedBindable_delete(void* self);

#ifdef __cplusplus
}
#endif
