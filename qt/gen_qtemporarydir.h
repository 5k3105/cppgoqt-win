#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTemporaryDir_new();
void* QTemporaryDir_new2(const char* templateName, long long templateName_len);
void QTemporaryDir_swap(void* self, void* other);
int QTemporaryDir_isValid(void* self);
char* QTemporaryDir_errorString(void* self, long long* __out_len);
int QTemporaryDir_autoRemove(void* self);
void QTemporaryDir_setAutoRemove(void* self, int b);
int QTemporaryDir_remove(void* self);
char* QTemporaryDir_path(void* self, long long* __out_len);
char* QTemporaryDir_filePath(void* self, const char* fileName, long long fileName_len, long long* __out_len);
void QTemporaryDir_delete(void* self);

#ifdef __cplusplus
}
#endif
