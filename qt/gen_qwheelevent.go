package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWheelEvent_chain uintptr
	addr_QWheelEvent_new uintptr
	addr_QWheelEvent_new2 uintptr
	addr_QWheelEvent_new3 uintptr
	addr_QWheelEvent_pixelDelta uintptr
	addr_QWheelEvent_angleDelta uintptr
	addr_QWheelEvent_phase uintptr
	addr_QWheelEvent_inverted uintptr
	addr_QWheelEvent_isInverted uintptr
	addr_QWheelEvent_hasPixelDelta uintptr
	addr_QWheelEvent_source uintptr
	addr_QWheelEvent_registerEventType uintptr
	addr_QWheelEvent_registerEventType2 uintptr
	pg_QWheelEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWheelEvent_chain, "QWheelEvent_chain")
	symaddr(&addr_QWheelEvent_new, "QWheelEvent_new")
	symaddr(&addr_QWheelEvent_new2, "QWheelEvent_new2")
	symaddr(&addr_QWheelEvent_new3, "QWheelEvent_new3")
	symaddr(&addr_QWheelEvent_pixelDelta, "QWheelEvent_pixelDelta")
	symaddr(&addr_QWheelEvent_angleDelta, "QWheelEvent_angleDelta")
	symaddr(&addr_QWheelEvent_phase, "QWheelEvent_phase")
	symaddr(&addr_QWheelEvent_inverted, "QWheelEvent_inverted")
	symaddr(&addr_QWheelEvent_isInverted, "QWheelEvent_isInverted")
	symaddr(&addr_QWheelEvent_hasPixelDelta, "QWheelEvent_hasPixelDelta")
	symaddr(&addr_QWheelEvent_source, "QWheelEvent_source")
	symaddr(&addr_QWheelEvent_registerEventType, "QWheelEvent_registerEventType")
	symaddr(&addr_QWheelEvent_registerEventType2, "QWheelEvent_registerEventType2")
	reg(&pg_QWheelEvent_delete, "QWheelEvent_delete")
}

type QWheelEvent struct {
	h uintptr
	*QSinglePointEvent
}

func bindQWheelEvent(h uintptr) *QWheelEvent {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QWheelEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QWheelEvent{h: h, QSinglePointEvent: &QSinglePointEvent{h: __c[0], QPointerEvent: &QPointerEvent{h: __c[1], QInputEvent: &QInputEvent{h: __c[2], QEvent: &QEvent{h: __c[3]}}}}}
}

func newownedQWheelEvent(h uintptr) *QWheelEvent {
	o := bindQWheelEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWheelEvent).Delete)
	}
	return o
}

func (this *QWheelEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWheelEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWheelEvent(h unsafe.Pointer) *QWheelEvent { return bindQWheelEvent(uintptr(h)) }
func UnsafeNewQWheelEventOwned(h unsafe.Pointer) *QWheelEvent { return newownedQWheelEvent(uintptr(h)) }

func (this *QWheelEvent) UnsafeRelease() {
	this.h = 0
	if this.QSinglePointEvent != nil {
		this.QSinglePointEvent.UnsafeRelease()
	}
}

func NewQWheelEvent(pos *QPointF, globalPos *QPointF, pixelDelta *QPoint, angleDelta *QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool) *QWheelEvent {
	__r, _, _ := purego.SyscallN(addr_QWheelEvent_new, uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(pixelDelta.UnsafePointer()), uintptr(angleDelta.UnsafePointer()), uintptr(buttons), uintptr(modifiers), uintptr(phase), uintptr(b2i(inverted)))
	__o := newownedQWheelEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(pixelDelta)
	runtime.KeepAlive(angleDelta)
	return __o
}

func NewQWheelEvent2(pos *QPointF, globalPos *QPointF, pixelDelta *QPoint, angleDelta *QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource) *QWheelEvent {
	__r, _, _ := purego.SyscallN(addr_QWheelEvent_new2, uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(pixelDelta.UnsafePointer()), uintptr(angleDelta.UnsafePointer()), uintptr(buttons), uintptr(modifiers), uintptr(phase), uintptr(b2i(inverted)), uintptr(source))
	__o := newownedQWheelEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(pixelDelta)
	runtime.KeepAlive(angleDelta)
	return __o
}

func NewQWheelEvent3(pos *QPointF, globalPos *QPointF, pixelDelta *QPoint, angleDelta *QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource, device *QPointingDevice) *QWheelEvent {
	__r, _, _ := purego.SyscallN(addr_QWheelEvent_new3, uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(pixelDelta.UnsafePointer()), uintptr(angleDelta.UnsafePointer()), uintptr(buttons), uintptr(modifiers), uintptr(phase), uintptr(b2i(inverted)), uintptr(source), uintptr(device.UnsafePointer()))
	__o := newownedQWheelEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(pixelDelta)
	runtime.KeepAlive(angleDelta)
	runtime.KeepAlive(device)
	return __o
}

func (this *QWheelEvent) PixelDelta() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWheelEvent_pixelDelta, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QWheelEvent) AngleDelta() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWheelEvent_angleDelta, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QWheelEvent) Phase() ScrollPhase {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_phase, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScrollPhase(__r)
}

func (this *QWheelEvent) Inverted() bool {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_inverted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWheelEvent) IsInverted() bool {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_isInverted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWheelEvent) HasPixelDelta() bool {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_hasPixelDelta, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWheelEvent) Source() MouseEventSource {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_source, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseEventSource(__r)
}

func QWheelEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QWheelEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QWheelEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QWheelEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWheelEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
