#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QNativeIpcKey_new();
void* QNativeIpcKey_new2(int type_);
void* QNativeIpcKey_new3(const char* k, long long k_len);
void* QNativeIpcKey_new4(void* other);
void* QNativeIpcKey_new5(const char* k, long long k_len, int type_);
int QNativeIpcKey_legacyDefaultTypeForOs();
void* QNativeIpcKey_OperatorAssign(void* self, void* other);
void QNativeIpcKey_swap(void* self, void* other);
int QNativeIpcKey_isEmpty(void* self);
int QNativeIpcKey_isValid(void* self);
int QNativeIpcKey_type(void* self);
void QNativeIpcKey_setType(void* self, int type_);
char* QNativeIpcKey_nativeKey(void* self, long long* __out_len);
void QNativeIpcKey_setNativeKey(void* self, const char* newKey, long long newKey_len);
char* QNativeIpcKey_toString(void* self, long long* __out_len);
void* QNativeIpcKey_fromString(const char* string_, long long string__len);
int QNativeIpcKey_OperatorEqual(void* self, void* rhs);
int QNativeIpcKey_OperatorNotEqual(void* self, void* rhs);
void QNativeIpcKey_delete(void* self);

#ifdef __cplusplus
}
#endif
