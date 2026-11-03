#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLine_new();
void* QLine_new2(void* pt1, void* pt2);
void* QLine_new3(int x1, int y1, int x2, int y2);
void* QLine_new4(void* arg1);
int QLine_isNull(void* self);
void* QLine_p1(void* self);
void* QLine_p2(void* self);
int QLine_x1(void* self);
int QLine_y1(void* self);
int QLine_x2(void* self);
int QLine_y2(void* self);
int QLine_dx(void* self);
int QLine_dy(void* self);
void QLine_translate(void* self, void* p);
void QLine_translate2(void* self, int dx, int dy);
void* QLine_translated(void* self, void* p);
void* QLine_translated2(void* self, int dx, int dy);
void* QLine_center(void* self);
void QLine_setP1(void* self, void* p1);
void QLine_setP2(void* self, void* p2);
void QLine_setPoints(void* self, void* p1, void* p2);
void QLine_setLine(void* self, int x1, int y1, int x2, int y2);
void* QLine_toLineF(void* self);
int QLine_OperatorEqual(void* self, void* rhs);
int QLine_OperatorNotEqual(void* self, void* rhs);
int QLine_OperatorEqual2(void* self, void* rhs);
int QLine_OperatorNotEqual2(void* self, void* rhs);
void* QLine_OperatorMultiply(void* self, void* m);
void QLine_delete(void* self);

#ifdef __cplusplus
}
#endif
