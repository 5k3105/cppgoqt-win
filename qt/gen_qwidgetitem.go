package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QWidgetItem_base func(uintptr) uintptr
	addr_QWidgetItem_new uintptr
	pg_QWidgetItem_delete func(uintptr)
)

func init() {
	reg(&pg_QWidgetItem_base, "QWidgetItem_base")
	symaddr(&addr_QWidgetItem_new, "QWidgetItem_new")
	reg(&pg_QWidgetItem_delete, "QWidgetItem_delete")
}

type QWidgetItem struct {
	h uintptr
	*QLayoutItem
}

func bindQWidgetItem(h uintptr) *QWidgetItem {
	if h == 0 {
		return nil
	}
	return &QWidgetItem{h: h, QLayoutItem: bindQLayoutItem(pg_QWidgetItem_base(h))}
}

func newownedQWidgetItem(h uintptr) *QWidgetItem {
	o := bindQWidgetItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWidgetItem).Delete)
	}
	return o
}

func (this *QWidgetItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWidgetItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWidgetItem(h unsafe.Pointer) *QWidgetItem { return bindQWidgetItem(uintptr(h)) }
func UnsafeNewQWidgetItemOwned(h unsafe.Pointer) *QWidgetItem { return newownedQWidgetItem(uintptr(h)) }

func (this *QWidgetItem) UnsafeRelease() {
	this.h = 0
	if this.QLayoutItem != nil {
		this.QLayoutItem.UnsafeRelease()
	}
}

func NewQWidgetItem(w *QWidget) *QWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QWidgetItem_new, uintptr(w.UnsafePointer()))
	__o := newownedQWidgetItem(__r)
	runtime.KeepAlive(w)
	return __o
}

func (this *QWidgetItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWidgetItem_delete(this.h)
	}
	this.UnsafeRelease()
}
