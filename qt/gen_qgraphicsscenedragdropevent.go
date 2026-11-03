package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSceneDragDropEvent_chain uintptr
	addr_QGraphicsSceneDragDropEvent_new uintptr
	addr_QGraphicsSceneDragDropEvent_new2 uintptr
	addr_QGraphicsSceneDragDropEvent_pos uintptr
	addr_QGraphicsSceneDragDropEvent_setPos uintptr
	addr_QGraphicsSceneDragDropEvent_scenePos uintptr
	addr_QGraphicsSceneDragDropEvent_setScenePos uintptr
	addr_QGraphicsSceneDragDropEvent_screenPos uintptr
	addr_QGraphicsSceneDragDropEvent_setScreenPos uintptr
	addr_QGraphicsSceneDragDropEvent_buttons uintptr
	addr_QGraphicsSceneDragDropEvent_setButtons uintptr
	addr_QGraphicsSceneDragDropEvent_modifiers uintptr
	addr_QGraphicsSceneDragDropEvent_setModifiers uintptr
	addr_QGraphicsSceneDragDropEvent_possibleActions uintptr
	addr_QGraphicsSceneDragDropEvent_setPossibleActions uintptr
	addr_QGraphicsSceneDragDropEvent_proposedAction uintptr
	addr_QGraphicsSceneDragDropEvent_setProposedAction uintptr
	addr_QGraphicsSceneDragDropEvent_acceptProposedAction uintptr
	addr_QGraphicsSceneDragDropEvent_dropAction uintptr
	addr_QGraphicsSceneDragDropEvent_setDropAction uintptr
	addr_QGraphicsSceneDragDropEvent_source uintptr
	addr_QGraphicsSceneDragDropEvent_setSource uintptr
	addr_QGraphicsSceneDragDropEvent_mimeData uintptr
	addr_QGraphicsSceneDragDropEvent_setMimeData uintptr
	addr_QGraphicsSceneDragDropEvent_registerEventType uintptr
	addr_QGraphicsSceneDragDropEvent_registerEventType2 uintptr
	addr_QGraphicsSceneDragDropEvent_setAccepted_default uintptr
	addr_QGraphicsSceneDragDropEvent_clone_default uintptr
	pg_QGraphicsSceneDragDropEvent__set_override func(uintptr, int32)
	pg_QGraphicsSceneDragDropEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSceneDragDropEvent_chain, "QGraphicsSceneDragDropEvent_chain")
	symaddr(&addr_QGraphicsSceneDragDropEvent_new, "QGraphicsSceneDragDropEvent_new")
	symaddr(&addr_QGraphicsSceneDragDropEvent_new2, "QGraphicsSceneDragDropEvent_new2")
	symaddr(&addr_QGraphicsSceneDragDropEvent_pos, "QGraphicsSceneDragDropEvent_pos")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setPos, "QGraphicsSceneDragDropEvent_setPos")
	symaddr(&addr_QGraphicsSceneDragDropEvent_scenePos, "QGraphicsSceneDragDropEvent_scenePos")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setScenePos, "QGraphicsSceneDragDropEvent_setScenePos")
	symaddr(&addr_QGraphicsSceneDragDropEvent_screenPos, "QGraphicsSceneDragDropEvent_screenPos")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setScreenPos, "QGraphicsSceneDragDropEvent_setScreenPos")
	symaddr(&addr_QGraphicsSceneDragDropEvent_buttons, "QGraphicsSceneDragDropEvent_buttons")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setButtons, "QGraphicsSceneDragDropEvent_setButtons")
	symaddr(&addr_QGraphicsSceneDragDropEvent_modifiers, "QGraphicsSceneDragDropEvent_modifiers")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setModifiers, "QGraphicsSceneDragDropEvent_setModifiers")
	symaddr(&addr_QGraphicsSceneDragDropEvent_possibleActions, "QGraphicsSceneDragDropEvent_possibleActions")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setPossibleActions, "QGraphicsSceneDragDropEvent_setPossibleActions")
	symaddr(&addr_QGraphicsSceneDragDropEvent_proposedAction, "QGraphicsSceneDragDropEvent_proposedAction")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setProposedAction, "QGraphicsSceneDragDropEvent_setProposedAction")
	symaddr(&addr_QGraphicsSceneDragDropEvent_acceptProposedAction, "QGraphicsSceneDragDropEvent_acceptProposedAction")
	symaddr(&addr_QGraphicsSceneDragDropEvent_dropAction, "QGraphicsSceneDragDropEvent_dropAction")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setDropAction, "QGraphicsSceneDragDropEvent_setDropAction")
	symaddr(&addr_QGraphicsSceneDragDropEvent_source, "QGraphicsSceneDragDropEvent_source")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setSource, "QGraphicsSceneDragDropEvent_setSource")
	symaddr(&addr_QGraphicsSceneDragDropEvent_mimeData, "QGraphicsSceneDragDropEvent_mimeData")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setMimeData, "QGraphicsSceneDragDropEvent_setMimeData")
	symaddr(&addr_QGraphicsSceneDragDropEvent_registerEventType, "QGraphicsSceneDragDropEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneDragDropEvent_registerEventType2, "QGraphicsSceneDragDropEvent_registerEventType2")
	symaddr(&addr_QGraphicsSceneDragDropEvent_setAccepted_default, "QGraphicsSceneDragDropEvent_setAccepted_default")
	symaddr(&addr_QGraphicsSceneDragDropEvent_clone_default, "QGraphicsSceneDragDropEvent_clone_default")
	reg(&pg_QGraphicsSceneDragDropEvent__set_override, "QGraphicsSceneDragDropEvent__set_override")
	reg(&pg_QGraphicsSceneDragDropEvent_delete, "QGraphicsSceneDragDropEvent_delete")
}

type QGraphicsSceneDragDropEvent struct {
	h uintptr
	*QGraphicsSceneEvent
}

func bindQGraphicsSceneDragDropEvent(h uintptr) *QGraphicsSceneDragDropEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSceneDragDropEvent{h: h, QGraphicsSceneEvent: &QGraphicsSceneEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQGraphicsSceneDragDropEvent(h uintptr) *QGraphicsSceneDragDropEvent {
	o := bindQGraphicsSceneDragDropEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneDragDropEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneDragDropEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneDragDropEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneDragDropEvent(h unsafe.Pointer) *QGraphicsSceneDragDropEvent { return bindQGraphicsSceneDragDropEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneDragDropEventOwned(h unsafe.Pointer) *QGraphicsSceneDragDropEvent { return newownedQGraphicsSceneDragDropEvent(uintptr(h)) }

func (this *QGraphicsSceneDragDropEvent) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsSceneEvent != nil {
		this.QGraphicsSceneEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneDragDropEvent() *QGraphicsSceneDragDropEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_new, uintptr(__goid))
	__o := newownedQGraphicsSceneDragDropEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneDragDropEvent_virtIdx, QGraphicsSceneDragDropEvent_enable)
	return __o
}

func NewQGraphicsSceneDragDropEvent2(type_ QEvent__Type) *QGraphicsSceneDragDropEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_new2, uintptr(__goid), uintptr(type_))
	__o := newownedQGraphicsSceneDragDropEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSceneDragDropEvent_virtIdx, QGraphicsSceneDragDropEvent_enable)
	return __o
}

func (this *QGraphicsSceneDragDropEvent) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneDragDropEvent) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetScenePos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setScenePos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneDragDropEvent) ScreenPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetScreenPos(pos *QPoint) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setScreenPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsSceneDragDropEvent) Buttons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_buttons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetButtons(buttons MouseButton) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetModifiers(modifiers KeyboardModifier) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setModifiers, this.h, uintptr(modifiers))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) PossibleActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_possibleActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetPossibleActions(actions DropAction) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setPossibleActions, this.h, uintptr(actions))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) ProposedAction() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_proposedAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetProposedAction(action DropAction) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setProposedAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_acceptProposedAction, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) DropAction() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_dropAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetDropAction(action DropAction) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setDropAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) Source() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_source, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetSource(source *QWidget) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setSource, this.h, uintptr(source.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
}

func (this *QGraphicsSceneDragDropEvent) MimeData() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_mimeData, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetMimeData(data *QMimeData) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setMimeData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func QGraphicsSceneDragDropEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneDragDropEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneDragDropEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QGraphicsSceneDragDropEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSceneDragDropEvent) CloneDefault() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneDragDropEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QGraphicsSceneDragDropEvent_virtIdx(slot int64) int32 {
	switch slot {
	case 6367896784077321574:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QGraphicsSceneDragDropEvent_enable(h uintptr, idx int32) { pg_QGraphicsSceneDragDropEvent__set_override(h, idx) }

func (this *QGraphicsSceneDragDropEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QGraphicsSceneDragDropEvent) OnClone(fn func() *QEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QGraphicsSceneDragDropEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneDragDropEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
