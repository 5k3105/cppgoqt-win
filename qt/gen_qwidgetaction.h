#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWidgetAction_base(void* self);
void QWidgetAction_chain(void* self, void** out);
void QWidgetAction__set_override(void* self, int idx);
void* QWidgetAction_fromQObject(void* o);
void* QWidgetAction_new(long long __goid, void* parent);
char* QWidgetAction_tr(const char* s, long long* __out_len);
void QWidgetAction_setDefaultWidget(void* self, void* w);
void* QWidgetAction_defaultWidget(void* self);
void* QWidgetAction_requestWidget(void* self, void* parent);
void QWidgetAction_releaseWidget(void* self, void* widget);
void** QWidgetAction_createdWidgets(void* self, long long* __out_len);
void* QWidgetAction_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QWidgetAction_connect2(void* sender, void* signal, void* receiver, void* method);
void* QWidgetAction_sender(void* self);
int QWidgetAction_senderSignalIndex(void* self);
int QWidgetAction_receivers(void* self, const char* signal);
int QWidgetAction_isSignalConnected(void* self, void* signal);
char* QWidgetAction_tr2(const char* s, const char* c, long long* __out_len);
char* QWidgetAction_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QWidgetAction_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QWidgetAction_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QWidgetAction_metaObject_default(void* self);
int QWidgetAction_event_default(void* self, void* arg1);
int QWidgetAction_eventFilter_default(void* self, void* arg1, void* arg2);
void* QWidgetAction_createWidget_default(void* self, void* parent);
void QWidgetAction_deleteWidget_default(void* self, void* widget);
void QWidgetAction_timerEvent_default(void* self, void* event);
void QWidgetAction_childEvent_default(void* self, void* event);
void QWidgetAction_customEvent_default(void* self, void* event);
void QWidgetAction_connectNotify_default(void* self, void* signal);
void QWidgetAction_disconnectNotify_default(void* self, void* signal);
void QWidgetAction_connect_changed(void* self, long long id);
void QWidgetAction_connect_enabledChanged(void* self, long long id);
void QWidgetAction_connect_checkableChanged(void* self, long long id);
void QWidgetAction_connect_visibleChanged(void* self, long long id);
void QWidgetAction_connect_triggered(void* self, long long id);
void QWidgetAction_connect_triggered2(void* self, long long id);
void QWidgetAction_connect_hovered(void* self, long long id);
void QWidgetAction_connect_toggled(void* self, long long id);
void QWidgetAction_connect_destroyed(void* self, long long id);
void QWidgetAction_connect_destroyed2(void* self, long long id);
void QWidgetAction_connect_objectNameChanged(void* self, long long id);
void QWidgetAction__disconnect(long long id);
void* QWidgetAction__Go_new(long long id, void* parent);
void QWidgetAction_delete(void* self);

#ifdef __cplusplus
}
#endif
