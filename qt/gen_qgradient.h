#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGradient_new();
void* QGradient_new2(int arg1);
void* QGradient_new3(void* arg1);
int QGradient_type(void* self);
void QGradient_setSpread(void* self, int spread);
int QGradient_spread(void* self);
void QGradient_setColorAt(void* self, double pos, void* color);
int QGradient_coordinateMode(void* self);
void QGradient_setCoordinateMode(void* self, int mode);
int QGradient_interpolationMode(void* self);
void QGradient_setInterpolationMode(void* self, int mode);
int QGradient_OperatorEqual(void* self, void* gradient);
int QGradient_OperatorNotEqual(void* self, void* other);
void QGradient_delete(void* self);

#ifdef __cplusplus
}
#endif
