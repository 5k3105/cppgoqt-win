#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSocketNotifier_base(void* self);
void* QSocketNotifier_fromQObject(void* o);
void* QSocketNotifier_new(int arg1);
void* QSocketNotifier_new2(long long socket, int arg2);
void* QSocketNotifier_new3(int arg1, void* parent);
void* QSocketNotifier_new4(long long socket, int arg2, void* parent);
char* QSocketNotifier_tr(const char* s, long long* __out_len);
void QSocketNotifier_setSocket(void* self, long long socket);
long long QSocketNotifier_socket(void* self);
int QSocketNotifier_type(void* self);
int QSocketNotifier_isValid(void* self);
int QSocketNotifier_isEnabled(void* self);
void QSocketNotifier_setEnabled(void* self, int arg1);
void* QSocketNotifier_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSocketNotifier_connect2(void* sender, void* signal, void* receiver, void* method);
char* QSocketNotifier_tr2(const char* s, const char* c, long long* __out_len);
char* QSocketNotifier_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSocketNotifier_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSocketNotifier_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QSocketNotifier_connect_activated(void* self, long long id);
void QSocketNotifier_connect_destroyed(void* self, long long id);
void QSocketNotifier_connect_destroyed2(void* self, long long id);
void QSocketNotifier_connect_objectNameChanged(void* self, long long id);
void QSocketNotifier__disconnect(long long id);
void QSocketNotifier_delete(void* self);

#ifdef __cplusplus
}
#endif
