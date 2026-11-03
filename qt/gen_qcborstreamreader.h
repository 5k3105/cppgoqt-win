#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCborStreamReader_new();
void* QCborStreamReader_new2(const char* data, long long len_);
void* QCborStreamReader_new3(const unsigned char* data, long long len_);
void* QCborStreamReader_new4(void* data);
void* QCborStreamReader_new5(void* device);
void QCborStreamReader_setDevice(void* self, void* device);
void* QCborStreamReader_device(void* self);
void QCborStreamReader_addData(void* self, void* data);
void QCborStreamReader_addData2(void* self, const char* data, long long len_);
void QCborStreamReader_addData3(void* self, const unsigned char* data, long long len_);
void QCborStreamReader_reparse(void* self);
void QCborStreamReader_clear(void* self);
void QCborStreamReader_reset(void* self);
void* QCborStreamReader_lastError(void* self);
long long QCborStreamReader_currentOffset(void* self);
int QCborStreamReader_isValid(void* self);
int QCborStreamReader_containerDepth(void* self);
int QCborStreamReader_parentContainerType(void* self);
int QCborStreamReader_hasNext(void* self);
int QCborStreamReader_next(void* self);
int QCborStreamReader_type(void* self);
int QCborStreamReader_isUnsignedInteger(void* self);
int QCborStreamReader_isNegativeInteger(void* self);
int QCborStreamReader_isInteger(void* self);
int QCborStreamReader_isByteArray(void* self);
int QCborStreamReader_isString(void* self);
int QCborStreamReader_isArray(void* self);
int QCborStreamReader_isMap(void* self);
int QCborStreamReader_isTag(void* self);
int QCborStreamReader_isSimpleType(void* self);
int QCborStreamReader_isFloat16(void* self);
int QCborStreamReader_isFloat(void* self);
int QCborStreamReader_isDouble(void* self);
int QCborStreamReader_isInvalid(void* self);
int QCborStreamReader_isSimpleType2(void* self, int st);
int QCborStreamReader_isFalse(void* self);
int QCborStreamReader_isTrue(void* self);
int QCborStreamReader_isBool(void* self);
int QCborStreamReader_isNull(void* self);
int QCborStreamReader_isUndefined(void* self);
int QCborStreamReader_isLengthKnown(void* self);
unsigned long long QCborStreamReader_length(void* self);
int QCborStreamReader_isContainer(void* self);
int QCborStreamReader_enterContainer(void* self);
int QCborStreamReader_leaveContainer(void* self);
int QCborStreamReader_readAndAppendToString(void* self, void* dst);
int QCborStreamReader_readAndAppendToUtf8String(void* self, void* dst);
int QCborStreamReader_readAndAppendToByteArray(void* self, void* dst);
long long QCborStreamReader_currentStringChunkSize(void* self);
int QCborStreamReader_toBool(void* self);
int QCborStreamReader_toTag(void* self);
unsigned long long QCborStreamReader_toUnsignedInteger(void* self);
int QCborStreamReader_toNegativeInteger(void* self);
int QCborStreamReader_toSimpleType(void* self);
void* QCborStreamReader_toFloat16(void* self);
float QCborStreamReader_toFloat(void* self);
double QCborStreamReader_toDouble(void* self);
long long QCborStreamReader_toInteger(void* self);
char* QCborStreamReader_readAllString(void* self, long long* __out_len);
void* QCborStreamReader_readAllUtf8String(void* self);
void* QCborStreamReader_readAllByteArray(void* self);
int QCborStreamReader_next2(void* self, int maxRecursion);
void QCborStreamReader_delete(void* self);

#ifdef __cplusplus
}
#endif
