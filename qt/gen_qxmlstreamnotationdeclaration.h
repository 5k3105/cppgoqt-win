#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamNotationDeclaration_new();
void* QXmlStreamNotationDeclaration_new2(void* arg1);
void* QXmlStreamNotationDeclaration_name(void* self);
void* QXmlStreamNotationDeclaration_systemId(void* self);
void* QXmlStreamNotationDeclaration_publicId(void* self);
int QXmlStreamNotationDeclaration_OperatorEqual(void* self, void* rhs);
int QXmlStreamNotationDeclaration_OperatorNotEqual(void* self, void* rhs);
void QXmlStreamNotationDeclaration_delete(void* self);

#ifdef __cplusplus
}
#endif
