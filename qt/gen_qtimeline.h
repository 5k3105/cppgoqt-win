#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTimeLine_base(void* self);
void QTimeLine__set_override(void* self, int idx);
void* QTimeLine_fromQObject(void* o);
void* QTimeLine_new(long long __goid);
void* QTimeLine_new2(long long __goid, int duration);
void* QTimeLine_new3(long long __goid, int duration, void* parent);
char* QTimeLine_tr(const char* s, long long* __out_len);
int QTimeLine_state(void* self);
int QTimeLine_loopCount(void* self);
void QTimeLine_setLoopCount(void* self, int count);
int QTimeLine_direction(void* self);
void QTimeLine_setDirection(void* self, int direction);
int QTimeLine_duration(void* self);
void QTimeLine_setDuration(void* self, int duration);
int QTimeLine_startFrame(void* self);
void QTimeLine_setStartFrame(void* self, int frame);
int QTimeLine_endFrame(void* self);
void QTimeLine_setEndFrame(void* self, int frame);
void QTimeLine_setFrameRange(void* self, int startFrame, int endFrame);
int QTimeLine_updateInterval(void* self);
void QTimeLine_setUpdateInterval(void* self, int interval);
void* QTimeLine_easingCurve(void* self);
void QTimeLine_setEasingCurve(void* self, void* curve);
int QTimeLine_currentTime(void* self);
int QTimeLine_currentFrame(void* self);
double QTimeLine_currentValue(void* self);
int QTimeLine_frameForTime(void* self, int msec);
double QTimeLine_valueForTime(void* self, int msec);
void QTimeLine_start(void* self);
void QTimeLine_resume(void* self);
void QTimeLine_stop(void* self);
void QTimeLine_setPaused(void* self, int paused);
void QTimeLine_setCurrentTime(void* self, int msec);
void QTimeLine_toggleDirection(void* self);
void* QTimeLine_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTimeLine_connect2(void* sender, void* signal, void* receiver, void* method);
void* QTimeLine_sender(void* self);
int QTimeLine_senderSignalIndex(void* self);
int QTimeLine_receivers(void* self, const char* signal);
int QTimeLine_isSignalConnected(void* self, void* signal);
char* QTimeLine_tr2(const char* s, const char* c, long long* __out_len);
char* QTimeLine_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTimeLine_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTimeLine_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
double QTimeLine_valueForTime_default(void* self, int msec);
void* QTimeLine_metaObject_default(void* self);
void QTimeLine_timerEvent_default(void* self, void* event);
int QTimeLine_event_default(void* self, void* event);
int QTimeLine_eventFilter_default(void* self, void* watched, void* event);
void QTimeLine_childEvent_default(void* self, void* event);
void QTimeLine_customEvent_default(void* self, void* event);
void QTimeLine_connectNotify_default(void* self, void* signal);
void QTimeLine_disconnectNotify_default(void* self, void* signal);
void QTimeLine_connect_valueChanged(void* self, long long id);
void QTimeLine_connect_frameChanged(void* self, long long id);
void QTimeLine_connect_stateChanged(void* self, long long id);
void QTimeLine_connect_finished(void* self, long long id);
void QTimeLine_connect_destroyed(void* self, long long id);
void QTimeLine_connect_destroyed2(void* self, long long id);
void QTimeLine_connect_objectNameChanged(void* self, long long id);
void QTimeLine__disconnect(long long id);
void* QTimeLine__Go_new(long long id);
void QTimeLine_delete(void* self);

#ifdef __cplusplus
}
#endif
