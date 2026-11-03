#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextFormat_new();
void* QTextFormat_new2(int type_);
void* QTextFormat_new3(void* rhs);
void* QTextFormat_OperatorAssign(void* self, void* rhs);
void QTextFormat_swap(void* self, void* other);
void QTextFormat_merge(void* self, void* other);
int QTextFormat_isValid(void* self);
int QTextFormat_isEmpty(void* self);
int QTextFormat_type(void* self);
int QTextFormat_objectIndex(void* self);
void QTextFormat_setObjectIndex(void* self, int object);
void* QTextFormat_property(void* self, int propertyId);
void QTextFormat_setProperty(void* self, int propertyId, void* value);
void QTextFormat_clearProperty(void* self, int propertyId);
int QTextFormat_hasProperty(void* self, int propertyId);
int QTextFormat_boolProperty(void* self, int propertyId);
int QTextFormat_intProperty(void* self, int propertyId);
double QTextFormat_doubleProperty(void* self, int propertyId);
char* QTextFormat_stringProperty(void* self, int propertyId, long long* __out_len);
void* QTextFormat_colorProperty(void* self, int propertyId);
void* QTextFormat_penProperty(void* self, int propertyId);
void* QTextFormat_brushProperty(void* self, int propertyId);
void* QTextFormat_lengthProperty(void* self, int propertyId);
void** QTextFormat_lengthVectorProperty(void* self, int propertyId, long long* __out_len);
void QTextFormat_setProperty2(void* self, int propertyId, void** lengths, long long lengths_len);
int* QTextFormat_properties(void* self, void** __out_vals, long long* __out_len);
int QTextFormat_propertyCount(void* self);
void QTextFormat_setObjectType(void* self, int type_);
int QTextFormat_objectType(void* self);
int QTextFormat_isCharFormat(void* self);
int QTextFormat_isBlockFormat(void* self);
int QTextFormat_isListFormat(void* self);
int QTextFormat_isFrameFormat(void* self);
int QTextFormat_isImageFormat(void* self);
int QTextFormat_isTableFormat(void* self);
int QTextFormat_isTableCellFormat(void* self);
void* QTextFormat_toBlockFormat(void* self);
void* QTextFormat_toCharFormat(void* self);
void* QTextFormat_toListFormat(void* self);
void* QTextFormat_toTableFormat(void* self);
void* QTextFormat_toFrameFormat(void* self);
void* QTextFormat_toImageFormat(void* self);
void* QTextFormat_toTableCellFormat(void* self);
int QTextFormat_OperatorEqual(void* self, void* rhs);
int QTextFormat_OperatorNotEqual(void* self, void* rhs);
void* QTextFormat_ToQVariant(void* self);
void QTextFormat_setLayoutDirection(void* self, int direction);
int QTextFormat_layoutDirection(void* self);
void QTextFormat_setBackground(void* self, void* brush);
void* QTextFormat_background(void* self);
void QTextFormat_clearBackground(void* self);
void QTextFormat_setForeground(void* self, void* brush);
void* QTextFormat_foreground(void* self);
void QTextFormat_clearForeground(void* self);
void* QTextFormat_toQVariant(void* self);
void QTextFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
