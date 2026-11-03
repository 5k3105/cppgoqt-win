#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIconEnginePlugin_base(void* self);
void QIconEnginePlugin__set_override(void* self, int idx);
void* QIconEnginePlugin_fromQObject(void* o);
char* QIconEnginePlugin_tr(const char* s, long long* __out_len);
void* QIconEnginePlugin_create(void* self, const char* filename, long long filename_len);
void* QIconEnginePlugin_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QIconEnginePlugin_connect2(void* sender, void* signal, void* receiver, void* method);
void* QIconEnginePlugin_sender(void* self);
int QIconEnginePlugin_senderSignalIndex(void* self);
int QIconEnginePlugin_receivers(void* self, const char* signal);
int QIconEnginePlugin_isSignalConnected(void* self, void* signal);
char* QIconEnginePlugin_tr2(const char* s, const char* c, long long* __out_len);
char* QIconEnginePlugin_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QIconEnginePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QIconEnginePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QIconEnginePlugin_metaObject_default(void* self);
int QIconEnginePlugin_event_default(void* self, void* event);
int QIconEnginePlugin_eventFilter_default(void* self, void* watched, void* event);
void QIconEnginePlugin_timerEvent_default(void* self, void* event);
void QIconEnginePlugin_childEvent_default(void* self, void* event);
void QIconEnginePlugin_customEvent_default(void* self, void* event);
void QIconEnginePlugin_connectNotify_default(void* self, void* signal);
void QIconEnginePlugin_disconnectNotify_default(void* self, void* signal);
void QIconEnginePlugin_connect_destroyed(void* self, long long id);
void QIconEnginePlugin_connect_destroyed2(void* self, long long id);
void QIconEnginePlugin_connect_objectNameChanged(void* self, long long id);
void QIconEnginePlugin__disconnect(long long id);
void* QIconEnginePlugin__Go_new(long long id, void* parent);
void QIconEnginePlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
