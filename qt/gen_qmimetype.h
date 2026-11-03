#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMimeType_new();
void* QMimeType_new2(void* other);
void* QMimeType_OperatorAssign(void* self, void* other);
void QMimeType_swap(void* self, void* other);
int QMimeType_isValid(void* self);
int QMimeType_isDefault(void* self);
char* QMimeType_name(void* self, long long* __out_len);
char* QMimeType_comment(void* self, long long* __out_len);
char* QMimeType_genericIconName(void* self, long long* __out_len);
char* QMimeType_iconName(void* self, long long* __out_len);
char** QMimeType_globPatterns(void* self, long long* __out_len);
char** QMimeType_parentMimeTypes(void* self, long long* __out_len);
char** QMimeType_allAncestors(void* self, long long* __out_len);
char** QMimeType_aliases(void* self, long long* __out_len);
char** QMimeType_suffixes(void* self, long long* __out_len);
char* QMimeType_preferredSuffix(void* self, long long* __out_len);
int QMimeType_inherits(void* self, const char* mimeTypeName, long long mimeTypeName_len);
char* QMimeType_filterString(void* self, long long* __out_len);
int QMimeType_OperatorEqual(void* self, void* rhs);
int QMimeType_OperatorNotEqual(void* self, void* rhs);
void QMimeType_delete(void* self);

#ifdef __cplusplus
}
#endif
