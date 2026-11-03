#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSharedMemory_base(void* self);
void QSharedMemory__set_override(void* self, int idx);
void* QSharedMemory_fromQObject(void* o);
void* QSharedMemory_new(long long __goid);
void* QSharedMemory_new2(long long __goid, void* key);
void* QSharedMemory_new3(long long __goid, const char* key, long long key_len);
void* QSharedMemory_new4(long long __goid, void* parent);
void* QSharedMemory_new5(long long __goid, void* key, void* parent);
void* QSharedMemory_new6(long long __goid, const char* key, long long key_len, void* parent);
char* QSharedMemory_tr(const char* s, long long* __out_len);
void QSharedMemory_setKey(void* self, const char* key, long long key_len);
char* QSharedMemory_key(void* self, long long* __out_len);
void QSharedMemory_setNativeKey(void* self, void* key);
void QSharedMemory_setNativeKey2(void* self, const char* key, long long key_len);
char* QSharedMemory_nativeKey(void* self, long long* __out_len);
void* QSharedMemory_nativeIpcKey(void* self);
int QSharedMemory_create(void* self, long long size);
long long QSharedMemory_size(void* self);
int QSharedMemory_attach(void* self);
int QSharedMemory_isAttached(void* self);
int QSharedMemory_detach(void* self);
int QSharedMemory_lock(void* self);
int QSharedMemory_unlock(void* self);
int QSharedMemory_error(void* self);
char* QSharedMemory_errorString(void* self, long long* __out_len);
int QSharedMemory_isKeyTypeSupported(int type_);
void* QSharedMemory_platformSafeKey(const char* key, long long key_len);
void* QSharedMemory_legacyNativeKey(const char* key, long long key_len);
void* QSharedMemory_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSharedMemory_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSharedMemory_sender(void* self);
int QSharedMemory_senderSignalIndex(void* self);
int QSharedMemory_receivers(void* self, const char* signal);
int QSharedMemory_isSignalConnected(void* self, void* signal);
char* QSharedMemory_tr2(const char* s, const char* c, long long* __out_len);
char* QSharedMemory_tr3(const char* s, const char* c, int n, long long* __out_len);
void QSharedMemory_setNativeKey3(void* self, const char* key, long long key_len, int type_);
int QSharedMemory_create2(void* self, long long size, int mode);
int QSharedMemory_attach2(void* self, int mode);
void* QSharedMemory_platformSafeKey2(const char* key, long long key_len, int type_);
void* QSharedMemory_legacyNativeKey2(const char* key, long long key_len, int type_);
void* QSharedMemory_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSharedMemory_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSharedMemory_metaObject_default(void* self);
int QSharedMemory_event_default(void* self, void* event);
int QSharedMemory_eventFilter_default(void* self, void* watched, void* event);
void QSharedMemory_timerEvent_default(void* self, void* event);
void QSharedMemory_childEvent_default(void* self, void* event);
void QSharedMemory_customEvent_default(void* self, void* event);
void QSharedMemory_connectNotify_default(void* self, void* signal);
void QSharedMemory_disconnectNotify_default(void* self, void* signal);
void QSharedMemory_connect_destroyed(void* self, long long id);
void QSharedMemory_connect_destroyed2(void* self, long long id);
void QSharedMemory_connect_objectNameChanged(void* self, long long id);
void QSharedMemory__disconnect(long long id);
void* QSharedMemory__Go_new(long long id, void* parent);
void QSharedMemory_delete(void* self);

#ifdef __cplusplus
}
#endif
