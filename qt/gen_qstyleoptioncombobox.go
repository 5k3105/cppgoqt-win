package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionComboBox_chain uintptr
	addr_QStyleOptionComboBox_new uintptr
	addr_QStyleOptionComboBox_new2 uintptr
	addr_QStyleOptionComboBox_OperatorAssign uintptr
	addr_QStyleOptionComboBox_editable_get uintptr
	addr_QStyleOptionComboBox_editable_set uintptr
	addr_QStyleOptionComboBox_popupRect_get uintptr
	addr_QStyleOptionComboBox_popupRect_set uintptr
	addr_QStyleOptionComboBox_frame_get uintptr
	addr_QStyleOptionComboBox_frame_set uintptr
	addr_QStyleOptionComboBox_currentText_get uintptr
	addr_QStyleOptionComboBox_currentText_set uintptr
	addr_QStyleOptionComboBox_currentIcon_get uintptr
	addr_QStyleOptionComboBox_currentIcon_set uintptr
	addr_QStyleOptionComboBox_iconSize_get uintptr
	addr_QStyleOptionComboBox_iconSize_set uintptr
	addr_QStyleOptionComboBox_textAlignment_get uintptr
	addr_QStyleOptionComboBox_textAlignment_set uintptr
	pg_QStyleOptionComboBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionComboBox_chain, "QStyleOptionComboBox_chain")
	symaddr(&addr_QStyleOptionComboBox_new, "QStyleOptionComboBox_new")
	symaddr(&addr_QStyleOptionComboBox_new2, "QStyleOptionComboBox_new2")
	symaddr(&addr_QStyleOptionComboBox_OperatorAssign, "QStyleOptionComboBox_OperatorAssign")
	symaddr(&addr_QStyleOptionComboBox_editable_get, "QStyleOptionComboBox_editable_get")
	symaddr(&addr_QStyleOptionComboBox_editable_set, "QStyleOptionComboBox_editable_set")
	symaddr(&addr_QStyleOptionComboBox_popupRect_get, "QStyleOptionComboBox_popupRect_get")
	symaddr(&addr_QStyleOptionComboBox_popupRect_set, "QStyleOptionComboBox_popupRect_set")
	symaddr(&addr_QStyleOptionComboBox_frame_get, "QStyleOptionComboBox_frame_get")
	symaddr(&addr_QStyleOptionComboBox_frame_set, "QStyleOptionComboBox_frame_set")
	symaddr(&addr_QStyleOptionComboBox_currentText_get, "QStyleOptionComboBox_currentText_get")
	symaddr(&addr_QStyleOptionComboBox_currentText_set, "QStyleOptionComboBox_currentText_set")
	symaddr(&addr_QStyleOptionComboBox_currentIcon_get, "QStyleOptionComboBox_currentIcon_get")
	symaddr(&addr_QStyleOptionComboBox_currentIcon_set, "QStyleOptionComboBox_currentIcon_set")
	symaddr(&addr_QStyleOptionComboBox_iconSize_get, "QStyleOptionComboBox_iconSize_get")
	symaddr(&addr_QStyleOptionComboBox_iconSize_set, "QStyleOptionComboBox_iconSize_set")
	symaddr(&addr_QStyleOptionComboBox_textAlignment_get, "QStyleOptionComboBox_textAlignment_get")
	symaddr(&addr_QStyleOptionComboBox_textAlignment_set, "QStyleOptionComboBox_textAlignment_set")
	reg(&pg_QStyleOptionComboBox_delete, "QStyleOptionComboBox_delete")
}

type QStyleOptionComboBox struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionComboBox(h uintptr) *QStyleOptionComboBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionComboBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionComboBox{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionComboBox(h uintptr) *QStyleOptionComboBox {
	o := bindQStyleOptionComboBox(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionComboBox).Delete)
	}
	return o
}

func (this *QStyleOptionComboBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionComboBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionComboBox(h unsafe.Pointer) *QStyleOptionComboBox { return bindQStyleOptionComboBox(uintptr(h)) }
func UnsafeNewQStyleOptionComboBoxOwned(h unsafe.Pointer) *QStyleOptionComboBox { return newownedQStyleOptionComboBox(uintptr(h)) }

func (this *QStyleOptionComboBox) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionComboBox() *QStyleOptionComboBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_new)
	__o := newownedQStyleOptionComboBox(__r)
	return __o
}

func NewQStyleOptionComboBox2(other *QStyleOptionComboBox) *QStyleOptionComboBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionComboBox(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionComboBox) OperatorAssign(arg1 *QStyleOptionComboBox) *QStyleOptionComboBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionComboBox(__r)
}

func (this *QStyleOptionComboBox) Editable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_editable_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionComboBox) SetEditable(editable bool) {
	purego.SyscallN(addr_QStyleOptionComboBox_editable_set, this.h, uintptr(b2i(editable)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionComboBox) PopupRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_popupRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOptionComboBox) SetPopupRect(popupRect *QRect) {
	purego.SyscallN(addr_QStyleOptionComboBox_popupRect_set, this.h, uintptr(popupRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(popupRect)
}

func (this *QStyleOptionComboBox) Frame() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_frame_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionComboBox) SetFrame(frame bool) {
	purego.SyscallN(addr_QStyleOptionComboBox_frame_set, this.h, uintptr(b2i(frame)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionComboBox) CurrentText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_currentText_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionComboBox) SetCurrentText(currentText string) {
	var currentText_p uintptr
	if len(currentText) > 0 {
		currentText_p = uintptr(unsafe.Pointer(unsafe.StringData(currentText)))
	}
	purego.SyscallN(addr_QStyleOptionComboBox_currentText_set, this.h, currentText_p, uintptr(len(currentText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(currentText)
}

func (this *QStyleOptionComboBox) CurrentIcon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_currentIcon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionComboBox) SetCurrentIcon(currentIcon *QIcon) {
	purego.SyscallN(addr_QStyleOptionComboBox_currentIcon_set, this.h, uintptr(currentIcon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(currentIcon)
}

func (this *QStyleOptionComboBox) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_iconSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionComboBox) SetIconSize(iconSize *QSize) {
	purego.SyscallN(addr_QStyleOptionComboBox_iconSize_set, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QStyleOptionComboBox) TextAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionComboBox_textAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionComboBox) SetTextAlignment(textAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionComboBox_textAlignment_set, this.h, uintptr(textAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionComboBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionComboBox_delete(this.h)
	}
	this.UnsafeRelease()
}
