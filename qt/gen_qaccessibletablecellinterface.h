#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleTableCellInterface__set_override(void* self, int idx);
int QAccessibleTableCellInterface_isSelected(void* self);
void** QAccessibleTableCellInterface_columnHeaderCells(void* self, long long* __out_len);
void** QAccessibleTableCellInterface_rowHeaderCells(void* self, long long* __out_len);
int QAccessibleTableCellInterface_columnIndex(void* self);
int QAccessibleTableCellInterface_rowIndex(void* self);
int QAccessibleTableCellInterface_columnExtent(void* self);
int QAccessibleTableCellInterface_rowExtent(void* self);
void* QAccessibleTableCellInterface_table(void* self);
void* QAccessibleTableCellInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleTableCellInterface__Go_new(long long id);
void QAccessibleTableCellInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
