#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsRotation_base(void* self);
void QGraphicsRotation_chain(void* self, void** out);
void QGraphicsRotation__set_override(void* self, int idx);
void* QGraphicsRotation_fromQObject(void* o);
void* QGraphicsRotation_new(long long __goid);
void* QGraphicsRotation_new2(long long __goid, void* parent);
char* QGraphicsRotation_tr(const char* s, long long* __out_len);
void* QGraphicsRotation_origin(void* self);
void QGraphicsRotation_setOrigin(void* self, void* point);
double QGraphicsRotation_angle(void* self);
void QGraphicsRotation_setAngle(void* self, double arg1);
void* QGraphicsRotation_axis(void* self);
void QGraphicsRotation_setAxis(void* self, void* axis);
void QGraphicsRotation_setAxis2(void* self, int axis);
void QGraphicsRotation_originChanged(void* self);
void QGraphicsRotation_angleChanged(void* self);
void QGraphicsRotation_axisChanged(void* self);
void QGraphicsRotation_update(void* self);
void* QGraphicsRotation_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsRotation_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsRotation_sender(void* self);
int QGraphicsRotation_senderSignalIndex(void* self);
int QGraphicsRotation_receivers(void* self, const char* signal);
int QGraphicsRotation_isSignalConnected(void* self, void* signal);
char* QGraphicsRotation_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsRotation_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsRotation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsRotation_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsRotation_metaObject_default(void* self);
void QGraphicsRotation_applyTo_default(void* self, void* matrix);
int QGraphicsRotation_event_default(void* self, void* event);
int QGraphicsRotation_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsRotation_timerEvent_default(void* self, void* event);
void QGraphicsRotation_childEvent_default(void* self, void* event);
void QGraphicsRotation_customEvent_default(void* self, void* event);
void QGraphicsRotation_connectNotify_default(void* self, void* signal);
void QGraphicsRotation_disconnectNotify_default(void* self, void* signal);
void QGraphicsRotation_connect_originChanged(void* self, long long id);
void QGraphicsRotation_connect_angleChanged(void* self, long long id);
void QGraphicsRotation_connect_axisChanged(void* self, long long id);
void QGraphicsRotation_connect_destroyed(void* self, long long id);
void QGraphicsRotation_connect_destroyed2(void* self, long long id);
void QGraphicsRotation_connect_objectNameChanged(void* self, long long id);
void QGraphicsRotation__disconnect(long long id);
void* QGraphicsRotation__Go_new(long long id, void* parent);
void QGraphicsRotation_delete(void* self);

#ifdef __cplusplus
}
#endif
