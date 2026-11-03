#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QEasingCurve_new();
void* QEasingCurve_new2(void* other);
void* QEasingCurve_new3(int type_);
void* QEasingCurve_OperatorAssign(void* self, void* other);
void QEasingCurve_swap(void* self, void* other);
double QEasingCurve_amplitude(void* self);
void QEasingCurve_setAmplitude(void* self, double amplitude);
double QEasingCurve_period(void* self);
void QEasingCurve_setPeriod(void* self, double period);
double QEasingCurve_overshoot(void* self);
void QEasingCurve_setOvershoot(void* self, double overshoot);
void QEasingCurve_addCubicBezierSegment(void* self, void* c1, void* c2, void* endPoint);
void QEasingCurve_addTCBSegment(void* self, void* nextPoint, double t, double c, double b);
void** QEasingCurve_toCubicSpline(void* self, long long* __out_len);
int QEasingCurve_type(void* self);
void QEasingCurve_setType(void* self, int type_);
double QEasingCurve_valueForProgress(void* self, double progress);
int QEasingCurve_OperatorEqual(void* self, void* rhs);
int QEasingCurve_OperatorNotEqual(void* self, void* rhs);
void QEasingCurve_delete(void* self);

#ifdef __cplusplus
}
#endif
