#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPoint_new();
void* QPoint_new2(int xpos, int ypos);
void* QPoint_new3(void* arg1);
int QPoint_isNull(void* self);
int QPoint_x(void* self);
int QPoint_y(void* self);
void QPoint_setX(void* self, int x);
void QPoint_setY(void* self, int y);
int QPoint_manhattanLength(void* self);
void* QPoint_transposed(void* self);
void* QPoint_OperatorPlusAssign(void* self, void* p);
void* QPoint_OperatorMinusAssign(void* self, void* p);
void* QPoint_OperatorMultiplyAssign(void* self, float factor);
void* QPoint_OperatorMultiplyAssign2(void* self, double factor);
void* QPoint_OperatorMultiplyAssign3(void* self, int factor);
void* QPoint_OperatorDivideAssign(void* self, double divisor);
int QPoint_dotProduct(void* p1, void* p2);
void* QPoint_toPointF(void* self);
void* QPoint_OperatorAssign(void* self, void* arg1);
int QPoint_OperatorEqual(void* self, void* rhs);
int QPoint_OperatorNotEqual(void* self, void* rhs);
void* QPoint_OperatorPlus(void* self, void* p2);
void* QPoint_OperatorMinus(void* self, void* p2);
void* QPoint_OperatorMultiply(void* self, float factor);
void* QPoint_OperatorMultiply2(void* self, double factor);
void* QPoint_OperatorMultiply3(void* self, int factor);
void* QPoint_OperatorPlus2(void* self);
void* QPoint_OperatorMinus2(void* self);
void* QPoint_OperatorDivide(void* self, double c);
int QPoint_OperatorEqual2(void* self, void* rhs);
int QPoint_OperatorNotEqual2(void* self, void* rhs);
void* QPoint_OperatorMultiply4(void* self, void* m);
void* QPoint_OperatorMultiply5(void* self, void* matrix);
void QPoint_delete(void* self);

#ifdef __cplusplus
}
#endif
