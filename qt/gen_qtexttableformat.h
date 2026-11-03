#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextTableFormat_base(void* self);
void QTextTableFormat_chain(void* self, void** out);
void* QTextTableFormat_new();
void* QTextTableFormat_new2(void* arg1);
int QTextTableFormat_isValid(void* self);
int QTextTableFormat_columns(void* self);
void QTextTableFormat_setColumns(void* self, int columns);
void QTextTableFormat_setColumnWidthConstraints(void* self, void** constraints, long long constraints_len);
void** QTextTableFormat_columnWidthConstraints(void* self, long long* __out_len);
void QTextTableFormat_clearColumnWidthConstraints(void* self);
double QTextTableFormat_cellSpacing(void* self);
void QTextTableFormat_setCellSpacing(void* self, double spacing);
double QTextTableFormat_cellPadding(void* self);
void QTextTableFormat_setCellPadding(void* self, double padding);
void QTextTableFormat_setAlignment(void* self, int alignment);
int QTextTableFormat_alignment(void* self);
void QTextTableFormat_setHeaderRowCount(void* self, int count);
int QTextTableFormat_headerRowCount(void* self);
void QTextTableFormat_setBorderCollapse(void* self, int borderCollapse);
int QTextTableFormat_borderCollapse(void* self);
void QTextTableFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
