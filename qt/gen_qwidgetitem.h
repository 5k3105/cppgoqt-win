#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWidgetItem_base(void* self);
void* QWidgetItem_new(void* w);
void QWidgetItem_delete(void* self);

#ifdef __cplusplus
}
#endif
