package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneResizeEvent_chain uintptr
	addr_QGraphicsSceneResizeEvent_new uintptr
	addr_QGraphicsSceneResizeEvent_oldSize uintptr
	addr_QGraphicsSceneResizeEvent_setOldSize uintptr
	addr_QGraphicsSceneResizeEvent_newSize uintptr
	addr_QGraphicsSceneResizeEvent_setNewSize uintptr
	addr_QGraphicsSceneResizeEvent_registerEventType uintptr
	addr_QGraphicsSceneResizeEvent_registerEventType2 uintptr
	addr_QGraphicsSceneResizeEvent_setAccepted_default uintptr
	addr_QGraphicsSceneResizeEvent_clone_default uintptr
	pg_QGraphicsSceneResizeEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneResizeEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneResizeEvent_chain, "QGraphicsSceneResizeEvent_chain")
	symaddr(&addr_QGraphicsSceneResizeEvent_new, "QGraphicsSceneResizeEvent_new")
	symaddr(&addr_QGraphicsSceneResizeEvent_oldSize, "QGraphicsSceneResizeEvent_oldSize")
	symaddr(&addr_QGraphicsSceneResizeEvent_setOldSize, "QGraphicsSceneResizeEvent_setOldSize")
	symaddr(&addr_QGraphicsSceneResizeEvent_newSize, "QGraphicsSceneResizeEvent_newSize")
	symaddr(&addr_QGraphicsSceneResizeEvent_setNewSize, "QGraphicsSceneResizeEvent_setNewSize")
	symaddr(&addr_QGraphicsSceneResizeEvent_registerEventType, "QGraphicsSceneResizeEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneResizeEvent_registerEventType2, "QGraphicsSceneResizeEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneResizeEvent_setAccepted_default, "QGraphicsSceneResizeEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneResizeEvent_clone_default, "QGraphicsSceneResizeEvent_clone_default")
	reg(&pg_QGraphicsSceneResizeEvent__set_override, "QGraphicsSceneResizeEvent__set_override")
	reg(&pg_QGraphicsSceneResizeEvent_delete, "QGraphicsSceneResizeEvent_delete")
}

type QGraphicsSceneResizeEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneResizeEvent(h uintptr) *QGraphicsSceneResizeEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneResizeEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneResizeEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneResizeEvent(h uintptr) *QGraphicsSceneResizeEvent {
	o := bindQGraphicsSceneResizeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneResizeEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneResizeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneResizeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneResizeEvent(h unsafe.Pointer) *QGraphicsSceneResizeEvent { return bindQGraphicsSceneResizeEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneResizeEventOwned(h unsafe.Pointer) *QGraphicsSceneResizeEvent { return newownedQGraphicsSceneResizeEvent(uintptr(h)) }

func (this *QGraphicsSceneResizeEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneResizeEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneResizeEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneResizeEvent_virtIdx, QGraphicsSceneResizeEvent_enable)
	return __o
}

func (this *QGraphicsSceneResizeEvent) OldSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneResizeEvent_oldSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsSceneResizeEvent) SetOldSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsSceneResizeEvent_setOldSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsSceneResizeEvent) NewSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneResizeEvent_newSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsSceneResizeEvent) SetNewSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsSceneResizeEvent_setNewSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func QGraphicsSceneResizeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneResizeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneResizeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneResizeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneResizeEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneResizeEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneResizeEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneResizeEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneResizeEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneResizeEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneResizeEvent__set_override(h, idx) }

func (this *QGraphicsSceneResizeEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneResizeEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneResizeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneResizeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
