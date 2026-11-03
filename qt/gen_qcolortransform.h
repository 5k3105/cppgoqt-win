#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QColorTransform_new();
void* QColorTransform_new2(void* colorTransform);
void* QColorTransform_OperatorAssign(void* self, void* other);
void QColorTransform_swap(void* self, void* other);
int QColorTransform_isIdentity(void* self);
unsigned int QColorTransform_map(void* self, unsigned int argb);
void* QColorTransform_map2(void* self, void* rgba64);
void* QColorTransform_map3(void* self, void* color);
int QColorTransform_OperatorEqual(void* self, void* ct2);
int QColorTransform_OperatorNotEqual(void* self, void* ct2);
void QColorTransform_delete(void* self);

#ifdef __cplusplus
}
#endif
