#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionTabWidgetFrame_base(void* self);
void* QStyleOptionTabWidgetFrame_new();
void* QStyleOptionTabWidgetFrame_new2(void* other);
void* QStyleOptionTabWidgetFrame_OperatorAssign(void* self, void* arg1);
int QStyleOptionTabWidgetFrame_lineWidth_get(void* self);
void QStyleOptionTabWidgetFrame_lineWidth_set(void* self, int lineWidth);
int QStyleOptionTabWidgetFrame_midLineWidth_get(void* self);
void QStyleOptionTabWidgetFrame_midLineWidth_set(void* self, int midLineWidth);
int QStyleOptionTabWidgetFrame_shape_get(void* self);
void QStyleOptionTabWidgetFrame_shape_set(void* self, int shape);
void* QStyleOptionTabWidgetFrame_tabBarSize_get(void* self);
void QStyleOptionTabWidgetFrame_tabBarSize_set(void* self, void* tabBarSize);
void* QStyleOptionTabWidgetFrame_rightCornerWidgetSize_get(void* self);
void QStyleOptionTabWidgetFrame_rightCornerWidgetSize_set(void* self, void* rightCornerWidgetSize);
void* QStyleOptionTabWidgetFrame_leftCornerWidgetSize_get(void* self);
void QStyleOptionTabWidgetFrame_leftCornerWidgetSize_set(void* self, void* leftCornerWidgetSize);
void* QStyleOptionTabWidgetFrame_tabBarRect_get(void* self);
void QStyleOptionTabWidgetFrame_tabBarRect_set(void* self, void* tabBarRect);
void* QStyleOptionTabWidgetFrame_selectedTabRect_get(void* self);
void QStyleOptionTabWidgetFrame_selectedTabRect_set(void* self, void* selectedTabRect);
void QStyleOptionTabWidgetFrame_delete(void* self);

#ifdef __cplusplus
}
#endif
