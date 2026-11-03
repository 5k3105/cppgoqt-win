#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionButton_base(void* self);
void* QStyleOptionButton_new();
void* QStyleOptionButton_new2(void* other);
void* QStyleOptionButton_OperatorAssign(void* self, void* arg1);
int QStyleOptionButton_features_get(void* self);
void QStyleOptionButton_features_set(void* self, int features);
char* QStyleOptionButton_text_get(void* self, long long* __out_len);
void QStyleOptionButton_text_set(void* self, const char* text, long long text_len);
void* QStyleOptionButton_icon_get(void* self);
void QStyleOptionButton_icon_set(void* self, void* icon);
void* QStyleOptionButton_iconSize_get(void* self);
void QStyleOptionButton_iconSize_set(void* self, void* iconSize);
void QStyleOptionButton_delete(void* self);

#ifdef __cplusplus
}
#endif
