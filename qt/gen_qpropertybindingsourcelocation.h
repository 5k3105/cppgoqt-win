#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyBindingSourceLocation_new();
void* QPropertyBindingSourceLocation_new2(void* arg1);
const char* QPropertyBindingSourceLocation_fileName_get(void* self);
void QPropertyBindingSourceLocation_fileName_set(void* self, const char* fileName);
const char* QPropertyBindingSourceLocation_functionName_get(void* self);
void QPropertyBindingSourceLocation_functionName_set(void* self, const char* functionName);
unsigned int QPropertyBindingSourceLocation_line_get(void* self);
void QPropertyBindingSourceLocation_line_set(void* self, unsigned int line);
unsigned int QPropertyBindingSourceLocation_column_get(void* self);
void QPropertyBindingSourceLocation_column_set(void* self, unsigned int column);
void QPropertyBindingSourceLocation_delete(void* self);

#ifdef __cplusplus
}
#endif
