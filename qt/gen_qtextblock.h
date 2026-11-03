#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextBlock_new();
void* QTextBlock_new2(void* o);
void* QTextBlock_OperatorAssign(void* self, void* o);
int QTextBlock_isValid(void* self);
int QTextBlock_OperatorEqual(void* self, void* o);
int QTextBlock_OperatorNotEqual(void* self, void* o);
int QTextBlock_OperatorLesser(void* self, void* o);
int QTextBlock_position(void* self);
int QTextBlock_length(void* self);
int QTextBlock_contains(void* self, int position);
void* QTextBlock_layout(void* self);
void QTextBlock_clearLayout(void* self);
void* QTextBlock_blockFormat(void* self);
int QTextBlock_blockFormatIndex(void* self);
void* QTextBlock_charFormat(void* self);
int QTextBlock_charFormatIndex(void* self);
int QTextBlock_textDirection(void* self);
char* QTextBlock_text(void* self, long long* __out_len);
void** QTextBlock_textFormats(void* self, long long* __out_len);
void* QTextBlock_document(void* self);
void* QTextBlock_textList(void* self);
void* QTextBlock_userData(void* self);
void QTextBlock_setUserData(void* self, void* data);
int QTextBlock_userState(void* self);
void QTextBlock_setUserState(void* self, int state);
int QTextBlock_revision(void* self);
void QTextBlock_setRevision(void* self, int rev);
int QTextBlock_isVisible(void* self);
void QTextBlock_setVisible(void* self, int visible);
int QTextBlock_blockNumber(void* self);
int QTextBlock_firstLineNumber(void* self);
void QTextBlock_setLineCount(void* self, int count);
int QTextBlock_lineCount(void* self);
void* QTextBlock_next(void* self);
void* QTextBlock_previous(void* self);
int QTextBlock_fragmentIndex(void* self);
void QTextBlock_delete(void* self);

#ifdef __cplusplus
}
#endif
