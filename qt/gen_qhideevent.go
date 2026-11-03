package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QHideEvent_base func(uintptr) uintptr
	addr_QHideEvent_new uintptr
	addr_QHideEvent_registerEventType uintptr
	addr_QHideEvent_registerEventType2 uintptr
	addr_QHideEvent_clone_default uintptr
	addr_QHideEvent_setAccepted_default uintptr
	pg_QHideEvent__set_override func(uintptr, int32)
	pg_QHideEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QHideEvent_base, "QHideEvent_base")
	symaddr(&addr_QHideEvent_new, "QHideEvent_new")
	symaddr(&addr_QHideEvent_registerEventType, "QHideEvent_registerEventType")
	symaddr(&addr_QHideEvent_registerEventType2, "QHideEvent_registerEventType2")
	symaddr(&addr_QHideEvent_clone_default, "QHideEvent_clone_default")
	symaddr(&addr_QHideEvent_setAccepted_default, "QHideEvent_setAccepted_default")
	reg(&pg_QHideEvent__set_override, "QHideEvent__set_override")
	reg(&pg_QHideEvent_delete, "QHideEvent_delete")
}

type QHideEvent struct {
	h uintptr
	*QEvent
}

func bindQHideEvent(h uintptr) *QHideEvent {
	if h == 0 {
		return nil
	}
	return &QHideEvent{h: h, QEvent: bindQEvent(pg_QHideEvent_base(h))}
}

func newownedQHideEvent(h uintptr) *QHideEvent {
	o := bindQHideEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QHideEvent).Delete)
	}
	return o
}

func (this *QHideEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QHideEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQHideEvent(h unsafe.Pointer) *QHideEvent { return bindQHideEvent(uintptr(h)) }
func UnsafeNewQHideEventOwned(h unsafe.Pointer) *QHideEvent { return newownedQHideEvent(uintptr(h)) }

func (this *QHideEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQHideEvent() *QHideEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QHideEvent_new, uintptr(__goid))
	__o := newownedQHideEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QHideEvent_virtIdx, QHideEvent_enable)
	return __o
}

func QHideEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QHideEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QHideEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QHideEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QHideEvent) CloneDefault() *QHideEvent {
	__r, _, _ := purego.SyscallN(addr_QHideEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQHideEvent(__r)
}

func (this *QHideEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QHideEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func QHideEvent_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case 6367896784077321574:
		return 1
	}
	return -1
}

func QHideEvent_enable(h uintptr, idx int32) { pg_QHideEvent__set_override(h, idx) }

func (this *QHideEvent) OnClone(fn func() *QHideEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QHideEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QHideEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QHideEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
