#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QObjectCleanupHandler_base(void* self);
void QObjectCleanupHandler__set_override(void* self, int idx);
void* QObjectCleanupHandler_fromQObject(void* o);
void* QObjectCleanupHandler_new(long long __goid);
char* QObjectCleanupHandler_tr(const char* s, long long* __out_len);
void* QObjectCleanupHandler_add(void* self, void* object);
void QObjectCleanupHandler_remove(void* self, void* object);
int QObjectCleanupHandler_isEmpty(void* self);
void QObjectCleanupHandler_clear(void* self);
void* QObjectCleanupHandler_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QObjectCleanupHandler_connect2(void* sender, void* signal, void* receiver, void* method);
void* QObjectCleanupHandler_sender(void* self);
int QObjectCleanupHandler_senderSignalIndex(void* self);
int QObjectCleanupHandler_receivers(void* self, const char* signal);
int QObjectCleanupHandler_isSignalConnected(void* self, void* signal);
char* QObjectCleanupHandler_tr2(const char* s, const char* c, long long* __out_len);
char* QObjectCleanupHandler_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QObjectCleanupHandler_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QObjectCleanupHandler_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QObjectCleanupHandler_metaObject_default(void* self);
int QObjectCleanupHandler_event_default(void* self, void* event);
int QObjectCleanupHandler_eventFilter_default(void* self, void* watched, void* event);
void QObjectCleanupHandler_timerEvent_default(void* self, void* event);
void QObjectCleanupHandler_childEvent_default(void* self, void* event);
void QObjectCleanupHandler_customEvent_default(void* self, void* event);
void QObjectCleanupHandler_connectNotify_default(void* self, void* signal);
void QObjectCleanupHandler_disconnectNotify_default(void* self, void* signal);
void QObjectCleanupHandler_connect_destroyed(void* self, long long id);
void QObjectCleanupHandler_connect_destroyed2(void* self, long long id);
void QObjectCleanupHandler_connect_objectNameChanged(void* self, long long id);
void QObjectCleanupHandler__disconnect(long long id);
void* QObjectCleanupHandler__Go_new(long long id);
void QObjectCleanupHandler_delete(void* self);

#ifdef __cplusplus
}
#endif
