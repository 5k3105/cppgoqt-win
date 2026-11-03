#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamNamespaceDeclaration_new();
void* QXmlStreamNamespaceDeclaration_new2(const char* prefix, long long prefix_len, const char* namespaceUri, long long namespaceUri_len);
void* QXmlStreamNamespaceDeclaration_new3(void* arg1);
void* QXmlStreamNamespaceDeclaration_prefix(void* self);
void* QXmlStreamNamespaceDeclaration_namespaceUri(void* self);
int QXmlStreamNamespaceDeclaration_OperatorEqual(void* self, void* rhs);
int QXmlStreamNamespaceDeclaration_OperatorNotEqual(void* self, void* rhs);
void QXmlStreamNamespaceDeclaration_delete(void* self);

#ifdef __cplusplus
}
#endif
