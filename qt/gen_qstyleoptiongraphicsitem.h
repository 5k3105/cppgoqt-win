#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionGraphicsItem_base(void* self);
void* QStyleOptionGraphicsItem_new();
void* QStyleOptionGraphicsItem_new2(void* other);
void* QStyleOptionGraphicsItem_OperatorAssign(void* self, void* arg1);
double QStyleOptionGraphicsItem_levelOfDetailFromTransform(void* worldTransform);
void* QStyleOptionGraphicsItem_exposedRect_get(void* self);
void QStyleOptionGraphicsItem_exposedRect_set(void* self, void* exposedRect);
void QStyleOptionGraphicsItem_delete(void* self);

#ifdef __cplusplus
}
#endif
