#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTileRules_new(int horizontalRule, int verticalRule);
void* QTileRules_new2();
void* QTileRules_new3(void* arg1);
void* QTileRules_new4(int rule);
int QTileRules_horizontal_get(void* self);
void QTileRules_horizontal_set(void* self, int horizontal);
int QTileRules_vertical_get(void* self);
void QTileRules_vertical_set(void* self, int vertical);
void QTileRules_delete(void* self);

#ifdef __cplusplus
}
#endif
