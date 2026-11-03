#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QQuaternion_new();
void* QQuaternion_new2(int arg1);
void* QQuaternion_new3(float scalar, float xpos, float ypos, float zpos);
void* QQuaternion_new4(float scalar, void* vector);
void* QQuaternion_new5(void* vector);
void* QQuaternion_new6(void* arg1);
int QQuaternion_isNull(void* self);
int QQuaternion_isIdentity(void* self);
void* QQuaternion_vector(void* self);
void QQuaternion_setVector(void* self, void* vector);
void QQuaternion_setVector2(void* self, float x, float y, float z);
float QQuaternion_x(void* self);
float QQuaternion_y(void* self);
float QQuaternion_z(void* self);
float QQuaternion_scalar(void* self);
void QQuaternion_setX(void* self, float x);
void QQuaternion_setY(void* self, float y);
void QQuaternion_setZ(void* self, float z);
void QQuaternion_setScalar(void* self, float scalar);
float QQuaternion_dotProduct(void* q1, void* q2);
float QQuaternion_length(void* self);
float QQuaternion_lengthSquared(void* self);
void* QQuaternion_normalized(void* self);
void QQuaternion_normalize(void* self);
void* QQuaternion_inverted(void* self);
void* QQuaternion_conjugated(void* self);
void* QQuaternion_rotatedVector(void* self, void* vector);
void* QQuaternion_OperatorPlusAssign(void* self, void* quaternion);
void* QQuaternion_OperatorMinusAssign(void* self, void* quaternion);
void* QQuaternion_OperatorMultiplyAssign(void* self, float factor);
void* QQuaternion_OperatorMultiplyAssign2(void* self, void* quaternion);
void* QQuaternion_OperatorDivideAssign(void* self, float divisor);
void* QQuaternion_toVector4D(void* self);
void* QQuaternion_ToQVariant(void* self);
void QQuaternion_getAxisAndAngle(void* self, void* axis, float* angle);
void* QQuaternion_fromAxisAndAngle(void* axis, float angle);
void QQuaternion_getAxisAndAngle2(void* self, float* x, float* y, float* z, float* angle);
void* QQuaternion_fromAxisAndAngle2(float x, float y, float z, float angle);
void* QQuaternion_toEulerAngles(void* self);
void QQuaternion_getEulerAngles(void* self, float* pitch, float* yaw, float* roll);
void* QQuaternion_fromEulerAngles(float pitch, float yaw, float roll);
void* QQuaternion_toAxes(void* self);
void* QQuaternion_fromAxes(void* axes);
void QQuaternion_getAxes(void* self, void* xAxis, void* yAxis, void* zAxis);
void* QQuaternion_fromAxes2(void* xAxis, void* yAxis, void* zAxis);
void* QQuaternion_fromDirection(void* direction, void* up);
void* QQuaternion_rotationTo(void* from, void* to);
void* QQuaternion_slerp(void* q1, void* q2, float t);
void* QQuaternion_nlerp(void* q1, void* q2, float t);
void* QQuaternion_OperatorAssign(void* self, void* arg1);
int QQuaternion_OperatorEqual(void* self, void* q2);
int QQuaternion_OperatorNotEqual(void* self, void* q2);
void* QQuaternion_OperatorPlus(void* self, void* q2);
void* QQuaternion_OperatorMinus(void* self, void* q2);
void* QQuaternion_OperatorMultiply(void* self, float factor);
void* QQuaternion_OperatorMultiply2(void* self, void* q2);
void* QQuaternion_OperatorMinus2(void* self);
void* QQuaternion_OperatorDivide(void* self, float divisor);
void* QQuaternion_OperatorMultiply3(void* self, void* vec);
void QQuaternion_delete(void* self);

#ifdef __cplusplus
}
#endif
