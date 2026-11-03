package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QEnterEvent_chain uintptr
	addr_QEnterEvent_new uintptr
	addr_QEnterEvent_new2 uintptr
	addr_QEnterEvent_pos uintptr
	addr_QEnterEvent_globalPos uintptr
	addr_QEnterEvent_x uintptr
	addr_QEnterEvent_y uintptr
	addr_QEnterEvent_globalX uintptr
	addr_QEnterEvent_globalY uintptr
	addr_QEnterEvent_localPos uintptr
	addr_QEnterEvent_windowPos uintptr
	addr_QEnterEvent_screenPos uintptr
	addr_QEnterEvent_registerEventType uintptr
	addr_QEnterEvent_registerEventType2 uintptr
	pg_QEnterEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QEnterEvent_chain, "QEnterEvent_chain")
	symaddr(&addr_QEnterEvent_new, "QEnterEvent_new")
	symaddr(&addr_QEnterEvent_new2, "QEnterEvent_new2")
	symaddr(&addr_QEnterEvent_pos, "QEnterEvent_pos")
	symaddr(&addr_QEnterEvent_globalPos, "QEnterEvent_globalPos")
	symaddr(&addr_QEnterEvent_x, "QEnterEvent_x")
	symaddr(&addr_QEnterEvent_y, "QEnterEvent_y")
	symaddr(&addr_QEnterEvent_globalX, "QEnterEvent_globalX")
	symaddr(&addr_QEnterEvent_globalY, "QEnterEvent_globalY")
	symaddr(&addr_QEnterEvent_localPos, "QEnterEvent_localPos")
	symaddr(&addr_QEnterEvent_windowPos, "QEnterEvent_windowPos")
	symaddr(&addr_QEnterEvent_screenPos, "QEnterEvent_screenPos")
	symaddr(&addr_QEnterEvent_registerEventType, "QEnterEvent_registerEventType")
	symaddr(&addr_QEnterEvent_registerEventType2, "QEnterEvent_registerEventType2")
	reg(&pg_QEnterEvent_delete, "QEnterEvent_delete")
}

type QEnterEvent struct {
	h uintptr
	*QSinglePointEvent
}

func bindQEnterEvent(h uintptr) *QEnterEvent {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QEnterEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QEnterEvent{h: h, QSinglePointEvent: &QSinglePointEvent{h: __c[0], QPointerEvent: &QPointerEvent{h: __c[1], QInputEvent: &QInputEvent{h: __c[2], QEvent: &QEvent{h: __c[3]}}}}}
}

func newownedQEnterEvent(h uintptr) *QEnterEvent {
	o := bindQEnterEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QEnterEvent).Delete)
	}
	return o
}

func (this *QEnterEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QEnterEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQEnterEvent(h unsafe.Pointer) *QEnterEvent { return bindQEnterEvent(uintptr(h)) }
func UnsafeNewQEnterEventOwned(h unsafe.Pointer) *QEnterEvent { return newownedQEnterEvent(uintptr(h)) }

func (this *QEnterEvent) UnsafeRelease() {
	this.h = 0
	if this.QSinglePointEvent != nil {
		this.QSinglePointEvent.UnsafeRelease()
	}
}

func NewQEnterEvent(localPos *QPointF, scenePos *QPointF, globalPos *QPointF) *QEnterEvent {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_new, uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()))
	__o := newownedQEnterEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	return __o
}

func NewQEnterEvent2(localPos *QPointF, scenePos *QPointF, globalPos *QPointF, device *QPointingDevice) *QEnterEvent {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_new2, uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(device.UnsafePointer()))
	__o := newownedQEnterEvent(__r)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(device)
	return __o
}

func (this *QEnterEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QEnterEvent) GlobalPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_globalPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QEnterEvent) X() int {
	__r0, _, _ := purego.SyscallN(addr_QEnterEvent_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEnterEvent) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QEnterEvent_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEnterEvent) GlobalX() int {
	__r0, _, _ := purego.SyscallN(addr_QEnterEvent_globalX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEnterEvent) GlobalY() int {
	__r0, _, _ := purego.SyscallN(addr_QEnterEvent_globalY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEnterEvent) LocalPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_localPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEnterEvent) WindowPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_windowPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEnterEvent) ScreenPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEnterEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func QEnterEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QEnterEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QEnterEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QEnterEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QEnterEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QEnterEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
