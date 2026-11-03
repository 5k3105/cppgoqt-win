#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStorageInfo_new();
void* QStorageInfo_new2(const char* path, long long path_len);
void* QStorageInfo_new3(void* dir);
void* QStorageInfo_new4(void* other);
void* QStorageInfo_OperatorAssign(void* self, void* other);
void QStorageInfo_swap(void* self, void* other);
void QStorageInfo_setPath(void* self, const char* path, long long path_len);
char* QStorageInfo_rootPath(void* self, long long* __out_len);
void* QStorageInfo_device(void* self);
void* QStorageInfo_subvolume(void* self);
void* QStorageInfo_fileSystemType(void* self);
char* QStorageInfo_name(void* self, long long* __out_len);
char* QStorageInfo_displayName(void* self, long long* __out_len);
long long QStorageInfo_bytesTotal(void* self);
long long QStorageInfo_bytesFree(void* self);
long long QStorageInfo_bytesAvailable(void* self);
int QStorageInfo_blockSize(void* self);
int QStorageInfo_isRoot(void* self);
int QStorageInfo_isReadOnly(void* self);
int QStorageInfo_isReady(void* self);
int QStorageInfo_isValid(void* self);
void QStorageInfo_refresh(void* self);
void** QStorageInfo_mountedVolumes(long long* __out_len);
void* QStorageInfo_root();
int QStorageInfo_OperatorEqual(void* self, void* rhs);
int QStorageInfo_OperatorNotEqual(void* self, void* rhs);
void QStorageInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
