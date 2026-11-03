#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFutureInterfaceBase_new();
void* QFutureInterfaceBase_new2(void* other);
void* QFutureInterfaceBase_new3(int initialState);
void* QFutureInterfaceBase_OperatorAssign(void* self, void* other);
void QFutureInterfaceBase_reportStarted(void* self);
void QFutureInterfaceBase_reportFinished(void* self);
void QFutureInterfaceBase_reportCanceled(void* self);
void QFutureInterfaceBase_reportException(void* self, void* e);
void QFutureInterfaceBase_reportResultsReady(void* self, int beginIndex, int endIndex);
void QFutureInterfaceBase_setRunnable(void* self, void* runnable);
void QFutureInterfaceBase_setThreadPool(void* self, void* pool);
void* QFutureInterfaceBase_threadPool(void* self);
void QFutureInterfaceBase_setFilterMode(void* self, int enable);
void QFutureInterfaceBase_setProgressRange(void* self, int minimum, int maximum);
int QFutureInterfaceBase_progressMinimum(void* self);
int QFutureInterfaceBase_progressMaximum(void* self);
int QFutureInterfaceBase_isProgressUpdateNeeded(void* self);
void QFutureInterfaceBase_setProgressValue(void* self, int progressValue);
int QFutureInterfaceBase_progressValue(void* self);
void QFutureInterfaceBase_setProgressValueAndText(void* self, int progressValue, const char* progressText, long long progressText_len);
char* QFutureInterfaceBase_progressText(void* self, long long* __out_len);
void QFutureInterfaceBase_setExpectedResultCount(void* self, int resultCount);
int QFutureInterfaceBase_expectedResultCount(void* self);
int QFutureInterfaceBase_resultCount(void* self);
int QFutureInterfaceBase_queryState(void* self, int state);
int QFutureInterfaceBase_isRunning(void* self);
int QFutureInterfaceBase_isStarted(void* self);
int QFutureInterfaceBase_isCanceled(void* self);
int QFutureInterfaceBase_isFinished(void* self);
int QFutureInterfaceBase_isPaused(void* self);
void QFutureInterfaceBase_setPaused(void* self, int paused);
void QFutureInterfaceBase_togglePaused(void* self);
int QFutureInterfaceBase_isSuspending(void* self);
int QFutureInterfaceBase_isSuspended(void* self);
int QFutureInterfaceBase_isThrottled(void* self);
int QFutureInterfaceBase_isResultReadyAt(void* self, int index);
int QFutureInterfaceBase_isValid(void* self);
int QFutureInterfaceBase_loadState(void* self);
void QFutureInterfaceBase_cancel(void* self);
void QFutureInterfaceBase_cancelAndFinish(void* self);
void QFutureInterfaceBase_cancelChain(void* self);
void QFutureInterfaceBase_setAddResultsIfCanceledEnabled(void* self, int enable);
int QFutureInterfaceBase_isAddResultsIfCanceledEnabled(void* self);
void QFutureInterfaceBase_setSuspended(void* self, int suspend);
void QFutureInterfaceBase_toggleSuspended(void* self);
void QFutureInterfaceBase_reportSuspended(void* self);
void QFutureInterfaceBase_setThrottled(void* self, int enable);
void QFutureInterfaceBase_waitForFinished(void* self);
int QFutureInterfaceBase_waitForNextResult(void* self);
void QFutureInterfaceBase_waitForResult(void* self, int resultIndex);
void QFutureInterfaceBase_waitForResume(void* self);
void QFutureInterfaceBase_suspendIfRequested(void* self);
void* QFutureInterfaceBase_mutex(void* self);
int QFutureInterfaceBase_hasException(void* self);
int QFutureInterfaceBase_OperatorEqual(void* self, void* other);
int QFutureInterfaceBase_OperatorNotEqual(void* self, void* other);
void QFutureInterfaceBase_swap(void* self, void* other);
int QFutureInterfaceBase_isChainCanceled(void* self);
void QFutureInterfaceBase_delete(void* self);

#ifdef __cplusplus
}
#endif
