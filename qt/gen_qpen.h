#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPen_new();
void* QPen_new2(int arg1);
void* QPen_new3(void* color);
void* QPen_new4(void* brush, double width);
void* QPen_new5(void* pen);
void* QPen_new6(void* brush, double width, int s);
void* QPen_new7(void* brush, double width, int s, int c);
void* QPen_new8(void* brush, double width, int s, int c, int j);
void* QPen_OperatorAssign(void* self, void* pen);
void QPen_swap(void* self, void* other);
void* QPen_OperatorAssign2(void* self, void* color);
void* QPen_OperatorAssign3(void* self, int style);
int QPen_style(void* self);
void QPen_setStyle(void* self, int arg1);
double* QPen_dashPattern(void* self, long long* __out_len);
void QPen_setDashPattern(void* self, double* pattern, long long pattern_len);
double QPen_dashOffset(void* self);
void QPen_setDashOffset(void* self, double doffset);
double QPen_miterLimit(void* self);
void QPen_setMiterLimit(void* self, double limit);
double QPen_widthF(void* self);
void QPen_setWidthF(void* self, double width);
int QPen_width(void* self);
void QPen_setWidth(void* self, int width);
void* QPen_color(void* self);
void QPen_setColor(void* self, void* color);
void* QPen_brush(void* self);
void QPen_setBrush(void* self, void* brush);
int QPen_isSolid(void* self);
int QPen_capStyle(void* self);
void QPen_setCapStyle(void* self, int pcs);
int QPen_joinStyle(void* self);
void QPen_setJoinStyle(void* self, int pcs);
int QPen_isCosmetic(void* self);
void QPen_setCosmetic(void* self, int cosmetic);
int QPen_OperatorEqual(void* self, void* p);
int QPen_OperatorNotEqual(void* self, void* p);
void* QPen_ToQVariant(void* self);
int QPen_isDetached(void* self);
int QPen_OperatorEqual2(void* self, void* rhs);
int QPen_OperatorNotEqual2(void* self, void* rhs);
void* QPen_toQVariant(void* self);
void QPen_delete(void* self);

#ifdef __cplusplus
}
#endif
