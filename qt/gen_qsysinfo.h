#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

char* QSysInfo_buildCpuArchitecture(long long* __out_len);
char* QSysInfo_currentCpuArchitecture(long long* __out_len);
char* QSysInfo_buildAbi(long long* __out_len);
char* QSysInfo_kernelType(long long* __out_len);
char* QSysInfo_kernelVersion(long long* __out_len);
char* QSysInfo_productType(long long* __out_len);
char* QSysInfo_productVersion(long long* __out_len);
char* QSysInfo_prettyProductName(long long* __out_len);
char* QSysInfo_machineHostName(long long* __out_len);
void* QSysInfo_machineUniqueId();
void* QSysInfo_bootUniqueId();
void QSysInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
