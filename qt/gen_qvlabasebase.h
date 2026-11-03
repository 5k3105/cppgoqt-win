#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

long long QVLABaseBase_capacity(void* self);
long long QVLABaseBase_size(void* self);
int QVLABaseBase_empty(void* self);
void QVLABaseBase_delete(void* self);

#ifdef __cplusplus
}
#endif
