#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSemaphoreReleaser_new();
void* QSemaphoreReleaser_new2(void* sem);
void* QSemaphoreReleaser_new3(void* sem);
void* QSemaphoreReleaser_new4(void* sem, int n);
void* QSemaphoreReleaser_new5(void* sem, int n);
void QSemaphoreReleaser_swap(void* self, void* other);
void* QSemaphoreReleaser_semaphore(void* self);
void* QSemaphoreReleaser_cancel(void* self);
void QSemaphoreReleaser_delete(void* self);

#ifdef __cplusplus
}
#endif
