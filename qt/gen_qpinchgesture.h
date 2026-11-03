#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPinchGesture_base(void* self);
void QPinchGesture_chain(void* self, void** out);
void QPinchGesture__set_override(void* self, int idx);
void* QPinchGesture_fromQObject(void* o);
void* QPinchGesture_new(long long __goid);
void* QPinchGesture_new2(long long __goid, void* parent);
char* QPinchGesture_tr(const char* s, long long* __out_len);
int QPinchGesture_totalChangeFlags(void* self);
void QPinchGesture_setTotalChangeFlags(void* self, int value);
int QPinchGesture_changeFlags(void* self);
void QPinchGesture_setChangeFlags(void* self, int value);
void* QPinchGesture_startCenterPoint(void* self);
void* QPinchGesture_lastCenterPoint(void* self);
void* QPinchGesture_centerPoint(void* self);
void QPinchGesture_setStartCenterPoint(void* self, void* value);
void QPinchGesture_setLastCenterPoint(void* self, void* value);
void QPinchGesture_setCenterPoint(void* self, void* value);
double QPinchGesture_totalScaleFactor(void* self);
double QPinchGesture_lastScaleFactor(void* self);
double QPinchGesture_scaleFactor(void* self);
void QPinchGesture_setTotalScaleFactor(void* self, double value);
void QPinchGesture_setLastScaleFactor(void* self, double value);
void QPinchGesture_setScaleFactor(void* self, double value);
double QPinchGesture_totalRotationAngle(void* self);
double QPinchGesture_lastRotationAngle(void* self);
double QPinchGesture_rotationAngle(void* self);
void QPinchGesture_setTotalRotationAngle(void* self, double value);
void QPinchGesture_setLastRotationAngle(void* self, double value);
void QPinchGesture_setRotationAngle(void* self, double value);
void* QPinchGesture_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPinchGesture_connect2(void* sender, void* signal, void* receiver, void* method);
void* QPinchGesture_sender(void* self);
int QPinchGesture_senderSignalIndex(void* self);
int QPinchGesture_receivers(void* self, const char* signal);
int QPinchGesture_isSignalConnected(void* self, void* signal);
char* QPinchGesture_tr2(const char* s, const char* c, long long* __out_len);
char* QPinchGesture_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPinchGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPinchGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPinchGesture_metaObject_default(void* self);
int QPinchGesture_event_default(void* self, void* event);
int QPinchGesture_eventFilter_default(void* self, void* watched, void* event);
void QPinchGesture_timerEvent_default(void* self, void* event);
void QPinchGesture_childEvent_default(void* self, void* event);
void QPinchGesture_customEvent_default(void* self, void* event);
void QPinchGesture_connectNotify_default(void* self, void* signal);
void QPinchGesture_disconnectNotify_default(void* self, void* signal);
void QPinchGesture_connect_destroyed(void* self, long long id);
void QPinchGesture_connect_destroyed2(void* self, long long id);
void QPinchGesture_connect_objectNameChanged(void* self, long long id);
void QPinchGesture__disconnect(long long id);
void* QPinchGesture__Go_new(long long id, void* parent);
void QPinchGesture_delete(void* self);

#ifdef __cplusplus
}
#endif
