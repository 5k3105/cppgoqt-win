#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleTextInterface__set_override(void* self, int idx);
void QAccessibleTextInterface_selection(void* self, int selectionIndex, int* startOffset, int* endOffset);
int QAccessibleTextInterface_selectionCount(void* self);
void QAccessibleTextInterface_addSelection(void* self, int startOffset, int endOffset);
void QAccessibleTextInterface_removeSelection(void* self, int selectionIndex);
void QAccessibleTextInterface_setSelection(void* self, int selectionIndex, int startOffset, int endOffset);
int QAccessibleTextInterface_cursorPosition(void* self);
void QAccessibleTextInterface_setCursorPosition(void* self, int position);
char* QAccessibleTextInterface_text(void* self, int startOffset, int endOffset, long long* __out_len);
char* QAccessibleTextInterface_textBeforeOffset(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len);
char* QAccessibleTextInterface_textAfterOffset(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len);
char* QAccessibleTextInterface_textAtOffset(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len);
int QAccessibleTextInterface_characterCount(void* self);
void* QAccessibleTextInterface_characterRect(void* self, int offset);
int QAccessibleTextInterface_offsetAtPoint(void* self, void* point);
void QAccessibleTextInterface_scrollToSubstring(void* self, int startIndex, int endIndex);
char* QAccessibleTextInterface_attributes(void* self, int offset, int* startOffset, int* endOffset, long long* __out_len);
void* QAccessibleTextInterface_OperatorAssign(void* self, void* arg1);
char* QAccessibleTextInterface_textBeforeOffset_default(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len);
char* QAccessibleTextInterface_textAfterOffset_default(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len);
char* QAccessibleTextInterface_textAtOffset_default(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len);
void* QAccessibleTextInterface__Go_new(long long id);
void QAccessibleTextInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
