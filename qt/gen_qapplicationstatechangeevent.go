package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QApplicationStateChangeEvent_base func(uintptr) uintptr
	addr_QApplicationStateChangeEvent_new uintptr
	addr_QApplicationStateChangeEvent_applicationState uintptr
	addr_QApplicationStateChangeEvent_registerEventType uintptr
	addr_QApplicationStateChangeEvent_registerEventType2 uintptr
	pg_QApplicationStateChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QApplicationStateChangeEvent_base, "QApplicationStateChangeEvent_base")
	symaddr(&addr_QApplicationStateChangeEvent_new, "QApplicationStateChangeEvent_new")
	symaddr(&addr_QApplicationStateChangeEvent_applicationState, "QApplicationStateChangeEvent_applicationState")
	symaddr(&addr_QApplicationStateChangeEvent_registerEventType, "QApplicationStateChangeEvent_registerEventType")
	symaddr(&addr_QApplicationStateChangeEvent_registerEventType2, "QApplicationStateChangeEvent_registerEventType2")
	reg(&pg_QApplicationStateChangeEvent_delete, "QApplicationStateChangeEvent_delete")
}

type QApplicationStateChangeEvent struct {
	h uintptr
	*QEvent
}

func bindQApplicationStateChangeEvent(h uintptr) *QApplicationStateChangeEvent {
	if h == 0 {
		return nil
	}
	return &QApplicationStateChangeEvent{h: h, QEvent: bindQEvent(pg_QApplicationStateChangeEvent_base(h))}
}

func newownedQApplicationStateChangeEvent(h uintptr) *QApplicationStateChangeEvent {
	o := bindQApplicationStateChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QApplicationStateChangeEvent).Delete)
	}
	return o
}

func (this *QApplicationStateChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QApplicationStateChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQApplicationStateChangeEvent(h unsafe.Pointer) *QApplicationStateChangeEvent { return bindQApplicationStateChangeEvent(uintptr(h)) }
func UnsafeNewQApplicationStateChangeEventOwned(h unsafe.Pointer) *QApplicationStateChangeEvent { return newownedQApplicationStateChangeEvent(uintptr(h)) }

func (this *QApplicationStateChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQApplicationStateChangeEvent(state ApplicationState) *QApplicationStateChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QApplicationStateChangeEvent_new, uintptr(state))
	__o := newownedQApplicationStateChangeEvent(__r)
	return __o
}

func (this *QApplicationStateChangeEvent) ApplicationState() ApplicationState {
	__r0, _, _ := purego.SyscallN(addr_QApplicationStateChangeEvent_applicationState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ApplicationState(__r)
}

func QApplicationStateChangeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QApplicationStateChangeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplicationStateChangeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QApplicationStateChangeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QApplicationStateChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QApplicationStateChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
