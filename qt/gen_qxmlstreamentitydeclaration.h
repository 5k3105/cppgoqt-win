#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamEntityDeclaration_new();
void* QXmlStreamEntityDeclaration_new2(void* arg1);
void* QXmlStreamEntityDeclaration_name(void* self);
void* QXmlStreamEntityDeclaration_notationName(void* self);
void* QXmlStreamEntityDeclaration_systemId(void* self);
void* QXmlStreamEntityDeclaration_publicId(void* self);
void* QXmlStreamEntityDeclaration_value(void* self);
int QXmlStreamEntityDeclaration_OperatorEqual(void* self, void* rhs);
int QXmlStreamEntityDeclaration_OperatorNotEqual(void* self, void* rhs);
void QXmlStreamEntityDeclaration_delete(void* self);

#ifdef __cplusplus
}
#endif
