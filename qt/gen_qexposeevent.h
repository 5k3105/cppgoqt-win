#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QExposeEvent_base(void* self);
void* QExposeEvent_new(void* m_region);
void* QExposeEvent_region(void* self);
int QExposeEvent_registerEventType();
int QExposeEvent_registerEventType2(int hint);
void QExposeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
