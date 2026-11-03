package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QChildWindowEvent_base func(uintptr) uintptr
	addr_QChildWindowEvent_new uintptr
	addr_QChildWindowEvent_child uintptr
	addr_QChildWindowEvent_registerEventType uintptr
	addr_QChildWindowEvent_registerEventType2 uintptr
	pg_QChildWindowEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QChildWindowEvent_base, "QChildWindowEvent_base")
	symaddr(&addr_QChildWindowEvent_new, "QChildWindowEvent_new")
	symaddr(&addr_QChildWindowEvent_child, "QChildWindowEvent_child")
	symaddr(&addr_QChildWindowEvent_registerEventType, "QChildWindowEvent_registerEventType")
	symaddr(&addr_QChildWindowEvent_registerEventType2, "QChildWindowEvent_registerEventType2")
	reg(&pg_QChildWindowEvent_delete, "QChildWindowEvent_delete")
}

type QChildWindowEvent struct {
	h uintptr
	*QEvent
}

func bindQChildWindowEvent(h uintptr) *QChildWindowEvent {
	if h == 0 {
		return nil
	}
	return &QChildWindowEvent{h: h, QEvent: bindQEvent(pg_QChildWindowEvent_base(h))}
}

func newownedQChildWindowEvent(h uintptr) *QChildWindowEvent {
	o := bindQChildWindowEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QChildWindowEvent).Delete)
	}
	return o
}

func (this *QChildWindowEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QChildWindowEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQChildWindowEvent(h unsafe.Pointer) *QChildWindowEvent { return bindQChildWindowEvent(uintptr(h)) }
func UnsafeNewQChildWindowEventOwned(h unsafe.Pointer) *QChildWindowEvent { return newownedQChildWindowEvent(uintptr(h)) }

func (this *QChildWindowEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQChildWindowEvent(type_ QEvent__Type, childWindow *QWindow) *QChildWindowEvent {
	__r, _, _ := purego.SyscallN(addr_QChildWindowEvent_new, uintptr(type_), uintptr(childWindow.UnsafePointer()))
	__o := newownedQChildWindowEvent(__r)
	runtime.KeepAlive(childWindow)
	return __o
}

func (this *QChildWindowEvent) Child() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QChildWindowEvent_child, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func QChildWindowEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QChildWindowEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QChildWindowEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QChildWindowEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QChildWindowEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QChildWindowEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
