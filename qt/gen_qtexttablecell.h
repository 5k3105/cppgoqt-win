#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextTableCell_new();
void* QTextTableCell_new2(void* o);
void* QTextTableCell_OperatorAssign(void* self, void* o);
void QTextTableCell_setFormat(void* self, void* format);
void* QTextTableCell_format(void* self);
int QTextTableCell_row(void* self);
int QTextTableCell_column(void* self);
int QTextTableCell_rowSpan(void* self);
int QTextTableCell_columnSpan(void* self);
int QTextTableCell_isValid(void* self);
void* QTextTableCell_firstCursorPosition(void* self);
void* QTextTableCell_lastCursorPosition(void* self);
int QTextTableCell_firstPosition(void* self);
int QTextTableCell_lastPosition(void* self);
int QTextTableCell_OperatorEqual(void* self, void* other);
int QTextTableCell_OperatorNotEqual(void* self, void* other);
int QTextTableCell_tableCellFormatIndex(void* self);
void QTextTableCell_delete(void* self);

#ifdef __cplusplus
}
#endif
