package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDragLeaveEvent_base func(uintptr) uintptr
	addr_QDragLeaveEvent_new uintptr
	addr_QDragLeaveEvent_registerEventType uintptr
	addr_QDragLeaveEvent_registerEventType2 uintptr
	addr_QDragLeaveEvent_clone_default uintptr
	addr_QDragLeaveEvent_setAccepted_default uintptr
	pg_QDragLeaveEvent__set_override func(uintptr, int32)
	pg_QDragLeaveEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QDragLeaveEvent_base, "QDragLeaveEvent_base")
	symaddr(&addr_QDragLeaveEvent_new, "QDragLeaveEvent_new")
	symaddr(&addr_QDragLeaveEvent_registerEventType, "QDragLeaveEvent_registerEventType")
	symaddr(&addr_QDragLeaveEvent_registerEventType2, "QDragLeaveEvent_registerEventType2")
	symaddr(&addr_QDragLeaveEvent_clone_default, "QDragLeaveEvent_clone_default")
	symaddr(&addr_QDragLeaveEvent_setAccepted_default, "QDragLeaveEvent_setAccepted_default")
	reg(&pg_QDragLeaveEvent__set_override, "QDragLeaveEvent__set_override")
	reg(&pg_QDragLeaveEvent_delete, "QDragLeaveEvent_delete")
}

type QDragLeaveEvent struct {
	h uintptr
	*QEvent
}

func bindQDragLeaveEvent(h uintptr) *QDragLeaveEvent {
	if h == 0 {
		return nil
	}
	return &QDragLeaveEvent{h: h, QEvent: bindQEvent(pg_QDragLeaveEvent_base(h))}
}

func newownedQDragLeaveEvent(h uintptr) *QDragLeaveEvent {
	o := bindQDragLeaveEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDragLeaveEvent).Delete)
	}
	return o
}

func (this *QDragLeaveEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDragLeaveEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDragLeaveEvent(h unsafe.Pointer) *QDragLeaveEvent { return bindQDragLeaveEvent(uintptr(h)) }
func UnsafeNewQDragLeaveEventOwned(h unsafe.Pointer) *QDragLeaveEvent { return newownedQDragLeaveEvent(uintptr(h)) }

func (this *QDragLeaveEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQDragLeaveEvent() *QDragLeaveEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDragLeaveEvent_new, uintptr(__goid))
	__o := newownedQDragLeaveEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDragLeaveEvent_virtIdx, QDragLeaveEvent_enable)
	return __o
}

func QDragLeaveEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QDragLeaveEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QDragLeaveEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QDragLeaveEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QDragLeaveEvent) CloneDefault() *QDragLeaveEvent {
	__r, _, _ := purego.SyscallN(addr_QDragLeaveEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQDragLeaveEvent(__r)
}

func (this *QDragLeaveEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QDragLeaveEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func QDragLeaveEvent_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case 6367896784077321574:
		return 1
	}
	return -1
}

func QDragLeaveEvent_enable(h uintptr, idx int32) { pg_QDragLeaveEvent__set_override(h, idx) }

func (this *QDragLeaveEvent) OnClone(fn func() *QDragLeaveEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QDragLeaveEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QDragLeaveEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDragLeaveEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
