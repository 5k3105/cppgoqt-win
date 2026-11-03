#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFontInfo_new(void* arg1);
void* QFontInfo_new2(void* arg1);
void* QFontInfo_OperatorAssign(void* self, void* arg1);
void QFontInfo_swap(void* self, void* other);
char* QFontInfo_family(void* self, long long* __out_len);
char* QFontInfo_styleName(void* self, long long* __out_len);
int QFontInfo_pixelSize(void* self);
int QFontInfo_pointSize(void* self);
double QFontInfo_pointSizeF(void* self);
int QFontInfo_italic(void* self);
int QFontInfo_style(void* self);
int QFontInfo_weight(void* self);
int QFontInfo_bold(void* self);
int QFontInfo_underline(void* self);
int QFontInfo_overline(void* self);
int QFontInfo_strikeOut(void* self);
int QFontInfo_fixedPitch(void* self);
int QFontInfo_styleHint(void* self);
void** QFontInfo_variableAxes(void* self, long long* __out_len);
int QFontInfo_legacyWeight(void* self);
int QFontInfo_exactMatch(void* self);
void QFontInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
