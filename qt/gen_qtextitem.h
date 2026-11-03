#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextItem_new();
void* QTextItem_new2(void* arg1);
double QTextItem_descent(void* self);
double QTextItem_ascent(void* self);
double QTextItem_width(void* self);
int QTextItem_renderFlags(void* self);
char* QTextItem_text(void* self, long long* __out_len);
void* QTextItem_font(void* self);
void QTextItem_delete(void* self);

#ifdef __cplusplus
}
#endif
