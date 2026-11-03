package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTimerEvent_base func(uintptr) uintptr
	addr_QTimerEvent_new uintptr
	addr_QTimerEvent_new2 uintptr
	addr_QTimerEvent_timerId uintptr
	addr_QTimerEvent_id uintptr
	addr_QTimerEvent_matches uintptr
	addr_QTimerEvent_registerEventType uintptr
	addr_QTimerEvent_registerEventType2 uintptr
	pg_QTimerEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QTimerEvent_base, "QTimerEvent_base")
	symaddr(&addr_QTimerEvent_new, "QTimerEvent_new")
	symaddr(&addr_QTimerEvent_new2, "QTimerEvent_new2")
	symaddr(&addr_QTimerEvent_timerId, "QTimerEvent_timerId")
	symaddr(&addr_QTimerEvent_id, "QTimerEvent_id")
	symaddr(&addr_QTimerEvent_matches, "QTimerEvent_matches")
	symaddr(&addr_QTimerEvent_registerEventType, "QTimerEvent_registerEventType")
	symaddr(&addr_QTimerEvent_registerEventType2, "QTimerEvent_registerEventType2")
	reg(&pg_QTimerEvent_delete, "QTimerEvent_delete")
}

type QTimerEvent struct {
	h uintptr
	*QEvent
}

func bindQTimerEvent(h uintptr) *QTimerEvent {
	if h == 0 {
		return nil
	}
	return &QTimerEvent{h: h, QEvent: bindQEvent(pg_QTimerEvent_base(h))}
}

func newownedQTimerEvent(h uintptr) *QTimerEvent {
	o := bindQTimerEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTimerEvent).Delete)
	}
	return o
}

func (this *QTimerEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTimerEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTimerEvent(h unsafe.Pointer) *QTimerEvent { return bindQTimerEvent(uintptr(h)) }
func UnsafeNewQTimerEventOwned(h unsafe.Pointer) *QTimerEvent { return newownedQTimerEvent(uintptr(h)) }

func (this *QTimerEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQTimerEvent(timerId int) *QTimerEvent {
	__r, _, _ := purego.SyscallN(addr_QTimerEvent_new, uintptr(timerId))
	__o := newownedQTimerEvent(__r)
	return __o
}

func NewQTimerEvent2(timerId TimerId) *QTimerEvent {
	__r, _, _ := purego.SyscallN(addr_QTimerEvent_new2, uintptr(timerId))
	__o := newownedQTimerEvent(__r)
	return __o
}

func (this *QTimerEvent) TimerId() int {
	__r0, _, _ := purego.SyscallN(addr_QTimerEvent_timerId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimerEvent) Id() TimerId {
	__r0, _, _ := purego.SyscallN(addr_QTimerEvent_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerId(__r)
}

func (this *QTimerEvent) Matches(timer *QBasicTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimerEvent_matches, this.h, uintptr(timer.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timer)
	return __r != 0
}

func QTimerEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QTimerEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QTimerEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QTimerEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QTimerEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTimerEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
