#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QResource_new();
void* QResource_new2(const char* file, long long file_len);
void* QResource_new3(const char* file, long long file_len, void* locale);
void QResource_setFileName(void* self, const char* file, long long file_len);
char* QResource_fileName(void* self, long long* __out_len);
char* QResource_absoluteFilePath(void* self, long long* __out_len);
void QResource_setLocale(void* self, void* locale);
void* QResource_locale(void* self);
int QResource_isValid(void* self);
int QResource_compressionAlgorithm(void* self);
long long QResource_size(void* self);
long long QResource_uncompressedSize(void* self);
void* QResource_uncompressedData(void* self);
void* QResource_lastModified(void* self);
int QResource_registerResource(const char* rccFilename, long long rccFilename_len);
int QResource_unregisterResource(const char* rccFilename, long long rccFilename_len);
int QResource_registerResource2(const unsigned char* rccData);
int QResource_unregisterResource2(const unsigned char* rccData);
int QResource_registerResource3(const char* rccFilename, long long rccFilename_len, const char* resourceRoot, long long resourceRoot_len);
int QResource_unregisterResource3(const char* rccFilename, long long rccFilename_len, const char* resourceRoot, long long resourceRoot_len);
int QResource_registerResource4(const unsigned char* rccData, const char* resourceRoot, long long resourceRoot_len);
int QResource_unregisterResource4(const unsigned char* rccData, const char* resourceRoot, long long resourceRoot_len);
void QResource_delete(void* self);

#ifdef __cplusplus
}
#endif
