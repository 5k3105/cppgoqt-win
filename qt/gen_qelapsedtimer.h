#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QElapsedTimer_new();
int QElapsedTimer_clockType();
int QElapsedTimer_isMonotonic();
void QElapsedTimer_start(void* self);
long long QElapsedTimer_restart(void* self);
void QElapsedTimer_invalidate(void* self);
int QElapsedTimer_isValid(void* self);
long long QElapsedTimer_nsecsElapsed(void* self);
long long QElapsedTimer_elapsed(void* self);
int QElapsedTimer_hasExpired(void* self, long long timeout);
long long QElapsedTimer_msecsSinceReference(void* self);
long long QElapsedTimer_msecsTo(void* self, void* other);
long long QElapsedTimer_secsTo(void* self, void* other);
int QElapsedTimer_OperatorLesser(void* self, void* rhs);
int QElapsedTimer_OperatorEqual(void* self, void* rhs);
int QElapsedTimer_OperatorNotEqual(void* self, void* rhs);
int QElapsedTimer_OperatorGreater(void* self, void* rhs);
int QElapsedTimer_OperatorLesserOrEqual(void* self, void* rhs);
int QElapsedTimer_OperatorGreaterOrEqual(void* self, void* rhs);
void QElapsedTimer_delete(void* self);

#ifdef __cplusplus
}
#endif
