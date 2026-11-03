#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFont__Tag_new();
void* QFont__Tag_new2(void* arg1);
int QFont__Tag_isValid(void* self);
unsigned int QFont__Tag_value(void* self);
void* QFont__Tag_toString(void* self);
void* QFont__Tag_fromValue(unsigned int value, int* __ok);
void* QFont__Tag_fromString(const char* view, long long view_len, int* __ok);
void* QFont__Tag_OperatorAssign(void* self, void* arg1);
int QFont__Tag_OperatorEqual(void* self, void* rhs);
int QFont__Tag_OperatorNotEqual(void* self, void* rhs);
int QFont__Tag_OperatorLesser(void* self, void* rhs);
int QFont__Tag_OperatorGreater(void* self, void* rhs);
int QFont__Tag_OperatorLesserOrEqual(void* self, void* rhs);
int QFont__Tag_OperatorGreaterOrEqual(void* self, void* rhs);
void QFont__Tag_delete(void* self);

#ifdef __cplusplus
}
#endif
