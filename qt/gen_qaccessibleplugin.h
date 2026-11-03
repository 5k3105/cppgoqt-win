#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessiblePlugin_base(void* self);
void QAccessiblePlugin__set_override(void* self, int idx);
void* QAccessiblePlugin_fromQObject(void* o);
char* QAccessiblePlugin_tr(const char* s, long long* __out_len);
void* QAccessiblePlugin_create(void* self, const char* key, long long key_len, void* object);
void* QAccessiblePlugin_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAccessiblePlugin_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAccessiblePlugin_sender(void* self);
int QAccessiblePlugin_senderSignalIndex(void* self);
int QAccessiblePlugin_receivers(void* self, const char* signal);
int QAccessiblePlugin_isSignalConnected(void* self, void* signal);
char* QAccessiblePlugin_tr2(const char* s, const char* c, long long* __out_len);
char* QAccessiblePlugin_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAccessiblePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAccessiblePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAccessiblePlugin_metaObject_default(void* self);
int QAccessiblePlugin_event_default(void* self, void* event);
int QAccessiblePlugin_eventFilter_default(void* self, void* watched, void* event);
void QAccessiblePlugin_timerEvent_default(void* self, void* event);
void QAccessiblePlugin_childEvent_default(void* self, void* event);
void QAccessiblePlugin_customEvent_default(void* self, void* event);
void QAccessiblePlugin_connectNotify_default(void* self, void* signal);
void QAccessiblePlugin_disconnectNotify_default(void* self, void* signal);
void QAccessiblePlugin_connect_destroyed(void* self, long long id);
void QAccessiblePlugin_connect_destroyed2(void* self, long long id);
void QAccessiblePlugin_connect_objectNameChanged(void* self, long long id);
void QAccessiblePlugin__disconnect(long long id);
void* QAccessiblePlugin__Go_new(long long id, void* parent);
void QAccessiblePlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
