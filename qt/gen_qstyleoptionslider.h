#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionSlider_base(void* self);
void QStyleOptionSlider_chain(void* self, void** out);
void* QStyleOptionSlider_new();
void* QStyleOptionSlider_new2(void* other);
void* QStyleOptionSlider_OperatorAssign(void* self, void* arg1);
int QStyleOptionSlider_orientation_get(void* self);
void QStyleOptionSlider_orientation_set(void* self, int orientation);
int QStyleOptionSlider_minimum_get(void* self);
void QStyleOptionSlider_minimum_set(void* self, int minimum);
int QStyleOptionSlider_maximum_get(void* self);
void QStyleOptionSlider_maximum_set(void* self, int maximum);
int QStyleOptionSlider_tickPosition_get(void* self);
void QStyleOptionSlider_tickPosition_set(void* self, int tickPosition);
int QStyleOptionSlider_tickInterval_get(void* self);
void QStyleOptionSlider_tickInterval_set(void* self, int tickInterval);
int QStyleOptionSlider_upsideDown_get(void* self);
void QStyleOptionSlider_upsideDown_set(void* self, int upsideDown);
int QStyleOptionSlider_sliderPosition_get(void* self);
void QStyleOptionSlider_sliderPosition_set(void* self, int sliderPosition);
int QStyleOptionSlider_sliderValue_get(void* self);
void QStyleOptionSlider_sliderValue_set(void* self, int sliderValue);
int QStyleOptionSlider_singleStep_get(void* self);
void QStyleOptionSlider_singleStep_set(void* self, int singleStep);
int QStyleOptionSlider_pageStep_get(void* self);
void QStyleOptionSlider_pageStep_set(void* self, int pageStep);
double QStyleOptionSlider_notchTarget_get(void* self);
void QStyleOptionSlider_notchTarget_set(void* self, double notchTarget);
int QStyleOptionSlider_dialWrapping_get(void* self);
void QStyleOptionSlider_dialWrapping_set(void* self, int dialWrapping);
int QStyleOptionSlider_keyboardModifiers_get(void* self);
void QStyleOptionSlider_keyboardModifiers_set(void* self, int keyboardModifiers);
void QStyleOptionSlider_delete(void* self);

#ifdef __cplusplus
}
#endif
