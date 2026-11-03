#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaProperty_new();
const char* QMetaProperty_name(void* self);
const char* QMetaProperty_typeName(void* self);
int QMetaProperty_type(void* self);
int QMetaProperty_userType(void* self);
int QMetaProperty_typeId(void* self);
void* QMetaProperty_metaType(void* self);
int QMetaProperty_propertyIndex(void* self);
int QMetaProperty_relativePropertyIndex(void* self);
int QMetaProperty_isReadable(void* self);
int QMetaProperty_isWritable(void* self);
int QMetaProperty_isResettable(void* self);
int QMetaProperty_isDesignable(void* self);
int QMetaProperty_isScriptable(void* self);
int QMetaProperty_isStored(void* self);
int QMetaProperty_isUser(void* self);
int QMetaProperty_isConstant(void* self);
int QMetaProperty_isFinal(void* self);
int QMetaProperty_isVirtual(void* self);
int QMetaProperty_isOverride(void* self);
int QMetaProperty_isRequired(void* self);
int QMetaProperty_isBindable(void* self);
int QMetaProperty_isFlagType(void* self);
int QMetaProperty_isEnumType(void* self);
void* QMetaProperty_enumerator(void* self);
int QMetaProperty_hasNotifySignal(void* self);
void* QMetaProperty_notifySignal(void* self);
int QMetaProperty_notifySignalIndex(void* self);
int QMetaProperty_revision(void* self);
void* QMetaProperty_read(void* self, void* obj);
int QMetaProperty_write(void* self, void* obj, void* value);
int QMetaProperty_reset(void* self, void* obj);
void* QMetaProperty_bindable(void* self, void* object);
int QMetaProperty_hasStdCppSet(void* self);
int QMetaProperty_isAlias(void* self);
int QMetaProperty_isValid(void* self);
void* QMetaProperty_enclosingMetaObject(void* self);
void QMetaProperty_delete(void* self);

#ifdef __cplusplus
}
#endif
