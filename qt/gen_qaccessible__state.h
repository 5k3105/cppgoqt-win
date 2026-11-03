#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessible__State_new();
int QAccessible__State_OperatorEqual(void* self, void* second);
unsigned long long QAccessible__State_disabled_get(void* self);
void QAccessible__State_disabled_set(void* self, unsigned long long disabled);
unsigned long long QAccessible__State_selected_get(void* self);
void QAccessible__State_selected_set(void* self, unsigned long long selected);
unsigned long long QAccessible__State_focusable_get(void* self);
void QAccessible__State_focusable_set(void* self, unsigned long long focusable);
unsigned long long QAccessible__State_focused_get(void* self);
void QAccessible__State_focused_set(void* self, unsigned long long focused);
unsigned long long QAccessible__State_pressed_get(void* self);
void QAccessible__State_pressed_set(void* self, unsigned long long pressed);
unsigned long long QAccessible__State_checkable_get(void* self);
void QAccessible__State_checkable_set(void* self, unsigned long long checkable);
unsigned long long QAccessible__State_checked_get(void* self);
void QAccessible__State_checked_set(void* self, unsigned long long checked);
unsigned long long QAccessible__State_checkStateMixed_get(void* self);
void QAccessible__State_checkStateMixed_set(void* self, unsigned long long checkStateMixed);
unsigned long long QAccessible__State_readOnly_get(void* self);
void QAccessible__State_readOnly_set(void* self, unsigned long long readOnly);
unsigned long long QAccessible__State_hotTracked_get(void* self);
void QAccessible__State_hotTracked_set(void* self, unsigned long long hotTracked);
unsigned long long QAccessible__State_defaultButton_get(void* self);
void QAccessible__State_defaultButton_set(void* self, unsigned long long defaultButton);
unsigned long long QAccessible__State_expanded_get(void* self);
void QAccessible__State_expanded_set(void* self, unsigned long long expanded);
unsigned long long QAccessible__State_collapsed_get(void* self);
void QAccessible__State_collapsed_set(void* self, unsigned long long collapsed);
unsigned long long QAccessible__State_busy_get(void* self);
void QAccessible__State_busy_set(void* self, unsigned long long busy);
unsigned long long QAccessible__State_expandable_get(void* self);
void QAccessible__State_expandable_set(void* self, unsigned long long expandable);
unsigned long long QAccessible__State_marqueed_get(void* self);
void QAccessible__State_marqueed_set(void* self, unsigned long long marqueed);
unsigned long long QAccessible__State_animated_get(void* self);
void QAccessible__State_animated_set(void* self, unsigned long long animated);
unsigned long long QAccessible__State_invisible_get(void* self);
void QAccessible__State_invisible_set(void* self, unsigned long long invisible);
unsigned long long QAccessible__State_offscreen_get(void* self);
void QAccessible__State_offscreen_set(void* self, unsigned long long offscreen);
unsigned long long QAccessible__State_sizeable_get(void* self);
void QAccessible__State_sizeable_set(void* self, unsigned long long sizeable);
unsigned long long QAccessible__State_movable_get(void* self);
void QAccessible__State_movable_set(void* self, unsigned long long movable);
unsigned long long QAccessible__State_selfVoicing_get(void* self);
void QAccessible__State_selfVoicing_set(void* self, unsigned long long selfVoicing);
unsigned long long QAccessible__State_selectable_get(void* self);
void QAccessible__State_selectable_set(void* self, unsigned long long selectable);
unsigned long long QAccessible__State_linked_get(void* self);
void QAccessible__State_linked_set(void* self, unsigned long long linked);
unsigned long long QAccessible__State_traversed_get(void* self);
void QAccessible__State_traversed_set(void* self, unsigned long long traversed);
unsigned long long QAccessible__State_multiSelectable_get(void* self);
void QAccessible__State_multiSelectable_set(void* self, unsigned long long multiSelectable);
unsigned long long QAccessible__State_extSelectable_get(void* self);
void QAccessible__State_extSelectable_set(void* self, unsigned long long extSelectable);
unsigned long long QAccessible__State_passwordEdit_get(void* self);
void QAccessible__State_passwordEdit_set(void* self, unsigned long long passwordEdit);
unsigned long long QAccessible__State_hasPopup_get(void* self);
void QAccessible__State_hasPopup_set(void* self, unsigned long long hasPopup);
unsigned long long QAccessible__State_modal_get(void* self);
void QAccessible__State_modal_set(void* self, unsigned long long modal);
unsigned long long QAccessible__State_active_get(void* self);
void QAccessible__State_active_set(void* self, unsigned long long active);
unsigned long long QAccessible__State_invalid_get(void* self);
void QAccessible__State_invalid_set(void* self, unsigned long long invalid);
unsigned long long QAccessible__State_editable_get(void* self);
void QAccessible__State_editable_set(void* self, unsigned long long editable);
unsigned long long QAccessible__State_multiLine_get(void* self);
void QAccessible__State_multiLine_set(void* self, unsigned long long multiLine);
unsigned long long QAccessible__State_selectableText_get(void* self);
void QAccessible__State_selectableText_set(void* self, unsigned long long selectableText);
unsigned long long QAccessible__State_supportsAutoCompletion_get(void* self);
void QAccessible__State_supportsAutoCompletion_set(void* self, unsigned long long supportsAutoCompletion);
unsigned long long QAccessible__State_searchEdit_get(void* self);
void QAccessible__State_searchEdit_set(void* self, unsigned long long searchEdit);
unsigned long long QAccessible__State_qt_reserved_get(void* self);
void QAccessible__State_qt_reserved_set(void* self, unsigned long long qt_reserved);
void QAccessible__State_delete(void* self);

#ifdef __cplusplus
}
#endif
