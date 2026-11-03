#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaAssociation_base(void* self);
void* QMetaAssociation_new();
void* QMetaAssociation_new2(void* arg1);
void* QMetaAssociation_keyMetaType(void* self);
void* QMetaAssociation_mappedMetaType(void* self);
int QMetaAssociation_canInsertKey(void* self);
int QMetaAssociation_canRemoveKey(void* self);
int QMetaAssociation_canContainsKey(void* self);
int QMetaAssociation_canGetMappedAtKey(void* self);
int QMetaAssociation_canSetMappedAtKey(void* self);
int QMetaAssociation_canGetKeyAtIterator(void* self);
int QMetaAssociation_canGetKeyAtConstIterator(void* self);
int QMetaAssociation_canGetMappedAtIterator(void* self);
int QMetaAssociation_canGetMappedAtConstIterator(void* self);
int QMetaAssociation_canSetMappedAtIterator(void* self);
int QMetaAssociation_canCreateIteratorAtKey(void* self);
int QMetaAssociation_canCreateConstIteratorAtKey(void* self);
int QMetaAssociation_OperatorEqual(void* self, void* rhs);
int QMetaAssociation_OperatorNotEqual(void* self, void* rhs);
void QMetaAssociation_delete(void* self);

#ifdef __cplusplus
}
#endif
