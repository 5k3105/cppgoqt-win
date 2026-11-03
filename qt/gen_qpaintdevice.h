#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QPaintDevice_devType(void* self);
int QPaintDevice_paintingActive(void* self);
void* QPaintDevice_paintEngine(void* self);
int QPaintDevice_width(void* self);
int QPaintDevice_height(void* self);
int QPaintDevice_widthMM(void* self);
int QPaintDevice_heightMM(void* self);
int QPaintDevice_logicalDpiX(void* self);
int QPaintDevice_logicalDpiY(void* self);
int QPaintDevice_physicalDpiX(void* self);
int QPaintDevice_physicalDpiY(void* self);
double QPaintDevice_devicePixelRatio(void* self);
double QPaintDevice_devicePixelRatioF(void* self);
int QPaintDevice_colorCount(void* self);
int QPaintDevice_depth(void* self);
double QPaintDevice_devicePixelRatioFScale();
int QPaintDevice_encodeMetricF(int metric, double value);
void QPaintDevice_delete(void* self);

#ifdef __cplusplus
}
#endif
