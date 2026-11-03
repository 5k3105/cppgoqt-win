#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIconEngine__ScaledPixmapArgument_new(void* arg1);
void* QIconEngine__ScaledPixmapArgument_OperatorAssign(void* self, void* arg1);
void* QIconEngine__ScaledPixmapArgument_size_get(void* self);
void QIconEngine__ScaledPixmapArgument_size_set(void* self, void* size);
int QIconEngine__ScaledPixmapArgument_mode_get(void* self);
void QIconEngine__ScaledPixmapArgument_mode_set(void* self, int mode);
int QIconEngine__ScaledPixmapArgument_state_get(void* self);
void QIconEngine__ScaledPixmapArgument_state_set(void* self, int state);
double QIconEngine__ScaledPixmapArgument_scale_get(void* self);
void QIconEngine__ScaledPixmapArgument_scale_set(void* self, double scale);
void* QIconEngine__ScaledPixmapArgument_pixmap_get(void* self);
void QIconEngine__ScaledPixmapArgument_pixmap_set(void* self, void* pixmap);
void QIconEngine__ScaledPixmapArgument_delete(void* self);

#ifdef __cplusplus
}
#endif
