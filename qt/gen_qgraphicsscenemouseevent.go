package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneMouseEvent_chain uintptr
	addr_QGraphicsSceneMouseEvent_new uintptr
	addr_QGraphicsSceneMouseEvent_new2 uintptr
	addr_QGraphicsSceneMouseEvent_pos uintptr
	addr_QGraphicsSceneMouseEvent_setPos uintptr
	addr_QGraphicsSceneMouseEvent_scenePos uintptr
	addr_QGraphicsSceneMouseEvent_setScenePos uintptr
	addr_QGraphicsSceneMouseEvent_screenPos uintptr
	addr_QGraphicsSceneMouseEvent_setScreenPos uintptr
	addr_QGraphicsSceneMouseEvent_buttonDownPos uintptr
	addr_QGraphicsSceneMouseEvent_setButtonDownPos uintptr
	addr_QGraphicsSceneMouseEvent_buttonDownScenePos uintptr
	addr_QGraphicsSceneMouseEvent_setButtonDownScenePos uintptr
	addr_QGraphicsSceneMouseEvent_buttonDownScreenPos uintptr
	addr_QGraphicsSceneMouseEvent_setButtonDownScreenPos uintptr
	addr_QGraphicsSceneMouseEvent_lastPos uintptr
	addr_QGraphicsSceneMouseEvent_setLastPos uintptr
	addr_QGraphicsSceneMouseEvent_lastScenePos uintptr
	addr_QGraphicsSceneMouseEvent_setLastScenePos uintptr
	addr_QGraphicsSceneMouseEvent_lastScreenPos uintptr
	addr_QGraphicsSceneMouseEvent_setLastScreenPos uintptr
	addr_QGraphicsSceneMouseEvent_buttons uintptr
	addr_QGraphicsSceneMouseEvent_setButtons uintptr
	addr_QGraphicsSceneMouseEvent_button uintptr
	addr_QGraphicsSceneMouseEvent_setButton uintptr
	addr_QGraphicsSceneMouseEvent_modifiers uintptr
	addr_QGraphicsSceneMouseEvent_setModifiers uintptr
	addr_QGraphicsSceneMouseEvent_source uintptr
	addr_QGraphicsSceneMouseEvent_setSource uintptr
	addr_QGraphicsSceneMouseEvent_flags uintptr
	addr_QGraphicsSceneMouseEvent_setFlags uintptr
	addr_QGraphicsSceneMouseEvent_registerEventType uintptr
	addr_QGraphicsSceneMouseEvent_registerEventType2 uintptr
	addr_QGraphicsSceneMouseEvent_setAccepted_default uintptr
	addr_QGraphicsSceneMouseEvent_clone_default uintptr
	pg_QGraphicsSceneMouseEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneMouseEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneMouseEvent_chain, "QGraphicsSceneMouseEvent_chain")
	symaddr(&addr_QGraphicsSceneMouseEvent_new, "QGraphicsSceneMouseEvent_new")
	symaddr(&addr_QGraphicsSceneMouseEvent_new2, "QGraphicsSceneMouseEvent_new2")
	symaddr(&addr_QGraphicsSceneMouseEvent_pos, "QGraphicsSceneMouseEvent_pos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setPos, "QGraphicsSceneMouseEvent_setPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_scenePos, "QGraphicsSceneMouseEvent_scenePos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setScenePos, "QGraphicsSceneMouseEvent_setScenePos")
	symaddr(&addr_QGraphicsSceneMouseEvent_screenPos, "QGraphicsSceneMouseEvent_screenPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setScreenPos, "QGraphicsSceneMouseEvent_setScreenPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_buttonDownPos, "QGraphicsSceneMouseEvent_buttonDownPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setButtonDownPos, "QGraphicsSceneMouseEvent_setButtonDownPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_buttonDownScenePos, "QGraphicsSceneMouseEvent_buttonDownScenePos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setButtonDownScenePos, "QGraphicsSceneMouseEvent_setButtonDownScenePos")
	symaddr(&addr_QGraphicsSceneMouseEvent_buttonDownScreenPos, "QGraphicsSceneMouseEvent_buttonDownScreenPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setButtonDownScreenPos, "QGraphicsSceneMouseEvent_setButtonDownScreenPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_lastPos, "QGraphicsSceneMouseEvent_lastPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setLastPos, "QGraphicsSceneMouseEvent_setLastPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_lastScenePos, "QGraphicsSceneMouseEvent_lastScenePos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setLastScenePos, "QGraphicsSceneMouseEvent_setLastScenePos")
	symaddr(&addr_QGraphicsSceneMouseEvent_lastScreenPos, "QGraphicsSceneMouseEvent_lastScreenPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_setLastScreenPos, "QGraphicsSceneMouseEvent_setLastScreenPos")
	symaddr(&addr_QGraphicsSceneMouseEvent_buttons, "QGraphicsSceneMouseEvent_buttons")
	symaddr(&addr_QGraphicsSceneMouseEvent_setButtons, "QGraphicsSceneMouseEvent_setButtons")
	symaddr(&addr_QGraphicsSceneMouseEvent_button, "QGraphicsSceneMouseEvent_button")
	symaddr(&addr_QGraphicsSceneMouseEvent_setButton, "QGraphicsSceneMouseEvent_setButton")
	symaddr(&addr_QGraphicsSceneMouseEvent_modifiers, "QGraphicsSceneMouseEvent_modifiers")
	symaddr(&addr_QGraphicsSceneMouseEvent_setModifiers, "QGraphicsSceneMouseEvent_setModifiers")
	symaddr(&addr_QGraphicsSceneMouseEvent_source, "QGraphicsSceneMouseEvent_source")
	symaddr(&addr_QGraphicsSceneMouseEvent_setSource, "QGraphicsSceneMouseEvent_setSource")
	symaddr(&addr_QGraphicsSceneMouseEvent_flags, "QGraphicsSceneMouseEvent_flags")
	symaddr(&addr_QGraphicsSceneMouseEvent_setFlags, "QGraphicsSceneMouseEvent_setFlags")
	symaddr(&addr_QGraphicsSceneMouseEvent_registerEventType, "QGraphicsSceneMouseEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneMouseEvent_registerEventType2, "QGraphicsSceneMouseEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneMouseEvent_setAccepted_default, "QGraphicsSceneMouseEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneMouseEvent_clone_default, "QGraphicsSceneMouseEvent_clone_default")
	reg(&pg_QGraphicsSceneMouseEvent__set_override, "QGraphicsSceneMouseEvent__set_override")
	reg(&pg_QGraphicsSceneMouseEvent_delete, "QGraphicsSceneMouseEvent_delete")
}

type QGraphicsSceneMouseEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneMouseEvent(h uintptr) *QGraphicsSceneMouseEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneMouseEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneMouseEvent(h uintptr) *QGraphicsSceneMouseEvent {
	o := bindQGraphicsSceneMouseEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneMouseEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneMouseEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneMouseEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneMouseEvent(h unsafe.Pointer) *QGraphicsSceneMouseEvent { return bindQGraphicsSceneMouseEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneMouseEventOwned(h unsafe.Pointer) *QGraphicsSceneMouseEvent { return newownedQGraphicsSceneMouseEvent(uintptr(h)) }

func (this *QGraphicsSceneMouseEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneMouseEvent() *QGraphicsSceneMouseEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneMouseEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneMouseEvent_virtIdx, QGraphicsSceneMouseEvent_enable)
	return __o
}

func NewQGraphicsSceneMouseEvent2(type_ QEvent__Type) *QGraphicsSceneMouseEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_new2, uintptr(__goid), uintptr(type_))
	__o := newownedQGraphicsSceneMouseEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneMouseEvent_virtIdx, QGraphicsSceneMouseEvent_enable)
	return __o
}

func (this *QGraphicsSceneMouseEvent) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMouseEvent) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMouseEvent) SetScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) ScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneMouseEvent) SetScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) ButtonDownPos(button MouseButton) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_buttonDownPos, this.h, uintptr(button))
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMouseEvent) SetButtonDownPos(button MouseButton, pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setButtonDownPos, this.h, uintptr(button), uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) ButtonDownScenePos(button MouseButton) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_buttonDownScenePos, this.h, uintptr(button))
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMouseEvent) SetButtonDownScenePos(button MouseButton, pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setButtonDownScenePos, this.h, uintptr(button), uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button MouseButton) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_buttonDownScreenPos, this.h, uintptr(button))
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneMouseEvent) SetButtonDownScreenPos(button MouseButton, pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setButtonDownScreenPos, this.h, uintptr(button), uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) LastPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_lastPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMouseEvent) SetLastPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setLastPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) LastScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_lastScenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneMouseEvent) SetLastScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setLastScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) LastScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_lastScreenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneMouseEvent) SetLastScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setLastScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneMouseEvent) Buttons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_buttons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QGraphicsSceneMouseEvent) SetButtons(buttons MouseButton) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneMouseEvent) Button() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_button, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QGraphicsSceneMouseEvent) SetButton(button MouseButton) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setButton, this.h, uintptr(button))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneMouseEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QGraphicsSceneMouseEvent) SetModifiers(modifiers KeyboardModifier) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setModifiers, this.h, uintptr(modifiers))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneMouseEvent) Source() MouseEventSource {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_source, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseEventSource(__r)
}

func (this *QGraphicsSceneMouseEvent) SetSource(source MouseEventSource) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setSource, this.h, uintptr(source))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneMouseEvent) Flags() MouseEventFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseEventFlag(__r)
}

func (this *QGraphicsSceneMouseEvent) SetFlags(arg1 MouseEventFlag) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setFlags, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func QGraphicsSceneMouseEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneMouseEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneMouseEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneMouseEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneMouseEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneMouseEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneMouseEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneMouseEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneMouseEvent__set_override(h, idx) }

func (this *QGraphicsSceneMouseEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneMouseEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneMouseEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneMouseEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
