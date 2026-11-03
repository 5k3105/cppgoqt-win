#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextCharFormat_base(void* self);
void* QTextCharFormat_new();
void* QTextCharFormat_new2(void* arg1);
int QTextCharFormat_isValid(void* self);
void QTextCharFormat_setFont(void* self, void* font);
void* QTextCharFormat_font(void* self);
void QTextCharFormat_setFontFamily(void* self, const char* family, long long family_len);
char* QTextCharFormat_fontFamily(void* self, long long* __out_len);
void QTextCharFormat_setFontFamilies(void* self, char** families, long long families_len);
void* QTextCharFormat_fontFamilies(void* self);
void QTextCharFormat_setFontStyleName(void* self, const char* styleName, long long styleName_len);
void* QTextCharFormat_fontStyleName(void* self);
void QTextCharFormat_setFontPointSize(void* self, double size);
double QTextCharFormat_fontPointSize(void* self);
void QTextCharFormat_setFontWeight(void* self, int weight);
int QTextCharFormat_fontWeight(void* self);
void QTextCharFormat_setFontItalic(void* self, int italic);
int QTextCharFormat_fontItalic(void* self);
void QTextCharFormat_setFontCapitalization(void* self, int capitalization);
int QTextCharFormat_fontCapitalization(void* self);
void QTextCharFormat_setFontLetterSpacingType(void* self, int letterSpacingType);
int QTextCharFormat_fontLetterSpacingType(void* self);
void QTextCharFormat_setFontLetterSpacing(void* self, double spacing);
double QTextCharFormat_fontLetterSpacing(void* self);
void QTextCharFormat_setFontWordSpacing(void* self, double spacing);
double QTextCharFormat_fontWordSpacing(void* self);
void QTextCharFormat_setFontUnderline(void* self, int underline);
int QTextCharFormat_fontUnderline(void* self);
void QTextCharFormat_setFontOverline(void* self, int overline);
int QTextCharFormat_fontOverline(void* self);
void QTextCharFormat_setFontStrikeOut(void* self, int strikeOut);
int QTextCharFormat_fontStrikeOut(void* self);
void QTextCharFormat_setUnderlineColor(void* self, void* color);
void* QTextCharFormat_underlineColor(void* self);
void QTextCharFormat_setFontFixedPitch(void* self, int fixedPitch);
int QTextCharFormat_fontFixedPitch(void* self);
void QTextCharFormat_setFontStretch(void* self, int factor);
int QTextCharFormat_fontStretch(void* self);
void QTextCharFormat_setFontStyleHint(void* self, int hint);
void QTextCharFormat_setFontStyleStrategy(void* self, int strategy);
int QTextCharFormat_fontStyleHint(void* self);
int QTextCharFormat_fontStyleStrategy(void* self);
void QTextCharFormat_setFontHintingPreference(void* self, int hintingPreference);
int QTextCharFormat_fontHintingPreference(void* self);
void QTextCharFormat_setFontKerning(void* self, int enable);
int QTextCharFormat_fontKerning(void* self);
void QTextCharFormat_setUnderlineStyle(void* self, int style);
int QTextCharFormat_underlineStyle(void* self);
void QTextCharFormat_setVerticalAlignment(void* self, int alignment);
int QTextCharFormat_verticalAlignment(void* self);
void QTextCharFormat_setTextOutline(void* self, void* pen);
void* QTextCharFormat_textOutline(void* self);
void QTextCharFormat_setToolTip(void* self, const char* tip, long long tip_len);
char* QTextCharFormat_toolTip(void* self, long long* __out_len);
void QTextCharFormat_setSuperScriptBaseline(void* self, double baseline);
double QTextCharFormat_superScriptBaseline(void* self);
void QTextCharFormat_setSubScriptBaseline(void* self, double baseline);
double QTextCharFormat_subScriptBaseline(void* self);
void QTextCharFormat_setBaselineOffset(void* self, double baseline);
double QTextCharFormat_baselineOffset(void* self);
void QTextCharFormat_setAnchor(void* self, int anchor);
int QTextCharFormat_isAnchor(void* self);
void QTextCharFormat_setAnchorHref(void* self, const char* value, long long value_len);
char* QTextCharFormat_anchorHref(void* self, long long* __out_len);
void QTextCharFormat_setAnchorNames(void* self, char** names, long long names_len);
char** QTextCharFormat_anchorNames(void* self, long long* __out_len);
void QTextCharFormat_setTableCellRowSpan(void* self, int tableCellRowSpan);
int QTextCharFormat_tableCellRowSpan(void* self);
void QTextCharFormat_setTableCellColumnSpan(void* self, int tableCellColumnSpan);
int QTextCharFormat_tableCellColumnSpan(void* self);
void* QTextCharFormat_OperatorAssign(void* self, void* arg1);
void QTextCharFormat_setFont2(void* self, void* font, int behavior);
void QTextCharFormat_setFontStyleHint2(void* self, int hint, int strategy);
void QTextCharFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
