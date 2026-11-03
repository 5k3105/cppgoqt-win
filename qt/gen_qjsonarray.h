#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QJsonArray_new();
void* QJsonArray_new2(void* other);
void* QJsonArray_OperatorAssign(void* self, void* other);
void* QJsonArray_fromStringList(char** list, long long list_len);
void* QJsonArray_fromVariantList(void** list, long long list_len);
void** QJsonArray_toVariantList(void* self, long long* __out_len);
long long QJsonArray_size(void* self);
long long QJsonArray_count(void* self);
int QJsonArray_isEmpty(void* self);
void* QJsonArray_at(void* self, long long i);
void* QJsonArray_first(void* self);
void* QJsonArray_last(void* self);
void QJsonArray_prepend(void* self, void* value);
void QJsonArray_append(void* self, void* value);
void QJsonArray_removeAt(void* self, long long i);
void* QJsonArray_takeAt(void* self, long long i);
void QJsonArray_removeFirst(void* self);
void QJsonArray_removeLast(void* self);
void QJsonArray_insert(void* self, long long i, void* value);
void QJsonArray_replace(void* self, long long i, void* value);
int QJsonArray_contains(void* self, void* element);
void* QJsonArray_OperatorSubscript(void* self, long long i);
void* QJsonArray_OperatorSubscript2(void* self, long long i);
void QJsonArray_swap(void* self, void* other);
void* QJsonArray_OperatorPlus(void* self, void* v);
void* QJsonArray_OperatorPlusAssign(void* self, void* v);
void* QJsonArray_OperatorShiftLeft(void* self, void* v);
void QJsonArray_push_back(void* self, void* t);
void QJsonArray_push_front(void* self, void* t);
void QJsonArray_pop_front(void* self);
void QJsonArray_pop_back(void* self);
int QJsonArray_empty(void* self);
int QJsonArray_OperatorEqual(void* self, void* rhs);
int QJsonArray_OperatorNotEqual(void* self, void* rhs);
int QJsonArray_OperatorEqual2(void* self, void* rhs);
int QJsonArray_OperatorNotEqual2(void* self, void* rhs);
void QJsonArray_delete(void* self);

#ifdef __cplusplus
}
#endif
