#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFile_base(void* self);
void QFile_chain(void* self, void** out);
void QFile__set_override(void* self, int idx);
void* QFile_fromQObject(void* o);
void* QFile_new(long long __goid);
void* QFile_new2(long long __goid, const char* name, long long name_len);
void* QFile_new3(long long __goid, void* parent);
void* QFile_new4(long long __goid, const char* name, long long name_len, void* parent);
char* QFile_tr(const char* s, long long* __out_len);
void QFile_setFileName(void* self, const char* name, long long name_len);
void* QFile_encodeName(const char* fileName, long long fileName_len);
char* QFile_decodeName(void* localFileName, long long* __out_len);
char* QFile_decodeName2(const char* localFileName, long long* __out_len);
int QFile_exists(void* self);
int QFile_exists2(const char* fileName, long long fileName_len);
char* QFile_symLinkTarget(void* self, long long* __out_len);
char* QFile_symLinkTarget2(const char* fileName, long long fileName_len, long long* __out_len);
int QFile_remove(void* self);
int QFile_remove2(const char* fileName, long long fileName_len);
int QFile_supportsMoveToTrash();
int QFile_moveToTrash(void* self);
int QFile_moveToTrash2(const char* fileName, long long fileName_len);
int QFile_rename(void* self, const char* newName, long long newName_len);
int QFile_rename2(const char* oldName, long long oldName_len, const char* newName, long long newName_len);
int QFile_link(void* self, const char* newName, long long newName_len);
int QFile_link2(const char* fileName, long long fileName_len, const char* newName, long long newName_len);
int QFile_copy(void* self, const char* newName, long long newName_len);
int QFile_copy2(const char* fileName, long long fileName_len, const char* newName, long long newName_len);
int QFile_open(void* self, int flags);
int QFile_open2(void* self, int flags, int permissions);
int QFile_open3(void* self, int fd, int ioFlags);
int QFile_resize(const char* filename, long long filename_len, long long sz);
int QFile_permissions(const char* filename, long long filename_len);
int QFile_setPermissions(const char* filename, long long filename_len, int permissionSpec);
void QFile_setOpenMode(void* self, int openMode);
void QFile_setErrorString(void* self, const char* errorString, long long errorString_len);
void* QFile_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QFile_connect2(void* sender, void* signal, void* receiver, void* method);
void* QFile_sender(void* self);
int QFile_senderSignalIndex(void* self);
int QFile_receivers(void* self, const char* signal);
int QFile_isSignalConnected(void* self, void* signal);
char* QFile_tr2(const char* s, const char* c, long long* __out_len);
char* QFile_tr3(const char* s, const char* c, int n, long long* __out_len);
int QFile_moveToTrash3(const char* fileName, long long fileName_len, void* pathInTrash);
int QFile_open4(void* self, int fd, int ioFlags, int handleFlags);
void* QFile_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QFile_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
int QFile_open_default(void* self, int flags);
int QFile_open_default2(void* self, int flags, int permissions);
int QFile_open_default3(void* self, int fd, int ioFlags);
int QFile_open_default4(void* self, int fd, int ioFlags, int handleFlags);
void* QFile_metaObject_default(void* self);
char* QFile_fileName_default(void* self, long long* __out_len);
long long QFile_size_default(void* self);
int QFile_resize_default(void* self, long long sz);
int QFile_permissions_default(void* self);
int QFile_setPermissions_default(void* self, int permissionSpec);
void QFile_close_default(void* self);
int QFile_isSequential_default(void* self);
long long QFile_pos_default(void* self);
int QFile_seek_default(void* self, long long offset);
int QFile_atEnd_default(void* self);
int QFile_reset_default(void* self);
long long QFile_bytesAvailable_default(void* self);
long long QFile_bytesToWrite_default(void* self);
int QFile_canReadLine_default(void* self);
int QFile_waitForReadyRead_default(void* self, int msecs);
int QFile_waitForBytesWritten_default(void* self, int msecs);
long long QFile_skipData_default(void* self, long long maxSize);
int QFile_event_default(void* self, void* event);
int QFile_eventFilter_default(void* self, void* watched, void* event);
void QFile_timerEvent_default(void* self, void* event);
void QFile_childEvent_default(void* self, void* event);
void QFile_customEvent_default(void* self, void* event);
void QFile_connectNotify_default(void* self, void* signal);
void QFile_disconnectNotify_default(void* self, void* signal);
void QFile_connect_readyRead(void* self, long long id);
void QFile_connect_channelReadyRead(void* self, long long id);
void QFile_connect_bytesWritten(void* self, long long id);
void QFile_connect_channelBytesWritten(void* self, long long id);
void QFile_connect_aboutToClose(void* self, long long id);
void QFile_connect_readChannelFinished(void* self, long long id);
void QFile_connect_destroyed(void* self, long long id);
void QFile_connect_destroyed2(void* self, long long id);
void QFile_connect_objectNameChanged(void* self, long long id);
void QFile__disconnect(long long id);
void* QFile__Go_new(long long id, void* parent);
void QFile_delete(void* self);

#ifdef __cplusplus
}
#endif
