#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMatrix4x4_new();
void* QMatrix4x4_new2(int arg1);
void* QMatrix4x4_new3(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44);
void* QMatrix4x4_new4(void* transform);
void* QMatrix4x4_new5(void* arg1);
void* QMatrix4x4_column(void* self, int index);
void QMatrix4x4_setColumn(void* self, int index, void* value);
void* QMatrix4x4_row(void* self, int index);
void QMatrix4x4_setRow(void* self, int index, void* value);
int QMatrix4x4_isAffine(void* self);
int QMatrix4x4_isIdentity(void* self);
void QMatrix4x4_setToIdentity(void* self);
void QMatrix4x4_fill(void* self, float value);
double QMatrix4x4_determinant(void* self);
void* QMatrix4x4_inverted(void* self);
void* QMatrix4x4_transposed(void* self);
void* QMatrix4x4_OperatorPlusAssign(void* self, void* other);
void* QMatrix4x4_OperatorMinusAssign(void* self, void* other);
void* QMatrix4x4_OperatorMultiplyAssign(void* self, void* other);
void* QMatrix4x4_OperatorMultiplyAssign2(void* self, float factor);
void* QMatrix4x4_OperatorDivideAssign(void* self, float divisor);
int QMatrix4x4_OperatorEqual(void* self, void* other);
int QMatrix4x4_OperatorNotEqual(void* self, void* other);
void QMatrix4x4_scale(void* self, void* vector);
void QMatrix4x4_translate(void* self, void* vector);
void QMatrix4x4_rotate(void* self, float angle, void* vector);
void QMatrix4x4_scale2(void* self, float x, float y);
void QMatrix4x4_scale3(void* self, float x, float y, float z);
void QMatrix4x4_scale4(void* self, float factor);
void QMatrix4x4_translate2(void* self, float x, float y);
void QMatrix4x4_translate3(void* self, float x, float y, float z);
void QMatrix4x4_rotate2(void* self, float angle, float x, float y);
void QMatrix4x4_rotate3(void* self, void* quaternion);
void QMatrix4x4_ortho(void* self, void* rect);
void QMatrix4x4_ortho2(void* self, void* rect);
void QMatrix4x4_ortho3(void* self, float left, float right, float bottom, float top, float nearPlane, float farPlane);
void QMatrix4x4_frustum(void* self, float left, float right, float bottom, float top, float nearPlane, float farPlane);
void QMatrix4x4_perspective(void* self, float verticalAngle, float aspectRatio, float nearPlane, float farPlane);
void QMatrix4x4_lookAt(void* self, void* eye, void* center, void* up);
void QMatrix4x4_viewport(void* self, void* rect);
void QMatrix4x4_viewport2(void* self, float left, float bottom, float width, float height);
void QMatrix4x4_flipCoordinates(void* self);
void QMatrix4x4_copyDataTo(void* self, float* values);
void* QMatrix4x4_toTransform(void* self);
void* QMatrix4x4_toTransform2(void* self, float distanceToPlane);
void* QMatrix4x4_map(void* self, void* point);
void* QMatrix4x4_map2(void* self, void* point);
void* QMatrix4x4_map3(void* self, void* point);
void* QMatrix4x4_mapVector(void* self, void* vector);
void* QMatrix4x4_map4(void* self, void* point);
void* QMatrix4x4_mapRect(void* self, void* rect);
void* QMatrix4x4_mapRect2(void* self, void* rect);
void QMatrix4x4_optimize(void* self);
void* QMatrix4x4_ToQVariant(void* self);
void QMatrix4x4_projectedRotate(void* self, float angle, float x, float y, float z, float distanceToPlane);
void QMatrix4x4_projectedRotate2(void* self, float angle, float x, float y, float z);
int QMatrix4x4_flags(void* self);
void* QMatrix4x4_inverted2(void* self, int* invertible);
void QMatrix4x4_rotate4(void* self, float angle, float x, float y, float z);
void QMatrix4x4_viewport3(void* self, float left, float bottom, float width, float height, float nearPlane);
void QMatrix4x4_viewport4(void* self, float left, float bottom, float width, float height, float nearPlane, float farPlane);
void* QMatrix4x4_OperatorMultiply(void* self, void* vector);
void* QMatrix4x4_OperatorMultiply2(void* self, void* vector);
void* QMatrix4x4_OperatorPlus(void* self, void* m2);
void* QMatrix4x4_OperatorMinus(void* self, void* m2);
void* QMatrix4x4_OperatorMultiply3(void* self, void* m2);
void* QMatrix4x4_OperatorMinus2(void* self);
void* QMatrix4x4_OperatorMultiply4(void* self, void* point);
void* QMatrix4x4_OperatorMultiply5(void* self, void* point);
void* QMatrix4x4_OperatorMultiply6(void* self, float factor);
void* QMatrix4x4_OperatorDivide(void* self, float divisor);
void QMatrix4x4_delete(void* self);

#ifdef __cplusplus
}
#endif
