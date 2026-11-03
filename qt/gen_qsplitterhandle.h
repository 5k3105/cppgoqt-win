#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSplitterHandle_base(void* self);
void QSplitterHandle_chain(void* self, void** out);
void* QSplitterHandle_fromQObject(void* o);
void* QSplitterHandle_new(int o, void* parent);
char* QSplitterHandle_tr(const char* s, long long* __out_len);
void QSplitterHandle_setOrientation(void* self, int o);
int QSplitterHandle_orientation(void* self);
int QSplitterHandle_opaqueResize(void* self);
void* QSplitterHandle_splitter(void* self);
void QSplitterHandle_setTabOrder(void* arg1, void* arg2);
void* QSplitterHandle_mouseGrabber();
void* QSplitterHandle_keyboardGrabber();
void* QSplitterHandle_find(unsigned long long arg1);
void* QSplitterHandle_createWindowContainer(void* window);
void* QSplitterHandle_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSplitterHandle_connect2(void* sender, void* signal, void* receiver, void* method);
double QSplitterHandle_devicePixelRatioFScale();
int QSplitterHandle_encodeMetricF(int metric, double value);
char* QSplitterHandle_tr2(const char* s, const char* c, long long* __out_len);
char* QSplitterHandle_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSplitterHandle_createWindowContainer2(void* window, void* parent);
void* QSplitterHandle_createWindowContainer3(void* window, void* parent, int flags);
void* QSplitterHandle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSplitterHandle_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QSplitterHandle_connect_windowTitleChanged(void* self, long long id);
void QSplitterHandle_connect_windowIconChanged(void* self, long long id);
void QSplitterHandle_connect_windowIconTextChanged(void* self, long long id);
void QSplitterHandle_connect_customContextMenuRequested(void* self, long long id);
void QSplitterHandle_connect_destroyed(void* self, long long id);
void QSplitterHandle_connect_destroyed2(void* self, long long id);
void QSplitterHandle_connect_objectNameChanged(void* self, long long id);
void QSplitterHandle__disconnect(long long id);
void QSplitterHandle_delete(void* self);

#ifdef __cplusplus
}
#endif
