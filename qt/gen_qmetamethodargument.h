#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

const char* QMetaMethodArgument_name_get(void* self);
void QMetaMethodArgument_name_set(void* self, const char* name);
void QMetaMethodArgument_delete(void* self);

#ifdef __cplusplus
}
#endif
