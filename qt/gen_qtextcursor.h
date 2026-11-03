#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextCursor_new();
void* QTextCursor_new2(void* document);
void* QTextCursor_new3(void* frame);
void* QTextCursor_new4(void* block);
void* QTextCursor_new5(void* cursor);
void* QTextCursor_OperatorAssign(void* self, void* other);
void QTextCursor_swap(void* self, void* other);
int QTextCursor_isNull(void* self);
void QTextCursor_setPosition(void* self, int pos);
int QTextCursor_position(void* self);
int QTextCursor_positionInBlock(void* self);
int QTextCursor_anchor(void* self);
void QTextCursor_insertText(void* self, const char* text, long long text_len);
void QTextCursor_insertText2(void* self, const char* text, long long text_len, void* format);
int QTextCursor_movePosition(void* self, int op);
int QTextCursor_visualNavigation(void* self);
void QTextCursor_setVisualNavigation(void* self, int b);
void QTextCursor_setVerticalMovementX(void* self, int x);
int QTextCursor_verticalMovementX(void* self);
void QTextCursor_setKeepPositionOnInsert(void* self, int b);
int QTextCursor_keepPositionOnInsert(void* self);
void QTextCursor_deleteChar(void* self);
void QTextCursor_deletePreviousChar(void* self);
void QTextCursor_select(void* self, int selection);
int QTextCursor_hasSelection(void* self);
int QTextCursor_hasComplexSelection(void* self);
void QTextCursor_removeSelectedText(void* self);
void QTextCursor_clearSelection(void* self);
int QTextCursor_selectionStart(void* self);
int QTextCursor_selectionEnd(void* self);
char* QTextCursor_selectedText(void* self, long long* __out_len);
void* QTextCursor_selection(void* self);
void QTextCursor_selectedTableCells(void* self, int* firstRow, int* numRows, int* firstColumn, int* numColumns);
void* QTextCursor_block(void* self);
void* QTextCursor_charFormat(void* self);
void QTextCursor_setCharFormat(void* self, void* format);
void QTextCursor_mergeCharFormat(void* self, void* modifier);
void* QTextCursor_blockFormat(void* self);
void QTextCursor_setBlockFormat(void* self, void* format);
void QTextCursor_mergeBlockFormat(void* self, void* modifier);
void* QTextCursor_blockCharFormat(void* self);
void QTextCursor_setBlockCharFormat(void* self, void* format);
void QTextCursor_mergeBlockCharFormat(void* self, void* modifier);
int QTextCursor_atBlockStart(void* self);
int QTextCursor_atBlockEnd(void* self);
int QTextCursor_atStart(void* self);
int QTextCursor_atEnd(void* self);
void QTextCursor_insertBlock(void* self);
void QTextCursor_insertBlock2(void* self, void* format);
void QTextCursor_insertBlock3(void* self, void* format, void* charFormat);
void* QTextCursor_insertList(void* self, void* format);
void* QTextCursor_insertList2(void* self, int style);
void* QTextCursor_createList(void* self, void* format);
void* QTextCursor_createList2(void* self, int style);
void* QTextCursor_currentList(void* self);
void* QTextCursor_insertTable(void* self, int rows, int cols, void* format);
void* QTextCursor_insertTable2(void* self, int rows, int cols);
void* QTextCursor_currentTable(void* self);
void* QTextCursor_insertFrame(void* self, void* format);
void* QTextCursor_currentFrame(void* self);
void QTextCursor_insertFragment(void* self, void* fragment);
void QTextCursor_insertHtml(void* self, const char* html, long long html_len);
void QTextCursor_insertMarkdown(void* self, const char* markdown, long long markdown_len);
void QTextCursor_insertImage(void* self, void* format, int alignment);
void QTextCursor_insertImage2(void* self, void* format);
void QTextCursor_insertImage3(void* self, const char* name, long long name_len);
void QTextCursor_insertImage4(void* self, void* image);
void QTextCursor_beginEditBlock(void* self);
void QTextCursor_joinPreviousEditBlock(void* self);
void QTextCursor_endEditBlock(void* self);
int QTextCursor_OperatorNotEqual(void* self, void* rhs);
int QTextCursor_OperatorLesser(void* self, void* rhs);
int QTextCursor_OperatorLesserOrEqual(void* self, void* rhs);
int QTextCursor_OperatorEqual(void* self, void* rhs);
int QTextCursor_OperatorGreaterOrEqual(void* self, void* rhs);
int QTextCursor_OperatorGreater(void* self, void* rhs);
int QTextCursor_isCopyOf(void* self, void* other);
int QTextCursor_blockNumber(void* self);
int QTextCursor_columnNumber(void* self);
void* QTextCursor_document(void* self);
void QTextCursor_setPosition2(void* self, int pos, int mode);
int QTextCursor_movePosition2(void* self, int op, int arg2);
int QTextCursor_movePosition3(void* self, int op, int arg2, int n);
void QTextCursor_insertMarkdown2(void* self, const char* markdown, long long markdown_len, int features);
void QTextCursor_insertImage5(void* self, void* image, const char* name, long long name_len);
void QTextCursor_delete(void* self);

#ifdef __cplusplus
}
#endif
