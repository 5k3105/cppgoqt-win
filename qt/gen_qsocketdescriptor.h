#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSocketDescriptor_new();
void* QSocketDescriptor_new2(long long desc);
long long QSocketDescriptor_ToQintptr(void* self);
int QSocketDescriptor_isValid(void* self);
int QSocketDescriptor_OperatorEqual(void* self, void* rhs);
int QSocketDescriptor_OperatorNotEqual(void* self, void* rhs);
void QSocketDescriptor_delete(void* self);

#ifdef __cplusplus
}
#endif
