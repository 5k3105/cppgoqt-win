#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScopedPropertyUpdateGroup_new();
void QScopedPropertyUpdateGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
