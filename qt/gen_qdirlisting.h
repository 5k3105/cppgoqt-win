#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDirListing_new(const char* path, long long path_len);
void* QDirListing_new2(const char* path, long long path_len, char** nameFilters, long long nameFilters_len);
void* QDirListing_new3(const char* path, long long path_len, int flags);
void* QDirListing_new4(const char* path, long long path_len, char** nameFilters, long long nameFilters_len, int flags);
void QDirListing_swap(void* self, void* other);
char* QDirListing_iteratorPath(void* self, long long* __out_len);
int QDirListing_iteratorFlags(void* self);
char** QDirListing_nameFilters(void* self, long long* __out_len);
void* QDirListing_end(void* self);
void* QDirListing_cend(void* self);
void* QDirListing_constEnd(void* self);
void QDirListing_delete(void* self);

#ifdef __cplusplus
}
#endif
