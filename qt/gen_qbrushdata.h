#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBrushData_ref_get(void* self);
void QBrushData_ref_set(void* self, void* ref);
int QBrushData_style_get(void* self);
void QBrushData_style_set(void* self, int style);
void* QBrushData_color_get(void* self);
void QBrushData_color_set(void* self, void* color);
void* QBrushData_transform_get(void* self);
void QBrushData_transform_set(void* self, void* transform);
void QBrushData_delete(void* self);

#ifdef __cplusplus
}
#endif
