#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWidgetItemV2_base(void* self);
void QWidgetItemV2_chain(void* self, void** out);
void* QWidgetItemV2_new(void* widget);
void QWidgetItemV2_delete(void* self);

#ifdef __cplusplus
}
#endif
