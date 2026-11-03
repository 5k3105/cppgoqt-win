#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPolygonF_new();
void* QPolygonF_new2(void** v, long long v_len);
void* QPolygonF_new3(void* r);
void* QPolygonF_new4(void* a);
void* QPolygonF_new5(void* arg1);
void QPolygonF_swap(void* self, void* other);
void* QPolygonF_ToQVariant(void* self);
void QPolygonF_translate(void* self, double dx, double dy);
void QPolygonF_translate2(void* self, void* offset);
void* QPolygonF_translated(void* self, double dx, double dy);
void* QPolygonF_translated2(void* self, void* offset);
void* QPolygonF_toPolygon(void* self);
int QPolygonF_isClosed(void* self);
void* QPolygonF_boundingRect(void* self);
int QPolygonF_containsPoint(void* self, void* pt, int fillRule);
void* QPolygonF_united(void* self, void* r);
void* QPolygonF_intersected(void* self, void* r);
void* QPolygonF_subtracted(void* self, void* r);
int QPolygonF_intersects(void* self, void* r);
long long QPolygonF_maxSize();
long long QPolygonF_size(void* self);
long long QPolygonF_count(void* self);
long long QPolygonF_length(void* self);
int QPolygonF_isEmpty(void* self);
void QPolygonF_resize(void* self, long long size);
void QPolygonF_resize2(void* self, long long size, void* c);
void QPolygonF_resizeForOverwrite(void* self, long long size);
long long QPolygonF_capacity(void* self);
void QPolygonF_reserve(void* self, long long size);
void QPolygonF_squeeze(void* self);
void QPolygonF_detach(void* self);
int QPolygonF_isDetached(void* self);
int QPolygonF_isSharedWith(void* self, void** other, long long other_len);
void* QPolygonF_data(void* self);
void* QPolygonF_data2(void* self);
void* QPolygonF_constData(void* self);
void QPolygonF_clear(void* self);
void* QPolygonF_at(void* self, long long i);
void* QPolygonF_OperatorSubscript(void* self, long long i);
void* QPolygonF_OperatorSubscript2(void* self, long long i);
void QPolygonF_append(void* self, void* t);
void QPolygonF_append2(void* self, void** l, long long l_len);
void QPolygonF_prepend(void* self, void* t);
void QPolygonF_replace(void* self, long long i, void* t);
void QPolygonF_remove(void* self, long long i);
void QPolygonF_removeFirst(void* self);
void QPolygonF_removeLast(void* self);
void* QPolygonF_takeFirst(void* self);
void* QPolygonF_takeLast(void* self);
void QPolygonF_removeAt(void* self, long long i);
void* QPolygonF_takeAt(void* self, long long i);
void QPolygonF_move(void* self, long long from, long long to);
void* QPolygonF_first(void* self);
void* QPolygonF_first2(void* self);
void* QPolygonF_constFirst(void* self);
void* QPolygonF_last(void* self);
void* QPolygonF_last2(void* self);
void* QPolygonF_constLast(void* self);
int QPolygonF_startsWith(void* self, void* t);
int QPolygonF_endsWith(void* self, void* t);
void** QPolygonF_mid(void* self, long long pos, long long* __out_len);
void** QPolygonF_first3(void* self, long long n, long long* __out_len);
void** QPolygonF_last3(void* self, long long n, long long* __out_len);
void** QPolygonF_sliced(void* self, long long pos, long long* __out_len);
void** QPolygonF_sliced2(void* self, long long pos, long long n, long long* __out_len);
void* QPolygonF_value(void* self, long long i);
void* QPolygonF_value2(void* self, long long i, void* defaultValue);
void QPolygonF_swapItemsAt(void* self, long long i, long long j);
void QPolygonF_push_back(void* self, void* t);
void QPolygonF_push_front(void* self, void* t);
void QPolygonF_pop_back(void* self);
void QPolygonF_pop_front(void* self);
int QPolygonF_empty(void* self);
void* QPolygonF_front(void* self);
void* QPolygonF_front2(void* self);
void* QPolygonF_back(void* self);
void* QPolygonF_back2(void* self);
void QPolygonF_shrink_to_fit(void* self);
long long QPolygonF_max_size(void* self);
void** QPolygonF_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len);
void** QPolygonF_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len);
void** QPolygonF_fromList(void** list, long long list_len, long long* __out_len);
void** QPolygonF_toList(void* self, long long* __out_len);
void** QPolygonF_fromVector(void** vector, long long vector_len, long long* __out_len);
void** QPolygonF_toVector(void* self, long long* __out_len);
void QPolygonF_remove2(void* self, long long i, long long n);
void** QPolygonF_mid2(void* self, long long pos, long long len_, long long* __out_len);
void* QPolygonF_OperatorMultiply(void* self, void* m);
void* QPolygonF_toQVariant(void* self);
void QPolygonF_delete(void* self);

#ifdef __cplusplus
}
#endif
