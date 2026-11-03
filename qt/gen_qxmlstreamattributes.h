#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamAttributes_new();
void* QXmlStreamAttributes_value(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len);
void* QXmlStreamAttributes_value2(void* self, const char* qualifiedName, long long qualifiedName_len);
void QXmlStreamAttributes_append(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* value, long long value_len);
void QXmlStreamAttributes_append2(void* self, const char* qualifiedName, long long qualifiedName_len, const char* value, long long value_len);
int QXmlStreamAttributes_hasAttribute(void* self, const char* qualifiedName, long long qualifiedName_len);
int QXmlStreamAttributes_hasAttribute2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len);
long long QXmlStreamAttributes_maxSize();
long long QXmlStreamAttributes_size(void* self);
long long QXmlStreamAttributes_count(void* self);
long long QXmlStreamAttributes_length(void* self);
int QXmlStreamAttributes_isEmpty(void* self);
void QXmlStreamAttributes_resize(void* self, long long size);
void QXmlStreamAttributes_resize2(void* self, long long size, void* c);
void QXmlStreamAttributes_resizeForOverwrite(void* self, long long size);
long long QXmlStreamAttributes_capacity(void* self);
void QXmlStreamAttributes_reserve(void* self, long long size);
void QXmlStreamAttributes_squeeze(void* self);
void QXmlStreamAttributes_detach(void* self);
int QXmlStreamAttributes_isDetached(void* self);
int QXmlStreamAttributes_isSharedWith(void* self, void** other, long long other_len);
void* QXmlStreamAttributes_data(void* self);
void* QXmlStreamAttributes_data2(void* self);
void* QXmlStreamAttributes_constData(void* self);
void QXmlStreamAttributes_clear(void* self);
void* QXmlStreamAttributes_at(void* self, long long i);
void* QXmlStreamAttributes_OperatorSubscript(void* self, long long i);
void* QXmlStreamAttributes_OperatorSubscript2(void* self, long long i);
void QXmlStreamAttributes_prepend(void* self, void* t);
void QXmlStreamAttributes_replace(void* self, long long i, void* t);
void QXmlStreamAttributes_remove(void* self, long long i);
void QXmlStreamAttributes_removeFirst(void* self);
void QXmlStreamAttributes_removeLast(void* self);
void* QXmlStreamAttributes_takeFirst(void* self);
void* QXmlStreamAttributes_takeLast(void* self);
void QXmlStreamAttributes_removeAt(void* self, long long i);
void* QXmlStreamAttributes_takeAt(void* self, long long i);
void QXmlStreamAttributes_move(void* self, long long from, long long to);
void* QXmlStreamAttributes_first(void* self);
void* QXmlStreamAttributes_first2(void* self);
void* QXmlStreamAttributes_constFirst(void* self);
void* QXmlStreamAttributes_last(void* self);
void* QXmlStreamAttributes_last2(void* self);
void* QXmlStreamAttributes_constLast(void* self);
int QXmlStreamAttributes_startsWith(void* self, void* t);
int QXmlStreamAttributes_endsWith(void* self, void* t);
void** QXmlStreamAttributes_mid(void* self, long long pos, long long* __out_len);
void** QXmlStreamAttributes_first3(void* self, long long n, long long* __out_len);
void** QXmlStreamAttributes_last3(void* self, long long n, long long* __out_len);
void** QXmlStreamAttributes_sliced(void* self, long long pos, long long* __out_len);
void** QXmlStreamAttributes_sliced2(void* self, long long pos, long long n, long long* __out_len);
void QXmlStreamAttributes_swapItemsAt(void* self, long long i, long long j);
void QXmlStreamAttributes_push_back(void* self, void* t);
void QXmlStreamAttributes_push_front(void* self, void* t);
void QXmlStreamAttributes_pop_back(void* self);
void QXmlStreamAttributes_pop_front(void* self);
int QXmlStreamAttributes_empty(void* self);
void* QXmlStreamAttributes_front(void* self);
void* QXmlStreamAttributes_front2(void* self);
void* QXmlStreamAttributes_back(void* self);
void* QXmlStreamAttributes_back2(void* self);
void QXmlStreamAttributes_shrink_to_fit(void* self);
long long QXmlStreamAttributes_max_size(void* self);
void** QXmlStreamAttributes_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len);
void** QXmlStreamAttributes_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len);
void** QXmlStreamAttributes_fromList(void** list, long long list_len, long long* __out_len);
void** QXmlStreamAttributes_toList(void* self, long long* __out_len);
void** QXmlStreamAttributes_fromVector(void** vector, long long vector_len, long long* __out_len);
void** QXmlStreamAttributes_toVector(void* self, long long* __out_len);
void QXmlStreamAttributes_remove2(void* self, long long i, long long n);
void** QXmlStreamAttributes_mid2(void* self, long long pos, long long len_, long long* __out_len);
void QXmlStreamAttributes_delete(void* self);

#ifdef __cplusplus
}
#endif
