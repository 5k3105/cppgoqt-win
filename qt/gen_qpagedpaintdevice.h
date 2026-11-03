#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPagedPaintDevice_base(void* self);
int QPagedPaintDevice_newPage(void* self);
int QPagedPaintDevice_setPageLayout(void* self, void* pageLayout);
int QPagedPaintDevice_setPageSize(void* self, void* pageSize);
int QPagedPaintDevice_setPageOrientation(void* self, int orientation);
int QPagedPaintDevice_setPageMargins(void* self, void* margins, int units);
void* QPagedPaintDevice_pageLayout(void* self);
void QPagedPaintDevice_setPageRanges(void* self, void* ranges);
void* QPagedPaintDevice_pageRanges(void* self);
double QPagedPaintDevice_devicePixelRatioFScale();
int QPagedPaintDevice_encodeMetricF(int metric, double value);
void QPagedPaintDevice_delete(void* self);

#ifdef __cplusplus
}
#endif
