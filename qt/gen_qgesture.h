#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGesture_base(void* self);
void QGesture__set_override(void* self, int idx);
void* QGesture_fromQObject(void* o);
void* QGesture_new(long long __goid);
void* QGesture_new2(long long __goid, void* parent);
char* QGesture_tr(const char* s, long long* __out_len);
int QGesture_gestureType(void* self);
int QGesture_state(void* self);
void* QGesture_hotSpot(void* self);
void QGesture_setHotSpot(void* self, void* value);
int QGesture_hasHotSpot(void* self);
void QGesture_unsetHotSpot(void* self);
void QGesture_setGestureCancelPolicy(void* self, int policy);
int QGesture_gestureCancelPolicy(void* self);
void* QGesture_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGesture_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGesture_sender(void* self);
int QGesture_senderSignalIndex(void* self);
int QGesture_receivers(void* self, const char* signal);
int QGesture_isSignalConnected(void* self, void* signal);
char* QGesture_tr2(const char* s, const char* c, long long* __out_len);
char* QGesture_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGesture_metaObject_default(void* self);
int QGesture_event_default(void* self, void* event);
int QGesture_eventFilter_default(void* self, void* watched, void* event);
void QGesture_timerEvent_default(void* self, void* event);
void QGesture_childEvent_default(void* self, void* event);
void QGesture_customEvent_default(void* self, void* event);
void QGesture_connectNotify_default(void* self, void* signal);
void QGesture_disconnectNotify_default(void* self, void* signal);
void QGesture_connect_destroyed(void* self, long long id);
void QGesture_connect_destroyed2(void* self, long long id);
void QGesture_connect_objectNameChanged(void* self, long long id);
void QGesture__disconnect(long long id);
void* QGesture__Go_new(long long id, void* parent);
void QGesture_delete(void* self);

#ifdef __cplusplus
}
#endif
