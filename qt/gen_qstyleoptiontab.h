#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionTab_base(void* self);
void* QStyleOptionTab_new();
void* QStyleOptionTab_new2(void* other);
void* QStyleOptionTab_OperatorAssign(void* self, void* arg1);
int QStyleOptionTab_shape_get(void* self);
void QStyleOptionTab_shape_set(void* self, int shape);
char* QStyleOptionTab_text_get(void* self, long long* __out_len);
void QStyleOptionTab_text_set(void* self, const char* text, long long text_len);
void* QStyleOptionTab_icon_get(void* self);
void QStyleOptionTab_icon_set(void* self, void* icon);
int QStyleOptionTab_row_get(void* self);
void QStyleOptionTab_row_set(void* self, int row);
int QStyleOptionTab_position_get(void* self);
void QStyleOptionTab_position_set(void* self, int position);
int QStyleOptionTab_selectedPosition_get(void* self);
void QStyleOptionTab_selectedPosition_set(void* self, int selectedPosition);
int QStyleOptionTab_cornerWidgets_get(void* self);
void QStyleOptionTab_cornerWidgets_set(void* self, int cornerWidgets);
void* QStyleOptionTab_iconSize_get(void* self);
void QStyleOptionTab_iconSize_set(void* self, void* iconSize);
int QStyleOptionTab_documentMode_get(void* self);
void QStyleOptionTab_documentMode_set(void* self, int documentMode);
void* QStyleOptionTab_leftButtonSize_get(void* self);
void QStyleOptionTab_leftButtonSize_set(void* self, void* leftButtonSize);
void* QStyleOptionTab_rightButtonSize_get(void* self);
void QStyleOptionTab_rightButtonSize_set(void* self, void* rightButtonSize);
int QStyleOptionTab_features_get(void* self);
void QStyleOptionTab_features_set(void* self, int features);
int QStyleOptionTab_tabIndex_get(void* self);
void QStyleOptionTab_tabIndex_set(void* self, int tabIndex);
void QStyleOptionTab_delete(void* self);

#ifdef __cplusplus
}
#endif
