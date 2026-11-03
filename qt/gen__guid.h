#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* _GUID_new();
void* _GUID_new2(void* arg1);
unsigned long _GUID_Data1_get(void* self);
void _GUID_Data1_set(void* self, unsigned long Data1);
unsigned short _GUID_Data2_get(void* self);
void _GUID_Data2_set(void* self, unsigned short Data2);
unsigned short _GUID_Data3_get(void* self);
void _GUID_Data3_set(void* self, unsigned short Data3);
void _GUID_delete(void* self);

#ifdef __cplusplus
}
#endif
