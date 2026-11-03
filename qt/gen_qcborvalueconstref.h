#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCborValueConstRef_new(void* arg1);
void* QCborValueConstRef_ToQCborValue(void* self);
int QCborValueConstRef_type(void* self);
int QCborValueConstRef_isInteger(void* self);
int QCborValueConstRef_isByteArray(void* self);
int QCborValueConstRef_isString(void* self);
int QCborValueConstRef_isArray(void* self);
int QCborValueConstRef_isMap(void* self);
int QCborValueConstRef_isTag(void* self);
int QCborValueConstRef_isFalse(void* self);
int QCborValueConstRef_isTrue(void* self);
int QCborValueConstRef_isBool(void* self);
int QCborValueConstRef_isNull(void* self);
int QCborValueConstRef_isUndefined(void* self);
int QCborValueConstRef_isDouble(void* self);
int QCborValueConstRef_isDateTime(void* self);
int QCborValueConstRef_isUrl(void* self);
int QCborValueConstRef_isRegularExpression(void* self);
int QCborValueConstRef_isUuid(void* self);
int QCborValueConstRef_isInvalid(void* self);
int QCborValueConstRef_isContainer(void* self);
int QCborValueConstRef_isSimpleType(void* self);
int QCborValueConstRef_isSimpleType2(void* self, int st);
int QCborValueConstRef_toSimpleType(void* self);
int QCborValueConstRef_tag(void* self);
void* QCborValueConstRef_taggedValue(void* self);
long long QCborValueConstRef_toInteger(void* self);
int QCborValueConstRef_toBool(void* self);
double QCborValueConstRef_toDouble(void* self);
void* QCborValueConstRef_toByteArray(void* self);
char* QCborValueConstRef_toString(void* self, long long* __out_len);
void* QCborValueConstRef_toStringView(void* self);
void* QCborValueConstRef_toDateTime(void* self);
void* QCborValueConstRef_toUrl(void* self);
void* QCborValueConstRef_toRegularExpression(void* self);
void* QCborValueConstRef_toUuid(void* self);
void* QCborValueConstRef_toArray(void* self);
void* QCborValueConstRef_toArray2(void* self, void* a);
void* QCborValueConstRef_toMap(void* self);
void* QCborValueConstRef_toMap2(void* self, void* m);
void* QCborValueConstRef_OperatorSubscript(void* self, const char* key, long long key_len);
void* QCborValueConstRef_OperatorSubscript2(void* self, void* key);
void* QCborValueConstRef_OperatorSubscript3(void* self, long long key);
int QCborValueConstRef_compare(void* self, void* other);
void* QCborValueConstRef_toVariant(void* self);
void* QCborValueConstRef_toJsonValue(void* self);
void* QCborValueConstRef_toCbor(void* self);
void QCborValueConstRef_toCbor2(void* self, void* writer);
char* QCborValueConstRef_toDiagnosticNotation(void* self, long long* __out_len);
int QCborValueConstRef_toSimpleType2(void* self, int defaultValue);
int QCborValueConstRef_tag2(void* self, int defaultValue);
void* QCborValueConstRef_taggedValue2(void* self, void* defaultValue);
long long QCborValueConstRef_toInteger2(void* self, long long defaultValue);
int QCborValueConstRef_toBool2(void* self, int defaultValue);
double QCborValueConstRef_toDouble2(void* self, double defaultValue);
void* QCborValueConstRef_toByteArray2(void* self, void* defaultValue);
char* QCborValueConstRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len);
void* QCborValueConstRef_toStringView2(void* self, const char* defaultValue, long long defaultValue_len);
void* QCborValueConstRef_toDateTime2(void* self, void* defaultValue);
void* QCborValueConstRef_toUrl2(void* self, void* defaultValue);
void* QCborValueConstRef_toRegularExpression2(void* self, void* defaultValue);
void* QCborValueConstRef_toUuid2(void* self, void* defaultValue);
void* QCborValueConstRef_toCbor3(void* self, int opt);
void QCborValueConstRef_toCbor4(void* self, void* writer, int opt);
char* QCborValueConstRef_toDiagnosticNotation2(void* self, int opt, long long* __out_len);
int QCborValueConstRef_OperatorEqual(void* self, void* rhs);
int QCborValueConstRef_OperatorNotEqual(void* self, void* rhs);
int QCborValueConstRef_OperatorLesser(void* self, void* rhs);
int QCborValueConstRef_OperatorGreater(void* self, void* rhs);
int QCborValueConstRef_OperatorLesserOrEqual(void* self, void* rhs);
int QCborValueConstRef_OperatorGreaterOrEqual(void* self, void* rhs);
int QCborValueConstRef_OperatorEqual2(void* self, void* rhs);
int QCborValueConstRef_OperatorNotEqual2(void* self, void* rhs);
int QCborValueConstRef_OperatorLesser2(void* self, void* rhs);
int QCborValueConstRef_OperatorGreater2(void* self, void* rhs);
int QCborValueConstRef_OperatorLesserOrEqual2(void* self, void* rhs);
int QCborValueConstRef_OperatorGreaterOrEqual2(void* self, void* rhs);
int QCborValueConstRef_OperatorEqual3(void* self, void* rhs);
int QCborValueConstRef_OperatorNotEqual3(void* self, void* rhs);
int QCborValueConstRef_OperatorLesser3(void* self, void* rhs);
int QCborValueConstRef_OperatorGreater3(void* self, void* rhs);
int QCborValueConstRef_OperatorLesserOrEqual3(void* self, void* rhs);
int QCborValueConstRef_OperatorGreaterOrEqual3(void* self, void* rhs);
int QCborValueConstRef_OperatorEqual4(void* self, void* rhs);
int QCborValueConstRef_OperatorNotEqual4(void* self, void* rhs);
int QCborValueConstRef_OperatorLesser4(void* self, void* rhs);
int QCborValueConstRef_OperatorGreater4(void* self, void* rhs);
int QCborValueConstRef_OperatorLesserOrEqual4(void* self, void* rhs);
int QCborValueConstRef_OperatorGreaterOrEqual4(void* self, void* rhs);
void QCborValueConstRef_delete(void* self);

#ifdef __cplusplus
}
#endif
