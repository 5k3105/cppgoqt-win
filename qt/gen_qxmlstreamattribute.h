#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamAttribute_new();
void* QXmlStreamAttribute_new2(const char* qualifiedName, long long qualifiedName_len, const char* value, long long value_len);
void* QXmlStreamAttribute_new3(const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* value, long long value_len);
void* QXmlStreamAttribute_new4(void* arg1);
void* QXmlStreamAttribute_namespaceUri(void* self);
void* QXmlStreamAttribute_name(void* self);
void* QXmlStreamAttribute_qualifiedName(void* self);
void* QXmlStreamAttribute_prefix(void* self);
void* QXmlStreamAttribute_value(void* self);
int QXmlStreamAttribute_isDefault(void* self);
int QXmlStreamAttribute_OperatorEqual(void* self, void* rhs);
int QXmlStreamAttribute_OperatorNotEqual(void* self, void* rhs);
void QXmlStreamAttribute_delete(void* self);

#ifdef __cplusplus
}
#endif
