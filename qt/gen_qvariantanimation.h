#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVariantAnimation_base(void* self);
void QVariantAnimation_chain(void* self, void** out);
void QVariantAnimation__set_override(void* self, int idx);
void* QVariantAnimation_fromQObject(void* o);
void* QVariantAnimation_new(long long __goid);
void* QVariantAnimation_new2(long long __goid, void* parent);
char* QVariantAnimation_tr(const char* s, long long* __out_len);
void* QVariantAnimation_startValue(void* self);
void QVariantAnimation_setStartValue(void* self, void* value);
void* QVariantAnimation_endValue(void* self);
void QVariantAnimation_setEndValue(void* self, void* value);
void* QVariantAnimation_keyValueAt(void* self, double step);
void QVariantAnimation_setKeyValueAt(void* self, double step, void* value);
void* QVariantAnimation_currentValue(void* self);
void QVariantAnimation_setDuration(void* self, int msecs);
void* QVariantAnimation_easingCurve(void* self);
void QVariantAnimation_setEasingCurve(void* self, void* easing);
void QVariantAnimation_valueChanged(void* self, void* value);
void* QVariantAnimation_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QVariantAnimation_connect2(void* sender, void* signal, void* receiver, void* method);
void* QVariantAnimation_sender(void* self);
int QVariantAnimation_senderSignalIndex(void* self);
int QVariantAnimation_receivers(void* self, const char* signal);
int QVariantAnimation_isSignalConnected(void* self, void* signal);
char* QVariantAnimation_tr2(const char* s, const char* c, long long* __out_len);
char* QVariantAnimation_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QVariantAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QVariantAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QVariantAnimation_metaObject_default(void* self);
int QVariantAnimation_duration_default(void* self);
int QVariantAnimation_event_default(void* self, void* event);
void QVariantAnimation_updateCurrentTime_default(void* self, int arg1);
void QVariantAnimation_updateState_default(void* self, int newState, int oldState);
void QVariantAnimation_updateCurrentValue_default(void* self, void* value);
void* QVariantAnimation_interpolated_default(void* self, void* from, void* to, double progress);
void QVariantAnimation_updateDirection_default(void* self, int direction);
int QVariantAnimation_eventFilter_default(void* self, void* watched, void* event);
void QVariantAnimation_timerEvent_default(void* self, void* event);
void QVariantAnimation_childEvent_default(void* self, void* event);
void QVariantAnimation_customEvent_default(void* self, void* event);
void QVariantAnimation_connectNotify_default(void* self, void* signal);
void QVariantAnimation_disconnectNotify_default(void* self, void* signal);
void QVariantAnimation_connect_valueChanged(void* self, long long id);
void QVariantAnimation_connect_finished(void* self, long long id);
void QVariantAnimation_connect_stateChanged(void* self, long long id);
void QVariantAnimation_connect_currentLoopChanged(void* self, long long id);
void QVariantAnimation_connect_directionChanged(void* self, long long id);
void QVariantAnimation_connect_destroyed(void* self, long long id);
void QVariantAnimation_connect_destroyed2(void* self, long long id);
void QVariantAnimation_connect_objectNameChanged(void* self, long long id);
void QVariantAnimation__disconnect(long long id);
void* QVariantAnimation__Go_new(long long id, void* parent);
void QVariantAnimation_delete(void* self);

#ifdef __cplusplus
}
#endif
