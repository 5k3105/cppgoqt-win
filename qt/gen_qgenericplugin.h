#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGenericPlugin_base(void* self);
void QGenericPlugin__set_override(void* self, int idx);
void* QGenericPlugin_fromQObject(void* o);
char* QGenericPlugin_tr(const char* s, long long* __out_len);
void* QGenericPlugin_create(void* self, const char* name, long long name_len, const char* spec, long long spec_len);
void* QGenericPlugin_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGenericPlugin_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGenericPlugin_sender(void* self);
int QGenericPlugin_senderSignalIndex(void* self);
int QGenericPlugin_receivers(void* self, const char* signal);
int QGenericPlugin_isSignalConnected(void* self, void* signal);
char* QGenericPlugin_tr2(const char* s, const char* c, long long* __out_len);
char* QGenericPlugin_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGenericPlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGenericPlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QGenericPlugin_metaObject_default(void* self);
int QGenericPlugin_event_default(void* self, void* event);
int QGenericPlugin_eventFilter_default(void* self, void* watched, void* event);
void QGenericPlugin_timerEvent_default(void* self, void* event);
void QGenericPlugin_childEvent_default(void* self, void* event);
void QGenericPlugin_customEvent_default(void* self, void* event);
void QGenericPlugin_connectNotify_default(void* self, void* signal);
void QGenericPlugin_disconnectNotify_default(void* self, void* signal);
void QGenericPlugin_connect_destroyed(void* self, long long id);
void QGenericPlugin_connect_destroyed2(void* self, long long id);
void QGenericPlugin_connect_objectNameChanged(void* self, long long id);
void QGenericPlugin__disconnect(long long id);
void* QGenericPlugin__Go_new(long long id, void* parent);
void QGenericPlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
