#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCborArray_new();
void* QCborArray_new2(void* other);
void* QCborArray_OperatorAssign(void* self, void* other);
void QCborArray_swap(void* self, void* other);
void* QCborArray_toCborValue(void* self);
long long QCborArray_size(void* self);
int QCborArray_isEmpty(void* self);
void QCborArray_clear(void* self);
void* QCborArray_at(void* self, long long i);
void* QCborArray_first(void* self);
void* QCborArray_last(void* self);
void* QCborArray_OperatorSubscript(void* self, long long i);
void* QCborArray_first2(void* self);
void* QCborArray_last2(void* self);
void* QCborArray_OperatorSubscript2(void* self, long long i);
void QCborArray_insert(void* self, long long i, void* value);
void QCborArray_prepend(void* self, void* value);
void QCborArray_append(void* self, void* value);
void QCborArray_removeAt(void* self, long long i);
void* QCborArray_takeAt(void* self, long long i);
void QCborArray_removeFirst(void* self);
void QCborArray_removeLast(void* self);
void* QCborArray_takeFirst(void* self);
void* QCborArray_takeLast(void* self);
int QCborArray_contains(void* self, void* value);
int QCborArray_compare(void* self, void* other);
void QCborArray_push_back(void* self, void* t);
void QCborArray_push_front(void* self, void* t);
void QCborArray_pop_front(void* self);
void QCborArray_pop_back(void* self);
int QCborArray_empty(void* self);
void* QCborArray_OperatorPlus(void* self, void* v);
void* QCborArray_OperatorPlusAssign(void* self, void* v);
void* QCborArray_OperatorShiftLeft(void* self, void* v);
void* QCborArray_fromStringList(char** list, long long list_len);
void* QCborArray_fromVariantList(void** list, long long list_len);
void* QCborArray_fromJsonArray(void* array);
void** QCborArray_toVariantList(void* self, long long* __out_len);
void* QCborArray_toJsonArray(void* self);
int QCborArray_OperatorEqual(void* self, void* rhs);
int QCborArray_OperatorNotEqual(void* self, void* rhs);
int QCborArray_OperatorLesser(void* self, void* rhs);
int QCborArray_OperatorGreater(void* self, void* rhs);
int QCborArray_OperatorLesserOrEqual(void* self, void* rhs);
int QCborArray_OperatorGreaterOrEqual(void* self, void* rhs);
int QCborArray_OperatorEqual2(void* self, void* rhs);
int QCborArray_OperatorNotEqual2(void* self, void* rhs);
int QCborArray_OperatorLesser2(void* self, void* rhs);
int QCborArray_OperatorGreater2(void* self, void* rhs);
int QCborArray_OperatorLesserOrEqual2(void* self, void* rhs);
int QCborArray_OperatorGreaterOrEqual2(void* self, void* rhs);
int QCborArray_OperatorEqual3(void* self, void* rhs);
int QCborArray_OperatorNotEqual3(void* self, void* rhs);
int QCborArray_OperatorLesser3(void* self, void* rhs);
int QCborArray_OperatorGreater3(void* self, void* rhs);
int QCborArray_OperatorLesserOrEqual3(void* self, void* rhs);
int QCborArray_OperatorGreaterOrEqual3(void* self, void* rhs);
void QCborArray_delete(void* self);

#ifdef __cplusplus
}
#endif
