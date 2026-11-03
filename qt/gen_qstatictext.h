#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStaticText_new();
void* QStaticText_new2(const char* text, long long text_len);
void* QStaticText_new3(void* other);
void* QStaticText_OperatorAssign(void* self, void* arg1);
void QStaticText_swap(void* self, void* other);
void QStaticText_setText(void* self, const char* text, long long text_len);
char* QStaticText_text(void* self, long long* __out_len);
void QStaticText_setTextFormat(void* self, int textFormat);
int QStaticText_textFormat(void* self);
void QStaticText_setTextWidth(void* self, double textWidth);
double QStaticText_textWidth(void* self);
void QStaticText_setTextOption(void* self, void* textOption);
void* QStaticText_textOption(void* self);
void* QStaticText_size(void* self);
void QStaticText_prepare(void* self);
void QStaticText_setPerformanceHint(void* self, int performanceHint);
int QStaticText_performanceHint(void* self);
int QStaticText_OperatorEqual(void* self, void* arg1);
int QStaticText_OperatorNotEqual(void* self, void* arg1);
void QStaticText_prepare2(void* self, void* matrix);
void QStaticText_prepare3(void* self, void* matrix, void* font);
void QStaticText_delete(void* self);

#ifdef __cplusplus
}
#endif
