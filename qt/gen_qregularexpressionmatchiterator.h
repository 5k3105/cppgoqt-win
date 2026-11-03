#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpressionMatchIterator_new();
void* QRegularExpressionMatchIterator_new2(void* iterator);
void* QRegularExpressionMatchIterator_OperatorAssign(void* self, void* iterator);
void QRegularExpressionMatchIterator_swap(void* self, void* other);
int QRegularExpressionMatchIterator_isValid(void* self);
int QRegularExpressionMatchIterator_hasNext(void* self);
void* QRegularExpressionMatchIterator_next(void* self);
void* QRegularExpressionMatchIterator_peekNext(void* self);
void* QRegularExpressionMatchIterator_regularExpression(void* self);
int QRegularExpressionMatchIterator_matchType(void* self);
int QRegularExpressionMatchIterator_matchOptions(void* self);
void QRegularExpressionMatchIterator_delete(void* self);

#ifdef __cplusplus
}
#endif
