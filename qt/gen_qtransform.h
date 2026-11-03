#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTransform_new(int arg1);
void* QTransform_new2();
void* QTransform_new3(double h11, double h12, double h13, double h21, double h22, double h23, double h31, double h32, double h33);
void* QTransform_new4(double h11, double h12, double h21, double h22, double dx, double dy);
void* QTransform_new5(void* other);
void* QTransform_OperatorAssign(void* self, void* arg1);
int QTransform_isAffine(void* self);
int QTransform_isIdentity(void* self);
int QTransform_isInvertible(void* self);
int QTransform_isScaling(void* self);
int QTransform_isRotating(void* self);
int QTransform_isTranslating(void* self);
int QTransform_type(void* self);
double QTransform_determinant(void* self);
double QTransform_m11(void* self);
double QTransform_m12(void* self);
double QTransform_m13(void* self);
double QTransform_m21(void* self);
double QTransform_m22(void* self);
double QTransform_m23(void* self);
double QTransform_m31(void* self);
double QTransform_m32(void* self);
double QTransform_m33(void* self);
double QTransform_dx(void* self);
double QTransform_dy(void* self);
void QTransform_setMatrix(void* self, double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33);
void* QTransform_inverted(void* self);
void* QTransform_adjoint(void* self);
void* QTransform_transposed(void* self);
void* QTransform_translate(void* self, double dx, double dy);
void* QTransform_scale(void* self, double sx, double sy);
void* QTransform_shear(void* self, double sh, double sv);
void* QTransform_rotate(void* self, double a, int axis, double distanceToPlane);
void* QTransform_rotate2(void* self, double a);
void* QTransform_rotateRadians(void* self, double a, int axis, double distanceToPlane);
void* QTransform_rotateRadians2(void* self, double a);
int QTransform_squareToQuad(void* square, void* result);
int QTransform_quadToSquare(void* quad, void* result);
int QTransform_quadToQuad(void* one, void* two, void* result);
int QTransform_OperatorEqual(void* self, void* arg1);
int QTransform_OperatorNotEqual(void* self, void* arg1);
void* QTransform_OperatorMultiplyAssign(void* self, void* arg1);
void* QTransform_OperatorMultiply(void* self, void* o);
void* QTransform_ToQVariant(void* self);
void QTransform_reset(void* self);
void* QTransform_map(void* self, void* p);
void* QTransform_map2(void* self, void* p);
void* QTransform_map3(void* self, void* l);
void* QTransform_map4(void* self, void* l);
void* QTransform_map5(void* self, void* a);
void* QTransform_map6(void* self, void* a);
void* QTransform_map7(void* self, void* r);
void* QTransform_map8(void* self, void* p);
void* QTransform_mapToPolygon(void* self, void* r);
void* QTransform_mapRect(void* self, void* arg1);
void* QTransform_mapRect2(void* self, void* arg1);
void QTransform_map9(void* self, int x, int y, int* tx, int* ty);
void QTransform_map10(void* self, double x, double y, double* tx, double* ty);
void* QTransform_OperatorMultiplyAssign2(void* self, double div);
void* QTransform_OperatorDivideAssign(void* self, double div);
void* QTransform_OperatorPlusAssign(void* self, double div);
void* QTransform_OperatorMinusAssign(void* self, double div);
void* QTransform_fromTranslate(double dx, double dy);
void* QTransform_fromScale(double dx, double dy);
void* QTransform_inverted2(void* self, int* invertible);
void* QTransform_rotate3(void* self, double a, int axis);
void* QTransform_rotateRadians3(void* self, double a, int axis);
void* QTransform_OperatorMultiply2(void* self, double n);
void* QTransform_OperatorDivide(void* self, double n);
void* QTransform_OperatorPlus(void* self, double n);
void* QTransform_OperatorMinus(void* self, double n);
void* QTransform_toQVariant(void* self);
void QTransform_delete(void* self);

#ifdef __cplusplus
}
#endif
