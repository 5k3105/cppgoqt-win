#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBrush_new();
void* QBrush_new2(int bs);
void* QBrush_new3(void* color);
void* QBrush_new4(int color);
void* QBrush_new5(void* color, void* pixmap);
void* QBrush_new6(int color, void* pixmap);
void* QBrush_new7(void* pixmap);
void* QBrush_new8(void* image);
void* QBrush_new9(void* brush);
void* QBrush_new10(void* gradient);
void* QBrush_new11(void* color, int bs);
void* QBrush_new12(int color, int bs);
void* QBrush_OperatorAssign(void* self, void* brush);
void QBrush_swap(void* self, void* other);
void* QBrush_OperatorAssign2(void* self, int style);
void* QBrush_OperatorAssign3(void* self, void* color);
void* QBrush_OperatorAssign4(void* self, int color);
void* QBrush_ToQVariant(void* self);
int QBrush_style(void* self);
void QBrush_setStyle(void* self, int arg1);
void* QBrush_transform(void* self);
void QBrush_setTransform(void* self, void* arg1);
void* QBrush_texture(void* self);
void QBrush_setTexture(void* self, void* pixmap);
void* QBrush_textureImage(void* self);
void QBrush_setTextureImage(void* self, void* image);
void* QBrush_color(void* self);
void QBrush_setColor(void* self, void* color);
void QBrush_setColor2(void* self, int color);
void* QBrush_gradient(void* self);
int QBrush_isOpaque(void* self);
int QBrush_OperatorEqual(void* self, void* b);
int QBrush_OperatorNotEqual(void* self, void* b);
int QBrush_isDetached(void* self);
int QBrush_OperatorEqual2(void* self, void* rhs);
int QBrush_OperatorNotEqual2(void* self, void* rhs);
void* QBrush_toQVariant(void* self);
void QBrush_delete(void* self);

#ifdef __cplusplus
}
#endif
