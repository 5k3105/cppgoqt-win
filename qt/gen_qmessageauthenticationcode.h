#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMessageAuthenticationCode_new(int method);
void* QMessageAuthenticationCode_new2(int method, void* key);
void QMessageAuthenticationCode_swap(void* self, void* other);
void QMessageAuthenticationCode_reset(void* self);
void QMessageAuthenticationCode_setKey(void* self, void* key);
void QMessageAuthenticationCode_addData(void* self, const char* data, long long length);
void QMessageAuthenticationCode_addData2(void* self, void* data);
int QMessageAuthenticationCode_addData3(void* self, void* device);
void* QMessageAuthenticationCode_resultView(void* self);
void* QMessageAuthenticationCode_result(void* self);
void* QMessageAuthenticationCode_hash(void* message, void* key, int method);
void QMessageAuthenticationCode_delete(void* self);

#ifdef __cplusplus
}
#endif
