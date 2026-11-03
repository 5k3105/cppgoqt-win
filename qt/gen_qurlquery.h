#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUrlQuery_new();
void* QUrlQuery_new2(void* url);
void* QUrlQuery_new3(const char* queryString, long long queryString_len);
void* QUrlQuery_new4(void* other);
void* QUrlQuery_OperatorAssign(void* self, void* other);
void QUrlQuery_swap(void* self, void* other);
int QUrlQuery_isEmpty(void* self);
int QUrlQuery_isDetached(void* self);
void QUrlQuery_clear(void* self);
char* QUrlQuery_query(void* self, long long* __out_len);
void QUrlQuery_setQuery(void* self, const char* queryString, long long queryString_len);
char* QUrlQuery_toString(void* self, long long* __out_len);
void QUrlQuery_setQueryDelimiters(void* self, void* valueDelimiter, void* pairDelimiter);
void* QUrlQuery_queryValueDelimiter(void* self);
void* QUrlQuery_queryPairDelimiter(void* self);
int QUrlQuery_hasQueryItem(void* self, const char* key, long long key_len);
void QUrlQuery_addQueryItem(void* self, const char* key, long long key_len, const char* value, long long value_len);
void QUrlQuery_removeQueryItem(void* self, const char* key, long long key_len);
char* QUrlQuery_queryItemValue(void* self, const char* key, long long key_len, long long* __out_len);
char** QUrlQuery_allQueryItemValues(void* self, const char* key, long long key_len, long long* __out_len);
void QUrlQuery_removeAllQueryItems(void* self, const char* key, long long key_len);
unsigned short QUrlQuery_defaultQueryValueDelimiter();
unsigned short QUrlQuery_defaultQueryPairDelimiter();
char* QUrlQuery_query2(void* self, int encoding, long long* __out_len);
char* QUrlQuery_toString2(void* self, int encoding, long long* __out_len);
char* QUrlQuery_queryItemValue2(void* self, const char* key, long long key_len, int encoding, long long* __out_len);
char** QUrlQuery_allQueryItemValues2(void* self, const char* key, long long key_len, int encoding, long long* __out_len);
int QUrlQuery_OperatorEqual(void* self, void* rhs);
int QUrlQuery_OperatorNotEqual(void* self, void* rhs);
void QUrlQuery_delete(void* self);

#ifdef __cplusplus
}
#endif
