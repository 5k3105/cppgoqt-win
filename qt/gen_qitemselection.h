#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QItemSelection_new(void* topLeft, void* bottomRight);
void* QItemSelection_new2();
void* QItemSelection_new3(void* arg1);
void QItemSelection_select(void* self, void* topLeft, void* bottomRight);
int QItemSelection_contains(void* self, void* index);
void** QItemSelection_indexes(void* self, long long* __out_len);
void QItemSelection_merge(void* self, void* other, int command);
void QItemSelection_split(void* range_, void* other, void* result);
long long QItemSelection_maxSize();
long long QItemSelection_size(void* self);
long long QItemSelection_count(void* self);
long long QItemSelection_length(void* self);
int QItemSelection_isEmpty(void* self);
void QItemSelection_resize(void* self, long long size);
void QItemSelection_resize2(void* self, long long size, void* c);
void QItemSelection_resizeForOverwrite(void* self, long long size);
long long QItemSelection_capacity(void* self);
void QItemSelection_reserve(void* self, long long size);
void QItemSelection_squeeze(void* self);
void QItemSelection_detach(void* self);
int QItemSelection_isDetached(void* self);
int QItemSelection_isSharedWith(void* self, void** other, long long other_len);
void* QItemSelection_data(void* self);
void* QItemSelection_data2(void* self);
void* QItemSelection_constData(void* self);
void QItemSelection_clear(void* self);
void* QItemSelection_at(void* self, long long i);
void* QItemSelection_OperatorSubscript(void* self, long long i);
void* QItemSelection_OperatorSubscript2(void* self, long long i);
void QItemSelection_append(void* self, void* t);
void QItemSelection_append2(void* self, void** l, long long l_len);
void QItemSelection_prepend(void* self, void* t);
void QItemSelection_replace(void* self, long long i, void* t);
void QItemSelection_remove(void* self, long long i);
void QItemSelection_removeFirst(void* self);
void QItemSelection_removeLast(void* self);
void* QItemSelection_takeFirst(void* self);
void* QItemSelection_takeLast(void* self);
void QItemSelection_removeAt(void* self, long long i);
void* QItemSelection_takeAt(void* self, long long i);
void QItemSelection_move(void* self, long long from, long long to);
void* QItemSelection_first(void* self);
void* QItemSelection_first2(void* self);
void* QItemSelection_constFirst(void* self);
void* QItemSelection_last(void* self);
void* QItemSelection_last2(void* self);
void* QItemSelection_constLast(void* self);
int QItemSelection_startsWith(void* self, void* t);
int QItemSelection_endsWith(void* self, void* t);
void** QItemSelection_mid(void* self, long long pos, long long* __out_len);
void** QItemSelection_first3(void* self, long long n, long long* __out_len);
void** QItemSelection_last3(void* self, long long n, long long* __out_len);
void** QItemSelection_sliced(void* self, long long pos, long long* __out_len);
void** QItemSelection_sliced2(void* self, long long pos, long long n, long long* __out_len);
void* QItemSelection_value(void* self, long long i);
void* QItemSelection_value2(void* self, long long i, void* defaultValue);
void QItemSelection_swapItemsAt(void* self, long long i, long long j);
void QItemSelection_push_back(void* self, void* t);
void QItemSelection_push_front(void* self, void* t);
void QItemSelection_pop_back(void* self);
void QItemSelection_pop_front(void* self);
int QItemSelection_empty(void* self);
void* QItemSelection_front(void* self);
void* QItemSelection_front2(void* self);
void* QItemSelection_back(void* self);
void* QItemSelection_back2(void* self);
void QItemSelection_shrink_to_fit(void* self);
long long QItemSelection_max_size(void* self);
void** QItemSelection_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len);
void** QItemSelection_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len);
void** QItemSelection_fromList(void** list, long long list_len, long long* __out_len);
void** QItemSelection_toList(void* self, long long* __out_len);
void** QItemSelection_fromVector(void** vector, long long vector_len, long long* __out_len);
void** QItemSelection_toVector(void* self, long long* __out_len);
void QItemSelection_remove2(void* self, long long i, long long n);
void** QItemSelection_mid2(void* self, long long pos, long long len_, long long* __out_len);
void QItemSelection_delete(void* self);

#ifdef __cplusplus
}
#endif
