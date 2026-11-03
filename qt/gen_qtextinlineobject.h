#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextInlineObject_new();
int QTextInlineObject_isValid(void* self);
void* QTextInlineObject_rect(void* self);
double QTextInlineObject_width(void* self);
double QTextInlineObject_ascent(void* self);
double QTextInlineObject_descent(void* self);
double QTextInlineObject_height(void* self);
int QTextInlineObject_textDirection(void* self);
void QTextInlineObject_setWidth(void* self, double w);
void QTextInlineObject_setAscent(void* self, double a);
void QTextInlineObject_setDescent(void* self, double d);
int QTextInlineObject_textPosition(void* self);
int QTextInlineObject_formatIndex(void* self);
void* QTextInlineObject_format(void* self);
void QTextInlineObject_delete(void* self);

#ifdef __cplusplus
}
#endif
