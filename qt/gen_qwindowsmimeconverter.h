#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QWindowsMimeConverter_registerMimeType(const char* mimeType, long long mimeType_len);
void QWindowsMimeConverter_delete(void* self);

#ifdef __cplusplus
}
#endif
