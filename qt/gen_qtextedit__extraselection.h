#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextEdit__ExtraSelection_new(void* arg1);
void* QTextEdit__ExtraSelection_OperatorAssign(void* self, void* arg1);
void* QTextEdit__ExtraSelection_cursor_get(void* self);
void QTextEdit__ExtraSelection_cursor_set(void* self, void* cursor);
void* QTextEdit__ExtraSelection_format_get(void* self);
void QTextEdit__ExtraSelection_format_set(void* self, void* format);
void QTextEdit__ExtraSelection_delete(void* self);

#ifdef __cplusplus
}
#endif
