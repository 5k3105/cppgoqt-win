#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QRunnable__set_override(void* self, int idx);
void QRunnable_run(void* self);
int QRunnable_autoDelete(void* self);
void QRunnable_setAutoDelete(void* self, int autoDelete);
void* QRunnable__Go_new(long long id);
void QRunnable_delete(void* self);

#ifdef __cplusplus
}
#endif
