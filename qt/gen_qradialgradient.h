#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRadialGradient_base(void* self);
void* QRadialGradient_new();
void* QRadialGradient_new2(void* center, double radius, void* focalPoint);
void* QRadialGradient_new3(double cx, double cy, double radius, double fx, double fy);
void* QRadialGradient_new4(void* center, double radius);
void* QRadialGradient_new5(double cx, double cy, double radius);
void* QRadialGradient_new6(void* center, double centerRadius, void* focalPoint, double focalRadius);
void* QRadialGradient_new7(double cx, double cy, double centerRadius, double fx, double fy, double focalRadius);
void* QRadialGradient_new8(void* arg1);
void* QRadialGradient_center(void* self);
void QRadialGradient_setCenter(void* self, void* center);
void QRadialGradient_setCenter2(void* self, double x, double y);
void* QRadialGradient_focalPoint(void* self);
void QRadialGradient_setFocalPoint(void* self, void* focalPoint);
void QRadialGradient_setFocalPoint2(void* self, double x, double y);
double QRadialGradient_radius(void* self);
void QRadialGradient_setRadius(void* self, double radius);
double QRadialGradient_centerRadius(void* self);
void QRadialGradient_setCenterRadius(void* self, double radius);
double QRadialGradient_focalRadius(void* self);
void QRadialGradient_setFocalRadius(void* self, double radius);
void QRadialGradient_delete(void* self);

#ifdef __cplusplus
}
#endif
