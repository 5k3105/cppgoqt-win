package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QWindowStateChangeEvent_base func(uintptr) uintptr
	addr_QWindowStateChangeEvent_new uintptr
	addr_QWindowStateChangeEvent_new2 uintptr
	addr_QWindowStateChangeEvent_oldState uintptr
	addr_QWindowStateChangeEvent_isOverride uintptr
	addr_QWindowStateChangeEvent_registerEventType uintptr
	addr_QWindowStateChangeEvent_registerEventType2 uintptr
	pg_QWindowStateChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QWindowStateChangeEvent_base, "QWindowStateChangeEvent_base")
	symaddr(&addr_QWindowStateChangeEvent_new, "QWindowStateChangeEvent_new")
	symaddr(&addr_QWindowStateChangeEvent_new2, "QWindowStateChangeEvent_new2")
	symaddr(&addr_QWindowStateChangeEvent_oldState, "QWindowStateChangeEvent_oldState")
	symaddr(&addr_QWindowStateChangeEvent_isOverride, "QWindowStateChangeEvent_isOverride")
	symaddr(&addr_QWindowStateChangeEvent_registerEventType, "QWindowStateChangeEvent_registerEventType")
	symaddr(&addr_QWindowStateChangeEvent_registerEventType2, "QWindowStateChangeEvent_registerEventType2")
	reg(&pg_QWindowStateChangeEvent_delete, "QWindowStateChangeEvent_delete")
}

type QWindowStateChangeEvent struct {
	h uintptr
	*QEvent
}

func bindQWindowStateChangeEvent(h uintptr) *QWindowStateChangeEvent {
	if h == 0 {
		return nil
	}
	return &QWindowStateChangeEvent{h: h, QEvent: bindQEvent(pg_QWindowStateChangeEvent_base(h))}
}

func newownedQWindowStateChangeEvent(h uintptr) *QWindowStateChangeEvent {
	o := bindQWindowStateChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWindowStateChangeEvent).Delete)
	}
	return o
}

func (this *QWindowStateChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWindowStateChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWindowStateChangeEvent(h unsafe.Pointer) *QWindowStateChangeEvent { return bindQWindowStateChangeEvent(uintptr(h)) }
func UnsafeNewQWindowStateChangeEventOwned(h unsafe.Pointer) *QWindowStateChangeEvent { return newownedQWindowStateChangeEvent(uintptr(h)) }

func (this *QWindowStateChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQWindowStateChangeEvent(oldState WindowState) *QWindowStateChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QWindowStateChangeEvent_new, uintptr(oldState))
	__o := newownedQWindowStateChangeEvent(__r)
	return __o
}

func NewQWindowStateChangeEvent2(oldState WindowState, isOverride bool) *QWindowStateChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QWindowStateChangeEvent_new2, uintptr(oldState), uintptr(b2i(isOverride)))
	__o := newownedQWindowStateChangeEvent(__r)
	return __o
}

func (this *QWindowStateChangeEvent) OldState() WindowState {
	__r0, _, _ := purego.SyscallN(addr_QWindowStateChangeEvent_oldState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowState(__r)
}

func (this *QWindowStateChangeEvent) IsOverride() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindowStateChangeEvent_isOverride, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QWindowStateChangeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QWindowStateChangeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QWindowStateChangeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QWindowStateChangeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QWindowStateChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWindowStateChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
