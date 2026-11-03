package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionMenuItem_base func(uintptr) uintptr
	addr_QStyleOptionMenuItem_new uintptr
	addr_QStyleOptionMenuItem_new2 uintptr
	addr_QStyleOptionMenuItem_OperatorAssign uintptr
	addr_QStyleOptionMenuItem_menuItemType_get uintptr
	addr_QStyleOptionMenuItem_menuItemType_set uintptr
	addr_QStyleOptionMenuItem_checkType_get uintptr
	addr_QStyleOptionMenuItem_checkType_set uintptr
	addr_QStyleOptionMenuItem_checked_get uintptr
	addr_QStyleOptionMenuItem_checked_set uintptr
	addr_QStyleOptionMenuItem_menuHasCheckableItems_get uintptr
	addr_QStyleOptionMenuItem_menuHasCheckableItems_set uintptr
	addr_QStyleOptionMenuItem_menuRect_get uintptr
	addr_QStyleOptionMenuItem_menuRect_set uintptr
	addr_QStyleOptionMenuItem_text_get uintptr
	addr_QStyleOptionMenuItem_text_set uintptr
	addr_QStyleOptionMenuItem_icon_get uintptr
	addr_QStyleOptionMenuItem_icon_set uintptr
	addr_QStyleOptionMenuItem_maxIconWidth_get uintptr
	addr_QStyleOptionMenuItem_maxIconWidth_set uintptr
	addr_QStyleOptionMenuItem_reservedShortcutWidth_get uintptr
	addr_QStyleOptionMenuItem_reservedShortcutWidth_set uintptr
	addr_QStyleOptionMenuItem_font_get uintptr
	addr_QStyleOptionMenuItem_font_set uintptr
	pg_QStyleOptionMenuItem_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionMenuItem_base, "QStyleOptionMenuItem_base")
	symaddr(&addr_QStyleOptionMenuItem_new, "QStyleOptionMenuItem_new")
	symaddr(&addr_QStyleOptionMenuItem_new2, "QStyleOptionMenuItem_new2")
	symaddr(&addr_QStyleOptionMenuItem_OperatorAssign, "QStyleOptionMenuItem_OperatorAssign")
	symaddr(&addr_QStyleOptionMenuItem_menuItemType_get, "QStyleOptionMenuItem_menuItemType_get")
	symaddr(&addr_QStyleOptionMenuItem_menuItemType_set, "QStyleOptionMenuItem_menuItemType_set")
	symaddr(&addr_QStyleOptionMenuItem_checkType_get, "QStyleOptionMenuItem_checkType_get")
	symaddr(&addr_QStyleOptionMenuItem_checkType_set, "QStyleOptionMenuItem_checkType_set")
	symaddr(&addr_QStyleOptionMenuItem_checked_get, "QStyleOptionMenuItem_checked_get")
	symaddr(&addr_QStyleOptionMenuItem_checked_set, "QStyleOptionMenuItem_checked_set")
	symaddr(&addr_QStyleOptionMenuItem_menuHasCheckableItems_get, "QStyleOptionMenuItem_menuHasCheckableItems_get")
	symaddr(&addr_QStyleOptionMenuItem_menuHasCheckableItems_set, "QStyleOptionMenuItem_menuHasCheckableItems_set")
	symaddr(&addr_QStyleOptionMenuItem_menuRect_get, "QStyleOptionMenuItem_menuRect_get")
	symaddr(&addr_QStyleOptionMenuItem_menuRect_set, "QStyleOptionMenuItem_menuRect_set")
	symaddr(&addr_QStyleOptionMenuItem_text_get, "QStyleOptionMenuItem_text_get")
	symaddr(&addr_QStyleOptionMenuItem_text_set, "QStyleOptionMenuItem_text_set")
	symaddr(&addr_QStyleOptionMenuItem_icon_get, "QStyleOptionMenuItem_icon_get")
	symaddr(&addr_QStyleOptionMenuItem_icon_set, "QStyleOptionMenuItem_icon_set")
	symaddr(&addr_QStyleOptionMenuItem_maxIconWidth_get, "QStyleOptionMenuItem_maxIconWidth_get")
	symaddr(&addr_QStyleOptionMenuItem_maxIconWidth_set, "QStyleOptionMenuItem_maxIconWidth_set")
	symaddr(&addr_QStyleOptionMenuItem_reservedShortcutWidth_get, "QStyleOptionMenuItem_reservedShortcutWidth_get")
	symaddr(&addr_QStyleOptionMenuItem_reservedShortcutWidth_set, "QStyleOptionMenuItem_reservedShortcutWidth_set")
	symaddr(&addr_QStyleOptionMenuItem_font_get, "QStyleOptionMenuItem_font_get")
	symaddr(&addr_QStyleOptionMenuItem_font_set, "QStyleOptionMenuItem_font_set")
	reg(&pg_QStyleOptionMenuItem_delete, "QStyleOptionMenuItem_delete")
}

type QStyleOptionMenuItem struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionMenuItem(h uintptr) *QStyleOptionMenuItem {
	if h == 0 {
		return nil
	}
	return &QStyleOptionMenuItem{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionMenuItem_base(h))}
}

func newownedQStyleOptionMenuItem(h uintptr) *QStyleOptionMenuItem {
	o := bindQStyleOptionMenuItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionMenuItem).Delete)
	}
	return o
}

func (this *QStyleOptionMenuItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionMenuItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionMenuItem(h unsafe.Pointer) *QStyleOptionMenuItem { return bindQStyleOptionMenuItem(uintptr(h)) }
func UnsafeNewQStyleOptionMenuItemOwned(h unsafe.Pointer) *QStyleOptionMenuItem { return newownedQStyleOptionMenuItem(uintptr(h)) }

func (this *QStyleOptionMenuItem) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_new)
	__o := newownedQStyleOptionMenuItem(__r)
	return __o
}

func NewQStyleOptionMenuItem2(other *QStyleOptionMenuItem) *QStyleOptionMenuItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionMenuItem(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionMenuItem) OperatorAssign(arg1 *QStyleOptionMenuItem) *QStyleOptionMenuItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionMenuItem(__r)
}

func (this *QStyleOptionMenuItem) MenuItemType() QStyleOptionMenuItem__MenuItemType {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_menuItemType_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionMenuItem__MenuItemType(__r)
}

func (this *QStyleOptionMenuItem) SetMenuItemType(menuItemType QStyleOptionMenuItem__MenuItemType) {
	purego.SyscallN(addr_QStyleOptionMenuItem_menuItemType_set, this.h, uintptr(menuItemType))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItem) CheckType() QStyleOptionMenuItem__CheckType {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_checkType_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionMenuItem__CheckType(__r)
}

func (this *QStyleOptionMenuItem) SetCheckType(checkType QStyleOptionMenuItem__CheckType) {
	purego.SyscallN(addr_QStyleOptionMenuItem_checkType_set, this.h, uintptr(checkType))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItem) Checked() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_checked_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionMenuItem) SetChecked(checked bool) {
	purego.SyscallN(addr_QStyleOptionMenuItem_checked_set, this.h, uintptr(b2i(checked)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItem) MenuHasCheckableItems() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_menuHasCheckableItems_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionMenuItem) SetMenuHasCheckableItems(menuHasCheckableItems bool) {
	purego.SyscallN(addr_QStyleOptionMenuItem_menuHasCheckableItems_set, this.h, uintptr(b2i(menuHasCheckableItems)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItem) MenuRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_menuRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOptionMenuItem) SetMenuRect(menuRect *QRect) {
	purego.SyscallN(addr_QStyleOptionMenuItem_menuRect_set, this.h, uintptr(menuRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menuRect)
}

func (this *QStyleOptionMenuItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionMenuItem) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionMenuItem_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionMenuItem) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionMenuItem) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionMenuItem_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionMenuItem) MaxIconWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_maxIconWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionMenuItem) SetMaxIconWidth(maxIconWidth int) {
	purego.SyscallN(addr_QStyleOptionMenuItem_maxIconWidth_set, this.h, uintptr(maxIconWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItem) ReservedShortcutWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_reservedShortcutWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionMenuItem) SetReservedShortcutWidth(reservedShortcutWidth int) {
	purego.SyscallN(addr_QStyleOptionMenuItem_reservedShortcutWidth_set, this.h, uintptr(reservedShortcutWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItem_font_get, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QStyleOptionMenuItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QStyleOptionMenuItem_font_set, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QStyleOptionMenuItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionMenuItem_delete(this.h)
	}
	this.UnsafeRelease()
}
