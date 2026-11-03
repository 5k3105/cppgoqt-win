package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontMetricsF_new uintptr
	addr_QFontMetricsF_new2 uintptr
	addr_QFontMetricsF_new3 uintptr
	addr_QFontMetricsF_new4 uintptr
	addr_QFontMetricsF_OperatorAssign uintptr
	addr_QFontMetricsF_OperatorAssign2 uintptr
	addr_QFontMetricsF_swap uintptr
	pg_QFontMetricsF_ascent func(uintptr) float64
	pg_QFontMetricsF_capHeight func(uintptr) float64
	pg_QFontMetricsF_descent func(uintptr) float64
	pg_QFontMetricsF_height func(uintptr) float64
	pg_QFontMetricsF_leading func(uintptr) float64
	pg_QFontMetricsF_lineSpacing func(uintptr) float64
	pg_QFontMetricsF_minLeftBearing func(uintptr) float64
	pg_QFontMetricsF_minRightBearing func(uintptr) float64
	pg_QFontMetricsF_maxWidth func(uintptr) float64
	pg_QFontMetricsF_xHeight func(uintptr) float64
	pg_QFontMetricsF_averageCharWidth func(uintptr) float64
	addr_QFontMetricsF_inFont uintptr
	addr_QFontMetricsF_inFontUcs4 uintptr
	pg_QFontMetricsF_leftBearing func(uintptr, uintptr) float64
	pg_QFontMetricsF_rightBearing func(uintptr, uintptr) float64
	pg_QFontMetricsF_horizontalAdvance func(uintptr, string, int64) float64
	pg_QFontMetricsF_horizontalAdvance2 func(uintptr, uintptr) float64
	pg_QFontMetricsF_horizontalAdvance3 func(uintptr, string, int64, uintptr) float64
	addr_QFontMetricsF_boundingRect uintptr
	addr_QFontMetricsF_boundingRect2 uintptr
	addr_QFontMetricsF_boundingRect3 uintptr
	addr_QFontMetricsF_boundingRect4 uintptr
	addr_QFontMetricsF_size uintptr
	addr_QFontMetricsF_tightBoundingRect uintptr
	addr_QFontMetricsF_tightBoundingRect2 uintptr
	pg_QFontMetricsF_elidedText func(uintptr, string, int64, int32, float64, *int64) uintptr
	pg_QFontMetricsF_underlinePos func(uintptr) float64
	pg_QFontMetricsF_overlinePos func(uintptr) float64
	pg_QFontMetricsF_strikeOutPos func(uintptr) float64
	pg_QFontMetricsF_lineWidth func(uintptr) float64
	pg_QFontMetricsF_fontDpi func(uintptr) float64
	addr_QFontMetricsF_OperatorEqual uintptr
	addr_QFontMetricsF_OperatorNotEqual uintptr
	pg_QFontMetricsF_horizontalAdvance4 func(uintptr, string, int64, int32) float64
	addr_QFontMetricsF_boundingRect5 uintptr
	addr_QFontMetricsF_boundingRect6 uintptr
	addr_QFontMetricsF_size2 uintptr
	addr_QFontMetricsF_size3 uintptr
	pg_QFontMetricsF_elidedText2 func(uintptr, string, int64, int32, float64, int32, *int64) uintptr
	pg_QFontMetricsF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontMetricsF_new, "QFontMetricsF_new")
	symaddr(&addr_QFontMetricsF_new2, "QFontMetricsF_new2")
	symaddr(&addr_QFontMetricsF_new3, "QFontMetricsF_new3")
	symaddr(&addr_QFontMetricsF_new4, "QFontMetricsF_new4")
	symaddr(&addr_QFontMetricsF_OperatorAssign, "QFontMetricsF_OperatorAssign")
	symaddr(&addr_QFontMetricsF_OperatorAssign2, "QFontMetricsF_OperatorAssign2")
	symaddr(&addr_QFontMetricsF_swap, "QFontMetricsF_swap")
	reg(&pg_QFontMetricsF_ascent, "QFontMetricsF_ascent")
	reg(&pg_QFontMetricsF_capHeight, "QFontMetricsF_capHeight")
	reg(&pg_QFontMetricsF_descent, "QFontMetricsF_descent")
	reg(&pg_QFontMetricsF_height, "QFontMetricsF_height")
	reg(&pg_QFontMetricsF_leading, "QFontMetricsF_leading")
	reg(&pg_QFontMetricsF_lineSpacing, "QFontMetricsF_lineSpacing")
	reg(&pg_QFontMetricsF_minLeftBearing, "QFontMetricsF_minLeftBearing")
	reg(&pg_QFontMetricsF_minRightBearing, "QFontMetricsF_minRightBearing")
	reg(&pg_QFontMetricsF_maxWidth, "QFontMetricsF_maxWidth")
	reg(&pg_QFontMetricsF_xHeight, "QFontMetricsF_xHeight")
	reg(&pg_QFontMetricsF_averageCharWidth, "QFontMetricsF_averageCharWidth")
	symaddr(&addr_QFontMetricsF_inFont, "QFontMetricsF_inFont")
	symaddr(&addr_QFontMetricsF_inFontUcs4, "QFontMetricsF_inFontUcs4")
	reg(&pg_QFontMetricsF_leftBearing, "QFontMetricsF_leftBearing")
	reg(&pg_QFontMetricsF_rightBearing, "QFontMetricsF_rightBearing")
	reg(&pg_QFontMetricsF_horizontalAdvance, "QFontMetricsF_horizontalAdvance")
	reg(&pg_QFontMetricsF_horizontalAdvance2, "QFontMetricsF_horizontalAdvance2")
	reg(&pg_QFontMetricsF_horizontalAdvance3, "QFontMetricsF_horizontalAdvance3")
	symaddr(&addr_QFontMetricsF_boundingRect, "QFontMetricsF_boundingRect")
	symaddr(&addr_QFontMetricsF_boundingRect2, "QFontMetricsF_boundingRect2")
	symaddr(&addr_QFontMetricsF_boundingRect3, "QFontMetricsF_boundingRect3")
	symaddr(&addr_QFontMetricsF_boundingRect4, "QFontMetricsF_boundingRect4")
	symaddr(&addr_QFontMetricsF_size, "QFontMetricsF_size")
	symaddr(&addr_QFontMetricsF_tightBoundingRect, "QFontMetricsF_tightBoundingRect")
	symaddr(&addr_QFontMetricsF_tightBoundingRect2, "QFontMetricsF_tightBoundingRect2")
	reg(&pg_QFontMetricsF_elidedText, "QFontMetricsF_elidedText")
	reg(&pg_QFontMetricsF_underlinePos, "QFontMetricsF_underlinePos")
	reg(&pg_QFontMetricsF_overlinePos, "QFontMetricsF_overlinePos")
	reg(&pg_QFontMetricsF_strikeOutPos, "QFontMetricsF_strikeOutPos")
	reg(&pg_QFontMetricsF_lineWidth, "QFontMetricsF_lineWidth")
	reg(&pg_QFontMetricsF_fontDpi, "QFontMetricsF_fontDpi")
	symaddr(&addr_QFontMetricsF_OperatorEqual, "QFontMetricsF_OperatorEqual")
	symaddr(&addr_QFontMetricsF_OperatorNotEqual, "QFontMetricsF_OperatorNotEqual")
	reg(&pg_QFontMetricsF_horizontalAdvance4, "QFontMetricsF_horizontalAdvance4")
	symaddr(&addr_QFontMetricsF_boundingRect5, "QFontMetricsF_boundingRect5")
	symaddr(&addr_QFontMetricsF_boundingRect6, "QFontMetricsF_boundingRect6")
	symaddr(&addr_QFontMetricsF_size2, "QFontMetricsF_size2")
	symaddr(&addr_QFontMetricsF_size3, "QFontMetricsF_size3")
	reg(&pg_QFontMetricsF_elidedText2, "QFontMetricsF_elidedText2")
	reg(&pg_QFontMetricsF_delete, "QFontMetricsF_delete")
}

type QFontMetricsF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFontMetricsF) UnsafeGoID() int64     { return this.vid }
func (this *QFontMetricsF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFontMetricsF(h uintptr) *QFontMetricsF {
	if h == 0 {
		return nil
	}
	return &QFontMetricsF{h: h}
}

func newownedQFontMetricsF(h uintptr) *QFontMetricsF {
	o := bindQFontMetricsF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFontMetricsF).Delete)
	}
	return o
}

func (this *QFontMetricsF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontMetricsF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontMetricsF(h unsafe.Pointer) *QFontMetricsF { return bindQFontMetricsF(uintptr(h)) }
func UnsafeNewQFontMetricsFOwned(h unsafe.Pointer) *QFontMetricsF { return newownedQFontMetricsF(uintptr(h)) }

func (this *QFontMetricsF) UnsafeRelease() { this.h = 0 }

func NewQFontMetricsF(font *QFont) *QFontMetricsF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_new, uintptr(font.UnsafePointer()))
	__o := newownedQFontMetricsF(__r)
	runtime.KeepAlive(font)
	return __o
}

func NewQFontMetricsF2(font *QFont, pd *QPaintDevice) *QFontMetricsF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_new2, uintptr(font.UnsafePointer()), uintptr(pd.UnsafePointer()))
	__o := newownedQFontMetricsF(__r)
	runtime.KeepAlive(font)
	runtime.KeepAlive(pd)
	return __o
}

func NewQFontMetricsF3(arg1 *QFontMetrics) *QFontMetricsF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQFontMetricsF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQFontMetricsF4(arg1 *QFontMetricsF) *QFontMetricsF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQFontMetricsF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QFontMetricsF) OperatorAssign(arg1 *QFontMetricsF) *QFontMetricsF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFontMetricsF(__r)
}

func (this *QFontMetricsF) OperatorAssignWithQFontMetrics(arg1 *QFontMetrics) *QFontMetricsF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_OperatorAssign2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFontMetricsF(__r)
}

func (this *QFontMetricsF) Swap(other *QFontMetricsF) {
	purego.SyscallN(addr_QFontMetricsF_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFontMetricsF) Ascent() float64 {
	__r := pg_QFontMetricsF_ascent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) CapHeight() float64 {
	__r := pg_QFontMetricsF_capHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) Descent() float64 {
	__r := pg_QFontMetricsF_descent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) Height() float64 {
	__r := pg_QFontMetricsF_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) Leading() float64 {
	__r := pg_QFontMetricsF_leading(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) LineSpacing() float64 {
	__r := pg_QFontMetricsF_lineSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) MinLeftBearing() float64 {
	__r := pg_QFontMetricsF_minLeftBearing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) MinRightBearing() float64 {
	__r := pg_QFontMetricsF_minRightBearing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) MaxWidth() float64 {
	__r := pg_QFontMetricsF_maxWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) XHeight() float64 {
	__r := pg_QFontMetricsF_xHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) AverageCharWidth() float64 {
	__r := pg_QFontMetricsF_averageCharWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) InFont(arg1 *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetricsF_inFont, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QFontMetricsF) InFontUcs4(ucs4 uint) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetricsF_inFontUcs4, this.h, uintptr(ucs4))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontMetricsF) LeftBearing(arg1 *QChar) float64 {
	__r := pg_QFontMetricsF_leftBearing(this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func (this *QFontMetricsF) RightBearing(arg1 *QChar) float64 {
	__r := pg_QFontMetricsF_rightBearing(this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func (this *QFontMetricsF) HorizontalAdvance(string_ string) float64 {
	__r := pg_QFontMetricsF_horizontalAdvance(this.h, string_, int64(len(string_)))
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) HorizontalAdvanceWithQChar(arg1 *QChar) float64 {
	__r := pg_QFontMetricsF_horizontalAdvance2(this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func (this *QFontMetricsF) HorizontalAdvance2(string_ string, textOption *QTextOption) float64 {
	__r := pg_QFontMetricsF_horizontalAdvance3(this.h, string_, int64(len(string_)), uintptr(textOption.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(textOption)
	return __r
}

func (this *QFontMetricsF) BoundingRect(string_ string) *QRectF {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_boundingRect, this.h, string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) BoundingRect2(text string, textOption *QTextOption) *QRectF {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_boundingRect2, this.h, text_p, uintptr(len(text)), uintptr(textOption.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(textOption)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) BoundingRectWithQChar(arg1 *QChar) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_boundingRect3, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) BoundingRect3(r *QRectF, flags int, string_ string) *QRectF {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_boundingRect4, this.h, uintptr(r.UnsafePointer()), uintptr(flags), string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(string_)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) Size(flags int, str string) *QSizeF {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_size, this.h, uintptr(flags), str_p, uintptr(len(str)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return newownedQSizeF(__r)
}

func (this *QFontMetricsF) TightBoundingRect(text string) *QRectF {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_tightBoundingRect, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) TightBoundingRect2(text string, textOption *QTextOption) *QRectF {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_tightBoundingRect2, this.h, text_p, uintptr(len(text)), uintptr(textOption.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(textOption)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) ElidedText(text string, mode TextElideMode, width float64) string {
	var __n int64
	__r := pg_QFontMetricsF_elidedText(this.h, text, int64(len(text)), int32(mode), width, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFontMetricsF) UnderlinePos() float64 {
	__r := pg_QFontMetricsF_underlinePos(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) OverlinePos() float64 {
	__r := pg_QFontMetricsF_overlinePos(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) StrikeOutPos() float64 {
	__r := pg_QFontMetricsF_strikeOutPos(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) LineWidth() float64 {
	__r := pg_QFontMetricsF_lineWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) FontDpi() float64 {
	__r := pg_QFontMetricsF_fontDpi(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) OperatorEqual(other *QFontMetricsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetricsF_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QFontMetricsF) OperatorNotEqual(other *QFontMetricsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontMetricsF_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QFontMetricsF) HorizontalAdvance3(string_ string, length int) float64 {
	__r := pg_QFontMetricsF_horizontalAdvance4(this.h, string_, int64(len(string_)), int32(length))
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontMetricsF) BoundingRect4(r *QRectF, flags int, string_ string, tabstops int) *QRectF {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_boundingRect5, this.h, uintptr(r.UnsafePointer()), uintptr(flags), string__p, uintptr(len(string_)), uintptr(tabstops))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(string_)
	return newownedQRectF(__r)
}

func (this *QFontMetricsF) BoundingRect5(r *QRectF, flags int, string_ string, tabstops int) (*QRectF, int) {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var tabarray_out int32
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_boundingRect6, this.h, uintptr(r.UnsafePointer()), uintptr(flags), string__p, uintptr(len(string_)), uintptr(tabstops), uintptr(unsafe.Pointer(&tabarray_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(string_)
	return newownedQRectF(__r), int(tabarray_out)
}

func (this *QFontMetricsF) Size2(flags int, str string, tabstops int) *QSizeF {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_size2, this.h, uintptr(flags), str_p, uintptr(len(str)), uintptr(tabstops))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return newownedQSizeF(__r)
}

func (this *QFontMetricsF) Size3(flags int, str string, tabstops int) (*QSizeF, int) {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var tabarray_out int32
	__r, _, _ := purego.SyscallN(addr_QFontMetricsF_size3, this.h, uintptr(flags), str_p, uintptr(len(str)), uintptr(tabstops), uintptr(unsafe.Pointer(&tabarray_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return newownedQSizeF(__r), int(tabarray_out)
}

func (this *QFontMetricsF) ElidedText2(text string, mode TextElideMode, width float64, flags int) string {
	var __n int64
	__r := pg_QFontMetricsF_elidedText2(this.h, text, int64(len(text)), int32(mode), width, int32(flags), &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFontMetricsF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontMetricsF_delete(this.h)
	}
	this.UnsafeRelease()
}
