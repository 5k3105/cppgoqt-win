#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPanGesture_base(void* self);
void QPanGesture_chain(void* self, void** out);
void QPanGesture__set_override(void* self, int idx);
void* QPanGesture_fromQObject(void* o);
void* QPanGesture_new(long long __goid);
void* QPanGesture_new2(long long __goid, void* parent);
char* QPanGesture_tr(const char* s, long long* __out_len);
void* QPanGesture_lastOffset(void* self);
void* QPanGesture_offset(void* self);
void* QPanGesture_delta(void* self);
double QPanGesture_acceleration(void* self);
void QPanGesture_setLastOffset(void* self, void* value);
void QPanGesture_setOffset(void* self, void* value);
void QPanGesture_setAcceleration(void* self, double value);
void* QPanGesture_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPanGesture_connect2(void* sender, void* signal, void* receiver, void* method);
void* QPanGesture_sender(void* self);
int QPanGesture_senderSignalIndex(void* self);
int QPanGesture_receivers(void* self, const char* signal);
int QPanGesture_isSignalConnected(void* self, void* signal);
char* QPanGesture_tr2(const char* s, const char* c, long long* __out_len);
char* QPanGesture_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPanGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPanGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPanGesture_metaObject_default(void* self);
int QPanGesture_event_default(void* self, void* event);
int QPanGesture_eventFilter_default(void* self, void* watched, void* event);
void QPanGesture_timerEvent_default(void* self, void* event);
void QPanGesture_childEvent_default(void* self, void* event);
void QPanGesture_customEvent_default(void* self, void* event);
void QPanGesture_connectNotify_default(void* self, void* signal);
void QPanGesture_disconnectNotify_default(void* self, void* signal);
void QPanGesture_connect_destroyed(void* self, long long id);
void QPanGesture_connect_destroyed2(void* self, long long id);
void QPanGesture_connect_objectNameChanged(void* self, long long id);
void QPanGesture__disconnect(long long id);
void* QPanGesture__Go_new(long long id, void* parent);
void QPanGesture_delete(void* self);

#ifdef __cplusplus
}
#endif
