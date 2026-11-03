#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QPaintEngineState_state(void* self);
void* QPaintEngineState_pen(void* self);
void* QPaintEngineState_brush(void* self);
void* QPaintEngineState_brushOrigin(void* self);
void* QPaintEngineState_backgroundBrush(void* self);
int QPaintEngineState_backgroundMode(void* self);
void* QPaintEngineState_font(void* self);
void* QPaintEngineState_transform(void* self);
int QPaintEngineState_clipOperation(void* self);
void* QPaintEngineState_clipRegion(void* self);
void* QPaintEngineState_clipPath(void* self);
int QPaintEngineState_isClipEnabled(void* self);
int QPaintEngineState_renderHints(void* self);
int QPaintEngineState_compositionMode(void* self);
double QPaintEngineState_opacity(void* self);
void* QPaintEngineState_painter(void* self);
int QPaintEngineState_brushNeedsResolving(void* self);
int QPaintEngineState_penNeedsResolving(void* self);
void QPaintEngineState_delete(void* self);

#ifdef __cplusplus
}
#endif
