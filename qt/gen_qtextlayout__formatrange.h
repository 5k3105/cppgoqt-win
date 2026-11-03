#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextLayout__FormatRange_new();
void* QTextLayout__FormatRange_new2(void* arg1);
int QTextLayout__FormatRange_OperatorEqual(void* self, void* rhs);
int QTextLayout__FormatRange_OperatorNotEqual(void* self, void* rhs);
int QTextLayout__FormatRange_start_get(void* self);
void QTextLayout__FormatRange_start_set(void* self, int start);
int QTextLayout__FormatRange_length_get(void* self);
void QTextLayout__FormatRange_length_set(void* self, int length);
void* QTextLayout__FormatRange_format_get(void* self);
void QTextLayout__FormatRange_format_set(void* self, void* format);
void QTextLayout__FormatRange_delete(void* self);

#ifdef __cplusplus
}
#endif
