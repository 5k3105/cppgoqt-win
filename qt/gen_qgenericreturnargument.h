#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGenericReturnArgument_base(void* self);
void* QGenericReturnArgument_new();
void* QGenericReturnArgument_new2(const char* aName);
void QGenericReturnArgument_delete(void* self);

#ifdef __cplusplus
}
#endif
