#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVector3D_new();
void* QVector3D_new2(int arg1);
void* QVector3D_new3(float xpos, float ypos, float zpos);
void* QVector3D_new4(void* point);
void* QVector3D_new5(void* point);
void* QVector3D_new6(void* vector);
void* QVector3D_new7(void* vector, float zpos);
void* QVector3D_new8(void* vector);
void* QVector3D_new9(void* arg1);
int QVector3D_isNull(void* self);
float QVector3D_x(void* self);
float QVector3D_y(void* self);
float QVector3D_z(void* self);
void QVector3D_setX(void* self, float x);
void QVector3D_setY(void* self, float y);
void QVector3D_setZ(void* self, float z);
float QVector3D_OperatorSubscript(void* self, int i);
float QVector3D_length(void* self);
float QVector3D_lengthSquared(void* self);
void* QVector3D_normalized(void* self);
void QVector3D_normalize(void* self);
void* QVector3D_OperatorPlusAssign(void* self, void* vector);
void* QVector3D_OperatorMinusAssign(void* self, void* vector);
void* QVector3D_OperatorMultiplyAssign(void* self, float factor);
void* QVector3D_OperatorMultiplyAssign2(void* self, void* vector);
void* QVector3D_OperatorDivideAssign(void* self, float divisor);
void* QVector3D_OperatorDivideAssign2(void* self, void* vector);
float QVector3D_dotProduct(void* v1, void* v2);
void* QVector3D_crossProduct(void* v1, void* v2);
void* QVector3D_normal(void* v1, void* v2);
void* QVector3D_normal2(void* v1, void* v2, void* v3);
void* QVector3D_project(void* self, void* modelView, void* projection, void* viewport);
void* QVector3D_unproject(void* self, void* modelView, void* projection, void* viewport);
float QVector3D_distanceToPoint(void* self, void* point);
float QVector3D_distanceToPlane(void* self, void* plane, void* normal);
float QVector3D_distanceToPlane2(void* self, void* plane1, void* plane2, void* plane3);
float QVector3D_distanceToLine(void* self, void* point, void* direction);
void* QVector3D_toVector2D(void* self);
void* QVector3D_toVector4D(void* self);
void* QVector3D_toPoint(void* self);
void* QVector3D_toPointF(void* self);
void* QVector3D_ToQVariant(void* self);
int QVector3D_OperatorEqual(void* self, void* v2);
int QVector3D_OperatorNotEqual(void* self, void* v2);
void* QVector3D_OperatorPlus(void* self, void* v2);
void* QVector3D_OperatorMinus(void* self, void* v2);
void* QVector3D_OperatorMultiply(void* self, float factor);
void* QVector3D_OperatorMultiply2(void* self, void* v2);
void* QVector3D_OperatorMinus2(void* self);
void* QVector3D_OperatorDivide(void* self, float divisor);
void* QVector3D_OperatorDivide2(void* self, void* divisor);
void* QVector3D_OperatorMultiply3(void* self, void* matrix);
void QVector3D_delete(void* self);

#ifdef __cplusplus
}
#endif
