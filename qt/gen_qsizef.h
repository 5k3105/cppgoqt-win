#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSizeF_new();
void* QSizeF_new2(void* sz);
void* QSizeF_new3(double w, double h);
void* QSizeF_new4(void* arg1);
int QSizeF_isNull(void* self);
int QSizeF_isEmpty(void* self);
int QSizeF_isValid(void* self);
double QSizeF_width(void* self);
double QSizeF_height(void* self);
void QSizeF_setWidth(void* self, double w);
void QSizeF_setHeight(void* self, double h);
void QSizeF_transpose(void* self);
void* QSizeF_transposed(void* self);
void QSizeF_scale(void* self, double w, double h, int mode);
void QSizeF_scale2(void* self, void* s, int mode);
void* QSizeF_scaled(void* self, double w, double h, int mode);
void* QSizeF_scaled2(void* self, void* s, int mode);
void* QSizeF_expandedTo(void* self, void* arg1);
void* QSizeF_boundedTo(void* self, void* arg1);
void* QSizeF_grownBy(void* self, void* m);
void* QSizeF_shrunkBy(void* self, void* m);
void* QSizeF_OperatorPlusAssign(void* self, void* arg1);
void* QSizeF_OperatorMinusAssign(void* self, void* arg1);
void* QSizeF_OperatorMultiplyAssign(void* self, double c);
void* QSizeF_OperatorDivideAssign(void* self, double c);
void* QSizeF_toSize(void* self);
void* QSizeF_OperatorAssign(void* self, void* arg1);
int QSizeF_OperatorEqual(void* self, void* rhs);
int QSizeF_OperatorNotEqual(void* self, void* rhs);
int QSizeF_OperatorEqual2(void* self, void* rhs);
int QSizeF_OperatorNotEqual2(void* self, void* rhs);
void* QSizeF_OperatorPlus(void* self, void* s2);
void* QSizeF_OperatorMinus(void* self, void* s2);
void* QSizeF_OperatorMultiply(void* self, double c);
void* QSizeF_OperatorDivide(void* self, double c);
void QSizeF_delete(void* self);

#ifdef __cplusplus
}
#endif
