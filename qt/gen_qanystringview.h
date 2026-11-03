#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAnyStringView_new();
void* QAnyStringView_new2(void* str);
void* QAnyStringView_new3(const char* str, long long str_len);
void* QAnyStringView_new4(void* str);
void* QAnyStringView_new5(void* v);
void* QAnyStringView_new6(const char* arg1, long long arg1_len);
void* QAnyStringView_mid(void* self, long long pos);
void* QAnyStringView_left(void* self, long long n);
void* QAnyStringView_right(void* self, long long n);
void* QAnyStringView_sliced(void* self, long long pos);
void* QAnyStringView_sliced2(void* self, long long pos, long long n);
void* QAnyStringView_first(void* self, long long n);
void* QAnyStringView_last(void* self, long long n);
void* QAnyStringView_chopped(void* self, long long n);
void* QAnyStringView_slice(void* self, long long pos);
void* QAnyStringView_slice2(void* self, long long pos, long long n);
void QAnyStringView_truncate(void* self, long long n);
void QAnyStringView_chop(void* self, long long n);
char* QAnyStringView_toString(void* self, long long* __out_len);
long long QAnyStringView_size(void* self);
int QAnyStringView_compare(const char* lhs, long long lhs_len, const char* rhs, long long rhs_len);
int QAnyStringView_equal(const char* lhs, long long lhs_len, const char* rhs, long long rhs_len);
void* QAnyStringView_front(void* self);
void* QAnyStringView_back(void* self);
int QAnyStringView_empty(void* self);
long long QAnyStringView_size_bytes(void* self);
long long QAnyStringView_max_size(void* self);
int QAnyStringView_isNull(void* self);
int QAnyStringView_isEmpty(void* self);
long long QAnyStringView_length(void* self);
void* QAnyStringView_OperatorAssign(void* self, const char* arg1, long long arg1_len);
void* QAnyStringView_mid2(void* self, long long pos, long long n);
int QAnyStringView_compare2(const char* lhs, long long lhs_len, const char* rhs, long long rhs_len, int cs);
int QAnyStringView_OperatorEqual(void* self, const char* rhs, long long rhs_len);
int QAnyStringView_OperatorNotEqual(void* self, const char* rhs, long long rhs_len);
int QAnyStringView_OperatorLesser(void* self, const char* rhs, long long rhs_len);
int QAnyStringView_OperatorGreater(void* self, const char* rhs, long long rhs_len);
int QAnyStringView_OperatorLesserOrEqual(void* self, const char* rhs, long long rhs_len);
int QAnyStringView_OperatorGreaterOrEqual(void* self, const char* rhs, long long rhs_len);
void QAnyStringView_delete(void* self);

#ifdef __cplusplus
}
#endif
