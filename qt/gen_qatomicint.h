#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAtomicInt_new();
void* QAtomicInt_new2(void* arg1);
void* QAtomicInt_new3(int value);
void* QAtomicInt_OperatorAssign(void* self, void* arg1);
int QAtomicInt_loadRelaxed(void* self);
void QAtomicInt_storeRelaxed(void* self, int newValue);
int QAtomicInt_loadAcquire(void* self);
void QAtomicInt_storeRelease(void* self, int newValue);
int QAtomicInt_ToInt(void* self);
int QAtomicInt_isReferenceCountingNative();
int QAtomicInt_isReferenceCountingWaitFree();
int QAtomicInt_ref(void* self);
void QAtomicInt_refRelaxed(void* self);
int QAtomicInt_deref(void* self);
int QAtomicInt_isTestAndSetNative();
int QAtomicInt_isTestAndSetWaitFree();
int QAtomicInt_testAndSetRelaxed(void* self, int expectedValue, int newValue);
int QAtomicInt_testAndSetAcquire(void* self, int expectedValue, int newValue);
int QAtomicInt_testAndSetRelease(void* self, int expectedValue, int newValue);
int QAtomicInt_testAndSetOrdered(void* self, int expectedValue, int newValue);
int QAtomicInt_testAndSetRelaxed2(void* self, int expectedValue, int newValue, int* currentValue);
int QAtomicInt_testAndSetAcquire2(void* self, int expectedValue, int newValue, int* currentValue);
int QAtomicInt_testAndSetRelease2(void* self, int expectedValue, int newValue, int* currentValue);
int QAtomicInt_testAndSetOrdered2(void* self, int expectedValue, int newValue, int* currentValue);
int QAtomicInt_isFetchAndStoreNative();
int QAtomicInt_isFetchAndStoreWaitFree();
int QAtomicInt_fetchAndStoreRelaxed(void* self, int newValue);
int QAtomicInt_fetchAndStoreAcquire(void* self, int newValue);
int QAtomicInt_fetchAndStoreRelease(void* self, int newValue);
int QAtomicInt_fetchAndStoreOrdered(void* self, int newValue);
int QAtomicInt_isFetchAndAddNative();
int QAtomicInt_isFetchAndAddWaitFree();
int QAtomicInt_fetchAndAddRelaxed(void* self, int valueToAdd);
int QAtomicInt_fetchAndAddAcquire(void* self, int valueToAdd);
int QAtomicInt_fetchAndAddRelease(void* self, int valueToAdd);
int QAtomicInt_fetchAndAddOrdered(void* self, int valueToAdd);
int QAtomicInt_fetchAndSubRelaxed(void* self, int valueToAdd);
int QAtomicInt_fetchAndSubAcquire(void* self, int valueToAdd);
int QAtomicInt_fetchAndSubRelease(void* self, int valueToAdd);
int QAtomicInt_fetchAndSubOrdered(void* self, int valueToAdd);
int QAtomicInt_fetchAndAndRelaxed(void* self, int valueToAdd);
int QAtomicInt_fetchAndAndAcquire(void* self, int valueToAdd);
int QAtomicInt_fetchAndAndRelease(void* self, int valueToAdd);
int QAtomicInt_fetchAndAndOrdered(void* self, int valueToAdd);
int QAtomicInt_fetchAndOrRelaxed(void* self, int valueToAdd);
int QAtomicInt_fetchAndOrAcquire(void* self, int valueToAdd);
int QAtomicInt_fetchAndOrRelease(void* self, int valueToAdd);
int QAtomicInt_fetchAndOrOrdered(void* self, int valueToAdd);
int QAtomicInt_fetchAndXorRelaxed(void* self, int valueToAdd);
int QAtomicInt_fetchAndXorAcquire(void* self, int valueToAdd);
int QAtomicInt_fetchAndXorRelease(void* self, int valueToAdd);
int QAtomicInt_fetchAndXorOrdered(void* self, int valueToAdd);
int QAtomicInt_OperatorPlusPlus(void* self);
int QAtomicInt_OperatorPlusPlus2(void* self, int arg1);
int QAtomicInt_OperatorMinusMinus(void* self);
int QAtomicInt_OperatorMinusMinus2(void* self, int arg1);
int QAtomicInt_OperatorPlusAssign(void* self, int v);
int QAtomicInt_OperatorMinusAssign(void* self, int v);
int QAtomicInt_OperatorBitwiseAndAssign(void* self, int v);
int QAtomicInt_OperatorBitwiseOrAssign(void* self, int v);
int QAtomicInt_OperatorBitwiseXorAssign(void* self, int v);
void QAtomicInt_delete(void* self);

#ifdef __cplusplus
}
#endif
