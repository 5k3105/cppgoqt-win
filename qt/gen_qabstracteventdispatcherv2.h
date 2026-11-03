#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractEventDispatcherV2_base(void* self);
void QAbstractEventDispatcherV2_chain(void* self, void** out);
void* QAbstractEventDispatcherV2_fromQObject(void* o);
char* QAbstractEventDispatcherV2_tr(const char* s, long long* __out_len);
int QAbstractEventDispatcherV2_processEventsWithDeadline(void* self, int flags, void* deadline);
void* QAbstractEventDispatcherV2_instance();
void* QAbstractEventDispatcherV2_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractEventDispatcherV2_connect2(void* sender, void* signal, void* receiver, void* method);
char* QAbstractEventDispatcherV2_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractEventDispatcherV2_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAbstractEventDispatcherV2_instance2(void* thread);
void* QAbstractEventDispatcherV2_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractEventDispatcherV2_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QAbstractEventDispatcherV2_connect_aboutToBlock(void* self, long long id);
void QAbstractEventDispatcherV2_connect_awake(void* self, long long id);
void QAbstractEventDispatcherV2_connect_destroyed(void* self, long long id);
void QAbstractEventDispatcherV2_connect_destroyed2(void* self, long long id);
void QAbstractEventDispatcherV2_connect_objectNameChanged(void* self, long long id);
void QAbstractEventDispatcherV2__disconnect(long long id);
void QAbstractEventDispatcherV2_delete(void* self);

#ifdef __cplusplus
}
#endif
