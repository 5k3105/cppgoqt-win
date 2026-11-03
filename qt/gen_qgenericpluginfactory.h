#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

char** QGenericPluginFactory_keys(long long* __out_len);
void* QGenericPluginFactory_create(const char* arg1, long long arg1_len, const char* arg2, long long arg2_len);
void QGenericPluginFactory_delete(void* self);

#ifdef __cplusplus
}
#endif
