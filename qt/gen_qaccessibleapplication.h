#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleApplication_base(void* self);
void QAccessibleApplication_chain(void* self, void** out);
void QAccessibleApplication__set_override(void* self, int idx);
void* QAccessibleApplication_new(long long __goid);
void* QAccessibleApplication_window_default(void* self);
int QAccessibleApplication_childCount_default(void* self);
int QAccessibleApplication_indexOfChild_default(void* self, void* arg1);
void* QAccessibleApplication_focusChild_default(void* self);
void* QAccessibleApplication_parent_default(void* self);
void* QAccessibleApplication_child_default(void* self, int index);
char* QAccessibleApplication_text_default(void* self, int t, long long* __out_len);
int QAccessibleApplication_role_default(void* self);
void* QAccessibleApplication_state_default(void* self);
int QAccessibleApplication_isValid_default(void* self);
void* QAccessibleApplication_object_default(void* self);
void* QAccessibleApplication_rect_default(void* self);
void QAccessibleApplication_setText_default(void* self, int t, const char* text, long long text_len);
void* QAccessibleApplication_childAt_default(void* self, int x, int y);
void* QAccessibleApplication_foregroundColor_default(void* self);
void* QAccessibleApplication_backgroundColor_default(void* self);
void* QAccessibleApplication__Go_new(long long id);
void QAccessibleApplication_delete(void* self);

#ifdef __cplusplus
}
#endif
