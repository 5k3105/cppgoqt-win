#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QException__set_override(void* self, int idx);
void* QException_new(long long __goid);
void* QException_new2(long long __goid, void* arg1);
void* QException_OperatorAssign(void* self, void* arg1);
void QException_raise(void* self);
void* QException_clone(void* self);
const char* QException_what(void* self);
void QException_raise_default(void* self);
void* QException_clone_default(void* self);
void* QException__Go_new(long long id);
void QException_delete(void* self);

#ifdef __cplusplus
}
#endif
