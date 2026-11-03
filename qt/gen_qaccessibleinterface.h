#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QAccessibleInterface_isValid(void* self);
void* QAccessibleInterface_object(void* self);
void* QAccessibleInterface_window(void* self);
void* QAccessibleInterface_focusChild(void* self);
void* QAccessibleInterface_childAt(void* self, int x, int y);
void* QAccessibleInterface_parent(void* self);
void* QAccessibleInterface_child(void* self, int index);
int QAccessibleInterface_childCount(void* self);
int QAccessibleInterface_indexOfChild(void* self, void* arg1);
char* QAccessibleInterface_text(void* self, int t, long long* __out_len);
void QAccessibleInterface_setText(void* self, int t, const char* text, long long text_len);
void* QAccessibleInterface_rect(void* self);
int QAccessibleInterface_role(void* self);
void* QAccessibleInterface_state(void* self);
void* QAccessibleInterface_foregroundColor(void* self);
void* QAccessibleInterface_backgroundColor(void* self);
void* QAccessibleInterface_textInterface(void* self);
void* QAccessibleInterface_editableTextInterface(void* self);
void* QAccessibleInterface_valueInterface(void* self);
void* QAccessibleInterface_actionInterface(void* self);
void* QAccessibleInterface_imageInterface(void* self);
void* QAccessibleInterface_tableInterface(void* self);
void* QAccessibleInterface_tableCellInterface(void* self);
void* QAccessibleInterface_hyperlinkInterface(void* self);
void* QAccessibleInterface_selectionInterface(void* self);
void* QAccessibleInterface_attributesInterface(void* self);
void* QAccessibleInterface_OperatorAssign(void* self, void* arg1);
void QAccessibleInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
