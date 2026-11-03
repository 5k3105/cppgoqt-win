#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRawFont_new();
void* QRawFont_new2(const char* fileName, long long fileName_len, double pixelSize);
void* QRawFont_new3(void* fontData, double pixelSize);
void* QRawFont_new4(void* other);
void* QRawFont_new5(const char* fileName, long long fileName_len, double pixelSize, int hintingPreference);
void* QRawFont_new6(void* fontData, double pixelSize, int hintingPreference);
void* QRawFont_OperatorAssign(void* self, void* other);
void QRawFont_swap(void* self, void* other);
int QRawFont_isValid(void* self);
int QRawFont_OperatorEqual(void* self, void* other);
int QRawFont_OperatorNotEqual(void* self, void* other);
unsigned int QRawFont_glyphCount(void* self);
char* QRawFont_familyName(void* self, long long* __out_len);
char* QRawFont_styleName(void* self, long long* __out_len);
int QRawFont_style(void* self);
int QRawFont_weight(void* self);
unsigned int* QRawFont_glyphIndexesForString(void* self, const char* text, long long text_len, long long* __out_len);
void** QRawFont_advancesForGlyphIndexes(void* self, unsigned int* glyphIndexes, long long glyphIndexes_len, long long* __out_len);
void** QRawFont_advancesForGlyphIndexes2(void* self, unsigned int* glyphIndexes, long long glyphIndexes_len, int layoutFlags, long long* __out_len);
int QRawFont_glyphIndexesForChars(void* self, void* chars, int numChars, unsigned int* glyphIndexes, int* numGlyphs);
void* QRawFont_alphaMapForGlyph(void* self, unsigned int glyphIndex);
void* QRawFont_pathForGlyph(void* self, unsigned int glyphIndex);
void* QRawFont_boundingRect(void* self, unsigned int glyphIndex);
void QRawFont_setPixelSize(void* self, double pixelSize);
double QRawFont_pixelSize(void* self);
int QRawFont_hintingPreference(void* self);
double QRawFont_ascent(void* self);
double QRawFont_capHeight(void* self);
double QRawFont_descent(void* self);
double QRawFont_leading(void* self);
double QRawFont_xHeight(void* self);
double QRawFont_averageCharWidth(void* self);
double QRawFont_maxCharWidth(void* self);
double QRawFont_lineThickness(void* self);
double QRawFont_underlinePosition(void* self);
double QRawFont_unitsPerEm(void* self);
char* QRawFont_glyphName(void* self, unsigned int glyphIndex, long long* __out_len);
void QRawFont_loadFromFile(void* self, const char* fileName, long long fileName_len, double pixelSize, int hintingPreference);
void QRawFont_loadFromData(void* self, void* fontData, double pixelSize, int hintingPreference);
int QRawFont_supportsCharacter(void* self, unsigned int ucs4);
int QRawFont_supportsCharacter2(void* self, void* character);
int* QRawFont_supportedWritingSystems(void* self, long long* __out_len);
void* QRawFont_fontTable(void* self, const char* tagName);
void* QRawFont_fontTable2(void* self, void* tag);
void* QRawFont_fromFont(void* font);
void* QRawFont_alphaMapForGlyph2(void* self, unsigned int glyphIndex, int antialiasingType);
void* QRawFont_alphaMapForGlyph3(void* self, unsigned int glyphIndex, int antialiasingType, void* transform);
void* QRawFont_fromFont2(void* font, int writingSystem);
void QRawFont_delete(void* self);

#ifdef __cplusplus
}
#endif
