package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontMetrics_new uintptr
	addr_QFontMetrics_new2 uintptr
	addr_QFontMetrics_new3 uintptr
	addr_QFontMetrics_OperatorAssign uintptr
	addr_QFontMetrics_swap uintptr
	addr_QFontMetrics_ascent uintptr
	addr_QFontMetrics_capHeight uintptr
	addr_QFontMetrics_descent uintptr
	addr_QFontMetrics_height uintptr
	addr_QFontMetrics_leading uintptr
	addr_QFontMetrics_lineSpacing uintptr
	addr_QFontMetrics_minLeftBearing uintptr
	addr_QFontMetrics_minRightBearing uintptr
	addr_QFontMetrics_maxWidth uintptr
	addr_QFontMetrics_xHeight uintptr
	addr_QFontMetrics_averageCharWidth uintptr
	addr_QFontMetrics_inFont uintptr
	addr_QFontMetrics_inFontUcs4 uintptr
	addr_QFontMetrics_leftBearing uintptr
	addr_QFontMetrics_rightBearing uintptr
	addr_QFontMetrics_horizontalAdvance uintptr
	addr_QFontMetrics_horizontalAdvance2 uintptr
	addr_QFontMetrics_horizontalAdvance3 uintptr
	addr_QFontMetrics_boundingRect uintptr
	addr_QFontMetrics_boundingRect2 uintptr
	addr_QFontMetrics_boundingRect3 uintptr
	addr_QFontMetrics_boundingRect4 uintptr
	addr_QFontMetrics_boundingRect5 uintptr
	addr_QFontMetrics_size uintptr
	addr_QFontMetrics_tightBoundingRect uintptr
	addr_QFontMetrics_tightBoundingRect2 uintptr
	addr_QFontMetrics_elidedText uintptr
	addr_QFontMetrics_underlinePos uintptr
	addr_QFontMetrics_overlinePos uintptr
	addr_QFontMetrics_strikeOutPos uintptr
	addr_QFontMetrics_lineWidth uintptr
	pg_QFontMetrics_fontDpi func(uintptr) float64
	addr_QFontMetrics_OperatorEqual uintptr
	addr_QFontMetrics_OperatorNotEqual uintptr
	addr_QFontMetrics_horizontalAdvance4 uintptr
	addr_QFontMetrics_boundingRect6 uintptr
	addr_QFontMetrics_boundingRect7 uintptr
	addr_QFontMetrics_boundingRect8 uintptr
	addr_QFontMetrics_boundingRect9 uintptr
	addr_QFontMetrics_size2 uintptr
	addr_QFontMetrics_size3 uintptr
	addr_QFontMetrics_elidedText2 uintptr
	pg_QFontMetrics_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontMetrics_new, "QFontMetrics_new")
	symaddr(&addr_QFontMetrics_new2, "QFontMetrics_new2")
	symaddr(&addr_QFontMetrics_new3, "QFontMetrics_new3")
	symaddr(&addr_QFontMetrics_OperatorAssign, "QFontMetrics_OperatorAssign")
	symaddr(&addr_QFontMetrics_swap, "QFontMetrics_swap")
	symaddr(&addr_QFontMetrics_ascent, "QFontMetrics_ascent")
	symaddr(&addr_QFontMetrics_capHeight, "QFontMetrics_capHeight")
	symaddr(&addr_QFontMetrics_descent, "QFontMetrics_descent")
	symaddr(&addr_QFontMetrics_height, "QFontMetrics_height")
	symaddr(&addr_QFontMetrics_leading, "QFontMetrics_leading")
	symaddr(&addr_QFontMetrics_lineSpacing, "QFontMetrics_lineSpacing")
	symaddr(&addr_QFontMetrics_minLeftBearing, "QFontMetrics_minLeftBearing")
	symaddr(&addr_QFontMetrics_minRightBearing, "QFontMetrics_minRightBearing")
	symaddr(&addr_QFontMetrics_maxWidth, "QFontMetrics_maxWidth")
	symaddr(&addr_QFontMetrics_xHeight, "QFontMetrics_xHeight")
	symaddr(&addr_QFontMetrics_averageCharWidth, "QFontMetrics_averageCharWidth")
	symaddr(&addr_QFontMetrics_inFont, "QFontMetrics_inFont")
	symaddr(&addr_QFontMetrics_inFontUcs4, "QFontMetrics_inFontUcs4")
	symaddr(&addr_QFontMetrics_leftBearing, "QFontMetrics_leftBearing")
	symaddr(&addr_QFontMetrics_rightBearing, "QFontMetrics_rightBearing")
	symaddr(&addr_QFontMetrics_horizontalAdvance, "QFontMetrics_horizontalAdvance")
	symaddr(&addr_QFontMetrics_horizontalAdvance2, "QFontMetrics_horizontalAdvance2")
	symaddr(&addr_QFontMetrics_horizontalAdvance3, "QFontMetrics_horizontalAdvance3")
	symaddr(&addr_QFontMetrics_boundingRect, "QFontMetrics_boundingRect")
	symaddr(&addr_QFontMetrics_boundingRect2, "QFontMetrics_boundingRect2")
	symaddr(&addr_QFontMetrics_boundingRect3, "QFontMetrics_boundingRect3")
	symaddr(&addr_QFontMetrics_boundingRect4, "QFontMetrics_boundingRect4")
	symaddr(&addr_QFontMetrics_boundingRect5, "QFontMetrics_boundingRect5")
	symaddr(&addr_QFontMetrics_size, "QFontMetrics_size")
	symaddr(&addr_QFontMetrics_tightBoundingRect, "QFontMetrics_tightBoundingRect")
	symaddr(&addr_QFontMetrics_tightBoundingRect2, "QFontMetrics_tightBoundingRect2")
	symaddr(&addr_QFontMetrics_elidedText, "QFontMetrics_elidedText")
	symaddr(&addr_QFontMetrics_underlinePos, "QFontMetrics_underlinePos")
	symaddr(&addr_QFontMetrics_overlinePos, "QFontMetrics_overlinePos")
	symaddr(&addr_QFontMetrics_strikeOutPos, "QFontMetrics_strikeOutPos")
	symaddr(&addr_QFontMetrics_lineWidth, "QFontMetrics_lineWidth")
	reg(&pg_QFontMetrics_fontDpi, "QFontMetrics_fontDpi")
	symaddr(&addr_QFontMetrics_OperatorEqual, "QFontMetrics_OperatorEqual")
	symaddr(&addr_QFontMetrics_OperatorNotEqual, "QFontMetrics_OperatorNotEqual")
	symaddr(&addr_QFontMetrics_horizontalAdvance4, "QFontMetrics_horizontalAdvance4")
	symaddr(&addr_QFontMetrics_boundingRect6, "QFontMetrics_boundingRect6")
	symaddr(&addr_QFontMetrics_boundingRect7, "QFontMetrics_boundingRect7")
	symaddr(&addr_QFontMetrics_boundingRect8, "QFontMetrics_boundingRect8")
	symaddr(&addr_QFontMetrics_boundingRect9, "QFontMetrics_boundingRect9")
	symaddr(&addr_QFontMetrics_size2, "QFontMetrics_size2")
	symaddr(&addr_QFontMetrics_size3, "QFontMetrics_size3")
	symaddr(&addr_QFontMetrics_elidedText2, "QFontMetrics_elidedText2")
	reg(&pg_QFontMetrics_delete, "QFontMetrics_delete")
}

type QFontMetrics struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFontMetrics) UnsafeGoID() int64     { return this.vid }
func (this *QFontMetrics) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFontMetrics(h uintptr) *QFontMetrics {
	if h == 0 {
		return nil
	}
	return &QFontMetrics{h: h}
}

func newownedQFontMetrics(h uintptr) *QFontMetrics {
	o := bindQFontMetrics(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFontMetrics).Delete)
	}
	return o
}

func (this *QFontMetrics) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontMetrics) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontMetrics(h unsafe.Pointer) *QFontMetrics { return bindQFontMetrics(uintptr(h)) }
func UnsafeNewQFontMetricsOwned(h unsafe.Pointer) *QFontMetrics { return newownedQFontMetrics(uintptr(h)) }

func (this *QFontMetrics) UnsafeRelease() { this.h = 0 }

func NewQFontMetrics(arg1 *QFont) *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQFontMetrics(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQFontMetrics2(font *QFont, pd *QPaintDevice) *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_new2, uintptr(font.UnsafePointer()), uintptr(pd.UnsafePointer()))
	__o := newownedQFontMetrics(__r)
	runtime.KeepAlive(font)
	runtime.KeepAlive(pd)
	return __o
}

func NewQFontMetrics3(arg1 *QFontMetrics) *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQFontMetrics(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QFontMetrics) OperatorAssign(arg1 *QFontMetrics) *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFontMetrics(__r)
}

func (this *QFontMetrics) Swap(other *QFontMetrics) {
	purego.SyscallN(addr_QFontMetrics_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFontMetrics) Ascent() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_ascent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) CapHeight() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_capHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) Descent() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_descent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) Leading() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_leading, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) LineSpacing() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_lineSpacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) MinLeftBearing() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_minLeftBearing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) MinRightBearing() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_minRightBearing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) MaxWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_maxWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) XHeight() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_xHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) AverageCharWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_averageCharWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) InFont(arg1 *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_inFont, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QFontMetrics) InFontUcs4(ucs4 uint) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_inFontUcs4, this.h, uintptr(ucs4))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontMetrics) LeftBearing(arg1 *QChar) int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_leftBearing, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QFontMetrics) RightBearing(arg1 *QChar) int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_rightBearing, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QFontMetrics) HorizontalAdvance(arg1 string) int {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_horizontalAdvance, this.h, arg1_p, uintptr(len(arg1)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QFontMetrics) HorizontalAdvance2(arg1 string, textOption *QTextOption) int {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_horizontalAdvance2, this.h, arg1_p, uintptr(len(arg1)), uintptr(textOption.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(textOption)
	return int(__r)
}

func (this *QFontMetrics) HorizontalAdvanceWithQChar(arg1 *QChar) int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_horizontalAdvance3, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QFontMetrics) BoundingRect(arg1 *QChar) *QRect {
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func (this *QFontMetrics) BoundingRectWithText(text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect2, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QFontMetrics) BoundingRect2(text string, textOption *QTextOption) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect3, this.h, text_p, uintptr(len(text)), uintptr(textOption.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(textOption)
	return newownedQRect(__r)
}

func (this *QFontMetrics) BoundingRect3(r *QRect, flags int, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect4, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QFontMetrics) BoundingRect4(x int, y int, w int, h int, flags int, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect5, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QFontMetrics) Size(flags int, str string) *QSize {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_size, this.h, uintptr(flags), str_p, uintptr(len(str)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return newownedQSize(__r)
}

func (this *QFontMetrics) TightBoundingRect(text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_tightBoundingRect, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QFontMetrics) TightBoundingRect2(text string, textOption *QTextOption) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_tightBoundingRect2, this.h, text_p, uintptr(len(text)), uintptr(textOption.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(textOption)
	return newownedQRect(__r)
}

func (this *QFontMetrics) ElidedText(text string, mode TextElideMode, width int) string {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_elidedText, this.h, text_p, uintptr(len(text)), uintptr(mode), uintptr(width), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return pgStr(__r, __n)
}

func (this *QFontMetrics) UnderlinePos() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_underlinePos, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) OverlinePos() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_overlinePos, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) StrikeOutPos() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_strikeOutPos, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) LineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_lineWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontMetrics) FontDpi() float64 {
	__r := pg_QFontMetrics_fontDpi(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetrics) OperatorEqual(other *QFontMetrics) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QFontMetrics) OperatorNotEqual(other *QFontMetrics) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QFontMetrics) HorizontalAdvance3(arg1 string, len_ int) int {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontMetrics_horizontalAdvance4, this.h, arg1_p, uintptr(len(arg1)), uintptr(len_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QFontMetrics) BoundingRect5(r *QRect, flags int, text string, tabstops int) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect6, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)), uintptr(tabstops))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QFontMetrics) BoundingRect6(r *QRect, flags int, text string, tabstops int) (*QRect, int) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var tabarray_out int32
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect7, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)), uintptr(tabstops), uintptr(unsafe.Pointer(&tabarray_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r), int(tabarray_out)
}

func (this *QFontMetrics) BoundingRect7(x int, y int, w int, h int, flags int, text string, tabstops int) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect8, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(flags), text_p, uintptr(len(text)), uintptr(tabstops))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QFontMetrics) BoundingRect8(x int, y int, w int, h int, flags int, text string, tabstops int) (*QRect, int) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var tabarray_out int32
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_boundingRect9, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(flags), text_p, uintptr(len(text)), uintptr(tabstops), uintptr(unsafe.Pointer(&tabarray_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRect(__r), int(tabarray_out)
}

func (this *QFontMetrics) Size2(flags int, str string, tabstops int) *QSize {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_size2, this.h, uintptr(flags), str_p, uintptr(len(str)), uintptr(tabstops))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return newownedQSize(__r)
}

func (this *QFontMetrics) Size3(flags int, str string, tabstops int) (*QSize, int) {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var tabarray_out int32
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_size3, this.h, uintptr(flags), str_p, uintptr(len(str)), uintptr(tabstops), uintptr(unsafe.Pointer(&tabarray_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return newownedQSize(__r), int(tabarray_out)
}

func (this *QFontMetrics) ElidedText2(text string, mode TextElideMode, width int, flags int) string {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontMetrics_elidedText2, this.h, text_p, uintptr(len(text)), uintptr(mode), uintptr(width), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return pgStr(__r, __n)
}

func (this *QFontMetrics) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontMetrics_delete(this.h)
	}
	this.UnsafeRelease()
}
