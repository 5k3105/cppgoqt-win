#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsItemAnimation_base(void* self);
void QGraphicsItemAnimation__set_override(void* self, int idx);
void* QGraphicsItemAnimation_fromQObject(void* o);
void* QGraphicsItemAnimation_new(long long __goid);
void* QGraphicsItemAnimation_new2(long long __goid, void* parent);
char* QGraphicsItemAnimation_tr(const char* s, long long* __out_len);
void* QGraphicsItemAnimation_item(void* self);
void QGraphicsItemAnimation_setItem(void* self, void* item);
void* QGraphicsItemAnimation_timeLine(void* self);
void QGraphicsItemAnimation_setTimeLine(void* self, void* timeLine);
void* QGraphicsItemAnimation_posAt(void* self, double step);
void QGraphicsItemAnimation_setPosAt(void* self, double step, void* pos);
void* QGraphicsItemAnimation_transformAt(void* self, double step);
double QGraphicsItemAnimation_rotationAt(void* self, double step);
void QGraphicsItemAnimation_setRotationAt(void* self, double step, double angle);
double QGraphicsItemAnimation_xTranslationAt(void* self, double step);
double QGraphicsItemAnimation_yTranslationAt(void* self, double step);
void QGraphicsItemAnimation_setTranslationAt(void* self, double step, double dx, double dy);
double QGraphicsItemAnimation_verticalScaleAt(void* self, double step);
double QGraphicsItemAnimation_horizontalScaleAt(void* self, double step);
void QGraphicsItemAnimation_setScaleAt(void* self, double step, double sx, double sy);
double QGraphicsItemAnimation_verticalShearAt(void* self, double step);
double QGraphicsItemAnimation_horizontalShearAt(void* self, double step);
void QGraphicsItemAnimation_setShearAt(void* self, double step, double sh, double sv);
void QGraphicsItemAnimation_clear(void* self);
void QGraphicsItemAnimation_setStep(void* self, double x);
void* QGraphicsItemAnimation_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsItemAnimation_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGraphicsItemAnimation_sender(void* self);
int QGraphicsItemAnimation_senderSignalIndex(void* self);
int QGraphicsItemAnimation_receivers(void* self, const char* signal);
int QGraphicsItemAnimation_isSignalConnected(void* self, void* signal);
char* QGraphicsItemAnimation_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsItemAnimation_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsItemAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsItemAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGraphicsItemAnimation_metaObject_default(void* self);
void QGraphicsItemAnimation_beforeAnimationStep_default(void* self, double step);
void QGraphicsItemAnimation_afterAnimationStep_default(void* self, double step);
int QGraphicsItemAnimation_event_default(void* self, void* event);
int QGraphicsItemAnimation_eventFilter_default(void* self, void* watched, void* event);
void QGraphicsItemAnimation_timerEvent_default(void* self, void* event);
void QGraphicsItemAnimation_childEvent_default(void* self, void* event);
void QGraphicsItemAnimation_customEvent_default(void* self, void* event);
void QGraphicsItemAnimation_connectNotify_default(void* self, void* signal);
void QGraphicsItemAnimation_disconnectNotify_default(void* self, void* signal);
void QGraphicsItemAnimation_connect_destroyed(void* self, long long id);
void QGraphicsItemAnimation_connect_destroyed2(void* self, long long id);
void QGraphicsItemAnimation_connect_objectNameChanged(void* self, long long id);
void QGraphicsItemAnimation__disconnect(long long id);
void* QGraphicsItemAnimation__Go_new(long long id, void* parent);
void QGraphicsItemAnimation_delete(void* self);

#ifdef __cplusplus
}
#endif
