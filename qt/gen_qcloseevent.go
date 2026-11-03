package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QCloseEvent_base func(uintptr) uintptr
	addr_QCloseEvent_new uintptr
	addr_QCloseEvent_registerEventType uintptr
	addr_QCloseEvent_registerEventType2 uintptr
	addr_QCloseEvent_clone_default uintptr
	addr_QCloseEvent_setAccepted_default uintptr
	pg_QCloseEvent__set_override func(uintptr, int32)
	pg_QCloseEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QCloseEvent_base, "QCloseEvent_base")
	symaddr(&addr_QCloseEvent_new, "QCloseEvent_new")
	symaddr(&addr_QCloseEvent_registerEventType, "QCloseEvent_registerEventType")
	symaddr(&addr_QCloseEvent_registerEventType2, "QCloseEvent_registerEventType2")
	symaddr(&addr_QCloseEvent_clone_default, "QCloseEvent_clone_default")
	symaddr(&addr_QCloseEvent_setAccepted_default, "QCloseEvent_setAccepted_default")
	reg(&pg_QCloseEvent__set_override, "QCloseEvent__set_override")
	reg(&pg_QCloseEvent_delete, "QCloseEvent_delete")
}

type QCloseEvent struct {
	h uintptr
	*QEvent
}

func bindQCloseEvent(h uintptr) *QCloseEvent {
	if h == 0 {
		return nil
	}
	return &QCloseEvent{h: h, QEvent: bindQEvent(pg_QCloseEvent_base(h))}
}

func newownedQCloseEvent(h uintptr) *QCloseEvent {
	o := bindQCloseEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCloseEvent).Delete)
	}
	return o
}

func (this *QCloseEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCloseEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCloseEvent(h unsafe.Pointer) *QCloseEvent { return bindQCloseEvent(uintptr(h)) }
func UnsafeNewQCloseEventOwned(h unsafe.Pointer) *QCloseEvent { return newownedQCloseEvent(uintptr(h)) }

func (this *QCloseEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQCloseEvent() *QCloseEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCloseEvent_new, uintptr(__goid))
	__o := newownedQCloseEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCloseEvent_virtIdx, QCloseEvent_enable)
	return __o
}

func QCloseEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QCloseEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCloseEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QCloseEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QCloseEvent) CloneDefault() *QCloseEvent {
	__r, _, _ := purego.SyscallN(addr_QCloseEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQCloseEvent(__r)
}

func (this *QCloseEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QCloseEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func QCloseEvent_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case 6367896784077321574:
		return 1
	}
	return -1
}

func QCloseEvent_enable(h uintptr, idx int32) { pg_QCloseEvent__set_override(h, idx) }

func (this *QCloseEvent) OnClone(fn func() *QCloseEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QCloseEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QCloseEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCloseEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
