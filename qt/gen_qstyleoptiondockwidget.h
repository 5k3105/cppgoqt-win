#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionDockWidget_base(void* self);
void* QStyleOptionDockWidget_new();
void* QStyleOptionDockWidget_new2(void* other);
void* QStyleOptionDockWidget_OperatorAssign(void* self, void* arg1);
char* QStyleOptionDockWidget_title_get(void* self, long long* __out_len);
void QStyleOptionDockWidget_title_set(void* self, const char* title, long long title_len);
int QStyleOptionDockWidget_closable_get(void* self);
void QStyleOptionDockWidget_closable_set(void* self, int closable);
int QStyleOptionDockWidget_movable_get(void* self);
void QStyleOptionDockWidget_movable_set(void* self, int movable);
int QStyleOptionDockWidget_floatable_get(void* self);
void QStyleOptionDockWidget_floatable_set(void* self, int floatable);
int QStyleOptionDockWidget_verticalTitleBar_get(void* self);
void QStyleOptionDockWidget_verticalTitleBar_set(void* self, int verticalTitleBar);
void QStyleOptionDockWidget_delete(void* self);

#ifdef __cplusplus
}
#endif
