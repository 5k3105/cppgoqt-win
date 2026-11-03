#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSignalBlocker_new(void* o);
void* QSignalBlocker_new2(void* o);
void QSignalBlocker_reblock(void* self);
void QSignalBlocker_unblock(void* self);
void QSignalBlocker_dismiss(void* self);
void QSignalBlocker_delete(void* self);

#ifdef __cplusplus
}
#endif
