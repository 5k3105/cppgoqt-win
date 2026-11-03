package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDropEvent_base func(uintptr) uintptr
	addr_QDropEvent_new uintptr
	addr_QDropEvent_new2 uintptr
	addr_QDropEvent_pos uintptr
	addr_QDropEvent_posF uintptr
	addr_QDropEvent_mouseButtons uintptr
	addr_QDropEvent_keyboardModifiers uintptr
	addr_QDropEvent_position uintptr
	addr_QDropEvent_buttons uintptr
	addr_QDropEvent_modifiers uintptr
	addr_QDropEvent_possibleActions uintptr
	addr_QDropEvent_proposedAction uintptr
	addr_QDropEvent_acceptProposedAction uintptr
	addr_QDropEvent_dropAction uintptr
	addr_QDropEvent_setDropAction uintptr
	addr_QDropEvent_source uintptr
	addr_QDropEvent_mimeData uintptr
	addr_QDropEvent_registerEventType uintptr
	addr_QDropEvent_registerEventType2 uintptr
	pg_QDropEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QDropEvent_base, "QDropEvent_base")
	symaddr(&addr_QDropEvent_new, "QDropEvent_new")
	symaddr(&addr_QDropEvent_new2, "QDropEvent_new2")
	symaddr(&addr_QDropEvent_pos, "QDropEvent_pos")
	symaddr(&addr_QDropEvent_posF, "QDropEvent_posF")
	symaddr(&addr_QDropEvent_mouseButtons, "QDropEvent_mouseButtons")
	symaddr(&addr_QDropEvent_keyboardModifiers, "QDropEvent_keyboardModifiers")
	symaddr(&addr_QDropEvent_position, "QDropEvent_position")
	symaddr(&addr_QDropEvent_buttons, "QDropEvent_buttons")
	symaddr(&addr_QDropEvent_modifiers, "QDropEvent_modifiers")
	symaddr(&addr_QDropEvent_possibleActions, "QDropEvent_possibleActions")
	symaddr(&addr_QDropEvent_proposedAction, "QDropEvent_proposedAction")
	symaddr(&addr_QDropEvent_acceptProposedAction, "QDropEvent_acceptProposedAction")
	symaddr(&addr_QDropEvent_dropAction, "QDropEvent_dropAction")
	symaddr(&addr_QDropEvent_setDropAction, "QDropEvent_setDropAction")
	symaddr(&addr_QDropEvent_source, "QDropEvent_source")
	symaddr(&addr_QDropEvent_mimeData, "QDropEvent_mimeData")
	symaddr(&addr_QDropEvent_registerEventType, "QDropEvent_registerEventType")
	symaddr(&addr_QDropEvent_registerEventType2, "QDropEvent_registerEventType2")
	reg(&pg_QDropEvent_delete, "QDropEvent_delete")
}

type QDropEvent struct {
	h uintptr
	*QEvent
}

func bindQDropEvent(h uintptr) *QDropEvent {
	if h == 0 {
		return nil
	}
	return &QDropEvent{h: h, QEvent: bindQEvent(pg_QDropEvent_base(h))}
}

func newownedQDropEvent(h uintptr) *QDropEvent {
	o := bindQDropEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDropEvent).Delete)
	}
	return o
}

func (this *QDropEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDropEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDropEvent(h unsafe.Pointer) *QDropEvent { return bindQDropEvent(uintptr(h)) }
func UnsafeNewQDropEventOwned(h unsafe.Pointer) *QDropEvent { return newownedQDropEvent(uintptr(h)) }

func (this *QDropEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQDropEvent(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDropEvent {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_new, uintptr(pos.UnsafePointer()), uintptr(actions), uintptr(data.UnsafePointer()), uintptr(buttons), uintptr(modifiers))
	__o := newownedQDropEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(data)
	return __o
}

func NewQDropEvent2(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, type_ QEvent__Type) *QDropEvent {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_new2, uintptr(pos.UnsafePointer()), uintptr(actions), uintptr(data.UnsafePointer()), uintptr(buttons), uintptr(modifiers), uintptr(type_))
	__o := newownedQDropEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(data)
	return __o
}

func (this *QDropEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QDropEvent) PosF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_posF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QDropEvent) MouseButtons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_mouseButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QDropEvent) KeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_keyboardModifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QDropEvent) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QDropEvent) Buttons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_buttons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QDropEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QDropEvent) PossibleActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_possibleActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QDropEvent) ProposedAction() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_proposedAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QDropEvent) AcceptProposedAction() {
	purego.SyscallN(addr_QDropEvent_acceptProposedAction, this.h)
	runtime.KeepAlive(this)
}

func (this *QDropEvent) DropAction() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_dropAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QDropEvent) SetDropAction(action DropAction) {
	purego.SyscallN(addr_QDropEvent_setDropAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QDropEvent) Source() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_source, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDropEvent) MimeData() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QDropEvent_mimeData, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func QDropEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QDropEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QDropEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QDropEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDropEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
