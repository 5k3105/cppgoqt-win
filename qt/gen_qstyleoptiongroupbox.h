#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionGroupBox_base(void* self);
void QStyleOptionGroupBox_chain(void* self, void** out);
void* QStyleOptionGroupBox_new();
void* QStyleOptionGroupBox_new2(void* other);
void* QStyleOptionGroupBox_OperatorAssign(void* self, void* arg1);
int QStyleOptionGroupBox_features_get(void* self);
void QStyleOptionGroupBox_features_set(void* self, int features);
char* QStyleOptionGroupBox_text_get(void* self, long long* __out_len);
void QStyleOptionGroupBox_text_set(void* self, const char* text, long long text_len);
int QStyleOptionGroupBox_textAlignment_get(void* self);
void QStyleOptionGroupBox_textAlignment_set(void* self, int textAlignment);
void* QStyleOptionGroupBox_textColor_get(void* self);
void QStyleOptionGroupBox_textColor_set(void* self, void* textColor);
int QStyleOptionGroupBox_lineWidth_get(void* self);
void QStyleOptionGroupBox_lineWidth_set(void* self, int lineWidth);
int QStyleOptionGroupBox_midLineWidth_get(void* self);
void QStyleOptionGroupBox_midLineWidth_set(void* self, int midLineWidth);
void QStyleOptionGroupBox_delete(void* self);

#ifdef __cplusplus
}
#endif
