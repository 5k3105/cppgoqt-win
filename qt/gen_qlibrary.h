#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLibrary_base(void* self);
void QLibrary__set_override(void* self, int idx);
void* QLibrary_fromQObject(void* o);
void* QLibrary_new(long long __goid);
void* QLibrary_new2(long long __goid, const char* fileName, long long fileName_len);
void* QLibrary_new3(long long __goid, const char* fileName, long long fileName_len, int verNum);
void* QLibrary_new4(long long __goid, const char* fileName, long long fileName_len, const char* version, long long version_len);
void* QLibrary_new5(long long __goid, void* parent);
void* QLibrary_new6(long long __goid, const char* fileName, long long fileName_len, void* parent);
void* QLibrary_new7(long long __goid, const char* fileName, long long fileName_len, int verNum, void* parent);
void* QLibrary_new8(long long __goid, const char* fileName, long long fileName_len, const char* version, long long version_len, void* parent);
char* QLibrary_tr(const char* s, long long* __out_len);
int QLibrary_load(void* self);
int QLibrary_unload(void* self);
int QLibrary_isLoaded(void* self);
int QLibrary_isLibrary(const char* fileName, long long fileName_len);
void QLibrary_setFileName(void* self, const char* fileName, long long fileName_len);
char* QLibrary_fileName(void* self, long long* __out_len);
void QLibrary_setFileNameAndVersion(void* self, const char* fileName, long long fileName_len, int verNum);
void QLibrary_setFileNameAndVersion2(void* self, const char* fileName, long long fileName_len, const char* version, long long version_len);
char* QLibrary_errorString(void* self, long long* __out_len);
void QLibrary_setLoadHints(void* self, int hints);
int QLibrary_loadHints(void* self);
void* QLibrary_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QLibrary_connect2(void* sender, void* signal, void* receiver, void* method);
void* QLibrary_sender(void* self);
int QLibrary_senderSignalIndex(void* self);
int QLibrary_receivers(void* self, const char* signal);
int QLibrary_isSignalConnected(void* self, void* signal);
char* QLibrary_tr2(const char* s, const char* c, long long* __out_len);
char* QLibrary_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QLibrary_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QLibrary_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QLibrary_metaObject_default(void* self);
int QLibrary_event_default(void* self, void* event);
int QLibrary_eventFilter_default(void* self, void* watched, void* event);
void QLibrary_timerEvent_default(void* self, void* event);
void QLibrary_childEvent_default(void* self, void* event);
void QLibrary_customEvent_default(void* self, void* event);
void QLibrary_connectNotify_default(void* self, void* signal);
void QLibrary_disconnectNotify_default(void* self, void* signal);
void QLibrary_connect_destroyed(void* self, long long id);
void QLibrary_connect_destroyed2(void* self, long long id);
void QLibrary_connect_objectNameChanged(void* self, long long id);
void QLibrary__disconnect(long long id);
void* QLibrary__Go_new(long long id, void* parent);
void QLibrary_delete(void* self);

#ifdef __cplusplus
}
#endif
