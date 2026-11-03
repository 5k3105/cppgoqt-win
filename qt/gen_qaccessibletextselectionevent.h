#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextSelectionEvent_base(void* self);
void QAccessibleTextSelectionEvent_chain(void* self, void** out);
void* QAccessibleTextSelectionEvent_new(void* obj, int start, int end);
void* QAccessibleTextSelectionEvent_new2(void* iface, int start, int end);
void QAccessibleTextSelectionEvent_setSelection(void* self, int start, int end);
int QAccessibleTextSelectionEvent_selectionStart(void* self);
int QAccessibleTextSelectionEvent_selectionEnd(void* self);
void QAccessibleTextSelectionEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
