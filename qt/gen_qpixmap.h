#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPixmap_base(void* self);
void QPixmap__set_override(void* self, int idx);
void* QPixmap_new(long long __goid);
void* QPixmap_new2(long long __goid, int w, int h);
void* QPixmap_new3(long long __goid, void* arg1);
void* QPixmap_new4(long long __goid, const char* fileName, long long fileName_len);
void* QPixmap_new5(long long __goid, void* arg1);
void* QPixmap_new6(long long __goid, const char* fileName, long long fileName_len, const char* format);
void* QPixmap_new7(long long __goid, const char* fileName, long long fileName_len, const char* format, int flags);
void* QPixmap_OperatorAssign(void* self, void* arg1);
void QPixmap_swap(void* self, void* other);
void* QPixmap_ToQVariant(void* self);
int QPixmap_isNull(void* self);
int QPixmap_width(void* self);
int QPixmap_height(void* self);
void* QPixmap_size(void* self);
void* QPixmap_rect(void* self);
int QPixmap_depth(void* self);
int QPixmap_defaultDepth();
void QPixmap_fill(void* self);
void* QPixmap_mask(void* self);
void QPixmap_setMask(void* self, void* arg1);
double QPixmap_devicePixelRatio(void* self);
void QPixmap_setDevicePixelRatio(void* self, double scaleFactor);
void* QPixmap_deviceIndependentSize(void* self);
int QPixmap_hasAlpha(void* self);
int QPixmap_hasAlphaChannel(void* self);
void* QPixmap_createHeuristicMask(void* self);
void* QPixmap_createMaskFromColor(void* self, void* maskColor);
void* QPixmap_scaled(void* self, int w, int h);
void* QPixmap_scaled2(void* self, void* s);
void* QPixmap_scaledToWidth(void* self, int w);
void* QPixmap_scaledToHeight(void* self, int h);
void* QPixmap_transformed(void* self, void* arg1);
void* QPixmap_trueMatrix(void* m, int w, int h);
void* QPixmap_toImage(void* self);
void* QPixmap_fromImage(void* image);
void* QPixmap_fromImageReader(void* imageReader);
int QPixmap_load(void* self, const char* fileName, long long fileName_len);
int QPixmap_loadFromData(void* self, const unsigned char* buf, unsigned int len_);
int QPixmap_loadFromData2(void* self, void* data);
int QPixmap_save(void* self, const char* fileName, long long fileName_len);
int QPixmap_save2(void* self, void* device);
int QPixmap_convertFromImage(void* self, void* img);
void* QPixmap_copy(void* self, int x, int y, int width, int height);
void* QPixmap_copy2(void* self);
void QPixmap_scroll(void* self, int dx, int dy, int x, int y, int width, int height);
void QPixmap_scroll2(void* self, int dx, int dy, void* rect);
long long QPixmap_cacheKey(void* self);
int QPixmap_isDetached(void* self);
void QPixmap_detach(void* self);
int QPixmap_isQBitmap(void* self);
int QPixmap_OperatorNot(void* self);
double QPixmap_devicePixelRatioFScale();
int QPixmap_encodeMetricF(int metric, double value);
double QPixmap_getDecodedMetricF(void* self, int metricA, int metricB);
void QPixmap_fill2(void* self, void* fillColor);
void* QPixmap_createHeuristicMask2(void* self, int clipTight);
void* QPixmap_createMaskFromColor2(void* self, void* maskColor, int mode);
void* QPixmap_scaled3(void* self, int w, int h, int aspectMode);
void* QPixmap_scaled4(void* self, int w, int h, int aspectMode, int mode);
void* QPixmap_scaled5(void* self, void* s, int aspectMode);
void* QPixmap_scaled6(void* self, void* s, int aspectMode, int mode);
void* QPixmap_scaledToWidth2(void* self, int w, int mode);
void* QPixmap_scaledToHeight2(void* self, int h, int mode);
void* QPixmap_transformed2(void* self, void* arg1, int mode);
void* QPixmap_fromImage2(void* image, int flags);
void* QPixmap_fromImageReader2(void* imageReader, int flags);
int QPixmap_load2(void* self, const char* fileName, long long fileName_len, const char* format);
int QPixmap_load3(void* self, const char* fileName, long long fileName_len, const char* format, int flags);
int QPixmap_loadFromData3(void* self, const unsigned char* buf, unsigned int len_, const char* format);
int QPixmap_loadFromData4(void* self, const unsigned char* buf, unsigned int len_, const char* format, int flags);
int QPixmap_loadFromData5(void* self, void* data, const char* format);
int QPixmap_loadFromData6(void* self, void* data, const char* format, int flags);
int QPixmap_save3(void* self, const char* fileName, long long fileName_len, const char* format);
int QPixmap_save4(void* self, const char* fileName, long long fileName_len, const char* format, int quality);
int QPixmap_save5(void* self, void* device, const char* format);
int QPixmap_save6(void* self, void* device, const char* format, int quality);
int QPixmap_convertFromImage2(void* self, void* img, int flags);
void* QPixmap_copy3(void* self, void* rect);
void QPixmap_scroll3(void* self, int dx, int dy, int x, int y, int width, int height, void* exposed);
void QPixmap_scroll4(void* self, int dx, int dy, void* rect, void* exposed);
void* QPixmap_toQVariant(void* self);
int QPixmap_devType_default(void* self);
void* QPixmap_paintEngine_default(void* self);
int QPixmap_metric_default(void* self, int arg1);
void QPixmap_initPainter_default(void* self, void* painter);
void* QPixmap_redirected_default(void* self, void* offset);
void* QPixmap_sharedPainter_default(void* self);
void* QPixmap__Go_new(long long id);
void QPixmap_delete(void* self);

#ifdef __cplusplus
}
#endif
