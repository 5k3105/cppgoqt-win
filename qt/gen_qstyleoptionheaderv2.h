#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionHeaderV2_base(void* self);
void QStyleOptionHeaderV2_chain(void* self, void** out);
void* QStyleOptionHeaderV2_new();
void* QStyleOptionHeaderV2_new2(void* other);
void* QStyleOptionHeaderV2_OperatorAssign(void* self, void* arg1);
int QStyleOptionHeaderV2_textElideMode_get(void* self);
void QStyleOptionHeaderV2_textElideMode_set(void* self, int textElideMode);
int QStyleOptionHeaderV2_isSectionDragTarget_get(void* self);
void QStyleOptionHeaderV2_isSectionDragTarget_set(void* self, int isSectionDragTarget);
int QStyleOptionHeaderV2_unused_get(void* self);
void QStyleOptionHeaderV2_unused_set(void* self, int unused);
void QStyleOptionHeaderV2_delete(void* self);

#ifdef __cplusplus
}
#endif
