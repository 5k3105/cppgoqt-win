#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPlatformSurfaceEvent_base(void* self);
void* QPlatformSurfaceEvent_new(int surfaceEventType);
int QPlatformSurfaceEvent_surfaceEventType(void* self);
int QPlatformSurfaceEvent_registerEventType();
int QPlatformSurfaceEvent_registerEventType2(int hint);
void QPlatformSurfaceEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
