#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

char* QCborParserError_errorString(void* self, long long* __out_len);
long long QCborParserError_offset_get(void* self);
void QCborParserError_offset_set(void* self, long long offset);
void* QCborParserError_error_get(void* self);
void QCborParserError_error_set(void* self, void* error_);
void QCborParserError_delete(void* self);

#ifdef __cplusplus
}
#endif
