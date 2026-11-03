#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QContiguousCacheData_allocateData(long long size, long long alignment);
void QContiguousCacheData_freeData(void* data);
long long QContiguousCacheData_alloc_get(void* self);
void QContiguousCacheData_alloc_set(void* self, long long alloc);
long long QContiguousCacheData_count_get(void* self);
void QContiguousCacheData_count_set(void* self, long long count);
long long QContiguousCacheData_start_get(void* self);
void QContiguousCacheData_start_set(void* self, long long start);
long long QContiguousCacheData_offset_get(void* self);
void QContiguousCacheData_offset_set(void* self, long long offset);
void QContiguousCacheData_delete(void* self);

#ifdef __cplusplus
}
#endif
