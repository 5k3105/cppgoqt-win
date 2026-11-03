#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSizeGrip_base(void* self);
void QSizeGrip_chain(void* self, void** out);
void* QSizeGrip_fromQObject(void* o);
void* QSizeGrip_new(void* parent);
char* QSizeGrip_tr(const char* s, long long* __out_len);
void QSizeGrip_setTabOrder(void* arg1, void* arg2);
void* QSizeGrip_mouseGrabber();
void* QSizeGrip_keyboardGrabber();
void* QSizeGrip_find(unsigned long long arg1);
void* QSizeGrip_createWindowContainer(void* window);
void* QSizeGrip_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSizeGrip_connect2(void* sender, void* signal, void* receiver, void* method);
double QSizeGrip_devicePixelRatioFScale();
int QSizeGrip_encodeMetricF(int metric, double value);
char* QSizeGrip_tr2(const char* s, const char* c, long long* __out_len);
char* QSizeGrip_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSizeGrip_createWindowContainer2(void* window, void* parent);
void* QSizeGrip_createWindowContainer3(void* window, void* parent, int flags);
void* QSizeGrip_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSizeGrip_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QSizeGrip_connect_windowTitleChanged(void* self, long long id);
void QSizeGrip_connect_windowIconChanged(void* self, long long id);
void QSizeGrip_connect_windowIconTextChanged(void* self, long long id);
void QSizeGrip_connect_customContextMenuRequested(void* self, long long id);
void QSizeGrip_connect_destroyed(void* self, long long id);
void QSizeGrip_connect_destroyed2(void* self, long long id);
void QSizeGrip_connect_objectNameChanged(void* self, long long id);
void QSizeGrip__disconnect(long long id);
void QSizeGrip_delete(void* self);

#ifdef __cplusplus
}
#endif
