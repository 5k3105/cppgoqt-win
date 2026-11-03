#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPainterPathStroker_new();
void* QPainterPathStroker_new2(void* pen);
void QPainterPathStroker_setWidth(void* self, double width);
double QPainterPathStroker_width(void* self);
void QPainterPathStroker_setCapStyle(void* self, int style);
int QPainterPathStroker_capStyle(void* self);
void QPainterPathStroker_setJoinStyle(void* self, int style);
int QPainterPathStroker_joinStyle(void* self);
void QPainterPathStroker_setMiterLimit(void* self, double length);
double QPainterPathStroker_miterLimit(void* self);
void QPainterPathStroker_setCurveThreshold(void* self, double threshold);
double QPainterPathStroker_curveThreshold(void* self);
void QPainterPathStroker_setDashPattern(void* self, int arg1);
void QPainterPathStroker_setDashPattern2(void* self, double* dashPattern, long long dashPattern_len);
double* QPainterPathStroker_dashPattern(void* self, long long* __out_len);
void QPainterPathStroker_setDashOffset(void* self, double offset);
double QPainterPathStroker_dashOffset(void* self);
void* QPainterPathStroker_createStroke(void* self, void* path);
void QPainterPathStroker_delete(void* self);

#ifdef __cplusplus
}
#endif
