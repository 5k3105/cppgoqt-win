package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionGraphicsItem_base func(uintptr) uintptr
	addr_QStyleOptionGraphicsItem_new uintptr
	addr_QStyleOptionGraphicsItem_new2 uintptr
	addr_QStyleOptionGraphicsItem_OperatorAssign uintptr
	pg_QStyleOptionGraphicsItem_levelOfDetailFromTransform func(uintptr) float64
	addr_QStyleOptionGraphicsItem_exposedRect_get uintptr
	addr_QStyleOptionGraphicsItem_exposedRect_set uintptr
	pg_QStyleOptionGraphicsItem_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionGraphicsItem_base, "QStyleOptionGraphicsItem_base")
	symaddr(&addr_QStyleOptionGraphicsItem_new, "QStyleOptionGraphicsItem_new")
	symaddr(&addr_QStyleOptionGraphicsItem_new2, "QStyleOptionGraphicsItem_new2")
	symaddr(&addr_QStyleOptionGraphicsItem_OperatorAssign, "QStyleOptionGraphicsItem_OperatorAssign")
	reg(&pg_QStyleOptionGraphicsItem_levelOfDetailFromTransform, "QStyleOptionGraphicsItem_levelOfDetailFromTransform")
	symaddr(&addr_QStyleOptionGraphicsItem_exposedRect_get, "QStyleOptionGraphicsItem_exposedRect_get")
	symaddr(&addr_QStyleOptionGraphicsItem_exposedRect_set, "QStyleOptionGraphicsItem_exposedRect_set")
	reg(&pg_QStyleOptionGraphicsItem_delete, "QStyleOptionGraphicsItem_delete")
}

type QStyleOptionGraphicsItem struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionGraphicsItem(h uintptr) *QStyleOptionGraphicsItem {
	if h == 0 {
		return nil
	}
	return &QStyleOptionGraphicsItem{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionGraphicsItem_base(h))}
}

func newownedQStyleOptionGraphicsItem(h uintptr) *QStyleOptionGraphicsItem {
	o := bindQStyleOptionGraphicsItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionGraphicsItem).Delete)
	}
	return o
}

func (this *QStyleOptionGraphicsItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionGraphicsItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionGraphicsItem(h unsafe.Pointer) *QStyleOptionGraphicsItem { return bindQStyleOptionGraphicsItem(uintptr(h)) }
func UnsafeNewQStyleOptionGraphicsItemOwned(h unsafe.Pointer) *QStyleOptionGraphicsItem { return newownedQStyleOptionGraphicsItem(uintptr(h)) }

func (this *QStyleOptionGraphicsItem) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGraphicsItem_new)
	__o := newownedQStyleOptionGraphicsItem(__r)
	return __o
}

func NewQStyleOptionGraphicsItem2(other *QStyleOptionGraphicsItem) *QStyleOptionGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGraphicsItem_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionGraphicsItem(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionGraphicsItem) OperatorAssign(arg1 *QStyleOptionGraphicsItem) *QStyleOptionGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGraphicsItem_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionGraphicsItem(__r)
}

func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform *QTransform) float64 {
	__r := pg_QStyleOptionGraphicsItem_levelOfDetailFromTransform(uintptr(worldTransform.UnsafePointer()))
	runtime.KeepAlive(worldTransform)
	return __r
}

func (this *QStyleOptionGraphicsItem) ExposedRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionGraphicsItem_exposedRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QStyleOptionGraphicsItem) SetExposedRect(exposedRect *QRectF) {
	purego.SyscallN(addr_QStyleOptionGraphicsItem_exposedRect_set, this.h, uintptr(exposedRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exposedRect)
}

func (this *QStyleOptionGraphicsItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionGraphicsItem_delete(this.h)
	}
	this.UnsafeRelease()
}
