#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionProgressBar_base(void* self);
void* QStyleOptionProgressBar_new();
void* QStyleOptionProgressBar_new2(void* other);
void* QStyleOptionProgressBar_OperatorAssign(void* self, void* arg1);
int QStyleOptionProgressBar_minimum_get(void* self);
void QStyleOptionProgressBar_minimum_set(void* self, int minimum);
int QStyleOptionProgressBar_maximum_get(void* self);
void QStyleOptionProgressBar_maximum_set(void* self, int maximum);
int QStyleOptionProgressBar_progress_get(void* self);
void QStyleOptionProgressBar_progress_set(void* self, int progress);
char* QStyleOptionProgressBar_text_get(void* self, long long* __out_len);
void QStyleOptionProgressBar_text_set(void* self, const char* text, long long text_len);
int QStyleOptionProgressBar_textAlignment_get(void* self);
void QStyleOptionProgressBar_textAlignment_set(void* self, int textAlignment);
int QStyleOptionProgressBar_textVisible_get(void* self);
void QStyleOptionProgressBar_textVisible_set(void* self, int textVisible);
int QStyleOptionProgressBar_invertedAppearance_get(void* self);
void QStyleOptionProgressBar_invertedAppearance_set(void* self, int invertedAppearance);
int QStyleOptionProgressBar_bottomToTop_get(void* self);
void QStyleOptionProgressBar_bottomToTop_set(void* self, int bottomToTop);
void QStyleOptionProgressBar_delete(void* self);

#ifdef __cplusplus
}
#endif
