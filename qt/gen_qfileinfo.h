#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFileInfo_new();
void* QFileInfo_new2(const char* file, long long file_len);
void* QFileInfo_new3(void* file);
void* QFileInfo_new4(void* dir, const char* file, long long file_len);
void* QFileInfo_new5(void* fileinfo);
void* QFileInfo_OperatorAssign(void* self, void* fileinfo);
void QFileInfo_swap(void* self, void* other);
void QFileInfo_setFile(void* self, const char* file, long long file_len);
void QFileInfo_setFile2(void* self, void* file);
void QFileInfo_setFile3(void* self, void* dir, const char* file, long long file_len);
int QFileInfo_exists(void* self);
int QFileInfo_exists2(const char* file, long long file_len);
void QFileInfo_refresh(void* self);
char* QFileInfo_filePath(void* self, long long* __out_len);
char* QFileInfo_absoluteFilePath(void* self, long long* __out_len);
char* QFileInfo_canonicalFilePath(void* self, long long* __out_len);
char* QFileInfo_fileName(void* self, long long* __out_len);
char* QFileInfo_baseName(void* self, long long* __out_len);
char* QFileInfo_completeBaseName(void* self, long long* __out_len);
char* QFileInfo_suffix(void* self, long long* __out_len);
char* QFileInfo_bundleName(void* self, long long* __out_len);
char* QFileInfo_completeSuffix(void* self, long long* __out_len);
char* QFileInfo_path(void* self, long long* __out_len);
char* QFileInfo_absolutePath(void* self, long long* __out_len);
char* QFileInfo_canonicalPath(void* self, long long* __out_len);
void* QFileInfo_dir(void* self);
void* QFileInfo_absoluteDir(void* self);
int QFileInfo_isReadable(void* self);
int QFileInfo_isWritable(void* self);
int QFileInfo_isExecutable(void* self);
int QFileInfo_isHidden(void* self);
int QFileInfo_isNativePath(void* self);
int QFileInfo_isRelative(void* self);
int QFileInfo_isAbsolute(void* self);
int QFileInfo_makeAbsolute(void* self);
int QFileInfo_isFile(void* self);
int QFileInfo_isDir(void* self);
int QFileInfo_isSymLink(void* self);
int QFileInfo_isSymbolicLink(void* self);
int QFileInfo_isOther(void* self);
int QFileInfo_isShortcut(void* self);
int QFileInfo_isAlias(void* self);
int QFileInfo_isJunction(void* self);
int QFileInfo_isRoot(void* self);
int QFileInfo_isBundle(void* self);
char* QFileInfo_symLinkTarget(void* self, long long* __out_len);
char* QFileInfo_readSymLink(void* self, long long* __out_len);
char* QFileInfo_junctionTarget(void* self, long long* __out_len);
char* QFileInfo_owner(void* self, long long* __out_len);
unsigned int QFileInfo_ownerId(void* self);
char* QFileInfo_group(void* self, long long* __out_len);
unsigned int QFileInfo_groupId(void* self);
int QFileInfo_permission(void* self, int permissions);
int QFileInfo_permissions(void* self);
long long QFileInfo_size(void* self);
void* QFileInfo_birthTime(void* self);
void* QFileInfo_metadataChangeTime(void* self);
void* QFileInfo_lastModified(void* self);
void* QFileInfo_lastRead(void* self);
void* QFileInfo_fileTime(void* self, int time);
void* QFileInfo_birthTime2(void* self, void* tz);
void* QFileInfo_metadataChangeTime2(void* self, void* tz);
void* QFileInfo_lastModified2(void* self, void* tz);
void* QFileInfo_lastRead2(void* self, void* tz);
void* QFileInfo_fileTime2(void* self, int time, void* tz);
int QFileInfo_caching(void* self);
void QFileInfo_setCaching(void* self, int on);
void QFileInfo_stat(void* self);
int QFileInfo_OperatorEqual(void* self, void* rhs);
int QFileInfo_OperatorNotEqual(void* self, void* rhs);
void QFileInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
