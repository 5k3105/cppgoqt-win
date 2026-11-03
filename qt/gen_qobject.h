#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QObject__set_override(void* self, int idx);
void* QObject_new(long long __goid);
void* QObject_new2(long long __goid, void* parent);
void* QObject_metaObject(void* self);
char* QObject_tr(const char* s, long long* __out_len);
int QObject_event(void* self, void* event);
int QObject_eventFilter(void* self, void* watched, void* event);
char* QObject_objectName(void* self, long long* __out_len);
void QObject_setObjectName(void* self, const char* name, long long name_len);
int QObject_isWidgetType(void* self);
int QObject_isWindowType(void* self);
int QObject_isQuickItemType(void* self);
int QObject_isQmlExposed(void* self);
int QObject_signalsBlocked(void* self);
int QObject_blockSignals(void* self, int b);
void* QObject_thread(void* self);
int QObject_moveToThread(void* self, void* thread);
int QObject_startTimer(void* self, int interval);
int QObject_startTimer2(void* self, long long time);
void QObject_killTimer(void* self, int id);
void QObject_killTimer2(void* self, int id);
void** QObject_children(void* self, long long* __out_len);
void QObject_setParent(void* self, void* parent);
void QObject_installEventFilter(void* self, void* filterObj);
void QObject_removeEventFilter(void* self, void* obj);
void* QObject_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QObject_connect2(void* sender, void* signal, void* receiver, void* method);
void* QObject_connect3(void* self, void* sender, const char* signal, const char* member);
void QObject_dumpObjectTree(void* self);
void QObject_dumpObjectInfo(void* self);
int QObject_setProperty(void* self, const char* name, void* value);
void* QObject_property(void* self, const char* name);
void** QObject_dynamicPropertyNames(void* self, long long* __out_len);
void* QObject_bindingStorage(void* self);
void* QObject_bindingStorage2(void* self);
void QObject_destroyed(void* self);
void* QObject_parent(void* self);
int QObject_inherits(void* self, const char* classname);
void QObject_deleteLater(void* self);
void* QObject_sender(void* self);
int QObject_senderSignalIndex(void* self);
int QObject_receivers(void* self, const char* signal);
int QObject_isSignalConnected(void* self, void* signal);
char* QObject_tr2(const char* s, const char* c, long long* __out_len);
char* QObject_tr3(const char* s, const char* c, int n, long long* __out_len);
int QObject_startTimer3(void* self, int interval, int timerType);
int QObject_startTimer4(void* self, long long time, int timerType);
void* QObject_connect4(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QObject_connect5(void* sender, void* signal, void* receiver, void* method, int type_);
void* QObject_connect6(void* self, void* sender, const char* signal, const char* member, int type_);
void QObject_destroyed2(void* self, void* arg1);
void* QObject_metaObject_default(void* self);
int QObject_event_default(void* self, void* event);
int QObject_eventFilter_default(void* self, void* watched, void* event);
void QObject_timerEvent_default(void* self, void* event);
void QObject_childEvent_default(void* self, void* event);
void QObject_customEvent_default(void* self, void* event);
void QObject_connectNotify_default(void* self, void* signal);
void QObject_disconnectNotify_default(void* self, void* signal);
void QObject_connect_destroyed(void* self, long long id);
void QObject_connect_destroyed2(void* self, long long id);
void QObject_connect_objectNameChanged(void* self, long long id);
void QObject__disconnect(long long id);
void* QObject__Go_new(long long id, void* parent);
void QObject_delete(void* self);

#ifdef __cplusplus
}
#endif
