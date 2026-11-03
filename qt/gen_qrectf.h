#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRectF_new();
void* QRectF_new2(void* topleft, void* size);
void* QRectF_new3(void* topleft, void* bottomRight);
void* QRectF_new4(double left, double top, double width, double height);
void* QRectF_new5(void* rect);
void* QRectF_new6(void* arg1);
int QRectF_isNull(void* self);
int QRectF_isEmpty(void* self);
int QRectF_isValid(void* self);
void* QRectF_normalized(void* self);
double QRectF_left(void* self);
double QRectF_top(void* self);
double QRectF_right(void* self);
double QRectF_bottom(void* self);
double QRectF_x(void* self);
double QRectF_y(void* self);
void QRectF_setLeft(void* self, double pos);
void QRectF_setTop(void* self, double pos);
void QRectF_setRight(void* self, double pos);
void QRectF_setBottom(void* self, double pos);
void QRectF_setX(void* self, double pos);
void QRectF_setY(void* self, double pos);
void* QRectF_topLeft(void* self);
void* QRectF_bottomRight(void* self);
void* QRectF_topRight(void* self);
void* QRectF_bottomLeft(void* self);
void* QRectF_center(void* self);
void QRectF_setTopLeft(void* self, void* p);
void QRectF_setBottomRight(void* self, void* p);
void QRectF_setTopRight(void* self, void* p);
void QRectF_setBottomLeft(void* self, void* p);
void QRectF_moveLeft(void* self, double pos);
void QRectF_moveTop(void* self, double pos);
void QRectF_moveRight(void* self, double pos);
void QRectF_moveBottom(void* self, double pos);
void QRectF_moveTopLeft(void* self, void* p);
void QRectF_moveBottomRight(void* self, void* p);
void QRectF_moveTopRight(void* self, void* p);
void QRectF_moveBottomLeft(void* self, void* p);
void QRectF_moveCenter(void* self, void* p);
void QRectF_translate(void* self, double dx, double dy);
void QRectF_translate2(void* self, void* p);
void* QRectF_translated(void* self, double dx, double dy);
void* QRectF_translated2(void* self, void* p);
void* QRectF_transposed(void* self);
void QRectF_moveTo(void* self, double x, double y);
void QRectF_moveTo2(void* self, void* p);
void QRectF_setRect(void* self, double x, double y, double w, double h);
void QRectF_getRect(void* self, double* x, double* y, double* w, double* h);
void QRectF_setCoords(void* self, double x1, double y1, double x2, double y2);
void QRectF_getCoords(void* self, double* x1, double* y1, double* x2, double* y2);
void QRectF_adjust(void* self, double x1, double y1, double x2, double y2);
void* QRectF_adjusted(void* self, double x1, double y1, double x2, double y2);
void* QRectF_size(void* self);
double QRectF_width(void* self);
double QRectF_height(void* self);
void QRectF_setWidth(void* self, double w);
void QRectF_setHeight(void* self, double h);
void QRectF_setSize(void* self, void* s);
void* QRectF_OperatorBitwiseOr(void* self, void* r);
void* QRectF_OperatorBitwiseAnd(void* self, void* r);
void* QRectF_OperatorBitwiseOrAssign(void* self, void* r);
void* QRectF_OperatorBitwiseAndAssign(void* self, void* r);
int QRectF_contains(void* self, void* r);
int QRectF_contains2(void* self, void* p);
int QRectF_contains3(void* self, double x, double y);
void* QRectF_united(void* self, void* other);
void* QRectF_intersected(void* self, void* other);
int QRectF_intersects(void* self, void* r);
void* QRectF_marginsAdded(void* self, void* margins);
void* QRectF_marginsRemoved(void* self, void* margins);
void* QRectF_OperatorPlusAssign(void* self, void* margins);
void* QRectF_OperatorMinusAssign(void* self, void* margins);
void* QRectF_toRect(void* self);
void* QRectF_toAlignedRect(void* self);
void* QRectF_OperatorAssign(void* self, void* arg1);
int QRectF_OperatorEqual(void* self, void* rhs);
int QRectF_OperatorNotEqual(void* self, void* rhs);
int QRectF_OperatorEqual2(void* self, void* rhs);
int QRectF_OperatorNotEqual2(void* self, void* rhs);
void* QRectF_OperatorPlus(void* self, void* rhs);
void* QRectF_OperatorMinus(void* self, void* rhs);
void QRectF_delete(void* self);

#ifdef __cplusplus
}
#endif
