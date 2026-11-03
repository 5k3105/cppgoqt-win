#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFileIconProvider_base(void* self);
void QFileIconProvider__set_override(void* self, int idx);
void* QFileIconProvider_new(long long __goid);
void* QFileIconProvider_icon_default(void* self, int type_);
char* QFileIconProvider_type_default(void* self, void* arg1, long long* __out_len);
void QFileIconProvider_setOptions_default(void* self, int arg1);
int QFileIconProvider_options_default(void* self);
void* QFileIconProvider__Go_new(long long id);
void QFileIconProvider_delete(void* self);

#ifdef __cplusplus
}
#endif
