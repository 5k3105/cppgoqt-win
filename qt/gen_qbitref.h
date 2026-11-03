#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QBitRef_ToBool(void* self);
int QBitRef_OperatorNot(void* self);
void* QBitRef_OperatorAssign(void* self, void* val);
void* QBitRef_OperatorAssign2(void* self, int val);
void QBitRef_delete(void* self);

#ifdef __cplusplus
}
#endif
