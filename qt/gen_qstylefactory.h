#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

char** QStyleFactory_keys(long long* __out_len);
void* QStyleFactory_create(const char* arg1, long long arg1_len);
void QStyleFactory_delete(void* self);

#ifdef __cplusplus
}
#endif
