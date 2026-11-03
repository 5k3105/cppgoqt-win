#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QConicalGradient_base(void* self);
void* QConicalGradient_new();
void* QConicalGradient_new2(void* center, double startAngle);
void* QConicalGradient_new3(double cx, double cy, double startAngle);
void* QConicalGradient_new4(void* arg1);
void* QConicalGradient_center(void* self);
void QConicalGradient_setCenter(void* self, void* center);
void QConicalGradient_setCenter2(void* self, double x, double y);
double QConicalGradient_angle(void* self);
void QConicalGradient_setAngle(void* self, double angle);
void QConicalGradient_delete(void* self);

#ifdef __cplusplus
}
#endif
