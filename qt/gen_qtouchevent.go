package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTouchEvent_chain uintptr
	addr_QTouchEvent_new uintptr
	addr_QTouchEvent_new2 uintptr
	addr_QTouchEvent_new3 uintptr
	addr_QTouchEvent_new4 uintptr
	addr_QTouchEvent_new5 uintptr
	addr_QTouchEvent_new6 uintptr
	addr_QTouchEvent_target uintptr
	addr_QTouchEvent_touchPointStates uintptr
	addr_QTouchEvent_touchPoints uintptr
	addr_QTouchEvent_registerEventType uintptr
	addr_QTouchEvent_registerEventType2 uintptr
	pg_QTouchEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTouchEvent_chain, "QTouchEvent_chain")
	symaddr(&addr_QTouchEvent_new, "QTouchEvent_new")
	symaddr(&addr_QTouchEvent_new2, "QTouchEvent_new2")
	symaddr(&addr_QTouchEvent_new3, "QTouchEvent_new3")
	symaddr(&addr_QTouchEvent_new4, "QTouchEvent_new4")
	symaddr(&addr_QTouchEvent_new5, "QTouchEvent_new5")
	symaddr(&addr_QTouchEvent_new6, "QTouchEvent_new6")
	symaddr(&addr_QTouchEvent_target, "QTouchEvent_target")
	symaddr(&addr_QTouchEvent_touchPointStates, "QTouchEvent_touchPointStates")
	symaddr(&addr_QTouchEvent_touchPoints, "QTouchEvent_touchPoints")
	symaddr(&addr_QTouchEvent_registerEventType, "QTouchEvent_registerEventType")
	symaddr(&addr_QTouchEvent_registerEventType2, "QTouchEvent_registerEventType2")
	reg(&pg_QTouchEvent_delete, "QTouchEvent_delete")
}

type QTouchEvent struct {
	h uintptr
	*QPointerEvent
}

func bindQTouchEvent(h uintptr) *QTouchEvent {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QTouchEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTouchEvent{h: h, QPointerEvent: &QPointerEvent{h: __c[0], QInputEvent: &QInputEvent{h: __c[1], QEvent: &QEvent{h: __c[2]}}}}
}

func newownedQTouchEvent(h uintptr) *QTouchEvent {
	o := bindQTouchEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTouchEvent).Delete)
	}
	return o
}

func (this *QTouchEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTouchEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTouchEvent(h unsafe.Pointer) *QTouchEvent { return bindQTouchEvent(uintptr(h)) }
func UnsafeNewQTouchEventOwned(h unsafe.Pointer) *QTouchEvent { return newownedQTouchEvent(uintptr(h)) }

func (this *QTouchEvent) UnsafeRelease() {
	this.h = 0
	if this.QPointerEvent != nil {
		this.QPointerEvent.UnsafeRelease()
	}
}

func NewQTouchEvent(eventType QEvent__Type) *QTouchEvent {
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_new, uintptr(eventType))
	__o := newownedQTouchEvent(__r)
	return __o
}

func NewQTouchEvent2(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates QEventPoint__State) *QTouchEvent {
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_new2, uintptr(eventType), uintptr(device.UnsafePointer()), uintptr(modifiers), uintptr(touchPointStates))
	__o := newownedQTouchEvent(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQTouchEvent3(eventType QEvent__Type, device *QPointingDevice) *QTouchEvent {
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_new3, uintptr(eventType), uintptr(device.UnsafePointer()))
	__o := newownedQTouchEvent(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQTouchEvent4(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier) *QTouchEvent {
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_new4, uintptr(eventType), uintptr(device.UnsafePointer()), uintptr(modifiers))
	__o := newownedQTouchEvent(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQTouchEvent5(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPoints []*QEventPoint) *QTouchEvent {
	touchPoints_arr := make([]uintptr, len(touchPoints))
	for __i, __e := range touchPoints {
		touchPoints_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var touchPoints_p uintptr
	if len(touchPoints_arr) > 0 {
		touchPoints_p = uintptr(unsafe.Pointer(&touchPoints_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_new5, uintptr(eventType), uintptr(device.UnsafePointer()), uintptr(modifiers), touchPoints_p, uintptr(len(touchPoints)))
	__o := newownedQTouchEvent(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(touchPoints)
	runtime.KeepAlive(touchPoints_arr)
	return __o
}

func NewQTouchEvent6(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates QEventPoint__State, touchPoints []*QEventPoint) *QTouchEvent {
	touchPoints_arr := make([]uintptr, len(touchPoints))
	for __i, __e := range touchPoints {
		touchPoints_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var touchPoints_p uintptr
	if len(touchPoints_arr) > 0 {
		touchPoints_p = uintptr(unsafe.Pointer(&touchPoints_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_new6, uintptr(eventType), uintptr(device.UnsafePointer()), uintptr(modifiers), uintptr(touchPointStates), touchPoints_p, uintptr(len(touchPoints)))
	__o := newownedQTouchEvent(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(touchPoints)
	runtime.KeepAlive(touchPoints_arr)
	return __o
}

func (this *QTouchEvent) Target() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_target, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTouchEvent) TouchPointStates() QEventPoint__State {
	__r0, _, _ := purego.SyscallN(addr_QTouchEvent_touchPointStates, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QEventPoint__State(__r)
}

func (this *QTouchEvent) TouchPoints() []*QEventPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTouchEvent_touchPoints, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QEventPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQEventPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTouchEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QTouchEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QTouchEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QTouchEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QTouchEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTouchEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
