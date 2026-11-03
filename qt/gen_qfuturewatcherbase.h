#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFutureWatcherBase_base(void* self);
void* QFutureWatcherBase_fromQObject(void* o);
char* QFutureWatcherBase_tr(const char* s, long long* __out_len);
int QFutureWatcherBase_progressValue(void* self);
int QFutureWatcherBase_progressMinimum(void* self);
int QFutureWatcherBase_progressMaximum(void* self);
char* QFutureWatcherBase_progressText(void* self, long long* __out_len);
int QFutureWatcherBase_isStarted(void* self);
int QFutureWatcherBase_isFinished(void* self);
int QFutureWatcherBase_isRunning(void* self);
int QFutureWatcherBase_isCanceled(void* self);
int QFutureWatcherBase_isPaused(void* self);
int QFutureWatcherBase_isSuspending(void* self);
int QFutureWatcherBase_isSuspended(void* self);
void QFutureWatcherBase_waitForFinished(void* self);
void QFutureWatcherBase_setPendingResultsLimit(void* self, int limit);
void QFutureWatcherBase_started(void* self);
void QFutureWatcherBase_finished(void* self);
void QFutureWatcherBase_canceled(void* self);
void QFutureWatcherBase_paused(void* self);
void QFutureWatcherBase_suspending(void* self);
void QFutureWatcherBase_suspended(void* self);
void QFutureWatcherBase_resumed(void* self);
void QFutureWatcherBase_resultReadyAt(void* self, int resultIndex);
void QFutureWatcherBase_resultsReadyAt(void* self, int beginIndex, int endIndex);
void QFutureWatcherBase_progressRangeChanged(void* self, int minimum, int maximum);
void QFutureWatcherBase_progressValueChanged(void* self, int progressValue);
void QFutureWatcherBase_progressTextChanged(void* self, const char* progressText, long long progressText_len);
void QFutureWatcherBase_cancel(void* self);
void QFutureWatcherBase_setSuspended(void* self, int suspend);
void QFutureWatcherBase_suspend(void* self);
void QFutureWatcherBase_resume(void* self);
void QFutureWatcherBase_toggleSuspended(void* self);
void QFutureWatcherBase_setPaused(void* self, int paused);
void QFutureWatcherBase_pause(void* self);
void QFutureWatcherBase_togglePaused(void* self);
void* QFutureWatcherBase_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QFutureWatcherBase_connect2(void* sender, void* signal, void* receiver, void* method);
char* QFutureWatcherBase_tr2(const char* s, const char* c, long long* __out_len);
char* QFutureWatcherBase_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QFutureWatcherBase_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QFutureWatcherBase_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QFutureWatcherBase_connect_started(void* self, long long id);
void QFutureWatcherBase_connect_finished(void* self, long long id);
void QFutureWatcherBase_connect_canceled(void* self, long long id);
void QFutureWatcherBase_connect_paused(void* self, long long id);
void QFutureWatcherBase_connect_suspending(void* self, long long id);
void QFutureWatcherBase_connect_suspended(void* self, long long id);
void QFutureWatcherBase_connect_resumed(void* self, long long id);
void QFutureWatcherBase_connect_resultReadyAt(void* self, long long id);
void QFutureWatcherBase_connect_resultsReadyAt(void* self, long long id);
void QFutureWatcherBase_connect_progressRangeChanged(void* self, long long id);
void QFutureWatcherBase_connect_progressValueChanged(void* self, long long id);
void QFutureWatcherBase_connect_progressTextChanged(void* self, long long id);
void QFutureWatcherBase_connect_destroyed(void* self, long long id);
void QFutureWatcherBase_connect_destroyed2(void* self, long long id);
void QFutureWatcherBase_connect_objectNameChanged(void* self, long long id);
void QFutureWatcherBase__disconnect(long long id);
void QFutureWatcherBase_delete(void* self);

#ifdef __cplusplus
}
#endif
