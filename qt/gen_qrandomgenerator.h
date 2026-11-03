#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRandomGenerator_new();
void* QRandomGenerator_new2(void* other);
void* QRandomGenerator_new3(unsigned int seedValue);
void* QRandomGenerator_OperatorAssign(void* self, void* other);
unsigned int QRandomGenerator_generate(void* self);
unsigned long long QRandomGenerator_generate64(void* self);
double QRandomGenerator_generateDouble(void* self);
double QRandomGenerator_bounded(void* self, double highest);
unsigned int QRandomGenerator_bounded2(void* self, unsigned int highest);
unsigned int QRandomGenerator_bounded3(void* self, unsigned int lowest, unsigned int highest);
int QRandomGenerator_bounded4(void* self, int highest);
int QRandomGenerator_bounded5(void* self, int lowest, int highest);
unsigned long long QRandomGenerator_bounded6(void* self, unsigned long long highest);
unsigned long long QRandomGenerator_bounded7(void* self, unsigned long long lowest, unsigned long long highest);
long long QRandomGenerator_bounded8(void* self, long long highest);
long long QRandomGenerator_bounded9(void* self, long long lowest, long long highest);
long long QRandomGenerator_bounded10(void* self, int lowest, long long highest);
long long QRandomGenerator_bounded11(void* self, long long lowest, int highest);
unsigned long long QRandomGenerator_bounded12(void* self, unsigned int lowest, unsigned long long highest);
unsigned long long QRandomGenerator_bounded13(void* self, unsigned long long lowest, unsigned int highest);
void QRandomGenerator_generate2(void* self, unsigned int* begin, unsigned int* end);
unsigned int QRandomGenerator_OperatorCall(void* self);
void QRandomGenerator_seed(void* self);
void QRandomGenerator_discard(void* self, unsigned long long z);
unsigned int QRandomGenerator_min();
unsigned int QRandomGenerator_max();
void* QRandomGenerator_system();
void* QRandomGenerator_global();
void* QRandomGenerator_securelySeeded();
void QRandomGenerator_seed2(void* self, unsigned int s);
int QRandomGenerator_OperatorEqual(void* self, void* rng2);
int QRandomGenerator_OperatorNotEqual(void* self, void* rng2);
void QRandomGenerator_delete(void* self);

#ifdef __cplusplus
}
#endif
