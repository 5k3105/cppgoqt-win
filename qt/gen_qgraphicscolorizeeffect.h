#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsColorizeEffect_base(void* self);
void QGraphicsColorizeEffect_chain(void* self, void** out);
void QGraphicsColorizeEffect__set_override(void* self, int idx);
void* QGraphicsColorizeEffect_fromQObject(void* o);
void* QGraphicsColorizeEffect_new(long long __goid);
void* QGraphicsColorizeEffect_new2(long long __goid, void* parent);
char* QGraphicsColorizeEffect_tr(const char* s, long long* __out_len);
void* QGraphicsColorizeEffect_color(void* self);
double QGraphicsColorizeEffect_strength(void* self);
void QGraphicsColorizeEffect_setColor(void* self, void* c);
void QGraphicsColorizeEffect_setStrength(void* self, double strength);
void QGraphicsColorizeEffect_colorChanged(void* self, void* color);
void QGraphicsColorizeEffect_strengthChanged(void* self, double strength);
void QGraphicsColorizeEffect_updateBoundingRect(void* self);
int QGraphicsColorizeEffect_sourceIsPixmap(void* self);
void* QGraphicsColorizeEffect_sourceBoundingRect(void* self);
void QGraphicsColorizeEffect_drawSource(void* self, void* painter);
void* QGraphicsColorizeEffect_sourcePixmap(void* self);
void* QGraphicsColorizeEffect_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsColorizeEffect_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsColorizeEffect_sender(void* self);
int QGraphicsColorizeEffect_senderSignalIndex(void* self);
int QGraphicsColorizeEffect_receivers(void* self, const char* signal);
int QGraphicsColorizeEffect_isSignalConnected(void* self, void* signal);
char* QGraphicsColorizeEffect_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsColorizeEffect_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsColorizeEffect_sourceBoundingRect2(void* self, int system);
void* QGraphicsColorizeEffect_sourcePixmap2(void* self, int system);
void* QGraphicsColorizeEffect_sourcePixmap3(void* self, int system, void* offset);
void* QGraphicsColorizeEffect_sourcePixmap4(void* self, int system, void* offset, int mode);
void* QGraphicsColorizeEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsColorizeEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsColorizeEffect_metaObject_default(void* self);
void QGraphicsColorizeEffect_draw_default(void* self, void* painter);
void* QGraphicsColorizeEffect_boundingRectFor_default(void* self, void* sourceRect);
void QGraphicsColorizeEffect_sourceChanged_default(void* self, int flags);
int QGraphicsColorizeEffect_event_default(void* self, void* event);
int QGraphicsColorizeEffect_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsColorizeEffect_timerEvent_default(void* self, void* event);
void QGraphicsColorizeEffect_childEvent_default(void* self, void* event);
void QGraphicsColorizeEffect_customEvent_default(void* self, void* event);
void QGraphicsColorizeEffect_connectNotify_default(void* self, void* signal);
void QGraphicsColorizeEffect_disconnectNotify_default(void* self, void* signal);
void QGraphicsColorizeEffect_connect_colorChanged(void* self, long long id);
void QGraphicsColorizeEffect_connect_strengthChanged(void* self, long long id);
void QGraphicsColorizeEffect_connect_enabledChanged(void* self, long long id);
void QGraphicsColorizeEffect_connect_destroyed(void* self, long long id);
void QGraphicsColorizeEffect_connect_destroyed2(void* self, long long id);
void QGraphicsColorizeEffect_connect_objectNameChanged(void* self, long long id);
void QGraphicsColorizeEffect__disconnect(long long id);
void* QGraphicsColorizeEffect__Go_new(long long id, void* parent);
void QGraphicsColorizeEffect_delete(void* self);

#ifdef __cplusplus
}
#endif
