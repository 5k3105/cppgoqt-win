package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QNativeGestureEvent_chain uintptr
	pg_QNativeGestureEvent_new func(int32, uintptr, uintptr, uintptr, uintptr, float64, int64, int64) uintptr
	pg_QNativeGestureEvent_new2 func(int32, uintptr, int32, uintptr, uintptr, uintptr, float64, uintptr) uintptr
	pg_QNativeGestureEvent_new3 func(int32, uintptr, int32, uintptr, uintptr, uintptr, float64, uintptr, int64) uintptr
	addr_QNativeGestureEvent_gestureType uintptr
	addr_QNativeGestureEvent_fingerCount uintptr
	pg_QNativeGestureEvent_value func(uintptr) float64
	addr_QNativeGestureEvent_delta uintptr
	addr_QNativeGestureEvent_pos uintptr
	addr_QNativeGestureEvent_globalPos uintptr
	addr_QNativeGestureEvent_localPos uintptr
	addr_QNativeGestureEvent_windowPos uintptr
	addr_QNativeGestureEvent_screenPos uintptr
	addr_QNativeGestureEvent_registerEventType uintptr
	addr_QNativeGestureEvent_registerEventType2 uintptr
	pg_QNativeGestureEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QNativeGestureEvent_chain, "QNativeGestureEvent_chain")
	reg(&pg_QNativeGestureEvent_new, "QNativeGestureEvent_new")
	reg(&pg_QNativeGestureEvent_new2, "QNativeGestureEvent_new2")
	reg(&pg_QNativeGestureEvent_new3, "QNativeGestureEvent_new3")
	symaddr(&addr_QNativeGestureEvent_gestureType, "QNativeGestureEvent_gestureType")
	symaddr(&addr_QNativeGestureEvent_fingerCount, "QNativeGestureEvent_fingerCount")
	reg(&pg_QNativeGestureEvent_value, "QNativeGestureEvent_value")
	symaddr(&addr_QNativeGestureEvent_delta, "QNativeGestureEvent_delta")
	symaddr(&addr_QNativeGestureEvent_pos, "QNativeGestureEvent_pos")
	symaddr(&addr_QNativeGestureEvent_globalPos, "QNativeGestureEvent_globalPos")
	symaddr(&addr_QNativeGestureEvent_localPos, "QNativeGestureEvent_localPos")
	symaddr(&addr_QNativeGestureEvent_windowPos, "QNativeGestureEvent_windowPos")
	symaddr(&addr_QNativeGestureEvent_screenPos, "QNativeGestureEvent_screenPos")
	symaddr(&addr_QNativeGestureEvent_registerEventType, "QNativeGestureEvent_registerEventType")
	symaddr(&addr_QNativeGestureEvent_registerEventType2, "QNativeGestureEvent_registerEventType2")
	reg(&pg_QNativeGestureEvent_delete, "QNativeGestureEvent_delete")
}

type QNativeGestureEvent struct {
	h uintptr
	*QSinglePointEvent
}

func bindQNativeGestureEvent(h uintptr) *QNativeGestureEvent {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QNativeGestureEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QNativeGestureEvent{h: h, QSinglePointEvent: &QSinglePointEvent{h: __c[0], QPointerEvent: &QPointerEvent{h: __c[1], QInputEvent: &QInputEvent{h: __c[2], QEvent: &QEvent{h: __c[3]}}}}}
}

func newownedQNativeGestureEvent(h uintptr) *QNativeGestureEvent {
	o := bindQNativeGestureEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QNativeGestureEvent).Delete)
	}
	return o
}

func (this *QNativeGestureEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QNativeGestureEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQNativeGestureEvent(h unsafe.Pointer) *QNativeGestureEvent { return bindQNativeGestureEvent(uintptr(h)) }
func UnsafeNewQNativeGestureEventOwned(h unsafe.Pointer) *QNativeGestureEvent { return newownedQNativeGestureEvent(uintptr(h)) }

func (this *QNativeGestureEvent) UnsafeRelease() {
	this.h = 0
	if this.QSinglePointEvent != nil {
		this.QSinglePointEvent.UnsafeRelease()
	}
}

func NewQNativeGestureEvent(type_ NativeGestureType, dev *QPointingDevice, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, sequenceId uint64, intArgument uint64) *QNativeGestureEvent {
	__o := newownedQNativeGestureEvent(pg_QNativeGestureEvent_new(int32(type_), uintptr(dev.UnsafePointer()), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), value, int64(sequenceId), int64(intArgument)))
	runtime.KeepAlive(dev)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	return __o
}

func NewQNativeGestureEvent2(type_ NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF) *QNativeGestureEvent {
	__o := newownedQNativeGestureEvent(pg_QNativeGestureEvent_new2(int32(type_), uintptr(dev.UnsafePointer()), int32(fingerCount), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), value, uintptr(delta.UnsafePointer())))
	runtime.KeepAlive(dev)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(delta)
	return __o
}

func NewQNativeGestureEvent3(type_ NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF, sequenceId uint64) *QNativeGestureEvent {
	__o := newownedQNativeGestureEvent(pg_QNativeGestureEvent_new3(int32(type_), uintptr(dev.UnsafePointer()), int32(fingerCount), uintptr(localPos.UnsafePointer()), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), value, uintptr(delta.UnsafePointer()), int64(sequenceId)))
	runtime.KeepAlive(dev)
	runtime.KeepAlive(localPos)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(delta)
	return __o
}

func (this *QNativeGestureEvent) GestureType() NativeGestureType {
	__r0, _, _ := purego.SyscallN(addr_QNativeGestureEvent_gestureType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return NativeGestureType(__r)
}

func (this *QNativeGestureEvent) FingerCount() int {
	__r0, _, _ := purego.SyscallN(addr_QNativeGestureEvent_fingerCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QNativeGestureEvent) Value() float64 {
	__r := pg_QNativeGestureEvent_value(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QNativeGestureEvent) Delta() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QNativeGestureEvent_delta, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QNativeGestureEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QNativeGestureEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QNativeGestureEvent) GlobalPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QNativeGestureEvent_globalPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QNativeGestureEvent) LocalPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QNativeGestureEvent_localPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QNativeGestureEvent) WindowPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QNativeGestureEvent_windowPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QNativeGestureEvent) ScreenPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QNativeGestureEvent_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func QNativeGestureEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QNativeGestureEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QNativeGestureEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QNativeGestureEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QNativeGestureEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QNativeGestureEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
