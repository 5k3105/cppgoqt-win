#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStaticPlugin_new(void* arg1);
void* QStaticPlugin_metaData(void* self);
void QStaticPlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
