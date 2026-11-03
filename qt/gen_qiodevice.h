#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIODevice_base(void* self);
void* QIODevice_fromQObject(void* o);
char* QIODevice_tr(const char* s, long long* __out_len);
int QIODevice_openMode(void* self);
void QIODevice_setTextModeEnabled(void* self, int enabled);
int QIODevice_isTextModeEnabled(void* self);
int QIODevice_isOpen(void* self);
int QIODevice_isReadable(void* self);
int QIODevice_isWritable(void* self);
int QIODevice_isSequential(void* self);
int QIODevice_readChannelCount(void* self);
int QIODevice_writeChannelCount(void* self);
int QIODevice_currentReadChannel(void* self);
void QIODevice_setCurrentReadChannel(void* self, int channel);
int QIODevice_currentWriteChannel(void* self);
void QIODevice_setCurrentWriteChannel(void* self, int channel);
int QIODevice_open(void* self, int mode);
void QIODevice_close(void* self);
long long QIODevice_pos(void* self);
long long QIODevice_size(void* self);
int QIODevice_seek(void* self, long long pos);
int QIODevice_atEnd(void* self);
int QIODevice_reset(void* self);
long long QIODevice_bytesAvailable(void* self);
long long QIODevice_bytesToWrite(void* self);
long long QIODevice_read(void* self, char* data, long long maxlen);
void* QIODevice_read2(void* self, long long maxlen);
void* QIODevice_readAll(void* self);
long long QIODevice_readLine(void* self, char* data, long long maxlen);
void* QIODevice_readLine2(void* self);
int QIODevice_readLineInto(void* self, void* result);
int QIODevice_canReadLine(void* self);
void QIODevice_startTransaction(void* self);
void QIODevice_commitTransaction(void* self);
void QIODevice_rollbackTransaction(void* self);
int QIODevice_isTransactionStarted(void* self);
long long QIODevice_write(void* self, const char* data, long long len_);
long long QIODevice_write2(void* self, const char* data);
long long QIODevice_write3(void* self, void* data);
long long QIODevice_peek(void* self, char* data, long long maxlen);
void* QIODevice_peek2(void* self, long long maxlen);
long long QIODevice_skip(void* self, long long maxSize);
int QIODevice_waitForReadyRead(void* self, int msecs);
int QIODevice_waitForBytesWritten(void* self, int msecs);
void QIODevice_ungetChar(void* self, char c);
int QIODevice_putChar(void* self, char c);
int QIODevice_getChar(void* self, char* c);
char* QIODevice_errorString(void* self, long long* __out_len);
void QIODevice_readyRead(void* self);
void QIODevice_channelReadyRead(void* self, int channel);
void QIODevice_bytesWritten(void* self, long long bytes);
void QIODevice_channelBytesWritten(void* self, int channel, long long bytes);
void QIODevice_aboutToClose(void* self);
void QIODevice_readChannelFinished(void* self);
void* QIODevice_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QIODevice_connect2(void* sender, void* signal, void* receiver, void* method);
char* QIODevice_tr2(const char* s, const char* c, long long* __out_len);
char* QIODevice_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QIODevice_readLine3(void* self, long long maxlen);
int QIODevice_readLineInto2(void* self, void* result, long long maxlen);
void* QIODevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QIODevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QIODevice_asQIODeviceBase(void* self);
void QIODevice_connect_readyRead(void* self, long long id);
void QIODevice_connect_channelReadyRead(void* self, long long id);
void QIODevice_connect_bytesWritten(void* self, long long id);
void QIODevice_connect_channelBytesWritten(void* self, long long id);
void QIODevice_connect_aboutToClose(void* self, long long id);
void QIODevice_connect_readChannelFinished(void* self, long long id);
void QIODevice_connect_destroyed(void* self, long long id);
void QIODevice_connect_destroyed2(void* self, long long id);
void QIODevice_connect_objectNameChanged(void* self, long long id);
void QIODevice__disconnect(long long id);
void QIODevice_delete(void* self);

#ifdef __cplusplus
}
#endif
