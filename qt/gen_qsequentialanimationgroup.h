#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSequentialAnimationGroup_base(void* self);
void QSequentialAnimationGroup_chain(void* self, void** out);
void QSequentialAnimationGroup__set_override(void* self, int idx);
void* QSequentialAnimationGroup_fromQObject(void* o);
void* QSequentialAnimationGroup_new(long long __goid);
void* QSequentialAnimationGroup_new2(long long __goid, void* parent);
char* QSequentialAnimationGroup_tr(const char* s, long long* __out_len);
void* QSequentialAnimationGroup_addPause(void* self, int msecs);
void* QSequentialAnimationGroup_insertPause(void* self, int index, int msecs);
void* QSequentialAnimationGroup_currentAnimation(void* self);
void QSequentialAnimationGroup_currentAnimationChanged(void* self, void* current);
void* QSequentialAnimationGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSequentialAnimationGroup_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSequentialAnimationGroup_sender(void* self);
int QSequentialAnimationGroup_senderSignalIndex(void* self);
int QSequentialAnimationGroup_receivers(void* self, const char* signal);
int QSequentialAnimationGroup_isSignalConnected(void* self, void* signal);
char* QSequentialAnimationGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QSequentialAnimationGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSequentialAnimationGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSequentialAnimationGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSequentialAnimationGroup_metaObject_default(void* self);
int QSequentialAnimationGroup_duration_default(void* self);
int QSequentialAnimationGroup_event_default(void* self, void* event);
void QSequentialAnimationGroup_updateCurrentTime_default(void* self, int arg1);
void QSequentialAnimationGroup_updateState_default(void* self, int newState, int oldState);
void QSequentialAnimationGroup_updateDirection_default(void* self, int direction);
int QSequentialAnimationGroup_eventFilter_default(void* self, void* watched, void* event);
void QSequentialAnimationGroup_timerEvent_default(void* self, void* event);
void QSequentialAnimationGroup_childEvent_default(void* self, void* event);
void QSequentialAnimationGroup_customEvent_default(void* self, void* event);
void QSequentialAnimationGroup_connectNotify_default(void* self, void* signal);
void QSequentialAnimationGroup_disconnectNotify_default(void* self, void* signal);
void QSequentialAnimationGroup_connect_currentAnimationChanged(void* self, long long id);
void QSequentialAnimationGroup_connect_finished(void* self, long long id);
void QSequentialAnimationGroup_connect_stateChanged(void* self, long long id);
void QSequentialAnimationGroup_connect_currentLoopChanged(void* self, long long id);
void QSequentialAnimationGroup_connect_directionChanged(void* self, long long id);
void QSequentialAnimationGroup_connect_destroyed(void* self, long long id);
void QSequentialAnimationGroup_connect_destroyed2(void* self, long long id);
void QSequentialAnimationGroup_connect_objectNameChanged(void* self, long long id);
void QSequentialAnimationGroup__disconnect(long long id);
void* QSequentialAnimationGroup__Go_new(long long id, void* parent);
void QSequentialAnimationGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
