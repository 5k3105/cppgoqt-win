#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIcon_new();
void* QIcon_new2(void* pixmap);
void* QIcon_new3(void* other);
void* QIcon_new4(const char* fileName, long long fileName_len);
void* QIcon_new5(void* engine);
void* QIcon_OperatorAssign(void* self, void* other);
void QIcon_swap(void* self, void* other);
void* QIcon_ToQVariant(void* self);
void* QIcon_pixmap(void* self, void* size);
void* QIcon_pixmap2(void* self, int w, int h);
void* QIcon_pixmap3(void* self, int extent);
void* QIcon_pixmap4(void* self, void* size, double devicePixelRatio);
void* QIcon_pixmap5(void* self, void* window, void* size);
void* QIcon_actualSize(void* self, void* size);
void* QIcon_actualSize2(void* self, void* window, void* size);
char* QIcon_name(void* self, long long* __out_len);
void QIcon_paint(void* self, void* painter, void* rect);
void QIcon_paint2(void* self, void* painter, int x, int y, int w, int h);
int QIcon_isNull(void* self);
int QIcon_isDetached(void* self);
void QIcon_detach(void* self);
long long QIcon_cacheKey(void* self);
void QIcon_addPixmap(void* self, void* pixmap);
void QIcon_addFile(void* self, const char* fileName, long long fileName_len);
void** QIcon_availableSizes(void* self, long long* __out_len);
void QIcon_setIsMask(void* self, int isMask);
int QIcon_isMask(void* self);
void* QIcon_fromTheme(const char* name, long long name_len);
void* QIcon_fromTheme2(const char* name, long long name_len, void* fallback);
int QIcon_hasThemeIcon(const char* name, long long name_len);
void* QIcon_fromTheme3(int icon);
void* QIcon_fromTheme4(int icon, void* fallback);
int QIcon_hasThemeIcon2(int icon);
char** QIcon_themeSearchPaths(long long* __out_len);
void QIcon_setThemeSearchPaths(char** searchpath, long long searchpath_len);
char** QIcon_fallbackSearchPaths(long long* __out_len);
void QIcon_setFallbackSearchPaths(char** paths, long long paths_len);
char* QIcon_themeName(long long* __out_len);
void QIcon_setThemeName(const char* path, long long path_len);
char* QIcon_fallbackThemeName(long long* __out_len);
void QIcon_setFallbackThemeName(const char* name, long long name_len);
void* QIcon_pixmap6(void* self, void* size, int mode);
void* QIcon_pixmap7(void* self, void* size, int mode, int state);
void* QIcon_pixmap8(void* self, int w, int h, int mode);
void* QIcon_pixmap9(void* self, int w, int h, int mode, int state);
void* QIcon_pixmap10(void* self, int extent, int mode);
void* QIcon_pixmap11(void* self, int extent, int mode, int state);
void* QIcon_pixmap12(void* self, void* size, double devicePixelRatio, int mode);
void* QIcon_pixmap13(void* self, void* size, double devicePixelRatio, int mode, int state);
void* QIcon_pixmap14(void* self, void* window, void* size, int mode);
void* QIcon_pixmap15(void* self, void* window, void* size, int mode, int state);
void* QIcon_actualSize3(void* self, void* size, int mode);
void* QIcon_actualSize4(void* self, void* size, int mode, int state);
void* QIcon_actualSize5(void* self, void* window, void* size, int mode);
void* QIcon_actualSize6(void* self, void* window, void* size, int mode, int state);
void QIcon_paint3(void* self, void* painter, void* rect, int alignment);
void QIcon_paint4(void* self, void* painter, void* rect, int alignment, int mode);
void QIcon_paint5(void* self, void* painter, void* rect, int alignment, int mode, int state);
void QIcon_paint6(void* self, void* painter, int x, int y, int w, int h, int alignment);
void QIcon_paint7(void* self, void* painter, int x, int y, int w, int h, int alignment, int mode);
void QIcon_paint8(void* self, void* painter, int x, int y, int w, int h, int alignment, int mode, int state);
void QIcon_addPixmap2(void* self, void* pixmap, int mode);
void QIcon_addPixmap3(void* self, void* pixmap, int mode, int state);
void QIcon_addFile2(void* self, const char* fileName, long long fileName_len, void* size);
void QIcon_addFile3(void* self, const char* fileName, long long fileName_len, void* size, int mode);
void QIcon_addFile4(void* self, const char* fileName, long long fileName_len, void* size, int mode, int state);
void** QIcon_availableSizes2(void* self, int mode, long long* __out_len);
void** QIcon_availableSizes3(void* self, int mode, int state, long long* __out_len);
void* QIcon_toQVariant(void* self);
void QIcon_delete(void* self);

#ifdef __cplusplus
}
#endif
