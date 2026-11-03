#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSwipeGesture_base(void* self);
void QSwipeGesture_chain(void* self, void** out);
void QSwipeGesture__set_override(void* self, int idx);
void* QSwipeGesture_fromQObject(void* o);
void* QSwipeGesture_new(long long __goid);
void* QSwipeGesture_new2(long long __goid, void* parent);
char* QSwipeGesture_tr(const char* s, long long* __out_len);
int QSwipeGesture_horizontalDirection(void* self);
int QSwipeGesture_verticalDirection(void* self);
double QSwipeGesture_swipeAngle(void* self);
void QSwipeGesture_setSwipeAngle(void* self, double value);
void* QSwipeGesture_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSwipeGesture_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSwipeGesture_sender(void* self);
int QSwipeGesture_senderSignalIndex(void* self);
int QSwipeGesture_receivers(void* self, const char* signal);
int QSwipeGesture_isSignalConnected(void* self, void* signal);
char* QSwipeGesture_tr2(const char* s, const char* c, long long* __out_len);
char* QSwipeGesture_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSwipeGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSwipeGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSwipeGesture_metaObject_default(void* self);
int QSwipeGesture_event_default(void* self, void* event);
int QSwipeGesture_eventFilter_default(void* self, void* watched, void* event);
void QSwipeGesture_timerEvent_default(void* self, void* event);
void QSwipeGesture_childEvent_default(void* self, void* event);
void QSwipeGesture_customEvent_default(void* self, void* event);
void QSwipeGesture_connectNotify_default(void* self, void* signal);
void QSwipeGesture_disconnectNotify_default(void* self, void* signal);
void QSwipeGesture_connect_destroyed(void* self, long long id);
void QSwipeGesture_connect_destroyed2(void* self, long long id);
void QSwipeGesture_connect_objectNameChanged(void* self, long long id);
void QSwipeGesture__disconnect(long long id);
void* QSwipeGesture__Go_new(long long id, void* parent);
void QSwipeGesture_delete(void* self);

#ifdef __cplusplus
}
#endif
