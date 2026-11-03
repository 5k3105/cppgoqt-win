#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionFrame_base(void* self);
void* QStyleOptionFrame_new();
void* QStyleOptionFrame_new2(void* other);
void* QStyleOptionFrame_OperatorAssign(void* self, void* arg1);
int QStyleOptionFrame_lineWidth_get(void* self);
void QStyleOptionFrame_lineWidth_set(void* self, int lineWidth);
int QStyleOptionFrame_midLineWidth_get(void* self);
void QStyleOptionFrame_midLineWidth_set(void* self, int midLineWidth);
int QStyleOptionFrame_features_get(void* self);
void QStyleOptionFrame_features_set(void* self, int features);
int QStyleOptionFrame_frameShape_get(void* self);
void QStyleOptionFrame_frameShape_set(void* self, int frameShape);
void QStyleOptionFrame_delete(void* self);

#ifdef __cplusplus
}
#endif
