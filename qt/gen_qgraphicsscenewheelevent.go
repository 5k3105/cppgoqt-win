package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneWheelEvent_chain uintptr
	addr_QGraphicsSceneWheelEvent_new uintptr
	addr_QGraphicsSceneWheelEvent_new2 uintptr
	addr_QGraphicsSceneWheelEvent_pos uintptr
	addr_QGraphicsSceneWheelEvent_setPos uintptr
	addr_QGraphicsSceneWheelEvent_scenePos uintptr
	addr_QGraphicsSceneWheelEvent_setScenePos uintptr
	addr_QGraphicsSceneWheelEvent_screenPos uintptr
	addr_QGraphicsSceneWheelEvent_setScreenPos uintptr
	addr_QGraphicsSceneWheelEvent_buttons uintptr
	addr_QGraphicsSceneWheelEvent_setButtons uintptr
	addr_QGraphicsSceneWheelEvent_modifiers uintptr
	addr_QGraphicsSceneWheelEvent_setModifiers uintptr
	addr_QGraphicsSceneWheelEvent_delta uintptr
	addr_QGraphicsSceneWheelEvent_setDelta uintptr
	addr_QGraphicsSceneWheelEvent_orientation uintptr
	addr_QGraphicsSceneWheelEvent_setOrientation uintptr
	addr_QGraphicsSceneWheelEvent_phase uintptr
	addr_QGraphicsSceneWheelEvent_setPhase uintptr
	addr_QGraphicsSceneWheelEvent_pixelDelta uintptr
	addr_QGraphicsSceneWheelEvent_setPixelDelta uintptr
	addr_QGraphicsSceneWheelEvent_isInverted uintptr
	addr_QGraphicsSceneWheelEvent_setInverted uintptr
	addr_QGraphicsSceneWheelEvent_registerEventType uintptr
	addr_QGraphicsSceneWheelEvent_registerEventType2 uintptr
	addr_QGraphicsSceneWheelEvent_setAccepted_default uintptr
	addr_QGraphicsSceneWheelEvent_clone_default uintptr
	pg_QGraphicsSceneWheelEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneWheelEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneWheelEvent_chain, "QGraphicsSceneWheelEvent_chain")
	symaddr(&addr_QGraphicsSceneWheelEvent_new, "QGraphicsSceneWheelEvent_new")
	symaddr(&addr_QGraphicsSceneWheelEvent_new2, "QGraphicsSceneWheelEvent_new2")
	symaddr(&addr_QGraphicsSceneWheelEvent_pos, "QGraphicsSceneWheelEvent_pos")
	symaddr(&addr_QGraphicsSceneWheelEvent_setPos, "QGraphicsSceneWheelEvent_setPos")
	symaddr(&addr_QGraphicsSceneWheelEvent_scenePos, "QGraphicsSceneWheelEvent_scenePos")
	symaddr(&addr_QGraphicsSceneWheelEvent_setScenePos, "QGraphicsSceneWheelEvent_setScenePos")
	symaddr(&addr_QGraphicsSceneWheelEvent_screenPos, "QGraphicsSceneWheelEvent_screenPos")
	symaddr(&addr_QGraphicsSceneWheelEvent_setScreenPos, "QGraphicsSceneWheelEvent_setScreenPos")
	symaddr(&addr_QGraphicsSceneWheelEvent_buttons, "QGraphicsSceneWheelEvent_buttons")
	symaddr(&addr_QGraphicsSceneWheelEvent_setButtons, "QGraphicsSceneWheelEvent_setButtons")
	symaddr(&addr_QGraphicsSceneWheelEvent_modifiers, "QGraphicsSceneWheelEvent_modifiers")
	symaddr(&addr_QGraphicsSceneWheelEvent_setModifiers, "QGraphicsSceneWheelEvent_setModifiers")
	symaddr(&addr_QGraphicsSceneWheelEvent_delta, "QGraphicsSceneWheelEvent_delta")
	symaddr(&addr_QGraphicsSceneWheelEvent_setDelta, "QGraphicsSceneWheelEvent_setDelta")
	symaddr(&addr_QGraphicsSceneWheelEvent_orientation, "QGraphicsSceneWheelEvent_orientation")
	symaddr(&addr_QGraphicsSceneWheelEvent_setOrientation, "QGraphicsSceneWheelEvent_setOrientation")
	symaddr(&addr_QGraphicsSceneWheelEvent_phase, "QGraphicsSceneWheelEvent_phase")
	symaddr(&addr_QGraphicsSceneWheelEvent_setPhase, "QGraphicsSceneWheelEvent_setPhase")
	symaddr(&addr_QGraphicsSceneWheelEvent_pixelDelta, "QGraphicsSceneWheelEvent_pixelDelta")
	symaddr(&addr_QGraphicsSceneWheelEvent_setPixelDelta, "QGraphicsSceneWheelEvent_setPixelDelta")
	symaddr(&addr_QGraphicsSceneWheelEvent_isInverted, "QGraphicsSceneWheelEvent_isInverted")
	symaddr(&addr_QGraphicsSceneWheelEvent_setInverted, "QGraphicsSceneWheelEvent_setInverted")
	symaddr(&addr_QGraphicsSceneWheelEvent_registerEventType, "QGraphicsSceneWheelEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneWheelEvent_registerEventType2, "QGraphicsSceneWheelEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneWheelEvent_setAccepted_default, "QGraphicsSceneWheelEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneWheelEvent_clone_default, "QGraphicsSceneWheelEvent_clone_default")
	reg(&pg_QGraphicsSceneWheelEvent__set_override, "QGraphicsSceneWheelEvent__set_override")
	reg(&pg_QGraphicsSceneWheelEvent_delete, "QGraphicsSceneWheelEvent_delete")
}

type QGraphicsSceneWheelEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneWheelEvent(h uintptr) *QGraphicsSceneWheelEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneWheelEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneWheelEvent(h uintptr) *QGraphicsSceneWheelEvent {
	o := bindQGraphicsSceneWheelEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneWheelEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneWheelEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneWheelEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneWheelEvent(h unsafe.Pointer) *QGraphicsSceneWheelEvent { return bindQGraphicsSceneWheelEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneWheelEventOwned(h unsafe.Pointer) *QGraphicsSceneWheelEvent { return newownedQGraphicsSceneWheelEvent(uintptr(h)) }

func (this *QGraphicsSceneWheelEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneWheelEvent() *QGraphicsSceneWheelEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneWheelEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneWheelEvent_virtIdx, QGraphicsSceneWheelEvent_enable)
	return __o
}

func NewQGraphicsSceneWheelEvent2(type_ QEvent__Type) *QGraphicsSceneWheelEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_new2, uintptr(__goid), uintptr(type_))
	__o := newownedQGraphicsSceneWheelEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneWheelEvent_virtIdx, QGraphicsSceneWheelEvent_enable)
	return __o
}

func (this *QGraphicsSceneWheelEvent) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneWheelEvent) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneWheelEvent) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneWheelEvent) SetScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneWheelEvent) ScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneWheelEvent) SetScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneWheelEvent) Buttons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_buttons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QGraphicsSceneWheelEvent) SetButtons(buttons MouseButton) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneWheelEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QGraphicsSceneWheelEvent) SetModifiers(modifiers KeyboardModifier) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setModifiers, this.h, uintptr(modifiers))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneWheelEvent) Delta() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_delta, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsSceneWheelEvent) SetDelta(delta int) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setDelta, this.h, uintptr(delta))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneWheelEvent) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QGraphicsSceneWheelEvent) SetOrientation(orientation Orientation) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setOrientation, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneWheelEvent) Phase() ScrollPhase {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_phase, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScrollPhase(__r)
}

func (this *QGraphicsSceneWheelEvent) SetPhase(scrollPhase ScrollPhase) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setPhase, this.h, uintptr(scrollPhase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneWheelEvent) PixelDelta() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_pixelDelta, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneWheelEvent) SetPixelDelta(delta *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setPixelDelta, this.h, uintptr(delta.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delta)
}

func (this *QGraphicsSceneWheelEvent) IsInverted() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_isInverted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsSceneWheelEvent) SetInverted(inverted bool) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setInverted, this.h, uintptr(b2i(inverted)))
	runtime.KeepAlive(this)
}

func QGraphicsSceneWheelEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneWheelEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneWheelEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneWheelEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneWheelEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneWheelEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneWheelEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneWheelEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneWheelEvent__set_override(h, idx) }

func (this *QGraphicsSceneWheelEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneWheelEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneWheelEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneWheelEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
