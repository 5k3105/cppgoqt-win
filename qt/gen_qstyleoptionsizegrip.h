#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionSizeGrip_base(void* self);
void QStyleOptionSizeGrip_chain(void* self, void** out);
void* QStyleOptionSizeGrip_new();
void* QStyleOptionSizeGrip_new2(void* other);
void* QStyleOptionSizeGrip_OperatorAssign(void* self, void* arg1);
int QStyleOptionSizeGrip_corner_get(void* self);
void QStyleOptionSizeGrip_corner_set(void* self, int corner);
void QStyleOptionSizeGrip_delete(void* self);

#ifdef __cplusplus
}
#endif
