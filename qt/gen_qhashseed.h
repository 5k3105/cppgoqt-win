#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QHashSeed_new();
void* QHashSeed_new2(unsigned long long d);
unsigned long long QHashSeed_ToSizeT(void* self);
void* QHashSeed_globalSeed();
void QHashSeed_setDeterministicGlobalSeed();
void QHashSeed_resetRandomGlobalSeed();
void QHashSeed_delete(void* self);

#ifdef __cplusplus
}
#endif
