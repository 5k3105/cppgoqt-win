#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFocusFrame_base(void* self);
void QFocusFrame_chain(void* self, void** out);
void QFocusFrame__set_override(void* self, int idx);
void* QFocusFrame_fromQObject(void* o);
void* QFocusFrame_new(long long __goid, void* parent);
void* QFocusFrame_new2(long long __goid);
char* QFocusFrame_tr(const char* s, long long* __out_len);
void QFocusFrame_setWidget(void* self, void* widget);
void* QFocusFrame_widget(void* self);
void QFocusFrame_setTabOrder(void* arg1, void* arg2);
void* QFocusFrame_mouseGrabber();
void* QFocusFrame_keyboardGrabber();
void* QFocusFrame_find(unsigned long long arg1);
void* QFocusFrame_createWindowContainer(void* window);
void QFocusFrame_updateMicroFocus(void* self);
void QFocusFrame_create(void* self);
void QFocusFrame_destroy(void* self);
int QFocusFrame_focusNextChild(void* self);
int QFocusFrame_focusPreviousChild(void* self);
void* QFocusFrame_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QFocusFrame_connect2(void* sender, void* signal, void* receiver, void* method);
void* QFocusFrame_sender(void* self);
int QFocusFrame_senderSignalIndex(void* self);
int QFocusFrame_receivers(void* self, const char* signal);
int QFocusFrame_isSignalConnected(void* self, void* signal);
double QFocusFrame_devicePixelRatioFScale();
int QFocusFrame_encodeMetricF(int metric, double value);
double QFocusFrame_getDecodedMetricF(void* self, int metricA, int metricB);
char* QFocusFrame_tr2(const char* s, const char* c, long long* __out_len);
char* QFocusFrame_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QFocusFrame_createWindowContainer2(void* window, void* parent);
void* QFocusFrame_createWindowContainer3(void* window, void* parent, int flags);
void QFocusFrame_updateMicroFocus2(void* self, int query);
void QFocusFrame_create2(void* self, unsigned long long arg1);
void QFocusFrame_create3(void* self, unsigned long long arg1, int initializeWindow);
void QFocusFrame_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow);
void QFocusFrame_destroy2(void* self, int destroyWindow);
void QFocusFrame_destroy3(void* self, int destroyWindow, int destroySubWindows);
void* QFocusFrame_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QFocusFrame_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QFocusFrame_metaObject_default(void* self);
int QFocusFrame_event_default(void* self, void* e);
int QFocusFrame_eventFilter_default(void* self, void* arg1, void* arg2);
void QFocusFrame_paintEvent_default(void* self, void* arg1);
void QFocusFrame_initStyleOption_default(void* self, void* option);
int QFocusFrame_devType_default(void* self);
void QFocusFrame_setVisible_default(void* self, int visible);
void* QFocusFrame_sizeHint_default(void* self);
void* QFocusFrame_minimumSizeHint_default(void* self);
int QFocusFrame_heightForWidth_default(void* self, int arg1);
int QFocusFrame_hasHeightForWidth_default(void* self);
void* QFocusFrame_paintEngine_default(void* self);
void QFocusFrame_mousePressEvent_default(void* self, void* event);
void QFocusFrame_mouseReleaseEvent_default(void* self, void* event);
void QFocusFrame_mouseDoubleClickEvent_default(void* self, void* event);
void QFocusFrame_mouseMoveEvent_default(void* self, void* event);
void QFocusFrame_wheelEvent_default(void* self, void* event);
void QFocusFrame_keyPressEvent_default(void* self, void* event);
void QFocusFrame_keyReleaseEvent_default(void* self, void* event);
void QFocusFrame_focusInEvent_default(void* self, void* event);
void QFocusFrame_focusOutEvent_default(void* self, void* event);
void QFocusFrame_enterEvent_default(void* self, void* event);
void QFocusFrame_leaveEvent_default(void* self, void* event);
void QFocusFrame_moveEvent_default(void* self, void* event);
void QFocusFrame_resizeEvent_default(void* self, void* event);
void QFocusFrame_closeEvent_default(void* self, void* event);
void QFocusFrame_contextMenuEvent_default(void* self, void* event);
void QFocusFrame_tabletEvent_default(void* self, void* event);
void QFocusFrame_actionEvent_default(void* self, void* event);
void QFocusFrame_dragEnterEvent_default(void* self, void* event);
void QFocusFrame_dragMoveEvent_default(void* self, void* event);
void QFocusFrame_dragLeaveEvent_default(void* self, void* event);
void QFocusFrame_dropEvent_default(void* self, void* event);
void QFocusFrame_showEvent_default(void* self, void* event);
void QFocusFrame_hideEvent_default(void* self, void* event);
void QFocusFrame_changeEvent_default(void* self, void* arg1);
int QFocusFrame_metric_default(void* self, int arg1);
void QFocusFrame_initPainter_default(void* self, void* painter);
void* QFocusFrame_redirected_default(void* self, void* offset);
void* QFocusFrame_sharedPainter_default(void* self);
void QFocusFrame_inputMethodEvent_default(void* self, void* arg1);
void* QFocusFrame_inputMethodQuery_default(void* self, int arg1);
int QFocusFrame_focusNextPrevChild_default(void* self, int next);
void QFocusFrame_timerEvent_default(void* self, void* event);
void QFocusFrame_childEvent_default(void* self, void* event);
void QFocusFrame_customEvent_default(void* self, void* event);
void QFocusFrame_connectNotify_default(void* self, void* signal);
void QFocusFrame_disconnectNotify_default(void* self, void* signal);
void QFocusFrame_connect_windowTitleChanged(void* self, long long id);
void QFocusFrame_connect_windowIconChanged(void* self, long long id);
void QFocusFrame_connect_windowIconTextChanged(void* self, long long id);
void QFocusFrame_connect_customContextMenuRequested(void* self, long long id);
void QFocusFrame_connect_destroyed(void* self, long long id);
void QFocusFrame_connect_destroyed2(void* self, long long id);
void QFocusFrame_connect_objectNameChanged(void* self, long long id);
void QFocusFrame__disconnect(long long id);
void* QFocusFrame__Go_new(long long id);
void QFocusFrame_delete(void* self);

#ifdef __cplusplus
}
#endif
