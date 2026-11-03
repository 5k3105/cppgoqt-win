#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QColormap_new(void* colormap);
void QColormap_initialize();
void QColormap_cleanup();
void* QColormap_instance();
void* QColormap_OperatorAssign(void* self, void* colormap);
int QColormap_mode(void* self);
int QColormap_depth(void* self);
int QColormap_size(void* self);
unsigned int QColormap_pixel(void* self, void* color);
void* QColormap_colorAt(void* self, unsigned int pixel);
void** QColormap_colormap(void* self, long long* __out_len);
void* QColormap_instance2(int screen);
void QColormap_delete(void* self);

#ifdef __cplusplus
}
#endif
