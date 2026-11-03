#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleImageInterface__set_override(void* self, int idx);
char* QAccessibleImageInterface_imageDescription(void* self, long long* __out_len);
void* QAccessibleImageInterface_imageSize(void* self);
void* QAccessibleImageInterface_imagePosition(void* self);
void* QAccessibleImageInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleImageInterface__Go_new(long long id);
void QAccessibleImageInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
