#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaType_new(int type_);
void* QMetaType_new2();
void* QMetaType_new3(void* arg1);
void QMetaType_registerNormalizedTypedef(void* normalizedTypeName, void* type_);
int QMetaType_type(const char* typeName);
int QMetaType_type2(void* typeName);
const char* QMetaType_typeName(int type_);
int QMetaType_sizeOf(int type_);
int QMetaType_typeFlags(int type_);
void* QMetaType_metaObjectForType(int type_);
int QMetaType_isRegistered(int type_);
int QMetaType_isValid(void* self);
int QMetaType_isRegistered2(void* self);
void QMetaType_registerType(void* self);
int QMetaType_id(void* self);
long long QMetaType_sizeOf2(void* self);
long long QMetaType_alignOf(void* self);
int QMetaType_flags(void* self);
void* QMetaType_metaObject(void* self);
const char* QMetaType_name(void* self);
int QMetaType_isDefaultConstructible(void* self);
int QMetaType_isCopyConstructible(void* self);
int QMetaType_isMoveConstructible(void* self);
int QMetaType_isDestructible(void* self);
int QMetaType_isEqualityComparable(void* self);
int QMetaType_isOrdered(void* self);
int QMetaType_hasRegisteredDataStreamOperators(void* self);
void* QMetaType_underlyingType(void* self);
void* QMetaType_fromName(void* name);
int QMetaType_hasRegisteredDebugStreamOperator(void* self);
int QMetaType_hasRegisteredDebugStreamOperator2(int typeId);
int QMetaType_canConvert(void* fromType, void* toType);
int QMetaType_canView(void* fromType, void* toType);
int QMetaType_hasRegisteredConverterFunction(void* fromType, void* toType);
int QMetaType_hasRegisteredMutableViewFunction(void* fromType, void* toType);
void QMetaType_unregisterConverterFunction(void* from, void* to);
void QMetaType_unregisterMutableViewFunction(void* from, void* to);
void QMetaType_unregisterMetaType(void* type_);
int QMetaType_id2(void* self, int arg1);
int QMetaType_OperatorEqual(void* self, void* rhs);
int QMetaType_OperatorNotEqual(void* self, void* rhs);
void QMetaType_delete(void* self);

#ifdef __cplusplus
}
#endif
