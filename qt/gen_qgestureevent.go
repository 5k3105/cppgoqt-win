package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGestureEvent_base func(uintptr) uintptr
	addr_QGestureEvent_new uintptr
	addr_QGestureEvent_new2 uintptr
	addr_QGestureEvent_gestures uintptr
	addr_QGestureEvent_gesture uintptr
	addr_QGestureEvent_activeGestures uintptr
	addr_QGestureEvent_canceledGestures uintptr
	addr_QGestureEvent_setAccepted uintptr
	addr_QGestureEvent_accept uintptr
	addr_QGestureEvent_ignore uintptr
	addr_QGestureEvent_isAccepted uintptr
	addr_QGestureEvent_setAccepted2 uintptr
	addr_QGestureEvent_accept2 uintptr
	addr_QGestureEvent_ignore2 uintptr
	addr_QGestureEvent_isAccepted2 uintptr
	addr_QGestureEvent_setWidget uintptr
	addr_QGestureEvent_widget uintptr
	addr_QGestureEvent_mapToGraphicsScene uintptr
	addr_QGestureEvent_OperatorAssign uintptr
	addr_QGestureEvent_registerEventType uintptr
	addr_QGestureEvent_registerEventType2 uintptr
	pg_QGestureEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QGestureEvent_base, "QGestureEvent_base")
	symaddr(&addr_QGestureEvent_new, "QGestureEvent_new")
	symaddr(&addr_QGestureEvent_new2, "QGestureEvent_new2")
	symaddr(&addr_QGestureEvent_gestures, "QGestureEvent_gestures")
	symaddr(&addr_QGestureEvent_gesture, "QGestureEvent_gesture")
	symaddr(&addr_QGestureEvent_activeGestures, "QGestureEvent_activeGestures")
	symaddr(&addr_QGestureEvent_canceledGestures, "QGestureEvent_canceledGestures")
	symaddr(&addr_QGestureEvent_setAccepted, "QGestureEvent_setAccepted")
	symaddr(&addr_QGestureEvent_accept, "QGestureEvent_accept")
	symaddr(&addr_QGestureEvent_ignore, "QGestureEvent_ignore")
	symaddr(&addr_QGestureEvent_isAccepted, "QGestureEvent_isAccepted")
	symaddr(&addr_QGestureEvent_setAccepted2, "QGestureEvent_setAccepted2")
	symaddr(&addr_QGestureEvent_accept2, "QGestureEvent_accept2")
	symaddr(&addr_QGestureEvent_ignore2, "QGestureEvent_ignore2")
	symaddr(&addr_QGestureEvent_isAccepted2, "QGestureEvent_isAccepted2")
	symaddr(&addr_QGestureEvent_setWidget, "QGestureEvent_setWidget")
	symaddr(&addr_QGestureEvent_widget, "QGestureEvent_widget")
	symaddr(&addr_QGestureEvent_mapToGraphicsScene, "QGestureEvent_mapToGraphicsScene")
	symaddr(&addr_QGestureEvent_OperatorAssign, "QGestureEvent_OperatorAssign")
	symaddr(&addr_QGestureEvent_registerEventType, "QGestureEvent_registerEventType")
	symaddr(&addr_QGestureEvent_registerEventType2, "QGestureEvent_registerEventType2")
	reg(&pg_QGestureEvent_delete, "QGestureEvent_delete")
}

type QGestureEvent struct {
	h uintptr
	*QEvent
}

func bindQGestureEvent(h uintptr) *QGestureEvent {
	if h == 0 {
		return nil
	}
	return &QGestureEvent{h: h, QEvent: bindQEvent(pg_QGestureEvent_base(h))}
}

func newownedQGestureEvent(h uintptr) *QGestureEvent {
	o := bindQGestureEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGestureEvent).Delete)
	}
	return o
}

func (this *QGestureEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGestureEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGestureEvent(h unsafe.Pointer) *QGestureEvent { return bindQGestureEvent(uintptr(h)) }
func UnsafeNewQGestureEventOwned(h unsafe.Pointer) *QGestureEvent { return newownedQGestureEvent(uintptr(h)) }

func (this *QGestureEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQGestureEvent(gestures []*QGesture) *QGestureEvent {
	gestures_arr := make([]uintptr, len(gestures))
	for __i, __e := range gestures {
		gestures_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var gestures_p uintptr
	if len(gestures_arr) > 0 {
		gestures_p = uintptr(unsafe.Pointer(&gestures_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_new, gestures_p, uintptr(len(gestures)))
	__o := newownedQGestureEvent(__r)
	runtime.KeepAlive(gestures)
	runtime.KeepAlive(gestures_arr)
	return __o
}

func NewQGestureEvent2(arg1 *QGestureEvent) *QGestureEvent {
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQGestureEvent(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QGestureEvent) Gestures() []*QGesture {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_gestures, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGesture, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGesture(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGestureEvent) Gesture(type_ GestureType) *QGesture {
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_gesture, this.h, uintptr(type_))
	runtime.KeepAlive(this)
	return bindQGesture(__r)
}

func (this *QGestureEvent) ActiveGestures() []*QGesture {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_activeGestures, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGesture, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGesture(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGestureEvent) CanceledGestures() []*QGesture {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_canceledGestures, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGesture, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGesture(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGestureEvent) SetAccepted(arg1 *QGesture, arg2 bool) {
	purego.SyscallN(addr_QGestureEvent_setAccepted, this.h, uintptr(arg1.UnsafePointer()), uintptr(b2i(arg2)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGestureEvent) Accept(arg1 *QGesture) {
	purego.SyscallN(addr_QGestureEvent_accept, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGestureEvent) Ignore(arg1 *QGesture) {
	purego.SyscallN(addr_QGestureEvent_ignore, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGestureEvent) IsAccepted(arg1 *QGesture) bool {
	__r0, _, _ := purego.SyscallN(addr_QGestureEvent_isAccepted, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QGestureEvent) SetAccepted2(arg1 GestureType, arg2 bool) {
	purego.SyscallN(addr_QGestureEvent_setAccepted2, this.h, uintptr(arg1), uintptr(b2i(arg2)))
	runtime.KeepAlive(this)
}

func (this *QGestureEvent) AcceptWithQtGestureType(arg1 GestureType) {
	purego.SyscallN(addr_QGestureEvent_accept2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QGestureEvent) IgnoreWithQtGestureType(arg1 GestureType) {
	purego.SyscallN(addr_QGestureEvent_ignore2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QGestureEvent) IsAcceptedWithQtGestureType(arg1 GestureType) bool {
	__r0, _, _ := purego.SyscallN(addr_QGestureEvent_isAccepted2, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGestureEvent) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QGestureEvent_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QGestureEvent) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QGestureEvent) MapToGraphicsScene(gesturePoint *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_mapToGraphicsScene, this.h, uintptr(gesturePoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(gesturePoint)
	return newownedQPointF(__r)
}

func (this *QGestureEvent) OperatorAssign(arg1 *QGestureEvent) *QGestureEvent {
	__r, _, _ := purego.SyscallN(addr_QGestureEvent_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQGestureEvent(__r)
}

func QGestureEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGestureEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGestureEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGestureEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGestureEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGestureEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
