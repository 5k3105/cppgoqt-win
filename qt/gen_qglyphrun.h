#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGlyphRun_new();
void* QGlyphRun_new2(void* other);
void* QGlyphRun_OperatorAssign(void* self, void* other);
void QGlyphRun_swap(void* self, void* other);
void* QGlyphRun_rawFont(void* self);
void QGlyphRun_setRawFont(void* self, void* rawFont);
unsigned int* QGlyphRun_glyphIndexes(void* self, long long* __out_len);
void QGlyphRun_setGlyphIndexes(void* self, unsigned int* glyphIndexes, long long glyphIndexes_len);
void** QGlyphRun_positions(void* self, long long* __out_len);
void QGlyphRun_setPositions(void* self, void** positions, long long positions_len);
void QGlyphRun_clear(void* self);
int QGlyphRun_OperatorEqual(void* self, void* other);
int QGlyphRun_OperatorNotEqual(void* self, void* other);
void QGlyphRun_setOverline(void* self, int overline);
int QGlyphRun_overline(void* self);
void QGlyphRun_setUnderline(void* self, int underline);
int QGlyphRun_underline(void* self);
void QGlyphRun_setStrikeOut(void* self, int strikeOut);
int QGlyphRun_strikeOut(void* self);
void QGlyphRun_setRightToLeft(void* self, int on);
int QGlyphRun_isRightToLeft(void* self);
void QGlyphRun_setFlag(void* self, int flag);
void QGlyphRun_setFlags(void* self, int flags);
int QGlyphRun_flags(void* self);
void QGlyphRun_setBoundingRect(void* self, void* boundingRect);
void* QGlyphRun_boundingRect(void* self);
long long* QGlyphRun_stringIndexes(void* self, long long* __out_len);
void QGlyphRun_setStringIndexes(void* self, long long* stringIndexes, long long stringIndexes_len);
void QGlyphRun_setSourceString(void* self, const char* sourceString, long long sourceString_len);
char* QGlyphRun_sourceString(void* self, long long* __out_len);
int QGlyphRun_isEmpty(void* self);
void QGlyphRun_setFlag2(void* self, int flag, int enabled);
void QGlyphRun_delete(void* self);

#ifdef __cplusplus
}
#endif
