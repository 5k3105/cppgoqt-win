#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLockFile_new(const char* fileName, long long fileName_len);
char* QLockFile_fileName(void* self, long long* __out_len);
int QLockFile_lock(void* self);
int QLockFile_tryLock(void* self, int timeout);
void QLockFile_unlock(void* self);
void QLockFile_setStaleLockTime(void* self, int arg1);
int QLockFile_staleLockTime(void* self);
int QLockFile_tryLock2(void* self);
void QLockFile_setStaleLockTime2(void* self, long long value);
long long QLockFile_staleLockTimeAsDuration(void* self);
int QLockFile_isLocked(void* self);
int QLockFile_getLockInfo(void* self, long long* pid, void* hostname, void* appname);
int QLockFile_removeStaleLockFile(void* self);
int QLockFile_error(void* self);
int QLockFile_tryLock3(void* self, long long timeout);
void QLockFile_delete(void* self);

#ifdef __cplusplus
}
#endif
