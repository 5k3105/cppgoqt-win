package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QActionEvent_base func(uintptr) uintptr
	addr_QActionEvent_new uintptr
	addr_QActionEvent_new2 uintptr
	addr_QActionEvent_action uintptr
	addr_QActionEvent_before uintptr
	addr_QActionEvent_registerEventType uintptr
	addr_QActionEvent_registerEventType2 uintptr
	pg_QActionEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QActionEvent_base, "QActionEvent_base")
	symaddr(&addr_QActionEvent_new, "QActionEvent_new")
	symaddr(&addr_QActionEvent_new2, "QActionEvent_new2")
	symaddr(&addr_QActionEvent_action, "QActionEvent_action")
	symaddr(&addr_QActionEvent_before, "QActionEvent_before")
	symaddr(&addr_QActionEvent_registerEventType, "QActionEvent_registerEventType")
	symaddr(&addr_QActionEvent_registerEventType2, "QActionEvent_registerEventType2")
	reg(&pg_QActionEvent_delete, "QActionEvent_delete")
}

type QActionEvent struct {
	h uintptr
	*QEvent
}

func bindQActionEvent(h uintptr) *QActionEvent {
	if h == 0 {
		return nil
	}
	return &QActionEvent{h: h, QEvent: bindQEvent(pg_QActionEvent_base(h))}
}

func newownedQActionEvent(h uintptr) *QActionEvent {
	o := bindQActionEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QActionEvent).Delete)
	}
	return o
}

func (this *QActionEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QActionEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQActionEvent(h unsafe.Pointer) *QActionEvent { return bindQActionEvent(uintptr(h)) }
func UnsafeNewQActionEventOwned(h unsafe.Pointer) *QActionEvent { return newownedQActionEvent(uintptr(h)) }

func (this *QActionEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQActionEvent(type_ int, action *QAction) *QActionEvent {
	__r, _, _ := purego.SyscallN(addr_QActionEvent_new, uintptr(type_), uintptr(action.UnsafePointer()))
	__o := newownedQActionEvent(__r)
	runtime.KeepAlive(action)
	return __o
}

func NewQActionEvent2(type_ int, action *QAction, before *QAction) *QActionEvent {
	__r, _, _ := purego.SyscallN(addr_QActionEvent_new2, uintptr(type_), uintptr(action.UnsafePointer()), uintptr(before.UnsafePointer()))
	__o := newownedQActionEvent(__r)
	runtime.KeepAlive(action)
	runtime.KeepAlive(before)
	return __o
}

func (this *QActionEvent) Action() *QAction {
	__r, _, _ := purego.SyscallN(addr_QActionEvent_action, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QActionEvent) Before() *QAction {
	__r, _, _ := purego.SyscallN(addr_QActionEvent_before, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func QActionEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QActionEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QActionEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QActionEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QActionEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QActionEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
