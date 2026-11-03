#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QQuaternion__Axis_fromVector3D(void* v);
void* QQuaternion__Axis_toVector3D(void* self);
float QQuaternion__Axis_x_get(void* self);
void QQuaternion__Axis_x_set(void* self, float x);
float QQuaternion__Axis_y_get(void* self);
void QQuaternion__Axis_y_set(void* self, float y);
float QQuaternion__Axis_z_get(void* self);
void QQuaternion__Axis_z_set(void* self, float z);
void QQuaternion__Axis_delete(void* self);

#ifdef __cplusplus
}
#endif
