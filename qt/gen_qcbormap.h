#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCborMap_new();
void* QCborMap_new2(void* other);
void* QCborMap_OperatorAssign(void* self, void* other);
void QCborMap_swap(void* self, void* other);
void* QCborMap_toCborValue(void* self);
long long QCborMap_size(void* self);
int QCborMap_isEmpty(void* self);
void QCborMap_clear(void* self);
void** QCborMap_keys(void* self, long long* __out_len);
void* QCborMap_value(void* self, long long key);
void* QCborMap_value2(void* self, void* key);
void* QCborMap_value3(void* self, const char* key, long long key_len);
void* QCborMap_value4(void* self, void* key);
void* QCborMap_OperatorSubscript(void* self, long long key);
void* QCborMap_OperatorSubscript2(void* self, void* key);
void* QCborMap_OperatorSubscript3(void* self, const char* key, long long key_len);
void* QCborMap_OperatorSubscript4(void* self, void* key);
void* QCborMap_OperatorSubscript5(void* self, long long key);
void* QCborMap_OperatorSubscript6(void* self, void* key);
void* QCborMap_OperatorSubscript7(void* self, const char* key, long long key_len);
void* QCborMap_OperatorSubscript8(void* self, void* key);
void* QCborMap_take(void* self, long long key);
void* QCborMap_take2(void* self, void* key);
void* QCborMap_take3(void* self, const char* key, long long key_len);
void* QCborMap_take4(void* self, void* key);
void QCborMap_remove(void* self, long long key);
void QCborMap_remove2(void* self, void* key);
void QCborMap_remove3(void* self, const char* key, long long key_len);
void QCborMap_remove4(void* self, void* key);
int QCborMap_contains(void* self, long long key);
int QCborMap_contains2(void* self, void* key);
int QCborMap_contains3(void* self, const char* key, long long key_len);
int QCborMap_contains4(void* self, void* key);
int QCborMap_compare(void* self, void* other);
int QCborMap_empty(void* self);
void* QCborMap_fromVariantMap(char** map__keys, void* map__vals, long long map__len);
void* QCborMap_fromVariantHash(char** hash_keys, void* hash_vals, long long hash_len);
void* QCborMap_fromJsonObject(void* o);
void* QCborMap_toJsonObject(void* self);
int QCborMap_OperatorEqual(void* self, void* rhs);
int QCborMap_OperatorNotEqual(void* self, void* rhs);
int QCborMap_OperatorLesser(void* self, void* rhs);
int QCborMap_OperatorGreater(void* self, void* rhs);
int QCborMap_OperatorLesserOrEqual(void* self, void* rhs);
int QCborMap_OperatorGreaterOrEqual(void* self, void* rhs);
int QCborMap_OperatorEqual2(void* self, void* rhs);
int QCborMap_OperatorNotEqual2(void* self, void* rhs);
int QCborMap_OperatorLesser2(void* self, void* rhs);
int QCborMap_OperatorGreater2(void* self, void* rhs);
int QCborMap_OperatorLesserOrEqual2(void* self, void* rhs);
int QCborMap_OperatorGreaterOrEqual2(void* self, void* rhs);
int QCborMap_OperatorEqual3(void* self, void* rhs);
int QCborMap_OperatorNotEqual3(void* self, void* rhs);
int QCborMap_OperatorLesser3(void* self, void* rhs);
int QCborMap_OperatorGreater3(void* self, void* rhs);
int QCborMap_OperatorLesserOrEqual3(void* self, void* rhs);
int QCborMap_OperatorGreaterOrEqual3(void* self, void* rhs);
void QCborMap_delete(void* self);

#ifdef __cplusplus
}
#endif
