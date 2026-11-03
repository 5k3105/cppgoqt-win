#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionMenuItem_base(void* self);
void* QStyleOptionMenuItem_new();
void* QStyleOptionMenuItem_new2(void* other);
void* QStyleOptionMenuItem_OperatorAssign(void* self, void* arg1);
int QStyleOptionMenuItem_menuItemType_get(void* self);
void QStyleOptionMenuItem_menuItemType_set(void* self, int menuItemType);
int QStyleOptionMenuItem_checkType_get(void* self);
void QStyleOptionMenuItem_checkType_set(void* self, int checkType);
int QStyleOptionMenuItem_checked_get(void* self);
void QStyleOptionMenuItem_checked_set(void* self, int checked);
int QStyleOptionMenuItem_menuHasCheckableItems_get(void* self);
void QStyleOptionMenuItem_menuHasCheckableItems_set(void* self, int menuHasCheckableItems);
void* QStyleOptionMenuItem_menuRect_get(void* self);
void QStyleOptionMenuItem_menuRect_set(void* self, void* menuRect);
char* QStyleOptionMenuItem_text_get(void* self, long long* __out_len);
void QStyleOptionMenuItem_text_set(void* self, const char* text, long long text_len);
void* QStyleOptionMenuItem_icon_get(void* self);
void QStyleOptionMenuItem_icon_set(void* self, void* icon);
int QStyleOptionMenuItem_maxIconWidth_get(void* self);
void QStyleOptionMenuItem_maxIconWidth_set(void* self, int maxIconWidth);
int QStyleOptionMenuItem_reservedShortcutWidth_get(void* self);
void QStyleOptionMenuItem_reservedShortcutWidth_set(void* self, int reservedShortcutWidth);
void* QStyleOptionMenuItem_font_get(void* self);
void QStyleOptionMenuItem_font_set(void* self, void* font);
void QStyleOptionMenuItem_delete(void* self);

#ifdef __cplusplus
}
#endif
