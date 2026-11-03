#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleAttributesInterface__set_override(void* self, int idx);
int* QAccessibleAttributesInterface_attributeKeys(void* self, long long* __out_len);
void* QAccessibleAttributesInterface_attributeValue(void* self, int key);
void* QAccessibleAttributesInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleAttributesInterface__Go_new(long long id);
void QAccessibleAttributesInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
