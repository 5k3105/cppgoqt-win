#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaEnum_new();
void* QMetaEnum_new2(void* arg1);
const char* QMetaEnum_name(void* self);
const char* QMetaEnum_enumName(void* self);
void* QMetaEnum_metaType(void* self);
int QMetaEnum_isFlag(void* self);
int QMetaEnum_isScoped(void* self);
int QMetaEnum_is64Bit(void* self);
int QMetaEnum_keyCount(void* self);
const char* QMetaEnum_key(void* self, int index);
int QMetaEnum_value(void* self, int index);
unsigned long long QMetaEnum_value64(void* self, int index, int* __ok);
const char* QMetaEnum_scope(void* self);
int QMetaEnum_keyToValue(void* self, const char* key);
int QMetaEnum_keysToValue(void* self, const char* keys);
unsigned long long QMetaEnum_keyToValue64(void* self, const char* key, int* __ok);
unsigned long long QMetaEnum_keysToValue64(void* self, const char* keys, int* __ok);
const char* QMetaEnum_valueToKey(void* self, unsigned long long value);
void* QMetaEnum_valueToKeys(void* self, unsigned long long value);
void* QMetaEnum_enclosingMetaObject(void* self);
int QMetaEnum_isValid(void* self);
int QMetaEnum_keyToValue2(void* self, const char* key, int* ok);
int QMetaEnum_keysToValue2(void* self, const char* keys, int* ok);
void QMetaEnum_delete(void* self);

#ifdef __cplusplus
}
#endif
