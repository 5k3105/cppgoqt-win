#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPaintDeviceWindow_base(void* self);
void QPaintDeviceWindow_chain(void* self, void** out);
void* QPaintDeviceWindow_fromQObject(void* o);
char* QPaintDeviceWindow_tr(const char* s, long long* __out_len);
void QPaintDeviceWindow_update(void* self, void* rect);
void QPaintDeviceWindow_update2(void* self, void* region);
void QPaintDeviceWindow_update3(void* self);
void* QPaintDeviceWindow_fromWinId(unsigned long long id);
void* QPaintDeviceWindow_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPaintDeviceWindow_connect2(void* sender, void* signal, void* receiver, void* method);
int QPaintDeviceWindow_devType(void* self);
int QPaintDeviceWindow_paintingActive(void* self);
int QPaintDeviceWindow_widthMM(void* self);
int QPaintDeviceWindow_heightMM(void* self);
int QPaintDeviceWindow_logicalDpiX(void* self);
int QPaintDeviceWindow_logicalDpiY(void* self);
int QPaintDeviceWindow_physicalDpiX(void* self);
int QPaintDeviceWindow_physicalDpiY(void* self);
double QPaintDeviceWindow_devicePixelRatioF(void* self);
int QPaintDeviceWindow_colorCount(void* self);
int QPaintDeviceWindow_depth(void* self);
double QPaintDeviceWindow_devicePixelRatioFScale();
int QPaintDeviceWindow_encodeMetricF(int metric, double value);
char* QPaintDeviceWindow_tr2(const char* s, const char* c, long long* __out_len);
char* QPaintDeviceWindow_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPaintDeviceWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPaintDeviceWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPaintDeviceWindow_asQPaintDevice(void* self);
void QPaintDeviceWindow_connect_screenChanged(void* self, long long id);
void QPaintDeviceWindow_connect_modalityChanged(void* self, long long id);
void QPaintDeviceWindow_connect_flagsChanged(void* self, long long id);
void QPaintDeviceWindow_connect_windowStateChanged(void* self, long long id);
void QPaintDeviceWindow_connect_windowTitleChanged(void* self, long long id);
void QPaintDeviceWindow_connect_xChanged(void* self, long long id);
void QPaintDeviceWindow_connect_yChanged(void* self, long long id);
void QPaintDeviceWindow_connect_widthChanged(void* self, long long id);
void QPaintDeviceWindow_connect_heightChanged(void* self, long long id);
void QPaintDeviceWindow_connect_minimumWidthChanged(void* self, long long id);
void QPaintDeviceWindow_connect_minimumHeightChanged(void* self, long long id);
void QPaintDeviceWindow_connect_maximumWidthChanged(void* self, long long id);
void QPaintDeviceWindow_connect_maximumHeightChanged(void* self, long long id);
void QPaintDeviceWindow_connect_safeAreaMarginsChanged(void* self, long long id);
void QPaintDeviceWindow_connect_visibleChanged(void* self, long long id);
void QPaintDeviceWindow_connect_visibilityChanged(void* self, long long id);
void QPaintDeviceWindow_connect_activeChanged(void* self, long long id);
void QPaintDeviceWindow_connect_contentOrientationChanged(void* self, long long id);
void QPaintDeviceWindow_connect_focusObjectChanged(void* self, long long id);
void QPaintDeviceWindow_connect_opacityChanged(void* self, long long id);
void QPaintDeviceWindow_connect_transientParentChanged(void* self, long long id);
void QPaintDeviceWindow_connect_destroyed(void* self, long long id);
void QPaintDeviceWindow_connect_destroyed2(void* self, long long id);
void QPaintDeviceWindow_connect_objectNameChanged(void* self, long long id);
void QPaintDeviceWindow__disconnect(long long id);
void QPaintDeviceWindow_delete(void* self);

#ifdef __cplusplus
}
#endif
