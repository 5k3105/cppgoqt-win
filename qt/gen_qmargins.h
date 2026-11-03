#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMargins_new();
void* QMargins_new2(int left, int top, int right, int bottom);
void* QMargins_new3(void* arg1);
int QMargins_isNull(void* self);
int QMargins_left(void* self);
int QMargins_top(void* self);
int QMargins_right(void* self);
int QMargins_bottom(void* self);
void QMargins_setLeft(void* self, int left);
void QMargins_setTop(void* self, int top);
void QMargins_setRight(void* self, int right);
void QMargins_setBottom(void* self, int bottom);
void* QMargins_OperatorPlusAssign(void* self, void* margins);
void* QMargins_OperatorMinusAssign(void* self, void* margins);
void* QMargins_OperatorPlusAssign2(void* self, int arg1);
void* QMargins_OperatorMinusAssign2(void* self, int arg1);
void* QMargins_OperatorMultiplyAssign(void* self, int arg1);
void* QMargins_OperatorDivideAssign(void* self, int arg1);
void* QMargins_OperatorMultiplyAssign2(void* self, double arg1);
void* QMargins_OperatorDivideAssign2(void* self, double arg1);
void* QMargins_toMarginsF(void* self);
void* QMargins_OperatorAssign(void* self, void* arg1);
void* QMargins_OperatorPlus(void* self, void* m2);
void* QMargins_OperatorMinus(void* self, void* m2);
void* QMargins_OperatorPlus2(void* self, int rhs);
void* QMargins_OperatorMinus2(void* self, int rhs);
void* QMargins_OperatorMultiply(void* self, int factor);
void* QMargins_OperatorMultiply2(void* self, double factor);
void* QMargins_OperatorDivide(void* self, int divisor);
void* QMargins_OperatorDivide2(void* self, double divisor);
void* QMargins_OperatorBitwiseOr(void* self, void* m2);
int QMargins_OperatorEqual(void* self, void* rhs);
int QMargins_OperatorNotEqual(void* self, void* rhs);
void* QMargins_OperatorPlus3(void* self);
void* QMargins_OperatorMinus3(void* self);
int QMargins_OperatorEqual2(void* self, void* rhs);
int QMargins_OperatorNotEqual2(void* self, void* rhs);
void* QMargins_OperatorPlus4(void* self, void* rectangle);
void QMargins_delete(void* self);

#ifdef __cplusplus
}
#endif
