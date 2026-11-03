#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStylePainter_base(void* self);
void* QStylePainter_new(void* w);
void* QStylePainter_new2();
void* QStylePainter_new3(void* pd, void* w);
int QStylePainter_begin(void* self, void* w);
int QStylePainter_begin2(void* self, void* pd, void* w);
void QStylePainter_drawPrimitive(void* self, int pe, void* opt);
void QStylePainter_drawControl(void* self, int ce, void* opt);
void QStylePainter_drawComplexControl(void* self, int cc, void* opt);
void QStylePainter_drawItemText(void* self, void* r, int flags, void* pal, int enabled, const char* text, long long text_len);
void QStylePainter_drawItemPixmap(void* self, void* r, int flags, void* pixmap);
void* QStylePainter_style(void* self);
void QStylePainter_drawItemText2(void* self, void* r, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole);
void QStylePainter_delete(void* self);

#ifdef __cplusplus
}
#endif
