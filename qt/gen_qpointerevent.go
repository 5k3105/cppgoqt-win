package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPointerEvent_chain uintptr
	addr_QPointerEvent_new uintptr
	addr_QPointerEvent_new2 uintptr
	addr_QPointerEvent_new3 uintptr
	addr_QPointerEvent_pointingDevice uintptr
	addr_QPointerEvent_pointerType uintptr
	addr_QPointerEvent_pointCount uintptr
	addr_QPointerEvent_point uintptr
	addr_QPointerEvent_points uintptr
	addr_QPointerEvent_pointById uintptr
	addr_QPointerEvent_allPointsGrabbed uintptr
	addr_QPointerEvent_isBeginEvent uintptr
	addr_QPointerEvent_isUpdateEvent uintptr
	addr_QPointerEvent_isEndEvent uintptr
	addr_QPointerEvent_allPointsAccepted uintptr
	addr_QPointerEvent_exclusiveGrabber uintptr
	addr_QPointerEvent_setExclusiveGrabber uintptr
	addr_QPointerEvent_clearPassiveGrabbers uintptr
	addr_QPointerEvent_addPassiveGrabber uintptr
	addr_QPointerEvent_removePassiveGrabber uintptr
	addr_QPointerEvent_registerEventType uintptr
	addr_QPointerEvent_registerEventType2 uintptr
	pg_QPointerEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPointerEvent_chain, "QPointerEvent_chain")
	symaddr(&addr_QPointerEvent_new, "QPointerEvent_new")
	symaddr(&addr_QPointerEvent_new2, "QPointerEvent_new2")
	symaddr(&addr_QPointerEvent_new3, "QPointerEvent_new3")
	symaddr(&addr_QPointerEvent_pointingDevice, "QPointerEvent_pointingDevice")
	symaddr(&addr_QPointerEvent_pointerType, "QPointerEvent_pointerType")
	symaddr(&addr_QPointerEvent_pointCount, "QPointerEvent_pointCount")
	symaddr(&addr_QPointerEvent_point, "QPointerEvent_point")
	symaddr(&addr_QPointerEvent_points, "QPointerEvent_points")
	symaddr(&addr_QPointerEvent_pointById, "QPointerEvent_pointById")
	symaddr(&addr_QPointerEvent_allPointsGrabbed, "QPointerEvent_allPointsGrabbed")
	symaddr(&addr_QPointerEvent_isBeginEvent, "QPointerEvent_isBeginEvent")
	symaddr(&addr_QPointerEvent_isUpdateEvent, "QPointerEvent_isUpdateEvent")
	symaddr(&addr_QPointerEvent_isEndEvent, "QPointerEvent_isEndEvent")
	symaddr(&addr_QPointerEvent_allPointsAccepted, "QPointerEvent_allPointsAccepted")
	symaddr(&addr_QPointerEvent_exclusiveGrabber, "QPointerEvent_exclusiveGrabber")
	symaddr(&addr_QPointerEvent_setExclusiveGrabber, "QPointerEvent_setExclusiveGrabber")
	symaddr(&addr_QPointerEvent_clearPassiveGrabbers, "QPointerEvent_clearPassiveGrabbers")
	symaddr(&addr_QPointerEvent_addPassiveGrabber, "QPointerEvent_addPassiveGrabber")
	symaddr(&addr_QPointerEvent_removePassiveGrabber, "QPointerEvent_removePassiveGrabber")
	symaddr(&addr_QPointerEvent_registerEventType, "QPointerEvent_registerEventType")
	symaddr(&addr_QPointerEvent_registerEventType2, "QPointerEvent_registerEventType2")
	reg(&pg_QPointerEvent_delete, "QPointerEvent_delete")
}

type QPointerEvent struct {
	h uintptr
	*QInputEvent
}

func bindQPointerEvent(h uintptr) *QPointerEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPointerEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPointerEvent{h: h, QInputEvent: &QInputEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQPointerEvent(h uintptr) *QPointerEvent {
	o := bindQPointerEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPointerEvent).Delete)
	}
	return o
}

func (this *QPointerEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPointerEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPointerEvent(h unsafe.Pointer) *QPointerEvent { return bindQPointerEvent(uintptr(h)) }
func UnsafeNewQPointerEventOwned(h unsafe.Pointer) *QPointerEvent { return newownedQPointerEvent(uintptr(h)) }

func (this *QPointerEvent) UnsafeRelease() {
	this.h = 0
	if this.QInputEvent != nil {
		this.QInputEvent.UnsafeRelease()
	}
}

func NewQPointerEvent(type_ QEvent__Type, dev *QPointingDevice) *QPointerEvent {
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_new, uintptr(type_), uintptr(dev.UnsafePointer()))
	__o := newownedQPointerEvent(__r)
	runtime.KeepAlive(dev)
	return __o
}

func NewQPointerEvent2(type_ QEvent__Type, dev *QPointingDevice, modifiers KeyboardModifier) *QPointerEvent {
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_new2, uintptr(type_), uintptr(dev.UnsafePointer()), uintptr(modifiers))
	__o := newownedQPointerEvent(__r)
	runtime.KeepAlive(dev)
	return __o
}

func NewQPointerEvent3(type_ QEvent__Type, dev *QPointingDevice, modifiers KeyboardModifier, points []*QEventPoint) *QPointerEvent {
	points_arr := make([]uintptr, len(points))
	for __i, __e := range points {
		points_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var points_p uintptr
	if len(points_arr) > 0 {
		points_p = uintptr(unsafe.Pointer(&points_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_new3, uintptr(type_), uintptr(dev.UnsafePointer()), uintptr(modifiers), points_p, uintptr(len(points)))
	__o := newownedQPointerEvent(__r)
	runtime.KeepAlive(dev)
	runtime.KeepAlive(points)
	runtime.KeepAlive(points_arr)
	return __o
}

func (this *QPointerEvent) PointingDevice() *QPointingDevice {
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_pointingDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPointingDevice(__r)
}

func (this *QPointerEvent) PointerType() QPointingDevice__PointerType {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_pointerType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPointingDevice__PointerType(__r)
}

func (this *QPointerEvent) PointCount() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_pointCount, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPointerEvent) Point(i int64) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_point, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQEventPoint(__r)
}

func (this *QPointerEvent) Points() []*QEventPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_points, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QEventPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQEventPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPointerEvent) PointById(id int) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_pointById, this.h, uintptr(id))
	runtime.KeepAlive(this)
	return bindQEventPoint(__r)
}

func (this *QPointerEvent) AllPointsGrabbed() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_allPointsGrabbed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointerEvent) IsBeginEvent() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_isBeginEvent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointerEvent) IsUpdateEvent() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_isUpdateEvent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointerEvent) IsEndEvent() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_isEndEvent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointerEvent) AllPointsAccepted() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_allPointsAccepted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointerEvent) ExclusiveGrabber(point *QEventPoint) *QObject {
	__r, _, _ := purego.SyscallN(addr_QPointerEvent_exclusiveGrabber, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return bindQObject(__r)
}

func (this *QPointerEvent) SetExclusiveGrabber(point *QEventPoint, exclusiveGrabber *QObject) {
	purego.SyscallN(addr_QPointerEvent_setExclusiveGrabber, this.h, uintptr(point.UnsafePointer()), uintptr(exclusiveGrabber.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	runtime.KeepAlive(exclusiveGrabber)
}

func (this *QPointerEvent) ClearPassiveGrabbers(point *QEventPoint) {
	purego.SyscallN(addr_QPointerEvent_clearPassiveGrabbers, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
}

func (this *QPointerEvent) AddPassiveGrabber(point *QEventPoint, grabber *QObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_addPassiveGrabber, this.h, uintptr(point.UnsafePointer()), uintptr(grabber.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	runtime.KeepAlive(grabber)
	return __r != 0
}

func (this *QPointerEvent) RemovePassiveGrabber(point *QEventPoint, grabber *QObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_removePassiveGrabber, this.h, uintptr(point.UnsafePointer()), uintptr(grabber.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	runtime.KeepAlive(grabber)
	return __r != 0
}

func QPointerEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QPointerEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QPointerEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QPointerEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPointerEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
