#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QModelIndex_new();
void* QModelIndex_new2(void* arg1);
int QModelIndex_row(void* self);
int QModelIndex_column(void* self);
unsigned long long QModelIndex_internalId(void* self);
void* QModelIndex_parent(void* self);
void* QModelIndex_sibling(void* self, int row, int column);
void* QModelIndex_siblingAtColumn(void* self, int column);
void* QModelIndex_siblingAtRow(void* self, int row);
void* QModelIndex_data(void* self);
void QModelIndex_multiData(void* self, void* roleDataSpan);
int QModelIndex_flags(void* self);
void* QModelIndex_model(void* self);
int QModelIndex_isValid(void* self);
void* QModelIndex_OperatorAssign(void* self, void* arg1);
void* QModelIndex_data2(void* self, int role);
int QModelIndex_OperatorEqual(void* self, void* rhs);
int QModelIndex_OperatorNotEqual(void* self, void* rhs);
int QModelIndex_OperatorLesser(void* self, void* rhs);
int QModelIndex_OperatorGreater(void* self, void* rhs);
int QModelIndex_OperatorLesserOrEqual(void* self, void* rhs);
int QModelIndex_OperatorGreaterOrEqual(void* self, void* rhs);
int QModelIndex_OperatorEqual2(void* self, void* rhs);
int QModelIndex_OperatorNotEqual2(void* self, void* rhs);
int QModelIndex_OperatorLesser2(void* self, void* rhs);
int QModelIndex_OperatorGreater2(void* self, void* rhs);
int QModelIndex_OperatorLesserOrEqual2(void* self, void* rhs);
int QModelIndex_OperatorGreaterOrEqual2(void* self, void* rhs);
void QModelIndex_delete(void* self);

#ifdef __cplusplus
}
#endif
