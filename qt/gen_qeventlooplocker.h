#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QEventLoopLocker_new();
void* QEventLoopLocker_new2(void* loop);
void* QEventLoopLocker_new3(void* thread);
void QEventLoopLocker_swap(void* self, void* other);
void QEventLoopLocker_delete(void* self);

#ifdef __cplusplus
}
#endif
