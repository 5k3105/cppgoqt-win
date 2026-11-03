#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsScale_base(void* self);
void QGraphicsScale_chain(void* self, void** out);
void QGraphicsScale__set_override(void* self, int idx);
void* QGraphicsScale_fromQObject(void* o);
void* QGraphicsScale_new(long long __goid);
void* QGraphicsScale_new2(long long __goid, void* parent);
char* QGraphicsScale_tr(const char* s, long long* __out_len);
void* QGraphicsScale_origin(void* self);
void QGraphicsScale_setOrigin(void* self, void* point);
double QGraphicsScale_xScale(void* self);
void QGraphicsScale_setXScale(void* self, double arg1);
double QGraphicsScale_yScale(void* self);
void QGraphicsScale_setYScale(void* self, double arg1);
double QGraphicsScale_zScale(void* self);
void QGraphicsScale_setZScale(void* self, double arg1);
void QGraphicsScale_originChanged(void* self);
void QGraphicsScale_xScaleChanged(void* self);
void QGraphicsScale_yScaleChanged(void* self);
void QGraphicsScale_zScaleChanged(void* self);
void QGraphicsScale_scaleChanged(void* self);
void QGraphicsScale_update(void* self);
void* QGraphicsScale_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsScale_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsScale_sender(void* self);
int QGraphicsScale_senderSignalIndex(void* self);
int QGraphicsScale_receivers(void* self, const char* signal);
int QGraphicsScale_isSignalConnected(void* self, void* signal);
char* QGraphicsScale_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsScale_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsScale_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsScale_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsScale_metaObject_default(void* self);
void QGraphicsScale_applyTo_default(void* self, void* matrix);
int QGraphicsScale_event_default(void* self, void* event);
int QGraphicsScale_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsScale_timerEvent_default(void* self, void* event);
void QGraphicsScale_childEvent_default(void* self, void* event);
void QGraphicsScale_customEvent_default(void* self, void* event);
void QGraphicsScale_connectNotify_default(void* self, void* signal);
void QGraphicsScale_disconnectNotify_default(void* self, void* signal);
void QGraphicsScale_connect_originChanged(void* self, long long id);
void QGraphicsScale_connect_xScaleChanged(void* self, long long id);
void QGraphicsScale_connect_yScaleChanged(void* self, long long id);
void QGraphicsScale_connect_zScaleChanged(void* self, long long id);
void QGraphicsScale_connect_scaleChanged(void* self, long long id);
void QGraphicsScale_connect_destroyed(void* self, long long id);
void QGraphicsScale_connect_destroyed2(void* self, long long id);
void QGraphicsScale_connect_objectNameChanged(void* self, long long id);
void QGraphicsScale__disconnect(long long id);
void* QGraphicsScale__Go_new(long long id, void* parent);
void QGraphicsScale_delete(void* self);

#ifdef __cplusplus
}
#endif
