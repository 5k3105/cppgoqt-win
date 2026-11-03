#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QEventLoop_base(void* self);
void QEventLoop__set_override(void* self, int idx);
void* QEventLoop_fromQObject(void* o);
void* QEventLoop_new(long long __goid);
void* QEventLoop_new2(long long __goid, void* parent);
char* QEventLoop_tr(const char* s, long long* __out_len);
int QEventLoop_processEvents(void* self);
void QEventLoop_processEvents2(void* self, int flags, int maximumTime);
void QEventLoop_processEvents3(void* self, int flags, void* deadline);
int QEventLoop_exec(void* self);
int QEventLoop_isRunning(void* self);
void QEventLoop_wakeUp(void* self);
void QEventLoop_exit(void* self);
void QEventLoop_quit(void* self);
void* QEventLoop_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QEventLoop_connect2(void* sender, void* signal, void* receiver, void* method);
void* QEventLoop_sender(void* self);
int QEventLoop_senderSignalIndex(void* self);
int QEventLoop_receivers(void* self, const char* signal);
int QEventLoop_isSignalConnected(void* self, void* signal);
char* QEventLoop_tr2(const char* s, const char* c, long long* __out_len);
char* QEventLoop_tr3(const char* s, const char* c, int n, long long* __out_len);
int QEventLoop_processEvents4(void* self, int flags);
int QEventLoop_exec2(void* self, int flags);
void QEventLoop_exit2(void* self, int returnCode);
void* QEventLoop_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QEventLoop_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QEventLoop_metaObject_default(void* self);
int QEventLoop_event_default(void* self, void* event);
int QEventLoop_eventFilter_default(void* self, void* watched, void* event);
void QEventLoop_timerEvent_default(void* self, void* event);
void QEventLoop_childEvent_default(void* self, void* event);
void QEventLoop_customEvent_default(void* self, void* event);
void QEventLoop_connectNotify_default(void* self, void* signal);
void QEventLoop_disconnectNotify_default(void* self, void* signal);
void QEventLoop_connect_destroyed(void* self, long long id);
void QEventLoop_connect_destroyed2(void* self, long long id);
void QEventLoop_connect_objectNameChanged(void* self, long long id);
void QEventLoop__disconnect(long long id);
void* QEventLoop__Go_new(long long id, void* parent);
void QEventLoop_delete(void* self);

#ifdef __cplusplus
}
#endif
