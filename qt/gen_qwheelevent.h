#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWheelEvent_base(void* self);
void QWheelEvent_chain(void* self, void** out);
void* QWheelEvent_new(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int buttons, int modifiers, int phase, int inverted);
void* QWheelEvent_new2(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int buttons, int modifiers, int phase, int inverted, int source);
void* QWheelEvent_new3(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int buttons, int modifiers, int phase, int inverted, int source, void* device);
void* QWheelEvent_pixelDelta(void* self);
void* QWheelEvent_angleDelta(void* self);
int QWheelEvent_phase(void* self);
int QWheelEvent_inverted(void* self);
int QWheelEvent_isInverted(void* self);
int QWheelEvent_hasPixelDelta(void* self);
int QWheelEvent_source(void* self);
int QWheelEvent_registerEventType();
int QWheelEvent_registerEventType2(int hint);
void QWheelEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
