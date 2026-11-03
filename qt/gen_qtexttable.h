#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextTable_base(void* self);
void QTextTable_chain(void* self, void** out);
void* QTextTable_fromQObject(void* o);
void* QTextTable_new(void* doc);
char* QTextTable_tr(const char* s, long long* __out_len);
void QTextTable_resize(void* self, int rows, int cols);
void QTextTable_insertRows(void* self, int pos, int num);
void QTextTable_insertColumns(void* self, int pos, int num);
void QTextTable_appendRows(void* self, int count);
void QTextTable_appendColumns(void* self, int count);
void QTextTable_removeRows(void* self, int pos, int num);
void QTextTable_removeColumns(void* self, int pos, int num);
void QTextTable_mergeCells(void* self, int row, int col, int numRows, int numCols);
void QTextTable_mergeCells2(void* self, void* cursor);
void QTextTable_splitCell(void* self, int row, int col, int numRows, int numCols);
int QTextTable_rows(void* self);
int QTextTable_columns(void* self);
void* QTextTable_cellAt(void* self, int row, int col);
void* QTextTable_cellAt2(void* self, int position);
void* QTextTable_cellAt3(void* self, void* c);
void* QTextTable_rowStart(void* self, void* c);
void* QTextTable_rowEnd(void* self, void* c);
void QTextTable_setFormat(void* self, void* format);
void* QTextTable_format(void* self);
void* QTextTable_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTextTable_connect2(void* sender, void* signal, void* receiver, void* method);
char* QTextTable_tr2(const char* s, const char* c, long long* __out_len);
char* QTextTable_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTextTable_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTextTable_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QTextTable_connect_destroyed(void* self, long long id);
void QTextTable_connect_destroyed2(void* self, long long id);
void QTextTable_connect_objectNameChanged(void* self, long long id);
void QTextTable__disconnect(long long id);
void QTextTable_delete(void* self);

#ifdef __cplusplus
}
#endif
