#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAbstractFileIconProvider__set_override(void* self, int idx);
void* QAbstractFileIconProvider_new(long long __goid);
void* QAbstractFileIconProvider_icon(void* self, int arg1);
void* QAbstractFileIconProvider_icon2(void* self, void* arg1);
char* QAbstractFileIconProvider_type(void* self, void* arg1, long long* __out_len);
void QAbstractFileIconProvider_setOptions(void* self, int arg1);
int QAbstractFileIconProvider_options(void* self);
void* QAbstractFileIconProvider_icon_default(void* self, int arg1);
void* QAbstractFileIconProvider_icon_default2(void* self, void* arg1);
char* QAbstractFileIconProvider_type_default(void* self, void* arg1, long long* __out_len);
void QAbstractFileIconProvider_setOptions_default(void* self, int arg1);
int QAbstractFileIconProvider_options_default(void* self);
void* QAbstractFileIconProvider__Go_new(long long id);
void QAbstractFileIconProvider_delete(void* self);

#ifdef __cplusplus
}
#endif
