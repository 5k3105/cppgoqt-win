#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMicrophonePermission_new();
void* QMicrophonePermission_new2(void* other);
void* QMicrophonePermission_OperatorAssign(void* self, void* other);
void QMicrophonePermission_swap(void* self, void* other);
void QMicrophonePermission_delete(void* self);

#ifdef __cplusplus
}
#endif
