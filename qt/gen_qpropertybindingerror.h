#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyBindingError_new();
void* QPropertyBindingError_new2(int type_);
void* QPropertyBindingError_new3(void* other);
void* QPropertyBindingError_new4(int type_, const char* description, long long description_len);
void* QPropertyBindingError_OperatorAssign(void* self, void* other);
int QPropertyBindingError_hasError(void* self);
int QPropertyBindingError_type(void* self);
char* QPropertyBindingError_description(void* self, long long* __out_len);
void QPropertyBindingError_delete(void* self);

#ifdef __cplusplus
}
#endif
