#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsBlurEffect_base(void* self);
void QGraphicsBlurEffect_chain(void* self, void** out);
void QGraphicsBlurEffect__set_override(void* self, int idx);
void* QGraphicsBlurEffect_fromQObject(void* o);
void* QGraphicsBlurEffect_new(long long __goid);
void* QGraphicsBlurEffect_new2(long long __goid, void* parent);
char* QGraphicsBlurEffect_tr(const char* s, long long* __out_len);
double QGraphicsBlurEffect_blurRadius(void* self);
int QGraphicsBlurEffect_blurHints(void* self);
void QGraphicsBlurEffect_setBlurRadius(void* self, double blurRadius);
void QGraphicsBlurEffect_setBlurHints(void* self, int hints);
void QGraphicsBlurEffect_blurRadiusChanged(void* self, double blurRadius);
void QGraphicsBlurEffect_blurHintsChanged(void* self, int hints);
void QGraphicsBlurEffect_updateBoundingRect(void* self);
int QGraphicsBlurEffect_sourceIsPixmap(void* self);
void* QGraphicsBlurEffect_sourceBoundingRect(void* self);
void QGraphicsBlurEffect_drawSource(void* self, void* painter);
void* QGraphicsBlurEffect_sourcePixmap(void* self);
void* QGraphicsBlurEffect_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsBlurEffect_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsBlurEffect_sender(void* self);
int QGraphicsBlurEffect_senderSignalIndex(void* self);
int QGraphicsBlurEffect_receivers(void* self, const char* signal);
int QGraphicsBlurEffect_isSignalConnected(void* self, void* signal);
char* QGraphicsBlurEffect_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsBlurEffect_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsBlurEffect_sourceBoundingRect2(void* self, int system);
void* QGraphicsBlurEffect_sourcePixmap2(void* self, int system);
void* QGraphicsBlurEffect_sourcePixmap3(void* self, int system, void* offset);
void* QGraphicsBlurEffect_sourcePixmap4(void* self, int system, void* offset, int mode);
void* QGraphicsBlurEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsBlurEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsBlurEffect_metaObject_default(void* self);
void* QGraphicsBlurEffect_boundingRectFor_default(void* self, void* rect);
void QGraphicsBlurEffect_draw_default(void* self, void* painter);
void QGraphicsBlurEffect_sourceChanged_default(void* self, int flags);
int QGraphicsBlurEffect_event_default(void* self, void* event);
int QGraphicsBlurEffect_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsBlurEffect_timerEvent_default(void* self, void* event);
void QGraphicsBlurEffect_childEvent_default(void* self, void* event);
void QGraphicsBlurEffect_customEvent_default(void* self, void* event);
void QGraphicsBlurEffect_connectNotify_default(void* self, void* signal);
void QGraphicsBlurEffect_disconnectNotify_default(void* self, void* signal);
void QGraphicsBlurEffect_connect_blurRadiusChanged(void* self, long long id);
void QGraphicsBlurEffect_connect_blurHintsChanged(void* self, long long id);
void QGraphicsBlurEffect_connect_enabledChanged(void* self, long long id);
void QGraphicsBlurEffect_connect_destroyed(void* self, long long id);
void QGraphicsBlurEffect_connect_destroyed2(void* self, long long id);
void QGraphicsBlurEffect_connect_objectNameChanged(void* self, long long id);
void QGraphicsBlurEffect__disconnect(long long id);
void* QGraphicsBlurEffect__Go_new(long long id, void* parent);
void QGraphicsBlurEffect_delete(void* self);

#ifdef __cplusplus
}
#endif
