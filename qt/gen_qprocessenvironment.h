#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QProcessEnvironment_new();
void* QProcessEnvironment_new2(int arg1);
void* QProcessEnvironment_new3(void* other);
void* QProcessEnvironment_OperatorAssign(void* self, void* other);
void QProcessEnvironment_swap(void* self, void* other);
int QProcessEnvironment_isEmpty(void* self);
int QProcessEnvironment_inheritsFromParent(void* self);
void QProcessEnvironment_clear(void* self);
int QProcessEnvironment_contains(void* self, const char* name, long long name_len);
void QProcessEnvironment_insert(void* self, const char* name, long long name_len, const char* value, long long value_len);
void QProcessEnvironment_remove(void* self, const char* name, long long name_len);
char* QProcessEnvironment_value(void* self, const char* name, long long name_len, long long* __out_len);
char** QProcessEnvironment_toStringList(void* self, long long* __out_len);
char** QProcessEnvironment_keys(void* self, long long* __out_len);
void QProcessEnvironment_insert2(void* self, void* e);
void* QProcessEnvironment_systemEnvironment();
char* QProcessEnvironment_value2(void* self, const char* name, long long name_len, const char* defaultValue, long long defaultValue_len, long long* __out_len);
int QProcessEnvironment_OperatorEqual(void* self, void* rhs);
int QProcessEnvironment_OperatorNotEqual(void* self, void* rhs);
void QProcessEnvironment_delete(void* self);

#ifdef __cplusplus
}
#endif
