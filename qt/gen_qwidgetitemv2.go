package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWidgetItemV2_chain uintptr
	addr_QWidgetItemV2_new uintptr
	pg_QWidgetItemV2_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWidgetItemV2_chain, "QWidgetItemV2_chain")
	symaddr(&addr_QWidgetItemV2_new, "QWidgetItemV2_new")
	reg(&pg_QWidgetItemV2_delete, "QWidgetItemV2_delete")
}

type QWidgetItemV2 struct {
	h uintptr
	*QWidgetItem
}

func bindQWidgetItemV2(h uintptr) *QWidgetItemV2 {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QWidgetItemV2_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QWidgetItemV2{h: h, QWidgetItem: &QWidgetItem{h: __c[0], QLayoutItem: &QLayoutItem{h: __c[1]}}}
}

func newownedQWidgetItemV2(h uintptr) *QWidgetItemV2 {
	o := bindQWidgetItemV2(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWidgetItemV2).Delete)
	}
	return o
}

func (this *QWidgetItemV2) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWidgetItemV2) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWidgetItemV2(h unsafe.Pointer) *QWidgetItemV2 { return bindQWidgetItemV2(uintptr(h)) }
func UnsafeNewQWidgetItemV2Owned(h unsafe.Pointer) *QWidgetItemV2 { return newownedQWidgetItemV2(uintptr(h)) }

func (this *QWidgetItemV2) UnsafeRelease() {
	this.h = 0
	if this.QWidgetItem != nil {
		this.QWidgetItem.UnsafeRelease()
	}
}

func NewQWidgetItemV2(widget *QWidget) *QWidgetItemV2 {
	__r, _, _ := purego.SyscallN(addr_QWidgetItemV2_new, uintptr(widget.UnsafePointer()))
	__o := newownedQWidgetItemV2(__r)
	runtime.KeepAlive(widget)
	return __o
}

func (this *QWidgetItemV2) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWidgetItemV2_delete(this.h)
	}
	this.UnsafeRelease()
}
