#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMimeDatabase_new();
void* QMimeDatabase_mimeTypeForName(void* self, const char* nameOrAlias, long long nameOrAlias_len);
void* QMimeDatabase_mimeTypeForFile(void* self, const char* fileName, long long fileName_len);
void* QMimeDatabase_mimeTypeForFile2(void* self, void* fileInfo);
void** QMimeDatabase_mimeTypesForFileName(void* self, const char* fileName, long long fileName_len, long long* __out_len);
void* QMimeDatabase_mimeTypeForData(void* self, void* data);
void* QMimeDatabase_mimeTypeForData2(void* self, void* device);
void* QMimeDatabase_mimeTypeForUrl(void* self, void* url);
void* QMimeDatabase_mimeTypeForFileNameAndData(void* self, const char* fileName, long long fileName_len, void* device);
void* QMimeDatabase_mimeTypeForFileNameAndData2(void* self, const char* fileName, long long fileName_len, void* data);
char* QMimeDatabase_suffixForFileName(void* self, const char* fileName, long long fileName_len, long long* __out_len);
void** QMimeDatabase_allMimeTypes(void* self, long long* __out_len);
void* QMimeDatabase_mimeTypeForFile3(void* self, const char* fileName, long long fileName_len, int mode);
void* QMimeDatabase_mimeTypeForFile4(void* self, void* fileInfo, int mode);
void QMimeDatabase_delete(void* self);

#ifdef __cplusplus
}
#endif
