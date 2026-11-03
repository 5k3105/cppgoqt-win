#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPluginMetaData__MagicHeader_new();
void* QPluginMetaData__MagicHeader_header_get(void* self);
void QPluginMetaData__MagicHeader_header_set(void* self, void* header);
void QPluginMetaData__MagicHeader_delete(void* self);

#ifdef __cplusplus
}
#endif
