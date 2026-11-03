#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPolygon_new();
void* QPolygon_new2(void** v, long long v_len);
void* QPolygon_new3(void* r);
void* QPolygon_new4(void* arg1);
void* QPolygon_new5(void* r, int closed);
void QPolygon_swap(void* self, void* other);
void* QPolygon_ToQVariant(void* self);
void QPolygon_translate(void* self, int dx, int dy);
void QPolygon_translate2(void* self, void* offset);
void* QPolygon_translated(void* self, int dx, int dy);
void* QPolygon_translated2(void* self, void* offset);
void* QPolygon_boundingRect(void* self);
void QPolygon_point(void* self, int i, int* x, int* y);
void* QPolygon_point2(void* self, int i);
void QPolygon_setPoint(void* self, int index, int x, int y);
void QPolygon_setPoint2(void* self, int index, void* p);
void QPolygon_setPoints(void* self, int nPoints, int firstx, int firsty);
void QPolygon_putPoints(void* self, int index, int nPoints, int firstx, int firsty);
void QPolygon_putPoints2(void* self, int index, int nPoints, void* from);
int QPolygon_containsPoint(void* self, void* pt, int fillRule);
void* QPolygon_united(void* self, void* r);
void* QPolygon_intersected(void* self, void* r);
void* QPolygon_subtracted(void* self, void* r);
int QPolygon_intersects(void* self, void* r);
void* QPolygon_toPolygonF(void* self);
long long QPolygon_maxSize();
long long QPolygon_size(void* self);
long long QPolygon_count(void* self);
long long QPolygon_length(void* self);
int QPolygon_isEmpty(void* self);
void QPolygon_resize(void* self, long long size);
void QPolygon_resize2(void* self, long long size, void* c);
void QPolygon_resizeForOverwrite(void* self, long long size);
long long QPolygon_capacity(void* self);
void QPolygon_reserve(void* self, long long size);
void QPolygon_squeeze(void* self);
void QPolygon_detach(void* self);
int QPolygon_isDetached(void* self);
int QPolygon_isSharedWith(void* self, void** other, long long other_len);
void* QPolygon_data(void* self);
void* QPolygon_data2(void* self);
void* QPolygon_constData(void* self);
void QPolygon_clear(void* self);
void* QPolygon_at(void* self, long long i);
void* QPolygon_OperatorSubscript(void* self, long long i);
void* QPolygon_OperatorSubscript2(void* self, long long i);
void QPolygon_append(void* self, void* t);
void QPolygon_append2(void* self, void** l, long long l_len);
void QPolygon_prepend(void* self, void* t);
void QPolygon_replace(void* self, long long i, void* t);
void QPolygon_remove(void* self, long long i);
void QPolygon_removeFirst(void* self);
void QPolygon_removeLast(void* self);
void* QPolygon_takeFirst(void* self);
void* QPolygon_takeLast(void* self);
void QPolygon_removeAt(void* self, long long i);
void* QPolygon_takeAt(void* self, long long i);
void QPolygon_move(void* self, long long from, long long to);
void* QPolygon_first(void* self);
void* QPolygon_first2(void* self);
void* QPolygon_constFirst(void* self);
void* QPolygon_last(void* self);
void* QPolygon_last2(void* self);
void* QPolygon_constLast(void* self);
int QPolygon_startsWith(void* self, void* t);
int QPolygon_endsWith(void* self, void* t);
void** QPolygon_mid(void* self, long long pos, long long* __out_len);
void** QPolygon_first3(void* self, long long n, long long* __out_len);
void** QPolygon_last3(void* self, long long n, long long* __out_len);
void** QPolygon_sliced(void* self, long long pos, long long* __out_len);
void** QPolygon_sliced2(void* self, long long pos, long long n, long long* __out_len);
void* QPolygon_value(void* self, long long i);
void* QPolygon_value2(void* self, long long i, void* defaultValue);
void QPolygon_swapItemsAt(void* self, long long i, long long j);
void QPolygon_push_back(void* self, void* t);
void QPolygon_push_front(void* self, void* t);
void QPolygon_pop_back(void* self);
void QPolygon_pop_front(void* self);
int QPolygon_empty(void* self);
void* QPolygon_front(void* self);
void* QPolygon_front2(void* self);
void* QPolygon_back(void* self);
void* QPolygon_back2(void* self);
void QPolygon_shrink_to_fit(void* self);
long long QPolygon_max_size(void* self);
void** QPolygon_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len);
void** QPolygon_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len);
void** QPolygon_fromList(void** list, long long list_len, long long* __out_len);
void** QPolygon_toList(void* self, long long* __out_len);
void** QPolygon_fromVector(void** vector, long long vector_len, long long* __out_len);
void** QPolygon_toVector(void* self, long long* __out_len);
void QPolygon_putPoints3(void* self, int index, int nPoints, void* from, int fromIndex);
void QPolygon_remove2(void* self, long long i, long long n);
void** QPolygon_mid2(void* self, long long pos, long long len_, long long* __out_len);
void* QPolygon_OperatorMultiply(void* self, void* m);
void* QPolygon_toQVariant(void* self);
void QPolygon_delete(void* self);

#ifdef __cplusplus
}
#endif
