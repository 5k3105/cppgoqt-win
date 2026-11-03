#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAnimationDriver_base(void* self);
void QAnimationDriver__set_override(void* self, int idx);
void* QAnimationDriver_fromQObject(void* o);
void* QAnimationDriver_new(long long __goid);
void* QAnimationDriver_new2(long long __goid, void* parent);
char* QAnimationDriver_tr(const char* s, long long* __out_len);
void QAnimationDriver_advance(void* self);
void QAnimationDriver_install(void* self);
void QAnimationDriver_uninstall(void* self);
int QAnimationDriver_isRunning(void* self);
long long QAnimationDriver_elapsed(void* self);
void QAnimationDriver_started(void* self);
void QAnimationDriver_stopped(void* self);
void QAnimationDriver_advanceAnimation(void* self);
void* QAnimationDriver_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAnimationDriver_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAnimationDriver_sender(void* self);
int QAnimationDriver_senderSignalIndex(void* self);
int QAnimationDriver_receivers(void* self, const char* signal);
int QAnimationDriver_isSignalConnected(void* self, void* signal);
char* QAnimationDriver_tr2(const char* s, const char* c, long long* __out_len);
char* QAnimationDriver_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAnimationDriver_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAnimationDriver_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QAnimationDriver_advance_default(void* self);
long long QAnimationDriver_elapsed_default(void* self);
void* QAnimationDriver_metaObject_default(void* self);
void QAnimationDriver_start_default(void* self);
void QAnimationDriver_stop_default(void* self);
int QAnimationDriver_event_default(void* self, void* event);
int QAnimationDriver_eventFilter_default(void* self, void* watched, void* event);
void QAnimationDriver_timerEvent_default(void* self, void* event);
void QAnimationDriver_childEvent_default(void* self, void* event);
void QAnimationDriver_customEvent_default(void* self, void* event);
void QAnimationDriver_connectNotify_default(void* self, void* signal);
void QAnimationDriver_disconnectNotify_default(void* self, void* signal);
void QAnimationDriver_connect_started(void* self, long long id);
void QAnimationDriver_connect_stopped(void* self, long long id);
void QAnimationDriver_connect_destroyed(void* self, long long id);
void QAnimationDriver_connect_destroyed2(void* self, long long id);
void QAnimationDriver_connect_objectNameChanged(void* self, long long id);
void QAnimationDriver__disconnect(long long id);
void* QAnimationDriver__Go_new(long long id, void* parent);
void QAnimationDriver_delete(void* self);

#ifdef __cplusplus
}
#endif
