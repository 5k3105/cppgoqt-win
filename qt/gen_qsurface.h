#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QSurface_surfaceClass(void* self);
void* QSurface_format(void* self);
int QSurface_surfaceType(void* self);
int QSurface_supportsOpenGL(void* self);
void* QSurface_size(void* self);
void* QSurface_OperatorAssign(void* self, void* arg1);
void QSurface_delete(void* self);

#ifdef __cplusplus
}
#endif
