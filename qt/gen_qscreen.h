#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScreen_base(void* self);
void* QScreen_fromQObject(void* o);
char* QScreen_tr(const char* s, long long* __out_len);
char* QScreen_name(void* self, long long* __out_len);
char* QScreen_manufacturer(void* self, long long* __out_len);
char* QScreen_model(void* self, long long* __out_len);
char* QScreen_serialNumber(void* self, long long* __out_len);
int QScreen_depth(void* self);
void* QScreen_size(void* self);
void* QScreen_geometry(void* self);
void* QScreen_physicalSize(void* self);
double QScreen_physicalDotsPerInchX(void* self);
double QScreen_physicalDotsPerInchY(void* self);
double QScreen_physicalDotsPerInch(void* self);
double QScreen_logicalDotsPerInchX(void* self);
double QScreen_logicalDotsPerInchY(void* self);
double QScreen_logicalDotsPerInch(void* self);
double QScreen_devicePixelRatio(void* self);
void* QScreen_availableSize(void* self);
void* QScreen_availableGeometry(void* self);
void** QScreen_virtualSiblings(void* self, long long* __out_len);
void* QScreen_virtualSiblingAt(void* self, void* point);
void* QScreen_virtualSize(void* self);
void* QScreen_virtualGeometry(void* self);
void* QScreen_availableVirtualSize(void* self);
void* QScreen_availableVirtualGeometry(void* self);
int QScreen_primaryOrientation(void* self);
int QScreen_orientation(void* self);
int QScreen_nativeOrientation(void* self);
int QScreen_angleBetween(void* self, int a, int b);
void* QScreen_transformBetween(void* self, int a, int b, void* target);
void* QScreen_mapBetween(void* self, int a, int b, void* rect);
int QScreen_isPortrait(void* self, int orientation);
int QScreen_isLandscape(void* self, int orientation);
void* QScreen_grabWindow(void* self);
double QScreen_refreshRate(void* self);
void QScreen_geometryChanged(void* self, void* geometry);
void QScreen_availableGeometryChanged(void* self, void* geometry);
void QScreen_physicalSizeChanged(void* self, void* size);
void QScreen_physicalDotsPerInchChanged(void* self, double dpi);
void QScreen_logicalDotsPerInchChanged(void* self, double dpi);
void QScreen_virtualGeometryChanged(void* self, void* rect);
void QScreen_primaryOrientationChanged(void* self, int orientation);
void QScreen_orientationChanged(void* self, int orientation);
void QScreen_refreshRateChanged(void* self, double refreshRate);
void* QScreen_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QScreen_connect2(void* sender, void* signal, void* receiver, void* method);
char* QScreen_tr2(const char* s, const char* c, long long* __out_len);
char* QScreen_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QScreen_grabWindow2(void* self, unsigned long long window);
void* QScreen_grabWindow3(void* self, unsigned long long window, int x);
void* QScreen_grabWindow4(void* self, unsigned long long window, int x, int y);
void* QScreen_grabWindow5(void* self, unsigned long long window, int x, int y, int w);
void* QScreen_grabWindow6(void* self, unsigned long long window, int x, int y, int w, int h);
void* QScreen_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QScreen_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QScreen_connect_geometryChanged(void* self, long long id);
void QScreen_connect_availableGeometryChanged(void* self, long long id);
void QScreen_connect_physicalSizeChanged(void* self, long long id);
void QScreen_connect_physicalDotsPerInchChanged(void* self, long long id);
void QScreen_connect_logicalDotsPerInchChanged(void* self, long long id);
void QScreen_connect_virtualGeometryChanged(void* self, long long id);
void QScreen_connect_primaryOrientationChanged(void* self, long long id);
void QScreen_connect_orientationChanged(void* self, long long id);
void QScreen_connect_refreshRateChanged(void* self, long long id);
void QScreen_connect_destroyed(void* self, long long id);
void QScreen_connect_destroyed2(void* self, long long id);
void QScreen_connect_objectNameChanged(void* self, long long id);
void QScreen__disconnect(long long id);
void QScreen_delete(void* self);

#ifdef __cplusplus
}
#endif
