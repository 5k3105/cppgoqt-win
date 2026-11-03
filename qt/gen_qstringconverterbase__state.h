#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStringConverterBase__State_new();
void* QStringConverterBase__State_new2(int f);
void QStringConverterBase__State_clear(void* self);
void QStringConverterBase__State_reset(void* self);
int QStringConverterBase__State_flags_get(void* self);
void QStringConverterBase__State_flags_set(void* self, int flags);
int QStringConverterBase__State_internalState_get(void* self);
void QStringConverterBase__State_internalState_set(void* self, int internalState);
long long QStringConverterBase__State_remainingChars_get(void* self);
void QStringConverterBase__State_remainingChars_set(void* self, long long remainingChars);
long long QStringConverterBase__State_invalidChars_get(void* self);
void QStringConverterBase__State_invalidChars_set(void* self, long long invalidChars);
void QStringConverterBase__State_delete(void* self);

#ifdef __cplusplus
}
#endif
