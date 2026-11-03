#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFlag_new(int value);
void* QFlag_new2(long value);
void* QFlag_new3(unsigned long value);
void* QFlag_new4(unsigned int value);
void* QFlag_new5(short value);
void* QFlag_new6(unsigned short value);
void* QFlag_new7(void* arg1);
int QFlag_ToInt(void* self);
unsigned int QFlag_ToUint(void* self);
void QFlag_delete(void* self);

#ifdef __cplusplus
}
#endif
