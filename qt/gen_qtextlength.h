#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextLength_new();
void* QTextLength_new2(int type_, double value);
void* QTextLength_new3(void* arg1);
int QTextLength_type(void* self);
double QTextLength_value(void* self, double maximumLength);
double QTextLength_rawValue(void* self);
int QTextLength_OperatorEqual(void* self, void* other);
int QTextLength_OperatorNotEqual(void* self, void* other);
void* QTextLength_ToQVariant(void* self);
void* QTextLength_toQVariant(void* self);
void QTextLength_delete(void* self);

#ifdef __cplusplus
}
#endif
