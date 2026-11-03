#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextImageFormat_base(void* self);
void QTextImageFormat_chain(void* self, void** out);
void* QTextImageFormat_new();
void* QTextImageFormat_new2(void* arg1);
int QTextImageFormat_isValid(void* self);
void QTextImageFormat_setName(void* self, const char* name, long long name_len);
char* QTextImageFormat_name(void* self, long long* __out_len);
void QTextImageFormat_setWidth(void* self, double width);
double QTextImageFormat_width(void* self);
void QTextImageFormat_setMaximumWidth(void* self, void* maxWidth);
void* QTextImageFormat_maximumWidth(void* self);
void QTextImageFormat_setHeight(void* self, double height);
double QTextImageFormat_height(void* self);
void QTextImageFormat_setQuality(void* self, int quality);
void QTextImageFormat_setQuality2(void* self);
int QTextImageFormat_quality(void* self);
void QTextImageFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
