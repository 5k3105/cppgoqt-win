package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionToolButton_chain uintptr
	addr_QStyleOptionToolButton_new uintptr
	addr_QStyleOptionToolButton_new2 uintptr
	addr_QStyleOptionToolButton_OperatorAssign uintptr
	addr_QStyleOptionToolButton_features_get uintptr
	addr_QStyleOptionToolButton_features_set uintptr
	addr_QStyleOptionToolButton_icon_get uintptr
	addr_QStyleOptionToolButton_icon_set uintptr
	addr_QStyleOptionToolButton_iconSize_get uintptr
	addr_QStyleOptionToolButton_iconSize_set uintptr
	addr_QStyleOptionToolButton_text_get uintptr
	addr_QStyleOptionToolButton_text_set uintptr
	addr_QStyleOptionToolButton_arrowType_get uintptr
	addr_QStyleOptionToolButton_arrowType_set uintptr
	addr_QStyleOptionToolButton_toolButtonStyle_get uintptr
	addr_QStyleOptionToolButton_toolButtonStyle_set uintptr
	addr_QStyleOptionToolButton_pos_get uintptr
	addr_QStyleOptionToolButton_pos_set uintptr
	addr_QStyleOptionToolButton_font_get uintptr
	addr_QStyleOptionToolButton_font_set uintptr
	pg_QStyleOptionToolButton_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionToolButton_chain, "QStyleOptionToolButton_chain")
	symaddr(&addr_QStyleOptionToolButton_new, "QStyleOptionToolButton_new")
	symaddr(&addr_QStyleOptionToolButton_new2, "QStyleOptionToolButton_new2")
	symaddr(&addr_QStyleOptionToolButton_OperatorAssign, "QStyleOptionToolButton_OperatorAssign")
	symaddr(&addr_QStyleOptionToolButton_features_get, "QStyleOptionToolButton_features_get")
	symaddr(&addr_QStyleOptionToolButton_features_set, "QStyleOptionToolButton_features_set")
	symaddr(&addr_QStyleOptionToolButton_icon_get, "QStyleOptionToolButton_icon_get")
	symaddr(&addr_QStyleOptionToolButton_icon_set, "QStyleOptionToolButton_icon_set")
	symaddr(&addr_QStyleOptionToolButton_iconSize_get, "QStyleOptionToolButton_iconSize_get")
	symaddr(&addr_QStyleOptionToolButton_iconSize_set, "QStyleOptionToolButton_iconSize_set")
	symaddr(&addr_QStyleOptionToolButton_text_get, "QStyleOptionToolButton_text_get")
	symaddr(&addr_QStyleOptionToolButton_text_set, "QStyleOptionToolButton_text_set")
	symaddr(&addr_QStyleOptionToolButton_arrowType_get, "QStyleOptionToolButton_arrowType_get")
	symaddr(&addr_QStyleOptionToolButton_arrowType_set, "QStyleOptionToolButton_arrowType_set")
	symaddr(&addr_QStyleOptionToolButton_toolButtonStyle_get, "QStyleOptionToolButton_toolButtonStyle_get")
	symaddr(&addr_QStyleOptionToolButton_toolButtonStyle_set, "QStyleOptionToolButton_toolButtonStyle_set")
	symaddr(&addr_QStyleOptionToolButton_pos_get, "QStyleOptionToolButton_pos_get")
	symaddr(&addr_QStyleOptionToolButton_pos_set, "QStyleOptionToolButton_pos_set")
	symaddr(&addr_QStyleOptionToolButton_font_get, "QStyleOptionToolButton_font_get")
	symaddr(&addr_QStyleOptionToolButton_font_set, "QStyleOptionToolButton_font_set")
	reg(&pg_QStyleOptionToolButton_delete, "QStyleOptionToolButton_delete")
}

type QStyleOptionToolButton struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionToolButton(h uintptr) *QStyleOptionToolButton {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionToolButton_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionToolButton{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionToolButton(h uintptr) *QStyleOptionToolButton {
	o := bindQStyleOptionToolButton(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionToolButton).Delete)
	}
	return o
}

func (this *QStyleOptionToolButton) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionToolButton) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionToolButton(h unsafe.Pointer) *QStyleOptionToolButton { return bindQStyleOptionToolButton(uintptr(h)) }
func UnsafeNewQStyleOptionToolButtonOwned(h unsafe.Pointer) *QStyleOptionToolButton { return newownedQStyleOptionToolButton(uintptr(h)) }

func (this *QStyleOptionToolButton) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionToolButton() *QStyleOptionToolButton {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_new)
	__o := newownedQStyleOptionToolButton(__r)
	return __o
}

func NewQStyleOptionToolButton2(other *QStyleOptionToolButton) *QStyleOptionToolButton {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionToolButton(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionToolButton) OperatorAssign(arg1 *QStyleOptionToolButton) *QStyleOptionToolButton {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionToolButton(__r)
}

func (this *QStyleOptionToolButton) Features() QStyleOptionToolButton__ToolButtonFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionToolButton__ToolButtonFeature(__r)
}

func (this *QStyleOptionToolButton) SetFeatures(features QStyleOptionToolButton__ToolButtonFeature) {
	purego.SyscallN(addr_QStyleOptionToolButton_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolButton) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionToolButton) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionToolButton_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionToolButton) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_iconSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionToolButton) SetIconSize(iconSize *QSize) {
	purego.SyscallN(addr_QStyleOptionToolButton_iconSize_set, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QStyleOptionToolButton) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionToolButton) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionToolButton_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionToolButton) ArrowType() ArrowType {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_arrowType_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ArrowType(__r)
}

func (this *QStyleOptionToolButton) SetArrowType(arrowType ArrowType) {
	purego.SyscallN(addr_QStyleOptionToolButton_arrowType_set, this.h, uintptr(arrowType))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolButton) ToolButtonStyle() ToolButtonStyle {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_toolButtonStyle_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ToolButtonStyle(__r)
}

func (this *QStyleOptionToolButton) SetToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	purego.SyscallN(addr_QStyleOptionToolButton_toolButtonStyle_set, this.h, uintptr(toolButtonStyle))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolButton) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_pos_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QStyleOptionToolButton) SetPos(pos *QPoint) {
	purego.SyscallN(addr_QStyleOptionToolButton_pos_set, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QStyleOptionToolButton) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolButton_font_get, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QStyleOptionToolButton) SetFont(font *QFont) {
	purego.SyscallN(addr_QStyleOptionToolButton_font_set, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QStyleOptionToolButton) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionToolButton_delete(this.h)
	}
	this.UnsafeRelease()
}
