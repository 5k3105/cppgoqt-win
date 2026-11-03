#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QThread_base(void* self);
void QThread__set_override(void* self, int idx);
void* QThread_fromQObject(void* o);
void* QThread_new(long long __goid);
void* QThread_new2(long long __goid, void* parent);
char* QThread_tr(const char* s, long long* __out_len);
void* QThread_currentThread();
int QThread_isMainThread();
int QThread_idealThreadCount();
void QThread_yieldCurrentThread();
void QThread_setPriority(void* self, int priority);
int QThread_priority(void* self);
int QThread_isFinished(void* self);
int QThread_isRunning(void* self);
void QThread_requestInterruption(void* self);
int QThread_isInterruptionRequested(void* self);
void QThread_setStackSize(void* self, unsigned int stackSize);
unsigned int QThread_stackSize(void* self);
void* QThread_eventDispatcher(void* self);
void QThread_setEventDispatcher(void* self, void* eventDispatcher);
int QThread_loopLevel(void* self);
int QThread_isCurrentThread(void* self);
void QThread_setServiceLevel(void* self, int serviceLevel);
int QThread_serviceLevel(void* self);
void QThread_start(void* self);
void QThread_terminate(void* self);
void QThread_exit(void* self);
void QThread_quit(void* self);
int QThread_wait(void* self);
int QThread_wait2(void* self, unsigned long time);
void QThread_sleep(unsigned long arg1);
void QThread_msleep(unsigned long arg1);
void QThread_usleep(unsigned long arg1);
void QThread_sleep2(long long nsec);
int QThread_exec(void* self);
void* QThread_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QThread_connect2(void* sender, void* signal, void* receiver, void* method);
void* QThread_sender(void* self);
int QThread_senderSignalIndex(void* self);
int QThread_receivers(void* self, const char* signal);
int QThread_isSignalConnected(void* self, void* signal);
char* QThread_tr2(const char* s, const char* c, long long* __out_len);
char* QThread_tr3(const char* s, const char* c, int n, long long* __out_len);
void QThread_start2(void* self, int arg1);
void QThread_exit2(void* self, int retcode);
int QThread_wait3(void* self, void* deadline);
void* QThread_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QThread_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QThread_metaObject_default(void* self);
int QThread_event_default(void* self, void* event);
void QThread_run_default(void* self);
int QThread_eventFilter_default(void* self, void* watched, void* event);
void QThread_timerEvent_default(void* self, void* event);
void QThread_childEvent_default(void* self, void* event);
void QThread_customEvent_default(void* self, void* event);
void QThread_connectNotify_default(void* self, void* signal);
void QThread_disconnectNotify_default(void* self, void* signal);
void QThread_connect_started(void* self, long long id);
void QThread_connect_finished(void* self, long long id);
void QThread_connect_destroyed(void* self, long long id);
void QThread_connect_destroyed2(void* self, long long id);
void QThread_connect_objectNameChanged(void* self, long long id);
void QThread__disconnect(long long id);
void* QThread__Go_new(long long id, void* parent);
void QThread_delete(void* self);

#ifdef __cplusplus
}
#endif
