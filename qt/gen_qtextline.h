#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextLine_new();
int QTextLine_isValid(void* self);
void* QTextLine_rect(void* self);
double QTextLine_x(void* self);
double QTextLine_y(void* self);
double QTextLine_width(void* self);
double QTextLine_ascent(void* self);
double QTextLine_descent(void* self);
double QTextLine_height(void* self);
double QTextLine_leading(void* self);
void QTextLine_setLeadingIncluded(void* self, int included);
int QTextLine_leadingIncluded(void* self);
double QTextLine_naturalTextWidth(void* self);
double QTextLine_horizontalAdvance(void* self);
void* QTextLine_naturalTextRect(void* self);
double QTextLine_cursorToX(void* self, int* cursorPos);
double QTextLine_cursorToX2(void* self, int cursorPos);
int QTextLine_xToCursor(void* self, double x);
void QTextLine_setLineWidth(void* self, double width);
void QTextLine_setNumColumns(void* self, int columns);
void QTextLine_setNumColumns2(void* self, int columns, double alignmentWidth);
void QTextLine_setPosition(void* self, void* pos);
void* QTextLine_position(void* self);
int QTextLine_textStart(void* self);
int QTextLine_textLength(void* self);
int QTextLine_lineNumber(void* self);
void QTextLine_draw(void* self, void* painter, void* position);
void** QTextLine_glyphRuns(void* self, int from, int length, int flags, long long* __out_len);
void** QTextLine_glyphRuns2(void* self, long long* __out_len);
double QTextLine_cursorToX3(void* self, int* cursorPos, int edge);
double QTextLine_cursorToX4(void* self, int cursorPos, int edge);
int QTextLine_xToCursor2(void* self, double x, int arg2);
void** QTextLine_glyphRuns3(void* self, int from, long long* __out_len);
void** QTextLine_glyphRuns4(void* self, int from, int length, long long* __out_len);
void QTextLine_delete(void* self);

#ifdef __cplusplus
}
#endif
