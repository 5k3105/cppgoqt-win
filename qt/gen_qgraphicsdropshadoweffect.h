#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsDropShadowEffect_base(void* self);
void QGraphicsDropShadowEffect_chain(void* self, void** out);
void QGraphicsDropShadowEffect__set_override(void* self, int idx);
void* QGraphicsDropShadowEffect_fromQObject(void* o);
void* QGraphicsDropShadowEffect_new(long long __goid);
void* QGraphicsDropShadowEffect_new2(long long __goid, void* parent);
char* QGraphicsDropShadowEffect_tr(const char* s, long long* __out_len);
void* QGraphicsDropShadowEffect_offset(void* self);
double QGraphicsDropShadowEffect_xOffset(void* self);
double QGraphicsDropShadowEffect_yOffset(void* self);
double QGraphicsDropShadowEffect_blurRadius(void* self);
void* QGraphicsDropShadowEffect_color(void* self);
void QGraphicsDropShadowEffect_setOffset(void* self, void* ofs);
void QGraphicsDropShadowEffect_setOffset2(void* self, double dx, double dy);
void QGraphicsDropShadowEffect_setOffset3(void* self, double d);
void QGraphicsDropShadowEffect_setXOffset(void* self, double dx);
void QGraphicsDropShadowEffect_setYOffset(void* self, double dy);
void QGraphicsDropShadowEffect_setBlurRadius(void* self, double blurRadius);
void QGraphicsDropShadowEffect_setColor(void* self, void* color);
void QGraphicsDropShadowEffect_offsetChanged(void* self, void* offset);
void QGraphicsDropShadowEffect_blurRadiusChanged(void* self, double blurRadius);
void QGraphicsDropShadowEffect_colorChanged(void* self, void* color);
void QGraphicsDropShadowEffect_updateBoundingRect(void* self);
int QGraphicsDropShadowEffect_sourceIsPixmap(void* self);
void* QGraphicsDropShadowEffect_sourceBoundingRect(void* self);
void QGraphicsDropShadowEffect_drawSource(void* self, void* painter);
void* QGraphicsDropShadowEffect_sourcePixmap(void* self);
void* QGraphicsDropShadowEffect_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsDropShadowEffect_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsDropShadowEffect_sender(void* self);
int QGraphicsDropShadowEffect_senderSignalIndex(void* self);
int QGraphicsDropShadowEffect_receivers(void* self, const char* signal);
int QGraphicsDropShadowEffect_isSignalConnected(void* self, void* signal);
char* QGraphicsDropShadowEffect_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsDropShadowEffect_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsDropShadowEffect_sourceBoundingRect2(void* self, int system);
void* QGraphicsDropShadowEffect_sourcePixmap2(void* self, int system);
void* QGraphicsDropShadowEffect_sourcePixmap3(void* self, int system, void* offset);
void* QGraphicsDropShadowEffect_sourcePixmap4(void* self, int system, void* offset, int mode);
void* QGraphicsDropShadowEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsDropShadowEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsDropShadowEffect_metaObject_default(void* self);
void* QGraphicsDropShadowEffect_boundingRectFor_default(void* self, void* rect);
void QGraphicsDropShadowEffect_draw_default(void* self, void* painter);
void QGraphicsDropShadowEffect_sourceChanged_default(void* self, int flags);
int QGraphicsDropShadowEffect_event_default(void* self, void* event);
int QGraphicsDropShadowEffect_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsDropShadowEffect_timerEvent_default(void* self, void* event);
void QGraphicsDropShadowEffect_childEvent_default(void* self, void* event);
void QGraphicsDropShadowEffect_customEvent_default(void* self, void* event);
void QGraphicsDropShadowEffect_connectNotify_default(void* self, void* signal);
void QGraphicsDropShadowEffect_disconnectNotify_default(void* self, void* signal);
void QGraphicsDropShadowEffect_connect_offsetChanged(void* self, long long id);
void QGraphicsDropShadowEffect_connect_blurRadiusChanged(void* self, long long id);
void QGraphicsDropShadowEffect_connect_colorChanged(void* self, long long id);
void QGraphicsDropShadowEffect_connect_enabledChanged(void* self, long long id);
void QGraphicsDropShadowEffect_connect_destroyed(void* self, long long id);
void QGraphicsDropShadowEffect_connect_destroyed2(void* self, long long id);
void QGraphicsDropShadowEffect_connect_objectNameChanged(void* self, long long id);
void QGraphicsDropShadowEffect__disconnect(long long id);
void* QGraphicsDropShadowEffect__Go_new(long long id, void* parent);
void QGraphicsDropShadowEffect_delete(void* self);

#ifdef __cplusplus
}
#endif
