#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QCborError_ToQCborErrorCode(void* self);
char* QCborError_toString(void* self, long long* __out_len);
int QCborError_c_get(void* self);
void QCborError_c_set(void* self, int c);
void QCborError_delete(void* self);

#ifdef __cplusplus
}
#endif
