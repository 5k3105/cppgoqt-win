#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QByteArray__FromBase64Result_new(void* arg1);
void* QByteArray__FromBase64Result_new2();
void QByteArray__FromBase64Result_swap(void* self, void* other);
int QByteArray__FromBase64Result_ToBool(void* self);
void* QByteArray__FromBase64Result_OperatorMultiply(void* self);
void* QByteArray__FromBase64Result_OperatorMultiply2(void* self);
void* QByteArray__FromBase64Result_OperatorAssign(void* self, void* arg1);
int QByteArray__FromBase64Result_OperatorEqual(void* self, void* rhs);
int QByteArray__FromBase64Result_OperatorNotEqual(void* self, void* rhs);
void* QByteArray__FromBase64Result_decoded_get(void* self);
void QByteArray__FromBase64Result_decoded_set(void* self, void* decoded);
int QByteArray__FromBase64Result_decodingStatus_get(void* self);
void QByteArray__FromBase64Result_decodingStatus_set(void* self, int decodingStatus);
void QByteArray__FromBase64Result_delete(void* self);

#ifdef __cplusplus
}
#endif
