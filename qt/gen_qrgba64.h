#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRgba64_new();
void* QRgba64_new2(void* arg1);
void* QRgba64_fromRgba64(unsigned long long c);
void* QRgba64_fromRgba642(unsigned short red, unsigned short green, unsigned short blue, unsigned short alpha);
void* QRgba64_fromRgba(unsigned char red, unsigned char green, unsigned char blue, unsigned char alpha);
void* QRgba64_fromArgb32(unsigned int rgb);
int QRgba64_isOpaque(void* self);
int QRgba64_isTransparent(void* self);
unsigned short QRgba64_red(void* self);
unsigned short QRgba64_green(void* self);
unsigned short QRgba64_blue(void* self);
unsigned short QRgba64_alpha(void* self);
void QRgba64_setRed(void* self, unsigned short _red);
void QRgba64_setGreen(void* self, unsigned short _green);
void QRgba64_setBlue(void* self, unsigned short _blue);
void QRgba64_setAlpha(void* self, unsigned short _alpha);
unsigned char QRgba64_red8(void* self);
unsigned char QRgba64_green8(void* self);
unsigned char QRgba64_blue8(void* self);
unsigned char QRgba64_alpha8(void* self);
unsigned int QRgba64_toArgb32(void* self);
unsigned short QRgba64_toRgb16(void* self);
void* QRgba64_premultiplied(void* self);
void* QRgba64_unpremultiplied(void* self);
unsigned long long QRgba64_ToQuint64(void* self);
void* QRgba64_OperatorAssign(void* self, unsigned long long _rgba);
void QRgba64_delete(void* self);

#ifdef __cplusplus
}
#endif
