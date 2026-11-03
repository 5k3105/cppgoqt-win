#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* Qt__strong_ordering_new(void* arg1);
void* Qt__strong_ordering_ToQtPartialOrdering(void* self);
void* Qt__strong_ordering_ToQtWeakOrdering(void* self);
int Qt__strong_ordering_OperatorEqual(void* self, void* rhs);
int Qt__strong_ordering_OperatorNotEqual(void* self, void* rhs);
int Qt__strong_ordering_OperatorEqual2(void* self, void* rhs);
int Qt__strong_ordering_OperatorNotEqual2(void* self, void* rhs);
int Qt__strong_ordering_OperatorEqual3(void* self, void* rhs);
int Qt__strong_ordering_OperatorNotEqual3(void* self, void* rhs);
void Qt__strong_ordering_delete(void* self);

#ifdef __cplusplus
}
#endif
