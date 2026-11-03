#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QJsonDocument_new();
void* QJsonDocument_new2(void* object);
void* QJsonDocument_new3(void* array);
void* QJsonDocument_new4(void* other);
void* QJsonDocument_OperatorAssign(void* self, void* other);
void QJsonDocument_swap(void* self, void* other);
void* QJsonDocument_fromVariant(void* variant);
void* QJsonDocument_toVariant(void* self);
void* QJsonDocument_fromJson(void* json);
void* QJsonDocument_toJson(void* self);
int QJsonDocument_isEmpty(void* self);
int QJsonDocument_isArray(void* self);
int QJsonDocument_isObject(void* self);
void* QJsonDocument_object(void* self);
void* QJsonDocument_array(void* self);
void QJsonDocument_setObject(void* self, void* object);
void QJsonDocument_setArray(void* self, void* array);
void* QJsonDocument_OperatorSubscript(void* self, const char* key, long long key_len);
void* QJsonDocument_OperatorSubscript2(void* self, void* key);
void* QJsonDocument_OperatorSubscript3(void* self, void* key);
void* QJsonDocument_OperatorSubscript4(void* self, long long i);
int QJsonDocument_isNull(void* self);
void* QJsonDocument_fromJson2(void* json, void* error_);
void* QJsonDocument_toJson2(void* self, int format);
int QJsonDocument_OperatorEqual(void* self, void* rhs);
int QJsonDocument_OperatorNotEqual(void* self, void* rhs);
void QJsonDocument_delete(void* self);

#ifdef __cplusplus
}
#endif
