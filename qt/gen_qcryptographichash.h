#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCryptographicHash_new(int method);
void QCryptographicHash_swap(void* self, void* other);
void QCryptographicHash_reset(void* self);
int QCryptographicHash_algorithm(void* self);
void QCryptographicHash_addData(void* self, const char* data, long long length);
void QCryptographicHash_addData2(void* self, void* data);
int QCryptographicHash_addData3(void* self, void* device);
void* QCryptographicHash_result(void* self);
void* QCryptographicHash_resultView(void* self);
void* QCryptographicHash_hash(void* data, int method);
int QCryptographicHash_hashLength(int method);
int QCryptographicHash_supportsAlgorithm(int method);
void QCryptographicHash_delete(void* self);

#ifdef __cplusplus
}
#endif
