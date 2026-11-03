#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRubberBand_base(void* self);
void QRubberBand_chain(void* self, void** out);
void* QRubberBand_fromQObject(void* o);
void* QRubberBand_new(int arg1);
void* QRubberBand_new2(int arg1, void* arg2);
char* QRubberBand_tr(const char* s, long long* __out_len);
int QRubberBand_shape(void* self);
void QRubberBand_setGeometry(void* self, void* r);
void QRubberBand_setGeometry2(void* self, int x, int y, int w, int h);
void QRubberBand_move(void* self, int x, int y);
void QRubberBand_move2(void* self, void* p);
void QRubberBand_resize(void* self, int w, int h);
void QRubberBand_resize2(void* self, void* s);
void QRubberBand_setTabOrder(void* arg1, void* arg2);
void* QRubberBand_mouseGrabber();
void* QRubberBand_keyboardGrabber();
void* QRubberBand_find(unsigned long long arg1);
void* QRubberBand_createWindowContainer(void* window);
void* QRubberBand_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QRubberBand_connect2(void* sender, void* signal, void* receiver, void* method);
double QRubberBand_devicePixelRatioFScale();
int QRubberBand_encodeMetricF(int metric, double value);
char* QRubberBand_tr2(const char* s, const char* c, long long* __out_len);
char* QRubberBand_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QRubberBand_createWindowContainer2(void* window, void* parent);
void* QRubberBand_createWindowContainer3(void* window, void* parent, int flags);
void* QRubberBand_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QRubberBand_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QRubberBand_connect_windowTitleChanged(void* self, long long id);
void QRubberBand_connect_windowIconChanged(void* self, long long id);
void QRubberBand_connect_windowIconTextChanged(void* self, long long id);
void QRubberBand_connect_customContextMenuRequested(void* self, long long id);
void QRubberBand_connect_destroyed(void* self, long long id);
void QRubberBand_connect_destroyed2(void* self, long long id);
void QRubberBand_connect_objectNameChanged(void* self, long long id);
void QRubberBand__disconnect(long long id);
void QRubberBand_delete(void* self);

#ifdef __cplusplus
}
#endif
