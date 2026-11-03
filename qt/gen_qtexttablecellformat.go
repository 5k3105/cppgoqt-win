package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextTableCellFormat_chain uintptr
	addr_QTextTableCellFormat_new uintptr
	addr_QTextTableCellFormat_new2 uintptr
	addr_QTextTableCellFormat_isValid uintptr
	pg_QTextTableCellFormat_setTopPadding func(uintptr, float64)
	pg_QTextTableCellFormat_topPadding func(uintptr) float64
	pg_QTextTableCellFormat_setBottomPadding func(uintptr, float64)
	pg_QTextTableCellFormat_bottomPadding func(uintptr) float64
	pg_QTextTableCellFormat_setLeftPadding func(uintptr, float64)
	pg_QTextTableCellFormat_leftPadding func(uintptr) float64
	pg_QTextTableCellFormat_setRightPadding func(uintptr, float64)
	pg_QTextTableCellFormat_rightPadding func(uintptr) float64
	pg_QTextTableCellFormat_setPadding func(uintptr, float64)
	pg_QTextTableCellFormat_setTopBorder func(uintptr, float64)
	pg_QTextTableCellFormat_topBorder func(uintptr) float64
	pg_QTextTableCellFormat_setBottomBorder func(uintptr, float64)
	pg_QTextTableCellFormat_bottomBorder func(uintptr) float64
	pg_QTextTableCellFormat_setLeftBorder func(uintptr, float64)
	pg_QTextTableCellFormat_leftBorder func(uintptr) float64
	pg_QTextTableCellFormat_setRightBorder func(uintptr, float64)
	pg_QTextTableCellFormat_rightBorder func(uintptr) float64
	pg_QTextTableCellFormat_setBorder func(uintptr, float64)
	addr_QTextTableCellFormat_setTopBorderStyle uintptr
	addr_QTextTableCellFormat_topBorderStyle uintptr
	addr_QTextTableCellFormat_setBottomBorderStyle uintptr
	addr_QTextTableCellFormat_bottomBorderStyle uintptr
	addr_QTextTableCellFormat_setLeftBorderStyle uintptr
	addr_QTextTableCellFormat_leftBorderStyle uintptr
	addr_QTextTableCellFormat_setRightBorderStyle uintptr
	addr_QTextTableCellFormat_rightBorderStyle uintptr
	addr_QTextTableCellFormat_setBorderStyle uintptr
	addr_QTextTableCellFormat_setTopBorderBrush uintptr
	addr_QTextTableCellFormat_topBorderBrush uintptr
	addr_QTextTableCellFormat_setBottomBorderBrush uintptr
	addr_QTextTableCellFormat_bottomBorderBrush uintptr
	addr_QTextTableCellFormat_setLeftBorderBrush uintptr
	addr_QTextTableCellFormat_leftBorderBrush uintptr
	addr_QTextTableCellFormat_setRightBorderBrush uintptr
	addr_QTextTableCellFormat_rightBorderBrush uintptr
	addr_QTextTableCellFormat_setBorderBrush uintptr
	pg_QTextTableCellFormat_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextTableCellFormat_chain, "QTextTableCellFormat_chain")
	symaddr(&addr_QTextTableCellFormat_new, "QTextTableCellFormat_new")
	symaddr(&addr_QTextTableCellFormat_new2, "QTextTableCellFormat_new2")
	symaddr(&addr_QTextTableCellFormat_isValid, "QTextTableCellFormat_isValid")
	reg(&pg_QTextTableCellFormat_setTopPadding, "QTextTableCellFormat_setTopPadding")
	reg(&pg_QTextTableCellFormat_topPadding, "QTextTableCellFormat_topPadding")
	reg(&pg_QTextTableCellFormat_setBottomPadding, "QTextTableCellFormat_setBottomPadding")
	reg(&pg_QTextTableCellFormat_bottomPadding, "QTextTableCellFormat_bottomPadding")
	reg(&pg_QTextTableCellFormat_setLeftPadding, "QTextTableCellFormat_setLeftPadding")
	reg(&pg_QTextTableCellFormat_leftPadding, "QTextTableCellFormat_leftPadding")
	reg(&pg_QTextTableCellFormat_setRightPadding, "QTextTableCellFormat_setRightPadding")
	reg(&pg_QTextTableCellFormat_rightPadding, "QTextTableCellFormat_rightPadding")
	reg(&pg_QTextTableCellFormat_setPadding, "QTextTableCellFormat_setPadding")
	reg(&pg_QTextTableCellFormat_setTopBorder, "QTextTableCellFormat_setTopBorder")
	reg(&pg_QTextTableCellFormat_topBorder, "QTextTableCellFormat_topBorder")
	reg(&pg_QTextTableCellFormat_setBottomBorder, "QTextTableCellFormat_setBottomBorder")
	reg(&pg_QTextTableCellFormat_bottomBorder, "QTextTableCellFormat_bottomBorder")
	reg(&pg_QTextTableCellFormat_setLeftBorder, "QTextTableCellFormat_setLeftBorder")
	reg(&pg_QTextTableCellFormat_leftBorder, "QTextTableCellFormat_leftBorder")
	reg(&pg_QTextTableCellFormat_setRightBorder, "QTextTableCellFormat_setRightBorder")
	reg(&pg_QTextTableCellFormat_rightBorder, "QTextTableCellFormat_rightBorder")
	reg(&pg_QTextTableCellFormat_setBorder, "QTextTableCellFormat_setBorder")
	symaddr(&addr_QTextTableCellFormat_setTopBorderStyle, "QTextTableCellFormat_setTopBorderStyle")
	symaddr(&addr_QTextTableCellFormat_topBorderStyle, "QTextTableCellFormat_topBorderStyle")
	symaddr(&addr_QTextTableCellFormat_setBottomBorderStyle, "QTextTableCellFormat_setBottomBorderStyle")
	symaddr(&addr_QTextTableCellFormat_bottomBorderStyle, "QTextTableCellFormat_bottomBorderStyle")
	symaddr(&addr_QTextTableCellFormat_setLeftBorderStyle, "QTextTableCellFormat_setLeftBorderStyle")
	symaddr(&addr_QTextTableCellFormat_leftBorderStyle, "QTextTableCellFormat_leftBorderStyle")
	symaddr(&addr_QTextTableCellFormat_setRightBorderStyle, "QTextTableCellFormat_setRightBorderStyle")
	symaddr(&addr_QTextTableCellFormat_rightBorderStyle, "QTextTableCellFormat_rightBorderStyle")
	symaddr(&addr_QTextTableCellFormat_setBorderStyle, "QTextTableCellFormat_setBorderStyle")
	symaddr(&addr_QTextTableCellFormat_setTopBorderBrush, "QTextTableCellFormat_setTopBorderBrush")
	symaddr(&addr_QTextTableCellFormat_topBorderBrush, "QTextTableCellFormat_topBorderBrush")
	symaddr(&addr_QTextTableCellFormat_setBottomBorderBrush, "QTextTableCellFormat_setBottomBorderBrush")
	symaddr(&addr_QTextTableCellFormat_bottomBorderBrush, "QTextTableCellFormat_bottomBorderBrush")
	symaddr(&addr_QTextTableCellFormat_setLeftBorderBrush, "QTextTableCellFormat_setLeftBorderBrush")
	symaddr(&addr_QTextTableCellFormat_leftBorderBrush, "QTextTableCellFormat_leftBorderBrush")
	symaddr(&addr_QTextTableCellFormat_setRightBorderBrush, "QTextTableCellFormat_setRightBorderBrush")
	symaddr(&addr_QTextTableCellFormat_rightBorderBrush, "QTextTableCellFormat_rightBorderBrush")
	symaddr(&addr_QTextTableCellFormat_setBorderBrush, "QTextTableCellFormat_setBorderBrush")
	reg(&pg_QTextTableCellFormat_delete, "QTextTableCellFormat_delete")
}

type QTextTableCellFormat struct {
	h uintptr
	*QTextCharFormat
}

func bindQTextTableCellFormat(h uintptr) *QTextTableCellFormat {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTextTableCellFormat_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextTableCellFormat{h: h, QTextCharFormat: &QTextCharFormat{h: __c[0], QTextFormat: &QTextFormat{h: __c[1]}}}
}

func newownedQTextTableCellFormat(h uintptr) *QTextTableCellFormat {
	o := bindQTextTableCellFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextTableCellFormat).Delete)
	}
	return o
}

func (this *QTextTableCellFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextTableCellFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextTableCellFormat(h unsafe.Pointer) *QTextTableCellFormat { return bindQTextTableCellFormat(uintptr(h)) }
func UnsafeNewQTextTableCellFormatOwned(h unsafe.Pointer) *QTextTableCellFormat { return newownedQTextTableCellFormat(uintptr(h)) }

func (this *QTextTableCellFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextCharFormat != nil {
		this.QTextCharFormat.UnsafeRelease()
	}
}

func NewQTextTableCellFormat() *QTextTableCellFormat {
	__r, _, _ := purego.SyscallN(addr_QTextTableCellFormat_new)
	__o := newownedQTextTableCellFormat(__r)
	return __o
}

func NewQTextTableCellFormat2(arg1 *QTextTableCellFormat) *QTextTableCellFormat {
	__r, _, _ := purego.SyscallN(addr_QTextTableCellFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextTableCellFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextTableCellFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCellFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextTableCellFormat) SetTopPadding(padding float64) {
	pg_QTextTableCellFormat_setTopPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) TopPadding() float64 {
	__r := pg_QTextTableCellFormat_topPadding(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetBottomPadding(padding float64) {
	pg_QTextTableCellFormat_setBottomPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) BottomPadding() float64 {
	__r := pg_QTextTableCellFormat_bottomPadding(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetLeftPadding(padding float64) {
	pg_QTextTableCellFormat_setLeftPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) LeftPadding() float64 {
	__r := pg_QTextTableCellFormat_leftPadding(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetRightPadding(padding float64) {
	pg_QTextTableCellFormat_setRightPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) RightPadding() float64 {
	__r := pg_QTextTableCellFormat_rightPadding(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetPadding(padding float64) {
	pg_QTextTableCellFormat_setPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) SetTopBorder(width float64) {
	pg_QTextTableCellFormat_setTopBorder(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) TopBorder() float64 {
	__r := pg_QTextTableCellFormat_topBorder(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetBottomBorder(width float64) {
	pg_QTextTableCellFormat_setBottomBorder(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) BottomBorder() float64 {
	__r := pg_QTextTableCellFormat_bottomBorder(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetLeftBorder(width float64) {
	pg_QTextTableCellFormat_setLeftBorder(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) LeftBorder() float64 {
	__r := pg_QTextTableCellFormat_leftBorder(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetRightBorder(width float64) {
	pg_QTextTableCellFormat_setRightBorder(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) RightBorder() float64 {
	__r := pg_QTextTableCellFormat_rightBorder(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableCellFormat) SetBorder(width float64) {
	pg_QTextTableCellFormat_setBorder(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) SetTopBorderStyle(style QTextFrameFormat__BorderStyle) {
	purego.SyscallN(addr_QTextTableCellFormat_setTopBorderStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) TopBorderStyle() QTextFrameFormat__BorderStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCellFormat_topBorderStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFrameFormat__BorderStyle(__r)
}

func (this *QTextTableCellFormat) SetBottomBorderStyle(style QTextFrameFormat__BorderStyle) {
	purego.SyscallN(addr_QTextTableCellFormat_setBottomBorderStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) BottomBorderStyle() QTextFrameFormat__BorderStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCellFormat_bottomBorderStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFrameFormat__BorderStyle(__r)
}

func (this *QTextTableCellFormat) SetLeftBorderStyle(style QTextFrameFormat__BorderStyle) {
	purego.SyscallN(addr_QTextTableCellFormat_setLeftBorderStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) LeftBorderStyle() QTextFrameFormat__BorderStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCellFormat_leftBorderStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFrameFormat__BorderStyle(__r)
}

func (this *QTextTableCellFormat) SetRightBorderStyle(style QTextFrameFormat__BorderStyle) {
	purego.SyscallN(addr_QTextTableCellFormat_setRightBorderStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) RightBorderStyle() QTextFrameFormat__BorderStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCellFormat_rightBorderStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFrameFormat__BorderStyle(__r)
}

func (this *QTextTableCellFormat) SetBorderStyle(style QTextFrameFormat__BorderStyle) {
	purego.SyscallN(addr_QTextTableCellFormat_setBorderStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextTableCellFormat) SetTopBorderBrush(brush *QBrush) {
	purego.SyscallN(addr_QTextTableCellFormat_setTopBorderBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextTableCellFormat) TopBorderBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextTableCellFormat_topBorderBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextTableCellFormat) SetBottomBorderBrush(brush *QBrush) {
	purego.SyscallN(addr_QTextTableCellFormat_setBottomBorderBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextTableCellFormat) BottomBorderBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextTableCellFormat_bottomBorderBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextTableCellFormat) SetLeftBorderBrush(brush *QBrush) {
	purego.SyscallN(addr_QTextTableCellFormat_setLeftBorderBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextTableCellFormat) LeftBorderBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextTableCellFormat_leftBorderBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextTableCellFormat) SetRightBorderBrush(brush *QBrush) {
	purego.SyscallN(addr_QTextTableCellFormat_setRightBorderBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextTableCellFormat) RightBorderBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextTableCellFormat_rightBorderBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextTableCellFormat) SetBorderBrush(brush *QBrush) {
	purego.SyscallN(addr_QTextTableCellFormat_setBorderBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextTableCellFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextTableCellFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
