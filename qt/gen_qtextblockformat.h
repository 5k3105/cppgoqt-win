#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextBlockFormat_base(void* self);
void* QTextBlockFormat_new();
void* QTextBlockFormat_new2(void* arg1);
int QTextBlockFormat_isValid(void* self);
void QTextBlockFormat_setAlignment(void* self, int alignment);
int QTextBlockFormat_alignment(void* self);
void QTextBlockFormat_setTopMargin(void* self, double margin);
double QTextBlockFormat_topMargin(void* self);
void QTextBlockFormat_setBottomMargin(void* self, double margin);
double QTextBlockFormat_bottomMargin(void* self);
void QTextBlockFormat_setLeftMargin(void* self, double margin);
double QTextBlockFormat_leftMargin(void* self);
void QTextBlockFormat_setRightMargin(void* self, double margin);
double QTextBlockFormat_rightMargin(void* self);
void QTextBlockFormat_setTextIndent(void* self, double aindent);
double QTextBlockFormat_textIndent(void* self);
void QTextBlockFormat_setIndent(void* self, int indent);
int QTextBlockFormat_indent(void* self);
void QTextBlockFormat_setHeadingLevel(void* self, int alevel);
int QTextBlockFormat_headingLevel(void* self);
void QTextBlockFormat_setLineHeight(void* self, double height, int heightType);
double QTextBlockFormat_lineHeight(void* self, double scriptLineHeight, double scaling);
double QTextBlockFormat_lineHeight2(void* self);
int QTextBlockFormat_lineHeightType(void* self);
void QTextBlockFormat_setNonBreakableLines(void* self, int b);
int QTextBlockFormat_nonBreakableLines(void* self);
void QTextBlockFormat_setPageBreakPolicy(void* self, int flags);
int QTextBlockFormat_pageBreakPolicy(void* self);
void QTextBlockFormat_setTabPositions(void* self, void** tabs, long long tabs_len);
void** QTextBlockFormat_tabPositions(void* self, long long* __out_len);
void QTextBlockFormat_setMarker(void* self, int marker);
int QTextBlockFormat_marker(void* self);
void* QTextBlockFormat_OperatorAssign(void* self, void* arg1);
void QTextBlockFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
