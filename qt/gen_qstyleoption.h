#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOption_new();
void* QStyleOption_new2(void* other);
void* QStyleOption_new3(int version);
void* QStyleOption_new4(int version, int type_);
void QStyleOption_initFrom(void* self, void* w);
void* QStyleOption_OperatorAssign(void* self, void* other);
int QStyleOption_version_get(void* self);
void QStyleOption_version_set(void* self, int version);
int QStyleOption_type_get(void* self);
void QStyleOption_type_set(void* self, int type_);
int QStyleOption_state_get(void* self);
void QStyleOption_state_set(void* self, int state);
int QStyleOption_direction_get(void* self);
void QStyleOption_direction_set(void* self, int direction);
void* QStyleOption_rect_get(void* self);
void QStyleOption_rect_set(void* self, void* rect);
void* QStyleOption_fontMetrics_get(void* self);
void QStyleOption_fontMetrics_set(void* self, void* fontMetrics);
void* QStyleOption_palette_get(void* self);
void QStyleOption_palette_set(void* self, void* palette);
void* QStyleOption_styleObject_get(void* self);
void QStyleOption_styleObject_set(void* self, void* styleObject);
void QStyleOption_delete(void* self);

#ifdef __cplusplus
}
#endif
