#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QInputDevice_base(void* self);
void QInputDevice__set_override(void* self, int idx);
void* QInputDevice_fromQObject(void* o);
void* QInputDevice_new(long long __goid);
void* QInputDevice_new2(long long __goid, const char* name, long long name_len, long long systemId, int type_);
void* QInputDevice_new3(long long __goid, void* parent);
void* QInputDevice_new4(long long __goid, const char* name, long long name_len, long long systemId, int type_, const char* seatName, long long seatName_len);
void* QInputDevice_new5(long long __goid, const char* name, long long name_len, long long systemId, int type_, const char* seatName, long long seatName_len, void* parent);
char* QInputDevice_tr(const char* s, long long* __out_len);
char* QInputDevice_name(void* self, long long* __out_len);
int QInputDevice_type(void* self);
int QInputDevice_capabilities(void* self);
int QInputDevice_hasCapability(void* self, int cap_);
long long QInputDevice_systemId(void* self);
char* QInputDevice_seatName(void* self, long long* __out_len);
void* QInputDevice_availableVirtualGeometry(void* self);
char** QInputDevice_seatNames(long long* __out_len);
void** QInputDevice_devices(long long* __out_len);
void* QInputDevice_primaryKeyboard();
int QInputDevice_OperatorEqual(void* self, void* other);
void QInputDevice_availableVirtualGeometryChanged(void* self, void* area);
void QInputDevice_capabilitiesChanged(void* self, int capabilities);
void* QInputDevice_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QInputDevice_connect2(void* sender, void* signal, void* receiver, void* method);
void* QInputDevice_sender(void* self);
int QInputDevice_senderSignalIndex(void* self);
int QInputDevice_receivers(void* self, const char* signal);
int QInputDevice_isSignalConnected(void* self, void* signal);
char* QInputDevice_tr2(const char* s, const char* c, long long* __out_len);
char* QInputDevice_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QInputDevice_primaryKeyboard2(const char* seatName, long long seatName_len);
void* QInputDevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QInputDevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QInputDevice_metaObject_default(void* self);
int QInputDevice_event_default(void* self, void* event);
int QInputDevice_eventFilter_default(void* self, void* watched, void* event);
void QInputDevice_timerEvent_default(void* self, void* event);
void QInputDevice_childEvent_default(void* self, void* event);
void QInputDevice_customEvent_default(void* self, void* event);
void QInputDevice_connectNotify_default(void* self, void* signal);
void QInputDevice_disconnectNotify_default(void* self, void* signal);
void QInputDevice_connect_availableVirtualGeometryChanged(void* self, long long id);
void QInputDevice_connect_capabilitiesChanged(void* self, long long id);
void QInputDevice_connect_destroyed(void* self, long long id);
void QInputDevice_connect_destroyed2(void* self, long long id);
void QInputDevice_connect_objectNameChanged(void* self, long long id);
void QInputDevice__disconnect(long long id);
void* QInputDevice__Go_new(long long id, void* parent);
void QInputDevice_delete(void* self);

#ifdef __cplusplus
}
#endif
