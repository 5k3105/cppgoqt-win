#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QInputMethodEvent__Attribute_new(int typ, int s, int l, void* val);
void* QInputMethodEvent__Attribute_new2(int typ, int s, int l);
void* QInputMethodEvent__Attribute_new3(void* arg1);
void* QInputMethodEvent__Attribute_OperatorAssign(void* self, void* arg1);
int QInputMethodEvent__Attribute_OperatorEqual(void* self, void* rhs);
int QInputMethodEvent__Attribute_OperatorNotEqual(void* self, void* rhs);
int QInputMethodEvent__Attribute_type_get(void* self);
void QInputMethodEvent__Attribute_type_set(void* self, int type_);
int QInputMethodEvent__Attribute_start_get(void* self);
void QInputMethodEvent__Attribute_start_set(void* self, int start);
int QInputMethodEvent__Attribute_length_get(void* self);
void QInputMethodEvent__Attribute_length_set(void* self, int length);
void* QInputMethodEvent__Attribute_value_get(void* self);
void QInputMethodEvent__Attribute_value_set(void* self, void* value);
void QInputMethodEvent__Attribute_delete(void* self);

#ifdef __cplusplus
}
#endif
