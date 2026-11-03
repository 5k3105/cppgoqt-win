#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QJsonValueConstRef_new(void* arg1);
void* QJsonValueConstRef_ToQJsonValue(void* self);
void* QJsonValueConstRef_toVariant(void* self);
int QJsonValueConstRef_type(void* self);
int QJsonValueConstRef_isNull(void* self);
int QJsonValueConstRef_isBool(void* self);
int QJsonValueConstRef_isDouble(void* self);
int QJsonValueConstRef_isString(void* self);
int QJsonValueConstRef_isArray(void* self);
int QJsonValueConstRef_isObject(void* self);
int QJsonValueConstRef_isUndefined(void* self);
int QJsonValueConstRef_toBool(void* self);
int QJsonValueConstRef_toInt(void* self);
long long QJsonValueConstRef_toInteger(void* self);
double QJsonValueConstRef_toDouble(void* self);
char* QJsonValueConstRef_toString(void* self, long long* __out_len);
void* QJsonValueConstRef_toStringView(void* self);
void* QJsonValueConstRef_toArray(void* self);
void* QJsonValueConstRef_toObject(void* self);
void* QJsonValueConstRef_OperatorSubscript(void* self, void* key);
void* QJsonValueConstRef_OperatorSubscript2(void* self, void* key);
void* QJsonValueConstRef_OperatorSubscript3(void* self, long long i);
int QJsonValueConstRef_toBool2(void* self, int defaultValue);
int QJsonValueConstRef_toInt2(void* self, int defaultValue);
long long QJsonValueConstRef_toInteger2(void* self, long long defaultValue);
double QJsonValueConstRef_toDouble2(void* self, double defaultValue);
char* QJsonValueConstRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len);
void* QJsonValueConstRef_toStringView2(void* self, const char* defaultValue, long long defaultValue_len);
int QJsonValueConstRef_OperatorEqual(void* self, void* rhs);
int QJsonValueConstRef_OperatorNotEqual(void* self, void* rhs);
int QJsonValueConstRef_OperatorEqual2(void* self, void* rhs);
int QJsonValueConstRef_OperatorNotEqual2(void* self, void* rhs);
int QJsonValueConstRef_OperatorEqual3(void* self, void* rhs);
int QJsonValueConstRef_OperatorNotEqual3(void* self, void* rhs);
int QJsonValueConstRef_OperatorEqual4(void* self, void* rhs);
int QJsonValueConstRef_OperatorNotEqual4(void* self, void* rhs);
void QJsonValueConstRef_delete(void* self);

#ifdef __cplusplus
}
#endif
