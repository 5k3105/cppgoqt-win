#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTableWidgetSelectionRange_new();
void* QTableWidgetSelectionRange_new2(int top, int left, int bottom, int right);
int QTableWidgetSelectionRange_topRow(void* self);
int QTableWidgetSelectionRange_bottomRow(void* self);
int QTableWidgetSelectionRange_leftColumn(void* self);
int QTableWidgetSelectionRange_rightColumn(void* self);
int QTableWidgetSelectionRange_rowCount(void* self);
int QTableWidgetSelectionRange_columnCount(void* self);
int QTableWidgetSelectionRange_OperatorEqual(void* self, void* rhs);
int QTableWidgetSelectionRange_OperatorNotEqual(void* self, void* rhs);
void QTableWidgetSelectionRange_delete(void* self);

#ifdef __cplusplus
}
#endif
