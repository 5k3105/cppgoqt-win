#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFontMetrics_new(void* arg1);
void* QFontMetrics_new2(void* font, void* pd);
void* QFontMetrics_new3(void* arg1);
void* QFontMetrics_OperatorAssign(void* self, void* arg1);
void QFontMetrics_swap(void* self, void* other);
int QFontMetrics_ascent(void* self);
int QFontMetrics_capHeight(void* self);
int QFontMetrics_descent(void* self);
int QFontMetrics_height(void* self);
int QFontMetrics_leading(void* self);
int QFontMetrics_lineSpacing(void* self);
int QFontMetrics_minLeftBearing(void* self);
int QFontMetrics_minRightBearing(void* self);
int QFontMetrics_maxWidth(void* self);
int QFontMetrics_xHeight(void* self);
int QFontMetrics_averageCharWidth(void* self);
int QFontMetrics_inFont(void* self, void* arg1);
int QFontMetrics_inFontUcs4(void* self, unsigned int ucs4);
int QFontMetrics_leftBearing(void* self, void* arg1);
int QFontMetrics_rightBearing(void* self, void* arg1);
int QFontMetrics_horizontalAdvance(void* self, const char* arg1, long long arg1_len);
int QFontMetrics_horizontalAdvance2(void* self, const char* arg1, long long arg1_len, void* textOption);
int QFontMetrics_horizontalAdvance3(void* self, void* arg1);
void* QFontMetrics_boundingRect(void* self, void* arg1);
void* QFontMetrics_boundingRect2(void* self, const char* text, long long text_len);
void* QFontMetrics_boundingRect3(void* self, const char* text, long long text_len, void* textOption);
void* QFontMetrics_boundingRect4(void* self, void* r, int flags, const char* text, long long text_len);
void* QFontMetrics_boundingRect5(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len);
void* QFontMetrics_size(void* self, int flags, const char* str, long long str_len);
void* QFontMetrics_tightBoundingRect(void* self, const char* text, long long text_len);
void* QFontMetrics_tightBoundingRect2(void* self, const char* text, long long text_len, void* textOption);
char* QFontMetrics_elidedText(void* self, const char* text, long long text_len, int mode, int width, long long* __out_len);
int QFontMetrics_underlinePos(void* self);
int QFontMetrics_overlinePos(void* self);
int QFontMetrics_strikeOutPos(void* self);
int QFontMetrics_lineWidth(void* self);
double QFontMetrics_fontDpi(void* self);
int QFontMetrics_OperatorEqual(void* self, void* other);
int QFontMetrics_OperatorNotEqual(void* self, void* other);
int QFontMetrics_horizontalAdvance4(void* self, const char* arg1, long long arg1_len, int len_);
void* QFontMetrics_boundingRect6(void* self, void* r, int flags, const char* text, long long text_len, int tabstops);
void* QFontMetrics_boundingRect7(void* self, void* r, int flags, const char* text, long long text_len, int tabstops, int* tabarray);
void* QFontMetrics_boundingRect8(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len, int tabstops);
void* QFontMetrics_boundingRect9(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len, int tabstops, int* tabarray);
void* QFontMetrics_size2(void* self, int flags, const char* str, long long str_len, int tabstops);
void* QFontMetrics_size3(void* self, int flags, const char* str, long long str_len, int tabstops, int* tabarray);
char* QFontMetrics_elidedText2(void* self, const char* text, long long text_len, int mode, int width, int flags, long long* __out_len);
void QFontMetrics_delete(void* self);

#ifdef __cplusplus
}
#endif
