#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPointingDevice_base(void* self);
void QPointingDevice_chain(void* self, void** out);
void QPointingDevice__set_override(void* self, int idx);
void* QPointingDevice_fromQObject(void* o);
void* QPointingDevice_new(long long __goid);
void* QPointingDevice_new2(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount);
void* QPointingDevice_new3(long long __goid, void* parent);
void* QPointingDevice_new4(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount, const char* seatName, long long seatName_len);
void* QPointingDevice_new5(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount, const char* seatName, long long seatName_len, void* uniqueId);
void* QPointingDevice_new6(long long __goid, const char* name, long long name_len, long long systemId, int devType, int pType, int caps, int maxPoints, int buttonCount, const char* seatName, long long seatName_len, void* uniqueId, void* parent);
char* QPointingDevice_tr(const char* s, long long* __out_len);
void QPointingDevice_setType(void* self, int devType);
void QPointingDevice_setCapabilities(void* self, int caps);
void QPointingDevice_setMaximumTouchPoints(void* self, int c);
int QPointingDevice_pointerType(void* self);
int QPointingDevice_maximumPoints(void* self);
int QPointingDevice_buttonCount(void* self);
void* QPointingDevice_uniqueId(void* self);
void* QPointingDevice_primaryPointingDevice();
int QPointingDevice_OperatorEqual(void* self, void* other);
void QPointingDevice_grabChanged(void* self, void* grabber, int transition, void* event, void* point);
char** QPointingDevice_seatNames(long long* __out_len);
void** QPointingDevice_devices(long long* __out_len);
void* QPointingDevice_primaryKeyboard();
void* QPointingDevice_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPointingDevice_connect2(void* sender, void* signal, void* receiver, void* method);
void* QPointingDevice_sender(void* self);
int QPointingDevice_senderSignalIndex(void* self);
int QPointingDevice_receivers(void* self, const char* signal);
int QPointingDevice_isSignalConnected(void* self, void* signal);
char* QPointingDevice_tr2(const char* s, const char* c, long long* __out_len);
char* QPointingDevice_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPointingDevice_primaryPointingDevice2(const char* seatName, long long seatName_len);
void* QPointingDevice_primaryKeyboard2(const char* seatName, long long seatName_len);
void* QPointingDevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPointingDevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPointingDevice_metaObject_default(void* self);
int QPointingDevice_event_default(void* self, void* event);
int QPointingDevice_eventFilter_default(void* self, void* watched, void* event);
void QPointingDevice_timerEvent_default(void* self, void* event);
void QPointingDevice_childEvent_default(void* self, void* event);
void QPointingDevice_customEvent_default(void* self, void* event);
void QPointingDevice_connectNotify_default(void* self, void* signal);
void QPointingDevice_disconnectNotify_default(void* self, void* signal);
void QPointingDevice_connect_grabChanged(void* self, long long id);
void QPointingDevice_connect_availableVirtualGeometryChanged(void* self, long long id);
void QPointingDevice_connect_capabilitiesChanged(void* self, long long id);
void QPointingDevice_connect_destroyed(void* self, long long id);
void QPointingDevice_connect_destroyed2(void* self, long long id);
void QPointingDevice_connect_objectNameChanged(void* self, long long id);
void QPointingDevice__disconnect(long long id);
void* QPointingDevice__Go_new(long long id, void* parent);
void QPointingDevice_delete(void* self);

#ifdef __cplusplus
}
#endif
