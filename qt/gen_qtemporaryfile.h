#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTemporaryFile_base(void* self);
void QTemporaryFile_chain(void* self, void** out);
void QTemporaryFile__set_override(void* self, int idx);
void* QTemporaryFile_fromQObject(void* o);
void* QTemporaryFile_new(long long __goid);
void* QTemporaryFile_new2(long long __goid, const char* templateName, long long templateName_len);
void* QTemporaryFile_new3(long long __goid, void* parent);
void* QTemporaryFile_new4(long long __goid, const char* templateName, long long templateName_len, void* parent);
char* QTemporaryFile_tr(const char* s, long long* __out_len);
int QTemporaryFile_autoRemove(void* self);
void QTemporaryFile_setAutoRemove(void* self, int b);
int QTemporaryFile_open(void* self);
char* QTemporaryFile_fileTemplate(void* self, long long* __out_len);
void QTemporaryFile_setFileTemplate(void* self, const char* name, long long name_len);
int QTemporaryFile_rename(void* self, const char* newName, long long newName_len);
int QTemporaryFile_renameOverwrite(void* self, const char* newName, long long newName_len);
void* QTemporaryFile_createNativeFile(const char* fileName, long long fileName_len);
void* QTemporaryFile_createNativeFile2(void* file);
void* QTemporaryFile_encodeName(const char* fileName, long long fileName_len);
char* QTemporaryFile_decodeName(void* localFileName, long long* __out_len);
char* QTemporaryFile_decodeName2(const char* localFileName, long long* __out_len);
int QTemporaryFile_exists(const char* fileName, long long fileName_len);
char* QTemporaryFile_symLinkTarget(const char* fileName, long long fileName_len, long long* __out_len);
int QTemporaryFile_remove(const char* fileName, long long fileName_len);
int QTemporaryFile_supportsMoveToTrash();
int QTemporaryFile_moveToTrash(const char* fileName, long long fileName_len);
int QTemporaryFile_link(const char* fileName, long long fileName_len, const char* newName, long long newName_len);
int QTemporaryFile_copy(const char* fileName, long long fileName_len, const char* newName, long long newName_len);
int QTemporaryFile_resize(const char* filename, long long filename_len, long long sz);
int QTemporaryFile_permissions(const char* filename, long long filename_len);
int QTemporaryFile_setPermissions(const char* filename, long long filename_len, int permissionSpec);
void QTemporaryFile_setOpenMode(void* self, int openMode);
void QTemporaryFile_setErrorString(void* self, const char* errorString, long long errorString_len);
void* QTemporaryFile_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTemporaryFile_connect2(void* sender, void* signal, void* receiver, void* method);
void* QTemporaryFile_sender(void* self);
int QTemporaryFile_senderSignalIndex(void* self);
int QTemporaryFile_receivers(void* self, const char* signal);
int QTemporaryFile_isSignalConnected(void* self, void* signal);
char* QTemporaryFile_tr2(const char* s, const char* c, long long* __out_len);
char* QTemporaryFile_tr3(const char* s, const char* c, int n, long long* __out_len);
int QTemporaryFile_moveToTrash2(const char* fileName, long long fileName_len, void* pathInTrash);
void* QTemporaryFile_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTemporaryFile_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
int QTemporaryFile_open_default(void* self);
void* QTemporaryFile_metaObject_default(void* self);
char* QTemporaryFile_fileName_default(void* self, long long* __out_len);
long long QTemporaryFile_size_default(void* self);
int QTemporaryFile_resize_default(void* self, long long sz);
int QTemporaryFile_permissions_default(void* self);
int QTemporaryFile_setPermissions_default(void* self, int permissionSpec);
void QTemporaryFile_close_default(void* self);
int QTemporaryFile_isSequential_default(void* self);
long long QTemporaryFile_pos_default(void* self);
int QTemporaryFile_seek_default(void* self, long long offset);
int QTemporaryFile_atEnd_default(void* self);
int QTemporaryFile_reset_default(void* self);
long long QTemporaryFile_bytesAvailable_default(void* self);
long long QTemporaryFile_bytesToWrite_default(void* self);
int QTemporaryFile_canReadLine_default(void* self);
int QTemporaryFile_waitForReadyRead_default(void* self, int msecs);
int QTemporaryFile_waitForBytesWritten_default(void* self, int msecs);
long long QTemporaryFile_skipData_default(void* self, long long maxSize);
int QTemporaryFile_event_default(void* self, void* event);
int QTemporaryFile_eventFilter_default(void* self, void* watched, void* event);
void QTemporaryFile_timerEvent_default(void* self, void* event);
void QTemporaryFile_childEvent_default(void* self, void* event);
void QTemporaryFile_customEvent_default(void* self, void* event);
void QTemporaryFile_connectNotify_default(void* self, void* signal);
void QTemporaryFile_disconnectNotify_default(void* self, void* signal);
void QTemporaryFile_connect_readyRead(void* self, long long id);
void QTemporaryFile_connect_channelReadyRead(void* self, long long id);
void QTemporaryFile_connect_bytesWritten(void* self, long long id);
void QTemporaryFile_connect_channelBytesWritten(void* self, long long id);
void QTemporaryFile_connect_aboutToClose(void* self, long long id);
void QTemporaryFile_connect_readChannelFinished(void* self, long long id);
void QTemporaryFile_connect_destroyed(void* self, long long id);
void QTemporaryFile_connect_destroyed2(void* self, long long id);
void QTemporaryFile_connect_objectNameChanged(void* self, long long id);
void QTemporaryFile__disconnect(long long id);
void* QTemporaryFile__Go_new(long long id, void* parent);
void QTemporaryFile_delete(void* self);

#ifdef __cplusplus
}
#endif
