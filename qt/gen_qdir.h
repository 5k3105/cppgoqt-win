#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDir_new(void* arg1);
void* QDir_new2();
void* QDir_new3(const char* path, long long path_len, const char* nameFilter, long long nameFilter_len);
void* QDir_new4(const char* path, long long path_len);
void* QDir_new5(const char* path, long long path_len, const char* nameFilter, long long nameFilter_len, int sort);
void* QDir_new6(const char* path, long long path_len, const char* nameFilter, long long nameFilter_len, int sort, int filter);
void* QDir_OperatorAssign(void* self, void* arg1);
void QDir_swap(void* self, void* other);
void QDir_setPath(void* self, const char* path, long long path_len);
char* QDir_path(void* self, long long* __out_len);
char* QDir_absolutePath(void* self, long long* __out_len);
char* QDir_canonicalPath(void* self, long long* __out_len);
void QDir_setSearchPaths(const char* prefix, long long prefix_len, char** searchPaths, long long searchPaths_len);
void QDir_addSearchPath(const char* prefix, long long prefix_len, const char* path, long long path_len);
char** QDir_searchPaths(const char* prefix, long long prefix_len, long long* __out_len);
char* QDir_dirName(void* self, long long* __out_len);
char* QDir_filePath(void* self, const char* fileName, long long fileName_len, long long* __out_len);
char* QDir_absoluteFilePath(void* self, const char* fileName, long long fileName_len, long long* __out_len);
char* QDir_relativeFilePath(void* self, const char* fileName, long long fileName_len, long long* __out_len);
char* QDir_toNativeSeparators(const char* pathName, long long pathName_len, long long* __out_len);
char* QDir_fromNativeSeparators(const char* pathName, long long pathName_len, long long* __out_len);
int QDir_cd(void* self, const char* dirName, long long dirName_len);
int QDir_cdUp(void* self);
char** QDir_nameFilters(void* self, long long* __out_len);
void QDir_setNameFilters(void* self, char** nameFilters, long long nameFilters_len);
int QDir_filter(void* self);
void QDir_setFilter(void* self, int filter);
int QDir_sorting(void* self);
void QDir_setSorting(void* self, int sort);
long long QDir_count(void* self);
int QDir_isEmpty(void* self);
char* QDir_OperatorSubscript(void* self, long long arg1, long long* __out_len);
char** QDir_nameFiltersFromString(const char* nameFilter, long long nameFilter_len, long long* __out_len);
char** QDir_entryList(void* self, long long* __out_len);
char** QDir_entryList2(void* self, char** nameFilters, long long nameFilters_len, long long* __out_len);
void** QDir_entryInfoList(void* self, long long* __out_len);
void** QDir_entryInfoList2(void* self, char** nameFilters, long long nameFilters_len, long long* __out_len);
int QDir_mkdir(void* self, const char* dirName, long long dirName_len);
int QDir_rmdir(void* self, const char* dirName, long long dirName_len);
int QDir_mkpath(void* self, const char* dirPath, long long dirPath_len);
int QDir_rmpath(void* self, const char* dirPath, long long dirPath_len);
int QDir_removeRecursively(void* self);
int QDir_isReadable(void* self);
int QDir_exists(void* self);
int QDir_isRoot(void* self);
int QDir_isRelativePath(const char* path, long long path_len);
int QDir_isAbsolutePath(const char* path, long long path_len);
int QDir_isRelative(void* self);
int QDir_isAbsolute(void* self);
int QDir_makeAbsolute(void* self);
int QDir_remove(void* self, const char* fileName, long long fileName_len);
int QDir_rename(void* self, const char* oldName, long long oldName_len, const char* newName, long long newName_len);
int QDir_exists2(void* self, const char* name, long long name_len);
void** QDir_drives(long long* __out_len);
void* QDir_listSeparator();
void* QDir_separator();
int QDir_setCurrent(const char* path, long long path_len);
void* QDir_current();
char* QDir_currentPath(long long* __out_len);
void* QDir_home();
char* QDir_homePath(long long* __out_len);
void* QDir_root();
char* QDir_rootPath(long long* __out_len);
void* QDir_temp();
char* QDir_tempPath(long long* __out_len);
int QDir_match(char** filters, long long filters_len, const char* fileName, long long fileName_len);
int QDir_match2(const char* filter, long long filter_len, const char* fileName, long long fileName_len);
char* QDir_cleanPath(const char* path, long long path_len, long long* __out_len);
void QDir_refresh(void* self);
int QDir_isEmpty2(void* self, int filters);
char** QDir_entryList3(void* self, int filters, long long* __out_len);
char** QDir_entryList4(void* self, int filters, int sort, long long* __out_len);
char** QDir_entryList5(void* self, char** nameFilters, long long nameFilters_len, int filters, long long* __out_len);
char** QDir_entryList6(void* self, char** nameFilters, long long nameFilters_len, int filters, int sort, long long* __out_len);
void** QDir_entryInfoList3(void* self, int filters, long long* __out_len);
void** QDir_entryInfoList4(void* self, int filters, int sort, long long* __out_len);
void** QDir_entryInfoList5(void* self, char** nameFilters, long long nameFilters_len, int filters, long long* __out_len);
void** QDir_entryInfoList6(void* self, char** nameFilters, long long nameFilters_len, int filters, int sort, long long* __out_len);
int QDir_OperatorEqual(void* self, void* rhs);
int QDir_OperatorNotEqual(void* self, void* rhs);
void QDir_delete(void* self);

#ifdef __cplusplus
}
#endif
