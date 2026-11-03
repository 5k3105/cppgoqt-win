#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMarginsF_new();
void* QMarginsF_new2(double left, double top, double right, double bottom);
void* QMarginsF_new3(void* margins);
void* QMarginsF_new4(void* arg1);
int QMarginsF_isNull(void* self);
double QMarginsF_left(void* self);
double QMarginsF_top(void* self);
double QMarginsF_right(void* self);
double QMarginsF_bottom(void* self);
void QMarginsF_setLeft(void* self, double aleft);
void QMarginsF_setTop(void* self, double atop);
void QMarginsF_setRight(void* self, double aright);
void QMarginsF_setBottom(void* self, double abottom);
void* QMarginsF_OperatorPlusAssign(void* self, void* margins);
void* QMarginsF_OperatorMinusAssign(void* self, void* margins);
void* QMarginsF_OperatorPlusAssign2(void* self, double addend);
void* QMarginsF_OperatorMinusAssign2(void* self, double subtrahend);
void* QMarginsF_OperatorMultiplyAssign(void* self, double factor);
void* QMarginsF_OperatorDivideAssign(void* self, double divisor);
void* QMarginsF_toMargins(void* self);
void* QMarginsF_OperatorAssign(void* self, void* arg1);
int QMarginsF_OperatorEqual(void* self, void* rhs);
int QMarginsF_OperatorNotEqual(void* self, void* rhs);
int QMarginsF_OperatorEqual2(void* self, void* rhs);
int QMarginsF_OperatorNotEqual2(void* self, void* rhs);
void* QMarginsF_OperatorPlus(void* self, void* rhs);
void* QMarginsF_OperatorMinus(void* self, void* rhs);
void* QMarginsF_OperatorPlus2(void* self, double rhs);
void* QMarginsF_OperatorMinus2(void* self, double rhs);
void* QMarginsF_OperatorMultiply(void* self, double rhs);
void* QMarginsF_OperatorDivide(void* self, double divisor);
void* QMarginsF_OperatorBitwiseOr(void* self, void* m2);
void* QMarginsF_OperatorPlus3(void* self);
void* QMarginsF_OperatorMinus3(void* self);
void* QMarginsF_OperatorPlus4(void* self, void* rhs);
void QMarginsF_delete(void* self);

#ifdef __cplusplus
}
#endif
