#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyAnimation_base(void* self);
void QPropertyAnimation_chain(void* self, void** out);
void QPropertyAnimation__set_override(void* self, int idx);
void* QPropertyAnimation_fromQObject(void* o);
void* QPropertyAnimation_new(long long __goid);
void* QPropertyAnimation_new2(long long __goid, void* target, void* propertyName);
void* QPropertyAnimation_new3(long long __goid, void* parent);
void* QPropertyAnimation_new4(long long __goid, void* target, void* propertyName, void* parent);
char* QPropertyAnimation_tr(const char* s, long long* __out_len);
void* QPropertyAnimation_targetObject(void* self);
void QPropertyAnimation_setTargetObject(void* self, void* target);
void* QPropertyAnimation_propertyName(void* self);
void QPropertyAnimation_setPropertyName(void* self, void* propertyName);
void* QPropertyAnimation_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPropertyAnimation_connect2(void* sender, void* signal, void* receiver, void* method);
void* QPropertyAnimation_sender(void* self);
int QPropertyAnimation_senderSignalIndex(void* self);
int QPropertyAnimation_receivers(void* self, const char* signal);
int QPropertyAnimation_isSignalConnected(void* self, void* signal);
char* QPropertyAnimation_tr2(const char* s, const char* c, long long* __out_len);
char* QPropertyAnimation_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPropertyAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPropertyAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPropertyAnimation_metaObject_default(void* self);
int QPropertyAnimation_event_default(void* self, void* event);
void QPropertyAnimation_updateCurrentValue_default(void* self, void* value);
void QPropertyAnimation_updateState_default(void* self, int newState, int oldState);
int QPropertyAnimation_duration_default(void* self);
void QPropertyAnimation_updateCurrentTime_default(void* self, int arg1);
void* QPropertyAnimation_interpolated_default(void* self, void* from, void* to, double progress);
void QPropertyAnimation_updateDirection_default(void* self, int direction);
int QPropertyAnimation_eventFilter_default(void* self, void* watched, void* event);
void QPropertyAnimation_timerEvent_default(void* self, void* event);
void QPropertyAnimation_childEvent_default(void* self, void* event);
void QPropertyAnimation_customEvent_default(void* self, void* event);
void QPropertyAnimation_connectNotify_default(void* self, void* signal);
void QPropertyAnimation_disconnectNotify_default(void* self, void* signal);
void QPropertyAnimation_connect_valueChanged(void* self, long long id);
void QPropertyAnimation_connect_finished(void* self, long long id);
void QPropertyAnimation_connect_stateChanged(void* self, long long id);
void QPropertyAnimation_connect_currentLoopChanged(void* self, long long id);
void QPropertyAnimation_connect_directionChanged(void* self, long long id);
void QPropertyAnimation_connect_destroyed(void* self, long long id);
void QPropertyAnimation_connect_destroyed2(void* self, long long id);
void QPropertyAnimation_connect_objectNameChanged(void* self, long long id);
void QPropertyAnimation__disconnect(long long id);
void* QPropertyAnimation__Go_new(long long id, void* parent);
void QPropertyAnimation_delete(void* self);

#ifdef __cplusplus
}
#endif
