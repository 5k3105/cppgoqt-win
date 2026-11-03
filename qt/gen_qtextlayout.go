package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextLayout_new uintptr
	addr_QTextLayout_new2 uintptr
	addr_QTextLayout_new3 uintptr
	addr_QTextLayout_new4 uintptr
	addr_QTextLayout_new5 uintptr
	addr_QTextLayout_setFont uintptr
	addr_QTextLayout_font uintptr
	addr_QTextLayout_setRawFont uintptr
	addr_QTextLayout_setText uintptr
	addr_QTextLayout_text uintptr
	addr_QTextLayout_setTextOption uintptr
	addr_QTextLayout_textOption uintptr
	addr_QTextLayout_setPreeditArea uintptr
	addr_QTextLayout_preeditAreaPosition uintptr
	addr_QTextLayout_preeditAreaText uintptr
	addr_QTextLayout_setFormats uintptr
	addr_QTextLayout_formats uintptr
	addr_QTextLayout_clearFormats uintptr
	addr_QTextLayout_setCacheEnabled uintptr
	addr_QTextLayout_cacheEnabled uintptr
	addr_QTextLayout_setCursorMoveStyle uintptr
	addr_QTextLayout_cursorMoveStyle uintptr
	addr_QTextLayout_beginLayout uintptr
	addr_QTextLayout_endLayout uintptr
	addr_QTextLayout_clearLayout uintptr
	addr_QTextLayout_createLine uintptr
	addr_QTextLayout_lineCount uintptr
	addr_QTextLayout_lineAt uintptr
	addr_QTextLayout_lineForTextPosition uintptr
	addr_QTextLayout_isValidCursorPosition uintptr
	addr_QTextLayout_nextCursorPosition uintptr
	addr_QTextLayout_previousCursorPosition uintptr
	addr_QTextLayout_leftCursorPosition uintptr
	addr_QTextLayout_rightCursorPosition uintptr
	addr_QTextLayout_draw uintptr
	addr_QTextLayout_drawCursor uintptr
	addr_QTextLayout_drawCursor2 uintptr
	addr_QTextLayout_position uintptr
	addr_QTextLayout_setPosition uintptr
	addr_QTextLayout_boundingRect uintptr
	pg_QTextLayout_minimumWidth func(uintptr) float64
	pg_QTextLayout_maximumWidth func(uintptr) float64
	addr_QTextLayout_glyphRuns uintptr
	addr_QTextLayout_glyphRuns2 uintptr
	addr_QTextLayout_setFlags uintptr
	addr_QTextLayout_nextCursorPosition2 uintptr
	addr_QTextLayout_previousCursorPosition2 uintptr
	addr_QTextLayout_draw2 uintptr
	addr_QTextLayout_draw3 uintptr
	addr_QTextLayout_glyphRuns3 uintptr
	addr_QTextLayout_glyphRuns4 uintptr
	pg_QTextLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextLayout_new, "QTextLayout_new")
	symaddr(&addr_QTextLayout_new2, "QTextLayout_new2")
	symaddr(&addr_QTextLayout_new3, "QTextLayout_new3")
	symaddr(&addr_QTextLayout_new4, "QTextLayout_new4")
	symaddr(&addr_QTextLayout_new5, "QTextLayout_new5")
	symaddr(&addr_QTextLayout_setFont, "QTextLayout_setFont")
	symaddr(&addr_QTextLayout_font, "QTextLayout_font")
	symaddr(&addr_QTextLayout_setRawFont, "QTextLayout_setRawFont")
	symaddr(&addr_QTextLayout_setText, "QTextLayout_setText")
	symaddr(&addr_QTextLayout_text, "QTextLayout_text")
	symaddr(&addr_QTextLayout_setTextOption, "QTextLayout_setTextOption")
	symaddr(&addr_QTextLayout_textOption, "QTextLayout_textOption")
	symaddr(&addr_QTextLayout_setPreeditArea, "QTextLayout_setPreeditArea")
	symaddr(&addr_QTextLayout_preeditAreaPosition, "QTextLayout_preeditAreaPosition")
	symaddr(&addr_QTextLayout_preeditAreaText, "QTextLayout_preeditAreaText")
	symaddr(&addr_QTextLayout_setFormats, "QTextLayout_setFormats")
	symaddr(&addr_QTextLayout_formats, "QTextLayout_formats")
	symaddr(&addr_QTextLayout_clearFormats, "QTextLayout_clearFormats")
	symaddr(&addr_QTextLayout_setCacheEnabled, "QTextLayout_setCacheEnabled")
	symaddr(&addr_QTextLayout_cacheEnabled, "QTextLayout_cacheEnabled")
	symaddr(&addr_QTextLayout_setCursorMoveStyle, "QTextLayout_setCursorMoveStyle")
	symaddr(&addr_QTextLayout_cursorMoveStyle, "QTextLayout_cursorMoveStyle")
	symaddr(&addr_QTextLayout_beginLayout, "QTextLayout_beginLayout")
	symaddr(&addr_QTextLayout_endLayout, "QTextLayout_endLayout")
	symaddr(&addr_QTextLayout_clearLayout, "QTextLayout_clearLayout")
	symaddr(&addr_QTextLayout_createLine, "QTextLayout_createLine")
	symaddr(&addr_QTextLayout_lineCount, "QTextLayout_lineCount")
	symaddr(&addr_QTextLayout_lineAt, "QTextLayout_lineAt")
	symaddr(&addr_QTextLayout_lineForTextPosition, "QTextLayout_lineForTextPosition")
	symaddr(&addr_QTextLayout_isValidCursorPosition, "QTextLayout_isValidCursorPosition")
	symaddr(&addr_QTextLayout_nextCursorPosition, "QTextLayout_nextCursorPosition")
	symaddr(&addr_QTextLayout_previousCursorPosition, "QTextLayout_previousCursorPosition")
	symaddr(&addr_QTextLayout_leftCursorPosition, "QTextLayout_leftCursorPosition")
	symaddr(&addr_QTextLayout_rightCursorPosition, "QTextLayout_rightCursorPosition")
	symaddr(&addr_QTextLayout_draw, "QTextLayout_draw")
	symaddr(&addr_QTextLayout_drawCursor, "QTextLayout_drawCursor")
	symaddr(&addr_QTextLayout_drawCursor2, "QTextLayout_drawCursor2")
	symaddr(&addr_QTextLayout_position, "QTextLayout_position")
	symaddr(&addr_QTextLayout_setPosition, "QTextLayout_setPosition")
	symaddr(&addr_QTextLayout_boundingRect, "QTextLayout_boundingRect")
	reg(&pg_QTextLayout_minimumWidth, "QTextLayout_minimumWidth")
	reg(&pg_QTextLayout_maximumWidth, "QTextLayout_maximumWidth")
	symaddr(&addr_QTextLayout_glyphRuns, "QTextLayout_glyphRuns")
	symaddr(&addr_QTextLayout_glyphRuns2, "QTextLayout_glyphRuns2")
	symaddr(&addr_QTextLayout_setFlags, "QTextLayout_setFlags")
	symaddr(&addr_QTextLayout_nextCursorPosition2, "QTextLayout_nextCursorPosition2")
	symaddr(&addr_QTextLayout_previousCursorPosition2, "QTextLayout_previousCursorPosition2")
	symaddr(&addr_QTextLayout_draw2, "QTextLayout_draw2")
	symaddr(&addr_QTextLayout_draw3, "QTextLayout_draw3")
	symaddr(&addr_QTextLayout_glyphRuns3, "QTextLayout_glyphRuns3")
	symaddr(&addr_QTextLayout_glyphRuns4, "QTextLayout_glyphRuns4")
	reg(&pg_QTextLayout_delete, "QTextLayout_delete")
}

type QTextLayout struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextLayout) UnsafeGoID() int64     { return this.vid }
func (this *QTextLayout) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextLayout(h uintptr) *QTextLayout {
	if h == 0 {
		return nil
	}
	return &QTextLayout{h: h}
}

func newownedQTextLayout(h uintptr) *QTextLayout {
	o := bindQTextLayout(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextLayout).Delete)
	}
	return o
}

func (this *QTextLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextLayout(h unsafe.Pointer) *QTextLayout { return bindQTextLayout(uintptr(h)) }
func UnsafeNewQTextLayoutOwned(h unsafe.Pointer) *QTextLayout { return newownedQTextLayout(uintptr(h)) }

func (this *QTextLayout) UnsafeRelease() { this.h = 0 }

func NewQTextLayout() *QTextLayout {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_new)
	__o := newownedQTextLayout(__r)
	return __o
}

func NewQTextLayout2(text string) *QTextLayout {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextLayout_new2, text_p, uintptr(len(text)))
	__o := newownedQTextLayout(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQTextLayout3(text string, font *QFont) *QTextLayout {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextLayout_new3, text_p, uintptr(len(text)), uintptr(font.UnsafePointer()))
	__o := newownedQTextLayout(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(font)
	return __o
}

func NewQTextLayout4(b *QTextBlock) *QTextLayout {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_new4, uintptr(b.UnsafePointer()))
	__o := newownedQTextLayout(__r)
	runtime.KeepAlive(b)
	return __o
}

func NewQTextLayout5(text string, font *QFont, paintdevice *QPaintDevice) *QTextLayout {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextLayout_new5, text_p, uintptr(len(text)), uintptr(font.UnsafePointer()), uintptr(paintdevice.UnsafePointer()))
	__o := newownedQTextLayout(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(font)
	runtime.KeepAlive(paintdevice)
	return __o
}

func (this *QTextLayout) SetFont(f *QFont) {
	purego.SyscallN(addr_QTextLayout_setFont, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func (this *QTextLayout) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTextLayout) SetRawFont(rawFont *QRawFont) {
	purego.SyscallN(addr_QTextLayout_setRawFont, this.h, uintptr(rawFont.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rawFont)
}

func (this *QTextLayout) SetText(string_ string) {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	purego.SyscallN(addr_QTextLayout_setText, this.h, string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
}

func (this *QTextLayout) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextLayout) SetTextOption(option *QTextOption) {
	purego.SyscallN(addr_QTextLayout_setTextOption, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTextLayout) TextOption() *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_textOption, this.h)
	runtime.KeepAlive(this)
	return bindQTextOption(__r)
}

func (this *QTextLayout) SetPreeditArea(position int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextLayout_setPreeditArea, this.h, uintptr(position), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextLayout) PreeditAreaPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_preeditAreaPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) PreeditAreaText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_preeditAreaText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextLayout) SetFormats(overrides []*QTextLayout__FormatRange) {
	overrides_arr := make([]uintptr, len(overrides))
	for __i, __e := range overrides {
		overrides_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var overrides_p uintptr
	if len(overrides_arr) > 0 {
		overrides_p = uintptr(unsafe.Pointer(&overrides_arr[0]))
	}
	purego.SyscallN(addr_QTextLayout_setFormats, this.h, overrides_p, uintptr(len(overrides)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(overrides)
	runtime.KeepAlive(overrides_arr)
}

func (this *QTextLayout) Formats() []*QTextLayout__FormatRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_formats, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextLayout__FormatRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextLayout__FormatRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLayout) ClearFormats() {
	purego.SyscallN(addr_QTextLayout_clearFormats, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextLayout) SetCacheEnabled(enable bool) {
	purego.SyscallN(addr_QTextLayout_setCacheEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTextLayout) CacheEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_cacheEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextLayout) SetCursorMoveStyle(style CursorMoveStyle) {
	purego.SyscallN(addr_QTextLayout_setCursorMoveStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextLayout) CursorMoveStyle() CursorMoveStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_cursorMoveStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CursorMoveStyle(__r)
}

func (this *QTextLayout) BeginLayout() {
	purego.SyscallN(addr_QTextLayout_beginLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextLayout) EndLayout() {
	purego.SyscallN(addr_QTextLayout_endLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextLayout) ClearLayout() {
	purego.SyscallN(addr_QTextLayout_clearLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextLayout) CreateLine() *QTextLine {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_createLine, this.h)
	runtime.KeepAlive(this)
	return newownedQTextLine(__r)
}

func (this *QTextLayout) LineCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_lineCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) LineAt(i int) *QTextLine {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_lineAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQTextLine(__r)
}

func (this *QTextLayout) LineForTextPosition(pos int) *QTextLine {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_lineForTextPosition, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQTextLine(__r)
}

func (this *QTextLayout) IsValidCursorPosition(pos int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_isValidCursorPosition, this.h, uintptr(pos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextLayout) NextCursorPosition(oldPos int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_nextCursorPosition, this.h, uintptr(oldPos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) PreviousCursorPosition(oldPos int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_previousCursorPosition, this.h, uintptr(oldPos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) LeftCursorPosition(oldPos int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_leftCursorPosition, this.h, uintptr(oldPos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) RightCursorPosition(oldPos int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_rightCursorPosition, this.h, uintptr(oldPos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) Draw(p *QPainter, pos *QPointF) {
	purego.SyscallN(addr_QTextLayout_draw, this.h, uintptr(p.UnsafePointer()), uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pos)
}

func (this *QTextLayout) DrawCursor(p *QPainter, pos *QPointF, cursorPosition int) {
	purego.SyscallN(addr_QTextLayout_drawCursor, this.h, uintptr(p.UnsafePointer()), uintptr(pos.UnsafePointer()), uintptr(cursorPosition))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pos)
}

func (this *QTextLayout) DrawCursor2(p *QPainter, pos *QPointF, cursorPosition int, width int) {
	purego.SyscallN(addr_QTextLayout_drawCursor2, this.h, uintptr(p.UnsafePointer()), uintptr(pos.UnsafePointer()), uintptr(cursorPosition), uintptr(width))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pos)
}

func (this *QTextLayout) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QTextLayout) SetPosition(p *QPointF) {
	purego.SyscallN(addr_QTextLayout_setPosition, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QTextLayout) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QTextLayout_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QTextLayout) MinimumWidth() float64 {
	__r := pg_QTextLayout_minimumWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLayout) MaximumWidth() float64 {
	__r := pg_QTextLayout_maximumWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLayout) GlyphRuns(from int, length int, flags QTextLayout__GlyphRunRetrievalFlag) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_glyphRuns, this.h, uintptr(from), uintptr(length), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLayout) GlyphRuns2() []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_glyphRuns2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLayout) SetFlags(flags int) {
	purego.SyscallN(addr_QTextLayout_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTextLayout) NextCursorPosition2(oldPos int, mode QTextLayout__CursorMode) int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_nextCursorPosition2, this.h, uintptr(oldPos), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) PreviousCursorPosition2(oldPos int, mode QTextLayout__CursorMode) int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout_previousCursorPosition2, this.h, uintptr(oldPos), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout) Draw2(p *QPainter, pos *QPointF, selections []*QTextLayout__FormatRange) {
	selections_arr := make([]uintptr, len(selections))
	for __i, __e := range selections {
		selections_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var selections_p uintptr
	if len(selections_arr) > 0 {
		selections_p = uintptr(unsafe.Pointer(&selections_arr[0]))
	}
	purego.SyscallN(addr_QTextLayout_draw2, this.h, uintptr(p.UnsafePointer()), uintptr(pos.UnsafePointer()), selections_p, uintptr(len(selections)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(selections)
	runtime.KeepAlive(selections_arr)
}

func (this *QTextLayout) Draw3(p *QPainter, pos *QPointF, selections []*QTextLayout__FormatRange, clip *QRectF) {
	selections_arr := make([]uintptr, len(selections))
	for __i, __e := range selections {
		selections_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var selections_p uintptr
	if len(selections_arr) > 0 {
		selections_p = uintptr(unsafe.Pointer(&selections_arr[0]))
	}
	purego.SyscallN(addr_QTextLayout_draw3, this.h, uintptr(p.UnsafePointer()), uintptr(pos.UnsafePointer()), selections_p, uintptr(len(selections)), uintptr(clip.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(selections)
	runtime.KeepAlive(selections_arr)
	runtime.KeepAlive(clip)
}

func (this *QTextLayout) GlyphRunsWithFrom(from int) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_glyphRuns3, this.h, uintptr(from), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLayout) GlyphRuns3(from int, length int) []*QGlyphRun {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextLayout_glyphRuns4, this.h, uintptr(from), uintptr(length), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGlyphRun, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQGlyphRun(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
