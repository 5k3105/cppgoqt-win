package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDragMoveEvent_chain uintptr
	addr_QDragMoveEvent_new uintptr
	addr_QDragMoveEvent_new2 uintptr
	addr_QDragMoveEvent_answerRect uintptr
	addr_QDragMoveEvent_accept uintptr
	addr_QDragMoveEvent_ignore uintptr
	addr_QDragMoveEvent_accept2 uintptr
	addr_QDragMoveEvent_ignore2 uintptr
	addr_QDragMoveEvent_registerEventType uintptr
	addr_QDragMoveEvent_registerEventType2 uintptr
	pg_QDragMoveEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDragMoveEvent_chain, "QDragMoveEvent_chain")
	symaddr(&addr_QDragMoveEvent_new, "QDragMoveEvent_new")
	symaddr(&addr_QDragMoveEvent_new2, "QDragMoveEvent_new2")
	symaddr(&addr_QDragMoveEvent_answerRect, "QDragMoveEvent_answerRect")
	symaddr(&addr_QDragMoveEvent_accept, "QDragMoveEvent_accept")
	symaddr(&addr_QDragMoveEvent_ignore, "QDragMoveEvent_ignore")
	symaddr(&addr_QDragMoveEvent_accept2, "QDragMoveEvent_accept2")
	symaddr(&addr_QDragMoveEvent_ignore2, "QDragMoveEvent_ignore2")
	symaddr(&addr_QDragMoveEvent_registerEventType, "QDragMoveEvent_registerEventType")
	symaddr(&addr_QDragMoveEvent_registerEventType2, "QDragMoveEvent_registerEventType2")
	reg(&pg_QDragMoveEvent_delete, "QDragMoveEvent_delete")
}

type QDragMoveEvent struct {
	h uintptr
	*QDropEvent
}

func bindQDragMoveEvent(h uintptr) *QDragMoveEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QDragMoveEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDragMoveEvent{h: h, QDropEvent: &QDropEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQDragMoveEvent(h uintptr) *QDragMoveEvent {
	o := bindQDragMoveEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDragMoveEvent).Delete)
	}
	return o
}

func (this *QDragMoveEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDragMoveEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDragMoveEvent(h unsafe.Pointer) *QDragMoveEvent { return bindQDragMoveEvent(uintptr(h)) }
func UnsafeNewQDragMoveEventOwned(h unsafe.Pointer) *QDragMoveEvent { return newownedQDragMoveEvent(uintptr(h)) }

func (this *QDragMoveEvent) UnsafeRelease() {
	this.h = 0
	if this.QDropEvent != nil {
		this.QDropEvent.UnsafeRelease()
	}
}

func NewQDragMoveEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragMoveEvent {
	__r, _, _ := purego.SyscallN(addr_QDragMoveEvent_new, uintptr(pos.UnsafePointer()), uintptr(actions), uintptr(data.UnsafePointer()), uintptr(buttons), uintptr(modifiers))
	__o := newownedQDragMoveEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(data)
	return __o
}

func NewQDragMoveEvent2(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, type_ QEvent__Type) *QDragMoveEvent {
	__r, _, _ := purego.SyscallN(addr_QDragMoveEvent_new2, uintptr(pos.UnsafePointer()), uintptr(actions), uintptr(data.UnsafePointer()), uintptr(buttons), uintptr(modifiers), uintptr(type_))
	__o := newownedQDragMoveEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(data)
	return __o
}

func (this *QDragMoveEvent) AnswerRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QDragMoveEvent_answerRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QDragMoveEvent) Accept() {
	purego.SyscallN(addr_QDragMoveEvent_accept, this.h)
	runtime.KeepAlive(this)
}

func (this *QDragMoveEvent) Ignore() {
	purego.SyscallN(addr_QDragMoveEvent_ignore, this.h)
	runtime.KeepAlive(this)
}

func (this *QDragMoveEvent) AcceptWithQRect(r *QRect) {
	purego.SyscallN(addr_QDragMoveEvent_accept2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QDragMoveEvent) IgnoreWithQRect(r *QRect) {
	purego.SyscallN(addr_QDragMoveEvent_ignore2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func QDragMoveEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QDragMoveEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QDragMoveEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QDragMoveEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QDragMoveEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDragMoveEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
