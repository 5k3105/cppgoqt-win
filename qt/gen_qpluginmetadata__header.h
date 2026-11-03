#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPluginMetaData__Header_new(void* arg1);
unsigned char QPluginMetaData__Header_version_get(void* self);
void QPluginMetaData__Header_version_set(void* self, unsigned char version);
unsigned char QPluginMetaData__Header_qt_major_version_get(void* self);
void QPluginMetaData__Header_qt_major_version_set(void* self, unsigned char qt_major_version);
unsigned char QPluginMetaData__Header_qt_minor_version_get(void* self);
void QPluginMetaData__Header_qt_minor_version_set(void* self, unsigned char qt_minor_version);
unsigned char QPluginMetaData__Header_plugin_arch_requirements_get(void* self);
void QPluginMetaData__Header_plugin_arch_requirements_set(void* self, unsigned char plugin_arch_requirements);
void QPluginMetaData__Header_delete(void* self);

#ifdef __cplusplus
}
#endif
