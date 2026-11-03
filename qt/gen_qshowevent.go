package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QShowEvent_base func(uintptr) uintptr
	addr_QShowEvent_new uintptr
	addr_QShowEvent_registerEventType uintptr
	addr_QShowEvent_registerEventType2 uintptr
	addr_QShowEvent_clone_default uintptr
	addr_QShowEvent_setAccepted_default uintptr
	pg_QShowEvent__set_override func(uintptr, int32)
	pg_QShowEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QShowEvent_base, "QShowEvent_base")
	symaddr(&addr_QShowEvent_new, "QShowEvent_new")
	symaddr(&addr_QShowEvent_registerEventType, "QShowEvent_registerEventType")
	symaddr(&addr_QShowEvent_registerEventType2, "QShowEvent_registerEventType2")
	symaddr(&addr_QShowEvent_clone_default, "QShowEvent_clone_default")
	symaddr(&addr_QShowEvent_setAccepted_default, "QShowEvent_setAccepted_default")
	reg(&pg_QShowEvent__set_override, "QShowEvent__set_override")
	reg(&pg_QShowEvent_delete, "QShowEvent_delete")
}

type QShowEvent struct {
	h uintptr
	*QEvent
}

func bindQShowEvent(h uintptr) *QShowEvent {
	if h == 0 {
		return nil
	}
	return &QShowEvent{h: h, QEvent: bindQEvent(pg_QShowEvent_base(h))}
}

func newownedQShowEvent(h uintptr) *QShowEvent {
	o := bindQShowEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QShowEvent).Delete)
	}
	return o
}

func (this *QShowEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QShowEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQShowEvent(h unsafe.Pointer) *QShowEvent { return bindQShowEvent(uintptr(h)) }
func UnsafeNewQShowEventOwned(h unsafe.Pointer) *QShowEvent { return newownedQShowEvent(uintptr(h)) }

func (this *QShowEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQShowEvent() *QShowEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QShowEvent_new, uintptr(__goid))
	__o := newownedQShowEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShowEvent_virtIdx, QShowEvent_enable)
	return __o
}

func QShowEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QShowEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QShowEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QShowEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QShowEvent) CloneDefault() *QShowEvent {
	__r, _, _ := purego.SyscallN(addr_QShowEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQShowEvent(__r)
}

func (this *QShowEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QShowEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func QShowEvent_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case 6367896784077321574:
		return 1
	}
	return -1
}

func QShowEvent_enable(h uintptr, idx int32) { pg_QShowEvent__set_override(h, idx) }

func (this *QShowEvent) OnClone(fn func() *QShowEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QShowEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QShowEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QShowEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
