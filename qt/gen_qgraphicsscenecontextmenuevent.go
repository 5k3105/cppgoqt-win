package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneContextMenuEvent_chain uintptr
	addr_QGraphicsSceneContextMenuEvent_new uintptr
	addr_QGraphicsSceneContextMenuEvent_new2 uintptr
	addr_QGraphicsSceneContextMenuEvent_pos uintptr
	addr_QGraphicsSceneContextMenuEvent_setPos uintptr
	addr_QGraphicsSceneContextMenuEvent_scenePos uintptr
	addr_QGraphicsSceneContextMenuEvent_setScenePos uintptr
	addr_QGraphicsSceneContextMenuEvent_screenPos uintptr
	addr_QGraphicsSceneContextMenuEvent_setScreenPos uintptr
	addr_QGraphicsSceneContextMenuEvent_modifiers uintptr
	addr_QGraphicsSceneContextMenuEvent_setModifiers uintptr
	addr_QGraphicsSceneContextMenuEvent_reason uintptr
	addr_QGraphicsSceneContextMenuEvent_setReason uintptr
	addr_QGraphicsSceneContextMenuEvent_registerEventType uintptr
	addr_QGraphicsSceneContextMenuEvent_registerEventType2 uintptr
	addr_QGraphicsSceneContextMenuEvent_setAccepted_default uintptr
	addr_QGraphicsSceneContextMenuEvent_clone_default uintptr
	pg_QGraphicsSceneContextMenuEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneContextMenuEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneContextMenuEvent_chain, "QGraphicsSceneContextMenuEvent_chain")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_new, "QGraphicsSceneContextMenuEvent_new")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_new2, "QGraphicsSceneContextMenuEvent_new2")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_pos, "QGraphicsSceneContextMenuEvent_pos")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_setPos, "QGraphicsSceneContextMenuEvent_setPos")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_scenePos, "QGraphicsSceneContextMenuEvent_scenePos")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_setScenePos, "QGraphicsSceneContextMenuEvent_setScenePos")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_screenPos, "QGraphicsSceneContextMenuEvent_screenPos")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_setScreenPos, "QGraphicsSceneContextMenuEvent_setScreenPos")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_modifiers, "QGraphicsSceneContextMenuEvent_modifiers")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_setModifiers, "QGraphicsSceneContextMenuEvent_setModifiers")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_reason, "QGraphicsSceneContextMenuEvent_reason")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_setReason, "QGraphicsSceneContextMenuEvent_setReason")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_registerEventType, "QGraphicsSceneContextMenuEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_registerEventType2, "QGraphicsSceneContextMenuEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_setAccepted_default, "QGraphicsSceneContextMenuEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneContextMenuEvent_clone_default, "QGraphicsSceneContextMenuEvent_clone_default")
	reg(&pg_QGraphicsSceneContextMenuEvent__set_override, "QGraphicsSceneContextMenuEvent__set_override")
	reg(&pg_QGraphicsSceneContextMenuEvent_delete, "QGraphicsSceneContextMenuEvent_delete")
}

type QGraphicsSceneContextMenuEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneContextMenuEvent(h uintptr) *QGraphicsSceneContextMenuEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneContextMenuEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneContextMenuEvent(h uintptr) *QGraphicsSceneContextMenuEvent {
	o := bindQGraphicsSceneContextMenuEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneContextMenuEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneContextMenuEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneContextMenuEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneContextMenuEvent(h unsafe.Pointer) *QGraphicsSceneContextMenuEvent { return bindQGraphicsSceneContextMenuEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneContextMenuEventOwned(h unsafe.Pointer) *QGraphicsSceneContextMenuEvent { return newownedQGraphicsSceneContextMenuEvent(uintptr(h)) }

func (this *QGraphicsSceneContextMenuEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneContextMenuEvent() *QGraphicsSceneContextMenuEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneContextMenuEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneContextMenuEvent_virtIdx, QGraphicsSceneContextMenuEvent_enable)
	return __o
}

func NewQGraphicsSceneContextMenuEvent2(type_ QEvent__Type) *QGraphicsSceneContextMenuEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_new2, uintptr(__goid), uintptr(type_))
	__o := newownedQGraphicsSceneContextMenuEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneContextMenuEvent_virtIdx, QGraphicsSceneContextMenuEvent_enable)
	return __o
}

func (this *QGraphicsSceneContextMenuEvent) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneContextMenuEvent) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneContextMenuEvent) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneContextMenuEvent) SetScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_setScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneContextMenuEvent) ScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneContextMenuEvent) SetScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_setScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneContextMenuEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QGraphicsSceneContextMenuEvent) SetModifiers(modifiers KeyboardModifier) {
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_setModifiers, this.h, uintptr(modifiers))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneContextMenuEvent) Reason() QGraphicsSceneContextMenuEvent__Reason {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_reason, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsSceneContextMenuEvent__Reason(__r)
}

func (this *QGraphicsSceneContextMenuEvent) SetReason(reason QGraphicsSceneContextMenuEvent__Reason) {
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_setReason, this.h, uintptr(reason))
	runtime.KeepAlive(this)
}

func QGraphicsSceneContextMenuEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneContextMenuEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneContextMenuEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneContextMenuEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneContextMenuEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneContextMenuEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneContextMenuEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneContextMenuEvent__set_override(h, idx) }

func (this *QGraphicsSceneContextMenuEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneContextMenuEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneContextMenuEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneContextMenuEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
