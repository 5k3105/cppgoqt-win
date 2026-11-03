#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QPaintEngine__set_override(void* self, int idx);
int QPaintEngine_isActive(void* self);
void QPaintEngine_setActive(void* self, int newState);
int QPaintEngine_begin(void* self, void* pdev);
int QPaintEngine_end(void* self);
void QPaintEngine_updateState(void* self, void* state);
void QPaintEngine_drawRects(void* self, void* rects, int rectCount);
void QPaintEngine_drawRects2(void* self, void* rects, int rectCount);
void QPaintEngine_drawLines(void* self, void* lines, int lineCount);
void QPaintEngine_drawLines2(void* self, void* lines, int lineCount);
void QPaintEngine_drawEllipse(void* self, void* r);
void QPaintEngine_drawEllipse2(void* self, void* r);
void QPaintEngine_drawPath(void* self, void* path);
void QPaintEngine_drawPoints(void* self, void* points, int pointCount);
void QPaintEngine_drawPoints2(void* self, void* points, int pointCount);
void QPaintEngine_drawPolygon(void* self, void* points, int pointCount, int mode);
void QPaintEngine_drawPolygon2(void* self, void* points, int pointCount, int mode);
void QPaintEngine_drawPixmap(void* self, void* r, void* pm, void* sr);
void QPaintEngine_drawTextItem(void* self, void* p, void* textItem);
void QPaintEngine_drawTiledPixmap(void* self, void* r, void* pixmap, void* s);
void QPaintEngine_drawImage(void* self, void* r, void* pm, void* sr, int flags);
void QPaintEngine_setPaintDevice(void* self, void* device);
void* QPaintEngine_paintDevice(void* self);
void QPaintEngine_setSystemClip(void* self, void* baseClip);
void* QPaintEngine_systemClip(void* self);
void QPaintEngine_setSystemRect(void* self, void* rect);
void* QPaintEngine_systemRect(void* self);
void* QPaintEngine_coordinateOffset(void* self);
int QPaintEngine_type(void* self);
void QPaintEngine_fix_neg_rect(void* self, int* x, int* y, int* w, int* h);
int QPaintEngine_testDirty(void* self, int df);
void QPaintEngine_setDirty(void* self, int df);
void QPaintEngine_clearDirty(void* self, int df);
int QPaintEngine_hasFeature(void* self, int feature);
void* QPaintEngine_painter(void* self);
void QPaintEngine_syncState(void* self);
int QPaintEngine_isExtended(void* self);
void* QPaintEngine_createPixmap(void* self, void* size);
void* QPaintEngine_createPixmapFromImage(void* self, void* image, int flags);
void QPaintEngine_drawRects_default(void* self, void* rects, int rectCount);
void QPaintEngine_drawRects_default2(void* self, void* rects, int rectCount);
void QPaintEngine_drawLines_default(void* self, void* lines, int lineCount);
void QPaintEngine_drawLines_default2(void* self, void* lines, int lineCount);
void QPaintEngine_drawEllipse_default(void* self, void* r);
void QPaintEngine_drawEllipse_default2(void* self, void* r);
void QPaintEngine_drawPath_default(void* self, void* path);
void QPaintEngine_drawPoints_default(void* self, void* points, int pointCount);
void QPaintEngine_drawPoints_default2(void* self, void* points, int pointCount);
void QPaintEngine_drawPolygon_default(void* self, void* points, int pointCount, int mode);
void QPaintEngine_drawPolygon_default2(void* self, void* points, int pointCount, int mode);
void QPaintEngine_drawTextItem_default(void* self, void* p, void* textItem);
void QPaintEngine_drawTiledPixmap_default(void* self, void* r, void* pixmap, void* s);
void QPaintEngine_drawImage_default(void* self, void* r, void* pm, void* sr, int flags);
void* QPaintEngine_coordinateOffset_default(void* self);
void* QPaintEngine_createPixmap_default(void* self, void* size);
void* QPaintEngine_createPixmapFromImage_default(void* self, void* image, int flags);
void* QPaintEngine__Go_new(long long id);
void QPaintEngine_delete(void* self);

#ifdef __cplusplus
}
#endif
