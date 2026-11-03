#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMovie_base(void* self);
void QMovie__set_override(void* self, int idx);
void* QMovie_fromQObject(void* o);
void* QMovie_new(long long __goid);
void* QMovie_new2(long long __goid, void* device);
void* QMovie_new3(long long __goid, const char* fileName, long long fileName_len);
void* QMovie_new4(long long __goid, void* parent);
void* QMovie_new5(long long __goid, void* device, void* format);
void* QMovie_new6(long long __goid, void* device, void* format, void* parent);
void* QMovie_new7(long long __goid, const char* fileName, long long fileName_len, void* format);
void* QMovie_new8(long long __goid, const char* fileName, long long fileName_len, void* format, void* parent);
char* QMovie_tr(const char* s, long long* __out_len);
void** QMovie_supportedFormats(long long* __out_len);
void QMovie_setDevice(void* self, void* device);
void* QMovie_device(void* self);
void QMovie_setFileName(void* self, const char* fileName, long long fileName_len);
char* QMovie_fileName(void* self, long long* __out_len);
void QMovie_setFormat(void* self, void* format);
void* QMovie_format(void* self);
void QMovie_setBackgroundColor(void* self, void* color);
void* QMovie_backgroundColor(void* self);
int QMovie_state(void* self);
void* QMovie_frameRect(void* self);
void* QMovie_currentImage(void* self);
void* QMovie_currentPixmap(void* self);
int QMovie_isValid(void* self);
int QMovie_lastError(void* self);
char* QMovie_lastErrorString(void* self, long long* __out_len);
int QMovie_jumpToFrame(void* self, int frameNumber);
int QMovie_loopCount(void* self);
int QMovie_frameCount(void* self);
int QMovie_nextFrameDelay(void* self);
int QMovie_currentFrameNumber(void* self);
int QMovie_speed(void* self);
void* QMovie_scaledSize(void* self);
void QMovie_setScaledSize(void* self, void* size);
int QMovie_cacheMode(void* self);
void QMovie_setCacheMode(void* self, int mode);
void QMovie_started(void* self);
void QMovie_resized(void* self, void* size);
void QMovie_updated(void* self, void* rect);
void QMovie_stateChanged(void* self, int state);
void QMovie_error(void* self, int error_);
void QMovie_finished(void* self);
void QMovie_frameChanged(void* self, int frameNumber);
void QMovie_start(void* self);
int QMovie_jumpToNextFrame(void* self);
void QMovie_setPaused(void* self, int paused);
void QMovie_stop(void* self);
void QMovie_setSpeed(void* self, int percentSpeed);
void* QMovie_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QMovie_connect2(void* sender, void* signal, void* receiver, void* method);
void* QMovie_sender(void* self);
int QMovie_senderSignalIndex(void* self);
int QMovie_receivers(void* self, const char* signal);
int QMovie_isSignalConnected(void* self, void* signal);
char* QMovie_tr2(const char* s, const char* c, long long* __out_len);
char* QMovie_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QMovie_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QMovie_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QMovie_metaObject_default(void* self);
int QMovie_event_default(void* self, void* event);
int QMovie_eventFilter_default(void* self, void* watched, void* event);
void QMovie_timerEvent_default(void* self, void* event);
void QMovie_childEvent_default(void* self, void* event);
void QMovie_customEvent_default(void* self, void* event);
void QMovie_connectNotify_default(void* self, void* signal);
void QMovie_disconnectNotify_default(void* self, void* signal);
void QMovie_connect_started(void* self, long long id);
void QMovie_connect_resized(void* self, long long id);
void QMovie_connect_updated(void* self, long long id);
void QMovie_connect_stateChanged(void* self, long long id);
void QMovie_connect_error(void* self, long long id);
void QMovie_connect_finished(void* self, long long id);
void QMovie_connect_frameChanged(void* self, long long id);
void QMovie_connect_destroyed(void* self, long long id);
void QMovie_connect_destroyed2(void* self, long long id);
void QMovie_connect_objectNameChanged(void* self, long long id);
void QMovie__disconnect(long long id);
void* QMovie__Go_new(long long id, void* parent);
void QMovie_delete(void* self);

#ifdef __cplusplus
}
#endif
