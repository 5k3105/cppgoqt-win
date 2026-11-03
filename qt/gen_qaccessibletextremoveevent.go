package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTextRemoveEvent_chain uintptr
	addr_QAccessibleTextRemoveEvent_new uintptr
	addr_QAccessibleTextRemoveEvent_new2 uintptr
	addr_QAccessibleTextRemoveEvent_textRemoved uintptr
	addr_QAccessibleTextRemoveEvent_changePosition uintptr
	pg_QAccessibleTextRemoveEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTextRemoveEvent_chain, "QAccessibleTextRemoveEvent_chain")
	symaddr(&addr_QAccessibleTextRemoveEvent_new, "QAccessibleTextRemoveEvent_new")
	symaddr(&addr_QAccessibleTextRemoveEvent_new2, "QAccessibleTextRemoveEvent_new2")
	symaddr(&addr_QAccessibleTextRemoveEvent_textRemoved, "QAccessibleTextRemoveEvent_textRemoved")
	symaddr(&addr_QAccessibleTextRemoveEvent_changePosition, "QAccessibleTextRemoveEvent_changePosition")
	reg(&pg_QAccessibleTextRemoveEvent_delete, "QAccessibleTextRemoveEvent_delete")
}

type QAccessibleTextRemoveEvent struct {
	h uintptr
	*QAccessibleTextCursorEvent
}

func bindQAccessibleTextRemoveEvent(h uintptr) *QAccessibleTextRemoveEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAccessibleTextRemoveEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleTextRemoveEvent{h: h, QAccessibleTextCursorEvent: &QAccessibleTextCursorEvent{h: __c[0], QAccessibleEvent: &QAccessibleEvent{h: __c[1]}}}
}

func newownedQAccessibleTextRemoveEvent(h uintptr) *QAccessibleTextRemoveEvent {
	o := bindQAccessibleTextRemoveEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTextRemoveEvent).Delete)
	}
	return o
}

func (this *QAccessibleTextRemoveEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTextRemoveEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTextRemoveEvent(h unsafe.Pointer) *QAccessibleTextRemoveEvent { return bindQAccessibleTextRemoveEvent(uintptr(h)) }
func UnsafeNewQAccessibleTextRemoveEventOwned(h unsafe.Pointer) *QAccessibleTextRemoveEvent { return newownedQAccessibleTextRemoveEvent(uintptr(h)) }

func (this *QAccessibleTextRemoveEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleTextCursorEvent != nil {
		this.QAccessibleTextCursorEvent.UnsafeRelease()
	}
}

func NewQAccessibleTextRemoveEvent(obj *QObject, position int, text string) *QAccessibleTextRemoveEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextRemoveEvent_new, uintptr(obj.UnsafePointer()), uintptr(position), text_p, uintptr(len(text)))
	__o := newownedQAccessibleTextRemoveEvent(__r)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(text)
	return __o
}

func NewQAccessibleTextRemoveEvent2(iface *QAccessibleInterface, position int, text string) *QAccessibleTextRemoveEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextRemoveEvent_new2, uintptr(iface.UnsafePointer()), uintptr(position), text_p, uintptr(len(text)))
	__o := newownedQAccessibleTextRemoveEvent(__r)
	runtime.KeepAlive(iface)
	runtime.KeepAlive(text)
	return __o
}

func (this *QAccessibleTextRemoveEvent) TextRemoved() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextRemoveEvent_textRemoved, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTextRemoveEvent) ChangePosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextRemoveEvent_changePosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextRemoveEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTextRemoveEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
