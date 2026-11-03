package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionButton_base func(uintptr) uintptr
	addr_QStyleOptionButton_new uintptr
	addr_QStyleOptionButton_new2 uintptr
	addr_QStyleOptionButton_OperatorAssign uintptr
	addr_QStyleOptionButton_features_get uintptr
	addr_QStyleOptionButton_features_set uintptr
	addr_QStyleOptionButton_text_get uintptr
	addr_QStyleOptionButton_text_set uintptr
	addr_QStyleOptionButton_icon_get uintptr
	addr_QStyleOptionButton_icon_set uintptr
	addr_QStyleOptionButton_iconSize_get uintptr
	addr_QStyleOptionButton_iconSize_set uintptr
	pg_QStyleOptionButton_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionButton_base, "QStyleOptionButton_base")
	symaddr(&addr_QStyleOptionButton_new, "QStyleOptionButton_new")
	symaddr(&addr_QStyleOptionButton_new2, "QStyleOptionButton_new2")
	symaddr(&addr_QStyleOptionButton_OperatorAssign, "QStyleOptionButton_OperatorAssign")
	symaddr(&addr_QStyleOptionButton_features_get, "QStyleOptionButton_features_get")
	symaddr(&addr_QStyleOptionButton_features_set, "QStyleOptionButton_features_set")
	symaddr(&addr_QStyleOptionButton_text_get, "QStyleOptionButton_text_get")
	symaddr(&addr_QStyleOptionButton_text_set, "QStyleOptionButton_text_set")
	symaddr(&addr_QStyleOptionButton_icon_get, "QStyleOptionButton_icon_get")
	symaddr(&addr_QStyleOptionButton_icon_set, "QStyleOptionButton_icon_set")
	symaddr(&addr_QStyleOptionButton_iconSize_get, "QStyleOptionButton_iconSize_get")
	symaddr(&addr_QStyleOptionButton_iconSize_set, "QStyleOptionButton_iconSize_set")
	reg(&pg_QStyleOptionButton_delete, "QStyleOptionButton_delete")
}

type QStyleOptionButton struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionButton(h uintptr) *QStyleOptionButton {
	if h == 0 {
		return nil
	}
	return &QStyleOptionButton{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionButton_base(h))}
}

func newownedQStyleOptionButton(h uintptr) *QStyleOptionButton {
	o := bindQStyleOptionButton(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionButton).Delete)
	}
	return o
}

func (this *QStyleOptionButton) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionButton) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionButton(h unsafe.Pointer) *QStyleOptionButton { return bindQStyleOptionButton(uintptr(h)) }
func UnsafeNewQStyleOptionButtonOwned(h unsafe.Pointer) *QStyleOptionButton { return newownedQStyleOptionButton(uintptr(h)) }

func (this *QStyleOptionButton) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionButton() *QStyleOptionButton {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionButton_new)
	__o := newownedQStyleOptionButton(__r)
	return __o
}

func NewQStyleOptionButton2(other *QStyleOptionButton) *QStyleOptionButton {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionButton_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionButton(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionButton) OperatorAssign(arg1 *QStyleOptionButton) *QStyleOptionButton {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionButton_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionButton(__r)
}

func (this *QStyleOptionButton) Features() QStyleOptionButton__ButtonFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionButton_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionButton__ButtonFeature(__r)
}

func (this *QStyleOptionButton) SetFeatures(features QStyleOptionButton__ButtonFeature) {
	purego.SyscallN(addr_QStyleOptionButton_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionButton) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionButton_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionButton) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionButton_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionButton) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionButton_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionButton) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionButton_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionButton) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionButton_iconSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionButton) SetIconSize(iconSize *QSize) {
	purego.SyscallN(addr_QStyleOptionButton_iconSize_set, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QStyleOptionButton) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionButton_delete(this.h)
	}
	this.UnsafeRelease()
}
