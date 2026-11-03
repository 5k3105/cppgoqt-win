#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractEventDispatcher_base(void* self);
void QAbstractEventDispatcher__set_override(void* self, int idx);
void* QAbstractEventDispatcher_fromQObject(void* o);
char* QAbstractEventDispatcher_tr(const char* s, long long* __out_len);
void* QAbstractEventDispatcher_instance();
int QAbstractEventDispatcher_processEvents(void* self, int flags);
void QAbstractEventDispatcher_registerSocketNotifier(void* self, void* notifier);
void QAbstractEventDispatcher_unregisterSocketNotifier(void* self, void* notifier);
int QAbstractEventDispatcher_registerTimer(void* self, long long interval, int timerType, void* object);
void QAbstractEventDispatcher_registerTimer2(void* self, int timerId, long long interval, int timerType, void* object);
int QAbstractEventDispatcher_unregisterTimer(void* self, int timerId);
int QAbstractEventDispatcher_unregisterTimers(void* self, void* object);
void** QAbstractEventDispatcher_registeredTimers(void* self, void* object, long long* __out_len);
int QAbstractEventDispatcher_remainingTime(void* self, int timerId);
int QAbstractEventDispatcher_unregisterTimer2(void* self, int timerId);
void** QAbstractEventDispatcher_timersForObject(void* self, void* object, long long* __out_len);
void QAbstractEventDispatcher_wakeUp(void* self);
void QAbstractEventDispatcher_interrupt(void* self);
void QAbstractEventDispatcher_startingUp(void* self);
void QAbstractEventDispatcher_closingDown(void* self);
void QAbstractEventDispatcher_installNativeEventFilter(void* self, void* filterObj);
void QAbstractEventDispatcher_removeNativeEventFilter(void* self, void* filterObj);
void QAbstractEventDispatcher_aboutToBlock(void* self);
void QAbstractEventDispatcher_awake(void* self);
void* QAbstractEventDispatcher_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractEventDispatcher_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAbstractEventDispatcher_sender(void* self);
int QAbstractEventDispatcher_senderSignalIndex(void* self);
int QAbstractEventDispatcher_receivers(void* self, const char* signal);
int QAbstractEventDispatcher_isSignalConnected(void* self, void* signal);
char* QAbstractEventDispatcher_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractEventDispatcher_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAbstractEventDispatcher_instance2(void* thread);
void* QAbstractEventDispatcher_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractEventDispatcher_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QAbstractEventDispatcher_startingUp_default(void* self);
void QAbstractEventDispatcher_closingDown_default(void* self);
void* QAbstractEventDispatcher_metaObject_default(void* self);
int QAbstractEventDispatcher_event_default(void* self, void* event);
int QAbstractEventDispatcher_eventFilter_default(void* self, void* watched, void* event);
void QAbstractEventDispatcher_timerEvent_default(void* self, void* event);
void QAbstractEventDispatcher_childEvent_default(void* self, void* event);
void QAbstractEventDispatcher_customEvent_default(void* self, void* event);
void QAbstractEventDispatcher_connectNotify_default(void* self, void* signal);
void QAbstractEventDispatcher_disconnectNotify_default(void* self, void* signal);
void QAbstractEventDispatcher_connect_aboutToBlock(void* self, long long id);
void QAbstractEventDispatcher_connect_awake(void* self, long long id);
void QAbstractEventDispatcher_connect_destroyed(void* self, long long id);
void QAbstractEventDispatcher_connect_destroyed2(void* self, long long id);
void QAbstractEventDispatcher_connect_objectNameChanged(void* self, long long id);
void QAbstractEventDispatcher__disconnect(long long id);
void* QAbstractEventDispatcher__Go_new(long long id, void* parent);
void QAbstractEventDispatcher_delete(void* self);

#ifdef __cplusplus
}
#endif
