#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDirIterator_new(void* dir);
void* QDirIterator_new2(const char* path, long long path_len);
void* QDirIterator_new3(const char* path, long long path_len, int filter);
void* QDirIterator_new4(const char* path, long long path_len, char** nameFilters, long long nameFilters_len);
void* QDirIterator_new5(void* dir, int flags);
void* QDirIterator_new6(const char* path, long long path_len, int flags);
void* QDirIterator_new7(const char* path, long long path_len, int filter, int flags);
void* QDirIterator_new8(const char* path, long long path_len, char** nameFilters, long long nameFilters_len, int filters);
void* QDirIterator_new9(const char* path, long long path_len, char** nameFilters, long long nameFilters_len, int filters, int flags);
char* QDirIterator_next(void* self, long long* __out_len);
void* QDirIterator_nextFileInfo(void* self);
int QDirIterator_hasNext(void* self);
char* QDirIterator_fileName(void* self, long long* __out_len);
char* QDirIterator_filePath(void* self, long long* __out_len);
void* QDirIterator_fileInfo(void* self);
char* QDirIterator_path(void* self, long long* __out_len);
void QDirIterator_delete(void* self);

#ifdef __cplusplus
}
#endif
