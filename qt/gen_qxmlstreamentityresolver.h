#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QXmlStreamEntityResolver__set_override(void* self, int idx);
void* QXmlStreamEntityResolver_new(long long __goid);
char* QXmlStreamEntityResolver_resolveEntity(void* self, const char* publicId, long long publicId_len, const char* systemId, long long systemId_len, long long* __out_len);
char* QXmlStreamEntityResolver_resolveUndeclaredEntity(void* self, const char* name, long long name_len, long long* __out_len);
char* QXmlStreamEntityResolver_resolveEntity_default(void* self, const char* publicId, long long publicId_len, const char* systemId, long long systemId_len, long long* __out_len);
char* QXmlStreamEntityResolver_resolveUndeclaredEntity_default(void* self, const char* name, long long name_len, long long* __out_len);
void* QXmlStreamEntityResolver__Go_new(long long id);
void QXmlStreamEntityResolver_delete(void* self);

#ifdef __cplusplus
}
#endif
