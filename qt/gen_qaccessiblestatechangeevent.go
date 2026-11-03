package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibleStateChangeEvent_base func(uintptr) uintptr
	addr_QAccessibleStateChangeEvent_new uintptr
	addr_QAccessibleStateChangeEvent_new2 uintptr
	addr_QAccessibleStateChangeEvent_changedStates uintptr
	pg_QAccessibleStateChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibleStateChangeEvent_base, "QAccessibleStateChangeEvent_base")
	symaddr(&addr_QAccessibleStateChangeEvent_new, "QAccessibleStateChangeEvent_new")
	symaddr(&addr_QAccessibleStateChangeEvent_new2, "QAccessibleStateChangeEvent_new2")
	symaddr(&addr_QAccessibleStateChangeEvent_changedStates, "QAccessibleStateChangeEvent_changedStates")
	reg(&pg_QAccessibleStateChangeEvent_delete, "QAccessibleStateChangeEvent_delete")
}

type QAccessibleStateChangeEvent struct {
	h uintptr
	*QAccessibleEvent
}

func bindQAccessibleStateChangeEvent(h uintptr) *QAccessibleStateChangeEvent {
	if h == 0 {
		return nil
	}
	return &QAccessibleStateChangeEvent{h: h, QAccessibleEvent: bindQAccessibleEvent(pg_QAccessibleStateChangeEvent_base(h))}
}

func newownedQAccessibleStateChangeEvent(h uintptr) *QAccessibleStateChangeEvent {
	o := bindQAccessibleStateChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleStateChangeEvent).Delete)
	}
	return o
}

func (this *QAccessibleStateChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleStateChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleStateChangeEvent(h unsafe.Pointer) *QAccessibleStateChangeEvent { return bindQAccessibleStateChangeEvent(uintptr(h)) }
func UnsafeNewQAccessibleStateChangeEventOwned(h unsafe.Pointer) *QAccessibleStateChangeEvent { return newownedQAccessibleStateChangeEvent(uintptr(h)) }

func (this *QAccessibleStateChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleEvent != nil {
		this.QAccessibleEvent.UnsafeRelease()
	}
}

func NewQAccessibleStateChangeEvent(obj *QObject, state *QAccessible__State) *QAccessibleStateChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleStateChangeEvent_new, uintptr(obj.UnsafePointer()), uintptr(state.UnsafePointer()))
	__o := newownedQAccessibleStateChangeEvent(__r)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(state)
	return __o
}

func NewQAccessibleStateChangeEvent2(iface *QAccessibleInterface, state *QAccessible__State) *QAccessibleStateChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleStateChangeEvent_new2, uintptr(iface.UnsafePointer()), uintptr(state.UnsafePointer()))
	__o := newownedQAccessibleStateChangeEvent(__r)
	runtime.KeepAlive(iface)
	runtime.KeepAlive(state)
	return __o
}

func (this *QAccessibleStateChangeEvent) ChangedStates() *QAccessible__State {
	__r, _, _ := purego.SyscallN(addr_QAccessibleStateChangeEvent_changedStates, this.h)
	runtime.KeepAlive(this)
	return newownedQAccessible__State(__r)
}

func (this *QAccessibleStateChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleStateChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
