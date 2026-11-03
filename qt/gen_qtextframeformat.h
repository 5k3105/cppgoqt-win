#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextFrameFormat_base(void* self);
void* QTextFrameFormat_new();
void* QTextFrameFormat_new2(void* arg1);
int QTextFrameFormat_isValid(void* self);
void QTextFrameFormat_setPosition(void* self, int f);
int QTextFrameFormat_position(void* self);
void QTextFrameFormat_setBorder(void* self, double border);
double QTextFrameFormat_border(void* self);
void QTextFrameFormat_setBorderBrush(void* self, void* brush);
void* QTextFrameFormat_borderBrush(void* self);
void QTextFrameFormat_setBorderStyle(void* self, int style);
int QTextFrameFormat_borderStyle(void* self);
void QTextFrameFormat_setMargin(void* self, double margin);
double QTextFrameFormat_margin(void* self);
void QTextFrameFormat_setTopMargin(void* self, double margin);
double QTextFrameFormat_topMargin(void* self);
void QTextFrameFormat_setBottomMargin(void* self, double margin);
double QTextFrameFormat_bottomMargin(void* self);
void QTextFrameFormat_setLeftMargin(void* self, double margin);
double QTextFrameFormat_leftMargin(void* self);
void QTextFrameFormat_setRightMargin(void* self, double margin);
double QTextFrameFormat_rightMargin(void* self);
void QTextFrameFormat_setPadding(void* self, double padding);
double QTextFrameFormat_padding(void* self);
void QTextFrameFormat_setWidth(void* self, double width);
void QTextFrameFormat_setWidth2(void* self, void* length);
void* QTextFrameFormat_width(void* self);
void QTextFrameFormat_setHeight(void* self, double height);
void QTextFrameFormat_setHeight2(void* self, void* height);
void* QTextFrameFormat_height(void* self);
void QTextFrameFormat_setPageBreakPolicy(void* self, int flags);
int QTextFrameFormat_pageBreakPolicy(void* self);
void* QTextFrameFormat_OperatorAssign(void* self, void* arg1);
void QTextFrameFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
