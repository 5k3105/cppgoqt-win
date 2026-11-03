package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPaintEngine_isActive uintptr
	addr_QPaintEngine_setActive uintptr
	addr_QPaintEngine_begin uintptr
	addr_QPaintEngine_end uintptr
	addr_QPaintEngine_updateState uintptr
	addr_QPaintEngine_drawRects uintptr
	addr_QPaintEngine_drawRects2 uintptr
	addr_QPaintEngine_drawLines uintptr
	addr_QPaintEngine_drawLines2 uintptr
	addr_QPaintEngine_drawEllipse uintptr
	addr_QPaintEngine_drawEllipse2 uintptr
	addr_QPaintEngine_drawPath uintptr
	addr_QPaintEngine_drawPoints uintptr
	addr_QPaintEngine_drawPoints2 uintptr
	addr_QPaintEngine_drawPolygon uintptr
	addr_QPaintEngine_drawPolygon2 uintptr
	addr_QPaintEngine_drawPixmap uintptr
	addr_QPaintEngine_drawTextItem uintptr
	addr_QPaintEngine_drawTiledPixmap uintptr
	addr_QPaintEngine_drawImage uintptr
	addr_QPaintEngine_setPaintDevice uintptr
	addr_QPaintEngine_paintDevice uintptr
	addr_QPaintEngine_setSystemClip uintptr
	addr_QPaintEngine_systemClip uintptr
	addr_QPaintEngine_setSystemRect uintptr
	addr_QPaintEngine_systemRect uintptr
	addr_QPaintEngine_coordinateOffset uintptr
	addr_QPaintEngine_type uintptr
	addr_QPaintEngine_fix_neg_rect uintptr
	addr_QPaintEngine_testDirty uintptr
	addr_QPaintEngine_setDirty uintptr
	addr_QPaintEngine_clearDirty uintptr
	addr_QPaintEngine_hasFeature uintptr
	addr_QPaintEngine_painter uintptr
	addr_QPaintEngine_syncState uintptr
	addr_QPaintEngine_isExtended uintptr
	addr_QPaintEngine_createPixmap uintptr
	addr_QPaintEngine_createPixmapFromImage uintptr
	addr_QPaintEngine_drawRects_default uintptr
	addr_QPaintEngine_drawRects_default2 uintptr
	addr_QPaintEngine_drawLines_default uintptr
	addr_QPaintEngine_drawLines_default2 uintptr
	addr_QPaintEngine_drawEllipse_default uintptr
	addr_QPaintEngine_drawEllipse_default2 uintptr
	addr_QPaintEngine_drawPath_default uintptr
	addr_QPaintEngine_drawPoints_default uintptr
	addr_QPaintEngine_drawPoints_default2 uintptr
	addr_QPaintEngine_drawPolygon_default uintptr
	addr_QPaintEngine_drawPolygon_default2 uintptr
	addr_QPaintEngine_drawTextItem_default uintptr
	addr_QPaintEngine_drawTiledPixmap_default uintptr
	addr_QPaintEngine_drawImage_default uintptr
	addr_QPaintEngine_coordinateOffset_default uintptr
	addr_QPaintEngine_createPixmap_default uintptr
	addr_QPaintEngine_createPixmapFromImage_default uintptr
	pg_QPaintEngine__set_override func(uintptr, int32)
	pg_QPaintEngine__Go_new func(int64) uintptr
	pg_QPaintEngine_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPaintEngine_isActive, "QPaintEngine_isActive")
	symaddr(&addr_QPaintEngine_setActive, "QPaintEngine_setActive")
	symaddr(&addr_QPaintEngine_begin, "QPaintEngine_begin")
	symaddr(&addr_QPaintEngine_end, "QPaintEngine_end")
	symaddr(&addr_QPaintEngine_updateState, "QPaintEngine_updateState")
	symaddr(&addr_QPaintEngine_drawRects, "QPaintEngine_drawRects")
	symaddr(&addr_QPaintEngine_drawRects2, "QPaintEngine_drawRects2")
	symaddr(&addr_QPaintEngine_drawLines, "QPaintEngine_drawLines")
	symaddr(&addr_QPaintEngine_drawLines2, "QPaintEngine_drawLines2")
	symaddr(&addr_QPaintEngine_drawEllipse, "QPaintEngine_drawEllipse")
	symaddr(&addr_QPaintEngine_drawEllipse2, "QPaintEngine_drawEllipse2")
	symaddr(&addr_QPaintEngine_drawPath, "QPaintEngine_drawPath")
	symaddr(&addr_QPaintEngine_drawPoints, "QPaintEngine_drawPoints")
	symaddr(&addr_QPaintEngine_drawPoints2, "QPaintEngine_drawPoints2")
	symaddr(&addr_QPaintEngine_drawPolygon, "QPaintEngine_drawPolygon")
	symaddr(&addr_QPaintEngine_drawPolygon2, "QPaintEngine_drawPolygon2")
	symaddr(&addr_QPaintEngine_drawPixmap, "QPaintEngine_drawPixmap")
	symaddr(&addr_QPaintEngine_drawTextItem, "QPaintEngine_drawTextItem")
	symaddr(&addr_QPaintEngine_drawTiledPixmap, "QPaintEngine_drawTiledPixmap")
	symaddr(&addr_QPaintEngine_drawImage, "QPaintEngine_drawImage")
	symaddr(&addr_QPaintEngine_setPaintDevice, "QPaintEngine_setPaintDevice")
	symaddr(&addr_QPaintEngine_paintDevice, "QPaintEngine_paintDevice")
	symaddr(&addr_QPaintEngine_setSystemClip, "QPaintEngine_setSystemClip")
	symaddr(&addr_QPaintEngine_systemClip, "QPaintEngine_systemClip")
	symaddr(&addr_QPaintEngine_setSystemRect, "QPaintEngine_setSystemRect")
	symaddr(&addr_QPaintEngine_systemRect, "QPaintEngine_systemRect")
	symaddr(&addr_QPaintEngine_coordinateOffset, "QPaintEngine_coordinateOffset")
	symaddr(&addr_QPaintEngine_type, "QPaintEngine_type")
	symaddr(&addr_QPaintEngine_fix_neg_rect, "QPaintEngine_fix_neg_rect")
	symaddr(&addr_QPaintEngine_testDirty, "QPaintEngine_testDirty")
	symaddr(&addr_QPaintEngine_setDirty, "QPaintEngine_setDirty")
	symaddr(&addr_QPaintEngine_clearDirty, "QPaintEngine_clearDirty")
	symaddr(&addr_QPaintEngine_hasFeature, "QPaintEngine_hasFeature")
	symaddr(&addr_QPaintEngine_painter, "QPaintEngine_painter")
	symaddr(&addr_QPaintEngine_syncState, "QPaintEngine_syncState")
	symaddr(&addr_QPaintEngine_isExtended, "QPaintEngine_isExtended")
	symaddr(&addr_QPaintEngine_createPixmap, "QPaintEngine_createPixmap")
	symaddr(&addr_QPaintEngine_createPixmapFromImage, "QPaintEngine_createPixmapFromImage")
	symaddr(&addr_QPaintEngine_drawRects_default, "QPaintEngine_drawRects_default")
	symaddr(&addr_QPaintEngine_drawRects_default2, "QPaintEngine_drawRects_default2")
	symaddr(&addr_QPaintEngine_drawLines_default, "QPaintEngine_drawLines_default")
	symaddr(&addr_QPaintEngine_drawLines_default2, "QPaintEngine_drawLines_default2")
	symaddr(&addr_QPaintEngine_drawEllipse_default, "QPaintEngine_drawEllipse_default")
	symaddr(&addr_QPaintEngine_drawEllipse_default2, "QPaintEngine_drawEllipse_default2")
	symaddr(&addr_QPaintEngine_drawPath_default, "QPaintEngine_drawPath_default")
	symaddr(&addr_QPaintEngine_drawPoints_default, "QPaintEngine_drawPoints_default")
	symaddr(&addr_QPaintEngine_drawPoints_default2, "QPaintEngine_drawPoints_default2")
	symaddr(&addr_QPaintEngine_drawPolygon_default, "QPaintEngine_drawPolygon_default")
	symaddr(&addr_QPaintEngine_drawPolygon_default2, "QPaintEngine_drawPolygon_default2")
	symaddr(&addr_QPaintEngine_drawTextItem_default, "QPaintEngine_drawTextItem_default")
	symaddr(&addr_QPaintEngine_drawTiledPixmap_default, "QPaintEngine_drawTiledPixmap_default")
	symaddr(&addr_QPaintEngine_drawImage_default, "QPaintEngine_drawImage_default")
	symaddr(&addr_QPaintEngine_coordinateOffset_default, "QPaintEngine_coordinateOffset_default")
	symaddr(&addr_QPaintEngine_createPixmap_default, "QPaintEngine_createPixmap_default")
	symaddr(&addr_QPaintEngine_createPixmapFromImage_default, "QPaintEngine_createPixmapFromImage_default")
	reg(&pg_QPaintEngine__set_override, "QPaintEngine__set_override")
	reg(&pg_QPaintEngine__Go_new, "QPaintEngine__Go_new")
	reg(&pg_QPaintEngine_delete, "QPaintEngine_delete")
}

type QPaintEngine struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPaintEngine) UnsafeGoID() int64     { return this.vid }
func (this *QPaintEngine) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPaintEngine(h uintptr) *QPaintEngine {
	if h == 0 {
		return nil
	}
	return &QPaintEngine{h: h}
}

func newownedQPaintEngine(h uintptr) *QPaintEngine {
	o := bindQPaintEngine(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPaintEngine).Delete)
	}
	return o
}

func (this *QPaintEngine) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPaintEngine) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPaintEngine(h unsafe.Pointer) *QPaintEngine { return bindQPaintEngine(uintptr(h)) }
func UnsafeNewQPaintEngineOwned(h unsafe.Pointer) *QPaintEngine { return newownedQPaintEngine(uintptr(h)) }

func (this *QPaintEngine) UnsafeRelease() { this.h = 0 }

func (this *QPaintEngine) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngine) SetActive(newState bool) {
	purego.SyscallN(addr_QPaintEngine_setActive, this.h, uintptr(b2i(newState)))
	runtime.KeepAlive(this)
}

func (this *QPaintEngine) Begin(pdev *QPaintDevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_begin, this.h, uintptr(pdev.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pdev)
	return __r != 0
}

func (this *QPaintEngine) End() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_end, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngine) UpdateState(state *QPaintEngineState) {
	purego.SyscallN(addr_QPaintEngine_updateState, this.h, uintptr(state.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
}

func (this *QPaintEngine) DrawRects(rects *QRect, rectCount int) {
	purego.SyscallN(addr_QPaintEngine_drawRects, this.h, uintptr(rects.UnsafePointer()), uintptr(rectCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
}

func (this *QPaintEngine) DrawRects2(rects *QRectF, rectCount int) {
	purego.SyscallN(addr_QPaintEngine_drawRects2, this.h, uintptr(rects.UnsafePointer()), uintptr(rectCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
}

func (this *QPaintEngine) DrawLines(lines *QLine, lineCount int) {
	purego.SyscallN(addr_QPaintEngine_drawLines, this.h, uintptr(lines.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
}

func (this *QPaintEngine) DrawLines2(lines *QLineF, lineCount int) {
	purego.SyscallN(addr_QPaintEngine_drawLines2, this.h, uintptr(lines.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
}

func (this *QPaintEngine) DrawEllipse(r *QRectF) {
	purego.SyscallN(addr_QPaintEngine_drawEllipse, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPaintEngine) DrawEllipseWithQRect(r *QRect) {
	purego.SyscallN(addr_QPaintEngine_drawEllipse2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPaintEngine) DrawPath(path *QPainterPath) {
	purego.SyscallN(addr_QPaintEngine_drawPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPaintEngine) DrawPoints(points *QPointF, pointCount int) {
	purego.SyscallN(addr_QPaintEngine_drawPoints, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPoints2(points *QPoint, pointCount int) {
	purego.SyscallN(addr_QPaintEngine_drawPoints2, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPolygon(points *QPointF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	purego.SyscallN(addr_QPaintEngine_drawPolygon, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPolygon2(points *QPoint, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	purego.SyscallN(addr_QPaintEngine_drawPolygon2, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPixmap(r *QRectF, pm *QPixmap, sr *QRectF) {
	purego.SyscallN(addr_QPaintEngine_drawPixmap, this.h, uintptr(r.UnsafePointer()), uintptr(pm.UnsafePointer()), uintptr(sr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pm)
	runtime.KeepAlive(sr)
}

func (this *QPaintEngine) DrawTextItem(p *QPointF, textItem *QTextItem) {
	purego.SyscallN(addr_QPaintEngine_drawTextItem, this.h, uintptr(p.UnsafePointer()), uintptr(textItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(textItem)
}

func (this *QPaintEngine) DrawTiledPixmap(r *QRectF, pixmap *QPixmap, s *QPointF) {
	purego.SyscallN(addr_QPaintEngine_drawTiledPixmap, this.h, uintptr(r.UnsafePointer()), uintptr(pixmap.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(s)
}

func (this *QPaintEngine) DrawImage(r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPaintEngine_drawImage, this.h, uintptr(r.UnsafePointer()), uintptr(pm.UnsafePointer()), uintptr(sr.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pm)
	runtime.KeepAlive(sr)
}

func (this *QPaintEngine) SetPaintDevice(device *QPaintDevice) {
	purego.SyscallN(addr_QPaintEngine_setPaintDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QPaintEngine) PaintDevice() *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_paintDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPaintDevice(__r)
}

func (this *QPaintEngine) SetSystemClip(baseClip *QRegion) {
	purego.SyscallN(addr_QPaintEngine_setSystemClip, this.h, uintptr(baseClip.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(baseClip)
}

func (this *QPaintEngine) SystemClip() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_systemClip, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QPaintEngine) SetSystemRect(rect *QRect) {
	purego.SyscallN(addr_QPaintEngine_setSystemRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPaintEngine) SystemRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_systemRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPaintEngine) CoordinateOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_coordinateOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPaintEngine) Type() QPaintEngine__Type {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPaintEngine__Type(__r)
}

func (this *QPaintEngine) Fix_neg_rect() (int, int, int, int) {
	var x_out int32
	var y_out int32
	var w_out int32
	var h_out int32
	purego.SyscallN(addr_QPaintEngine_fix_neg_rect, this.h, uintptr(unsafe.Pointer(&x_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&w_out)), uintptr(unsafe.Pointer(&h_out)))
	runtime.KeepAlive(this)
	return int(x_out), int(y_out), int(w_out), int(h_out)
}

func (this *QPaintEngine) TestDirty(df QPaintEngine__DirtyFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_testDirty, this.h, uintptr(df))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngine) SetDirty(df QPaintEngine__DirtyFlag) {
	purego.SyscallN(addr_QPaintEngine_setDirty, this.h, uintptr(df))
	runtime.KeepAlive(this)
}

func (this *QPaintEngine) ClearDirty(df QPaintEngine__DirtyFlag) {
	purego.SyscallN(addr_QPaintEngine_clearDirty, this.h, uintptr(df))
	runtime.KeepAlive(this)
}

func (this *QPaintEngine) HasFeature(feature QPaintEngine__PaintEngineFeature) bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_hasFeature, this.h, uintptr(feature))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngine) Painter() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_painter, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QPaintEngine) SyncState() {
	purego.SyscallN(addr_QPaintEngine_syncState, this.h)
	runtime.KeepAlive(this)
}

func (this *QPaintEngine) IsExtended() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngine_isExtended, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngine) CreatePixmap(size *QSize) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_createPixmap, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QPaintEngine) CreatePixmapFromImage(image *QImage, flags ImageConversionFlag) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_createPixmapFromImage, this.h, uintptr(image.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return newownedQPixmap(__r)
}

func (this *QPaintEngine) DrawRectsDefault(rects *QRect, rectCount int) {
	purego.SyscallN(addr_QPaintEngine_drawRects_default, this.h, uintptr(rects.UnsafePointer()), uintptr(rectCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
}

func (this *QPaintEngine) DrawRects2Default(rects *QRectF, rectCount int) {
	purego.SyscallN(addr_QPaintEngine_drawRects_default2, this.h, uintptr(rects.UnsafePointer()), uintptr(rectCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
}

func (this *QPaintEngine) DrawLinesDefault(lines *QLine, lineCount int) {
	purego.SyscallN(addr_QPaintEngine_drawLines_default, this.h, uintptr(lines.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
}

func (this *QPaintEngine) DrawLines2Default(lines *QLineF, lineCount int) {
	purego.SyscallN(addr_QPaintEngine_drawLines_default2, this.h, uintptr(lines.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
}

func (this *QPaintEngine) DrawEllipseDefault(r *QRectF) {
	purego.SyscallN(addr_QPaintEngine_drawEllipse_default, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPaintEngine) DrawEllipseWithQRectDefault(r *QRect) {
	purego.SyscallN(addr_QPaintEngine_drawEllipse_default2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPaintEngine) DrawPathDefault(path *QPainterPath) {
	purego.SyscallN(addr_QPaintEngine_drawPath_default, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPaintEngine) DrawPointsDefault(points *QPointF, pointCount int) {
	purego.SyscallN(addr_QPaintEngine_drawPoints_default, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPoints2Default(points *QPoint, pointCount int) {
	purego.SyscallN(addr_QPaintEngine_drawPoints_default2, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPolygonDefault(points *QPointF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	purego.SyscallN(addr_QPaintEngine_drawPolygon_default, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawPolygon2Default(points *QPoint, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	purego.SyscallN(addr_QPaintEngine_drawPolygon_default2, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPaintEngine) DrawTextItemDefault(p *QPointF, textItem *QTextItem) {
	purego.SyscallN(addr_QPaintEngine_drawTextItem_default, this.h, uintptr(p.UnsafePointer()), uintptr(textItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(textItem)
}

func (this *QPaintEngine) DrawTiledPixmapDefault(r *QRectF, pixmap *QPixmap, s *QPointF) {
	purego.SyscallN(addr_QPaintEngine_drawTiledPixmap_default, this.h, uintptr(r.UnsafePointer()), uintptr(pixmap.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(s)
}

func (this *QPaintEngine) DrawImageDefault(r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPaintEngine_drawImage_default, this.h, uintptr(r.UnsafePointer()), uintptr(pm.UnsafePointer()), uintptr(sr.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pm)
	runtime.KeepAlive(sr)
}

func (this *QPaintEngine) CoordinateOffsetDefault() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_coordinateOffset_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPaintEngine) CreatePixmapDefault(size *QSize) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_createPixmap_default, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QPaintEngine) CreatePixmapFromImageDefault(image *QImage, flags ImageConversionFlag) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPaintEngine_createPixmapFromImage_default, this.h, uintptr(image.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return newownedQPixmap(__r)
}

func QPaintEngine_virtIdx(slot int64) int32 {
	switch slot {
	case -7312564458357607458:
		return 0
	case -4400013430294884342:
		return 1
	case 7692482089134926569:
		return 2
	case -6262278833611863026:
		return 3
	case -6371403872082420691:
		return 4
	case 6141870017934704802:
		return 5
	case 8374868496907056304:
		return 6
	case 3290090450242779774:
		return 7
	case 3888932263865552676:
		return 8
	case -4168977152563157681:
		return 9
	case 8175511122464863079:
		return 10
	case 2689733668603817204:
		return 11
	case 6207679886378890104:
		return 12
	case 6211818076911756990:
		return 13
	case 8440328321123084003:
		return 14
	case 4524770546924598563:
		return 15
	case -6193357004012257525:
		return 16
	case 8178762518988445540:
		return 17
	case 5783836824819670344:
		return 18
	case -4718660678191360312:
		return 19
	case 5331809331223620608:
		return 20
	case -6968771525665389363:
		return 21
	}
	return -1
}

func QPaintEngine_enable(h uintptr, idx int32) { pg_QPaintEngine__set_override(h, idx) }

func NewQPaintEngine() *QPaintEngine {
	id := nextVirtID()
	o := newownedQPaintEngine(pg_QPaintEngine__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QPaintEngine_virtIdx, QPaintEngine_enable)
	return o
}

func (this *QPaintEngine) OnBegin(fn func(*QPaintDevice) bool) {
	registerVirt(this.UnsafeGoID(), -7312564458357607458, func(pdev uintptr) int32 { return b2i(fn(UnsafeNewQPaintDevice(unsafe.Pointer(pdev)))) }) // slot "begin"
}

func (this *QPaintEngine) OnEnd(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -4400013430294884342, func() int32 { return b2i(fn()) }) // slot "end"
}

func (this *QPaintEngine) OnUpdateState(fn func(*QPaintEngineState)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(state uintptr) { fn(UnsafeNewQPaintEngineState(unsafe.Pointer(state))) }) // slot "updateState"
}

func (this *QPaintEngine) OnDrawPixmap(fn func(*QRectF, *QPixmap, *QRectF)) {
	registerVirt(this.UnsafeGoID(), -6262278833611863026, func(r uintptr, pm uintptr, sr uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(r)), UnsafeNewQPixmap(unsafe.Pointer(pm)), UnsafeNewQRectF(unsafe.Pointer(sr))) }) // slot "drawPixmap"
}

func (this *QPaintEngine) OnType(fn func() QPaintEngine__Type) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QPaintEngine) OnDrawRects(fn func(*QRect, int)) {
	registerVirt(this.UnsafeGoID(), 6141870017934704802, func(rects uintptr, rectCount int32) { fn(UnsafeNewQRect(unsafe.Pointer(rects)), int(rectCount)) }) // slot "drawRects"
}

func (this *QPaintEngine) OnDrawRects2(fn func(*QRectF, int)) {
	registerVirt(this.UnsafeGoID(), 8374868496907056304, func(rects uintptr, rectCount int32) { fn(UnsafeNewQRectF(unsafe.Pointer(rects)), int(rectCount)) }) // slot "drawRects2"
}

func (this *QPaintEngine) OnDrawLines(fn func(*QLine, int)) {
	registerVirt(this.UnsafeGoID(), 3290090450242779774, func(lines uintptr, lineCount int32) { fn(UnsafeNewQLine(unsafe.Pointer(lines)), int(lineCount)) }) // slot "drawLines"
}

func (this *QPaintEngine) OnDrawLines2(fn func(*QLineF, int)) {
	registerVirt(this.UnsafeGoID(), 3888932263865552676, func(lines uintptr, lineCount int32) { fn(UnsafeNewQLineF(unsafe.Pointer(lines)), int(lineCount)) }) // slot "drawLines2"
}

func (this *QPaintEngine) OnDrawEllipse(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), -4168977152563157681, func(r uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(r))) }) // slot "drawEllipse"
}

func (this *QPaintEngine) OnDrawEllipse2(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 8175511122464863079, func(r uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(r))) }) // slot "drawEllipse2"
}

func (this *QPaintEngine) OnDrawPath(fn func(*QPainterPath)) {
	registerVirt(this.UnsafeGoID(), 2689733668603817204, func(path uintptr) { fn(UnsafeNewQPainterPath(unsafe.Pointer(path))) }) // slot "drawPath"
}

func (this *QPaintEngine) OnDrawPoints(fn func(*QPointF, int)) {
	registerVirt(this.UnsafeGoID(), 6207679886378890104, func(points uintptr, pointCount int32) { fn(UnsafeNewQPointF(unsafe.Pointer(points)), int(pointCount)) }) // slot "drawPoints"
}

func (this *QPaintEngine) OnDrawPoints2(fn func(*QPoint, int)) {
	registerVirt(this.UnsafeGoID(), 6211818076911756990, func(points uintptr, pointCount int32) { fn(UnsafeNewQPoint(unsafe.Pointer(points)), int(pointCount)) }) // slot "drawPoints2"
}

func (this *QPaintEngine) OnDrawPolygon(fn func(*QPointF, int, QPaintEngine__PolygonDrawMode)) {
	registerVirt(this.UnsafeGoID(), 8440328321123084003, func(points uintptr, pointCount int32, mode int32) { fn(UnsafeNewQPointF(unsafe.Pointer(points)), int(pointCount), QPaintEngine__PolygonDrawMode(mode)) }) // slot "drawPolygon"
}

func (this *QPaintEngine) OnDrawPolygon2(fn func(*QPoint, int, QPaintEngine__PolygonDrawMode)) {
	registerVirt(this.UnsafeGoID(), 4524770546924598563, func(points uintptr, pointCount int32, mode int32) { fn(UnsafeNewQPoint(unsafe.Pointer(points)), int(pointCount), QPaintEngine__PolygonDrawMode(mode)) }) // slot "drawPolygon2"
}

func (this *QPaintEngine) OnDrawTextItem(fn func(*QPointF, *QTextItem)) {
	registerVirt(this.UnsafeGoID(), -6193357004012257525, func(p uintptr, textItem uintptr) { fn(UnsafeNewQPointF(unsafe.Pointer(p)), UnsafeNewQTextItem(unsafe.Pointer(textItem))) }) // slot "drawTextItem"
}

func (this *QPaintEngine) OnDrawTiledPixmap(fn func(*QRectF, *QPixmap, *QPointF)) {
	registerVirt(this.UnsafeGoID(), 8178762518988445540, func(r uintptr, pixmap uintptr, s uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(r)), UnsafeNewQPixmap(unsafe.Pointer(pixmap)), UnsafeNewQPointF(unsafe.Pointer(s))) }) // slot "drawTiledPixmap"
}

func (this *QPaintEngine) OnDrawImage(fn func(*QRectF, *QImage, *QRectF, ImageConversionFlag)) {
	registerVirt(this.UnsafeGoID(), 5783836824819670344, func(r uintptr, pm uintptr, sr uintptr, flags int32) { fn(UnsafeNewQRectF(unsafe.Pointer(r)), UnsafeNewQImage(unsafe.Pointer(pm)), UnsafeNewQRectF(unsafe.Pointer(sr)), ImageConversionFlag(flags)) }) // slot "drawImage"
}

func (this *QPaintEngine) OnCoordinateOffset(fn func() *QPoint) {
	registerVirt(this.UnsafeGoID(), -4718660678191360312, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "coordinateOffset"
}

func (this *QPaintEngine) OnCreatePixmap(fn func(*QSize) *QPixmap) {
	registerVirt(this.UnsafeGoID(), 5331809331223620608, func(size uintptr) uintptr { return uintptr(fn(UnsafeNewQSize(unsafe.Pointer(size))).UnsafePointer()) }) // slot "createPixmap"
}

func (this *QPaintEngine) OnCreatePixmapFromImage(fn func(*QImage, ImageConversionFlag) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -6968771525665389363, func(image uintptr, flags int32) uintptr { return uintptr(fn(UnsafeNewQImage(unsafe.Pointer(image)), ImageConversionFlag(flags)).UnsafePointer()) }) // slot "createPixmapFromImage"
}

func (this *QPaintEngine) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPaintEngine_delete(this.h)
	}
	this.UnsafeRelease()
}
