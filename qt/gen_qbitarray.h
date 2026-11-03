#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBitArray_new();
void* QBitArray_new2(long long size);
void* QBitArray_new3(void* other);
void* QBitArray_new4(long long size, int val);
void* QBitArray_OperatorAssign(void* self, void* other);
void QBitArray_swap(void* self, void* other);
long long QBitArray_size(void* self);
long long QBitArray_count(void* self);
long long QBitArray_count2(void* self, int on);
int QBitArray_isEmpty(void* self);
int QBitArray_isNull(void* self);
void QBitArray_resize(void* self, long long size);
void QBitArray_detach(void* self);
int QBitArray_isDetached(void* self);
void QBitArray_clear(void* self);
int QBitArray_testBit(void* self, long long i);
void QBitArray_setBit(void* self, long long i);
void QBitArray_setBit2(void* self, long long i, int val);
void QBitArray_clearBit(void* self, long long i);
int QBitArray_toggleBit(void* self, long long i);
int QBitArray_at(void* self, long long i);
void* QBitArray_OperatorSubscript(void* self, long long i);
int QBitArray_OperatorSubscript2(void* self, long long i);
void* QBitArray_OperatorBitwiseAndAssign(void* self, void* arg1);
void* QBitArray_OperatorBitwiseOrAssign(void* self, void* arg1);
void* QBitArray_OperatorBitwiseXorAssign(void* self, void* arg1);
int QBitArray_fill(void* self, int aval);
void QBitArray_fill2(void* self, int val, long long first, long long last);
void QBitArray_truncate(void* self, long long pos);
const char* QBitArray_bits(void* self);
void* QBitArray_fromBits(const char* data, long long len_);
unsigned int QBitArray_toUInt32(void* self, int endianness);
int QBitArray_fill3(void* self, int aval, long long asize);
unsigned int QBitArray_toUInt322(void* self, int endianness, int* ok);
void* QBitArray_OperatorBitwiseAnd(void* self, void* a2);
void* QBitArray_OperatorBitwiseOr(void* self, void* a2);
void* QBitArray_OperatorBitwiseXor(void* self, void* a2);
void* QBitArray_OperatorBitwiseNot(void* self);
int QBitArray_OperatorEqual(void* self, void* rhs);
int QBitArray_OperatorNotEqual(void* self, void* rhs);
void QBitArray_delete(void* self);

#ifdef __cplusplus
}
#endif
