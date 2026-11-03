#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QQuaternion__Axes_x_get(void* self);
void QQuaternion__Axes_x_set(void* self, void* x);
void* QQuaternion__Axes_y_get(void* self);
void QQuaternion__Axes_y_set(void* self, void* y);
void* QQuaternion__Axes_z_get(void* self);
void QQuaternion__Axes_z_set(void* self, void* z);
void QQuaternion__Axes_delete(void* self);

#ifdef __cplusplus
}
#endif
