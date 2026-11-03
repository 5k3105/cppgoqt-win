package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextBlockFormat_base func(uintptr) uintptr
	addr_QTextBlockFormat_new uintptr
	addr_QTextBlockFormat_new2 uintptr
	addr_QTextBlockFormat_isValid uintptr
	addr_QTextBlockFormat_setAlignment uintptr
	addr_QTextBlockFormat_alignment uintptr
	pg_QTextBlockFormat_setTopMargin func(uintptr, float64)
	pg_QTextBlockFormat_topMargin func(uintptr) float64
	pg_QTextBlockFormat_setBottomMargin func(uintptr, float64)
	pg_QTextBlockFormat_bottomMargin func(uintptr) float64
	pg_QTextBlockFormat_setLeftMargin func(uintptr, float64)
	pg_QTextBlockFormat_leftMargin func(uintptr) float64
	pg_QTextBlockFormat_setRightMargin func(uintptr, float64)
	pg_QTextBlockFormat_rightMargin func(uintptr) float64
	pg_QTextBlockFormat_setTextIndent func(uintptr, float64)
	pg_QTextBlockFormat_textIndent func(uintptr) float64
	addr_QTextBlockFormat_setIndent uintptr
	addr_QTextBlockFormat_indent uintptr
	addr_QTextBlockFormat_setHeadingLevel uintptr
	addr_QTextBlockFormat_headingLevel uintptr
	pg_QTextBlockFormat_setLineHeight func(uintptr, float64, int32)
	pg_QTextBlockFormat_lineHeight func(uintptr, float64, float64) float64
	pg_QTextBlockFormat_lineHeight2 func(uintptr) float64
	addr_QTextBlockFormat_lineHeightType uintptr
	addr_QTextBlockFormat_setNonBreakableLines uintptr
	addr_QTextBlockFormat_nonBreakableLines uintptr
	addr_QTextBlockFormat_setPageBreakPolicy uintptr
	addr_QTextBlockFormat_pageBreakPolicy uintptr
	addr_QTextBlockFormat_setTabPositions uintptr
	addr_QTextBlockFormat_tabPositions uintptr
	addr_QTextBlockFormat_setMarker uintptr
	addr_QTextBlockFormat_marker uintptr
	addr_QTextBlockFormat_OperatorAssign uintptr
	pg_QTextBlockFormat_delete func(uintptr)
)

func init() {
	reg(&pg_QTextBlockFormat_base, "QTextBlockFormat_base")
	symaddr(&addr_QTextBlockFormat_new, "QTextBlockFormat_new")
	symaddr(&addr_QTextBlockFormat_new2, "QTextBlockFormat_new2")
	symaddr(&addr_QTextBlockFormat_isValid, "QTextBlockFormat_isValid")
	symaddr(&addr_QTextBlockFormat_setAlignment, "QTextBlockFormat_setAlignment")
	symaddr(&addr_QTextBlockFormat_alignment, "QTextBlockFormat_alignment")
	reg(&pg_QTextBlockFormat_setTopMargin, "QTextBlockFormat_setTopMargin")
	reg(&pg_QTextBlockFormat_topMargin, "QTextBlockFormat_topMargin")
	reg(&pg_QTextBlockFormat_setBottomMargin, "QTextBlockFormat_setBottomMargin")
	reg(&pg_QTextBlockFormat_bottomMargin, "QTextBlockFormat_bottomMargin")
	reg(&pg_QTextBlockFormat_setLeftMargin, "QTextBlockFormat_setLeftMargin")
	reg(&pg_QTextBlockFormat_leftMargin, "QTextBlockFormat_leftMargin")
	reg(&pg_QTextBlockFormat_setRightMargin, "QTextBlockFormat_setRightMargin")
	reg(&pg_QTextBlockFormat_rightMargin, "QTextBlockFormat_rightMargin")
	reg(&pg_QTextBlockFormat_setTextIndent, "QTextBlockFormat_setTextIndent")
	reg(&pg_QTextBlockFormat_textIndent, "QTextBlockFormat_textIndent")
	symaddr(&addr_QTextBlockFormat_setIndent, "QTextBlockFormat_setIndent")
	symaddr(&addr_QTextBlockFormat_indent, "QTextBlockFormat_indent")
	symaddr(&addr_QTextBlockFormat_setHeadingLevel, "QTextBlockFormat_setHeadingLevel")
	symaddr(&addr_QTextBlockFormat_headingLevel, "QTextBlockFormat_headingLevel")
	reg(&pg_QTextBlockFormat_setLineHeight, "QTextBlockFormat_setLineHeight")
	reg(&pg_QTextBlockFormat_lineHeight, "QTextBlockFormat_lineHeight")
	reg(&pg_QTextBlockFormat_lineHeight2, "QTextBlockFormat_lineHeight2")
	symaddr(&addr_QTextBlockFormat_lineHeightType, "QTextBlockFormat_lineHeightType")
	symaddr(&addr_QTextBlockFormat_setNonBreakableLines, "QTextBlockFormat_setNonBreakableLines")
	symaddr(&addr_QTextBlockFormat_nonBreakableLines, "QTextBlockFormat_nonBreakableLines")
	symaddr(&addr_QTextBlockFormat_setPageBreakPolicy, "QTextBlockFormat_setPageBreakPolicy")
	symaddr(&addr_QTextBlockFormat_pageBreakPolicy, "QTextBlockFormat_pageBreakPolicy")
	symaddr(&addr_QTextBlockFormat_setTabPositions, "QTextBlockFormat_setTabPositions")
	symaddr(&addr_QTextBlockFormat_tabPositions, "QTextBlockFormat_tabPositions")
	symaddr(&addr_QTextBlockFormat_setMarker, "QTextBlockFormat_setMarker")
	symaddr(&addr_QTextBlockFormat_marker, "QTextBlockFormat_marker")
	symaddr(&addr_QTextBlockFormat_OperatorAssign, "QTextBlockFormat_OperatorAssign")
	reg(&pg_QTextBlockFormat_delete, "QTextBlockFormat_delete")
}

type QTextBlockFormat struct {
	h uintptr
	*QTextFormat
}

func bindQTextBlockFormat(h uintptr) *QTextBlockFormat {
	if h == 0 {
		return nil
	}
	return &QTextBlockFormat{h: h, QTextFormat: bindQTextFormat(pg_QTextBlockFormat_base(h))}
}

func newownedQTextBlockFormat(h uintptr) *QTextBlockFormat {
	o := bindQTextBlockFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextBlockFormat).Delete)
	}
	return o
}

func (this *QTextBlockFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextBlockFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextBlockFormat(h unsafe.Pointer) *QTextBlockFormat { return bindQTextBlockFormat(uintptr(h)) }
func UnsafeNewQTextBlockFormatOwned(h unsafe.Pointer) *QTextBlockFormat { return newownedQTextBlockFormat(uintptr(h)) }

func (this *QTextBlockFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextFormat != nil {
		this.QTextFormat.UnsafeRelease()
	}
}

func NewQTextBlockFormat() *QTextBlockFormat {
	__r, _, _ := purego.SyscallN(addr_QTextBlockFormat_new)
	__o := newownedQTextBlockFormat(__r)
	return __o
}

func NewQTextBlockFormat2(arg1 *QTextBlockFormat) *QTextBlockFormat {
	__r, _, _ := purego.SyscallN(addr_QTextBlockFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextBlockFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextBlockFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBlockFormat) SetAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QTextBlockFormat_setAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QTextBlockFormat) SetTopMargin(margin float64) {
	pg_QTextBlockFormat_setTopMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) TopMargin() float64 {
	__r := pg_QTextBlockFormat_topMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) SetBottomMargin(margin float64) {
	pg_QTextBlockFormat_setBottomMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) BottomMargin() float64 {
	__r := pg_QTextBlockFormat_bottomMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) SetLeftMargin(margin float64) {
	pg_QTextBlockFormat_setLeftMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) LeftMargin() float64 {
	__r := pg_QTextBlockFormat_leftMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) SetRightMargin(margin float64) {
	pg_QTextBlockFormat_setRightMargin(this.h, margin)
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) RightMargin() float64 {
	__r := pg_QTextBlockFormat_rightMargin(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) SetTextIndent(aindent float64) {
	pg_QTextBlockFormat_setTextIndent(this.h, aindent)
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) TextIndent() float64 {
	__r := pg_QTextBlockFormat_textIndent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) SetIndent(indent int) {
	purego.SyscallN(addr_QTextBlockFormat_setIndent, this.h, uintptr(indent))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) Indent() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_indent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlockFormat) SetHeadingLevel(alevel int) {
	purego.SyscallN(addr_QTextBlockFormat_setHeadingLevel, this.h, uintptr(alevel))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) HeadingLevel() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_headingLevel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlockFormat) SetLineHeight(height float64, heightType int) {
	pg_QTextBlockFormat_setLineHeight(this.h, height, int32(heightType))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) LineHeight(scriptLineHeight float64, scaling float64) float64 {
	__r := pg_QTextBlockFormat_lineHeight(this.h, scriptLineHeight, scaling)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) LineHeight2() float64 {
	__r := pg_QTextBlockFormat_lineHeight2(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBlockFormat) LineHeightType() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_lineHeightType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBlockFormat) SetNonBreakableLines(b bool) {
	purego.SyscallN(addr_QTextBlockFormat_setNonBreakableLines, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) NonBreakableLines() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_nonBreakableLines, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBlockFormat) SetPageBreakPolicy(flags QTextFormat__PageBreakFlag) {
	purego.SyscallN(addr_QTextBlockFormat_setPageBreakPolicy, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) PageBreakPolicy() QTextFormat__PageBreakFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_pageBreakPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextFormat__PageBreakFlag(__r)
}

func (this *QTextBlockFormat) SetTabPositions(tabs []*QTextOption__Tab) {
	tabs_arr := make([]uintptr, len(tabs))
	for __i, __e := range tabs {
		tabs_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var tabs_p uintptr
	if len(tabs_arr) > 0 {
		tabs_p = uintptr(unsafe.Pointer(&tabs_arr[0]))
	}
	purego.SyscallN(addr_QTextBlockFormat_setTabPositions, this.h, tabs_p, uintptr(len(tabs)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tabs)
	runtime.KeepAlive(tabs_arr)
}

func (this *QTextBlockFormat) TabPositions() []*QTextOption__Tab {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBlockFormat_tabPositions, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextOption__Tab, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextOption__Tab(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextBlockFormat) SetMarker(marker QTextBlockFormat__MarkerType) {
	purego.SyscallN(addr_QTextBlockFormat_setMarker, this.h, uintptr(marker))
	runtime.KeepAlive(this)
}

func (this *QTextBlockFormat) Marker() QTextBlockFormat__MarkerType {
	__r0, _, _ := purego.SyscallN(addr_QTextBlockFormat_marker, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextBlockFormat__MarkerType(__r)
}

func (this *QTextBlockFormat) OperatorAssign(arg1 *QTextBlockFormat) *QTextBlockFormat {
	__r, _, _ := purego.SyscallN(addr_QTextBlockFormat_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextBlockFormat(__r)
}

func (this *QTextBlockFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextBlockFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
