#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTapAndHoldGesture_base(void* self);
void QTapAndHoldGesture_chain(void* self, void** out);
void QTapAndHoldGesture__set_override(void* self, int idx);
void* QTapAndHoldGesture_fromQObject(void* o);
void* QTapAndHoldGesture_new(long long __goid);
void* QTapAndHoldGesture_new2(long long __goid, void* parent);
char* QTapAndHoldGesture_tr(const char* s, long long* __out_len);
void* QTapAndHoldGesture_position(void* self);
void QTapAndHoldGesture_setPosition(void* self, void* pos);
void QTapAndHoldGesture_setTimeout(int msecs);
int QTapAndHoldGesture_timeout();
void* QTapAndHoldGesture_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTapAndHoldGesture_connect2(void* sender, void* signal, void* receiver, void* method);
void* QTapAndHoldGesture_sender(void* self);
int QTapAndHoldGesture_senderSignalIndex(void* self);
int QTapAndHoldGesture_receivers(void* self, const char* signal);
int QTapAndHoldGesture_isSignalConnected(void* self, void* signal);
char* QTapAndHoldGesture_tr2(const char* s, const char* c, long long* __out_len);
char* QTapAndHoldGesture_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTapAndHoldGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTapAndHoldGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QTapAndHoldGesture_metaObject_default(void* self);
int QTapAndHoldGesture_event_default(void* self, void* event);
int QTapAndHoldGesture_eventFilter_default(void* self, void* watched, void* event);
void QTapAndHoldGesture_timerEvent_default(void* self, void* event);
void QTapAndHoldGesture_childEvent_default(void* self, void* event);
void QTapAndHoldGesture_customEvent_default(void* self, void* event);
void QTapAndHoldGesture_connectNotify_default(void* self, void* signal);
void QTapAndHoldGesture_disconnectNotify_default(void* self, void* signal);
void QTapAndHoldGesture_connect_destroyed(void* self, long long id);
void QTapAndHoldGesture_connect_destroyed2(void* self, long long id);
void QTapAndHoldGesture_connect_objectNameChanged(void* self, long long id);
void QTapAndHoldGesture__disconnect(long long id);
void* QTapAndHoldGesture__Go_new(long long id, void* parent);
void QTapAndHoldGesture_delete(void* self);

#ifdef __cplusplus
}
#endif
