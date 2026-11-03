#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QOperatingSystemVersionBase_new(int osType, int vmajor);
void* QOperatingSystemVersionBase_new2(void* arg1);
void* QOperatingSystemVersionBase_new3(int osType, int vmajor, int vminor);
void* QOperatingSystemVersionBase_new4(int osType, int vmajor, int vminor, int vmicro);
void* QOperatingSystemVersionBase_current();
char* QOperatingSystemVersionBase_name(void* osversion, long long* __out_len);
int QOperatingSystemVersionBase_currentType();
void* QOperatingSystemVersionBase_version(void* self);
int QOperatingSystemVersionBase_majorVersion(void* self);
int QOperatingSystemVersionBase_minorVersion(void* self);
int QOperatingSystemVersionBase_microVersion(void* self);
int QOperatingSystemVersionBase_segmentCount(void* self);
int QOperatingSystemVersionBase_type(void* self);
char* QOperatingSystemVersionBase_name2(void* self, long long* __out_len);
int QOperatingSystemVersionBase_OperatorGreater(void* self, void* rhs);
int QOperatingSystemVersionBase_OperatorGreaterOrEqual(void* self, void* rhs);
int QOperatingSystemVersionBase_OperatorLesser(void* self, void* rhs);
int QOperatingSystemVersionBase_OperatorLesserOrEqual(void* self, void* rhs);
void QOperatingSystemVersionBase_delete(void* self);

#ifdef __cplusplus
}
#endif
