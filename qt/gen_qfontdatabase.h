#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFontDatabase_new();
int* QFontDatabase_standardSizes(long long* __out_len);
int* QFontDatabase_writingSystems(long long* __out_len);
int* QFontDatabase_writingSystems2(const char* family, long long family_len, long long* __out_len);
char** QFontDatabase_families(long long* __out_len);
char** QFontDatabase_styles(const char* family, long long family_len, long long* __out_len);
int* QFontDatabase_pointSizes(const char* family, long long family_len, long long* __out_len);
int* QFontDatabase_smoothSizes(const char* family, long long family_len, const char* style, long long style_len, long long* __out_len);
char* QFontDatabase_styleString(void* font, long long* __out_len);
char* QFontDatabase_styleString2(void* fontInfo, long long* __out_len);
void* QFontDatabase_font(const char* family, long long family_len, const char* style, long long style_len, int pointSize);
int QFontDatabase_isBitmapScalable(const char* family, long long family_len);
int QFontDatabase_isSmoothlyScalable(const char* family, long long family_len);
int QFontDatabase_isScalable(const char* family, long long family_len);
int QFontDatabase_isFixedPitch(const char* family, long long family_len);
int QFontDatabase_italic(const char* family, long long family_len, const char* style, long long style_len);
int QFontDatabase_bold(const char* family, long long family_len, const char* style, long long style_len);
int QFontDatabase_weight(const char* family, long long family_len, const char* style, long long style_len);
int QFontDatabase_hasFamily(const char* family, long long family_len);
int QFontDatabase_isPrivateFamily(const char* family, long long family_len);
char* QFontDatabase_writingSystemName(int writingSystem, long long* __out_len);
char* QFontDatabase_writingSystemSample(int writingSystem, long long* __out_len);
int QFontDatabase_addApplicationFont(const char* fileName, long long fileName_len);
int QFontDatabase_addApplicationFontFromData(void* fontData);
char** QFontDatabase_applicationFontFamilies(int id, long long* __out_len);
int QFontDatabase_removeApplicationFont(int id);
int QFontDatabase_removeAllApplicationFonts();
void QFontDatabase_addApplicationFallbackFontFamily(int script, const char* familyName, long long familyName_len);
int QFontDatabase_removeApplicationFallbackFontFamily(int script, const char* familyName, long long familyName_len);
void QFontDatabase_setApplicationFallbackFontFamilies(int arg1, char** familyNames, long long familyNames_len);
char** QFontDatabase_applicationFallbackFontFamilies(int script, long long* __out_len);
void QFontDatabase_addApplicationEmojiFontFamily(const char* familyName, long long familyName_len);
int QFontDatabase_removeApplicationEmojiFontFamily(const char* familyName, long long familyName_len);
void QFontDatabase_setApplicationEmojiFontFamilies(char** familyNames, long long familyNames_len);
char** QFontDatabase_applicationEmojiFontFamilies(long long* __out_len);
void* QFontDatabase_systemFont(int type_);
char** QFontDatabase_families2(int writingSystem, long long* __out_len);
int* QFontDatabase_pointSizes2(const char* family, long long family_len, const char* style, long long style_len, long long* __out_len);
int QFontDatabase_isBitmapScalable2(const char* family, long long family_len, const char* style, long long style_len);
int QFontDatabase_isSmoothlyScalable2(const char* family, long long family_len, const char* style, long long style_len);
int QFontDatabase_isScalable2(const char* family, long long family_len, const char* style, long long style_len);
int QFontDatabase_isFixedPitch2(const char* family, long long family_len, const char* style, long long style_len);
void QFontDatabase_delete(void* self);

#ifdef __cplusplus
}
#endif
