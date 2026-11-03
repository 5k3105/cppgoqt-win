#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractAnimation_base(void* self);
void QAbstractAnimation__set_override(void* self, int idx);
void* QAbstractAnimation_fromQObject(void* o);
char* QAbstractAnimation_tr(const char* s, long long* __out_len);
int QAbstractAnimation_state(void* self);
void* QAbstractAnimation_group(void* self);
int QAbstractAnimation_direction(void* self);
void QAbstractAnimation_setDirection(void* self, int direction);
int QAbstractAnimation_currentTime(void* self);
int QAbstractAnimation_currentLoopTime(void* self);
int QAbstractAnimation_loopCount(void* self);
void QAbstractAnimation_setLoopCount(void* self, int loopCount);
int QAbstractAnimation_currentLoop(void* self);
int QAbstractAnimation_duration(void* self);
int QAbstractAnimation_totalDuration(void* self);
void QAbstractAnimation_finished(void* self);
void QAbstractAnimation_stateChanged(void* self, int newState, int oldState);
void QAbstractAnimation_currentLoopChanged(void* self, int currentLoop);
void QAbstractAnimation_directionChanged(void* self, int arg1);
void QAbstractAnimation_start(void* self);
void QAbstractAnimation_pause(void* self);
void QAbstractAnimation_resume(void* self);
void QAbstractAnimation_setPaused(void* self, int arg1);
void QAbstractAnimation_stop(void* self);
void QAbstractAnimation_setCurrentTime(void* self, int msecs);
void* QAbstractAnimation_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractAnimation_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAbstractAnimation_sender(void* self);
int QAbstractAnimation_senderSignalIndex(void* self);
int QAbstractAnimation_receivers(void* self, const char* signal);
int QAbstractAnimation_isSignalConnected(void* self, void* signal);
char* QAbstractAnimation_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractAnimation_tr3(const char* s, const char* c, int n, long long* __out_len);
void QAbstractAnimation_start2(void* self, int policy);
void* QAbstractAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAbstractAnimation_metaObject_default(void* self);
int QAbstractAnimation_event_default(void* self, void* event);
void QAbstractAnimation_updateState_default(void* self, int newState, int oldState);
void QAbstractAnimation_updateDirection_default(void* self, int direction);
int QAbstractAnimation_eventFilter_default(void* self, void* watched, void* event);
void QAbstractAnimation_timerEvent_default(void* self, void* event);
void QAbstractAnimation_childEvent_default(void* self, void* event);
void QAbstractAnimation_customEvent_default(void* self, void* event);
void QAbstractAnimation_connectNotify_default(void* self, void* signal);
void QAbstractAnimation_disconnectNotify_default(void* self, void* signal);
void QAbstractAnimation_connect_finished(void* self, long long id);
void QAbstractAnimation_connect_stateChanged(void* self, long long id);
void QAbstractAnimation_connect_currentLoopChanged(void* self, long long id);
void QAbstractAnimation_connect_directionChanged(void* self, long long id);
void QAbstractAnimation_connect_destroyed(void* self, long long id);
void QAbstractAnimation_connect_destroyed2(void* self, long long id);
void QAbstractAnimation_connect_objectNameChanged(void* self, long long id);
void QAbstractAnimation__disconnect(long long id);
void* QAbstractAnimation__Go_new(long long id, void* parent);
void QAbstractAnimation_delete(void* self);

#ifdef __cplusplus
}
#endif
