#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionViewItem_base(void* self);
void* QStyleOptionViewItem_new();
void* QStyleOptionViewItem_new2(void* other);
void* QStyleOptionViewItem_OperatorAssign(void* self, void* arg1);
int QStyleOptionViewItem_displayAlignment_get(void* self);
void QStyleOptionViewItem_displayAlignment_set(void* self, int displayAlignment);
int QStyleOptionViewItem_decorationAlignment_get(void* self);
void QStyleOptionViewItem_decorationAlignment_set(void* self, int decorationAlignment);
int QStyleOptionViewItem_textElideMode_get(void* self);
void QStyleOptionViewItem_textElideMode_set(void* self, int textElideMode);
int QStyleOptionViewItem_decorationPosition_get(void* self);
void QStyleOptionViewItem_decorationPosition_set(void* self, int decorationPosition);
void* QStyleOptionViewItem_decorationSize_get(void* self);
void QStyleOptionViewItem_decorationSize_set(void* self, void* decorationSize);
void* QStyleOptionViewItem_font_get(void* self);
void QStyleOptionViewItem_font_set(void* self, void* font);
int QStyleOptionViewItem_showDecorationSelected_get(void* self);
void QStyleOptionViewItem_showDecorationSelected_set(void* self, int showDecorationSelected);
int QStyleOptionViewItem_features_get(void* self);
void QStyleOptionViewItem_features_set(void* self, int features);
void* QStyleOptionViewItem_locale_get(void* self);
void QStyleOptionViewItem_locale_set(void* self, void* locale);
void* QStyleOptionViewItem_widget_get(void* self);
void QStyleOptionViewItem_widget_set(void* self, void* widget);
void* QStyleOptionViewItem_index_get(void* self);
void QStyleOptionViewItem_index_set(void* self, void* index);
int QStyleOptionViewItem_checkState_get(void* self);
void QStyleOptionViewItem_checkState_set(void* self, int checkState);
void* QStyleOptionViewItem_icon_get(void* self);
void QStyleOptionViewItem_icon_set(void* self, void* icon);
char* QStyleOptionViewItem_text_get(void* self, long long* __out_len);
void QStyleOptionViewItem_text_set(void* self, const char* text, long long text_len);
int QStyleOptionViewItem_viewItemPosition_get(void* self);
void QStyleOptionViewItem_viewItemPosition_set(void* self, int viewItemPosition);
void* QStyleOptionViewItem_backgroundBrush_get(void* self);
void QStyleOptionViewItem_backgroundBrush_set(void* self, void* backgroundBrush);
void QStyleOptionViewItem_delete(void* self);

#ifdef __cplusplus
}
#endif
