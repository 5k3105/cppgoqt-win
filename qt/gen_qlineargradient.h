#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLinearGradient_base(void* self);
void* QLinearGradient_new();
void* QLinearGradient_new2(void* start, void* finalStop);
void* QLinearGradient_new3(double xStart, double yStart, double xFinalStop, double yFinalStop);
void* QLinearGradient_new4(void* arg1);
void* QLinearGradient_start(void* self);
void QLinearGradient_setStart(void* self, void* start);
void QLinearGradient_setStart2(void* self, double x, double y);
void* QLinearGradient_finalStop(void* self);
void QLinearGradient_setFinalStop(void* self, void* stop);
void QLinearGradient_setFinalStop2(void* self, double x, double y);
void QLinearGradient_delete(void* self);

#ifdef __cplusplus
}
#endif
