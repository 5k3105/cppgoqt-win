#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPluginLoader_base(void* self);
void QPluginLoader__set_override(void* self, int idx);
void* QPluginLoader_fromQObject(void* o);
void* QPluginLoader_new(long long __goid);
void* QPluginLoader_new2(long long __goid, const char* fileName, long long fileName_len);
void* QPluginLoader_new3(long long __goid, void* parent);
void* QPluginLoader_new4(long long __goid, const char* fileName, long long fileName_len, void* parent);
char* QPluginLoader_tr(const char* s, long long* __out_len);
void* QPluginLoader_instance(void* self);
void* QPluginLoader_metaData(void* self);
void** QPluginLoader_staticInstances(long long* __out_len);
void** QPluginLoader_staticPlugins(long long* __out_len);
int QPluginLoader_load(void* self);
int QPluginLoader_unload(void* self);
int QPluginLoader_isLoaded(void* self);
void QPluginLoader_setFileName(void* self, const char* fileName, long long fileName_len);
char* QPluginLoader_fileName(void* self, long long* __out_len);
char* QPluginLoader_errorString(void* self, long long* __out_len);
void QPluginLoader_setLoadHints(void* self, int loadHints);
int QPluginLoader_loadHints(void* self);
void* QPluginLoader_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPluginLoader_connect2(void* sender, void* signal, void* receiver, void* method);
void* QPluginLoader_sender(void* self);
int QPluginLoader_senderSignalIndex(void* self);
int QPluginLoader_receivers(void* self, const char* signal);
int QPluginLoader_isSignalConnected(void* self, void* signal);
char* QPluginLoader_tr2(const char* s, const char* c, long long* __out_len);
char* QPluginLoader_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPluginLoader_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPluginLoader_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPluginLoader_metaObject_default(void* self);
int QPluginLoader_event_default(void* self, void* event);
int QPluginLoader_eventFilter_default(void* self, void* watched, void* event);
void QPluginLoader_timerEvent_default(void* self, void* event);
void QPluginLoader_childEvent_default(void* self, void* event);
void QPluginLoader_customEvent_default(void* self, void* event);
void QPluginLoader_connectNotify_default(void* self, void* signal);
void QPluginLoader_disconnectNotify_default(void* self, void* signal);
void QPluginLoader_connect_destroyed(void* self, long long id);
void QPluginLoader_connect_destroyed2(void* self, long long id);
void QPluginLoader_connect_objectNameChanged(void* self, long long id);
void QPluginLoader__disconnect(long long id);
void* QPluginLoader__Go_new(long long id, void* parent);
void QPluginLoader_delete(void* self);

#ifdef __cplusplus
}
#endif
