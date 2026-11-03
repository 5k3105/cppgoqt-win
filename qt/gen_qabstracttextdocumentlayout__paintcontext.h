#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractTextDocumentLayout__PaintContext_new();
void* QAbstractTextDocumentLayout__PaintContext_new2(void* arg1);
void* QAbstractTextDocumentLayout__PaintContext_OperatorAssign(void* self, void* arg1);
int QAbstractTextDocumentLayout__PaintContext_cursorPosition_get(void* self);
void QAbstractTextDocumentLayout__PaintContext_cursorPosition_set(void* self, int cursorPosition);
void* QAbstractTextDocumentLayout__PaintContext_palette_get(void* self);
void QAbstractTextDocumentLayout__PaintContext_palette_set(void* self, void* palette);
void* QAbstractTextDocumentLayout__PaintContext_clip_get(void* self);
void QAbstractTextDocumentLayout__PaintContext_clip_set(void* self, void* clip);
void** QAbstractTextDocumentLayout__PaintContext_selections_get(void* self, long long* __out_len);
void QAbstractTextDocumentLayout__PaintContext_selections_set(void* self, void** selections, long long selections_len);
void QAbstractTextDocumentLayout__PaintContext_delete(void* self);

#ifdef __cplusplus
}
#endif
