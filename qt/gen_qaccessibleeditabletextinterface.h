#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleEditableTextInterface__set_override(void* self, int idx);
void QAccessibleEditableTextInterface_deleteText(void* self, int startOffset, int endOffset);
void QAccessibleEditableTextInterface_insertText(void* self, int offset, const char* text, long long text_len);
void QAccessibleEditableTextInterface_replaceText(void* self, int startOffset, int endOffset, const char* text, long long text_len);
void* QAccessibleEditableTextInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleEditableTextInterface__Go_new(long long id);
void QAccessibleEditableTextInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
