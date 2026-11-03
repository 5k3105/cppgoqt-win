package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPanGesture_chain uintptr
	addr_QPanGesture_new uintptr
	addr_QPanGesture_new2 uintptr
	addr_QPanGesture_tr uintptr
	addr_QPanGesture_lastOffset uintptr
	addr_QPanGesture_offset uintptr
	addr_QPanGesture_delta uintptr
	pg_QPanGesture_acceleration func(uintptr) float64
	addr_QPanGesture_setLastOffset uintptr
	addr_QPanGesture_setOffset uintptr
	pg_QPanGesture_setAcceleration func(uintptr, float64)
	addr_QPanGesture_connect uintptr
	addr_QPanGesture_connect2 uintptr
	addr_QPanGesture_sender uintptr
	addr_QPanGesture_senderSignalIndex uintptr
	addr_QPanGesture_receivers uintptr
	addr_QPanGesture_isSignalConnected uintptr
	addr_QPanGesture_tr2 uintptr
	addr_QPanGesture_tr3 uintptr
	addr_QPanGesture_connect3 uintptr
	addr_QPanGesture_connect4 uintptr
	addr_QPanGesture_metaObject_default uintptr
	addr_QPanGesture_event_default uintptr
	addr_QPanGesture_eventFilter_default uintptr
	addr_QPanGesture_timerEvent_default uintptr
	addr_QPanGesture_childEvent_default uintptr
	addr_QPanGesture_customEvent_default uintptr
	addr_QPanGesture_connectNotify_default uintptr
	addr_QPanGesture_disconnectNotify_default uintptr
	pg_QPanGesture_connect_destroyed func(uintptr, int64)
	pg_QPanGesture_connect_destroyed2 func(uintptr, int64)
	pg_QPanGesture_connect_objectNameChanged func(uintptr, int64)
	pg_QPanGesture__disconnect func(int64)
	pg_QPanGesture__set_override func(uintptr, int32)
	pg_QPanGesture_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPanGesture_chain, "QPanGesture_chain")
	symaddr(&addr_QPanGesture_new, "QPanGesture_new")
	symaddr(&addr_QPanGesture_new2, "QPanGesture_new2")
	symaddr(&addr_QPanGesture_tr, "QPanGesture_tr")
	symaddr(&addr_QPanGesture_lastOffset, "QPanGesture_lastOffset")
	symaddr(&addr_QPanGesture_offset, "QPanGesture_offset")
	symaddr(&addr_QPanGesture_delta, "QPanGesture_delta")
	reg(&pg_QPanGesture_acceleration, "QPanGesture_acceleration")
	symaddr(&addr_QPanGesture_setLastOffset, "QPanGesture_setLastOffset")
	symaddr(&addr_QPanGesture_setOffset, "QPanGesture_setOffset")
	reg(&pg_QPanGesture_setAcceleration, "QPanGesture_setAcceleration")
	symaddr(&addr_QPanGesture_connect, "QPanGesture_connect")
	symaddr(&addr_QPanGesture_connect2, "QPanGesture_connect2")
	symaddr(&addr_QPanGesture_sender, "QPanGesture_sender")
	symaddr(&addr_QPanGesture_senderSignalIndex, "QPanGesture_senderSignalIndex")
	symaddr(&addr_QPanGesture_receivers, "QPanGesture_receivers")
	symaddr(&addr_QPanGesture_isSignalConnected, "QPanGesture_isSignalConnected")
	symaddr(&addr_QPanGesture_tr2, "QPanGesture_tr2")
	symaddr(&addr_QPanGesture_tr3, "QPanGesture_tr3")
	symaddr(&addr_QPanGesture_connect3, "QPanGesture_connect3")
	symaddr(&addr_QPanGesture_connect4, "QPanGesture_connect4")
	symaddr(&addr_QPanGesture_metaObject_default, "QPanGesture_metaObject_default")
	symaddr(&addr_QPanGesture_event_default, "QPanGesture_event_default")
	symaddr(&addr_QPanGesture_eventFilter_default, "QPanGesture_eventFilter_default")
	symaddr(&addr_QPanGesture_timerEvent_default, "QPanGesture_timerEvent_default")
	symaddr(&addr_QPanGesture_childEvent_default, "QPanGesture_childEvent_default")
	symaddr(&addr_QPanGesture_customEvent_default, "QPanGesture_customEvent_default")
	symaddr(&addr_QPanGesture_connectNotify_default, "QPanGesture_connectNotify_default")
	symaddr(&addr_QPanGesture_disconnectNotify_default, "QPanGesture_disconnectNotify_default")
	reg(&pg_QPanGesture_connect_destroyed, "QPanGesture_connect_destroyed")
	reg(&pg_QPanGesture_connect_destroyed2, "QPanGesture_connect_destroyed2")
	reg(&pg_QPanGesture_connect_objectNameChanged, "QPanGesture_connect_objectNameChanged")
	reg(&pg_QPanGesture__disconnect, "QPanGesture__disconnect")
	reg(&pg_QPanGesture__set_override, "QPanGesture__set_override")
	reg(&pg_QPanGesture_delete, "QPanGesture_delete")
}

type QPanGesture struct {
	h uintptr
	*QGesture
}

func bindQPanGesture(h uintptr) *QPanGesture {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPanGesture_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPanGesture{h: h, QGesture: &QGesture{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQPanGesture(h uintptr) *QPanGesture { return bindQPanGesture(h) }

func (this *QPanGesture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPanGesture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPanGesture(h unsafe.Pointer) *QPanGesture { return bindQPanGesture(uintptr(h)) }
func UnsafeNewQPanGestureOwned(h unsafe.Pointer) *QPanGesture { return newownedQPanGesture(uintptr(h)) }

func (this *QPanGesture) UnsafeRelease() {
	this.h = 0
	if this.QGesture != nil {
		this.QGesture.UnsafeRelease()
	}
}

func NewQPanGesture() *QPanGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPanGesture_new, uintptr(__goid))
	__o := newownedQPanGesture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPanGesture_virtIdx, QPanGesture_enable)
	return __o
}

func NewQPanGesture2(parent *QObject) *QPanGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPanGesture_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPanGesture(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPanGesture_virtIdx, QPanGesture_enable)
	return __o
}

func QPanGesture_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPanGesture_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPanGesture) LastOffset() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_lastOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPanGesture) Offset() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_offset, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPanGesture) Delta() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_delta, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPanGesture) Acceleration() float64 {
	__r := pg_QPanGesture_acceleration(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPanGesture) SetLastOffset(value *QPointF) {
	purego.SyscallN(addr_QPanGesture_setLastOffset, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QPanGesture) SetOffset(value *QPointF) {
	purego.SyscallN(addr_QPanGesture_setOffset, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QPanGesture) SetAcceleration(value float64) {
	pg_QPanGesture_setAcceleration(this.h, value)
	runtime.KeepAlive(this)
}

func QPanGesture_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPanGesture_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPanGesture_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPanGesture) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPanGesture) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPanGesture_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPanGesture) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPanGesture_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPanGesture) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPanGesture_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPanGesture_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPanGesture_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPanGesture_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPanGesture_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPanGesture_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPanGesture_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPanGesture_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPanGesture) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPanGesture_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPanGesture) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPanGesture_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPanGesture) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPanGesture_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPanGesture) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QPanGesture_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPanGesture) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPanGesture_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPanGesture) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPanGesture_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPanGesture) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPanGesture_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPanGesture) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPanGesture_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPanGesture) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPanGesture_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPanGesture) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPanGesture_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPanGesture) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPanGesture_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPanGesture) Disconnect(handle int64) {
	pg_QPanGesture__disconnect(handle)
	unregisterCallback(handle)
}

func QPanGesture_virtIdx(slot int64) int32 {
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

func QPanGesture_enable(h uintptr, idx int32) { pg_QPanGesture__set_override(h, idx) }

func (this *QPanGesture) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPanGesture) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QPanGesture) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPanGesture) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QPanGesture) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPanGesture) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPanGesture) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPanGesture) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPanGesture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPanGesture_delete(this.h)
	}
	this.UnsafeRelease()
}
