#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QProcess__CreateProcessArguments_inheritHandles_get(void* self);
void QProcess__CreateProcessArguments_inheritHandles_set(void* self, int inheritHandles);
unsigned long QProcess__CreateProcessArguments_flags_get(void* self);
void QProcess__CreateProcessArguments_flags_set(void* self, unsigned long flags);
void QProcess__CreateProcessArguments_delete(void* self);

#ifdef __cplusplus
}
#endif
