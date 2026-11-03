package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStaticText_new uintptr
	addr_QStaticText_new2 uintptr
	addr_QStaticText_new3 uintptr
	addr_QStaticText_OperatorAssign uintptr
	addr_QStaticText_swap uintptr
	addr_QStaticText_setText uintptr
	addr_QStaticText_text uintptr
	addr_QStaticText_setTextFormat uintptr
	addr_QStaticText_textFormat uintptr
	pg_QStaticText_setTextWidth func(uintptr, float64)
	pg_QStaticText_textWidth func(uintptr) float64
	addr_QStaticText_setTextOption uintptr
	addr_QStaticText_textOption uintptr
	addr_QStaticText_size uintptr
	addr_QStaticText_prepare uintptr
	addr_QStaticText_setPerformanceHint uintptr
	addr_QStaticText_performanceHint uintptr
	addr_QStaticText_OperatorEqual uintptr
	addr_QStaticText_OperatorNotEqual uintptr
	addr_QStaticText_prepare2 uintptr
	addr_QStaticText_prepare3 uintptr
	pg_QStaticText_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStaticText_new, "QStaticText_new")
	symaddr(&addr_QStaticText_new2, "QStaticText_new2")
	symaddr(&addr_QStaticText_new3, "QStaticText_new3")
	symaddr(&addr_QStaticText_OperatorAssign, "QStaticText_OperatorAssign")
	symaddr(&addr_QStaticText_swap, "QStaticText_swap")
	symaddr(&addr_QStaticText_setText, "QStaticText_setText")
	symaddr(&addr_QStaticText_text, "QStaticText_text")
	symaddr(&addr_QStaticText_setTextFormat, "QStaticText_setTextFormat")
	symaddr(&addr_QStaticText_textFormat, "QStaticText_textFormat")
	reg(&pg_QStaticText_setTextWidth, "QStaticText_setTextWidth")
	reg(&pg_QStaticText_textWidth, "QStaticText_textWidth")
	symaddr(&addr_QStaticText_setTextOption, "QStaticText_setTextOption")
	symaddr(&addr_QStaticText_textOption, "QStaticText_textOption")
	symaddr(&addr_QStaticText_size, "QStaticText_size")
	symaddr(&addr_QStaticText_prepare, "QStaticText_prepare")
	symaddr(&addr_QStaticText_setPerformanceHint, "QStaticText_setPerformanceHint")
	symaddr(&addr_QStaticText_performanceHint, "QStaticText_performanceHint")
	symaddr(&addr_QStaticText_OperatorEqual, "QStaticText_OperatorEqual")
	symaddr(&addr_QStaticText_OperatorNotEqual, "QStaticText_OperatorNotEqual")
	symaddr(&addr_QStaticText_prepare2, "QStaticText_prepare2")
	symaddr(&addr_QStaticText_prepare3, "QStaticText_prepare3")
	reg(&pg_QStaticText_delete, "QStaticText_delete")
}

type QStaticText struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStaticText) UnsafeGoID() int64     { return this.vid }
func (this *QStaticText) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStaticText(h uintptr) *QStaticText {
	if h == 0 {
		return nil
	}
	return &QStaticText{h: h}
}

func newownedQStaticText(h uintptr) *QStaticText {
	o := bindQStaticText(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStaticText).Delete)
	}
	return o
}

func (this *QStaticText) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStaticText) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStaticText(h unsafe.Pointer) *QStaticText { return bindQStaticText(uintptr(h)) }
func UnsafeNewQStaticTextOwned(h unsafe.Pointer) *QStaticText { return newownedQStaticText(uintptr(h)) }

func (this *QStaticText) UnsafeRelease() { this.h = 0 }

func NewQStaticText() *QStaticText {
	__r, _, _ := purego.SyscallN(addr_QStaticText_new)
	__o := newownedQStaticText(__r)
	return __o
}

func NewQStaticText2(text string) *QStaticText {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QStaticText_new2, text_p, uintptr(len(text)))
	__o := newownedQStaticText(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQStaticText3(other *QStaticText) *QStaticText {
	__r, _, _ := purego.SyscallN(addr_QStaticText_new3, uintptr(other.UnsafePointer()))
	__o := newownedQStaticText(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStaticText) OperatorAssign(arg1 *QStaticText) *QStaticText {
	__r, _, _ := purego.SyscallN(addr_QStaticText_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStaticText(__r)
}

func (this *QStaticText) Swap(other *QStaticText) {
	purego.SyscallN(addr_QStaticText_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QStaticText) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStaticText_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStaticText) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStaticText_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStaticText) SetTextFormat(textFormat TextFormat) {
	purego.SyscallN(addr_QStaticText_setTextFormat, this.h, uintptr(textFormat))
	runtime.KeepAlive(this)
}

func (this *QStaticText) TextFormat() TextFormat {
	__r0, _, _ := purego.SyscallN(addr_QStaticText_textFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextFormat(__r)
}

func (this *QStaticText) SetTextWidth(textWidth float64) {
	pg_QStaticText_setTextWidth(this.h, textWidth)
	runtime.KeepAlive(this)
}

func (this *QStaticText) TextWidth() float64 {
	__r := pg_QStaticText_textWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStaticText) SetTextOption(textOption *QTextOption) {
	purego.SyscallN(addr_QStaticText_setTextOption, this.h, uintptr(textOption.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(textOption)
}

func (this *QStaticText) TextOption() *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QStaticText_textOption, this.h)
	runtime.KeepAlive(this)
	return newownedQTextOption(__r)
}

func (this *QStaticText) Size() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QStaticText_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QStaticText) Prepare() {
	purego.SyscallN(addr_QStaticText_prepare, this.h)
	runtime.KeepAlive(this)
}

func (this *QStaticText) SetPerformanceHint(performanceHint QStaticText__PerformanceHint) {
	purego.SyscallN(addr_QStaticText_setPerformanceHint, this.h, uintptr(performanceHint))
	runtime.KeepAlive(this)
}

func (this *QStaticText) PerformanceHint() QStaticText__PerformanceHint {
	__r0, _, _ := purego.SyscallN(addr_QStaticText_performanceHint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStaticText__PerformanceHint(__r)
}

func (this *QStaticText) OperatorEqual(arg1 *QStaticText) bool {
	__r0, _, _ := purego.SyscallN(addr_QStaticText_OperatorEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QStaticText) OperatorNotEqual(arg1 *QStaticText) bool {
	__r0, _, _ := purego.SyscallN(addr_QStaticText_OperatorNotEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QStaticText) PrepareWithMatrix(matrix *QTransform) {
	purego.SyscallN(addr_QStaticText_prepare2, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QStaticText) Prepare2(matrix *QTransform, font *QFont) {
	purego.SyscallN(addr_QStaticText_prepare3, this.h, uintptr(matrix.UnsafePointer()), uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	runtime.KeepAlive(font)
}

func (this *QStaticText) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStaticText_delete(this.h)
	}
	this.UnsafeRelease()
}
