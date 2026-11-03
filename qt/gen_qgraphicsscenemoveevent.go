package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneMoveEvent_chain uintptr
	addr_QGraphicsSceneMoveEvent_new uintptr
	addr_QGraphicsSceneMoveEvent_oldPos uintptr
	addr_QGraphicsSceneMoveEvent_setOldPos uintptr
	addr_QGraphicsSceneMoveEvent_newPos uintptr
	addr_QGraphicsSceneMoveEvent_setNewPos uintptr
	addr_QGraphicsSceneMoveEvent_registerEventType uintptr
	addr_QGraphicsSceneMoveEvent_registerEventType2 uintptr
	addr_QGraphicsSceneMoveEvent_setAccepted_default uintptr
	addr_QGraphicsSceneMoveEvent_clone_default uintptr
	pg_QGraphicsSceneMoveEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneMoveEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneMoveEvent_chain, "QGraphicsSceneMoveEvent_chain")
	symaddr(&addr_QGraphicsSceneMoveEvent_new, "QGraphicsSceneMoveEvent_new")
	symaddr(&addr_QGraphicsSceneMoveEvent_oldPos, "QGraphicsSceneMoveEvent_oldPos")
	symaddr(&addr_QGraphicsSceneMoveEvent_setOldPos, "QGraphicsSceneMoveEvent_setOldPos")
	symaddr(&addr_QGraphicsSceneMoveEvent_newPos, "QGraphicsSceneMoveEvent_newPos")
	symaddr(&addr_QGraphicsSceneMoveEvent_setNewPos, "QGraphicsSceneMoveEvent_setNewPos")
	symaddr(&addr_QGraphicsSceneMoveEvent_registerEventType, "QGraphicsSceneMoveEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneMoveEvent_registerEventType2, "QGraphicsSceneMoveEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneMoveEvent_setAccepted_default, "QGraphicsSceneMoveEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneMoveEvent_clone_default, "QGraphicsSceneMoveEvent_clone_default")
	reg(&pg_QGraphicsSceneMoveEvent__set_override, "QGraphicsSceneMoveEvent__set_override")
	reg(&pg_QGraphicsSceneMoveEvent_delete, "QGraphicsSceneMoveEvent_delete")
}

type QGraphicsSceneMoveEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneMoveEvent(h uintptr) *QGraphicsSceneMoveEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneMoveEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneMoveEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneMoveEvent(h uintptr) *QGraphicsSceneMoveEvent {
	o := bindQGraphicsSceneMoveEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneMoveEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneMoveEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneMoveEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneMoveEvent(h unsafe.Pointer) *QGraphicsSceneMoveEvent { return bindQGraphicsSceneMoveEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneMoveEventOwned(h unsafe.Pointer) *QGraphicsSceneMoveEvent { return newownedQGraphicsSceneMoveEvent(uintptr(h)) }

func (this *QGraphicsSceneMoveEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMoveEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneMoveEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneMoveEvent_virtIdx, QGraphicsSceneMoveEvent_enable)
	return __o
}

func (this *QGraphicsSceneMoveEvent) OldPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMoveEvent_oldPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMoveEvent) SetOldPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMoveEvent_setOldPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMoveEvent) NewPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMoveEvent_newPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMoveEvent) SetNewPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMoveEvent_setNewPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func QGraphicsSceneMoveEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMoveEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneMoveEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMoveEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneMoveEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneMoveEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneMoveEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMoveEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneMoveEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneMoveEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneMoveEvent__set_override(h, idx) }

func (this *QGraphicsSceneMoveEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneMoveEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneMoveEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneMoveEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
