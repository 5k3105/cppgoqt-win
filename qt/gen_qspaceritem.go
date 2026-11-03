package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSpacerItem_base func(uintptr) uintptr
	addr_QSpacerItem_new uintptr
	addr_QSpacerItem_new2 uintptr
	addr_QSpacerItem_new3 uintptr
	addr_QSpacerItem_new4 uintptr
	addr_QSpacerItem_changeSize uintptr
	addr_QSpacerItem_sizePolicy uintptr
	addr_QSpacerItem_OperatorAssign uintptr
	addr_QSpacerItem_changeSize2 uintptr
	addr_QSpacerItem_changeSize3 uintptr
	pg_QSpacerItem_delete func(uintptr)
)

func init() {
	reg(&pg_QSpacerItem_base, "QSpacerItem_base")
	symaddr(&addr_QSpacerItem_new, "QSpacerItem_new")
	symaddr(&addr_QSpacerItem_new2, "QSpacerItem_new2")
	symaddr(&addr_QSpacerItem_new3, "QSpacerItem_new3")
	symaddr(&addr_QSpacerItem_new4, "QSpacerItem_new4")
	symaddr(&addr_QSpacerItem_changeSize, "QSpacerItem_changeSize")
	symaddr(&addr_QSpacerItem_sizePolicy, "QSpacerItem_sizePolicy")
	symaddr(&addr_QSpacerItem_OperatorAssign, "QSpacerItem_OperatorAssign")
	symaddr(&addr_QSpacerItem_changeSize2, "QSpacerItem_changeSize2")
	symaddr(&addr_QSpacerItem_changeSize3, "QSpacerItem_changeSize3")
	reg(&pg_QSpacerItem_delete, "QSpacerItem_delete")
}

type QSpacerItem struct {
	h uintptr
	*QLayoutItem
}

func bindQSpacerItem(h uintptr) *QSpacerItem {
	if h == 0 {
		return nil
	}
	return &QSpacerItem{h: h, QLayoutItem: bindQLayoutItem(pg_QSpacerItem_base(h))}
}

func newownedQSpacerItem(h uintptr) *QSpacerItem {
	o := bindQSpacerItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSpacerItem).Delete)
	}
	return o
}

func (this *QSpacerItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSpacerItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSpacerItem(h unsafe.Pointer) *QSpacerItem { return bindQSpacerItem(uintptr(h)) }
func UnsafeNewQSpacerItemOwned(h unsafe.Pointer) *QSpacerItem { return newownedQSpacerItem(uintptr(h)) }

func (this *QSpacerItem) UnsafeRelease() {
	this.h = 0
	if this.QLayoutItem != nil {
		this.QLayoutItem.UnsafeRelease()
	}
}

func NewQSpacerItem(w int, h int) *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QSpacerItem_new, uintptr(w), uintptr(h))
	__o := newownedQSpacerItem(__r)
	return __o
}

func NewQSpacerItem2(arg1 *QSpacerItem) *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QSpacerItem_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQSpacerItem(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQSpacerItem3(w int, h int, hData QSizePolicy__Policy) *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QSpacerItem_new3, uintptr(w), uintptr(h), uintptr(hData))
	__o := newownedQSpacerItem(__r)
	return __o
}

func NewQSpacerItem4(w int, h int, hData QSizePolicy__Policy, vData QSizePolicy__Policy) *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QSpacerItem_new4, uintptr(w), uintptr(h), uintptr(hData), uintptr(vData))
	__o := newownedQSpacerItem(__r)
	return __o
}

func (this *QSpacerItem) ChangeSize(w int, h int) {
	purego.SyscallN(addr_QSpacerItem_changeSize, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QSpacerItem) SizePolicy() *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSpacerItem_sizePolicy, this.h)
	runtime.KeepAlive(this)
	return newownedQSizePolicy(__r)
}

func (this *QSpacerItem) OperatorAssign(arg1 *QSpacerItem) *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QSpacerItem_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSpacerItem(__r)
}

func (this *QSpacerItem) ChangeSize2(w int, h int, hData QSizePolicy__Policy) {
	purego.SyscallN(addr_QSpacerItem_changeSize2, this.h, uintptr(w), uintptr(h), uintptr(hData))
	runtime.KeepAlive(this)
}

func (this *QSpacerItem) ChangeSize3(w int, h int, hData QSizePolicy__Policy, vData QSizePolicy__Policy) {
	purego.SyscallN(addr_QSpacerItem_changeSize3, this.h, uintptr(w), uintptr(h), uintptr(hData), uintptr(vData))
	runtime.KeepAlive(this)
}

func (this *QSpacerItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSpacerItem_delete(this.h)
	}
	this.UnsafeRelease()
}
