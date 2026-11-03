#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaClassInfo_new();
void* QMetaClassInfo_new2(void* arg1);
const char* QMetaClassInfo_name(void* self);
const char* QMetaClassInfo_value(void* self);
void* QMetaClassInfo_enclosingMetaObject(void* self);
void QMetaClassInfo_delete(void* self);

#ifdef __cplusplus
}
#endif
