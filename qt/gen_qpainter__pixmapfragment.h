#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPainter__PixmapFragment_new();
void* QPainter__PixmapFragment_new2(void* arg1);
void* QPainter__PixmapFragment_create(void* pos, void* sourceRect);
void* QPainter__PixmapFragment_create2(void* pos, void* sourceRect, double scaleX);
void* QPainter__PixmapFragment_create3(void* pos, void* sourceRect, double scaleX, double scaleY);
void* QPainter__PixmapFragment_create4(void* pos, void* sourceRect, double scaleX, double scaleY, double rotation);
void* QPainter__PixmapFragment_create5(void* pos, void* sourceRect, double scaleX, double scaleY, double rotation, double opacity);
double QPainter__PixmapFragment_x_get(void* self);
void QPainter__PixmapFragment_x_set(void* self, double x);
double QPainter__PixmapFragment_y_get(void* self);
void QPainter__PixmapFragment_y_set(void* self, double y);
double QPainter__PixmapFragment_sourceLeft_get(void* self);
void QPainter__PixmapFragment_sourceLeft_set(void* self, double sourceLeft);
double QPainter__PixmapFragment_sourceTop_get(void* self);
void QPainter__PixmapFragment_sourceTop_set(void* self, double sourceTop);
double QPainter__PixmapFragment_width_get(void* self);
void QPainter__PixmapFragment_width_set(void* self, double width);
double QPainter__PixmapFragment_height_get(void* self);
void QPainter__PixmapFragment_height_set(void* self, double height);
double QPainter__PixmapFragment_scaleX_get(void* self);
void QPainter__PixmapFragment_scaleX_set(void* self, double scaleX);
double QPainter__PixmapFragment_scaleY_get(void* self);
void QPainter__PixmapFragment_scaleY_set(void* self, double scaleY);
double QPainter__PixmapFragment_rotation_get(void* self);
void QPainter__PixmapFragment_rotation_set(void* self, double rotation);
double QPainter__PixmapFragment_opacity_get(void* self);
void QPainter__PixmapFragment_opacity_set(void* self, double opacity);
void QPainter__PixmapFragment_delete(void* self);

#ifdef __cplusplus
}
#endif
