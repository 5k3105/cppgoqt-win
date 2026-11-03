#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QColorSpace__PrimaryPoints_fromPrimaries(int primaries);
int QColorSpace__PrimaryPoints_isValid(void* self);
void* QColorSpace__PrimaryPoints_whitePoint_get(void* self);
void QColorSpace__PrimaryPoints_whitePoint_set(void* self, void* whitePoint);
void* QColorSpace__PrimaryPoints_redPoint_get(void* self);
void QColorSpace__PrimaryPoints_redPoint_set(void* self, void* redPoint);
void* QColorSpace__PrimaryPoints_greenPoint_get(void* self);
void QColorSpace__PrimaryPoints_greenPoint_set(void* self, void* greenPoint);
void* QColorSpace__PrimaryPoints_bluePoint_get(void* self);
void QColorSpace__PrimaryPoints_bluePoint_set(void* self, void* bluePoint);
void QColorSpace__PrimaryPoints_delete(void* self);

#ifdef __cplusplus
}
#endif
