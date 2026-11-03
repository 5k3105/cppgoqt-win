#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRasterWindow_base(void* self);
void QRasterWindow_chain(void* self, void** out);
void QRasterWindow__set_override(void* self, int idx);
void* QRasterWindow_fromQObject(void* o);
void* QRasterWindow_new(long long __goid);
void* QRasterWindow_new2(long long __goid, void* parent);
char* QRasterWindow_tr(const char* s, long long* __out_len);
void* QRasterWindow_fromWinId(unsigned long long id);
void* QRasterWindow_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QRasterWindow_connect2(void* sender, void* signal, void* receiver, void* method);
void* QRasterWindow_sender(void* self);
int QRasterWindow_senderSignalIndex(void* self);
int QRasterWindow_receivers(void* self, const char* signal);
int QRasterWindow_isSignalConnected(void* self, void* signal);
double QRasterWindow_devicePixelRatioFScale();
int QRasterWindow_encodeMetricF(int metric, double value);
double QRasterWindow_getDecodedMetricF(void* self, int metricA, int metricB);
char* QRasterWindow_tr2(const char* s, const char* c, long long* __out_len);
char* QRasterWindow_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QRasterWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QRasterWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QRasterWindow_metaObject_default(void* self);
int QRasterWindow_metric_default(void* self, int metric);
void* QRasterWindow_redirected_default(void* self, void* arg1);
void QRasterWindow_resizeEvent_default(void* self, void* event);
void QRasterWindow_exposeEvent_default(void* self, void* arg1);
void QRasterWindow_paintEvent_default(void* self, void* event);
int QRasterWindow_event_default(void* self, void* event);
int QRasterWindow_surfaceType_default(void* self);
void* QRasterWindow_format_default(void* self);
void* QRasterWindow_size_default(void* self);
void* QRasterWindow_accessibleRoot_default(void* self);
void* QRasterWindow_focusObject_default(void* self);
void QRasterWindow_moveEvent_default(void* self, void* arg1);
void QRasterWindow_focusInEvent_default(void* self, void* arg1);
void QRasterWindow_focusOutEvent_default(void* self, void* arg1);
void QRasterWindow_showEvent_default(void* self, void* arg1);
void QRasterWindow_hideEvent_default(void* self, void* arg1);
void QRasterWindow_closeEvent_default(void* self, void* arg1);
void QRasterWindow_keyPressEvent_default(void* self, void* arg1);
void QRasterWindow_keyReleaseEvent_default(void* self, void* arg1);
void QRasterWindow_mousePressEvent_default(void* self, void* arg1);
void QRasterWindow_mouseReleaseEvent_default(void* self, void* arg1);
void QRasterWindow_mouseDoubleClickEvent_default(void* self, void* arg1);
void QRasterWindow_mouseMoveEvent_default(void* self, void* arg1);
void QRasterWindow_wheelEvent_default(void* self, void* arg1);
void QRasterWindow_touchEvent_default(void* self, void* arg1);
void QRasterWindow_tabletEvent_default(void* self, void* arg1);
int QRasterWindow_eventFilter_default(void* self, void* watched, void* event);
void QRasterWindow_timerEvent_default(void* self, void* event);
void QRasterWindow_childEvent_default(void* self, void* event);
void QRasterWindow_customEvent_default(void* self, void* event);
void QRasterWindow_connectNotify_default(void* self, void* signal);
void QRasterWindow_disconnectNotify_default(void* self, void* signal);
int QRasterWindow_devType_default(void* self);
void QRasterWindow_initPainter_default(void* self, void* painter);
void* QRasterWindow_sharedPainter_default(void* self);
void QRasterWindow_connect_screenChanged(void* self, long long id);
void QRasterWindow_connect_modalityChanged(void* self, long long id);
void QRasterWindow_connect_flagsChanged(void* self, long long id);
void QRasterWindow_connect_windowStateChanged(void* self, long long id);
void QRasterWindow_connect_windowTitleChanged(void* self, long long id);
void QRasterWindow_connect_xChanged(void* self, long long id);
void QRasterWindow_connect_yChanged(void* self, long long id);
void QRasterWindow_connect_widthChanged(void* self, long long id);
void QRasterWindow_connect_heightChanged(void* self, long long id);
void QRasterWindow_connect_minimumWidthChanged(void* self, long long id);
void QRasterWindow_connect_minimumHeightChanged(void* self, long long id);
void QRasterWindow_connect_maximumWidthChanged(void* self, long long id);
void QRasterWindow_connect_maximumHeightChanged(void* self, long long id);
void QRasterWindow_connect_safeAreaMarginsChanged(void* self, long long id);
void QRasterWindow_connect_visibleChanged(void* self, long long id);
void QRasterWindow_connect_visibilityChanged(void* self, long long id);
void QRasterWindow_connect_activeChanged(void* self, long long id);
void QRasterWindow_connect_contentOrientationChanged(void* self, long long id);
void QRasterWindow_connect_focusObjectChanged(void* self, long long id);
void QRasterWindow_connect_opacityChanged(void* self, long long id);
void QRasterWindow_connect_transientParentChanged(void* self, long long id);
void QRasterWindow_connect_destroyed(void* self, long long id);
void QRasterWindow_connect_destroyed2(void* self, long long id);
void QRasterWindow_connect_objectNameChanged(void* self, long long id);
void QRasterWindow__disconnect(long long id);
void* QRasterWindow__Go_new(long long id);
void QRasterWindow_delete(void* self);

#ifdef __cplusplus
}
#endif
