package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSinglePointEvent_chain uintptr
	addr_QSinglePointEvent_button uintptr
	addr_QSinglePointEvent_buttons uintptr
	addr_QSinglePointEvent_position uintptr
	addr_QSinglePointEvent_scenePosition uintptr
	addr_QSinglePointEvent_globalPosition uintptr
	addr_QSinglePointEvent_exclusivePointGrabber uintptr
	addr_QSinglePointEvent_setExclusivePointGrabber uintptr
	addr_QSinglePointEvent_registerEventType uintptr
	addr_QSinglePointEvent_registerEventType2 uintptr
	pg_QSinglePointEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSinglePointEvent_chain, "QSinglePointEvent_chain")
	symaddr(&addr_QSinglePointEvent_button, "QSinglePointEvent_button")
	symaddr(&addr_QSinglePointEvent_buttons, "QSinglePointEvent_buttons")
	symaddr(&addr_QSinglePointEvent_position, "QSinglePointEvent_position")
	symaddr(&addr_QSinglePointEvent_scenePosition, "QSinglePointEvent_scenePosition")
	symaddr(&addr_QSinglePointEvent_globalPosition, "QSinglePointEvent_globalPosition")
	symaddr(&addr_QSinglePointEvent_exclusivePointGrabber, "QSinglePointEvent_exclusivePointGrabber")
	symaddr(&addr_QSinglePointEvent_setExclusivePointGrabber, "QSinglePointEvent_setExclusivePointGrabber")
	symaddr(&addr_QSinglePointEvent_registerEventType, "QSinglePointEvent_registerEventType")
	symaddr(&addr_QSinglePointEvent_registerEventType2, "QSinglePointEvent_registerEventType2")
	reg(&pg_QSinglePointEvent_delete, "QSinglePointEvent_delete")
}

type QSinglePointEvent struct {
	h uintptr
	*QPointerEvent
}

func bindQSinglePointEvent(h uintptr) *QSinglePointEvent {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSinglePointEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSinglePointEvent{h: h, QPointerEvent: &QPointerEvent{h: __c[0], QInputEvent: &QInputEvent{h: __c[1], QEvent: &QEvent{h: __c[2]}}}}
}

func newownedQSinglePointEvent(h uintptr) *QSinglePointEvent {
	o := bindQSinglePointEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSinglePointEvent).Delete)
	}
	return o
}

func (this *QSinglePointEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSinglePointEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSinglePointEvent(h unsafe.Pointer) *QSinglePointEvent { return bindQSinglePointEvent(uintptr(h)) }
func UnsafeNewQSinglePointEventOwned(h unsafe.Pointer) *QSinglePointEvent { return newownedQSinglePointEvent(uintptr(h)) }

func (this *QSinglePointEvent) UnsafeRelease() {
	this.h = 0
	if this.QPointerEvent != nil {
		this.QPointerEvent.UnsafeRelease()
	}
}

func (this *QSinglePointEvent) Button() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QSinglePointEvent_button, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QSinglePointEvent) Buttons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QSinglePointEvent_buttons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QSinglePointEvent) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QSinglePointEvent_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QSinglePointEvent) ScenePosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QSinglePointEvent_scenePosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QSinglePointEvent) GlobalPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QSinglePointEvent_globalPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QSinglePointEvent) ExclusivePointGrabber() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSinglePointEvent_exclusivePointGrabber, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSinglePointEvent) SetExclusivePointGrabber(exclusiveGrabber *QObject) {
	purego.SyscallN(addr_QSinglePointEvent_setExclusivePointGrabber, this.h, uintptr(exclusiveGrabber.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exclusiveGrabber)
}

func QSinglePointEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QSinglePointEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QSinglePointEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QSinglePointEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QSinglePointEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSinglePointEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
