#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionToolBar_base(void* self);
void* QStyleOptionToolBar_new();
void* QStyleOptionToolBar_new2(void* other);
void* QStyleOptionToolBar_OperatorAssign(void* self, void* arg1);
int QStyleOptionToolBar_positionOfLine_get(void* self);
void QStyleOptionToolBar_positionOfLine_set(void* self, int positionOfLine);
int QStyleOptionToolBar_positionWithinLine_get(void* self);
void QStyleOptionToolBar_positionWithinLine_set(void* self, int positionWithinLine);
int QStyleOptionToolBar_toolBarArea_get(void* self);
void QStyleOptionToolBar_toolBarArea_set(void* self, int toolBarArea);
int QStyleOptionToolBar_features_get(void* self);
void QStyleOptionToolBar_features_set(void* self, int features);
int QStyleOptionToolBar_lineWidth_get(void* self);
void QStyleOptionToolBar_lineWidth_set(void* self, int lineWidth);
int QStyleOptionToolBar_midLineWidth_get(void* self);
void QStyleOptionToolBar_midLineWidth_set(void* self, int midLineWidth);
void QStyleOptionToolBar_delete(void* self);

#ifdef __cplusplus
}
#endif
