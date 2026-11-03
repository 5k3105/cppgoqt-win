#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsEffect_base(void* self);
void QGraphicsEffect__set_override(void* self, int idx);
void* QGraphicsEffect_fromQObject(void* o);
char* QGraphicsEffect_tr(const char* s, long long* __out_len);
void* QGraphicsEffect_boundingRectFor(void* self, void* sourceRect);
void* QGraphicsEffect_boundingRect(void* self);
int QGraphicsEffect_isEnabled(void* self);
void QGraphicsEffect_setEnabled(void* self, int enable);
void QGraphicsEffect_update(void* self);
void QGraphicsEffect_enabledChanged(void* self, int enabled);
void QGraphicsEffect_updateBoundingRect(void* self);
int QGraphicsEffect_sourceIsPixmap(void* self);
void* QGraphicsEffect_sourceBoundingRect(void* self);
void QGraphicsEffect_drawSource(void* self, void* painter);
void* QGraphicsEffect_sourcePixmap(void* self);
void* QGraphicsEffect_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsEffect_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsEffect_sender(void* self);
int QGraphicsEffect_senderSignalIndex(void* self);
int QGraphicsEffect_receivers(void* self, const char* signal);
int QGraphicsEffect_isSignalConnected(void* self, void* signal);
char* QGraphicsEffect_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsEffect_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsEffect_sourceBoundingRect2(void* self, int system);
void* QGraphicsEffect_sourcePixmap2(void* self, int system);
void* QGraphicsEffect_sourcePixmap3(void* self, int system, void* offset);
void* QGraphicsEffect_sourcePixmap4(void* self, int system, void* offset, int mode);
void* QGraphicsEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsEffect_boundingRectFor_default(void* self, void* sourceRect);
void* QGraphicsEffect_metaObject_default(void* self);
void QGraphicsEffect_sourceChanged_default(void* self, int flags);
int QGraphicsEffect_event_default(void* self, void* event);
int QGraphicsEffect_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsEffect_timerEvent_default(void* self, void* event);
void QGraphicsEffect_childEvent_default(void* self, void* event);
void QGraphicsEffect_customEvent_default(void* self, void* event);
void QGraphicsEffect_connectNotify_default(void* self, void* signal);
void QGraphicsEffect_disconnectNotify_default(void* self, void* signal);
void QGraphicsEffect_connect_enabledChanged(void* self, long long id);
void QGraphicsEffect_connect_destroyed(void* self, long long id);
void QGraphicsEffect_connect_destroyed2(void* self, long long id);
void QGraphicsEffect_connect_objectNameChanged(void* self, long long id);
void QGraphicsEffect__disconnect(long long id);
void* QGraphicsEffect__Go_new(long long id, void* parent);
void QGraphicsEffect_delete(void* self);

#ifdef __cplusplus
}
#endif
