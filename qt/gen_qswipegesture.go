package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSwipeGesture_chain uintptr
	addr_QSwipeGesture_new uintptr
	addr_QSwipeGesture_new2 uintptr
	addr_QSwipeGesture_tr uintptr
	addr_QSwipeGesture_horizontalDirection uintptr
	addr_QSwipeGesture_verticalDirection uintptr
	pg_QSwipeGesture_swipeAngle func(uintptr) float64
	pg_QSwipeGesture_setSwipeAngle func(uintptr, float64)
	addr_QSwipeGesture_connect uintptr
	addr_QSwipeGesture_connect2 uintptr
	addr_QSwipeGesture_sender uintptr
	addr_QSwipeGesture_senderSignalIndex uintptr
	addr_QSwipeGesture_receivers uintptr
	addr_QSwipeGesture_isSignalConnected uintptr
	addr_QSwipeGesture_tr2 uintptr
	addr_QSwipeGesture_tr3 uintptr
	addr_QSwipeGesture_connect3 uintptr
	addr_QSwipeGesture_connect4 uintptr
	addr_QSwipeGesture_metaObject_default uintptr
	addr_QSwipeGesture_event_default uintptr
	addr_QSwipeGesture_eventFilter_default uintptr
	addr_QSwipeGesture_timerEvent_default uintptr
	addr_QSwipeGesture_childEvent_default uintptr
	addr_QSwipeGesture_customEvent_default uintptr
	addr_QSwipeGesture_connectNotify_default uintptr
	addr_QSwipeGesture_disconnectNotify_default uintptr
	pg_QSwipeGesture_connect_destroyed func(uintptr, int64)
	pg_QSwipeGesture_connect_destroyed2 func(uintptr, int64)
	pg_QSwipeGesture_connect_objectNameChanged func(uintptr, int64)
	pg_QSwipeGesture__disconnect func(int64)
	pg_QSwipeGesture__set_override func(uintptr, int32)
	pg_QSwipeGesture_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSwipeGesture_chain, "QSwipeGesture_chain")
	symaddr(&addr_QSwipeGesture_new, "QSwipeGesture_new")
	symaddr(&addr_QSwipeGesture_new2, "QSwipeGesture_new2")
	symaddr(&addr_QSwipeGesture_tr, "QSwipeGesture_tr")
	symaddr(&addr_QSwipeGesture_horizontalDirection, "QSwipeGesture_horizontalDirection")
	symaddr(&addr_QSwipeGesture_verticalDirection, "QSwipeGesture_verticalDirection")
	reg(&pg_QSwipeGesture_swipeAngle, "QSwipeGesture_swipeAngle")
	reg(&pg_QSwipeGesture_setSwipeAngle, "QSwipeGesture_setSwipeAngle")
	symaddr(&addr_QSwipeGesture_connect, "QSwipeGesture_connect")
	symaddr(&addr_QSwipeGesture_connect2, "QSwipeGesture_connect2")
	symaddr(&addr_QSwipeGesture_sender, "QSwipeGesture_sender")
	symaddr(&addr_QSwipeGesture_senderSignalIndex, "QSwipeGesture_senderSignalIndex")
	symaddr(&addr_QSwipeGesture_receivers, "QSwipeGesture_receivers")
	symaddr(&addr_QSwipeGesture_isSignalConnected, "QSwipeGesture_isSignalConnected")
	symaddr(&addr_QSwipeGesture_tr2, "QSwipeGesture_tr2")
	symaddr(&addr_QSwipeGesture_tr3, "QSwipeGesture_tr3")
	symaddr(&addr_QSwipeGesture_connect3, "QSwipeGesture_connect3")
	symaddr(&addr_QSwipeGesture_connect4, "QSwipeGesture_connect4")
	symaddr(&addr_QSwipeGesture_metaObject_default, "QSwipeGesture_metaObject_default")
	symaddr(&addr_QSwipeGesture_event_default, "QSwipeGesture_event_default")
	symaddr(&addr_QSwipeGesture_eventFilter_default, "QSwipeGesture_eventFilter_default")
	symaddr(&addr_QSwipeGesture_timerEvent_default, "QSwipeGesture_timerEvent_default")
	symaddr(&addr_QSwipeGesture_childEvent_default, "QSwipeGesture_childEvent_default")
	symaddr(&addr_QSwipeGesture_customEvent_default, "QSwipeGesture_customEvent_default")
	symaddr(&addr_QSwipeGesture_connectNotify_default, "QSwipeGesture_connectNotify_default")
	symaddr(&addr_QSwipeGesture_disconnectNotify_default, "QSwipeGesture_disconnectNotify_default")
	reg(&pg_QSwipeGesture_connect_destroyed, "QSwipeGesture_connect_destroyed")
	reg(&pg_QSwipeGesture_connect_destroyed2, "QSwipeGesture_connect_destroyed2")
	reg(&pg_QSwipeGesture_connect_objectNameChanged, "QSwipeGesture_connect_objectNameChanged")
	reg(&pg_QSwipeGesture__disconnect, "QSwipeGesture__disconnect")
	reg(&pg_QSwipeGesture__set_override, "QSwipeGesture__set_override")
	reg(&pg_QSwipeGesture_delete, "QSwipeGesture_delete")
}

type QSwipeGesture struct {
	h uintptr
	*QGesture
}

func bindQSwipeGesture(h uintptr) *QSwipeGesture {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QSwipeGesture_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSwipeGesture{h: h, QGesture: &QGesture{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQSwipeGesture(h uintptr) *QSwipeGesture { return bindQSwipeGesture(h) }

func (this *QSwipeGesture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSwipeGesture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSwipeGesture(h unsafe.Pointer) *QSwipeGesture { return bindQSwipeGesture(uintptr(h)) }
func UnsafeNewQSwipeGestureOwned(h unsafe.Pointer) *QSwipeGesture { return newownedQSwipeGesture(uintptr(h)) }

func (this *QSwipeGesture) UnsafeRelease() {
	this.h = 0
	if this.QGesture != nil {
		this.QGesture.UnsafeRelease()
	}
}

func NewQSwipeGesture() *QSwipeGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_new, uintptr(__goid))
	__o := newownedQSwipeGesture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSwipeGesture_virtIdx, QSwipeGesture_enable)
	return __o
}

func NewQSwipeGesture2(parent *QObject) *QSwipeGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSwipeGesture(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSwipeGesture_virtIdx, QSwipeGesture_enable)
	return __o
}

func QSwipeGesture_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSwipeGesture) HorizontalDirection() QSwipeGesture__SwipeDirection {
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_horizontalDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSwipeGesture__SwipeDirection(__r)
}

func (this *QSwipeGesture) VerticalDirection() QSwipeGesture__SwipeDirection {
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_verticalDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSwipeGesture__SwipeDirection(__r)
}

func (this *QSwipeGesture) SwipeAngle() float64 {
	__r := pg_QSwipeGesture_swipeAngle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSwipeGesture) SetSwipeAngle(value float64) {
	pg_QSwipeGesture_setSwipeAngle(this.h, value)
	runtime.KeepAlive(this)
}

func QSwipeGesture_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSwipeGesture_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSwipeGesture) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSwipeGesture) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSwipeGesture) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSwipeGesture) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSwipeGesture_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSwipeGesture_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSwipeGesture_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSwipeGesture_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSwipeGesture) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSwipeGesture_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSwipeGesture) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSwipeGesture) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSwipeGesture_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSwipeGesture) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSwipeGesture_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSwipeGesture) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSwipeGesture_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSwipeGesture) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSwipeGesture_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSwipeGesture) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSwipeGesture_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSwipeGesture) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSwipeGesture_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSwipeGesture) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSwipeGesture_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSwipeGesture) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSwipeGesture_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSwipeGesture) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSwipeGesture_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSwipeGesture) Disconnect(handle int64) {
	pg_QSwipeGesture__disconnect(handle)
	unregisterCallback(handle)
}

func QSwipeGesture_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QSwipeGesture_enable(h uintptr, idx int32) { pg_QSwipeGesture__set_override(h, idx) }

func (this *QSwipeGesture) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSwipeGesture) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSwipeGesture) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSwipeGesture) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSwipeGesture) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSwipeGesture) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSwipeGesture) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSwipeGesture) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSwipeGesture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSwipeGesture_delete(this.h)
	}
	this.UnsafeRelease()
}
