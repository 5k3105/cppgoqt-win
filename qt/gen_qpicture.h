#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPicture_base(void* self);
void QPicture__set_override(void* self, int idx);
void* QPicture_new(long long __goid);
void* QPicture_new2(long long __goid, void* arg1);
void* QPicture_new3(long long __goid, int formatVersion);
int QPicture_isNull(void* self);
unsigned int QPicture_size(void* self);
const char* QPicture_data(void* self);
void QPicture_setData(void* self, const char* data, unsigned int size);
int QPicture_play(void* self, void* p);
int QPicture_load(void* self, void* dev);
int QPicture_load2(void* self, const char* fileName, long long fileName_len);
int QPicture_save(void* self, void* dev);
int QPicture_save2(void* self, const char* fileName, long long fileName_len);
void* QPicture_boundingRect(void* self);
void QPicture_setBoundingRect(void* self, void* r);
void* QPicture_OperatorAssign(void* self, void* p);
void QPicture_swap(void* self, void* other);
void QPicture_detach(void* self);
int QPicture_isDetached(void* self);
double QPicture_devicePixelRatioFScale();
int QPicture_encodeMetricF(int metric, double value);
double QPicture_getDecodedMetricF(void* self, int metricA, int metricB);
int QPicture_devType_default(void* self);
void* QPicture_paintEngine_default(void* self);
int QPicture_metric_default(void* self, int m);
void QPicture_initPainter_default(void* self, void* painter);
void* QPicture_redirected_default(void* self, void* offset);
void* QPicture_sharedPainter_default(void* self);
void* QPicture__Go_new(long long id);
void QPicture_delete(void* self);

#ifdef __cplusplus
}
#endif
