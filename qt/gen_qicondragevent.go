package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QIconDragEvent_base func(uintptr) uintptr
	addr_QIconDragEvent_new uintptr
	addr_QIconDragEvent_registerEventType uintptr
	addr_QIconDragEvent_registerEventType2 uintptr
	addr_QIconDragEvent_clone_default uintptr
	addr_QIconDragEvent_setAccepted_default uintptr
	pg_QIconDragEvent__set_override func(uintptr, int32)
	pg_QIconDragEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QIconDragEvent_base, "QIconDragEvent_base")
	symaddr(&addr_QIconDragEvent_new, "QIconDragEvent_new")
	symaddr(&addr_QIconDragEvent_registerEventType, "QIconDragEvent_registerEventType")
	symaddr(&addr_QIconDragEvent_registerEventType2, "QIconDragEvent_registerEventType2")
	symaddr(&addr_QIconDragEvent_clone_default, "QIconDragEvent_clone_default")
	symaddr(&addr_QIconDragEvent_setAccepted_default, "QIconDragEvent_setAccepted_default")
	reg(&pg_QIconDragEvent__set_override, "QIconDragEvent__set_override")
	reg(&pg_QIconDragEvent_delete, "QIconDragEvent_delete")
}

type QIconDragEvent struct {
	h uintptr
	*QEvent
}

func bindQIconDragEvent(h uintptr) *QIconDragEvent {
	if h == 0 {
		return nil
	}
	return &QIconDragEvent{h: h, QEvent: bindQEvent(pg_QIconDragEvent_base(h))}
}

func newownedQIconDragEvent(h uintptr) *QIconDragEvent {
	o := bindQIconDragEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QIconDragEvent).Delete)
	}
	return o
}

func (this *QIconDragEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIconDragEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIconDragEvent(h unsafe.Pointer) *QIconDragEvent { return bindQIconDragEvent(uintptr(h)) }
func UnsafeNewQIconDragEventOwned(h unsafe.Pointer) *QIconDragEvent { return newownedQIconDragEvent(uintptr(h)) }

func (this *QIconDragEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQIconDragEvent() *QIconDragEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QIconDragEvent_new, uintptr(__goid))
	__o := newownedQIconDragEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QIconDragEvent_virtIdx, QIconDragEvent_enable)
	return __o
}

func QIconDragEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QIconDragEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QIconDragEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QIconDragEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QIconDragEvent) CloneDefault() *QIconDragEvent {
	__r, _, _ := purego.SyscallN(addr_QIconDragEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQIconDragEvent(__r)
}

func (this *QIconDragEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QIconDragEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func QIconDragEvent_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case 6367896784077321574:
		return 1
	}
	return -1
}

func QIconDragEvent_enable(h uintptr, idx int32) { pg_QIconDragEvent__set_override(h, idx) }

func (this *QIconDragEvent) OnClone(fn func() *QIconDragEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QIconDragEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QIconDragEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIconDragEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
