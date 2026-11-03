#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyNotifier_base(void* self);
void QPropertyNotifier_chain(void* self, void** out);
void* QPropertyNotifier_new();
void QPropertyNotifier_delete(void* self);

#ifdef __cplusplus
}
#endif
