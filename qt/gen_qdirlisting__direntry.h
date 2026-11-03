#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDirListing__DirEntry_new(void* arg1);
void* QDirListing__DirEntry_new2();
char* QDirListing__DirEntry_fileName(void* self, long long* __out_len);
char* QDirListing__DirEntry_baseName(void* self, long long* __out_len);
char* QDirListing__DirEntry_completeBaseName(void* self, long long* __out_len);
char* QDirListing__DirEntry_suffix(void* self, long long* __out_len);
char* QDirListing__DirEntry_bundleName(void* self, long long* __out_len);
char* QDirListing__DirEntry_completeSuffix(void* self, long long* __out_len);
char* QDirListing__DirEntry_filePath(void* self, long long* __out_len);
int QDirListing__DirEntry_isDir(void* self);
int QDirListing__DirEntry_isFile(void* self);
int QDirListing__DirEntry_isSymLink(void* self);
int QDirListing__DirEntry_exists(void* self);
int QDirListing__DirEntry_isHidden(void* self);
int QDirListing__DirEntry_isReadable(void* self);
int QDirListing__DirEntry_isWritable(void* self);
int QDirListing__DirEntry_isExecutable(void* self);
void* QDirListing__DirEntry_fileInfo(void* self);
char* QDirListing__DirEntry_canonicalFilePath(void* self, long long* __out_len);
char* QDirListing__DirEntry_absoluteFilePath(void* self, long long* __out_len);
char* QDirListing__DirEntry_absolutePath(void* self, long long* __out_len);
long long QDirListing__DirEntry_size(void* self);
void* QDirListing__DirEntry_birthTime(void* self, void* tz);
void* QDirListing__DirEntry_metadataChangeTime(void* self, void* tz);
void* QDirListing__DirEntry_lastModified(void* self, void* tz);
void* QDirListing__DirEntry_lastRead(void* self, void* tz);
void* QDirListing__DirEntry_fileTime(void* self, int type_, void* tz);
void* QDirListing__DirEntry_OperatorAssign(void* self, void* arg1);
void QDirListing__DirEntry_delete(void* self);

#ifdef __cplusplus
}
#endif
