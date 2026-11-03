#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVector4D_new();
void* QVector4D_new2(int arg1);
void* QVector4D_new3(float xpos, float ypos, float zpos, float wpos);
void* QVector4D_new4(void* point);
void* QVector4D_new5(void* point);
void* QVector4D_new6(void* vector);
void* QVector4D_new7(void* vector, float zpos, float wpos);
void* QVector4D_new8(void* vector);
void* QVector4D_new9(void* vector, float wpos);
void* QVector4D_new10(void* arg1);
int QVector4D_isNull(void* self);
float QVector4D_x(void* self);
float QVector4D_y(void* self);
float QVector4D_z(void* self);
float QVector4D_w(void* self);
void QVector4D_setX(void* self, float x);
void QVector4D_setY(void* self, float y);
void QVector4D_setZ(void* self, float z);
void QVector4D_setW(void* self, float w);
float QVector4D_OperatorSubscript(void* self, int i);
float QVector4D_length(void* self);
float QVector4D_lengthSquared(void* self);
void* QVector4D_normalized(void* self);
void QVector4D_normalize(void* self);
void* QVector4D_OperatorPlusAssign(void* self, void* vector);
void* QVector4D_OperatorMinusAssign(void* self, void* vector);
void* QVector4D_OperatorMultiplyAssign(void* self, float factor);
void* QVector4D_OperatorMultiplyAssign2(void* self, void* vector);
void* QVector4D_OperatorDivideAssign(void* self, float divisor);
void* QVector4D_OperatorDivideAssign2(void* self, void* vector);
float QVector4D_dotProduct(void* v1, void* v2);
void* QVector4D_toVector2D(void* self);
void* QVector4D_toVector2DAffine(void* self);
void* QVector4D_toVector3D(void* self);
void* QVector4D_toVector3DAffine(void* self);
void* QVector4D_toPoint(void* self);
void* QVector4D_toPointF(void* self);
void* QVector4D_ToQVariant(void* self);
int QVector4D_OperatorEqual(void* self, void* v2);
int QVector4D_OperatorNotEqual(void* self, void* v2);
void* QVector4D_OperatorPlus(void* self, void* v2);
void* QVector4D_OperatorMinus(void* self, void* v2);
void* QVector4D_OperatorMultiply(void* self, float factor);
void* QVector4D_OperatorMultiply2(void* self, void* v2);
void* QVector4D_OperatorMinus2(void* self);
void* QVector4D_OperatorDivide(void* self, float divisor);
void* QVector4D_OperatorDivide2(void* self, void* divisor);
void* QVector4D_OperatorMultiply3(void* self, void* matrix);
void QVector4D_delete(void* self);

#ifdef __cplusplus
}
#endif
