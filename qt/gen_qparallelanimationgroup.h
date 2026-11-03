#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QParallelAnimationGroup_base(void* self);
void QParallelAnimationGroup_chain(void* self, void** out);
void QParallelAnimationGroup__set_override(void* self, int idx);
void* QParallelAnimationGroup_fromQObject(void* o);
void* QParallelAnimationGroup_new(long long __goid);
void* QParallelAnimationGroup_new2(long long __goid, void* parent);
char* QParallelAnimationGroup_tr(const char* s, long long* __out_len);
void* QParallelAnimationGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QParallelAnimationGroup_connect2(void* sender, void* signal, void* receiver, void* method);
void* QParallelAnimationGroup_sender(void* self);
int QParallelAnimationGroup_senderSignalIndex(void* self);
int QParallelAnimationGroup_receivers(void* self, const char* signal);
int QParallelAnimationGroup_isSignalConnected(void* self, void* signal);
char* QParallelAnimationGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QParallelAnimationGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QParallelAnimationGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QParallelAnimationGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QParallelAnimationGroup_metaObject_default(void* self);
int QParallelAnimationGroup_duration_default(void* self);
int QParallelAnimationGroup_event_default(void* self, void* event);
void QParallelAnimationGroup_updateCurrentTime_default(void* self, int currentTime);
void QParallelAnimationGroup_updateState_default(void* self, int newState, int oldState);
void QParallelAnimationGroup_updateDirection_default(void* self, int direction);
int QParallelAnimationGroup_eventFilter_default(void* self, void* watched, void* event);
void QParallelAnimationGroup_timerEvent_default(void* self, void* event);
void QParallelAnimationGroup_childEvent_default(void* self, void* event);
void QParallelAnimationGroup_customEvent_default(void* self, void* event);
void QParallelAnimationGroup_connectNotify_default(void* self, void* signal);
void QParallelAnimationGroup_disconnectNotify_default(void* self, void* signal);
void QParallelAnimationGroup_connect_finished(void* self, long long id);
void QParallelAnimationGroup_connect_stateChanged(void* self, long long id);
void QParallelAnimationGroup_connect_currentLoopChanged(void* self, long long id);
void QParallelAnimationGroup_connect_directionChanged(void* self, long long id);
void QParallelAnimationGroup_connect_destroyed(void* self, long long id);
void QParallelAnimationGroup_connect_destroyed2(void* self, long long id);
void QParallelAnimationGroup_connect_objectNameChanged(void* self, long long id);
void QParallelAnimationGroup__disconnect(long long id);
void* QParallelAnimationGroup__Go_new(long long id, void* parent);
void QParallelAnimationGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
