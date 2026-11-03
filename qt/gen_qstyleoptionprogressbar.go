package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionProgressBar_base func(uintptr) uintptr
	addr_QStyleOptionProgressBar_new uintptr
	addr_QStyleOptionProgressBar_new2 uintptr
	addr_QStyleOptionProgressBar_OperatorAssign uintptr
	addr_QStyleOptionProgressBar_minimum_get uintptr
	addr_QStyleOptionProgressBar_minimum_set uintptr
	addr_QStyleOptionProgressBar_maximum_get uintptr
	addr_QStyleOptionProgressBar_maximum_set uintptr
	addr_QStyleOptionProgressBar_progress_get uintptr
	addr_QStyleOptionProgressBar_progress_set uintptr
	addr_QStyleOptionProgressBar_text_get uintptr
	addr_QStyleOptionProgressBar_text_set uintptr
	addr_QStyleOptionProgressBar_textAlignment_get uintptr
	addr_QStyleOptionProgressBar_textAlignment_set uintptr
	addr_QStyleOptionProgressBar_textVisible_get uintptr
	addr_QStyleOptionProgressBar_textVisible_set uintptr
	addr_QStyleOptionProgressBar_invertedAppearance_get uintptr
	addr_QStyleOptionProgressBar_invertedAppearance_set uintptr
	addr_QStyleOptionProgressBar_bottomToTop_get uintptr
	addr_QStyleOptionProgressBar_bottomToTop_set uintptr
	pg_QStyleOptionProgressBar_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionProgressBar_base, "QStyleOptionProgressBar_base")
	symaddr(&addr_QStyleOptionProgressBar_new, "QStyleOptionProgressBar_new")
	symaddr(&addr_QStyleOptionProgressBar_new2, "QStyleOptionProgressBar_new2")
	symaddr(&addr_QStyleOptionProgressBar_OperatorAssign, "QStyleOptionProgressBar_OperatorAssign")
	symaddr(&addr_QStyleOptionProgressBar_minimum_get, "QStyleOptionProgressBar_minimum_get")
	symaddr(&addr_QStyleOptionProgressBar_minimum_set, "QStyleOptionProgressBar_minimum_set")
	symaddr(&addr_QStyleOptionProgressBar_maximum_get, "QStyleOptionProgressBar_maximum_get")
	symaddr(&addr_QStyleOptionProgressBar_maximum_set, "QStyleOptionProgressBar_maximum_set")
	symaddr(&addr_QStyleOptionProgressBar_progress_get, "QStyleOptionProgressBar_progress_get")
	symaddr(&addr_QStyleOptionProgressBar_progress_set, "QStyleOptionProgressBar_progress_set")
	symaddr(&addr_QStyleOptionProgressBar_text_get, "QStyleOptionProgressBar_text_get")
	symaddr(&addr_QStyleOptionProgressBar_text_set, "QStyleOptionProgressBar_text_set")
	symaddr(&addr_QStyleOptionProgressBar_textAlignment_get, "QStyleOptionProgressBar_textAlignment_get")
	symaddr(&addr_QStyleOptionProgressBar_textAlignment_set, "QStyleOptionProgressBar_textAlignment_set")
	symaddr(&addr_QStyleOptionProgressBar_textVisible_get, "QStyleOptionProgressBar_textVisible_get")
	symaddr(&addr_QStyleOptionProgressBar_textVisible_set, "QStyleOptionProgressBar_textVisible_set")
	symaddr(&addr_QStyleOptionProgressBar_invertedAppearance_get, "QStyleOptionProgressBar_invertedAppearance_get")
	symaddr(&addr_QStyleOptionProgressBar_invertedAppearance_set, "QStyleOptionProgressBar_invertedAppearance_set")
	symaddr(&addr_QStyleOptionProgressBar_bottomToTop_get, "QStyleOptionProgressBar_bottomToTop_get")
	symaddr(&addr_QStyleOptionProgressBar_bottomToTop_set, "QStyleOptionProgressBar_bottomToTop_set")
	reg(&pg_QStyleOptionProgressBar_delete, "QStyleOptionProgressBar_delete")
}

type QStyleOptionProgressBar struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionProgressBar(h uintptr) *QStyleOptionProgressBar {
	if h == 0 {
		return nil
	}
	return &QStyleOptionProgressBar{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionProgressBar_base(h))}
}

func newownedQStyleOptionProgressBar(h uintptr) *QStyleOptionProgressBar {
	o := bindQStyleOptionProgressBar(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionProgressBar).Delete)
	}
	return o
}

func (this *QStyleOptionProgressBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionProgressBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionProgressBar(h unsafe.Pointer) *QStyleOptionProgressBar { return bindQStyleOptionProgressBar(uintptr(h)) }
func UnsafeNewQStyleOptionProgressBarOwned(h unsafe.Pointer) *QStyleOptionProgressBar { return newownedQStyleOptionProgressBar(uintptr(h)) }

func (this *QStyleOptionProgressBar) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionProgressBar() *QStyleOptionProgressBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_new)
	__o := newownedQStyleOptionProgressBar(__r)
	return __o
}

func NewQStyleOptionProgressBar2(other *QStyleOptionProgressBar) *QStyleOptionProgressBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionProgressBar(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionProgressBar) OperatorAssign(arg1 *QStyleOptionProgressBar) *QStyleOptionProgressBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionProgressBar(__r)
}

func (this *QStyleOptionProgressBar) Minimum() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_minimum_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionProgressBar) SetMinimum(minimum int) {
	purego.SyscallN(addr_QStyleOptionProgressBar_minimum_set, this.h, uintptr(minimum))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) Maximum() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_maximum_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionProgressBar) SetMaximum(maximum int) {
	purego.SyscallN(addr_QStyleOptionProgressBar_maximum_set, this.h, uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) Progress() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_progress_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionProgressBar) SetProgress(progress int) {
	purego.SyscallN(addr_QStyleOptionProgressBar_progress_set, this.h, uintptr(progress))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionProgressBar) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionProgressBar_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionProgressBar) TextAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_textAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionProgressBar) SetTextAlignment(textAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionProgressBar_textAlignment_set, this.h, uintptr(textAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) TextVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_textVisible_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionProgressBar) SetTextVisible(textVisible bool) {
	purego.SyscallN(addr_QStyleOptionProgressBar_textVisible_set, this.h, uintptr(b2i(textVisible)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) InvertedAppearance() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_invertedAppearance_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionProgressBar) SetInvertedAppearance(invertedAppearance bool) {
	purego.SyscallN(addr_QStyleOptionProgressBar_invertedAppearance_set, this.h, uintptr(b2i(invertedAppearance)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) BottomToTop() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionProgressBar_bottomToTop_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionProgressBar) SetBottomToTop(bottomToTop bool) {
	purego.SyscallN(addr_QStyleOptionProgressBar_bottomToTop_set, this.h, uintptr(b2i(bottomToTop)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionProgressBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionProgressBar_delete(this.h)
	}
	this.UnsafeRelease()
}
