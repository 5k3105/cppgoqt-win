package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMouseEvent_chain uintptr
	addr_QMouseEvent_new uintptr
	addr_QMouseEvent_new2 uintptr
	addr_QMouseEvent_new3 uintptr
	addr_QMouseEvent_new4 uintptr
	addr_QMouseEvent_new5 uintptr
	addr_QMouseEvent_new6 uintptr
	addr_QMouseEvent_new7 uintptr
	addr_QMouseEvent_new8 uintptr
	addr_QMouseEvent_pos uintptr
	addr_QMouseEvent_globalPos uintptr
	addr_QMouseEvent_x uintptr
	addr_QMouseEvent_y uintptr
	addr_QMouseEvent_globalX uintptr
	addr_QMouseEvent_globalY uintptr
	addr_QMouseEvent_localPos uintptr
	addr_QMouseEvent_windowPos uintptr
	addr_QMouseEvent_screenPos uintptr
	addr_QMouseEvent_source uintptr
	addr_QMouseEvent_flags uintptr
	addr_QMouseEvent_registerEventType uintptr
	addr_QMouseEvent_registerEventType2 uintptr
	pg_QMouseEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMouseEvent_chain, "QMouseEvent_chain")
	symaddr(&addr_QMouseEvent_new, "QMouseEvent_new")
	symaddr(&addr_QMouseEvent_new2, "QMouseEvent_new2")
	symaddr(&addr_QMouseEvent_new3, "QMouseEvent_new3")
	symaddr(&addr_QMouseEvent_new4, "QMouseEvent_new4")
	symaddr(&addr_QMouseEvent_new5, "QMouseEvent_new5")
	symaddr(&addr_QMouseEvent_new6, "QMouseEvent_new6")
	symaddr(&addr_QMouseEvent_new7, "QMouseEvent_new7")
	symaddr(&addr_QMouseEvent_new8, "QMouseEvent_new8")
	symaddr(&addr_QMouseEvent_pos, "QMouseEvent_pos")
	symaddr(&addr_QMouseEvent_globalPos, "QMouseEvent_globalPos")
	symaddr(&addr_QMouseEvent_x, "QMouseEvent_x")
	symaddr(&addr_QMouseEvent_y, "QMouseEvent_y")
	symaddr(&addr_QMouseEvent_globalX, "QMouseEvent_globalX")
	symaddr(&addr_QMouseEvent_globalY, "QMouseEvent_globalY")
	symaddr(&addr_QMouseEvent_localPos, "QMouseEvent_localPos")
	symaddr(&addr_QMouseEvent_windowPos, "QMouseEvent_windowPos")
	symaddr(&addr_QMouseEvent_screenPos, "QMouseEvent_screenPos")
	symaddr(&addr_QMouseEvent_source, "QMouseEvent_source")
	symaddr(&addr_QMouseEvent_flags, "QMouseEvent_flags")
	symaddr(&addr_QMouseEvent_registerEventType, "QMouseEvent_registerEventType")
	symaddr(&addr_QMouseEvent_registerEventType2, "QMouseEvent_registerEventType2")
	reg(&pg_QMouseEvent_delete, "QMouseEvent_delete")
}

type QMouseEvent struct {
	h uintptr
	*QSinglePointEvent
}

func bindQMouseEvent(h uintptr) *QMouseEvent {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QMouseEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMouseEvent{h: h, QSinglePointEvent: &QSinglePointEvent{h: __c[0], QPointerEvent: &QPointerEvent{h: __c[1], QInputEvent: &QInputEvent{h: __c[2], QEvent: &QEvent{h: __c[3]}}}}}
}

func newownedQMouseEvent(h uintptr) *QMouseEvent {
	o := bindQMouseEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMouseEvent).Delete)
	}
	return o
}

func (this *QMouseEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMouseEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMouseEvent(h unsafe.Pointer) *QMouseEvent { return bindQMouseEvent(uintptr(h)) }
func UnsafeNewQMouseEventOwned(h unsafe.Pointer) *QMouseEvent { return newownedQMouseEvent(uintptr(h)) }

func (this *QMouseEvent) UnsafeRelease() {
	this.h = 0
	if this.QSinglePointEvent != nil {
		this.QSinglePointEvent.UnsafeRelease()
	}
}

func NewQMouseEvent(type_ QEvent__Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	return __o
}

func NewQMouseEvent2(type_ QEvent__Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new2, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(globalPos)
	return __o
}

func NewQMouseEvent3(type_ QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new3, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	return __o
}

func NewQMouseEvent4(type_ QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new4, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers), uintptr(source))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	return __o
}

func NewQMouseEvent5(type_ QEvent__Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new5, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers), uintptr(device.UnsafePointer()))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(device)
	return __o
}

func NewQMouseEvent6(type_ QEvent__Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new6, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers), uintptr(device.UnsafePointer()))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(device)
	return __o
}

func NewQMouseEvent7(type_ QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new7, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers), uintptr(device.UnsafePointer()))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(device)
	return __o
}

func NewQMouseEvent8(type_ QEvent__Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource, device *QPointingDevice) *QMouseEvent {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_new8, uintptr(type_), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(button), uintptr(buttons), uintptr(modifiers), uintptr(source), uintptr(device.UnsafePointer()))
	__o := newownedQMouseEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(device)
	return __o
}

func (this *QMouseEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QMouseEvent) GlobalPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_globalPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QMouseEvent) X() int {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMouseEvent) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMouseEvent) GlobalX() int {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_globalX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMouseEvent) GlobalY() int {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_globalY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMouseEvent) LocalPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_localPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QMouseEvent) WindowPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_windowPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QMouseEvent) ScreenPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QMouseEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QMouseEvent) Source() MouseEventSource {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_source, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseEventSource(__r)
}

func (this *QMouseEvent) Flags() MouseEventFlag {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseEventFlag(__r)
}

func QMouseEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QMouseEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QMouseEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QMouseEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMouseEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
