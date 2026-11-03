#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSurfaceFormat_new();
void* QSurfaceFormat_new2(int options);
void* QSurfaceFormat_new3(void* other);
void* QSurfaceFormat_OperatorAssign(void* self, void* other);
void QSurfaceFormat_setDepthBufferSize(void* self, int size);
int QSurfaceFormat_depthBufferSize(void* self);
void QSurfaceFormat_setStencilBufferSize(void* self, int size);
int QSurfaceFormat_stencilBufferSize(void* self);
void QSurfaceFormat_setRedBufferSize(void* self, int size);
int QSurfaceFormat_redBufferSize(void* self);
void QSurfaceFormat_setGreenBufferSize(void* self, int size);
int QSurfaceFormat_greenBufferSize(void* self);
void QSurfaceFormat_setBlueBufferSize(void* self, int size);
int QSurfaceFormat_blueBufferSize(void* self);
void QSurfaceFormat_setAlphaBufferSize(void* self, int size);
int QSurfaceFormat_alphaBufferSize(void* self);
void QSurfaceFormat_setSamples(void* self, int numSamples);
int QSurfaceFormat_samples(void* self);
void QSurfaceFormat_setSwapBehavior(void* self, int behavior);
int QSurfaceFormat_swapBehavior(void* self);
int QSurfaceFormat_hasAlpha(void* self);
void QSurfaceFormat_setProfile(void* self, int profile);
int QSurfaceFormat_profile(void* self);
void QSurfaceFormat_setRenderableType(void* self, int type_);
int QSurfaceFormat_renderableType(void* self);
void QSurfaceFormat_setMajorVersion(void* self, int majorVersion);
int QSurfaceFormat_majorVersion(void* self);
void QSurfaceFormat_setMinorVersion(void* self, int minorVersion);
int QSurfaceFormat_minorVersion(void* self);
int QSurfaceFormat_version(void* self, int* __out_b);
void QSurfaceFormat_setVersion(void* self, int major, int minor);
int QSurfaceFormat_stereo(void* self);
void QSurfaceFormat_setStereo(void* self, int enable);
void QSurfaceFormat_setOptions(void* self, int options);
void QSurfaceFormat_setOption(void* self, int option);
int QSurfaceFormat_testOption(void* self, int option);
int QSurfaceFormat_options(void* self);
int QSurfaceFormat_swapInterval(void* self);
void QSurfaceFormat_setSwapInterval(void* self, int interval);
void* QSurfaceFormat_colorSpace(void* self);
void QSurfaceFormat_setColorSpace(void* self, void* colorSpace);
void QSurfaceFormat_setColorSpace2(void* self, int colorSpace);
void QSurfaceFormat_setDefaultFormat(void* format);
void* QSurfaceFormat_defaultFormat();
void QSurfaceFormat_setOption2(void* self, int option, int on);
int QSurfaceFormat_OperatorEqual(void* self, void* rhs);
int QSurfaceFormat_OperatorNotEqual(void* self, void* rhs);
void QSurfaceFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
