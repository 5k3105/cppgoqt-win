#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

char* QStandardPaths_writableLocation(int type_, long long* __out_len);
char** QStandardPaths_standardLocations(int type_, long long* __out_len);
char* QStandardPaths_locate(int type_, const char* fileName, long long fileName_len, long long* __out_len);
char** QStandardPaths_locateAll(int type_, const char* fileName, long long fileName_len, long long* __out_len);
char* QStandardPaths_displayName(int type_, long long* __out_len);
char* QStandardPaths_findExecutable(const char* executableName, long long executableName_len, long long* __out_len);
void QStandardPaths_setTestModeEnabled(int testMode);
int QStandardPaths_isTestModeEnabled();
char* QStandardPaths_locate2(int type_, const char* fileName, long long fileName_len, int options, long long* __out_len);
char** QStandardPaths_locateAll2(int type_, const char* fileName, long long fileName_len, int options, long long* __out_len);
char* QStandardPaths_findExecutable2(const char* executableName, long long executableName_len, char** paths, long long paths_len, long long* __out_len);
void QStandardPaths_delete(void* self);

#ifdef __cplusplus
}
#endif
