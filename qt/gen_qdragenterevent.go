package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDragEnterEvent_chain uintptr
	addr_QDragEnterEvent_new uintptr
	addr_QDragEnterEvent_registerEventType uintptr
	addr_QDragEnterEvent_registerEventType2 uintptr
	pg_QDragEnterEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDragEnterEvent_chain, "QDragEnterEvent_chain")
	symaddr(&addr_QDragEnterEvent_new, "QDragEnterEvent_new")
	symaddr(&addr_QDragEnterEvent_registerEventType, "QDragEnterEvent_registerEventType")
	symaddr(&addr_QDragEnterEvent_registerEventType2, "QDragEnterEvent_registerEventType2")
	reg(&pg_QDragEnterEvent_delete, "QDragEnterEvent_delete")
}

type QDragEnterEvent struct {
	h uintptr
	*QDragMoveEvent
}

func bindQDragEnterEvent(h uintptr) *QDragEnterEvent {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QDragEnterEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDragEnterEvent{h: h, QDragMoveEvent: &QDragMoveEvent{h: __c[0], QDropEvent: &QDropEvent{h: __c[1], QEvent: &QEvent{h: __c[2]}}}}
}

func newownedQDragEnterEvent(h uintptr) *QDragEnterEvent {
	o := bindQDragEnterEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDragEnterEvent).Delete)
	}
	return o
}

func (this *QDragEnterEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDragEnterEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDragEnterEvent(h unsafe.Pointer) *QDragEnterEvent { return bindQDragEnterEvent(uintptr(h)) }
func UnsafeNewQDragEnterEventOwned(h unsafe.Pointer) *QDragEnterEvent { return newownedQDragEnterEvent(uintptr(h)) }

func (this *QDragEnterEvent) UnsafeRelease() {
	this.h = 0
	if this.QDragMoveEvent != nil {
		this.QDragMoveEvent.UnsafeRelease()
	}
}

func NewQDragEnterEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragEnterEvent {
	__r, _, _ := purego.SyscallN(addr_QDragEnterEvent_new, uintptr(pos.UnsafePointer()), uintptr(actions), uintptr(data.UnsafePointer()), uintptr(buttons), uintptr(modifiers))
	__o := newownedQDragEnterEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(data)
	return __o
}

func QDragEnterEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QDragEnterEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QDragEnterEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QDragEnterEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QDragEnterEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDragEnterEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
