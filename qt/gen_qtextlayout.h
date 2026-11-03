#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextLayout_new();
void* QTextLayout_new2(const char* text, long long text_len);
void* QTextLayout_new3(const char* text, long long text_len, void* font);
void* QTextLayout_new4(void* b);
void* QTextLayout_new5(const char* text, long long text_len, void* font, void* paintdevice);
void QTextLayout_setFont(void* self, void* f);
void* QTextLayout_font(void* self);
void QTextLayout_setRawFont(void* self, void* rawFont);
void QTextLayout_setText(void* self, const char* string_, long long string__len);
char* QTextLayout_text(void* self, long long* __out_len);
void QTextLayout_setTextOption(void* self, void* option);
void* QTextLayout_textOption(void* self);
void QTextLayout_setPreeditArea(void* self, int position, const char* text, long long text_len);
int QTextLayout_preeditAreaPosition(void* self);
char* QTextLayout_preeditAreaText(void* self, long long* __out_len);
void QTextLayout_setFormats(void* self, void** overrides, long long overrides_len);
void** QTextLayout_formats(void* self, long long* __out_len);
void QTextLayout_clearFormats(void* self);
void QTextLayout_setCacheEnabled(void* self, int enable);
int QTextLayout_cacheEnabled(void* self);
void QTextLayout_setCursorMoveStyle(void* self, int style);
int QTextLayout_cursorMoveStyle(void* self);
void QTextLayout_beginLayout(void* self);
void QTextLayout_endLayout(void* self);
void QTextLayout_clearLayout(void* self);
void* QTextLayout_createLine(void* self);
int QTextLayout_lineCount(void* self);
void* QTextLayout_lineAt(void* self, int i);
void* QTextLayout_lineForTextPosition(void* self, int pos);
int QTextLayout_isValidCursorPosition(void* self, int pos);
int QTextLayout_nextCursorPosition(void* self, int oldPos);
int QTextLayout_previousCursorPosition(void* self, int oldPos);
int QTextLayout_leftCursorPosition(void* self, int oldPos);
int QTextLayout_rightCursorPosition(void* self, int oldPos);
void QTextLayout_draw(void* self, void* p, void* pos);
void QTextLayout_drawCursor(void* self, void* p, void* pos, int cursorPosition);
void QTextLayout_drawCursor2(void* self, void* p, void* pos, int cursorPosition, int width);
void* QTextLayout_position(void* self);
void QTextLayout_setPosition(void* self, void* p);
void* QTextLayout_boundingRect(void* self);
double QTextLayout_minimumWidth(void* self);
double QTextLayout_maximumWidth(void* self);
void** QTextLayout_glyphRuns(void* self, int from, int length, int flags, long long* __out_len);
void** QTextLayout_glyphRuns2(void* self, long long* __out_len);
void QTextLayout_setFlags(void* self, int flags);
int QTextLayout_nextCursorPosition2(void* self, int oldPos, int mode);
int QTextLayout_previousCursorPosition2(void* self, int oldPos, int mode);
void QTextLayout_draw2(void* self, void* p, void* pos, void** selections, long long selections_len);
void QTextLayout_draw3(void* self, void* p, void* pos, void** selections, long long selections_len, void* clip);
void** QTextLayout_glyphRuns3(void* self, int from, long long* __out_len);
void** QTextLayout_glyphRuns4(void* self, int from, int length, long long* __out_len);
void QTextLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
