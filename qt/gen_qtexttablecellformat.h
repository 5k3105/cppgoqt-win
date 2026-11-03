#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextTableCellFormat_base(void* self);
void QTextTableCellFormat_chain(void* self, void** out);
void* QTextTableCellFormat_new();
void* QTextTableCellFormat_new2(void* arg1);
int QTextTableCellFormat_isValid(void* self);
void QTextTableCellFormat_setTopPadding(void* self, double padding);
double QTextTableCellFormat_topPadding(void* self);
void QTextTableCellFormat_setBottomPadding(void* self, double padding);
double QTextTableCellFormat_bottomPadding(void* self);
void QTextTableCellFormat_setLeftPadding(void* self, double padding);
double QTextTableCellFormat_leftPadding(void* self);
void QTextTableCellFormat_setRightPadding(void* self, double padding);
double QTextTableCellFormat_rightPadding(void* self);
void QTextTableCellFormat_setPadding(void* self, double padding);
void QTextTableCellFormat_setTopBorder(void* self, double width);
double QTextTableCellFormat_topBorder(void* self);
void QTextTableCellFormat_setBottomBorder(void* self, double width);
double QTextTableCellFormat_bottomBorder(void* self);
void QTextTableCellFormat_setLeftBorder(void* self, double width);
double QTextTableCellFormat_leftBorder(void* self);
void QTextTableCellFormat_setRightBorder(void* self, double width);
double QTextTableCellFormat_rightBorder(void* self);
void QTextTableCellFormat_setBorder(void* self, double width);
void QTextTableCellFormat_setTopBorderStyle(void* self, int style);
int QTextTableCellFormat_topBorderStyle(void* self);
void QTextTableCellFormat_setBottomBorderStyle(void* self, int style);
int QTextTableCellFormat_bottomBorderStyle(void* self);
void QTextTableCellFormat_setLeftBorderStyle(void* self, int style);
int QTextTableCellFormat_leftBorderStyle(void* self);
void QTextTableCellFormat_setRightBorderStyle(void* self, int style);
int QTextTableCellFormat_rightBorderStyle(void* self);
void QTextTableCellFormat_setBorderStyle(void* self, int style);
void QTextTableCellFormat_setTopBorderBrush(void* self, void* brush);
void* QTextTableCellFormat_topBorderBrush(void* self);
void QTextTableCellFormat_setBottomBorderBrush(void* self, void* brush);
void* QTextTableCellFormat_bottomBorderBrush(void* self);
void QTextTableCellFormat_setLeftBorderBrush(void* self, void* brush);
void* QTextTableCellFormat_leftBorderBrush(void* self);
void QTextTableCellFormat_setRightBorderBrush(void* self, void* brush);
void* QTextTableCellFormat_rightBorderBrush(void* self);
void QTextTableCellFormat_setBorderBrush(void* self, void* brush);
void QTextTableCellFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
