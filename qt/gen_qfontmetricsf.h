#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFontMetricsF_new(void* font);
void* QFontMetricsF_new2(void* font, void* pd);
void* QFontMetricsF_new3(void* arg1);
void* QFontMetricsF_new4(void* arg1);
void* QFontMetricsF_OperatorAssign(void* self, void* arg1);
void* QFontMetricsF_OperatorAssign2(void* self, void* arg1);
void QFontMetricsF_swap(void* self, void* other);
double QFontMetricsF_ascent(void* self);
double QFontMetricsF_capHeight(void* self);
double QFontMetricsF_descent(void* self);
double QFontMetricsF_height(void* self);
double QFontMetricsF_leading(void* self);
double QFontMetricsF_lineSpacing(void* self);
double QFontMetricsF_minLeftBearing(void* self);
double QFontMetricsF_minRightBearing(void* self);
double QFontMetricsF_maxWidth(void* self);
double QFontMetricsF_xHeight(void* self);
double QFontMetricsF_averageCharWidth(void* self);
int QFontMetricsF_inFont(void* self, void* arg1);
int QFontMetricsF_inFontUcs4(void* self, unsigned int ucs4);
double QFontMetricsF_leftBearing(void* self, void* arg1);
double QFontMetricsF_rightBearing(void* self, void* arg1);
double QFontMetricsF_horizontalAdvance(void* self, const char* string_, long long string__len);
double QFontMetricsF_horizontalAdvance2(void* self, void* arg1);
double QFontMetricsF_horizontalAdvance3(void* self, const char* string_, long long string__len, void* textOption);
void* QFontMetricsF_boundingRect(void* self, const char* string_, long long string__len);
void* QFontMetricsF_boundingRect2(void* self, const char* text, long long text_len, void* textOption);
void* QFontMetricsF_boundingRect3(void* self, void* arg1);
void* QFontMetricsF_boundingRect4(void* self, void* r, int flags, const char* string_, long long string__len);
void* QFontMetricsF_size(void* self, int flags, const char* str, long long str_len);
void* QFontMetricsF_tightBoundingRect(void* self, const char* text, long long text_len);
void* QFontMetricsF_tightBoundingRect2(void* self, const char* text, long long text_len, void* textOption);
char* QFontMetricsF_elidedText(void* self, const char* text, long long text_len, int mode, double width, long long* __out_len);
double QFontMetricsF_underlinePos(void* self);
double QFontMetricsF_overlinePos(void* self);
double QFontMetricsF_strikeOutPos(void* self);
double QFontMetricsF_lineWidth(void* self);
double QFontMetricsF_fontDpi(void* self);
int QFontMetricsF_OperatorEqual(void* self, void* other);
int QFontMetricsF_OperatorNotEqual(void* self, void* other);
double QFontMetricsF_horizontalAdvance4(void* self, const char* string_, long long string__len, int length);
void* QFontMetricsF_boundingRect5(void* self, void* r, int flags, const char* string_, long long string__len, int tabstops);
void* QFontMetricsF_boundingRect6(void* self, void* r, int flags, const char* string_, long long string__len, int tabstops, int* tabarray);
void* QFontMetricsF_size2(void* self, int flags, const char* str, long long str_len, int tabstops);
void* QFontMetricsF_size3(void* self, int flags, const char* str, long long str_len, int tabstops, int* tabarray);
char* QFontMetricsF_elidedText2(void* self, const char* text, long long text_len, int mode, double width, int flags, long long* __out_len);
void QFontMetricsF_delete(void* self);

#ifdef __cplusplus
}
#endif
