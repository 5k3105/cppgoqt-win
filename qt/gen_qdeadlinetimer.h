#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDeadlineTimer_new();
void* QDeadlineTimer_new2(int type_);
void* QDeadlineTimer_new3(int arg1);
void* QDeadlineTimer_new4(long long msecs);
void* QDeadlineTimer_new5(void* arg1);
void* QDeadlineTimer_new6(int arg1, int type_);
void* QDeadlineTimer_new7(long long msecs, int type_);
void QDeadlineTimer_swap(void* self, void* other);
int QDeadlineTimer_isForever(void* self);
int QDeadlineTimer_hasExpired(void* self);
int QDeadlineTimer_timerType(void* self);
void QDeadlineTimer_setTimerType(void* self, int type_);
long long QDeadlineTimer_remainingTime(void* self);
long long QDeadlineTimer_remainingTimeNSecs(void* self);
void QDeadlineTimer_setRemainingTime(void* self, long long msecs);
void QDeadlineTimer_setPreciseRemainingTime(void* self, long long secs);
long long QDeadlineTimer_deadline(void* self);
long long QDeadlineTimer_deadlineNSecs(void* self);
void QDeadlineTimer_setDeadline(void* self, long long msecs);
void QDeadlineTimer_setPreciseDeadline(void* self, long long secs);
void* QDeadlineTimer_addNSecs(void* dt, long long nsecs);
void* QDeadlineTimer_current();
void* QDeadlineTimer_OperatorPlusAssign(void* self, long long msecs);
void* QDeadlineTimer_OperatorMinusAssign(void* self, long long msecs);
long long QDeadlineTimer_remainingTimeAsDuration(void* self);
void QDeadlineTimer_setRemainingTime2(void* self, long long msecs, int type_);
void QDeadlineTimer_setPreciseRemainingTime2(void* self, long long secs, long long nsecs);
void QDeadlineTimer_setPreciseRemainingTime3(void* self, long long secs, long long nsecs, int type_);
void QDeadlineTimer_setDeadline2(void* self, long long msecs, int timerType);
void QDeadlineTimer_setPreciseDeadline2(void* self, long long secs, long long nsecs);
void QDeadlineTimer_setPreciseDeadline3(void* self, long long secs, long long nsecs, int type_);
void* QDeadlineTimer_current2(int timerType);
void* QDeadlineTimer_OperatorPlus(void* self, long long msecs);
void* QDeadlineTimer_OperatorMinus(void* self, long long msecs);
long long QDeadlineTimer_OperatorMinus2(void* self, void* dt2);
int QDeadlineTimer_OperatorEqual(void* self, void* rhs);
int QDeadlineTimer_OperatorNotEqual(void* self, void* rhs);
int QDeadlineTimer_OperatorLesser(void* self, void* rhs);
int QDeadlineTimer_OperatorGreater(void* self, void* rhs);
int QDeadlineTimer_OperatorLesserOrEqual(void* self, void* rhs);
int QDeadlineTimer_OperatorGreaterOrEqual(void* self, void* rhs);
void QDeadlineTimer_delete(void* self);

#ifdef __cplusplus
}
#endif
