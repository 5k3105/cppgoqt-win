#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPointF_new();
void* QPointF_new2(void* p);
void* QPointF_new3(double xpos, double ypos);
void* QPointF_new4(void* arg1);
double QPointF_manhattanLength(void* self);
int QPointF_isNull(void* self);
double QPointF_x(void* self);
double QPointF_y(void* self);
void QPointF_setX(void* self, double x);
void QPointF_setY(void* self, double y);
void* QPointF_transposed(void* self);
void* QPointF_OperatorPlusAssign(void* self, void* p);
void* QPointF_OperatorMinusAssign(void* self, void* p);
void* QPointF_OperatorMultiplyAssign(void* self, double c);
void* QPointF_OperatorDivideAssign(void* self, double c);
double QPointF_dotProduct(void* p1, void* p2);
void* QPointF_toPoint(void* self);
void* QPointF_OperatorAssign(void* self, void* arg1);
int QPointF_OperatorEqual(void* self, void* rhs);
int QPointF_OperatorNotEqual(void* self, void* rhs);
int QPointF_OperatorEqual2(void* self, void* rhs);
int QPointF_OperatorNotEqual2(void* self, void* rhs);
void* QPointF_OperatorPlus(void* self, void* p2);
void* QPointF_OperatorMinus(void* self, void* p2);
void* QPointF_OperatorMultiply(void* self, double c);
void* QPointF_OperatorPlus2(void* self);
void* QPointF_OperatorMinus2(void* self);
void* QPointF_OperatorDivide(void* self, double divisor);
void* QPointF_OperatorMultiply2(void* self, void* m);
void* QPointF_OperatorMultiply3(void* self, void* matrix);
void QPointF_delete(void* self);

#ifdef __cplusplus
}
#endif
