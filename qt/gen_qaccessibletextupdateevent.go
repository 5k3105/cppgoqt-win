package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTextUpdateEvent_chain uintptr
	addr_QAccessibleTextUpdateEvent_new uintptr
	addr_QAccessibleTextUpdateEvent_new2 uintptr
	addr_QAccessibleTextUpdateEvent_textRemoved uintptr
	addr_QAccessibleTextUpdateEvent_textInserted uintptr
	addr_QAccessibleTextUpdateEvent_changePosition uintptr
	pg_QAccessibleTextUpdateEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTextUpdateEvent_chain, "QAccessibleTextUpdateEvent_chain")
	symaddr(&addr_QAccessibleTextUpdateEvent_new, "QAccessibleTextUpdateEvent_new")
	symaddr(&addr_QAccessibleTextUpdateEvent_new2, "QAccessibleTextUpdateEvent_new2")
	symaddr(&addr_QAccessibleTextUpdateEvent_textRemoved, "QAccessibleTextUpdateEvent_textRemoved")
	symaddr(&addr_QAccessibleTextUpdateEvent_textInserted, "QAccessibleTextUpdateEvent_textInserted")
	symaddr(&addr_QAccessibleTextUpdateEvent_changePosition, "QAccessibleTextUpdateEvent_changePosition")
	reg(&pg_QAccessibleTextUpdateEvent_delete, "QAccessibleTextUpdateEvent_delete")
}

type QAccessibleTextUpdateEvent struct {
	h uintptr
	*QAccessibleTextCursorEvent
}

func bindQAccessibleTextUpdateEvent(h uintptr) *QAccessibleTextUpdateEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAccessibleTextUpdateEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleTextUpdateEvent{h: h, QAccessibleTextCursorEvent: &QAccessibleTextCursorEvent{h: __c[0], QAccessibleEvent: &QAccessibleEvent{h: __c[1]}}}
}

func newownedQAccessibleTextUpdateEvent(h uintptr) *QAccessibleTextUpdateEvent {
	o := bindQAccessibleTextUpdateEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTextUpdateEvent).Delete)
	}
	return o
}

func (this *QAccessibleTextUpdateEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTextUpdateEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTextUpdateEvent(h unsafe.Pointer) *QAccessibleTextUpdateEvent { return bindQAccessibleTextUpdateEvent(uintptr(h)) }
func UnsafeNewQAccessibleTextUpdateEventOwned(h unsafe.Pointer) *QAccessibleTextUpdateEvent { return newownedQAccessibleTextUpdateEvent(uintptr(h)) }

func (this *QAccessibleTextUpdateEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleTextCursorEvent != nil {
		this.QAccessibleTextCursorEvent.UnsafeRelease()
	}
}

func NewQAccessibleTextUpdateEvent(obj *QObject, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	var oldText_p uintptr
	if len(oldText) > 0 {
		oldText_p = uintptr(unsafe.Pointer(unsafe.StringData(oldText)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextUpdateEvent_new, uintptr(obj.UnsafePointer()), uintptr(position), oldText_p, uintptr(len(oldText)), text_p, uintptr(len(text)))
	__o := newownedQAccessibleTextUpdateEvent(__r)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(oldText)
	runtime.KeepAlive(text)
	return __o
}

func NewQAccessibleTextUpdateEvent2(iface *QAccessibleInterface, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	var oldText_p uintptr
	if len(oldText) > 0 {
		oldText_p = uintptr(unsafe.Pointer(unsafe.StringData(oldText)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextUpdateEvent_new2, uintptr(iface.UnsafePointer()), uintptr(position), oldText_p, uintptr(len(oldText)), text_p, uintptr(len(text)))
	__o := newownedQAccessibleTextUpdateEvent(__r)
	runtime.KeepAlive(iface)
	runtime.KeepAlive(oldText)
	runtime.KeepAlive(text)
	return __o
}

func (this *QAccessibleTextUpdateEvent) TextRemoved() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextUpdateEvent_textRemoved, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTextUpdateEvent) TextInserted() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextUpdateEvent_textInserted, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTextUpdateEvent) ChangePosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextUpdateEvent_changePosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextUpdateEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTextUpdateEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
