#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPersistentModelIndex_new();
void* QPersistentModelIndex_new2(void* index);
void* QPersistentModelIndex_new3(void* other);
void* QPersistentModelIndex_OperatorAssign(void* self, void* other);
void QPersistentModelIndex_swap(void* self, void* other);
void* QPersistentModelIndex_OperatorAssign2(void* self, void* other);
void* QPersistentModelIndex_ToQModelIndex(void* self);
int QPersistentModelIndex_row(void* self);
int QPersistentModelIndex_column(void* self);
unsigned long long QPersistentModelIndex_internalId(void* self);
void* QPersistentModelIndex_parent(void* self);
void* QPersistentModelIndex_sibling(void* self, int row, int column);
void* QPersistentModelIndex_data(void* self);
void QPersistentModelIndex_multiData(void* self, void* roleDataSpan);
int QPersistentModelIndex_flags(void* self);
void* QPersistentModelIndex_model(void* self);
int QPersistentModelIndex_isValid(void* self);
void* QPersistentModelIndex_data2(void* self, int role);
int QPersistentModelIndex_OperatorEqual(void* self, void* rhs);
int QPersistentModelIndex_OperatorNotEqual(void* self, void* rhs);
int QPersistentModelIndex_OperatorLesser(void* self, void* rhs);
int QPersistentModelIndex_OperatorGreater(void* self, void* rhs);
int QPersistentModelIndex_OperatorLesserOrEqual(void* self, void* rhs);
int QPersistentModelIndex_OperatorGreaterOrEqual(void* self, void* rhs);
int QPersistentModelIndex_OperatorEqual2(void* self, void* rhs);
int QPersistentModelIndex_OperatorNotEqual2(void* self, void* rhs);
int QPersistentModelIndex_OperatorLesser2(void* self, void* rhs);
int QPersistentModelIndex_OperatorGreater2(void* self, void* rhs);
int QPersistentModelIndex_OperatorLesserOrEqual2(void* self, void* rhs);
int QPersistentModelIndex_OperatorGreaterOrEqual2(void* self, void* rhs);
void QPersistentModelIndex_delete(void* self);

#ifdef __cplusplus
}
#endif
