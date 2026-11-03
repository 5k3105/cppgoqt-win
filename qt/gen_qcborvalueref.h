#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCborValueRef_base(void* self);
void* QCborValueRef_new(void* arg1);
void* QCborValueRef_OperatorAssign(void* self, void* other);
void* QCborValueRef_OperatorAssign2(void* self, void* other);
void* QCborValueRef_OperatorSubscript(void* self, long long key);
void* QCborValueRef_OperatorSubscript2(void* self, void* key);
void* QCborValueRef_OperatorSubscript3(void* self, const char* key, long long key_len);
void* QCborValueRef_ToQCborValue(void* self);
int QCborValueRef_type(void* self);
int QCborValueRef_isInteger(void* self);
int QCborValueRef_isByteArray(void* self);
int QCborValueRef_isString(void* self);
int QCborValueRef_isArray(void* self);
int QCborValueRef_isMap(void* self);
int QCborValueRef_isTag(void* self);
int QCborValueRef_isFalse(void* self);
int QCborValueRef_isTrue(void* self);
int QCborValueRef_isBool(void* self);
int QCborValueRef_isNull(void* self);
int QCborValueRef_isUndefined(void* self);
int QCborValueRef_isDouble(void* self);
int QCborValueRef_isDateTime(void* self);
int QCborValueRef_isUrl(void* self);
int QCborValueRef_isRegularExpression(void* self);
int QCborValueRef_isUuid(void* self);
int QCborValueRef_isInvalid(void* self);
int QCborValueRef_isContainer(void* self);
int QCborValueRef_isSimpleType(void* self);
int QCborValueRef_isSimpleType2(void* self, int st);
int QCborValueRef_toSimpleType(void* self);
int QCborValueRef_tag(void* self);
void* QCborValueRef_taggedValue(void* self);
long long QCborValueRef_toInteger(void* self);
int QCborValueRef_toBool(void* self);
double QCborValueRef_toDouble(void* self);
void* QCborValueRef_toByteArray(void* self);
char* QCborValueRef_toString(void* self, long long* __out_len);
void* QCborValueRef_toDateTime(void* self);
void* QCborValueRef_toUrl(void* self);
void* QCborValueRef_toRegularExpression(void* self);
void* QCborValueRef_toUuid(void* self);
void* QCborValueRef_toArray(void* self);
void* QCborValueRef_toArray2(void* self, void* a);
void* QCborValueRef_toMap(void* self);
void* QCborValueRef_toMap2(void* self, void* m);
void* QCborValueRef_OperatorSubscript4(void* self, const char* key, long long key_len);
void* QCborValueRef_OperatorSubscript5(void* self, void* key);
void* QCborValueRef_OperatorSubscript6(void* self, long long key);
int QCborValueRef_compare(void* self, void* other);
void* QCborValueRef_toVariant(void* self);
void* QCborValueRef_toJsonValue(void* self);
void* QCborValueRef_toCbor(void* self);
void QCborValueRef_toCbor2(void* self, void* writer);
char* QCborValueRef_toDiagnosticNotation(void* self, long long* __out_len);
int QCborValueRef_toSimpleType2(void* self, int defaultValue);
int QCborValueRef_tag2(void* self, int defaultValue);
void* QCborValueRef_taggedValue2(void* self, void* defaultValue);
long long QCborValueRef_toInteger2(void* self, long long defaultValue);
int QCborValueRef_toBool2(void* self, int defaultValue);
double QCborValueRef_toDouble2(void* self, double defaultValue);
void* QCborValueRef_toByteArray2(void* self, void* defaultValue);
char* QCborValueRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len);
void* QCborValueRef_toDateTime2(void* self, void* defaultValue);
void* QCborValueRef_toUrl2(void* self, void* defaultValue);
void* QCborValueRef_toRegularExpression2(void* self, void* defaultValue);
void* QCborValueRef_toUuid2(void* self, void* defaultValue);
void* QCborValueRef_toCbor3(void* self, int opt);
void QCborValueRef_toCbor4(void* self, void* writer, int opt);
char* QCborValueRef_toDiagnosticNotation2(void* self, int opt, long long* __out_len);
void QCborValueRef_delete(void* self);

#ifdef __cplusplus
}
#endif
