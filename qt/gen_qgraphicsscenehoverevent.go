package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneHoverEvent_chain uintptr
	addr_QGraphicsSceneHoverEvent_new uintptr
	addr_QGraphicsSceneHoverEvent_new2 uintptr
	addr_QGraphicsSceneHoverEvent_pos uintptr
	addr_QGraphicsSceneHoverEvent_setPos uintptr
	addr_QGraphicsSceneHoverEvent_scenePos uintptr
	addr_QGraphicsSceneHoverEvent_setScenePos uintptr
	addr_QGraphicsSceneHoverEvent_screenPos uintptr
	addr_QGraphicsSceneHoverEvent_setScreenPos uintptr
	addr_QGraphicsSceneHoverEvent_lastPos uintptr
	addr_QGraphicsSceneHoverEvent_setLastPos uintptr
	addr_QGraphicsSceneHoverEvent_lastScenePos uintptr
	addr_QGraphicsSceneHoverEvent_setLastScenePos uintptr
	addr_QGraphicsSceneHoverEvent_lastScreenPos uintptr
	addr_QGraphicsSceneHoverEvent_setLastScreenPos uintptr
	addr_QGraphicsSceneHoverEvent_modifiers uintptr
	addr_QGraphicsSceneHoverEvent_setModifiers uintptr
	addr_QGraphicsSceneHoverEvent_registerEventType uintptr
	addr_QGraphicsSceneHoverEvent_registerEventType2 uintptr
	addr_QGraphicsSceneHoverEvent_setAccepted_default uintptr
	addr_QGraphicsSceneHoverEvent_clone_default uintptr
	pg_QGraphicsSceneHoverEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneHoverEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneHoverEvent_chain, "QGraphicsSceneHoverEvent_chain")
	symaddr(&addr_QGraphicsSceneHoverEvent_new, "QGraphicsSceneHoverEvent_new")
	symaddr(&addr_QGraphicsSceneHoverEvent_new2, "QGraphicsSceneHoverEvent_new2")
	symaddr(&addr_QGraphicsSceneHoverEvent_pos, "QGraphicsSceneHoverEvent_pos")
	symaddr(&addr_QGraphicsSceneHoverEvent_setPos, "QGraphicsSceneHoverEvent_setPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_scenePos, "QGraphicsSceneHoverEvent_scenePos")
	symaddr(&addr_QGraphicsSceneHoverEvent_setScenePos, "QGraphicsSceneHoverEvent_setScenePos")
	symaddr(&addr_QGraphicsSceneHoverEvent_screenPos, "QGraphicsSceneHoverEvent_screenPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_setScreenPos, "QGraphicsSceneHoverEvent_setScreenPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_lastPos, "QGraphicsSceneHoverEvent_lastPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_setLastPos, "QGraphicsSceneHoverEvent_setLastPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_lastScenePos, "QGraphicsSceneHoverEvent_lastScenePos")
	symaddr(&addr_QGraphicsSceneHoverEvent_setLastScenePos, "QGraphicsSceneHoverEvent_setLastScenePos")
	symaddr(&addr_QGraphicsSceneHoverEvent_lastScreenPos, "QGraphicsSceneHoverEvent_lastScreenPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_setLastScreenPos, "QGraphicsSceneHoverEvent_setLastScreenPos")
	symaddr(&addr_QGraphicsSceneHoverEvent_modifiers, "QGraphicsSceneHoverEvent_modifiers")
	symaddr(&addr_QGraphicsSceneHoverEvent_setModifiers, "QGraphicsSceneHoverEvent_setModifiers")
	symaddr(&addr_QGraphicsSceneHoverEvent_registerEventType, "QGraphicsSceneHoverEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneHoverEvent_registerEventType2, "QGraphicsSceneHoverEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneHoverEvent_setAccepted_default, "QGraphicsSceneHoverEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneHoverEvent_clone_default, "QGraphicsSceneHoverEvent_clone_default")
	reg(&pg_QGraphicsSceneHoverEvent__set_override, "QGraphicsSceneHoverEvent__set_override")
	reg(&pg_QGraphicsSceneHoverEvent_delete, "QGraphicsSceneHoverEvent_delete")
}

type QGraphicsSceneHoverEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneHoverEvent(h uintptr) *QGraphicsSceneHoverEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneHoverEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneHoverEvent(h uintptr) *QGraphicsSceneHoverEvent {
	o := bindQGraphicsSceneHoverEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneHoverEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneHoverEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneHoverEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneHoverEvent(h unsafe.Pointer) *QGraphicsSceneHoverEvent { return bindQGraphicsSceneHoverEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneHoverEventOwned(h unsafe.Pointer) *QGraphicsSceneHoverEvent { return newownedQGraphicsSceneHoverEvent(uintptr(h)) }

func (this *QGraphicsSceneHoverEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneHoverEvent() *QGraphicsSceneHoverEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneHoverEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneHoverEvent_virtIdx, QGraphicsSceneHoverEvent_enable)
	return __o
}

func NewQGraphicsSceneHoverEvent2(type_ QEvent__Type) *QGraphicsSceneHoverEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_new2, uintptr(__goid), uintptr(type_))
	__o := newownedQGraphicsSceneHoverEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneHoverEvent_virtIdx, QGraphicsSceneHoverEvent_enable)
	return __o
}

func (this *QGraphicsSceneHoverEvent) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneHoverEvent) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHoverEvent) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneHoverEvent) SetScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHoverEvent) ScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneHoverEvent) SetScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHoverEvent) LastPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_lastPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneHoverEvent) SetLastPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setLastPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHoverEvent) LastScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_lastScenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneHoverEvent) SetLastScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setLastScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHoverEvent) LastScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_lastScreenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneHoverEvent) SetLastScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setLastScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneHoverEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QGraphicsSceneHoverEvent) SetModifiers(modifiers KeyboardModifier) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setModifiers, this.h, uintptr(modifiers))
	runtime.KeepAlive(this)
}

func QGraphicsSceneHoverEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneHoverEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneHoverEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneHoverEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneHoverEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneHoverEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneHoverEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneHoverEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneHoverEvent__set_override(h, idx) }

func (this *QGraphicsSceneHoverEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneHoverEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneHoverEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneHoverEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
