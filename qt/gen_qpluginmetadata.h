#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

unsigned char QPluginMetaData_archRequirements();
unsigned long long QPluginMetaData_size_get(void* self);
void QPluginMetaData_size_set(void* self, unsigned long long size);
void QPluginMetaData_delete(void* self);

#ifdef __cplusplus
}
#endif
