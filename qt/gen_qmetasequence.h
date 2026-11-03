#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaSequence_base(void* self);
void* QMetaSequence_new();
void* QMetaSequence_new2(void* arg1);
void* QMetaSequence_valueMetaType(void* self);
int QMetaSequence_isSortable(void* self);
int QMetaSequence_canAddValueAtBegin(void* self);
int QMetaSequence_canAddValueAtEnd(void* self);
int QMetaSequence_canRemoveValueAtBegin(void* self);
int QMetaSequence_canRemoveValueAtEnd(void* self);
int QMetaSequence_canGetValueAtIndex(void* self);
int QMetaSequence_canSetValueAtIndex(void* self);
int QMetaSequence_canAddValue(void* self);
int QMetaSequence_canRemoveValue(void* self);
int QMetaSequence_canGetValueAtIterator(void* self);
int QMetaSequence_canSetValueAtIterator(void* self);
int QMetaSequence_canInsertValueAtIterator(void* self);
int QMetaSequence_canEraseValueAtIterator(void* self);
int QMetaSequence_canEraseRangeAtIterator(void* self);
int QMetaSequence_canGetValueAtConstIterator(void* self);
int QMetaSequence_OperatorEqual(void* self, void* rhs);
int QMetaSequence_OperatorNotEqual(void* self, void* rhs);
void QMetaSequence_delete(void* self);

#ifdef __cplusplus
}
#endif
