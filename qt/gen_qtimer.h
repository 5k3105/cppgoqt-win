#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTimer_base(void* self);
void QTimer__set_override(void* self, int idx);
void* QTimer_fromQObject(void* o);
void* QTimer_new(long long __goid);
void* QTimer_new2(long long __goid, void* parent);
char* QTimer_tr(const char* s, long long* __out_len);
int QTimer_isActive(void* self);
int QTimer_timerId(void* self);
int QTimer_id(void* self);
void QTimer_setInterval(void* self, int msec);
int QTimer_interval(void* self);
int QTimer_remainingTime(void* self);
void QTimer_setTimerType(void* self, int atype);
int QTimer_timerType(void* self);
void QTimer_setSingleShot(void* self, int singleShot);
int QTimer_isSingleShot(void* self);
void QTimer_singleShot(int msec, void* receiver, const char* member);
void QTimer_singleShot2(int msec, int timerType, void* receiver, const char* member);
void QTimer_start(void* self, int msec);
void QTimer_start2(void* self);
void QTimer_stop(void* self);
void QTimer_setInterval2(void* self, long long value);
long long QTimer_intervalAsDuration(void* self);
long long QTimer_remainingTimeAsDuration(void* self);
void QTimer_singleShot3(long long value, void* receiver, const char* member);
void QTimer_singleShot4(long long interval, int timerType, void* receiver, const char* member);
void QTimer_start3(void* self, long long value);
void* QTimer_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTimer_connect2(void* sender, void* signal, void* receiver, void* method);
void* QTimer_sender(void* self);
int QTimer_senderSignalIndex(void* self);
int QTimer_receivers(void* self, const char* signal);
int QTimer_isSignalConnected(void* self, void* signal);
char* QTimer_tr2(const char* s, const char* c, long long* __out_len);
char* QTimer_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTimer_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTimer_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QTimer_metaObject_default(void* self);
void QTimer_timerEvent_default(void* self, void* arg1);
int QTimer_event_default(void* self, void* event);
int QTimer_eventFilter_default(void* self, void* watched, void* event);
void QTimer_childEvent_default(void* self, void* event);
void QTimer_customEvent_default(void* self, void* event);
void QTimer_connectNotify_default(void* self, void* signal);
void QTimer_disconnectNotify_default(void* self, void* signal);
void QTimer_connect_timeout(void* self, long long id);
void QTimer_connect_destroyed(void* self, long long id);
void QTimer_connect_destroyed2(void* self, long long id);
void QTimer_connect_objectNameChanged(void* self, long long id);
void QTimer__disconnect(long long id);
void* QTimer__Go_new(long long id, void* parent);
void QTimer_delete(void* self);

#ifdef __cplusplus
}
#endif
