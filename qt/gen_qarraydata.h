#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

long long QArrayData_allocatedCapacity(void* self);
long long QArrayData_constAllocatedCapacity(void* self);
int QArrayData_ref(void* self);
int QArrayData_deref(void* self);
int QArrayData_isShared(void* self);
int QArrayData_needsDetach(void* self);
long long QArrayData_detachCapacity(void* self, long long newSize);
void QArrayData_deallocate(void* data, long long objectSize, long long alignment);
int QArrayData_flags_get(void* self);
void QArrayData_flags_set(void* self, int flags);
long long QArrayData_alloc_get(void* self);
void QArrayData_alloc_set(void* self, long long alloc);
void QArrayData_delete(void* self);

#ifdef __cplusplus
}
#endif
