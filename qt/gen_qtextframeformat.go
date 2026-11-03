package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextFrameFormat_base func(uintptr) uintptr
	addr_QTextFrameFormat_new uintptr
	addr_QTextFrameFormat_new2 uintptr
	addr_QTextFrameFormat_isValid uintptr
	addr_QTextFrameFormat_setPosition uintptr
	addr_QTextFrameFormat_position uintptr
	pg_QTextFrameFormat_setBorder func(uintptr, float64)
	pg_QTextFrameFormat_border func(uintptr) float64
	addr_QTextFrameFormat_setBorderBrush uintptr
	addr_QTextFrameFormat_borderBrush uintptr
	addr_QTextFrameFormat_setBorderStyle uintptr
	addr_QTextFrameFormat_borderStyle uintptr
	pg_QTextFrameFormat_setMargin func(uintptr, float64)
	pg_QTextFrameFormat_margin func(uintptr) float64
	pg_QTextFrameFormat_setTopMargin func(uintptr, float64)
	pg_QTextFrameFormat_topMargin func(uintptr) float64
	pg_QTextFrameFormat_setBottomMargin func(uintptr, float64)
	pg_QTextFrameFormat_bottomMargin func(uintptr) float64
	pg_QTextFrameFormat_setLeftMargin func(uintptr, float64)
	pg_QTextFrameFormat_leftMargin func(uintptr) float64
	pg_QTextFrameFormat_setRightMargin func(uintptr, float64)
	pg_QTextFrameFormat_rightMargin func(uintptr) float64
	pg_QTextFrameFormat_setPadding func(uintptr, float64)
	pg_QTextFrameFormat_padding func(uintptr) float64
	pg_QTextFrameFormat_setWidth func(uintptr, float64)
	addr_QTextFrameFormat_setWidth2 uintptr
	addr_QTextFrameFormat_width uintptr
	pg_QTextFrameFormat_setHeight func(uintptr, float64)
	addr_QTextFrameFormat_setHeight2 uintptr
	addr_QTextFrameFormat_height uintptr
	addr_QTextFrameFormat_setPageBreakPolicy uintptr
	addr_QTextFrameFormat_pageBreakPolicy uintptr
	addr_QTextFrameFormat_OperatorAssign uintptr
	pg_QTextFrameFormat_delete func(uintptr)
)

func init() {
	reg(&pg_QTextFrameFormat_base, "QTextFrameFormat_base")
	symaddr(&addr_QTextFrameFormat_new, "QTextFrameFormat_new")
	symaddr(&addr_QTextFrameFormat_new2, "QTextFrameFormat_new2")
	symaddr(&addr_QTextFrameFormat_isValid, "QTextFrameFormat_isValid")
	symaddr(&addr_QTextFrameFormat_setPosition, "QTextFrameFormat_setPosition")
	symaddr(&addr_QTextFrameFormat_position, "QTextFrameFormat_position")
	reg(&pg_QTextFrameFormat_setBorder, "QTextFrameFormat_setBorder")
	reg(&pg_QTextFrameFormat_border, "QTextFrameFormat_border")
	symaddr(&addr_QTextFrameFormat_setBorderBrush, "QTextFrameFormat_setBorderBrush")
	symaddr(&addr_QTextFrameFormat_borderBrush, "QTextFrameFormat_borderBrush")
	symaddr(&addr_QTextFrameFormat_setBorderStyle, "QTextFrameFormat_setBorderStyle")
	symaddr(&addr_QTextFrameFormat_borderStyle, "QTextFrameFormat_borderStyle")
	reg(&pg_QTextFrameFormat_setMargin, "QTextFrameFormat_setMargin")
	reg(&pg_QTextFrameFormat_margin, "QTextFrameFormat_margin")
	reg(&pg_QTextFrameFormat_setTopMargin, "QTextFrameFormat_setTopMargin")
	reg(&pg_QTextFrameFormat_topMargin, "QTextFrameFormat_topMargin")
	reg(&pg_QTextFrameFormat_setBottomMargin, "QTextFrameFormat_setBottomMargin")
	reg(&pg_QTextFrameFormat_bottomMargin, "QTextFrameFormat_bottomMargin")
	reg(&pg_QTextFrameFormat_setLeftMargin, "QTextFrameFormat_setLeftMargin")
	reg(&pg_QTextFrameFormat_leftMargin, "QTextFrameFormat_leftMargin")
	reg(&pg_QTextFrameFormat_setRightMargin, "QTextFrameFormat_setRightMargin")
	reg(&pg_QTextFrameFormat_rightMargin, "QTextFrameFormat_rightMargin")
	reg(&pg_QTextFrameFormat_setPadding, "QTextFrameFormat_setPadding")
	reg(&pg_QTextFrameFormat_padding, "QTextFrameFormat_padding")
	reg(&pg_QTextFrameFormat_setWidth, "QTextFrameFormat_setWidth")
	symaddr(&addr_QTextFrameFormat_setWidth2, "QTextFrameFormat_setWidth2")
	symaddr(&addr_QTextFrameFormat_width, "QTextFrameFormat_width")
	reg(&pg_QTextFrameFormat_setHeight, "QTextFrameFormat_setHeight")
	symaddr(&addr_QTextFrameFormat_setHeight2, "QTextFrameFormat_setHeight2")
	symaddr(&addr_QTextFrameFormat_height, "QTextFrameFormat_height")
	symaddr(&addr_QTextFrameFormat_setPageBreakPolicy, "QTextFrameFormat_setPageBreakPolicy")
	symaddr(&addr_QTextFrameFormat_pageBreakPolicy, "QTextFrameFormat_pageBreakPolicy")
	symaddr(&addr_QTextFrameFormat_OperatorAssign, "QTextFrameFormat_OperatorAssign")
	reg(&pg_QTextFrameFormat_delete, "QTextFrameFormat_delete")
}

type QTextFrameFormat struct {
	h uintptr
	*QTextFormat
}

func bindQTextFrameFormat(h uintptr) *QTextFrameFormat {
	if h == 0 {
		return nil
	}
	return &QTextFrameFormat{h: h, QTextFormat: bindQTextFormat(pg_QTextFrameFormat_base(h))}
}

func newownedQTextFrameFormat(h uintptr) *QTextFrameFormat {
	o := bindQTextFrameFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextFrameFormat).Delete)
	}
	return o
}

func (this *QTextFrameFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextFrameFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextFrameFormat(h unsafe.Pointer) *QTextFrameFormat { return bindQTextFrameFormat(uintptr(h)) }
func UnsafeNewQTextFrameFormatOwned(h unsafe.Pointer) *QTextFrameFormat { return newownedQTextFrameFormat(uintptr(h)) }

func (this *QTextFrameFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextFormat != nil {
		this.QTextFormat.UnsafeRelease()
	}
}

func NewQTextFrameFormat() *QTextFrameFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFrameFormat_new)
	__o := newownedQTextFrameFormat(__r)
	return __o
}

func NewQTextFrameFormat2(arg1 *QTextFrameFormat) *QTextFrameFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFrameFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextFrameFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextFrameFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFrameFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFrameFormat) SetPosition(f QTextFrameFormat__Position) {
	purego.SyscallN(addr_QTextFrameFormat_setPosition, this.h, uintptr(f))
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) Position() QTextFrameFormat__Position {
	__r0, _, _ := purego.SyscallN(addr_QTextFrameFormat_position, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFrameFormat__Position(__r)
}

func (this *QTextFrameFormat) SetBorder(border float64) {
	pg_QTextFrameFormat_setBorder(this.h, border)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) Border() float64 {
	__r := pg_QTextFrameFormat_border(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetBorderBrush(brush *QBrush) {
	purego.SyscallN(addr_QTextFrameFormat_setBorderBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextFrameFormat) BorderBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextFrameFormat_borderBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextFrameFormat) SetBorderStyle(style QTextFrameFormat__BorderStyle) {
	purego.SyscallN(addr_QTextFrameFormat_setBorderStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) BorderStyle() QTextFrameFormat__BorderStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextFrameFormat_borderStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFrameFormat__BorderStyle(__r)
}

func (this *QTextFrameFormat) SetMargin(margin float64) {
	pg_QTextFrameFormat_setMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) Margin() float64 {
	__r := pg_QTextFrameFormat_margin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetTopMargin(margin float64) {
	pg_QTextFrameFormat_setTopMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) TopMargin() float64 {
	__r := pg_QTextFrameFormat_topMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetBottomMargin(margin float64) {
	pg_QTextFrameFormat_setBottomMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) BottomMargin() float64 {
	__r := pg_QTextFrameFormat_bottomMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetLeftMargin(margin float64) {
	pg_QTextFrameFormat_setLeftMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) LeftMargin() float64 {
	__r := pg_QTextFrameFormat_leftMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetRightMargin(margin float64) {
	pg_QTextFrameFormat_setRightMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) RightMargin() float64 {
	__r := pg_QTextFrameFormat_rightMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetPadding(padding float64) {
	pg_QTextFrameFormat_setPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) Padding() float64 {
	__r := pg_QTextFrameFormat_padding(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFrameFormat) SetWidth(width float64) {
	pg_QTextFrameFormat_setWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) SetWidthWithLength(length *QTextLength) {
	purego.SyscallN(addr_QTextFrameFormat_setWidth2, this.h, uintptr(length.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(length)
}

func (this *QTextFrameFormat) Width() *QTextLength {
	__r, _, _ := purego.SyscallN(addr_QTextFrameFormat_width, this.h)
	runtime.KeepAlive(this)
	return newownedQTextLength(__r)
}

func (this *QTextFrameFormat) SetHeight(height float64) {
	pg_QTextFrameFormat_setHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) SetHeightWithHeight(height *QTextLength) {
	purego.SyscallN(addr_QTextFrameFormat_setHeight2, this.h, uintptr(height.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(height)
}

func (this *QTextFrameFormat) Height() *QTextLength {
	__r, _, _ := purego.SyscallN(addr_QTextFrameFormat_height, this.h)
	runtime.KeepAlive(this)
	return newownedQTextLength(__r)
}

func (this *QTextFrameFormat) SetPageBreakPolicy(flags QTextFormat__PageBreakFlag) {
	purego.SyscallN(addr_QTextFrameFormat_setPageBreakPolicy, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTextFrameFormat) PageBreakPolicy() QTextFormat__PageBreakFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextFrameFormat_pageBreakPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFormat__PageBreakFlag(__r)
}

func (this *QTextFrameFormat) OperatorAssign(arg1 *QTextFrameFormat) *QTextFrameFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFrameFormat_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextFrameFormat(__r)
}

func (this *QTextFrameFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextFrameFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
