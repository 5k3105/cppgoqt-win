package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTextInsertEvent_chain uintptr
	addr_QAccessibleTextInsertEvent_new uintptr
	addr_QAccessibleTextInsertEvent_new2 uintptr
	addr_QAccessibleTextInsertEvent_textInserted uintptr
	addr_QAccessibleTextInsertEvent_changePosition uintptr
	pg_QAccessibleTextInsertEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTextInsertEvent_chain, "QAccessibleTextInsertEvent_chain")
	symaddr(&addr_QAccessibleTextInsertEvent_new, "QAccessibleTextInsertEvent_new")
	symaddr(&addr_QAccessibleTextInsertEvent_new2, "QAccessibleTextInsertEvent_new2")
	symaddr(&addr_QAccessibleTextInsertEvent_textInserted, "QAccessibleTextInsertEvent_textInserted")
	symaddr(&addr_QAccessibleTextInsertEvent_changePosition, "QAccessibleTextInsertEvent_changePosition")
	reg(&pg_QAccessibleTextInsertEvent_delete, "QAccessibleTextInsertEvent_delete")
}

type QAccessibleTextInsertEvent struct {
	h uintptr
	*QAccessibleTextCursorEvent
}

func bindQAccessibleTextInsertEvent(h uintptr) *QAccessibleTextInsertEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAccessibleTextInsertEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleTextInsertEvent{h: h, QAccessibleTextCursorEvent: &QAccessibleTextCursorEvent{h: __c[0], QAccessibleEvent: &QAccessibleEvent{h: __c[1]}}}
}

func newownedQAccessibleTextInsertEvent(h uintptr) *QAccessibleTextInsertEvent {
	o := bindQAccessibleTextInsertEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTextInsertEvent).Delete)
	}
	return o
}

func (this *QAccessibleTextInsertEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTextInsertEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTextInsertEvent(h unsafe.Pointer) *QAccessibleTextInsertEvent { return bindQAccessibleTextInsertEvent(uintptr(h)) }
func UnsafeNewQAccessibleTextInsertEventOwned(h unsafe.Pointer) *QAccessibleTextInsertEvent { return newownedQAccessibleTextInsertEvent(uintptr(h)) }

func (this *QAccessibleTextInsertEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleTextCursorEvent != nil {
		this.QAccessibleTextCursorEvent.UnsafeRelease()
	}
}

func NewQAccessibleTextInsertEvent(obj *QObject, position int, text string) *QAccessibleTextInsertEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInsertEvent_new, uintptr(obj.UnsafePointer()), uintptr(position), text_p, uintptr(len(text)))
	__o := newownedQAccessibleTextInsertEvent(__r)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(text)
	return __o
}

func NewQAccessibleTextInsertEvent2(iface *QAccessibleInterface, position int, text string) *QAccessibleTextInsertEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInsertEvent_new2, uintptr(iface.UnsafePointer()), uintptr(position), text_p, uintptr(len(text)))
	__o := newownedQAccessibleTextInsertEvent(__r)
	runtime.KeepAlive(iface)
	runtime.KeepAlive(text)
	return __o
}

func (this *QAccessibleTextInsertEvent) TextInserted() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInsertEvent_textInserted, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTextInsertEvent) ChangePosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextInsertEvent_changePosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextInsertEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTextInsertEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
