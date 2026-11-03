#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QHashDummyValue_new();
int QHashDummyValue_OperatorEqual(void* self, void* arg1);
int QHashDummyValue_OperatorNotEqual(void* self, void* arg1);
void QHashDummyValue_delete(void* self);

#ifdef __cplusplus
}
#endif
