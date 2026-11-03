#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QOperatingSystemVersion_base(void* self);
void QOperatingSystemVersion_chain(void* self, void** out);
void* QOperatingSystemVersion_new(void* osversion);
void* QOperatingSystemVersion_new2(int osType, int vmajor);
void* QOperatingSystemVersion_new3(void* arg1);
void* QOperatingSystemVersion_new4(int osType, int vmajor, int vminor);
void* QOperatingSystemVersion_new5(int osType, int vmajor, int vminor, int vmicro);
void* QOperatingSystemVersion_current();
int QOperatingSystemVersion_currentType();
int QOperatingSystemVersion_type(void* self);
char* QOperatingSystemVersion_name(void* osversion, long long* __out_len);
void QOperatingSystemVersion_delete(void* self);

#ifdef __cplusplus
}
#endif
