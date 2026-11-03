#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPauseAnimation_base(void* self);
void QPauseAnimation_chain(void* self, void** out);
void QPauseAnimation__set_override(void* self, int idx);
void* QPauseAnimation_fromQObject(void* o);
void* QPauseAnimation_new(long long __goid);
void* QPauseAnimation_new2(long long __goid, int msecs);
void* QPauseAnimation_new3(long long __goid, void* parent);
void* QPauseAnimation_new4(long long __goid, int msecs, void* parent);
char* QPauseAnimation_tr(const char* s, long long* __out_len);
void QPauseAnimation_setDuration(void* self, int msecs);
void* QPauseAnimation_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPauseAnimation_connect2(void* sender, void* signal, void* receiver, void* method);
void* QPauseAnimation_sender(void* self);
int QPauseAnimation_senderSignalIndex(void* self);
int QPauseAnimation_receivers(void* self, const char* signal);
int QPauseAnimation_isSignalConnected(void* self, void* signal);
char* QPauseAnimation_tr2(const char* s, const char* c, long long* __out_len);
char* QPauseAnimation_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPauseAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPauseAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QPauseAnimation_metaObject_default(void* self);
int QPauseAnimation_duration_default(void* self);
int QPauseAnimation_event_default(void* self, void* e);
void QPauseAnimation_updateCurrentTime_default(void* self, int arg1);
void QPauseAnimation_updateState_default(void* self, int newState, int oldState);
void QPauseAnimation_updateDirection_default(void* self, int direction);
int QPauseAnimation_eventFilter_default(void* self, void* watched, void* event);
void QPauseAnimation_timerEvent_default(void* self, void* event);
void QPauseAnimation_childEvent_default(void* self, void* event);
void QPauseAnimation_customEvent_default(void* self, void* event);
void QPauseAnimation_connectNotify_default(void* self, void* signal);
void QPauseAnimation_disconnectNotify_default(void* self, void* signal);
void QPauseAnimation_connect_finished(void* self, long long id);
void QPauseAnimation_connect_stateChanged(void* self, long long id);
void QPauseAnimation_connect_currentLoopChanged(void* self, long long id);
void QPauseAnimation_connect_directionChanged(void* self, long long id);
void QPauseAnimation_connect_destroyed(void* self, long long id);
void QPauseAnimation_connect_destroyed2(void* self, long long id);
void QPauseAnimation_connect_objectNameChanged(void* self, long long id);
void QPauseAnimation__disconnect(long long id);
void* QPauseAnimation__Go_new(long long id, void* parent);
void QPauseAnimation_delete(void* self);

#ifdef __cplusplus
}
#endif
