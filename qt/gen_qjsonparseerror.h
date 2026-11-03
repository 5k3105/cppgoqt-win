#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

char* QJsonParseError_errorString(void* self, long long* __out_len);
int QJsonParseError_offset_get(void* self);
void QJsonParseError_offset_set(void* self, int offset);
int QJsonParseError_error_get(void* self);
void QJsonParseError_error_set(void* self, int error_);
void QJsonParseError_delete(void* self);

#ifdef __cplusplus
}
#endif
