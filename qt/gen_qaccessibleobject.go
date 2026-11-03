package qt

import (
	"unsafe"
	"runtime"
)

var (
	pg_QAccessibleObject_base func(uintptr) uintptr
)

func init() {
	reg(&pg_QAccessibleObject_base, "QAccessibleObject_base")
}

type QAccessibleObject struct {
	h uintptr
	*QAccessibleInterface
}

func bindQAccessibleObject(h uintptr) *QAccessibleObject {
	if h == 0 {
		return nil
	}
	return &QAccessibleObject{h: h, QAccessibleInterface: bindQAccessibleInterface(pg_QAccessibleObject_base(h))}
}

func newownedQAccessibleObject(h uintptr) *QAccessibleObject { return bindQAccessibleObject(h) }

func (this *QAccessibleObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleObject(h unsafe.Pointer) *QAccessibleObject { return bindQAccessibleObject(uintptr(h)) }
func UnsafeNewQAccessibleObjectOwned(h unsafe.Pointer) *QAccessibleObject { return newownedQAccessibleObject(uintptr(h)) }

func (this *QAccessibleObject) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleInterface != nil {
		this.QAccessibleInterface.UnsafeRelease()
	}
}

func (this *QAccessibleObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
