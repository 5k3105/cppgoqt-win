#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QJsonObject_new();
void* QJsonObject_new2(void* other);
void* QJsonObject_OperatorAssign(void* self, void* other);
void QJsonObject_swap(void* self, void* other);
void* QJsonObject_fromVariantMap(char** map__keys, void* map__vals, long long map__len);
void* QJsonObject_fromVariantHash(char** map__keys, void* map__vals, long long map__len);
char** QJsonObject_keys(void* self, long long* __out_len);
long long QJsonObject_size(void* self);
long long QJsonObject_count(void* self);
long long QJsonObject_length(void* self);
int QJsonObject_isEmpty(void* self);
void* QJsonObject_value(void* self, const char* key, long long key_len);
void* QJsonObject_OperatorSubscript(void* self, const char* key, long long key_len);
void* QJsonObject_OperatorSubscript2(void* self, const char* key, long long key_len);
void* QJsonObject_value2(void* self, void* key);
void* QJsonObject_value3(void* self, void* key);
void* QJsonObject_OperatorSubscript3(void* self, void* key);
void* QJsonObject_OperatorSubscript4(void* self, void* key);
void* QJsonObject_OperatorSubscript5(void* self, void* key);
void* QJsonObject_OperatorSubscript6(void* self, void* key);
void QJsonObject_remove(void* self, const char* key, long long key_len);
void* QJsonObject_take(void* self, const char* key, long long key_len);
int QJsonObject_contains(void* self, const char* key, long long key_len);
void QJsonObject_remove2(void* self, void* key);
void QJsonObject_remove3(void* self, void* key);
void* QJsonObject_take2(void* self, void* key);
void* QJsonObject_take3(void* self, void* key);
int QJsonObject_contains2(void* self, void* key);
int QJsonObject_contains3(void* self, void* key);
int QJsonObject_empty(void* self);
int QJsonObject_OperatorEqual(void* self, void* rhs);
int QJsonObject_OperatorNotEqual(void* self, void* rhs);
int QJsonObject_OperatorEqual2(void* self, void* rhs);
int QJsonObject_OperatorNotEqual2(void* self, void* rhs);
int QJsonObject_OperatorEqual3(void* self, void* rhs);
int QJsonObject_OperatorNotEqual3(void* self, void* rhs);
void QJsonObject_delete(void* self);

#ifdef __cplusplus
}
#endif
