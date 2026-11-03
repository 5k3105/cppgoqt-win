#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextDocumentWriter_new();
void* QTextDocumentWriter_new2(void* device, void* format);
void* QTextDocumentWriter_new3(const char* fileName, long long fileName_len);
void* QTextDocumentWriter_new4(const char* fileName, long long fileName_len, void* format);
void QTextDocumentWriter_setFormat(void* self, void* format);
void* QTextDocumentWriter_format(void* self);
void QTextDocumentWriter_setDevice(void* self, void* device);
void* QTextDocumentWriter_device(void* self);
void QTextDocumentWriter_setFileName(void* self, const char* fileName, long long fileName_len);
char* QTextDocumentWriter_fileName(void* self, long long* __out_len);
int QTextDocumentWriter_write(void* self, void* document);
int QTextDocumentWriter_write2(void* self, void* fragment);
void** QTextDocumentWriter_supportedDocumentFormats(long long* __out_len);
void QTextDocumentWriter_delete(void* self);

#ifdef __cplusplus
}
#endif
