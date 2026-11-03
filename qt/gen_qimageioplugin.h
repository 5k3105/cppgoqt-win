#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QImageIOPlugin_base(void* self);
void QImageIOPlugin__set_override(void* self, int idx);
void* QImageIOPlugin_fromQObject(void* o);
char* QImageIOPlugin_tr(const char* s, long long* __out_len);
int QImageIOPlugin_capabilities(void* self, void* device, void* format);
void* QImageIOPlugin_create(void* self, void* device, void* format);
void* QImageIOPlugin_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QImageIOPlugin_connect2(void* sender, void* signal, void* receiver, void* method);
void* QImageIOPlugin_sender(void* self);
int QImageIOPlugin_senderSignalIndex(void* self);
int QImageIOPlugin_receivers(void* self, const char* signal);
int QImageIOPlugin_isSignalConnected(void* self, void* signal);
char* QImageIOPlugin_tr2(const char* s, const char* c, long long* __out_len);
char* QImageIOPlugin_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QImageIOPlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QImageIOPlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QImageIOPlugin_metaObject_default(void* self);
int QImageIOPlugin_event_default(void* self, void* event);
int QImageIOPlugin_eventFilter_default(void* self, void* watched, void* event);
void QImageIOPlugin_timerEvent_default(void* self, void* event);
void QImageIOPlugin_childEvent_default(void* self, void* event);
void QImageIOPlugin_customEvent_default(void* self, void* event);
void QImageIOPlugin_connectNotify_default(void* self, void* signal);
void QImageIOPlugin_disconnectNotify_default(void* self, void* signal);
void QImageIOPlugin_connect_destroyed(void* self, long long id);
void QImageIOPlugin_connect_destroyed2(void* self, long long id);
void QImageIOPlugin_connect_objectNameChanged(void* self, long long id);
void QImageIOPlugin__disconnect(long long id);
void* QImageIOPlugin__Go_new(long long id, void* parent);
void QImageIOPlugin_delete(void* self);

#ifdef __cplusplus
}
#endif
