#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRect_new();
void* QRect_new2(void* topleft, void* bottomright);
void* QRect_new3(void* topleft, void* size);
void* QRect_new4(int left, int top, int width, int height);
void* QRect_new5(void* arg1);
int QRect_isNull(void* self);
int QRect_isEmpty(void* self);
int QRect_isValid(void* self);
int QRect_left(void* self);
int QRect_top(void* self);
int QRect_right(void* self);
int QRect_bottom(void* self);
void* QRect_normalized(void* self);
int QRect_x(void* self);
int QRect_y(void* self);
void QRect_setLeft(void* self, int pos);
void QRect_setTop(void* self, int pos);
void QRect_setRight(void* self, int pos);
void QRect_setBottom(void* self, int pos);
void QRect_setX(void* self, int x);
void QRect_setY(void* self, int y);
void QRect_setTopLeft(void* self, void* p);
void QRect_setBottomRight(void* self, void* p);
void QRect_setTopRight(void* self, void* p);
void QRect_setBottomLeft(void* self, void* p);
void* QRect_topLeft(void* self);
void* QRect_bottomRight(void* self);
void* QRect_topRight(void* self);
void* QRect_bottomLeft(void* self);
void* QRect_center(void* self);
void QRect_moveLeft(void* self, int pos);
void QRect_moveTop(void* self, int pos);
void QRect_moveRight(void* self, int pos);
void QRect_moveBottom(void* self, int pos);
void QRect_moveTopLeft(void* self, void* p);
void QRect_moveBottomRight(void* self, void* p);
void QRect_moveTopRight(void* self, void* p);
void QRect_moveBottomLeft(void* self, void* p);
void QRect_moveCenter(void* self, void* p);
void QRect_translate(void* self, int dx, int dy);
void QRect_translate2(void* self, void* p);
void* QRect_translated(void* self, int dx, int dy);
void* QRect_translated2(void* self, void* p);
void* QRect_transposed(void* self);
void QRect_moveTo(void* self, int x, int t);
void QRect_moveTo2(void* self, void* p);
void QRect_setRect(void* self, int x, int y, int w, int h);
void QRect_getRect(void* self, int* x, int* y, int* w, int* h);
void QRect_setCoords(void* self, int x1, int y1, int x2, int y2);
void QRect_getCoords(void* self, int* x1, int* y1, int* x2, int* y2);
void QRect_adjust(void* self, int x1, int y1, int x2, int y2);
void* QRect_adjusted(void* self, int x1, int y1, int x2, int y2);
void* QRect_size(void* self);
int QRect_width(void* self);
int QRect_height(void* self);
void QRect_setWidth(void* self, int w);
void QRect_setHeight(void* self, int h);
void QRect_setSize(void* self, void* s);
void* QRect_OperatorBitwiseOr(void* self, void* r);
void* QRect_OperatorBitwiseAnd(void* self, void* r);
void* QRect_OperatorBitwiseOrAssign(void* self, void* r);
void* QRect_OperatorBitwiseAndAssign(void* self, void* r);
int QRect_contains(void* self, void* r);
int QRect_contains2(void* self, void* p);
int QRect_contains3(void* self, int x, int y);
int QRect_contains4(void* self, int x, int y, int proper);
void* QRect_united(void* self, void* other);
void* QRect_intersected(void* self, void* other);
int QRect_intersects(void* self, void* r);
void* QRect_marginsAdded(void* self, void* margins);
void* QRect_marginsRemoved(void* self, void* margins);
void* QRect_OperatorPlusAssign(void* self, void* margins);
void* QRect_OperatorMinusAssign(void* self, void* margins);
void* QRect_span(void* p1, void* p2);
void* QRect_toRectF(void* self);
void* QRect_OperatorAssign(void* self, void* arg1);
int QRect_contains5(void* self, void* r, int proper);
int QRect_contains6(void* self, void* p, int proper);
int QRect_OperatorEqual(void* self, void* rhs);
int QRect_OperatorNotEqual(void* self, void* rhs);
void* QRect_OperatorPlus(void* self, void* margins);
void* QRect_OperatorMinus(void* self, void* rhs);
int QRect_OperatorEqual2(void* self, void* rhs);
int QRect_OperatorNotEqual2(void* self, void* rhs);
void QRect_delete(void* self);

#ifdef __cplusplus
}
#endif
