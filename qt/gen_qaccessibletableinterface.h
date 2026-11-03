#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleTableInterface__set_override(void* self, int idx);
void* QAccessibleTableInterface_caption(void* self);
void* QAccessibleTableInterface_summary(void* self);
void* QAccessibleTableInterface_cellAt(void* self, int row, int column);
int QAccessibleTableInterface_selectedCellCount(void* self);
void** QAccessibleTableInterface_selectedCells(void* self, long long* __out_len);
char* QAccessibleTableInterface_columnDescription(void* self, int column, long long* __out_len);
char* QAccessibleTableInterface_rowDescription(void* self, int row, long long* __out_len);
int QAccessibleTableInterface_selectedColumnCount(void* self);
int QAccessibleTableInterface_selectedRowCount(void* self);
int QAccessibleTableInterface_columnCount(void* self);
int QAccessibleTableInterface_rowCount(void* self);
int* QAccessibleTableInterface_selectedColumns(void* self, long long* __out_len);
int* QAccessibleTableInterface_selectedRows(void* self, long long* __out_len);
int QAccessibleTableInterface_isColumnSelected(void* self, int column);
int QAccessibleTableInterface_isRowSelected(void* self, int row);
int QAccessibleTableInterface_selectRow(void* self, int row);
int QAccessibleTableInterface_selectColumn(void* self, int column);
int QAccessibleTableInterface_unselectRow(void* self, int row);
int QAccessibleTableInterface_unselectColumn(void* self, int column);
void QAccessibleTableInterface_modelChange(void* self, void* event);
void* QAccessibleTableInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleTableInterface__Go_new(long long id);
void QAccessibleTableInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
