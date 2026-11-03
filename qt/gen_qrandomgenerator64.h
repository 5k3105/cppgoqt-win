#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRandomGenerator64_base(void* self);
void* QRandomGenerator64_new();
void* QRandomGenerator64_new2(void* other);
void* QRandomGenerator64_new3(void* arg1);
void* QRandomGenerator64_new4(unsigned int seedValue);
unsigned long long QRandomGenerator64_generate(void* self);
unsigned long long QRandomGenerator64_OperatorCall(void* self);
void QRandomGenerator64_discard(void* self, unsigned long long z);
unsigned long long QRandomGenerator64_min();
unsigned long long QRandomGenerator64_max();
void* QRandomGenerator64_system();
void* QRandomGenerator64_global();
void* QRandomGenerator64_securelySeeded();
void* QRandomGenerator64_OperatorAssign(void* self, void* arg1);
void QRandomGenerator64_delete(void* self);

#ifdef __cplusplus
}
#endif
