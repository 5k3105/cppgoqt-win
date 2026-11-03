#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextDocumentFragment_new();
void* QTextDocumentFragment_new2(void* document);
void* QTextDocumentFragment_new3(void* range_);
void* QTextDocumentFragment_new4(void* rhs);
void* QTextDocumentFragment_OperatorAssign(void* self, void* rhs);
int QTextDocumentFragment_isEmpty(void* self);
char* QTextDocumentFragment_toPlainText(void* self, long long* __out_len);
char* QTextDocumentFragment_toRawText(void* self, long long* __out_len);
char* QTextDocumentFragment_toHtml(void* self, long long* __out_len);
char* QTextDocumentFragment_toMarkdown(void* self, long long* __out_len);
void* QTextDocumentFragment_fromPlainText(const char* plainText, long long plainText_len);
void* QTextDocumentFragment_fromHtml(const char* html, long long html_len);
void* QTextDocumentFragment_fromMarkdown(const char* markdown, long long markdown_len);
char* QTextDocumentFragment_toMarkdown2(void* self, int features, long long* __out_len);
void* QTextDocumentFragment_fromHtml2(const char* html, long long html_len, void* resourceProvider);
void* QTextDocumentFragment_fromMarkdown2(const char* markdown, long long markdown_len, int features);
void QTextDocumentFragment_delete(void* self);

#ifdef __cplusplus
}
#endif
