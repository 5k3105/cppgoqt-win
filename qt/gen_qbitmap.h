#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBitmap_base(void* self);
void QBitmap_chain(void* self, void** out);
void QBitmap__set_override(void* self, int idx);
void* QBitmap_new(long long __goid);
void* QBitmap_new2(long long __goid, void* arg1);
void* QBitmap_new3(long long __goid, int w, int h);
void* QBitmap_new4(long long __goid, void* arg1);
void* QBitmap_new5(long long __goid, const char* fileName, long long fileName_len);
void* QBitmap_new6(long long __goid, void* arg1);
void* QBitmap_new7(long long __goid, const char* fileName, long long fileName_len, const char* format);
void* QBitmap_OperatorAssign(void* self, void* arg1);
void QBitmap_swap(void* self, void* other);
void* QBitmap_ToQVariant(void* self);
void QBitmap_clear(void* self);
void* QBitmap_fromImage(void* image);
void* QBitmap_fromData(void* size, const unsigned char* bits);
void* QBitmap_fromPixmap(void* pixmap);
void* QBitmap_transformed(void* self, void* matrix);
void* QBitmap_OperatorAssign2(void* self, void* arg1);
int QBitmap_defaultDepth();
void* QBitmap_trueMatrix(void* m, int w, int h);
void* QBitmap_fromImageReader(void* imageReader);
double QBitmap_devicePixelRatioFScale();
int QBitmap_encodeMetricF(int metric, double value);
double QBitmap_getDecodedMetricF(void* self, int metricA, int metricB);
void* QBitmap_fromImage2(void* image, int flags);
void* QBitmap_fromData2(void* size, const unsigned char* bits, int monoFormat);
void* QBitmap_fromImageReader2(void* imageReader, int flags);
void* QBitmap_toQVariant(void* self);
int QBitmap_devType_default(void* self);
void* QBitmap_paintEngine_default(void* self);
int QBitmap_metric_default(void* self, int arg1);
void QBitmap_initPainter_default(void* self, void* painter);
void* QBitmap_redirected_default(void* self, void* offset);
void* QBitmap_sharedPainter_default(void* self);
void* QBitmap__Go_new(long long id);
void QBitmap_delete(void* self);

#ifdef __cplusplus
}
#endif
