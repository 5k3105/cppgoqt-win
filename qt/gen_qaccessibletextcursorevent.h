#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextCursorEvent_base(void* self);
void* QAccessibleTextCursorEvent_new(void* obj, int cursorPos);
void* QAccessibleTextCursorEvent_new2(void* iface, int cursorPos);
void QAccessibleTextCursorEvent_setCursorPosition(void* self, int position);
int QAccessibleTextCursorEvent_cursorPosition(void* self);
void QAccessibleTextCursorEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
