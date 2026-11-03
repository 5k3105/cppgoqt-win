package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneHelpEvent_chain uintptr
	addr_QGraphicsSceneHelpEvent_new uintptr
	addr_QGraphicsSceneHelpEvent_new2 uintptr
	addr_QGraphicsSceneHelpEvent_scenePos uintptr
	addr_QGraphicsSceneHelpEvent_setScenePos uintptr
	addr_QGraphicsSceneHelpEvent_screenPos uintptr
	addr_QGraphicsSceneHelpEvent_setScreenPos uintptr
	addr_QGraphicsSceneHelpEvent_registerEventType uintptr
	addr_QGraphicsSceneHelpEvent_registerEventType2 uintptr
	addr_QGraphicsSceneHelpEvent_setAccepted_default uintptr
	addr_QGraphicsSceneHelpEvent_clone_default uintptr
	pg_QGraphicsSceneHelpEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneHelpEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneHelpEvent_chain, "QGraphicsSceneHelpEvent_chain")
	symaddr(&addr_QGraphicsSceneHelpEvent_new, "QGraphicsSceneHelpEvent_new")
	symaddr(&addr_QGraphicsSceneHelpEvent_new2, "QGraphicsSceneHelpEvent_new2")
	symaddr(&addr_QGraphicsSceneHelpEvent_scenePos, "QGraphicsSceneHelpEvent_scenePos")
	symaddr(&addr_QGraphicsSceneHelpEvent_setScenePos, "QGraphicsSceneHelpEvent_setScenePos")
	symaddr(&addr_QGraphicsSceneHelpEvent_screenPos, "QGraphicsSceneHelpEvent_screenPos")
	symaddr(&addr_QGraphicsSceneHelpEvent_setScreenPos, "QGraphicsSceneHelpEvent_setScreenPos")
	symaddr(&addr_QGraphicsSceneHelpEvent_registerEventType, "QGraphicsSceneHelpEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneHelpEvent_registerEventType2, "QGraphicsSceneHelpEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneHelpEvent_setAccepted_default, "QGraphicsSceneHelpEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneHelpEvent_clone_default, "QGraphicsSceneHelpEvent_clone_default")
	reg(&pg_QGraphicsSceneHelpEvent__set_override, "QGraphicsSceneHelpEvent__set_override")
	reg(&pg_QGraphicsSceneHelpEvent_delete, "QGraphicsSceneHelpEvent_delete")
}

type QGraphicsSceneHelpEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneHelpEvent(h uintptr) *QGraphicsSceneHelpEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneHelpEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneHelpEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneHelpEvent(h uintptr) *QGraphicsSceneHelpEvent {
	o := bindQGraphicsSceneHelpEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneHelpEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneHelpEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneHelpEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneHelpEvent(h unsafe.Pointer) *QGraphicsSceneHelpEvent { return bindQGraphicsSceneHelpEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneHelpEventOwned(h unsafe.Pointer) *QGraphicsSceneHelpEvent { return newownedQGraphicsSceneHelpEvent(uintptr(h)) }

func (this *QGraphicsSceneHelpEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneHelpEvent() *QGraphicsSceneHelpEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneHelpEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneHelpEvent_virtIdx, QGraphicsSceneHelpEvent_enable)
	return __o
}

func NewQGraphicsSceneHelpEvent2(type_ QEvent__Type) *QGraphicsSceneHelpEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_new2, uintptr(__goid), uintptr(type_))
	__o := newownedQGraphicsSceneHelpEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneHelpEvent_virtIdx, QGraphicsSceneHelpEvent_enable)
	return __o
}

func (this *QGraphicsSceneHelpEvent) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneHelpEvent) SetScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneHelpEvent_setScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHelpEvent) ScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneHelpEvent) SetScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneHelpEvent_setScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func QGraphicsSceneHelpEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneHelpEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneHelpEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneHelpEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneHelpEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHelpEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneHelpEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneHelpEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneHelpEvent__set_override(h, idx) }

func (this *QGraphicsSceneHelpEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneHelpEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneHelpEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneHelpEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
