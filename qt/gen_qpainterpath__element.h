#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPainterPath__Element_new();
void* QPainterPath__Element_new2(void* arg1);
int QPainterPath__Element_isMoveTo(void* self);
int QPainterPath__Element_isLineTo(void* self);
int QPainterPath__Element_isCurveTo(void* self);
void* QPainterPath__Element_ToQPointF(void* self);
int QPainterPath__Element_OperatorEqual(void* self, void* e);
int QPainterPath__Element_OperatorNotEqual(void* self, void* e);
double QPainterPath__Element_x_get(void* self);
void QPainterPath__Element_x_set(void* self, double x);
double QPainterPath__Element_y_get(void* self);
void QPainterPath__Element_y_set(void* self, double y);
int QPainterPath__Element_type_get(void* self);
void QPainterPath__Element_type_set(void* self, int type_);
void QPainterPath__Element_delete(void* self);

#ifdef __cplusplus
}
#endif
