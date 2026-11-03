#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWinEventNotifier_base(void* self);
void QWinEventNotifier__set_override(void* self, int idx);
void* QWinEventNotifier_fromQObject(void* o);
void* QWinEventNotifier_new(long long __goid);
void* QWinEventNotifier_new2(long long __goid, void* parent);
char* QWinEventNotifier_tr(const char* s, long long* __out_len);
int QWinEventNotifier_isEnabled(void* self);
void QWinEventNotifier_setEnabled(void* self, int enable);
void* QWinEventNotifier_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QWinEventNotifier_connect2(void* sender, void* signal, void* receiver, void* method);
void* QWinEventNotifier_sender(void* self);
int QWinEventNotifier_senderSignalIndex(void* self);
int QWinEventNotifier_receivers(void* self, const char* signal);
int QWinEventNotifier_isSignalConnected(void* self, void* signal);
char* QWinEventNotifier_tr2(const char* s, const char* c, long long* __out_len);
char* QWinEventNotifier_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QWinEventNotifier_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QWinEventNotifier_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QWinEventNotifier_metaObject_default(void* self);
int QWinEventNotifier_event_default(void* self, void* e);
int QWinEventNotifier_eventFilter_default(void* self, void* watched, void* event);
void QWinEventNotifier_timerEvent_default(void* self, void* event);
void QWinEventNotifier_childEvent_default(void* self, void* event);
void QWinEventNotifier_customEvent_default(void* self, void* event);
void QWinEventNotifier_connectNotify_default(void* self, void* signal);
void QWinEventNotifier_disconnectNotify_default(void* self, void* signal);
void QWinEventNotifier_connect_destroyed(void* self, long long id);
void QWinEventNotifier_connect_destroyed2(void* self, long long id);
void QWinEventNotifier_connect_objectNameChanged(void* self, long long id);
void QWinEventNotifier__disconnect(long long id);
void* QWinEventNotifier__Go_new(long long id, void* parent);
void QWinEventNotifier_delete(void* self);

#ifdef __cplusplus
}
#endif
