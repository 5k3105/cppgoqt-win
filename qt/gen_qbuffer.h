#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBuffer_base(void* self);
void QBuffer_chain(void* self, void** out);
void QBuffer__set_override(void* self, int idx);
void* QBuffer_fromQObject(void* o);
void* QBuffer_new(long long __goid);
void* QBuffer_new2(long long __goid, void* buf);
void* QBuffer_new3(long long __goid, void* parent);
void* QBuffer_new4(long long __goid, void* buf, void* parent);
char* QBuffer_tr(const char* s, long long* __out_len);
void* QBuffer_buffer(void* self);
void* QBuffer_buffer2(void* self);
void QBuffer_setBuffer(void* self, void* a);
void QBuffer_setData(void* self, void* data);
void QBuffer_setData2(void* self, const char* data, long long len_);
void* QBuffer_data(void* self);
void QBuffer_setOpenMode(void* self, int openMode);
void QBuffer_setErrorString(void* self, const char* errorString, long long errorString_len);
void* QBuffer_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QBuffer_connect2(void* sender, void* signal, void* receiver, void* method);
void* QBuffer_sender(void* self);
int QBuffer_senderSignalIndex(void* self);
int QBuffer_receivers(void* self, const char* signal);
int QBuffer_isSignalConnected(void* self, void* signal);
char* QBuffer_tr2(const char* s, const char* c, long long* __out_len);
char* QBuffer_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QBuffer_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QBuffer_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QBuffer_metaObject_default(void* self);
int QBuffer_open_default(void* self, int openMode);
void QBuffer_close_default(void* self);
long long QBuffer_size_default(void* self);
long long QBuffer_pos_default(void* self);
int QBuffer_seek_default(void* self, long long off);
int QBuffer_atEnd_default(void* self);
int QBuffer_canReadLine_default(void* self);
void QBuffer_connectNotify_default(void* self, void* arg1);
void QBuffer_disconnectNotify_default(void* self, void* arg1);
int QBuffer_isSequential_default(void* self);
int QBuffer_reset_default(void* self);
long long QBuffer_bytesAvailable_default(void* self);
long long QBuffer_bytesToWrite_default(void* self);
int QBuffer_waitForReadyRead_default(void* self, int msecs);
int QBuffer_waitForBytesWritten_default(void* self, int msecs);
long long QBuffer_skipData_default(void* self, long long maxSize);
int QBuffer_event_default(void* self, void* event);
int QBuffer_eventFilter_default(void* self, void* watched, void* event);
void QBuffer_timerEvent_default(void* self, void* event);
void QBuffer_childEvent_default(void* self, void* event);
void QBuffer_customEvent_default(void* self, void* event);
void QBuffer_connect_readyRead(void* self, long long id);
void QBuffer_connect_channelReadyRead(void* self, long long id);
void QBuffer_connect_bytesWritten(void* self, long long id);
void QBuffer_connect_channelBytesWritten(void* self, long long id);
void QBuffer_connect_aboutToClose(void* self, long long id);
void QBuffer_connect_readChannelFinished(void* self, long long id);
void QBuffer_connect_destroyed(void* self, long long id);
void QBuffer_connect_destroyed2(void* self, long long id);
void QBuffer_connect_objectNameChanged(void* self, long long id);
void QBuffer__disconnect(long long id);
void* QBuffer__Go_new(long long id, void* parent);
void QBuffer_delete(void* self);

#ifdef __cplusplus
}
#endif
