#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPaintEvent_base(void* self);
void* QPaintEvent_new(void* paintRegion);
void* QPaintEvent_new2(void* paintRect);
void* QPaintEvent_rect(void* self);
void* QPaintEvent_region(void* self);
int QPaintEvent_registerEventType();
int QPaintEvent_registerEventType2(int hint);
void QPaintEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
