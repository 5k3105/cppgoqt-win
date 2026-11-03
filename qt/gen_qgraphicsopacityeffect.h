#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsOpacityEffect_base(void* self);
void QGraphicsOpacityEffect_chain(void* self, void** out);
void QGraphicsOpacityEffect__set_override(void* self, int idx);
void* QGraphicsOpacityEffect_fromQObject(void* o);
void* QGraphicsOpacityEffect_new(long long __goid);
void* QGraphicsOpacityEffect_new2(long long __goid, void* parent);
char* QGraphicsOpacityEffect_tr(const char* s, long long* __out_len);
double QGraphicsOpacityEffect_opacity(void* self);
void* QGraphicsOpacityEffect_opacityMask(void* self);
void QGraphicsOpacityEffect_setOpacity(void* self, double opacity);
void QGraphicsOpacityEffect_setOpacityMask(void* self, void* mask);
void QGraphicsOpacityEffect_opacityChanged(void* self, double opacity);
void QGraphicsOpacityEffect_opacityMaskChanged(void* self, void* mask);
void QGraphicsOpacityEffect_updateBoundingRect(void* self);
int QGraphicsOpacityEffect_sourceIsPixmap(void* self);
void* QGraphicsOpacityEffect_sourceBoundingRect(void* self);
void QGraphicsOpacityEffect_drawSource(void* self, void* painter);
void* QGraphicsOpacityEffect_sourcePixmap(void* self);
void* QGraphicsOpacityEffect_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsOpacityEffect_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsOpacityEffect_sender(void* self);
int QGraphicsOpacityEffect_senderSignalIndex(void* self);
int QGraphicsOpacityEffect_receivers(void* self, const char* signal);
int QGraphicsOpacityEffect_isSignalConnected(void* self, void* signal);
char* QGraphicsOpacityEffect_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsOpacityEffect_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsOpacityEffect_sourceBoundingRect2(void* self, int system);
void* QGraphicsOpacityEffect_sourcePixmap2(void* self, int system);
void* QGraphicsOpacityEffect_sourcePixmap3(void* self, int system, void* offset);
void* QGraphicsOpacityEffect_sourcePixmap4(void* self, int system, void* offset, int mode);
void* QGraphicsOpacityEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsOpacityEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsOpacityEffect_metaObject_default(void* self);
void QGraphicsOpacityEffect_draw_default(void* self, void* painter);
void* QGraphicsOpacityEffect_boundingRectFor_default(void* self, void* sourceRect);
void QGraphicsOpacityEffect_sourceChanged_default(void* self, int flags);
int QGraphicsOpacityEffect_event_default(void* self, void* event);
int QGraphicsOpacityEffect_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsOpacityEffect_timerEvent_default(void* self, void* event);
void QGraphicsOpacityEffect_childEvent_default(void* self, void* event);
void QGraphicsOpacityEffect_customEvent_default(void* self, void* event);
void QGraphicsOpacityEffect_connectNotify_default(void* self, void* signal);
void QGraphicsOpacityEffect_disconnectNotify_default(void* self, void* signal);
void QGraphicsOpacityEffect_connect_opacityChanged(void* self, long long id);
void QGraphicsOpacityEffect_connect_opacityMaskChanged(void* self, long long id);
void QGraphicsOpacityEffect_connect_enabledChanged(void* self, long long id);
void QGraphicsOpacityEffect_connect_destroyed(void* self, long long id);
void QGraphicsOpacityEffect_connect_destroyed2(void* self, long long id);
void QGraphicsOpacityEffect_connect_objectNameChanged(void* self, long long id);
void QGraphicsOpacityEffect__disconnect(long long id);
void* QGraphicsOpacityEffect__Go_new(long long id, void* parent);
void QGraphicsOpacityEffect_delete(void* self);

#ifdef __cplusplus
}
#endif
