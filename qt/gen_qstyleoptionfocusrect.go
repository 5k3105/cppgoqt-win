package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionFocusRect_base func(uintptr) uintptr
	addr_QStyleOptionFocusRect_new uintptr
	addr_QStyleOptionFocusRect_new2 uintptr
	addr_QStyleOptionFocusRect_OperatorAssign uintptr
	addr_QStyleOptionFocusRect_backgroundColor_get uintptr
	addr_QStyleOptionFocusRect_backgroundColor_set uintptr
	pg_QStyleOptionFocusRect_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionFocusRect_base, "QStyleOptionFocusRect_base")
	symaddr(&addr_QStyleOptionFocusRect_new, "QStyleOptionFocusRect_new")
	symaddr(&addr_QStyleOptionFocusRect_new2, "QStyleOptionFocusRect_new2")
	symaddr(&addr_QStyleOptionFocusRect_OperatorAssign, "QStyleOptionFocusRect_OperatorAssign")
	symaddr(&addr_QStyleOptionFocusRect_backgroundColor_get, "QStyleOptionFocusRect_backgroundColor_get")
	symaddr(&addr_QStyleOptionFocusRect_backgroundColor_set, "QStyleOptionFocusRect_backgroundColor_set")
	reg(&pg_QStyleOptionFocusRect_delete, "QStyleOptionFocusRect_delete")
}

type QStyleOptionFocusRect struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionFocusRect(h uintptr) *QStyleOptionFocusRect {
	if h == 0 {
		return nil
	}
	return &QStyleOptionFocusRect{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionFocusRect_base(h))}
}

func newownedQStyleOptionFocusRect(h uintptr) *QStyleOptionFocusRect {
	o := bindQStyleOptionFocusRect(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionFocusRect).Delete)
	}
	return o
}

func (this *QStyleOptionFocusRect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionFocusRect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionFocusRect(h unsafe.Pointer) *QStyleOptionFocusRect { return bindQStyleOptionFocusRect(uintptr(h)) }
func UnsafeNewQStyleOptionFocusRectOwned(h unsafe.Pointer) *QStyleOptionFocusRect { return newownedQStyleOptionFocusRect(uintptr(h)) }

func (this *QStyleOptionFocusRect) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFocusRect_new)
	__o := newownedQStyleOptionFocusRect(__r)
	return __o
}

func NewQStyleOptionFocusRect2(other *QStyleOptionFocusRect) *QStyleOptionFocusRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFocusRect_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionFocusRect(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionFocusRect) OperatorAssign(arg1 *QStyleOptionFocusRect) *QStyleOptionFocusRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFocusRect_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionFocusRect(__r)
}

func (this *QStyleOptionFocusRect) BackgroundColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFocusRect_backgroundColor_get, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QStyleOptionFocusRect) SetBackgroundColor(backgroundColor *QColor) {
	purego.SyscallN(addr_QStyleOptionFocusRect_backgroundColor_set, this.h, uintptr(backgroundColor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(backgroundColor)
}

func (this *QStyleOptionFocusRect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionFocusRect_delete(this.h)
	}
	this.UnsafeRelease()
}
