#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPainterPath_new();
void* QPainterPath_new2(void* startPoint);
void* QPainterPath_new3(void* other);
void* QPainterPath_OperatorAssign(void* self, void* other);
void QPainterPath_swap(void* self, void* other);
void QPainterPath_clear(void* self);
void QPainterPath_reserve(void* self, int size);
int QPainterPath_capacity(void* self);
void QPainterPath_closeSubpath(void* self);
void QPainterPath_moveTo(void* self, void* p);
void QPainterPath_moveTo2(void* self, double x, double y);
void QPainterPath_lineTo(void* self, void* p);
void QPainterPath_lineTo2(void* self, double x, double y);
void QPainterPath_arcMoveTo(void* self, void* rect, double angle);
void QPainterPath_arcMoveTo2(void* self, double x, double y, double w, double h, double angle);
void QPainterPath_arcTo(void* self, void* rect, double startAngle, double arcLength);
void QPainterPath_arcTo2(void* self, double x, double y, double w, double h, double startAngle, double arcLength);
void QPainterPath_cubicTo(void* self, void* ctrlPt1, void* ctrlPt2, void* endPt);
void QPainterPath_cubicTo2(void* self, double ctrlPt1x, double ctrlPt1y, double ctrlPt2x, double ctrlPt2y, double endPtx, double endPty);
void QPainterPath_quadTo(void* self, void* ctrlPt, void* endPt);
void QPainterPath_quadTo2(void* self, double ctrlPtx, double ctrlPty, double endPtx, double endPty);
void* QPainterPath_currentPosition(void* self);
void QPainterPath_addRect(void* self, void* rect);
void QPainterPath_addRect2(void* self, double x, double y, double w, double h);
void QPainterPath_addEllipse(void* self, void* rect);
void QPainterPath_addEllipse2(void* self, double x, double y, double w, double h);
void QPainterPath_addEllipse3(void* self, void* center, double rx, double ry);
void QPainterPath_addPolygon(void* self, void* polygon);
void QPainterPath_addText(void* self, void* point, void* f, const char* text, long long text_len);
void QPainterPath_addText2(void* self, double x, double y, void* f, const char* text, long long text_len);
void QPainterPath_addPath(void* self, void* path);
void QPainterPath_addRegion(void* self, void* region);
void QPainterPath_addRoundedRect(void* self, void* rect, double xRadius, double yRadius);
void QPainterPath_addRoundedRect2(void* self, double x, double y, double w, double h, double xRadius, double yRadius);
void QPainterPath_connectPath(void* self, void* path);
int QPainterPath_contains(void* self, void* pt);
int QPainterPath_contains2(void* self, void* rect);
int QPainterPath_intersects(void* self, void* rect);
void QPainterPath_translate(void* self, double dx, double dy);
void QPainterPath_translate2(void* self, void* offset);
void* QPainterPath_translated(void* self, double dx, double dy);
void* QPainterPath_translated2(void* self, void* offset);
void* QPainterPath_boundingRect(void* self);
void* QPainterPath_controlPointRect(void* self);
int QPainterPath_fillRule(void* self);
void QPainterPath_setFillRule(void* self, int fillRule);
int QPainterPath_isEmpty(void* self);
void* QPainterPath_toReversed(void* self);
void** QPainterPath_toSubpathPolygons(void* self, long long* __out_len);
void** QPainterPath_toFillPolygons(void* self, long long* __out_len);
void* QPainterPath_toFillPolygon(void* self);
int QPainterPath_elementCount(void* self);
void* QPainterPath_elementAt(void* self, int i);
void QPainterPath_setElementPositionAt(void* self, int i, double x, double y);
int QPainterPath_isCachingEnabled(void* self);
void QPainterPath_setCachingEnabled(void* self, int enabled);
double QPainterPath_length(void* self);
double QPainterPath_percentAtLength(void* self, double len_);
void* QPainterPath_pointAtPercent(void* self, double t);
double QPainterPath_angleAtPercent(void* self, double t);
double QPainterPath_slopeAtPercent(void* self, double t);
void* QPainterPath_trimmed(void* self, double fromFraction, double toFraction);
int QPainterPath_intersects2(void* self, void* p);
int QPainterPath_contains3(void* self, void* p);
void* QPainterPath_united(void* self, void* r);
void* QPainterPath_intersected(void* self, void* r);
void* QPainterPath_subtracted(void* self, void* r);
void* QPainterPath_simplified(void* self);
int QPainterPath_OperatorEqual(void* self, void* other);
int QPainterPath_OperatorNotEqual(void* self, void* other);
void* QPainterPath_OperatorBitwiseAnd(void* self, void* other);
void* QPainterPath_OperatorBitwiseOr(void* self, void* other);
void* QPainterPath_OperatorPlus(void* self, void* other);
void* QPainterPath_OperatorMinus(void* self, void* other);
void* QPainterPath_OperatorBitwiseAndAssign(void* self, void* other);
void* QPainterPath_OperatorBitwiseOrAssign(void* self, void* other);
void* QPainterPath_OperatorPlusAssign(void* self, void* other);
void* QPainterPath_OperatorMinusAssign(void* self, void* other);
void QPainterPath_addRoundedRect3(void* self, void* rect, double xRadius, double yRadius, int mode);
void QPainterPath_addRoundedRect4(void* self, double x, double y, double w, double h, double xRadius, double yRadius, int mode);
void** QPainterPath_toSubpathPolygons2(void* self, void* matrix, long long* __out_len);
void** QPainterPath_toFillPolygons2(void* self, void* matrix, long long* __out_len);
void* QPainterPath_toFillPolygon2(void* self, void* matrix);
void* QPainterPath_trimmed2(void* self, double fromFraction, double toFraction, double offset);
void* QPainterPath_OperatorMultiply(void* self, void* m);
void QPainterPath_delete(void* self);

#ifdef __cplusplus
}
#endif
