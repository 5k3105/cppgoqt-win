package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionToolBox_base func(uintptr) uintptr
	addr_QStyleOptionToolBox_new uintptr
	addr_QStyleOptionToolBox_new2 uintptr
	addr_QStyleOptionToolBox_OperatorAssign uintptr
	addr_QStyleOptionToolBox_text_get uintptr
	addr_QStyleOptionToolBox_text_set uintptr
	addr_QStyleOptionToolBox_icon_get uintptr
	addr_QStyleOptionToolBox_icon_set uintptr
	addr_QStyleOptionToolBox_position_get uintptr
	addr_QStyleOptionToolBox_position_set uintptr
	addr_QStyleOptionToolBox_selectedPosition_get uintptr
	addr_QStyleOptionToolBox_selectedPosition_set uintptr
	pg_QStyleOptionToolBox_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionToolBox_base, "QStyleOptionToolBox_base")
	symaddr(&addr_QStyleOptionToolBox_new, "QStyleOptionToolBox_new")
	symaddr(&addr_QStyleOptionToolBox_new2, "QStyleOptionToolBox_new2")
	symaddr(&addr_QStyleOptionToolBox_OperatorAssign, "QStyleOptionToolBox_OperatorAssign")
	symaddr(&addr_QStyleOptionToolBox_text_get, "QStyleOptionToolBox_text_get")
	symaddr(&addr_QStyleOptionToolBox_text_set, "QStyleOptionToolBox_text_set")
	symaddr(&addr_QStyleOptionToolBox_icon_get, "QStyleOptionToolBox_icon_get")
	symaddr(&addr_QStyleOptionToolBox_icon_set, "QStyleOptionToolBox_icon_set")
	symaddr(&addr_QStyleOptionToolBox_position_get, "QStyleOptionToolBox_position_get")
	symaddr(&addr_QStyleOptionToolBox_position_set, "QStyleOptionToolBox_position_set")
	symaddr(&addr_QStyleOptionToolBox_selectedPosition_get, "QStyleOptionToolBox_selectedPosition_get")
	symaddr(&addr_QStyleOptionToolBox_selectedPosition_set, "QStyleOptionToolBox_selectedPosition_set")
	reg(&pg_QStyleOptionToolBox_delete, "QStyleOptionToolBox_delete")
}

type QStyleOptionToolBox struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionToolBox(h uintptr) *QStyleOptionToolBox {
	if h == 0 {
		return nil
	}
	return &QStyleOptionToolBox{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionToolBox_base(h))}
}

func newownedQStyleOptionToolBox(h uintptr) *QStyleOptionToolBox {
	o := bindQStyleOptionToolBox(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionToolBox).Delete)
	}
	return o
}

func (this *QStyleOptionToolBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionToolBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionToolBox(h unsafe.Pointer) *QStyleOptionToolBox { return bindQStyleOptionToolBox(uintptr(h)) }
func UnsafeNewQStyleOptionToolBoxOwned(h unsafe.Pointer) *QStyleOptionToolBox { return newownedQStyleOptionToolBox(uintptr(h)) }

func (this *QStyleOptionToolBox) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_new)
	__o := newownedQStyleOptionToolBox(__r)
	return __o
}

func NewQStyleOptionToolBox2(other *QStyleOptionToolBox) *QStyleOptionToolBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionToolBox(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionToolBox) OperatorAssign(arg1 *QStyleOptionToolBox) *QStyleOptionToolBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionToolBox(__r)
}

func (this *QStyleOptionToolBox) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionToolBox) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionToolBox_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionToolBox) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionToolBox) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionToolBox_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionToolBox) Position() QStyleOptionToolBox__TabPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_position_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionToolBox__TabPosition(__r)
}

func (this *QStyleOptionToolBox) SetPosition(position QStyleOptionToolBox__TabPosition) {
	purego.SyscallN(addr_QStyleOptionToolBox_position_set, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBox) SelectedPosition() QStyleOptionToolBox__SelectedPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBox_selectedPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionToolBox__SelectedPosition(__r)
}

func (this *QStyleOptionToolBox) SetSelectedPosition(selectedPosition QStyleOptionToolBox__SelectedPosition) {
	purego.SyscallN(addr_QStyleOptionToolBox_selectedPosition_set, this.h, uintptr(selectedPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionToolBox_delete(this.h)
	}
	this.UnsafeRelease()
}
