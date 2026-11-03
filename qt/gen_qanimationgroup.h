#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAnimationGroup_base(void* self);
void QAnimationGroup_chain(void* self, void** out);
void QAnimationGroup__set_override(void* self, int idx);
void* QAnimationGroup_fromQObject(void* o);
char* QAnimationGroup_tr(const char* s, long long* __out_len);
void* QAnimationGroup_animationAt(void* self, int index);
int QAnimationGroup_animationCount(void* self);
int QAnimationGroup_indexOfAnimation(void* self, void* animation);
void QAnimationGroup_addAnimation(void* self, void* animation);
void QAnimationGroup_insertAnimation(void* self, int index, void* animation);
void QAnimationGroup_removeAnimation(void* self, void* animation);
void* QAnimationGroup_takeAnimation(void* self, int index);
void QAnimationGroup_clear(void* self);
void* QAnimationGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAnimationGroup_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAnimationGroup_sender(void* self);
int QAnimationGroup_senderSignalIndex(void* self);
int QAnimationGroup_receivers(void* self, const char* signal);
int QAnimationGroup_isSignalConnected(void* self, void* signal);
char* QAnimationGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QAnimationGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAnimationGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAnimationGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAnimationGroup_metaObject_default(void* self);
int QAnimationGroup_event_default(void* self, void* event);
void QAnimationGroup_updateState_default(void* self, int newState, int oldState);
void QAnimationGroup_updateDirection_default(void* self, int direction);
int QAnimationGroup_eventFilter_default(void* self, void* watched, void* event);
void QAnimationGroup_timerEvent_default(void* self, void* event);
void QAnimationGroup_childEvent_default(void* self, void* event);
void QAnimationGroup_customEvent_default(void* self, void* event);
void QAnimationGroup_connectNotify_default(void* self, void* signal);
void QAnimationGroup_disconnectNotify_default(void* self, void* signal);
void QAnimationGroup_connect_finished(void* self, long long id);
void QAnimationGroup_connect_stateChanged(void* self, long long id);
void QAnimationGroup_connect_currentLoopChanged(void* self, long long id);
void QAnimationGroup_connect_directionChanged(void* self, long long id);
void QAnimationGroup_connect_destroyed(void* self, long long id);
void QAnimationGroup_connect_destroyed2(void* self, long long id);
void QAnimationGroup_connect_objectNameChanged(void* self, long long id);
void QAnimationGroup__disconnect(long long id);
void* QAnimationGroup__Go_new(long long id, void* parent);
void QAnimationGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
