#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPartialOrdering_new(void* order);
void* QPartialOrdering_new2(void* stdorder);
void* QPartialOrdering_new3(void* stdorder);
void* QPartialOrdering_new4(void* arg1);
void* QPartialOrdering_ToQtPartialOrdering(void* self);
int QPartialOrdering_OperatorEqual(void* self, void* rhs);
int QPartialOrdering_OperatorNotEqual(void* self, void* rhs);
int QPartialOrdering_OperatorEqual2(void* self, void* rhs);
int QPartialOrdering_OperatorNotEqual2(void* self, void* rhs);
void QPartialOrdering_delete(void* self);

#ifdef __cplusplus
}
#endif
