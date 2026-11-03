#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGenericArgument_new();
void* QGenericArgument_new2(void* arg1);
void* QGenericArgument_new3(const char* aName);
const char* QGenericArgument_name(void* self);
void QGenericArgument_delete(void* self);

#ifdef __cplusplus
}
#endif
