#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFont_new();
void* QFont_new2(const char* family, long long family_len);
void* QFont_new3(char** families, long long families_len);
void* QFont_new4(void* font, void* pd);
void* QFont_new5(void* font);
void* QFont_new6(const char* family, long long family_len, int pointSize);
void* QFont_new7(const char* family, long long family_len, int pointSize, int weight);
void* QFont_new8(const char* family, long long family_len, int pointSize, int weight, int italic);
void* QFont_new9(char** families, long long families_len, int pointSize);
void* QFont_new10(char** families, long long families_len, int pointSize, int weight);
void* QFont_new11(char** families, long long families_len, int pointSize, int weight, int italic);
void QFont_swap(void* self, void* other);
char* QFont_family(void* self, long long* __out_len);
void QFont_setFamily(void* self, const char* arg1, long long arg1_len);
char** QFont_families(void* self, long long* __out_len);
void QFont_setFamilies(void* self, char** arg1, long long arg1_len);
char* QFont_styleName(void* self, long long* __out_len);
void QFont_setStyleName(void* self, const char* arg1, long long arg1_len);
int QFont_pointSize(void* self);
void QFont_setPointSize(void* self, int arg1);
double QFont_pointSizeF(void* self);
void QFont_setPointSizeF(void* self, double arg1);
int QFont_pixelSize(void* self);
void QFont_setPixelSize(void* self, int arg1);
int QFont_weight(void* self);
void QFont_setWeight(void* self, int weight);
int QFont_bold(void* self);
void QFont_setBold(void* self, int arg1);
void QFont_setStyle(void* self, int style);
int QFont_style(void* self);
int QFont_italic(void* self);
void QFont_setItalic(void* self, int b);
int QFont_underline(void* self);
void QFont_setUnderline(void* self, int arg1);
int QFont_overline(void* self);
void QFont_setOverline(void* self, int arg1);
int QFont_strikeOut(void* self);
void QFont_setStrikeOut(void* self, int arg1);
int QFont_fixedPitch(void* self);
void QFont_setFixedPitch(void* self, int arg1);
int QFont_kerning(void* self);
void QFont_setKerning(void* self, int arg1);
int QFont_styleHint(void* self);
int QFont_styleStrategy(void* self);
void QFont_setStyleHint(void* self, int arg1);
void QFont_setStyleStrategy(void* self, int s);
int QFont_stretch(void* self);
void QFont_setStretch(void* self, int arg1);
double QFont_letterSpacing(void* self);
int QFont_letterSpacingType(void* self);
void QFont_setLetterSpacing(void* self, int type_, double spacing);
double QFont_wordSpacing(void* self);
void QFont_setWordSpacing(void* self, double spacing);
void QFont_setCapitalization(void* self, int arg1);
int QFont_capitalization(void* self);
void QFont_setHintingPreference(void* self, int hintingPreference);
int QFont_hintingPreference(void* self);
void QFont_setFeature(void* self, void* tag, unsigned int value);
void QFont_unsetFeature(void* self, void* tag);
unsigned int QFont_featureValue(void* self, void* tag);
int QFont_isFeatureSet(void* self, void* tag);
void** QFont_featureTags(void* self, long long* __out_len);
void QFont_clearFeatures(void* self);
void QFont_setVariableAxis(void* self, void* tag, float value);
void QFont_unsetVariableAxis(void* self, void* tag);
int QFont_isVariableAxisSet(void* self, void* tag);
float QFont_variableAxisValue(void* self, void* tag);
void QFont_clearVariableAxes(void* self);
void** QFont_variableAxisTags(void* self, long long* __out_len);
int QFont_exactMatch(void* self);
void* QFont_OperatorAssign(void* self, void* arg1);
int QFont_OperatorEqual(void* self, void* arg1);
int QFont_OperatorNotEqual(void* self, void* arg1);
int QFont_OperatorLesser(void* self, void* arg1);
void* QFont_ToQVariant(void* self);
int QFont_isCopyOf(void* self, void* arg1);
char* QFont_key(void* self, long long* __out_len);
char* QFont_toString(void* self, long long* __out_len);
int QFont_fromString(void* self, const char* arg1, long long arg1_len);
char* QFont_substitute(const char* arg1, long long arg1_len, long long* __out_len);
char** QFont_substitutes(const char* arg1, long long arg1_len, long long* __out_len);
char** QFont_substitutions(long long* __out_len);
void QFont_insertSubstitution(const char* arg1, long long arg1_len, const char* arg2, long long arg2_len);
void QFont_insertSubstitutions(const char* arg1, long long arg1_len, char** arg2, long long arg2_len);
void QFont_removeSubstitutions(const char* arg1, long long arg1_len);
void QFont_initialize();
void QFont_cleanup();
void QFont_cacheStatistics();
char* QFont_defaultFamily(void* self, long long* __out_len);
void* QFont_resolve(void* self, void* arg1);
unsigned int QFont_resolveMask(void* self);
void QFont_setResolveMask(void* self, unsigned int mask);
void QFont_setLegacyWeight(void* self, int legacyWeight);
int QFont_legacyWeight(void* self);
void QFont_setStyleHint2(void* self, int arg1, int arg2);
void* QFont_toQVariant(void* self);
void QFont_delete(void* self);

#ifdef __cplusplus
}
#endif
