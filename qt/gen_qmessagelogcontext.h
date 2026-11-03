#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMessageLogContext_new();
void* QMessageLogContext_new2(const char* fileName, int lineNumber, const char* functionName, const char* categoryName);
int QMessageLogContext_version_get(void* self);
void QMessageLogContext_version_set(void* self, int version);
int QMessageLogContext_line_get(void* self);
void QMessageLogContext_line_set(void* self, int line);
const char* QMessageLogContext_file_get(void* self);
void QMessageLogContext_file_set(void* self, const char* file);
const char* QMessageLogContext_function_get(void* self);
void QMessageLogContext_function_set(void* self, const char* function);
const char* QMessageLogContext_category_get(void* self);
void QMessageLogContext_category_set(void* self, const char* category);
void QMessageLogContext_delete(void* self);

#ifdef __cplusplus
}
#endif
