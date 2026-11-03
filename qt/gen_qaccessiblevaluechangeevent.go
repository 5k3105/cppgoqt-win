package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibleValueChangeEvent_base func(uintptr) uintptr
	addr_QAccessibleValueChangeEvent_new uintptr
	addr_QAccessibleValueChangeEvent_new2 uintptr
	addr_QAccessibleValueChangeEvent_setValue uintptr
	addr_QAccessibleValueChangeEvent_value uintptr
	pg_QAccessibleValueChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibleValueChangeEvent_base, "QAccessibleValueChangeEvent_base")
	symaddr(&addr_QAccessibleValueChangeEvent_new, "QAccessibleValueChangeEvent_new")
	symaddr(&addr_QAccessibleValueChangeEvent_new2, "QAccessibleValueChangeEvent_new2")
	symaddr(&addr_QAccessibleValueChangeEvent_setValue, "QAccessibleValueChangeEvent_setValue")
	symaddr(&addr_QAccessibleValueChangeEvent_value, "QAccessibleValueChangeEvent_value")
	reg(&pg_QAccessibleValueChangeEvent_delete, "QAccessibleValueChangeEvent_delete")
}

type QAccessibleValueChangeEvent struct {
	h uintptr
	*QAccessibleEvent
}

func bindQAccessibleValueChangeEvent(h uintptr) *QAccessibleValueChangeEvent {
	if h == 0 {
		return nil
	}
	return &QAccessibleValueChangeEvent{h: h, QAccessibleEvent: bindQAccessibleEvent(pg_QAccessibleValueChangeEvent_base(h))}
}

func newownedQAccessibleValueChangeEvent(h uintptr) *QAccessibleValueChangeEvent {
	o := bindQAccessibleValueChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleValueChangeEvent).Delete)
	}
	return o
}

func (this *QAccessibleValueChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleValueChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleValueChangeEvent(h unsafe.Pointer) *QAccessibleValueChangeEvent { return bindQAccessibleValueChangeEvent(uintptr(h)) }
func UnsafeNewQAccessibleValueChangeEventOwned(h unsafe.Pointer) *QAccessibleValueChangeEvent { return newownedQAccessibleValueChangeEvent(uintptr(h)) }

func (this *QAccessibleValueChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleEvent != nil {
		this.QAccessibleEvent.UnsafeRelease()
	}
}

func NewQAccessibleValueChangeEvent(obj *QObject, val *QVariant) *QAccessibleValueChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueChangeEvent_new, uintptr(obj.UnsafePointer()), uintptr(val.UnsafePointer()))
	__o := newownedQAccessibleValueChangeEvent(__r)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(val)
	return __o
}

func NewQAccessibleValueChangeEvent2(iface *QAccessibleInterface, val *QVariant) *QAccessibleValueChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueChangeEvent_new2, uintptr(iface.UnsafePointer()), uintptr(val.UnsafePointer()))
	__o := newownedQAccessibleValueChangeEvent(__r)
	runtime.KeepAlive(iface)
	runtime.KeepAlive(val)
	return __o
}

func (this *QAccessibleValueChangeEvent) SetValue(val *QVariant) {
	purego.SyscallN(addr_QAccessibleValueChangeEvent_setValue, this.h, uintptr(val.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val)
}

func (this *QAccessibleValueChangeEvent) Value() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueChangeEvent_value, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAccessibleValueChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleValueChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
