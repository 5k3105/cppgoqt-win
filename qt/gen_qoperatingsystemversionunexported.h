#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QOperatingSystemVersionUnexported_base(void* self);
void* QOperatingSystemVersionUnexported_new(void* other);
void* QOperatingSystemVersionUnexported_new2();
void* QOperatingSystemVersionUnexported_new3(void* arg1);
void* QOperatingSystemVersionUnexported_new4(int arg1, int arg2, int arg3, int arg4);
void* QOperatingSystemVersionUnexported_current();
char* QOperatingSystemVersionUnexported_name(void* osversion, long long* __out_len);
int QOperatingSystemVersionUnexported_currentType();
void QOperatingSystemVersionUnexported_delete(void* self);

#ifdef __cplusplus
}
#endif
