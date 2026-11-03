#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

const char* QLibraryInfo_build();
int QLibraryInfo_isDebugBuild();
int QLibraryInfo_isSharedBuild();
void* QLibraryInfo_version();
char* QLibraryInfo_path(int p, long long* __out_len);
char** QLibraryInfo_paths(int p, long long* __out_len);
char* QLibraryInfo_location(int location, long long* __out_len);
char** QLibraryInfo_platformPluginArguments(const char* platformName, long long platformName_len, long long* __out_len);
void QLibraryInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
