#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUuid__Id128Bytes_new();
void* QUuid__Id128Bytes_new2(void* arg1);
void* QUuid__Id128Bytes_ToQByteArrayView(void* self);
void QUuid__Id128Bytes_delete(void* self);

#ifdef __cplusplus
}
#endif
