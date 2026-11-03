#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsTransform_base(void* self);
void QGraphicsTransform__set_override(void* self, int idx);
void* QGraphicsTransform_fromQObject(void* o);
char* QGraphicsTransform_tr(const char* s, long long* __out_len);
void QGraphicsTransform_applyTo(void* self, void* matrix);
void QGraphicsTransform_update(void* self);
void* QGraphicsTransform_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsTransform_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsTransform_sender(void* self);
int QGraphicsTransform_senderSignalIndex(void* self);
int QGraphicsTransform_receivers(void* self, const char* signal);
int QGraphicsTransform_isSignalConnected(void* self, void* signal);
char* QGraphicsTransform_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsTransform_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsTransform_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsTransform_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsTransform_metaObject_default(void* self);
int QGraphicsTransform_event_default(void* self, void* event);
int QGraphicsTransform_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsTransform_timerEvent_default(void* self, void* event);
void QGraphicsTransform_childEvent_default(void* self, void* event);
void QGraphicsTransform_customEvent_default(void* self, void* event);
void QGraphicsTransform_connectNotify_default(void* self, void* signal);
void QGraphicsTransform_disconnectNotify_default(void* self, void* signal);
void QGraphicsTransform_connect_destroyed(void* self, long long id);
void QGraphicsTransform_connect_destroyed2(void* self, long long id);
void QGraphicsTransform_connect_objectNameChanged(void* self, long long id);
void QGraphicsTransform__disconnect(long long id);
void* QGraphicsTransform__Go_new(long long id, void* parent);
void QGraphicsTransform_delete(void* self);

#ifdef __cplusplus
}
#endif
