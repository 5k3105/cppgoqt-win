#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSize_new();
void* QSize_new2(int w, int h);
void* QSize_new3(void* arg1);
int QSize_isNull(void* self);
int QSize_isEmpty(void* self);
int QSize_isValid(void* self);
int QSize_width(void* self);
int QSize_height(void* self);
void QSize_setWidth(void* self, int w);
void QSize_setHeight(void* self, int h);
void QSize_transpose(void* self);
void* QSize_transposed(void* self);
void QSize_scale(void* self, int w, int h, int mode);
void QSize_scale2(void* self, void* s, int mode);
void* QSize_scaled(void* self, int w, int h, int mode);
void* QSize_scaled2(void* self, void* s, int mode);
void* QSize_expandedTo(void* self, void* arg1);
void* QSize_boundedTo(void* self, void* arg1);
void* QSize_grownBy(void* self, void* m);
void* QSize_shrunkBy(void* self, void* m);
void* QSize_OperatorPlusAssign(void* self, void* arg1);
void* QSize_OperatorMinusAssign(void* self, void* arg1);
void* QSize_OperatorMultiplyAssign(void* self, double c);
void* QSize_OperatorDivideAssign(void* self, double c);
void* QSize_toSizeF(void* self);
void* QSize_OperatorAssign(void* self, void* arg1);
int QSize_OperatorEqual(void* self, void* rhs);
int QSize_OperatorNotEqual(void* self, void* rhs);
void* QSize_OperatorPlus(void* self, void* s2);
void* QSize_OperatorMinus(void* self, void* s2);
void* QSize_OperatorMultiply(void* self, double c);
void* QSize_OperatorDivide(void* self, double c);
int QSize_OperatorEqual2(void* self, void* rhs);
int QSize_OperatorNotEqual2(void* self, void* rhs);
void QSize_delete(void* self);

#ifdef __cplusplus
}
#endif
