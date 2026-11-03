#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSaveFile_base(void* self);
void QSaveFile_chain(void* self, void** out);
void QSaveFile__set_override(void* self, int idx);
void* QSaveFile_fromQObject(void* o);
void* QSaveFile_new(long long __goid);
void* QSaveFile_new2(long long __goid, const char* name, long long name_len);
void* QSaveFile_new3(long long __goid, void* parent);
void* QSaveFile_new4(long long __goid, const char* name, long long name_len, void* parent);
char* QSaveFile_tr(const char* s, long long* __out_len);
void QSaveFile_setFileName(void* self, const char* name, long long name_len);
int QSaveFile_commit(void* self);
void QSaveFile_cancelWriting(void* self);
void QSaveFile_setDirectWriteFallback(void* self, int enabled);
int QSaveFile_directWriteFallback(void* self);
void QSaveFile_setOpenMode(void* self, int openMode);
void QSaveFile_setErrorString(void* self, const char* errorString, long long errorString_len);
void* QSaveFile_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSaveFile_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSaveFile_sender(void* self);
int QSaveFile_senderSignalIndex(void* self);
int QSaveFile_receivers(void* self, const char* signal);
int QSaveFile_isSignalConnected(void* self, void* signal);
char* QSaveFile_tr2(const char* s, const char* c, long long* __out_len);
char* QSaveFile_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSaveFile_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSaveFile_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSaveFile_metaObject_default(void* self);
char* QSaveFile_fileName_default(void* self, long long* __out_len);
int QSaveFile_open_default(void* self, int flags);
int QSaveFile_isSequential_default(void* self);
long long QSaveFile_pos_default(void* self);
int QSaveFile_seek_default(void* self, long long offset);
int QSaveFile_atEnd_default(void* self);
long long QSaveFile_size_default(void* self);
int QSaveFile_resize_default(void* self, long long sz);
int QSaveFile_permissions_default(void* self);
int QSaveFile_setPermissions_default(void* self, int permissionSpec);
int QSaveFile_reset_default(void* self);
long long QSaveFile_bytesAvailable_default(void* self);
long long QSaveFile_bytesToWrite_default(void* self);
int QSaveFile_canReadLine_default(void* self);
int QSaveFile_waitForReadyRead_default(void* self, int msecs);
int QSaveFile_waitForBytesWritten_default(void* self, int msecs);
long long QSaveFile_skipData_default(void* self, long long maxSize);
int QSaveFile_event_default(void* self, void* event);
int QSaveFile_eventFilter_default(void* self, void* watched, void* event);
void QSaveFile_timerEvent_default(void* self, void* event);
void QSaveFile_childEvent_default(void* self, void* event);
void QSaveFile_customEvent_default(void* self, void* event);
void QSaveFile_connectNotify_default(void* self, void* signal);
void QSaveFile_disconnectNotify_default(void* self, void* signal);
void QSaveFile_connect_readyRead(void* self, long long id);
void QSaveFile_connect_channelReadyRead(void* self, long long id);
void QSaveFile_connect_bytesWritten(void* self, long long id);
void QSaveFile_connect_channelBytesWritten(void* self, long long id);
void QSaveFile_connect_aboutToClose(void* self, long long id);
void QSaveFile_connect_readChannelFinished(void* self, long long id);
void QSaveFile_connect_destroyed(void* self, long long id);
void QSaveFile_connect_destroyed2(void* self, long long id);
void QSaveFile_connect_objectNameChanged(void* self, long long id);
void QSaveFile__disconnect(long long id);
void* QSaveFile__Go_new(long long id, void* parent);
void QSaveFile_delete(void* self);

#ifdef __cplusplus
}
#endif
