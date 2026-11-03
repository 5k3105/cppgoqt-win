#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWidgetData_new(void* arg1);
void* QWidgetData_OperatorAssign(void* self, void* arg1);
unsigned long long QWidgetData_winid_get(void* self);
void QWidgetData_winid_set(void* self, unsigned long long winid);
unsigned int QWidgetData_widget_attributes_get(void* self);
void QWidgetData_widget_attributes_set(void* self, unsigned int widget_attributes);
int QWidgetData_window_flags_get(void* self);
void QWidgetData_window_flags_set(void* self, int window_flags);
unsigned int QWidgetData_window_state_get(void* self);
void QWidgetData_window_state_set(void* self, unsigned int window_state);
unsigned int QWidgetData_focus_policy_get(void* self);
void QWidgetData_focus_policy_set(void* self, unsigned int focus_policy);
unsigned int QWidgetData_sizehint_forced_get(void* self);
void QWidgetData_sizehint_forced_set(void* self, unsigned int sizehint_forced);
unsigned int QWidgetData_is_closing_get(void* self);
void QWidgetData_is_closing_set(void* self, unsigned int is_closing);
unsigned int QWidgetData_in_show_get(void* self);
void QWidgetData_in_show_set(void* self, unsigned int in_show);
unsigned int QWidgetData_in_set_window_state_get(void* self);
void QWidgetData_in_set_window_state_set(void* self, unsigned int in_set_window_state);
unsigned int QWidgetData_fstrut_dirty_get(void* self);
void QWidgetData_fstrut_dirty_set(void* self, unsigned int fstrut_dirty);
unsigned int QWidgetData_context_menu_policy_get(void* self);
void QWidgetData_context_menu_policy_set(void* self, unsigned int context_menu_policy);
unsigned int QWidgetData_window_modality_get(void* self);
void QWidgetData_window_modality_set(void* self, unsigned int window_modality);
unsigned int QWidgetData_in_destructor_get(void* self);
void QWidgetData_in_destructor_set(void* self, unsigned int in_destructor);
unsigned int QWidgetData_unused_get(void* self);
void QWidgetData_unused_set(void* self, unsigned int unused);
void* QWidgetData_crect_get(void* self);
void QWidgetData_crect_set(void* self, void* crect);
void* QWidgetData_pal_get(void* self);
void QWidgetData_pal_set(void* self, void* pal);
void* QWidgetData_fnt_get(void* self);
void QWidgetData_fnt_set(void* self, void* fnt);
void* QWidgetData_wrect_get(void* self);
void QWidgetData_wrect_set(void* self, void* wrect);
void QWidgetData_delete(void* self);

#ifdef __cplusplus
}
#endif
