#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTapGesture_base(void* self);
void QTapGesture_chain(void* self, void** out);
void QTapGesture__set_override(void* self, int idx);
void* QTapGesture_fromQObject(void* o);
void* QTapGesture_new(long long __goid);
void* QTapGesture_new2(long long __goid, void* parent);
char* QTapGesture_tr(const char* s, long long* __out_len);
void* QTapGesture_position(void* self);
void QTapGesture_setPosition(void* self, void* pos);
void* QTapGesture_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTapGesture_connect2(void* sender, void* signal, void* receiver, void* method);
void* QTapGesture_sender(void* self);
int QTapGesture_senderSignalIndex(void* self);
int QTapGesture_receivers(void* self, const char* signal);
int QTapGesture_isSignalConnected(void* self, void* signal);
char* QTapGesture_tr2(const char* s, const char* c, long long* __out_len);
char* QTapGesture_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTapGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTapGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QTapGesture_metaObject_default(void* self);
int QTapGesture_event_default(void* self, void* event);
int QTapGesture_eventFilter_default(void* self, void* watched, void* event);
void QTapGesture_timerEvent_default(void* self, void* event);
void QTapGesture_childEvent_default(void* self, void* event);
void QTapGesture_customEvent_default(void* self, void* event);
void QTapGesture_connectNotify_default(void* self, void* signal);
void QTapGesture_disconnectNotify_default(void* self, void* signal);
void QTapGesture_connect_destroyed(void* self, long long id);
void QTapGesture_connect_destroyed2(void* self, long long id);
void QTapGesture_connect_objectNameChanged(void* self, long long id);
void QTapGesture__disconnect(long long id);
void* QTapGesture__Go_new(long long id, void* parent);
void QTapGesture_delete(void* self);

#ifdef __cplusplus
}
#endif
