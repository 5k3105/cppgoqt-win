#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVector2D_new();
void* QVector2D_new2(int arg1);
void* QVector2D_new3(float xpos, float ypos);
void* QVector2D_new4(void* point);
void* QVector2D_new5(void* point);
void* QVector2D_new6(void* vector);
void* QVector2D_new7(void* vector);
void* QVector2D_new8(void* arg1);
int QVector2D_isNull(void* self);
float QVector2D_x(void* self);
float QVector2D_y(void* self);
void QVector2D_setX(void* self, float x);
void QVector2D_setY(void* self, float y);
float QVector2D_OperatorSubscript(void* self, int i);
float QVector2D_length(void* self);
float QVector2D_lengthSquared(void* self);
void* QVector2D_normalized(void* self);
void QVector2D_normalize(void* self);
float QVector2D_distanceToPoint(void* self, void* point);
float QVector2D_distanceToLine(void* self, void* point, void* direction);
void* QVector2D_OperatorPlusAssign(void* self, void* vector);
void* QVector2D_OperatorMinusAssign(void* self, void* vector);
void* QVector2D_OperatorMultiplyAssign(void* self, float factor);
void* QVector2D_OperatorMultiplyAssign2(void* self, void* vector);
void* QVector2D_OperatorDivideAssign(void* self, float divisor);
void* QVector2D_OperatorDivideAssign2(void* self, void* vector);
float QVector2D_dotProduct(void* v1, void* v2);
void* QVector2D_toVector3D(void* self);
void* QVector2D_toVector4D(void* self);
void* QVector2D_toPoint(void* self);
void* QVector2D_toPointF(void* self);
void* QVector2D_ToQVariant(void* self);
void* QVector2D_OperatorAssign(void* self, void* arg1);
int QVector2D_OperatorEqual(void* self, void* v2);
int QVector2D_OperatorNotEqual(void* self, void* v2);
void* QVector2D_OperatorPlus(void* self, void* v2);
void* QVector2D_OperatorMinus(void* self, void* v2);
void* QVector2D_OperatorMultiply(void* self, float factor);
void* QVector2D_OperatorMultiply2(void* self, void* v2);
void* QVector2D_OperatorMinus2(void* self);
void* QVector2D_OperatorDivide(void* self, float divisor);
void* QVector2D_OperatorDivide2(void* self, void* divisor);
void QVector2D_delete(void* self);

#ifdef __cplusplus
}
#endif
