#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStylePlugin_base(void* self);
void QStylePlugin__set_override(void* self, int idx);
void* QStylePlugin_fromQObject(void* o);
char* QStylePlugin_tr(const char* s, long long* __out_len);
void* QStylePlugin_create(void* self, const char* key, long long key_len);
void* QStylePlugin_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QStylePlugin_connect2(void* sender, void* signal, void* receiver, void* method);
void* QStylePlugin_sender(void* self);
int QStylePlugin_senderSignalIndex(void* self);
int QStylePlugin_receivers(void* self, const char* signal);
int QStylePlugin_isSignalConnected(void* self, void* signal);
char* QStylePlugin_tr2(const char* s, const char* c, long long* __out_len);
char* QStylePlugin_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QStylePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QStylePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QStylePlugin_metaObject_default(void* self);
int QStylePlugin_event_default(void* self, void* event);
int QStylePlugin_eventFilter_default(void* self, void* watched, void* event);
void QStylePlugin_timerEvent_default(void* self, void* event);
void QStylePlugin_childEvent_default(void* self, void* event);
void QStylePlugin_customEvent_default(void* self, void* event);
void QStylePlugin_connectNotify_default(void* self, void* signal);
void QStylePlugin_disconnectNotify_default(void* self, void* signal);
void QStylePlugin_connect_destroyed(void* self, long long id);
void QStylePlugin_connect_destroyed2(void* self, long long id);
void QStylePlugin_connect_objectNameChanged(void* self, long long id);
void QStylePlugin__disconnect(long long id);
void* QStylePlugin__Go_new(long long id, void* parent);
void QStylePlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
