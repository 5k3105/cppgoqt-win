#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextListFormat_base(void* self);
void* QTextListFormat_new();
void* QTextListFormat_new2(void* arg1);
int QTextListFormat_isValid(void* self);
void QTextListFormat_setStyle(void* self, int style);
int QTextListFormat_style(void* self);
void QTextListFormat_setIndent(void* self, int indent);
int QTextListFormat_indent(void* self);
void QTextListFormat_setNumberPrefix(void* self, const char* numberPrefix, long long numberPrefix_len);
char* QTextListFormat_numberPrefix(void* self, long long* __out_len);
void QTextListFormat_setNumberSuffix(void* self, const char* numberSuffix, long long numberSuffix_len);
char* QTextListFormat_numberSuffix(void* self, long long* __out_len);
void QTextListFormat_setStart(void* self, int indent);
int QTextListFormat_start(void* self);
void* QTextListFormat_OperatorAssign(void* self, void* arg1);
void QTextListFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
