#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QToolTip_showText(void* pos, const char* text, long long text_len);
void QToolTip_hideText();
int QToolTip_isVisible();
char* QToolTip_text(long long* __out_len);
void* QToolTip_palette();
void QToolTip_setPalette(void* arg1);
void* QToolTip_font();
void QToolTip_setFont(void* arg1);
void QToolTip_showText2(void* pos, const char* text, long long text_len, void* w);
void QToolTip_showText3(void* pos, const char* text, long long text_len, void* w, void* rect);
void QToolTip_showText4(void* pos, const char* text, long long text_len, void* w, void* rect, int msecShowTime);
void QToolTip_delete(void* self);

#ifdef __cplusplus
}
#endif
