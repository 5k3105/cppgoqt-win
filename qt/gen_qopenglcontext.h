#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QOpenGLContext_base(void* self);
void QOpenGLContext__set_override(void* self, int idx);
void* QOpenGLContext_fromQObject(void* o);
void* QOpenGLContext_new(long long __goid);
void* QOpenGLContext_new2(long long __goid, void* parent);
char* QOpenGLContext_tr(const char* s, long long* __out_len);
void QOpenGLContext_setFormat(void* self, void* format);
void QOpenGLContext_setShareContext(void* self, void* shareContext);
void QOpenGLContext_setScreen(void* self, void* screen);
int QOpenGLContext_create(void* self);
int QOpenGLContext_isValid(void* self);
void* QOpenGLContext_format(void* self);
void* QOpenGLContext_shareContext(void* self);
void* QOpenGLContext_shareGroup(void* self);
void* QOpenGLContext_screen(void* self);
unsigned int QOpenGLContext_defaultFramebufferObject(void* self);
int QOpenGLContext_makeCurrent(void* self, void* surface);
void QOpenGLContext_doneCurrent(void* self);
void QOpenGLContext_swapBuffers(void* self, void* surface);
void* QOpenGLContext_surface(void* self);
void* QOpenGLContext_currentContext();
int QOpenGLContext_areSharing(void* first, void* second);
void* QOpenGLContext_functions(void* self);
void* QOpenGLContext_extraFunctions(void* self);
void** QOpenGLContext_extensions(void* self, long long* __out_len);
int QOpenGLContext_hasExtension(void* self, void* extension);
int QOpenGLContext_openGLModuleType();
int QOpenGLContext_isOpenGLES(void* self);
int QOpenGLContext_supportsThreadedOpenGL();
void* QOpenGLContext_globalShareContext();
void QOpenGLContext_aboutToBeDestroyed(void* self);
void* QOpenGLContext_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QOpenGLContext_connect2(void* sender, void* signal, void* receiver, void* method);
void* QOpenGLContext_sender(void* self);
int QOpenGLContext_senderSignalIndex(void* self);
int QOpenGLContext_receivers(void* self, const char* signal);
int QOpenGLContext_isSignalConnected(void* self, void* signal);
char* QOpenGLContext_tr2(const char* s, const char* c, long long* __out_len);
char* QOpenGLContext_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QOpenGLContext_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QOpenGLContext_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QOpenGLContext_metaObject_default(void* self);
int QOpenGLContext_event_default(void* self, void* event);
int QOpenGLContext_eventFilter_default(void* self, void* watched, void* event);
void QOpenGLContext_timerEvent_default(void* self, void* event);
void QOpenGLContext_childEvent_default(void* self, void* event);
void QOpenGLContext_customEvent_default(void* self, void* event);
void QOpenGLContext_connectNotify_default(void* self, void* signal);
void QOpenGLContext_disconnectNotify_default(void* self, void* signal);
void QOpenGLContext_connect_aboutToBeDestroyed(void* self, long long id);
void QOpenGLContext_connect_destroyed(void* self, long long id);
void QOpenGLContext_connect_destroyed2(void* self, long long id);
void QOpenGLContext_connect_objectNameChanged(void* self, long long id);
void QOpenGLContext__disconnect(long long id);
void* QOpenGLContext__Go_new(long long id, void* parent);
void QOpenGLContext_delete(void* self);

#ifdef __cplusplus
}
#endif
