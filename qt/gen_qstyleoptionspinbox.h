#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionSpinBox_base(void* self);
void QStyleOptionSpinBox_chain(void* self, void** out);
void* QStyleOptionSpinBox_new();
void* QStyleOptionSpinBox_new2(void* other);
void* QStyleOptionSpinBox_OperatorAssign(void* self, void* arg1);
int QStyleOptionSpinBox_buttonSymbols_get(void* self);
void QStyleOptionSpinBox_buttonSymbols_set(void* self, int buttonSymbols);
int QStyleOptionSpinBox_stepEnabled_get(void* self);
void QStyleOptionSpinBox_stepEnabled_set(void* self, int stepEnabled);
int QStyleOptionSpinBox_frame_get(void* self);
void QStyleOptionSpinBox_frame_set(void* self, int frame);
void QStyleOptionSpinBox_delete(void* self);

#ifdef __cplusplus
}
#endif
