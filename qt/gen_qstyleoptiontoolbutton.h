#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionToolButton_base(void* self);
void QStyleOptionToolButton_chain(void* self, void** out);
void* QStyleOptionToolButton_new();
void* QStyleOptionToolButton_new2(void* other);
void* QStyleOptionToolButton_OperatorAssign(void* self, void* arg1);
int QStyleOptionToolButton_features_get(void* self);
void QStyleOptionToolButton_features_set(void* self, int features);
void* QStyleOptionToolButton_icon_get(void* self);
void QStyleOptionToolButton_icon_set(void* self, void* icon);
void* QStyleOptionToolButton_iconSize_get(void* self);
void QStyleOptionToolButton_iconSize_set(void* self, void* iconSize);
char* QStyleOptionToolButton_text_get(void* self, long long* __out_len);
void QStyleOptionToolButton_text_set(void* self, const char* text, long long text_len);
int QStyleOptionToolButton_arrowType_get(void* self);
void QStyleOptionToolButton_arrowType_set(void* self, int arrowType);
int QStyleOptionToolButton_toolButtonStyle_get(void* self);
void QStyleOptionToolButton_toolButtonStyle_set(void* self, int toolButtonStyle);
void* QStyleOptionToolButton_pos_get(void* self);
void QStyleOptionToolButton_pos_set(void* self, void* pos);
void* QStyleOptionToolButton_font_get(void* self);
void QStyleOptionToolButton_font_set(void* self, void* font);
void QStyleOptionToolButton_delete(void* self);

#ifdef __cplusplus
}
#endif
