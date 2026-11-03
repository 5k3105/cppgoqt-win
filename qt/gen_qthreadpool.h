#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QThreadPool_base(void* self);
void QThreadPool__set_override(void* self, int idx);
void* QThreadPool_fromQObject(void* o);
void* QThreadPool_new(long long __goid);
void* QThreadPool_new2(long long __goid, void* parent);
char* QThreadPool_tr(const char* s, long long* __out_len);
void* QThreadPool_globalInstance();
void QThreadPool_start(void* self, void* runnable);
int QThreadPool_tryStart(void* self, void* runnable);
void QThreadPool_startOnReservedThread(void* self, void* runnable);
int QThreadPool_expiryTimeout(void* self);
void QThreadPool_setExpiryTimeout(void* self, int expiryTimeout);
int QThreadPool_maxThreadCount(void* self);
void QThreadPool_setMaxThreadCount(void* self, int maxThreadCount);
int QThreadPool_activeThreadCount(void* self);
void QThreadPool_setStackSize(void* self, unsigned int stackSize);
unsigned int QThreadPool_stackSize(void* self);
void QThreadPool_setThreadPriority(void* self, int priority);
int QThreadPool_threadPriority(void* self);
void QThreadPool_reserveThread(void* self);
void QThreadPool_releaseThread(void* self);
void QThreadPool_setServiceLevel(void* self, int serviceLevel);
int QThreadPool_serviceLevel(void* self);
int QThreadPool_waitForDone(void* self, int msecs);
int QThreadPool_waitForDone2(void* self);
void QThreadPool_clear(void* self);
int QThreadPool_contains(void* self, void* thread);
int QThreadPool_tryTake(void* self, void* runnable);
void* QThreadPool_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QThreadPool_connect2(void* sender, void* signal, void* receiver, void* method);
void* QThreadPool_sender(void* self);
int QThreadPool_senderSignalIndex(void* self);
int QThreadPool_receivers(void* self, const char* signal);
int QThreadPool_isSignalConnected(void* self, void* signal);
char* QThreadPool_tr2(const char* s, const char* c, long long* __out_len);
char* QThreadPool_tr3(const char* s, const char* c, int n, long long* __out_len);
void QThreadPool_start2(void* self, void* runnable, int priority);
int QThreadPool_waitForDone3(void* self, void* deadline);
void* QThreadPool_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QThreadPool_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QThreadPool_metaObject_default(void* self);
int QThreadPool_event_default(void* self, void* event);
int QThreadPool_eventFilter_default(void* self, void* watched, void* event);
void QThreadPool_timerEvent_default(void* self, void* event);
void QThreadPool_childEvent_default(void* self, void* event);
void QThreadPool_customEvent_default(void* self, void* event);
void QThreadPool_connectNotify_default(void* self, void* signal);
void QThreadPool_disconnectNotify_default(void* self, void* signal);
void QThreadPool_connect_destroyed(void* self, long long id);
void QThreadPool_connect_destroyed2(void* self, long long id);
void QThreadPool_connect_objectNameChanged(void* self, long long id);
void QThreadPool__disconnect(long long id);
void* QThreadPool__Go_new(long long id, void* parent);
void QThreadPool_delete(void* self);

#ifdef __cplusplus
}
#endif
