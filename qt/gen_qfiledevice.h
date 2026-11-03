#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFileDevice_base(void* self);
void QFileDevice_chain(void* self, void** out);
void* QFileDevice_fromQObject(void* o);
char* QFileDevice_tr(const char* s, long long* __out_len);
int QFileDevice_error(void* self);
void QFileDevice_unsetError(void* self);
int QFileDevice_handle(void* self);
char* QFileDevice_fileName(void* self, long long* __out_len);
int QFileDevice_flush(void* self);
int QFileDevice_resize(void* self, long long sz);
int QFileDevice_permissions(void* self);
int QFileDevice_setPermissions(void* self, int permissionSpec);
int QFileDevice_unmap(void* self, unsigned char* address);
void* QFileDevice_fileTime(void* self, int time);
int QFileDevice_setFileTime(void* self, void* newDate, int fileTime);
void* QFileDevice_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QFileDevice_connect2(void* sender, void* signal, void* receiver, void* method);
char* QFileDevice_tr2(const char* s, const char* c, long long* __out_len);
char* QFileDevice_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QFileDevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QFileDevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QFileDevice_connect_readyRead(void* self, long long id);
void QFileDevice_connect_channelReadyRead(void* self, long long id);
void QFileDevice_connect_bytesWritten(void* self, long long id);
void QFileDevice_connect_channelBytesWritten(void* self, long long id);
void QFileDevice_connect_aboutToClose(void* self, long long id);
void QFileDevice_connect_readChannelFinished(void* self, long long id);
void QFileDevice_connect_destroyed(void* self, long long id);
void QFileDevice_connect_destroyed2(void* self, long long id);
void QFileDevice_connect_objectNameChanged(void* self, long long id);
void QFileDevice__disconnect(long long id);
void QFileDevice_delete(void* self);

#ifdef __cplusplus
}
#endif
