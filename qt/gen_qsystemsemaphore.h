#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSystemSemaphore_new(void* key);
void* QSystemSemaphore_new2(const char* key, long long key_len);
void* QSystemSemaphore_new3(void* key, int initialValue);
void* QSystemSemaphore_new4(void* key, int initialValue, int arg3);
void* QSystemSemaphore_new5(const char* key, long long key_len, int initialValue);
void* QSystemSemaphore_new6(const char* key, long long key_len, int initialValue, int mode);
char* QSystemSemaphore_tr(const char* sourceText, long long* __out_len);
void QSystemSemaphore_setNativeKey(void* self, void* key);
void QSystemSemaphore_setNativeKey2(void* self, const char* key, long long key_len);
void* QSystemSemaphore_nativeIpcKey(void* self);
void QSystemSemaphore_setKey(void* self, const char* key, long long key_len);
char* QSystemSemaphore_key(void* self, long long* __out_len);
int QSystemSemaphore_acquire(void* self);
int QSystemSemaphore_release(void* self);
int QSystemSemaphore_error(void* self);
char* QSystemSemaphore_errorString(void* self, long long* __out_len);
int QSystemSemaphore_isKeyTypeSupported(int type_);
void* QSystemSemaphore_platformSafeKey(const char* key, long long key_len);
void* QSystemSemaphore_legacyNativeKey(const char* key, long long key_len);
char* QSystemSemaphore_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QSystemSemaphore_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
void QSystemSemaphore_setNativeKey3(void* self, void* key, int initialValue);
void QSystemSemaphore_setNativeKey4(void* self, void* key, int initialValue, int arg3);
void QSystemSemaphore_setNativeKey5(void* self, const char* key, long long key_len, int initialValue);
void QSystemSemaphore_setNativeKey6(void* self, const char* key, long long key_len, int initialValue, int mode);
void QSystemSemaphore_setNativeKey7(void* self, const char* key, long long key_len, int initialValue, int mode, int type_);
void QSystemSemaphore_setKey2(void* self, const char* key, long long key_len, int initialValue);
void QSystemSemaphore_setKey3(void* self, const char* key, long long key_len, int initialValue, int mode);
int QSystemSemaphore_release2(void* self, int n);
void* QSystemSemaphore_platformSafeKey2(const char* key, long long key_len, int type_);
void* QSystemSemaphore_legacyNativeKey2(const char* key, long long key_len, int type_);
void QSystemSemaphore_delete(void* self);

#ifdef __cplusplus
}
#endif
