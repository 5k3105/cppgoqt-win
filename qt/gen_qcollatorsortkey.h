#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCollatorSortKey_new(void* other);
void* QCollatorSortKey_OperatorAssign(void* self, void* other);
void QCollatorSortKey_swap(void* self, void* other);
int QCollatorSortKey_compare(void* self, void* key);
int QCollatorSortKey_OperatorLesser(void* self, void* rhs);
void QCollatorSortKey_delete(void* self);

#ifdef __cplusplus
}
#endif
