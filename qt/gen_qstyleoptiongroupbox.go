package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionGroupBox_chain uintptr
	addr_QStyleOptionGroupBox_new uintptr
	addr_QStyleOptionGroupBox_new2 uintptr
	addr_QStyleOptionGroupBox_OperatorAssign uintptr
	addr_QStyleOptionGroupBox_features_get uintptr
	addr_QStyleOptionGroupBox_features_set uintptr
	addr_QStyleOptionGroupBox_text_get uintptr
	addr_QStyleOptionGroupBox_text_set uintptr
	addr_QStyleOptionGroupBox_textAlignment_get uintptr
	addr_QStyleOptionGroupBox_textAlignment_set uintptr
	addr_QStyleOptionGroupBox_textColor_get uintptr
	addr_QStyleOptionGroupBox_textColor_set uintptr
	addr_QStyleOptionGroupBox_lineWidth_get uintptr
	addr_QStyleOptionGroupBox_lineWidth_set uintptr
	addr_QStyleOptionGroupBox_midLineWidth_get uintptr
	addr_QStyleOptionGroupBox_midLineWidth_set uintptr
	pg_QStyleOptionGroupBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionGroupBox_chain, "QStyleOptionGroupBox_chain")
	symaddr(&addr_QStyleOptionGroupBox_new, "QStyleOptionGroupBox_new")
	symaddr(&addr_QStyleOptionGroupBox_new2, "QStyleOptionGroupBox_new2")
	symaddr(&addr_QStyleOptionGroupBox_OperatorAssign, "QStyleOptionGroupBox_OperatorAssign")
	symaddr(&addr_QStyleOptionGroupBox_features_get, "QStyleOptionGroupBox_features_get")
	symaddr(&addr_QStyleOptionGroupBox_features_set, "QStyleOptionGroupBox_features_set")
	symaddr(&addr_QStyleOptionGroupBox_text_get, "QStyleOptionGroupBox_text_get")
	symaddr(&addr_QStyleOptionGroupBox_text_set, "QStyleOptionGroupBox_text_set")
	symaddr(&addr_QStyleOptionGroupBox_textAlignment_get, "QStyleOptionGroupBox_textAlignment_get")
	symaddr(&addr_QStyleOptionGroupBox_textAlignment_set, "QStyleOptionGroupBox_textAlignment_set")
	symaddr(&addr_QStyleOptionGroupBox_textColor_get, "QStyleOptionGroupBox_textColor_get")
	symaddr(&addr_QStyleOptionGroupBox_textColor_set, "QStyleOptionGroupBox_textColor_set")
	symaddr(&addr_QStyleOptionGroupBox_lineWidth_get, "QStyleOptionGroupBox_lineWidth_get")
	symaddr(&addr_QStyleOptionGroupBox_lineWidth_set, "QStyleOptionGroupBox_lineWidth_set")
	symaddr(&addr_QStyleOptionGroupBox_midLineWidth_get, "QStyleOptionGroupBox_midLineWidth_get")
	symaddr(&addr_QStyleOptionGroupBox_midLineWidth_set, "QStyleOptionGroupBox_midLineWidth_set")
	reg(&pg_QStyleOptionGroupBox_delete, "QStyleOptionGroupBox_delete")
}

type QStyleOptionGroupBox struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionGroupBox(h uintptr) *QStyleOptionGroupBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionGroupBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionGroupBox{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionGroupBox(h uintptr) *QStyleOptionGroupBox {
	o := bindQStyleOptionGroupBox(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionGroupBox).Delete)
	}
	return o
}

func (this *QStyleOptionGroupBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionGroupBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionGroupBox(h unsafe.Pointer) *QStyleOptionGroupBox { return bindQStyleOptionGroupBox(uintptr(h)) }
func UnsafeNewQStyleOptionGroupBoxOwned(h unsafe.Pointer) *QStyleOptionGroupBox { return newownedQStyleOptionGroupBox(uintptr(h)) }

func (this *QStyleOptionGroupBox) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_new)
	__o := newownedQStyleOptionGroupBox(__r)
	return __o
}

func NewQStyleOptionGroupBox2(other *QStyleOptionGroupBox) *QStyleOptionGroupBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionGroupBox(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionGroupBox) OperatorAssign(arg1 *QStyleOptionGroupBox) *QStyleOptionGroupBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionGroupBox(__r)
}

func (this *QStyleOptionGroupBox) Features() QStyleOptionFrame__FrameFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionFrame__FrameFeature(__r)
}

func (this *QStyleOptionGroupBox) SetFeatures(features QStyleOptionFrame__FrameFeature) {
	purego.SyscallN(addr_QStyleOptionGroupBox_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionGroupBox) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionGroupBox) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionGroupBox_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionGroupBox) TextAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_textAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionGroupBox) SetTextAlignment(textAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionGroupBox_textAlignment_set, this.h, uintptr(textAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionGroupBox) TextColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_textColor_get, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QStyleOptionGroupBox) SetTextColor(textColor *QColor) {
	purego.SyscallN(addr_QStyleOptionGroupBox_textColor_set, this.h, uintptr(textColor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(textColor)
}

func (this *QStyleOptionGroupBox) LineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_lineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionGroupBox) SetLineWidth(lineWidth int) {
	purego.SyscallN(addr_QStyleOptionGroupBox_lineWidth_set, this.h, uintptr(lineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionGroupBox) MidLineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionGroupBox_midLineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionGroupBox) SetMidLineWidth(midLineWidth int) {
	purego.SyscallN(addr_QStyleOptionGroupBox_midLineWidth_set, this.h, uintptr(midLineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionGroupBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionGroupBox_delete(this.h)
	}
	this.UnsafeRelease()
}
