package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QFocusEvent_base func(uintptr) uintptr
	addr_QFocusEvent_new uintptr
	addr_QFocusEvent_new2 uintptr
	addr_QFocusEvent_gotFocus uintptr
	addr_QFocusEvent_lostFocus uintptr
	addr_QFocusEvent_reason uintptr
	addr_QFocusEvent_registerEventType uintptr
	addr_QFocusEvent_registerEventType2 uintptr
	pg_QFocusEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QFocusEvent_base, "QFocusEvent_base")
	symaddr(&addr_QFocusEvent_new, "QFocusEvent_new")
	symaddr(&addr_QFocusEvent_new2, "QFocusEvent_new2")
	symaddr(&addr_QFocusEvent_gotFocus, "QFocusEvent_gotFocus")
	symaddr(&addr_QFocusEvent_lostFocus, "QFocusEvent_lostFocus")
	symaddr(&addr_QFocusEvent_reason, "QFocusEvent_reason")
	symaddr(&addr_QFocusEvent_registerEventType, "QFocusEvent_registerEventType")
	symaddr(&addr_QFocusEvent_registerEventType2, "QFocusEvent_registerEventType2")
	reg(&pg_QFocusEvent_delete, "QFocusEvent_delete")
}

type QFocusEvent struct {
	h uintptr
	*QEvent
}

func bindQFocusEvent(h uintptr) *QFocusEvent {
	if h == 0 {
		return nil
	}
	return &QFocusEvent{h: h, QEvent: bindQEvent(pg_QFocusEvent_base(h))}
}

func newownedQFocusEvent(h uintptr) *QFocusEvent {
	o := bindQFocusEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFocusEvent).Delete)
	}
	return o
}

func (this *QFocusEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFocusEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFocusEvent(h unsafe.Pointer) *QFocusEvent { return bindQFocusEvent(uintptr(h)) }
func UnsafeNewQFocusEventOwned(h unsafe.Pointer) *QFocusEvent { return newownedQFocusEvent(uintptr(h)) }

func (this *QFocusEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQFocusEvent(type_ QEvent__Type) *QFocusEvent {
	__r, _, _ := purego.SyscallN(addr_QFocusEvent_new, uintptr(type_))
	__o := newownedQFocusEvent(__r)
	return __o
}

func NewQFocusEvent2(type_ QEvent__Type, reason FocusReason) *QFocusEvent {
	__r, _, _ := purego.SyscallN(addr_QFocusEvent_new2, uintptr(type_), uintptr(reason))
	__o := newownedQFocusEvent(__r)
	return __o
}

func (this *QFocusEvent) GotFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusEvent_gotFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFocusEvent) LostFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusEvent_lostFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFocusEvent) Reason() FocusReason {
	__r0, _, _ := purego.SyscallN(addr_QFocusEvent_reason, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return FocusReason(__r)
}

func QFocusEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QFocusEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QFocusEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QFocusEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QFocusEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFocusEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
