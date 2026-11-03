package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextLine_new uintptr
	addr_QTextLine_isValid uintptr
	addr_QTextLine_rect uintptr
	pg_QTextLine_x func(uintptr) float64
	pg_QTextLine_y func(uintptr) float64
	pg_QTextLine_width func(uintptr) float64
	pg_QTextLine_ascent func(uintptr) float64
	pg_QTextLine_descent func(uintptr) float64
	pg_QTextLine_height func(uintptr) float64
	pg_QTextLine_leading func(uintptr) float64
	addr_QTextLine_setLeadingIncluded uintptr
	addr_QTextLine_leadingIncluded uintptr
	pg_QTextLine_naturalTextWidth func(uintptr) float64
	pg_QTextLine_horizontalAdvance func(uintptr) float64
	addr_QTextLine_naturalTextRect uintptr
	pg_QTextLine_cursorToX func(uintptr, uintptr) float64
	pg_QTextLine_cursorToX2 func(uintptr, int32) float64
	pg_QTextLine_xToCursor func(uintptr, float64) int32
	pg_QTextLine_setLineWidth func(uintptr, float64)
	addr_QTextLine_setNumColumns uintptr
	pg_QTextLine_setNumColumns2 func(uintptr, int32, float64)
	addr_QTextLine_setPosition uintptr
	addr_QTextLine_position uintptr
	addr_QTextLine_textStart uintptr
	addr_QTextLine_textLength uintptr
	addr_QTextLine_lineNumber uintptr
	addr_QTextLine_draw uintptr
	addr_QTextLine_glyphRuns uintptr
	addr_QTextLine_glyphRuns2 uintptr
	pg_QTextLine_cursorToX3 func(uintptr, uintptr, int32) float64
	pg_QTextLine_cursorToX4 func(uintptr, int32, int32) float64
	pg_QTextLine_xToCursor2 func(uintptr, float64, int32) int32
	addr_QTextLine_glyphRuns3 uintptr
	addr_QTextLine_glyphRuns4 uintptr
	pg_QTextLine_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextLine_new, "QTextLine_new")
	symaddr(&addr_QTextLine_isValid, "QTextLine_isValid")
	symaddr(&addr_QTextLine_rect, "QTextLine_rect")
	reg(&pg_QTextLine_x, "QTextLine_x")
	reg(&pg_QTextLine_y, "QTextLine_y")
	reg(&pg_QTextLine_width, "QTextLine_width")
	reg(&pg_QTextLine_ascent, "QTextLine_ascent")
	reg(&pg_QTextLine_descent, "QTextLine_descent")
	reg(&pg_QTextLine_height, "QTextLine_height")
	reg(&pg_QTextLine_leading, "QTextLine_leading")
	symaddr(&addr_QTextLine_setLeadingIncluded, "QTextLine_setLeadingIncluded")
	symaddr(&addr_QTextLine_leadingIncluded, "QTextLine_leadingIncluded")
	reg(&pg_QTextLine_naturalTextWidth, "QTextLine_naturalTextWidth")
	reg(&pg_QTextLine_horizontalAdvance, "QTextLine_horizontalAdvance")
	symaddr(&addr_QTextLine_naturalTextRect, "QTextLine_naturalTextRect")
	reg(&pg_QTextLine_cursorToX, "QTextLine_cursorToX")
	reg(&pg_QTextLine_cursorToX2, "QTextLine_cursorToX2")
	reg(&pg_QTextLine_xToCursor, "QTextLine_xToCursor")
	reg(&pg_QTextLine_setLineWidth, "QTextLine_setLineWidth")
	symaddr(&addr_QTextLine_setNumColumns, "QTextLine_setNumColumns")
	reg(&pg_QTextLine_setNumColumns2, "QTextLine_setNumColumns2")
	symaddr(&addr_QTextLine_setPosition, "QTextLine_setPosition")
	symaddr(&addr_QTextLine_position, "QTextLine_position")
	symaddr(&addr_QTextLine_textStart, "QTextLine_textStart")
	symaddr(&addr_QTextLine_textLength, "QTextLine_textLength")
	symaddr(&addr_QTextLine_lineNumber, "QTextLine_lineNumber")
	symaddr(&addr_QTextLine_draw, "QTextLine_draw")
	symaddr(&addr_QTextLine_glyphRuns, "QTextLine_glyphRuns")
	symaddr(&addr_QTextLine_glyphRuns2, "QTextLine_glyphRuns2")
	reg(&pg_QTextLine_cursorToX3, "QTextLine_cursorToX3")
	reg(&pg_QTextLine_cursorToX4, "QTextLine_cursorToX4")
	reg(&pg_QTextLine_xToCursor2, "QTextLine_xToCursor2")
	symaddr(&addr_QTextLine_glyphRuns3, "QTextLine_glyphRuns3")
	symaddr(&addr_QTextLine_glyphRuns4, "QTextLine_glyphRuns4")
	reg(&pg_QTextLine_delete, "QTextLine_delete")
}

type QTextLine struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextLine) UnsafeGoID() int64     { return this.vid }
func (this *QTextLine) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextLine(h uintptr) *QTextLine {
	if h == 0 {
		return nil
	}
	return &QTextLine{h: h}
}

func newownedQTextLine(h uintptr) *QTextLine {
	o := bindQTextLine(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextLine).Delete)
	}
	return o
}

func (this *QTextLine) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextLine) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextLine(h unsafe.Pointer) *QTextLine { return bindQTextLine(uintptr(h)) }
func UnsafeNewQTextLineOwned(h unsafe.Pointer) *QTextLine { return newownedQTextLine(uintptr(h)) }

func (this *QTextLine) UnsafeRelease() { this.h = 0 }

func NewQTextLine() *QTextLine {
	__r, _, _ := purego.SyscallN(addr_QTextLine_new)
	__o := newownedQTextLine(__r)
	return __o
}

func (this *QTextLine) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLine_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextLine) Rect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QTextLine_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QTextLine) X() float64 {
	__r := pg_QTextLine_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) Y() float64 {
	__r := pg_QTextLine_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) Width() float64 {
	__r := pg_QTextLine_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) Ascent() float64 {
	__r := pg_QTextLine_ascent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) Descent() float64 {
	__r := pg_QTextLine_descent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) Height() float64 {
	__r := pg_QTextLine_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) Leading() float64 {
	__r := pg_QTextLine_leading(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) SetLeadingIncluded(included bool) {
	purego.SyscallN(addr_QTextLine_setLeadingIncluded, this.h, uintptr(b2i(included)))
	runtime.KeepAlive(this)
}

func (this *QTextLine) LeadingIncluded() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLine_leadingIncluded, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextLine) NaturalTextWidth() float64 {
	__r := pg_QTextLine_naturalTextWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) HorizontalAdvance() float64 {
	__r := pg_QTextLine_horizontalAdvance(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) NaturalTextRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QTextLine_naturalTextRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QTextLine) CursorToX() (float64, int) {
	var cursorPos_out int32
	__r := pg_QTextLine_cursorToX(this.h, uintptr(unsafe.Pointer(&cursorPos_out)))
	runtime.KeepAlive(this)
	return __r, int(cursorPos_out)
}

func (this *QTextLine) CursorToXWithCursorPos(cursorPos int) float64 {
	__r := pg_QTextLine_cursorToX2(this.h, int32(cursorPos))
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) XToCursor(x float64) int {
	__r := pg_QTextLine_xToCursor(this.h, x)
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLine) SetLineWidth(width float64) {
	pg_QTextLine_setLineWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextLine) SetNumColumns(columns int) {
	purego.SyscallN(addr_QTextLine_setNumColumns, this.h, uintptr(columns))
	runtime.KeepAlive(this)
}

func (this *QTextLine) SetNumColumns2(columns int, alignmentWidth float64) {
	pg_QTextLine_setNumColumns2(this.h, int32(columns), alignmentWidth)
	runtime.KeepAlive(this)
}

func (this *QTextLine) SetPosition(pos *QPointF) {
	purego.SyscallN(addr_QTextLine_setPosition, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QTextLine) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTextLine_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QTextLine) TextStart() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLine_textStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLine) TextLength() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLine_textLength, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLine) LineNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLine_lineNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLine) Draw(painter *QPainter, position *QPointF) {
	purego.SyscallN(addr_QTextLine_draw, this.h, uintptr(painter.UnsafePointer()), uintptr(position.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(position)
}

func (this *QTextLine) GlyphRuns(from int, length int, flags QTextLayout__GlyphRunRetrievalFlag) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLine_glyphRuns, this.h, uintptr(from), uintptr(length), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLine) GlyphRuns2() []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLine_glyphRuns2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLine) CursorToX2(edge QTextLine__Edge) (float64, int) {
	var cursorPos_out int32
	__r := pg_QTextLine_cursorToX3(this.h, uintptr(unsafe.Pointer(&cursorPos_out)), int32(edge))
	runtime.KeepAlive(this)
	return __r, int(cursorPos_out)
}

func (this *QTextLine) CursorToX3(cursorPos int, edge QTextLine__Edge) float64 {
	__r := pg_QTextLine_cursorToX4(this.h, int32(cursorPos), int32(edge))
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLine) XToCursor2(x float64, arg2 QTextLine__CursorPosition) int {
	__r := pg_QTextLine_xToCursor2(this.h, x, int32(arg2))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLine) GlyphRunsWithFrom(from int) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLine_glyphRuns3, this.h, uintptr(from), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLine) GlyphRuns3(from int, length int) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLine_glyphRuns4, this.h, uintptr(from), uintptr(length), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLine) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextLine_delete(this.h)
	}
	this.UnsafeRelease()
}
