#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUuid_new();
void* QUuid_new2(unsigned int l, unsigned short w1, unsigned short w2, unsigned char b1, unsigned char b2, unsigned char b3, unsigned char b4, unsigned char b5, unsigned char b6, unsigned char b7, unsigned char b8);
void* QUuid_new3(void* id128);
void* QUuid_new4(const char* string_, long long string__len);
void* QUuid_new5(void* arg1);
void* QUuid_new6(void* id128, int order);
void* QUuid_fromString(const char* string_, long long string__len);
char* QUuid_toString(void* self, long long* __out_len);
void* QUuid_toByteArray(void* self);
void* QUuid_toBytes(void* self);
void* QUuid_toRfc4122(void* self);
void* QUuid_fromRfc4122(void* arg1);
int QUuid_isNull(void* self);
void* QUuid_ToGUID(void* self);
void* QUuid_createUuid();
void* QUuid_createUuidV5(void* ns, void* baseData);
void* QUuid_createUuidV3(void* ns, void* baseData);
void* QUuid_createUuidV7();
int QUuid_variant(void* self);
int QUuid_version(void* self);
void* QUuid_OperatorAssign(void* self, void* arg1);
char* QUuid_toString2(void* self, int mode, long long* __out_len);
void* QUuid_toByteArray2(void* self, int mode);
void* QUuid_toBytes2(void* self, int order);
int QUuid_OperatorEqual(void* self, void* rhs);
int QUuid_OperatorNotEqual(void* self, void* rhs);
int QUuid_OperatorLesser(void* self, void* rhs);
int QUuid_OperatorGreater(void* self, void* rhs);
int QUuid_OperatorLesserOrEqual(void* self, void* rhs);
int QUuid_OperatorGreaterOrEqual(void* self, void* rhs);
unsigned int QUuid_data1_get(void* self);
void QUuid_data1_set(void* self, unsigned int data1);
unsigned short QUuid_data2_get(void* self);
void QUuid_data2_set(void* self, unsigned short data2);
unsigned short QUuid_data3_get(void* self);
void QUuid_data3_set(void* self, unsigned short data3);
void QUuid_delete(void* self);

#ifdef __cplusplus
}
#endif
