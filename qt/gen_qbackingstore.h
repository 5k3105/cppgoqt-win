#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBackingStore_new(void* window);
void* QBackingStore_window(void* self);
void* QBackingStore_paintDevice(void* self);
void QBackingStore_flush(void* self, void* region);
void QBackingStore_resize(void* self, void* size);
void* QBackingStore_size(void* self);
int QBackingStore_scroll(void* self, void* area, int dx, int dy);
void QBackingStore_beginPaint(void* self, void* arg1);
void QBackingStore_endPaint(void* self);
void QBackingStore_setStaticContents(void* self, void* region);
void* QBackingStore_staticContents(void* self);
int QBackingStore_hasStaticContents(void* self);
void QBackingStore_flush2(void* self, void* region, void* window);
void QBackingStore_flush3(void* self, void* region, void* window, void* offset);
void QBackingStore_delete(void* self);

#ifdef __cplusplus
}
#endif
