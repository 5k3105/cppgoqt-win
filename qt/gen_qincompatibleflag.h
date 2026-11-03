#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIncompatibleFlag_new(int i);
void* QIncompatibleFlag_new2(void* arg1);
int QIncompatibleFlag_ToInt(void* self);
void QIncompatibleFlag_delete(void* self);

#ifdef __cplusplus
}
#endif
