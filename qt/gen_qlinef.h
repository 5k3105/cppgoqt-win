#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLineF_new();
void* QLineF_new2(void* pt1, void* pt2);
void* QLineF_new3(double x1, double y1, double x2, double y2);
void* QLineF_new4(void* line);
void* QLineF_new5(void* arg1);
void* QLineF_fromPolar(double length, double angle);
int QLineF_isNull(void* self);
void* QLineF_p1(void* self);
void* QLineF_p2(void* self);
double QLineF_x1(void* self);
double QLineF_y1(void* self);
double QLineF_x2(void* self);
double QLineF_y2(void* self);
double QLineF_dx(void* self);
double QLineF_dy(void* self);
double QLineF_length(void* self);
void QLineF_setLength(void* self, double len_);
double QLineF_angle(void* self);
void QLineF_setAngle(void* self, double angle);
double QLineF_angleTo(void* self, void* l);
void* QLineF_unitVector(void* self);
void* QLineF_normalVector(void* self);
int QLineF_intersects(void* self, void* l);
void* QLineF_pointAt(void* self, double t);
void QLineF_translate(void* self, void* p);
void QLineF_translate2(void* self, double dx, double dy);
void* QLineF_translated(void* self, void* p);
void* QLineF_translated2(void* self, double dx, double dy);
void* QLineF_center(void* self);
void QLineF_setP1(void* self, void* p1);
void QLineF_setP2(void* self, void* p2);
void QLineF_setPoints(void* self, void* p1, void* p2);
void QLineF_setLine(void* self, double x1, double y1, double x2, double y2);
void* QLineF_toLine(void* self);
int QLineF_intersects2(void* self, void* l, void* intersectionPoint);
int QLineF_OperatorEqual(void* self, void* rhs);
int QLineF_OperatorNotEqual(void* self, void* rhs);
int QLineF_OperatorEqual2(void* self, void* rhs);
int QLineF_OperatorNotEqual2(void* self, void* rhs);
void* QLineF_OperatorMultiply(void* self, void* m);
void QLineF_delete(void* self);

#ifdef __cplusplus
}
#endif
