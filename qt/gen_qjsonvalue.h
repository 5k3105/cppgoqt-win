#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QJsonValue_new();
void* QJsonValue_new2(int b);
void* QJsonValue_new3(double n);
void* QJsonValue_new4(int n);
void* QJsonValue_new5(long long v);
void* QJsonValue_new6(const char* s, long long s_len);
void* QJsonValue_new7(void* s);
void* QJsonValue_new8(const char* s);
void* QJsonValue_new9(void* a);
void* QJsonValue_new10(void* o);
void* QJsonValue_new11(void* other);
void* QJsonValue_new12(int arg1);
void* QJsonValue_OperatorAssign(void* self, void* other);
void QJsonValue_swap(void* self, void* other);
void* QJsonValue_fromVariant(void* variant);
void* QJsonValue_toVariant(void* self);
void* QJsonValue_fromJson(void* json);
void* QJsonValue_toJson(void* self);
int QJsonValue_type(void* self);
int QJsonValue_isNull(void* self);
int QJsonValue_isBool(void* self);
int QJsonValue_isDouble(void* self);
int QJsonValue_isString(void* self);
int QJsonValue_isArray(void* self);
int QJsonValue_isObject(void* self);
int QJsonValue_isUndefined(void* self);
int QJsonValue_toBool(void* self);
int QJsonValue_toInt(void* self);
long long QJsonValue_toInteger(void* self);
double QJsonValue_toDouble(void* self);
char* QJsonValue_toString(void* self, long long* __out_len);
char* QJsonValue_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len);
void* QJsonValue_toStringView(void* self);
void* QJsonValue_toArray(void* self);
void* QJsonValue_toArray2(void* self, void* defaultValue);
void* QJsonValue_toObject(void* self);
void* QJsonValue_toObject2(void* self, void* defaultValue);
void* QJsonValue_OperatorSubscript(void* self, const char* key, long long key_len);
void* QJsonValue_OperatorSubscript2(void* self, void* key);
void* QJsonValue_OperatorSubscript3(void* self, void* key);
void* QJsonValue_OperatorSubscript4(void* self, long long i);
void* QJsonValue_fromJson2(void* json, void* error_);
void* QJsonValue_toJson2(void* self, int format);
int QJsonValue_toBool2(void* self, int defaultValue);
int QJsonValue_toInt2(void* self, int defaultValue);
long long QJsonValue_toInteger2(void* self, long long defaultValue);
double QJsonValue_toDouble2(void* self, double defaultValue);
void* QJsonValue_toStringView2(void* self, const char* defaultValue, long long defaultValue_len);
int QJsonValue_OperatorEqual(void* self, void* rhs);
int QJsonValue_OperatorNotEqual(void* self, void* rhs);
int QJsonValue_OperatorEqual2(void* self, void* rhs);
int QJsonValue_OperatorNotEqual2(void* self, void* rhs);
int QJsonValue_OperatorEqual3(void* self, void* rhs);
int QJsonValue_OperatorNotEqual3(void* self, void* rhs);
int QJsonValue_OperatorEqual4(void* self, void* rhs);
int QJsonValue_OperatorNotEqual4(void* self, void* rhs);
void QJsonValue_delete(void* self);

#ifdef __cplusplus
}
#endif
