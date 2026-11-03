#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleBridgePlugin_base(void* self);
void QAccessibleBridgePlugin__set_override(void* self, int idx);
void* QAccessibleBridgePlugin_fromQObject(void* o);
char* QAccessibleBridgePlugin_tr(const char* s, long long* __out_len);
void* QAccessibleBridgePlugin_create(void* self, const char* key, long long key_len);
void* QAccessibleBridgePlugin_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAccessibleBridgePlugin_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAccessibleBridgePlugin_sender(void* self);
int QAccessibleBridgePlugin_senderSignalIndex(void* self);
int QAccessibleBridgePlugin_receivers(void* self, const char* signal);
int QAccessibleBridgePlugin_isSignalConnected(void* self, void* signal);
char* QAccessibleBridgePlugin_tr2(const char* s, const char* c, long long* __out_len);
char* QAccessibleBridgePlugin_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAccessibleBridgePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAccessibleBridgePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAccessibleBridgePlugin_metaObject_default(void* self);
int QAccessibleBridgePlugin_event_default(void* self, void* event);
int QAccessibleBridgePlugin_eventFilter_default(void* self, void* watched, void* event);
void QAccessibleBridgePlugin_timerEvent_default(void* self, void* event);
void QAccessibleBridgePlugin_childEvent_default(void* self, void* event);
void QAccessibleBridgePlugin_customEvent_default(void* self, void* event);
void QAccessibleBridgePlugin_connectNotify_default(void* self, void* signal);
void QAccessibleBridgePlugin_disconnectNotify_default(void* self, void* signal);
void QAccessibleBridgePlugin_connect_destroyed(void* self, long long id);
void QAccessibleBridgePlugin_connect_destroyed2(void* self, long long id);
void QAccessibleBridgePlugin_connect_objectNameChanged(void* self, long long id);
void QAccessibleBridgePlugin__disconnect(long long id);
void* QAccessibleBridgePlugin__Go_new(long long id, void* parent);
void QAccessibleBridgePlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
