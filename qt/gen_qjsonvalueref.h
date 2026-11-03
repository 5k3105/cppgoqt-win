#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QJsonValueRef_base(void* self);
void* QJsonValueRef_new(void* arg1);
void* QJsonValueRef_new2(void* array, long long idx);
void* QJsonValueRef_new3(void* object, long long idx);
void* QJsonValueRef_OperatorAssign(void* self, void* val);
void* QJsonValueRef_OperatorAssign2(void* self, void* val);
void* QJsonValueRef_ToQJsonValue(void* self);
void* QJsonValueRef_toVariant(void* self);
int QJsonValueRef_type(void* self);
int QJsonValueRef_isNull(void* self);
int QJsonValueRef_isBool(void* self);
int QJsonValueRef_isDouble(void* self);
int QJsonValueRef_isString(void* self);
int QJsonValueRef_isArray(void* self);
int QJsonValueRef_isObject(void* self);
int QJsonValueRef_isUndefined(void* self);
int QJsonValueRef_toBool(void* self);
int QJsonValueRef_toInt(void* self);
long long QJsonValueRef_toInteger(void* self);
double QJsonValueRef_toDouble(void* self);
char* QJsonValueRef_toString(void* self, long long* __out_len);
void* QJsonValueRef_toStringView(void* self);
void* QJsonValueRef_toArray(void* self);
void* QJsonValueRef_toObject(void* self);
void* QJsonValueRef_OperatorSubscript(void* self, void* key);
void* QJsonValueRef_OperatorSubscript2(void* self, void* key);
void* QJsonValueRef_OperatorSubscript3(void* self, long long i);
int QJsonValueRef_toBool2(void* self, int defaultValue);
int QJsonValueRef_toInt2(void* self, int defaultValue);
long long QJsonValueRef_toInteger2(void* self, long long defaultValue);
double QJsonValueRef_toDouble2(void* self, double defaultValue);
char* QJsonValueRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len);
void* QJsonValueRef_toStringView2(void* self, const char* defaultValue, long long defaultValue_len);
int QJsonValueRef_OperatorEqual(void* self, void* rhs);
int QJsonValueRef_OperatorNotEqual(void* self, void* rhs);
int QJsonValueRef_OperatorEqual2(void* self, void* rhs);
int QJsonValueRef_OperatorNotEqual2(void* self, void* rhs);
void QJsonValueRef_delete(void* self);

#ifdef __cplusplus
}
#endif
