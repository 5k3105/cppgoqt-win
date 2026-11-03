#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionMenuItemV2_base(void* self);
void QStyleOptionMenuItemV2_chain(void* self, void** out);
void* QStyleOptionMenuItemV2_new();
void* QStyleOptionMenuItemV2_new2(void* other);
void* QStyleOptionMenuItemV2_OperatorAssign(void* self, void* arg1);
int QStyleOptionMenuItemV2_mouseDown_get(void* self);
void QStyleOptionMenuItemV2_mouseDown_set(void* self, int mouseDown);
int QStyleOptionMenuItemV2_unused_get(void* self);
void QStyleOptionMenuItemV2_unused_set(void* self, int unused);
void QStyleOptionMenuItemV2_delete(void* self);

#ifdef __cplusplus
}
#endif
