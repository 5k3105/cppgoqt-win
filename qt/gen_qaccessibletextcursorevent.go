package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibleTextCursorEvent_base func(uintptr) uintptr
	addr_QAccessibleTextCursorEvent_new uintptr
	addr_QAccessibleTextCursorEvent_new2 uintptr
	addr_QAccessibleTextCursorEvent_setCursorPosition uintptr
	addr_QAccessibleTextCursorEvent_cursorPosition uintptr
	pg_QAccessibleTextCursorEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibleTextCursorEvent_base, "QAccessibleTextCursorEvent_base")
	symaddr(&addr_QAccessibleTextCursorEvent_new, "QAccessibleTextCursorEvent_new")
	symaddr(&addr_QAccessibleTextCursorEvent_new2, "QAccessibleTextCursorEvent_new2")
	symaddr(&addr_QAccessibleTextCursorEvent_setCursorPosition, "QAccessibleTextCursorEvent_setCursorPosition")
	symaddr(&addr_QAccessibleTextCursorEvent_cursorPosition, "QAccessibleTextCursorEvent_cursorPosition")
	reg(&pg_QAccessibleTextCursorEvent_delete, "QAccessibleTextCursorEvent_delete")
}

type QAccessibleTextCursorEvent struct {
	h uintptr
	*QAccessibleEvent
}

func bindQAccessibleTextCursorEvent(h uintptr) *QAccessibleTextCursorEvent {
	if h == 0 {
		return nil
	}
	return &QAccessibleTextCursorEvent{h: h, QAccessibleEvent: bindQAccessibleEvent(pg_QAccessibleTextCursorEvent_base(h))}
}

func newownedQAccessibleTextCursorEvent(h uintptr) *QAccessibleTextCursorEvent {
	o := bindQAccessibleTextCursorEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTextCursorEvent).Delete)
	}
	return o
}

func (this *QAccessibleTextCursorEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTextCursorEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTextCursorEvent(h unsafe.Pointer) *QAccessibleTextCursorEvent { return bindQAccessibleTextCursorEvent(uintptr(h)) }
func UnsafeNewQAccessibleTextCursorEventOwned(h unsafe.Pointer) *QAccessibleTextCursorEvent { return newownedQAccessibleTextCursorEvent(uintptr(h)) }

func (this *QAccessibleTextCursorEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleEvent != nil {
		this.QAccessibleEvent.UnsafeRelease()
	}
}

func NewQAccessibleTextCursorEvent(obj *QObject, cursorPos int) *QAccessibleTextCursorEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextCursorEvent_new, uintptr(obj.UnsafePointer()), uintptr(cursorPos))
	__o := newownedQAccessibleTextCursorEvent(__r)
	runtime.KeepAlive(obj)
	return __o
}

func NewQAccessibleTextCursorEvent2(iface *QAccessibleInterface, cursorPos int) *QAccessibleTextCursorEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextCursorEvent_new2, uintptr(iface.UnsafePointer()), uintptr(cursorPos))
	__o := newownedQAccessibleTextCursorEvent(__r)
	runtime.KeepAlive(iface)
	return __o
}

func (this *QAccessibleTextCursorEvent) SetCursorPosition(position int) {
	purego.SyscallN(addr_QAccessibleTextCursorEvent_setCursorPosition, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextCursorEvent) CursorPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextCursorEvent_cursorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextCursorEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTextCursorEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
