#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionTabBarBase_base(void* self);
void* QStyleOptionTabBarBase_new();
void* QStyleOptionTabBarBase_new2(void* other);
void* QStyleOptionTabBarBase_OperatorAssign(void* self, void* arg1);
int QStyleOptionTabBarBase_shape_get(void* self);
void QStyleOptionTabBarBase_shape_set(void* self, int shape);
void* QStyleOptionTabBarBase_tabBarRect_get(void* self);
void QStyleOptionTabBarBase_tabBarRect_set(void* self, void* tabBarRect);
void* QStyleOptionTabBarBase_selectedTabRect_get(void* self);
void QStyleOptionTabBarBase_selectedTabRect_set(void* self, void* selectedTabRect);
int QStyleOptionTabBarBase_documentMode_get(void* self);
void QStyleOptionTabBarBase_documentMode_set(void* self, int documentMode);
void QStyleOptionTabBarBase_delete(void* self);

#ifdef __cplusplus
}
#endif
