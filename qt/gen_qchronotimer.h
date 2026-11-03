#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QChronoTimer_base(void* self);
void QChronoTimer__set_override(void* self, int idx);
void* QChronoTimer_fromQObject(void* o);
void* QChronoTimer_new(long long __goid, long long nsec);
void* QChronoTimer_new2(long long __goid);
void* QChronoTimer_new3(long long __goid, long long nsec, void* parent);
void* QChronoTimer_new4(long long __goid, void* parent);
char* QChronoTimer_tr(const char* s, long long* __out_len);
int QChronoTimer_isActive(void* self);
int QChronoTimer_id(void* self);
void QChronoTimer_setInterval(void* self, long long nsec);
long long QChronoTimer_interval(void* self);
long long QChronoTimer_remainingTime(void* self);
void QChronoTimer_setTimerType(void* self, int atype);
int QChronoTimer_timerType(void* self);
void QChronoTimer_setSingleShot(void* self, int singleShot);
int QChronoTimer_isSingleShot(void* self);
void QChronoTimer_start(void* self);
void QChronoTimer_stop(void* self);
void* QChronoTimer_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QChronoTimer_connect2(void* sender, void* signal, void* receiver, void* method);
void* QChronoTimer_sender(void* self);
int QChronoTimer_senderSignalIndex(void* self);
int QChronoTimer_receivers(void* self, const char* signal);
int QChronoTimer_isSignalConnected(void* self, void* signal);
char* QChronoTimer_tr2(const char* s, const char* c, long long* __out_len);
char* QChronoTimer_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QChronoTimer_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QChronoTimer_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QChronoTimer_metaObject_default(void* self);
void QChronoTimer_timerEvent_default(void* self, void* arg1);
int QChronoTimer_event_default(void* self, void* event);
int QChronoTimer_eventFilter_default(void* self, void* watched, void* event);
void QChronoTimer_childEvent_default(void* self, void* event);
void QChronoTimer_customEvent_default(void* self, void* event);
void QChronoTimer_connectNotify_default(void* self, void* signal);
void QChronoTimer_disconnectNotify_default(void* self, void* signal);
void QChronoTimer_connect_timeout(void* self, long long id);
void QChronoTimer_connect_destroyed(void* self, long long id);
void QChronoTimer_connect_destroyed2(void* self, long long id);
void QChronoTimer_connect_objectNameChanged(void* self, long long id);
void QChronoTimer__disconnect(long long id);
void* QChronoTimer__Go_new(long long id, void* parent);
void QChronoTimer_delete(void* self);

#ifdef __cplusplus
}
#endif
