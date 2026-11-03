package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTextSelectionEvent_chain uintptr
	addr_QAccessibleTextSelectionEvent_new uintptr
	addr_QAccessibleTextSelectionEvent_new2 uintptr
	addr_QAccessibleTextSelectionEvent_setSelection uintptr
	addr_QAccessibleTextSelectionEvent_selectionStart uintptr
	addr_QAccessibleTextSelectionEvent_selectionEnd uintptr
	pg_QAccessibleTextSelectionEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTextSelectionEvent_chain, "QAccessibleTextSelectionEvent_chain")
	symaddr(&addr_QAccessibleTextSelectionEvent_new, "QAccessibleTextSelectionEvent_new")
	symaddr(&addr_QAccessibleTextSelectionEvent_new2, "QAccessibleTextSelectionEvent_new2")
	symaddr(&addr_QAccessibleTextSelectionEvent_setSelection, "QAccessibleTextSelectionEvent_setSelection")
	symaddr(&addr_QAccessibleTextSelectionEvent_selectionStart, "QAccessibleTextSelectionEvent_selectionStart")
	symaddr(&addr_QAccessibleTextSelectionEvent_selectionEnd, "QAccessibleTextSelectionEvent_selectionEnd")
	reg(&pg_QAccessibleTextSelectionEvent_delete, "QAccessibleTextSelectionEvent_delete")
}

type QAccessibleTextSelectionEvent struct {
	h uintptr
	*QAccessibleTextCursorEvent
}

func bindQAccessibleTextSelectionEvent(h uintptr) *QAccessibleTextSelectionEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAccessibleTextSelectionEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleTextSelectionEvent{h: h, QAccessibleTextCursorEvent: &QAccessibleTextCursorEvent{h: __c[0], QAccessibleEvent: &QAccessibleEvent{h: __c[1]}}}
}

func newownedQAccessibleTextSelectionEvent(h uintptr) *QAccessibleTextSelectionEvent {
	o := bindQAccessibleTextSelectionEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTextSelectionEvent).Delete)
	}
	return o
}

func (this *QAccessibleTextSelectionEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTextSelectionEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTextSelectionEvent(h unsafe.Pointer) *QAccessibleTextSelectionEvent { return bindQAccessibleTextSelectionEvent(uintptr(h)) }
func UnsafeNewQAccessibleTextSelectionEventOwned(h unsafe.Pointer) *QAccessibleTextSelectionEvent { return newownedQAccessibleTextSelectionEvent(uintptr(h)) }

func (this *QAccessibleTextSelectionEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleTextCursorEvent != nil {
		this.QAccessibleTextCursorEvent.UnsafeRelease()
	}
}

func NewQAccessibleTextSelectionEvent(obj *QObject, start int, end int) *QAccessibleTextSelectionEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextSelectionEvent_new, uintptr(obj.UnsafePointer()), uintptr(start), uintptr(end))
	__o := newownedQAccessibleTextSelectionEvent(__r)
	runtime.KeepAlive(obj)
	return __o
}

func NewQAccessibleTextSelectionEvent2(iface *QAccessibleInterface, start int, end int) *QAccessibleTextSelectionEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextSelectionEvent_new2, uintptr(iface.UnsafePointer()), uintptr(start), uintptr(end))
	__o := newownedQAccessibleTextSelectionEvent(__r)
	runtime.KeepAlive(iface)
	return __o
}

func (this *QAccessibleTextSelectionEvent) SetSelection(start int, end int) {
	purego.SyscallN(addr_QAccessibleTextSelectionEvent_setSelection, this.h, uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextSelectionEvent) SelectionStart() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextSelectionEvent_selectionStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextSelectionEvent) SelectionEnd() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextSelectionEvent_selectionEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextSelectionEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTextSelectionEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
