#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractTextDocumentLayout__Selection_new(void* arg1);
void* QAbstractTextDocumentLayout__Selection_new2();
void* QAbstractTextDocumentLayout__Selection_OperatorAssign(void* self, void* arg1);
void* QAbstractTextDocumentLayout__Selection_cursor_get(void* self);
void QAbstractTextDocumentLayout__Selection_cursor_set(void* self, void* cursor);
void* QAbstractTextDocumentLayout__Selection_format_get(void* self);
void QAbstractTextDocumentLayout__Selection_format_set(void* self, void* format);
void QAbstractTextDocumentLayout__Selection_delete(void* self);

#ifdef __cplusplus
}
#endif
