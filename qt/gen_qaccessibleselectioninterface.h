#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleSelectionInterface__set_override(void* self, int idx);
int QAccessibleSelectionInterface_selectedItemCount(void* self);
void** QAccessibleSelectionInterface_selectedItems(void* self, long long* __out_len);
void* QAccessibleSelectionInterface_selectedItem(void* self, int selectionIndex);
int QAccessibleSelectionInterface_isSelected(void* self, void* childItem);
int QAccessibleSelectionInterface_select(void* self, void* childItem);
int QAccessibleSelectionInterface_unselect(void* self, void* childItem);
int QAccessibleSelectionInterface_selectAll(void* self);
int QAccessibleSelectionInterface_clear(void* self);
void* QAccessibleSelectionInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleSelectionInterface_selectedItem_default(void* self, int selectionIndex);
int QAccessibleSelectionInterface_isSelected_default(void* self, void* childItem);
void* QAccessibleSelectionInterface__Go_new(long long id);
void QAccessibleSelectionInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
