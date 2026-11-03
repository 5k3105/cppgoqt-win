#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QOffscreenSurface_base(void* self);
void QOffscreenSurface__set_override(void* self, int idx);
void* QOffscreenSurface_fromQObject(void* o);
void* QOffscreenSurface_new(long long __goid);
void* QOffscreenSurface_new2(long long __goid, void* screen);
void* QOffscreenSurface_new3(long long __goid, void* screen, void* parent);
char* QOffscreenSurface_tr(const char* s, long long* __out_len);
int QOffscreenSurface_surfaceType(void* self);
void QOffscreenSurface_create(void* self);
void QOffscreenSurface_destroy(void* self);
int QOffscreenSurface_isValid(void* self);
void QOffscreenSurface_setFormat(void* self, void* format);
void* QOffscreenSurface_format(void* self);
void* QOffscreenSurface_requestedFormat(void* self);
void* QOffscreenSurface_size(void* self);
void* QOffscreenSurface_screen(void* self);
void QOffscreenSurface_setScreen(void* self, void* screen);
void QOffscreenSurface_screenChanged(void* self, void* screen);
void* QOffscreenSurface_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QOffscreenSurface_connect2(void* sender, void* signal, void* receiver, void* method);
void* QOffscreenSurface_sender(void* self);
int QOffscreenSurface_senderSignalIndex(void* self);
int QOffscreenSurface_receivers(void* self, const char* signal);
int QOffscreenSurface_isSignalConnected(void* self, void* signal);
int QOffscreenSurface_surfaceClass(void* self);
int QOffscreenSurface_supportsOpenGL(void* self);
char* QOffscreenSurface_tr2(const char* s, const char* c, long long* __out_len);
char* QOffscreenSurface_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QOffscreenSurface_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QOffscreenSurface_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QOffscreenSurface_asQSurface(void* self);
int QOffscreenSurface_surfaceType_default(void* self);
void* QOffscreenSurface_format_default(void* self);
void* QOffscreenSurface_size_default(void* self);
void* QOffscreenSurface_metaObject_default(void* self);
int QOffscreenSurface_event_default(void* self, void* event);
int QOffscreenSurface_eventFilter_default(void* self, void* watched, void* event);
void QOffscreenSurface_timerEvent_default(void* self, void* event);
void QOffscreenSurface_childEvent_default(void* self, void* event);
void QOffscreenSurface_customEvent_default(void* self, void* event);
void QOffscreenSurface_connectNotify_default(void* self, void* signal);
void QOffscreenSurface_disconnectNotify_default(void* self, void* signal);
void QOffscreenSurface_connect_screenChanged(void* self, long long id);
void QOffscreenSurface_connect_destroyed(void* self, long long id);
void QOffscreenSurface_connect_destroyed2(void* self, long long id);
void QOffscreenSurface_connect_objectNameChanged(void* self, long long id);
void QOffscreenSurface__disconnect(long long id);
void* QOffscreenSurface__Go_new(long long id);
void QOffscreenSurface_delete(void* self);

#ifdef __cplusplus
}
#endif
