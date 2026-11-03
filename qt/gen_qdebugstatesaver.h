#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDebugStateSaver_new(void* dbg);
void QDebugStateSaver_delete(void* self);

#ifdef __cplusplus
}
#endif
