package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPainterPath_new uintptr
	addr_QPainterPath_new2 uintptr
	addr_QPainterPath_new3 uintptr
	addr_QPainterPath_OperatorAssign uintptr
	addr_QPainterPath_swap uintptr
	addr_QPainterPath_clear uintptr
	addr_QPainterPath_reserve uintptr
	addr_QPainterPath_capacity uintptr
	addr_QPainterPath_closeSubpath uintptr
	addr_QPainterPath_moveTo uintptr
	pg_QPainterPath_moveTo2 func(uintptr, float64, float64)
	addr_QPainterPath_lineTo uintptr
	pg_QPainterPath_lineTo2 func(uintptr, float64, float64)
	pg_QPainterPath_arcMoveTo func(uintptr, uintptr, float64)
	pg_QPainterPath_arcMoveTo2 func(uintptr, float64, float64, float64, float64, float64)
	pg_QPainterPath_arcTo func(uintptr, uintptr, float64, float64)
	pg_QPainterPath_arcTo2 func(uintptr, float64, float64, float64, float64, float64, float64)
	addr_QPainterPath_cubicTo uintptr
	pg_QPainterPath_cubicTo2 func(uintptr, float64, float64, float64, float64, float64, float64)
	addr_QPainterPath_quadTo uintptr
	pg_QPainterPath_quadTo2 func(uintptr, float64, float64, float64, float64)
	addr_QPainterPath_currentPosition uintptr
	addr_QPainterPath_addRect uintptr
	pg_QPainterPath_addRect2 func(uintptr, float64, float64, float64, float64)
	addr_QPainterPath_addEllipse uintptr
	pg_QPainterPath_addEllipse2 func(uintptr, float64, float64, float64, float64)
	pg_QPainterPath_addEllipse3 func(uintptr, uintptr, float64, float64)
	addr_QPainterPath_addPolygon uintptr
	addr_QPainterPath_addText uintptr
	pg_QPainterPath_addText2 func(uintptr, float64, float64, uintptr, string, int64)
	addr_QPainterPath_addPath uintptr
	addr_QPainterPath_addRegion uintptr
	pg_QPainterPath_addRoundedRect func(uintptr, uintptr, float64, float64)
	pg_QPainterPath_addRoundedRect2 func(uintptr, float64, float64, float64, float64, float64, float64)
	addr_QPainterPath_connectPath uintptr
	addr_QPainterPath_contains uintptr
	addr_QPainterPath_contains2 uintptr
	addr_QPainterPath_intersects uintptr
	pg_QPainterPath_translate func(uintptr, float64, float64)
	addr_QPainterPath_translate2 uintptr
	pg_QPainterPath_translated func(uintptr, float64, float64) uintptr
	addr_QPainterPath_translated2 uintptr
	addr_QPainterPath_boundingRect uintptr
	addr_QPainterPath_controlPointRect uintptr
	addr_QPainterPath_fillRule uintptr
	addr_QPainterPath_setFillRule uintptr
	addr_QPainterPath_isEmpty uintptr
	addr_QPainterPath_toReversed uintptr
	addr_QPainterPath_toSubpathPolygons uintptr
	addr_QPainterPath_toFillPolygons uintptr
	addr_QPainterPath_toFillPolygon uintptr
	addr_QPainterPath_elementCount uintptr
	addr_QPainterPath_elementAt uintptr
	pg_QPainterPath_setElementPositionAt func(uintptr, int32, float64, float64)
	addr_QPainterPath_isCachingEnabled uintptr
	addr_QPainterPath_setCachingEnabled uintptr
	pg_QPainterPath_length func(uintptr) float64
	pg_QPainterPath_percentAtLength func(uintptr, float64) float64
	pg_QPainterPath_pointAtPercent func(uintptr, float64) uintptr
	pg_QPainterPath_angleAtPercent func(uintptr, float64) float64
	pg_QPainterPath_slopeAtPercent func(uintptr, float64) float64
	pg_QPainterPath_trimmed func(uintptr, float64, float64) uintptr
	addr_QPainterPath_intersects2 uintptr
	addr_QPainterPath_contains3 uintptr
	addr_QPainterPath_united uintptr
	addr_QPainterPath_intersected uintptr
	addr_QPainterPath_subtracted uintptr
	addr_QPainterPath_simplified uintptr
	addr_QPainterPath_OperatorEqual uintptr
	addr_QPainterPath_OperatorNotEqual uintptr
	addr_QPainterPath_OperatorBitwiseAnd uintptr
	addr_QPainterPath_OperatorBitwiseOr uintptr
	addr_QPainterPath_OperatorPlus uintptr
	addr_QPainterPath_OperatorMinus uintptr
	addr_QPainterPath_OperatorBitwiseAndAssign uintptr
	addr_QPainterPath_OperatorBitwiseOrAssign uintptr
	addr_QPainterPath_OperatorPlusAssign uintptr
	addr_QPainterPath_OperatorMinusAssign uintptr
	pg_QPainterPath_addRoundedRect3 func(uintptr, uintptr, float64, float64, int32)
	pg_QPainterPath_addRoundedRect4 func(uintptr, float64, float64, float64, float64, float64, float64, int32)
	addr_QPainterPath_toSubpathPolygons2 uintptr
	addr_QPainterPath_toFillPolygons2 uintptr
	addr_QPainterPath_toFillPolygon2 uintptr
	pg_QPainterPath_trimmed2 func(uintptr, float64, float64, float64) uintptr
	addr_QPainterPath_OperatorMultiply uintptr
	pg_QPainterPath_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPainterPath_new, "QPainterPath_new")
	symaddr(&addr_QPainterPath_new2, "QPainterPath_new2")
	symaddr(&addr_QPainterPath_new3, "QPainterPath_new3")
	symaddr(&addr_QPainterPath_OperatorAssign, "QPainterPath_OperatorAssign")
	symaddr(&addr_QPainterPath_swap, "QPainterPath_swap")
	symaddr(&addr_QPainterPath_clear, "QPainterPath_clear")
	symaddr(&addr_QPainterPath_reserve, "QPainterPath_reserve")
	symaddr(&addr_QPainterPath_capacity, "QPainterPath_capacity")
	symaddr(&addr_QPainterPath_closeSubpath, "QPainterPath_closeSubpath")
	symaddr(&addr_QPainterPath_moveTo, "QPainterPath_moveTo")
	reg(&pg_QPainterPath_moveTo2, "QPainterPath_moveTo2")
	symaddr(&addr_QPainterPath_lineTo, "QPainterPath_lineTo")
	reg(&pg_QPainterPath_lineTo2, "QPainterPath_lineTo2")
	reg(&pg_QPainterPath_arcMoveTo, "QPainterPath_arcMoveTo")
	reg(&pg_QPainterPath_arcMoveTo2, "QPainterPath_arcMoveTo2")
	reg(&pg_QPainterPath_arcTo, "QPainterPath_arcTo")
	reg(&pg_QPainterPath_arcTo2, "QPainterPath_arcTo2")
	symaddr(&addr_QPainterPath_cubicTo, "QPainterPath_cubicTo")
	reg(&pg_QPainterPath_cubicTo2, "QPainterPath_cubicTo2")
	symaddr(&addr_QPainterPath_quadTo, "QPainterPath_quadTo")
	reg(&pg_QPainterPath_quadTo2, "QPainterPath_quadTo2")
	symaddr(&addr_QPainterPath_currentPosition, "QPainterPath_currentPosition")
	symaddr(&addr_QPainterPath_addRect, "QPainterPath_addRect")
	reg(&pg_QPainterPath_addRect2, "QPainterPath_addRect2")
	symaddr(&addr_QPainterPath_addEllipse, "QPainterPath_addEllipse")
	reg(&pg_QPainterPath_addEllipse2, "QPainterPath_addEllipse2")
	reg(&pg_QPainterPath_addEllipse3, "QPainterPath_addEllipse3")
	symaddr(&addr_QPainterPath_addPolygon, "QPainterPath_addPolygon")
	symaddr(&addr_QPainterPath_addText, "QPainterPath_addText")
	reg(&pg_QPainterPath_addText2, "QPainterPath_addText2")
	symaddr(&addr_QPainterPath_addPath, "QPainterPath_addPath")
	symaddr(&addr_QPainterPath_addRegion, "QPainterPath_addRegion")
	reg(&pg_QPainterPath_addRoundedRect, "QPainterPath_addRoundedRect")
	reg(&pg_QPainterPath_addRoundedRect2, "QPainterPath_addRoundedRect2")
	symaddr(&addr_QPainterPath_connectPath, "QPainterPath_connectPath")
	symaddr(&addr_QPainterPath_contains, "QPainterPath_contains")
	symaddr(&addr_QPainterPath_contains2, "QPainterPath_contains2")
	symaddr(&addr_QPainterPath_intersects, "QPainterPath_intersects")
	reg(&pg_QPainterPath_translate, "QPainterPath_translate")
	symaddr(&addr_QPainterPath_translate2, "QPainterPath_translate2")
	reg(&pg_QPainterPath_translated, "QPainterPath_translated")
	symaddr(&addr_QPainterPath_translated2, "QPainterPath_translated2")
	symaddr(&addr_QPainterPath_boundingRect, "QPainterPath_boundingRect")
	symaddr(&addr_QPainterPath_controlPointRect, "QPainterPath_controlPointRect")
	symaddr(&addr_QPainterPath_fillRule, "QPainterPath_fillRule")
	symaddr(&addr_QPainterPath_setFillRule, "QPainterPath_setFillRule")
	symaddr(&addr_QPainterPath_isEmpty, "QPainterPath_isEmpty")
	symaddr(&addr_QPainterPath_toReversed, "QPainterPath_toReversed")
	symaddr(&addr_QPainterPath_toSubpathPolygons, "QPainterPath_toSubpathPolygons")
	symaddr(&addr_QPainterPath_toFillPolygons, "QPainterPath_toFillPolygons")
	symaddr(&addr_QPainterPath_toFillPolygon, "QPainterPath_toFillPolygon")
	symaddr(&addr_QPainterPath_elementCount, "QPainterPath_elementCount")
	symaddr(&addr_QPainterPath_elementAt, "QPainterPath_elementAt")
	reg(&pg_QPainterPath_setElementPositionAt, "QPainterPath_setElementPositionAt")
	symaddr(&addr_QPainterPath_isCachingEnabled, "QPainterPath_isCachingEnabled")
	symaddr(&addr_QPainterPath_setCachingEnabled, "QPainterPath_setCachingEnabled")
	reg(&pg_QPainterPath_length, "QPainterPath_length")
	reg(&pg_QPainterPath_percentAtLength, "QPainterPath_percentAtLength")
	reg(&pg_QPainterPath_pointAtPercent, "QPainterPath_pointAtPercent")
	reg(&pg_QPainterPath_angleAtPercent, "QPainterPath_angleAtPercent")
	reg(&pg_QPainterPath_slopeAtPercent, "QPainterPath_slopeAtPercent")
	reg(&pg_QPainterPath_trimmed, "QPainterPath_trimmed")
	symaddr(&addr_QPainterPath_intersects2, "QPainterPath_intersects2")
	symaddr(&addr_QPainterPath_contains3, "QPainterPath_contains3")
	symaddr(&addr_QPainterPath_united, "QPainterPath_united")
	symaddr(&addr_QPainterPath_intersected, "QPainterPath_intersected")
	symaddr(&addr_QPainterPath_subtracted, "QPainterPath_subtracted")
	symaddr(&addr_QPainterPath_simplified, "QPainterPath_simplified")
	symaddr(&addr_QPainterPath_OperatorEqual, "QPainterPath_OperatorEqual")
	symaddr(&addr_QPainterPath_OperatorNotEqual, "QPainterPath_OperatorNotEqual")
	symaddr(&addr_QPainterPath_OperatorBitwiseAnd, "QPainterPath_OperatorBitwiseAnd")
	symaddr(&addr_QPainterPath_OperatorBitwiseOr, "QPainterPath_OperatorBitwiseOr")
	symaddr(&addr_QPainterPath_OperatorPlus, "QPainterPath_OperatorPlus")
	symaddr(&addr_QPainterPath_OperatorMinus, "QPainterPath_OperatorMinus")
	symaddr(&addr_QPainterPath_OperatorBitwiseAndAssign, "QPainterPath_OperatorBitwiseAndAssign")
	symaddr(&addr_QPainterPath_OperatorBitwiseOrAssign, "QPainterPath_OperatorBitwiseOrAssign")
	symaddr(&addr_QPainterPath_OperatorPlusAssign, "QPainterPath_OperatorPlusAssign")
	symaddr(&addr_QPainterPath_OperatorMinusAssign, "QPainterPath_OperatorMinusAssign")
	reg(&pg_QPainterPath_addRoundedRect3, "QPainterPath_addRoundedRect3")
	reg(&pg_QPainterPath_addRoundedRect4, "QPainterPath_addRoundedRect4")
	symaddr(&addr_QPainterPath_toSubpathPolygons2, "QPainterPath_toSubpathPolygons2")
	symaddr(&addr_QPainterPath_toFillPolygons2, "QPainterPath_toFillPolygons2")
	symaddr(&addr_QPainterPath_toFillPolygon2, "QPainterPath_toFillPolygon2")
	reg(&pg_QPainterPath_trimmed2, "QPainterPath_trimmed2")
	symaddr(&addr_QPainterPath_OperatorMultiply, "QPainterPath_OperatorMultiply")
	reg(&pg_QPainterPath_delete, "QPainterPath_delete")
}

type QPainterPath struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPainterPath) UnsafeGoID() int64     { return this.vid }
func (this *QPainterPath) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPainterPath(h uintptr) *QPainterPath {
	if h == 0 {
		return nil
	}
	return &QPainterPath{h: h}
}

func newownedQPainterPath(h uintptr) *QPainterPath {
	o := bindQPainterPath(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPainterPath).Delete)
	}
	return o
}

func (this *QPainterPath) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPainterPath) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPainterPath(h unsafe.Pointer) *QPainterPath { return bindQPainterPath(uintptr(h)) }
func UnsafeNewQPainterPathOwned(h unsafe.Pointer) *QPainterPath { return newownedQPainterPath(uintptr(h)) }

func (this *QPainterPath) UnsafeRelease() { this.h = 0 }

func NewQPainterPath() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_new)
	__o := newownedQPainterPath(__r)
	return __o
}

func NewQPainterPath2(startPoint *QPointF) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_new2, uintptr(startPoint.UnsafePointer()))
	__o := newownedQPainterPath(__r)
	runtime.KeepAlive(startPoint)
	return __o
}

func NewQPainterPath3(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_new3, uintptr(other.UnsafePointer()))
	__o := newownedQPainterPath(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QPainterPath) OperatorAssign(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPainterPath(__r)
}

func (this *QPainterPath) Swap(other *QPainterPath) {
	purego.SyscallN(addr_QPainterPath_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPainterPath) Clear() {
	purego.SyscallN(addr_QPainterPath_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) Reserve(size int) {
	purego.SyscallN(addr_QPainterPath_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPainterPath) Capacity() int {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_capacity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPainterPath) CloseSubpath() {
	purego.SyscallN(addr_QPainterPath_closeSubpath, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) MoveTo(p *QPointF) {
	purego.SyscallN(addr_QPainterPath_moveTo, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QPainterPath) MoveTo2(x float64, y float64) {
	pg_QPainterPath_moveTo2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) LineTo(p *QPointF) {
	purego.SyscallN(addr_QPainterPath_lineTo, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QPainterPath) LineTo2(x float64, y float64) {
	pg_QPainterPath_lineTo2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) ArcMoveTo(rect *QRectF, angle float64) {
	pg_QPainterPath_arcMoveTo(this.h, uintptr(rect.UnsafePointer()), angle)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainterPath) ArcMoveTo2(x float64, y float64, w float64, h float64, angle float64) {
	pg_QPainterPath_arcMoveTo2(this.h, x, y, w, h, angle)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) ArcTo(rect *QRectF, startAngle float64, arcLength float64) {
	pg_QPainterPath_arcTo(this.h, uintptr(rect.UnsafePointer()), startAngle, arcLength)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainterPath) ArcTo2(x float64, y float64, w float64, h float64, startAngle float64, arcLength float64) {
	pg_QPainterPath_arcTo2(this.h, x, y, w, h, startAngle, arcLength)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) CubicTo(ctrlPt1 *QPointF, ctrlPt2 *QPointF, endPt *QPointF) {
	purego.SyscallN(addr_QPainterPath_cubicTo, this.h, uintptr(ctrlPt1.UnsafePointer()), uintptr(ctrlPt2.UnsafePointer()), uintptr(endPt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ctrlPt1)
	runtime.KeepAlive(ctrlPt2)
	runtime.KeepAlive(endPt)
}

func (this *QPainterPath) CubicTo2(ctrlPt1x float64, ctrlPt1y float64, ctrlPt2x float64, ctrlPt2y float64, endPtx float64, endPty float64) {
	pg_QPainterPath_cubicTo2(this.h, ctrlPt1x, ctrlPt1y, ctrlPt2x, ctrlPt2y, endPtx, endPty)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) QuadTo(ctrlPt *QPointF, endPt *QPointF) {
	purego.SyscallN(addr_QPainterPath_quadTo, this.h, uintptr(ctrlPt.UnsafePointer()), uintptr(endPt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ctrlPt)
	runtime.KeepAlive(endPt)
}

func (this *QPainterPath) QuadTo2(ctrlPtx float64, ctrlPty float64, endPtx float64, endPty float64) {
	pg_QPainterPath_quadTo2(this.h, ctrlPtx, ctrlPty, endPtx, endPty)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) CurrentPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_currentPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPainterPath) AddRect(rect *QRectF) {
	purego.SyscallN(addr_QPainterPath_addRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainterPath) AddRect2(x float64, y float64, w float64, h float64) {
	pg_QPainterPath_addRect2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) AddEllipse(rect *QRectF) {
	purego.SyscallN(addr_QPainterPath_addEllipse, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainterPath) AddEllipse2(x float64, y float64, w float64, h float64) {
	pg_QPainterPath_addEllipse2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) AddEllipse3(center *QPointF, rx float64, ry float64) {
	pg_QPainterPath_addEllipse3(this.h, uintptr(center.UnsafePointer()), rx, ry)
	runtime.KeepAlive(this)
	runtime.KeepAlive(center)
}

func (this *QPainterPath) AddPolygon(polygon *QPolygonF) {
	purego.SyscallN(addr_QPainterPath_addPolygon, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainterPath) AddText(point *QPointF, f *QFont, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainterPath_addText, this.h, uintptr(point.UnsafePointer()), uintptr(f.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	runtime.KeepAlive(f)
	runtime.KeepAlive(text)
}

func (this *QPainterPath) AddText2(x float64, y float64, f *QFont, text string) {
	pg_QPainterPath_addText2(this.h, x, y, uintptr(f.UnsafePointer()), text, int64(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func (this *QPainterPath) AddPath(path *QPainterPath) {
	purego.SyscallN(addr_QPainterPath_addPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPainterPath) AddRegion(region *QRegion) {
	purego.SyscallN(addr_QPainterPath_addRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QPainterPath) AddRoundedRect(rect *QRectF, xRadius float64, yRadius float64) {
	pg_QPainterPath_addRoundedRect(this.h, uintptr(rect.UnsafePointer()), xRadius, yRadius)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainterPath) AddRoundedRect2(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64) {
	pg_QPainterPath_addRoundedRect2(this.h, x, y, w, h, xRadius, yRadius)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) ConnectPath(path *QPainterPath) {
	purego.SyscallN(addr_QPainterPath_connectPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPainterPath) Contains(pt *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_contains, this.h, uintptr(pt.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return __r != 0
}

func (this *QPainterPath) ContainsWithRect(rect *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_contains2, this.h, uintptr(rect.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return __r != 0
}

func (this *QPainterPath) Intersects(rect *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_intersects, this.h, uintptr(rect.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return __r != 0
}

func (this *QPainterPath) Translate(dx float64, dy float64) {
	pg_QPainterPath_translate(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) TranslateWithOffset(offset *QPointF) {
	purego.SyscallN(addr_QPainterPath_translate2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QPainterPath) Translated(dx float64, dy float64) *QPainterPath {
	__r := pg_QPainterPath_translated(this.h, dx, dy)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) TranslatedWithOffset(offset *QPointF) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_translated2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPainterPath) ControlPointRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_controlPointRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPainterPath) FillRule() FillRule {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_fillRule, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return FillRule(__r)
}

func (this *QPainterPath) SetFillRule(fillRule FillRule) {
	purego.SyscallN(addr_QPainterPath_setFillRule, this.h, uintptr(fillRule))
	runtime.KeepAlive(this)
}

func (this *QPainterPath) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainterPath) ToReversed() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toReversed, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) ToSubpathPolygons() []*QPolygonF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toSubpathPolygons, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPolygonF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPolygonF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPainterPath) ToFillPolygons() []*QPolygonF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toFillPolygons, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPolygonF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPolygonF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPainterPath) ToFillPolygon() *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toFillPolygon, this.h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QPainterPath) ElementCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_elementCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPainterPath) ElementAt(i int) *QPainterPath__Element {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_elementAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQPainterPath__Element(__r)
}

func (this *QPainterPath) SetElementPositionAt(i int, x float64, y float64) {
	pg_QPainterPath_setElementPositionAt(this.h, int32(i), x, y)
	runtime.KeepAlive(this)
}

func (this *QPainterPath) IsCachingEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_isCachingEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainterPath) SetCachingEnabled(enabled bool) {
	purego.SyscallN(addr_QPainterPath_setCachingEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QPainterPath) Length() float64 {
	__r := pg_QPainterPath_length(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPath) PercentAtLength(len_ float64) float64 {
	__r := pg_QPainterPath_percentAtLength(this.h, len_)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPath) PointAtPercent(t float64) *QPointF {
	__r := pg_QPainterPath_pointAtPercent(this.h, t)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPainterPath) AngleAtPercent(t float64) float64 {
	__r := pg_QPainterPath_angleAtPercent(this.h, t)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPath) SlopeAtPercent(t float64) float64 {
	__r := pg_QPainterPath_slopeAtPercent(this.h, t)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPath) Trimmed(fromFraction float64, toFraction float64) *QPainterPath {
	__r := pg_QPainterPath_trimmed(this.h, fromFraction, toFraction)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) IntersectsWithQPainterPath(p *QPainterPath) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_intersects2, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPainterPath) ContainsWithQPainterPath(p *QPainterPath) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_contains3, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPainterPath) United(r *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_united, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) Intersected(r *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_intersected, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) Subtracted(r *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_subtracted, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) Simplified() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_simplified, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) OperatorEqual(other *QPainterPath) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QPainterPath) OperatorNotEqual(other *QPainterPath) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QPainterPath) OperatorBitwiseAnd(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorBitwiseAnd, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) OperatorBitwiseOr(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorBitwiseOr, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) OperatorPlus(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorPlus, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) OperatorMinus(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorMinus, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) OperatorBitwiseAndAssign(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorBitwiseAndAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPainterPath(__r)
}

func (this *QPainterPath) OperatorBitwiseOrAssign(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorBitwiseOrAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPainterPath(__r)
}

func (this *QPainterPath) OperatorPlusAssign(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorPlusAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPainterPath(__r)
}

func (this *QPainterPath) OperatorMinusAssign(other *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorMinusAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPainterPath(__r)
}

func (this *QPainterPath) AddRoundedRect3(rect *QRectF, xRadius float64, yRadius float64, mode SizeMode) {
	pg_QPainterPath_addRoundedRect3(this.h, uintptr(rect.UnsafePointer()), xRadius, yRadius, int32(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainterPath) AddRoundedRect4(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode SizeMode) {
	pg_QPainterPath_addRoundedRect4(this.h, x, y, w, h, xRadius, yRadius, int32(mode))
	runtime.KeepAlive(this)
}

func (this *QPainterPath) ToSubpathPolygonsWithMatrix(matrix *QTransform) []*QPolygonF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toSubpathPolygons2, this.h, uintptr(matrix.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	__out := make([]*QPolygonF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPolygonF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPainterPath) ToFillPolygonsWithMatrix(matrix *QTransform) []*QPolygonF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toFillPolygons2, this.h, uintptr(matrix.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	__out := make([]*QPolygonF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPolygonF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPainterPath) ToFillPolygonWithMatrix(matrix *QTransform) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_toFillPolygon2, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQPolygonF(__r)
}

func (this *QPainterPath) Trimmed2(fromFraction float64, toFraction float64, offset float64) *QPainterPath {
	__r := pg_QPainterPath_trimmed2(this.h, fromFraction, toFraction, offset)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) OperatorMultiply(m *QTransform) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPath_OperatorMultiply, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQPainterPath(__r)
}

func (this *QPainterPath) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPainterPath_delete(this.h)
	}
	this.UnsafeRelease()
}
