#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QItemSelectionRange_new();
void* QItemSelectionRange_new2(void* topL, void* bottomR);
void* QItemSelectionRange_new3(void* index);
void* QItemSelectionRange_new4(void* arg1);
void QItemSelectionRange_swap(void* self, void* other);
int QItemSelectionRange_top(void* self);
int QItemSelectionRange_left(void* self);
int QItemSelectionRange_bottom(void* self);
int QItemSelectionRange_right(void* self);
int QItemSelectionRange_width(void* self);
int QItemSelectionRange_height(void* self);
void* QItemSelectionRange_topLeft(void* self);
void* QItemSelectionRange_bottomRight(void* self);
void* QItemSelectionRange_parent(void* self);
void* QItemSelectionRange_model(void* self);
int QItemSelectionRange_contains(void* self, void* index);
int QItemSelectionRange_contains2(void* self, int row, int column, void* parentIndex);
int QItemSelectionRange_intersects(void* self, void* other);
void* QItemSelectionRange_intersected(void* self, void* other);
int QItemSelectionRange_isValid(void* self);
int QItemSelectionRange_isEmpty(void* self);
void** QItemSelectionRange_indexes(void* self, long long* __out_len);
void* QItemSelectionRange_OperatorAssign(void* self, void* arg1);
int QItemSelectionRange_OperatorEqual(void* self, void* rhs);
int QItemSelectionRange_OperatorNotEqual(void* self, void* rhs);
void QItemSelectionRange_delete(void* self);

#ifdef __cplusplus
}
#endif
