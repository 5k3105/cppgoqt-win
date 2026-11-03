#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPageSize_new();
void* QPageSize_new2(int pageSizeId);
void* QPageSize_new3(void* pointSize);
void* QPageSize_new4(void* size, int units);
void* QPageSize_new5(void* other);
void* QPageSize_new6(void* pointSize, const char* name, long long name_len);
void* QPageSize_new7(void* pointSize, const char* name, long long name_len, int matchPolicy);
void* QPageSize_new8(void* size, int units, const char* name, long long name_len);
void* QPageSize_new9(void* size, int units, const char* name, long long name_len, int matchPolicy);
void* QPageSize_OperatorAssign(void* self, void* other);
void QPageSize_swap(void* self, void* other);
int QPageSize_isEquivalentTo(void* self, void* other);
int QPageSize_isValid(void* self);
char* QPageSize_key(void* self, long long* __out_len);
char* QPageSize_name(void* self, long long* __out_len);
int QPageSize_id(void* self);
int QPageSize_windowsId(void* self);
void* QPageSize_definitionSize(void* self);
int QPageSize_definitionUnits(void* self);
void* QPageSize_size(void* self, int units);
void* QPageSize_sizePoints(void* self);
void* QPageSize_sizePixels(void* self, int resolution);
void* QPageSize_rect(void* self, int units);
void* QPageSize_rectPoints(void* self);
void* QPageSize_rectPixels(void* self, int resolution);
char* QPageSize_key2(int pageSizeId, long long* __out_len);
char* QPageSize_name2(int pageSizeId, long long* __out_len);
int QPageSize_id2(void* pointSize);
int QPageSize_id3(void* size, int units);
int QPageSize_id4(int windowsId);
int QPageSize_windowsId2(int pageSizeId);
void* QPageSize_definitionSize2(int pageSizeId);
int QPageSize_definitionUnits2(int pageSizeId);
void* QPageSize_size2(int pageSizeId, int units);
void* QPageSize_sizePoints2(int pageSizeId);
void* QPageSize_sizePixels2(int pageSizeId, int resolution);
int QPageSize_id5(void* pointSize, int matchPolicy);
int QPageSize_id6(void* size, int units, int matchPolicy);
int QPageSize_OperatorEqual(void* self, void* rhs);
int QPageSize_OperatorNotEqual(void* self, void* rhs);
void QPageSize_delete(void* self);

#ifdef __cplusplus
}
#endif
