#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QFactoryInterface__set_override(void* self, int idx);
char** QFactoryInterface_keys(void* self, long long* __out_len);
void* QFactoryInterface_OperatorAssign(void* self, void* arg1);
void* QFactoryInterface__Go_new(long long id);
void QFactoryInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
