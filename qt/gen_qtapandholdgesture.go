package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTapAndHoldGesture_chain uintptr
	addr_QTapAndHoldGesture_new uintptr
	addr_QTapAndHoldGesture_new2 uintptr
	addr_QTapAndHoldGesture_tr uintptr
	addr_QTapAndHoldGesture_position uintptr
	addr_QTapAndHoldGesture_setPosition uintptr
	addr_QTapAndHoldGesture_setTimeout uintptr
	addr_QTapAndHoldGesture_timeout uintptr
	addr_QTapAndHoldGesture_connect uintptr
	addr_QTapAndHoldGesture_connect2 uintptr
	addr_QTapAndHoldGesture_sender uintptr
	addr_QTapAndHoldGesture_senderSignalIndex uintptr
	addr_QTapAndHoldGesture_receivers uintptr
	addr_QTapAndHoldGesture_isSignalConnected uintptr
	addr_QTapAndHoldGesture_tr2 uintptr
	addr_QTapAndHoldGesture_tr3 uintptr
	addr_QTapAndHoldGesture_connect3 uintptr
	addr_QTapAndHoldGesture_connect4 uintptr
	addr_QTapAndHoldGesture_metaObject_default uintptr
	addr_QTapAndHoldGesture_event_default uintptr
	addr_QTapAndHoldGesture_eventFilter_default uintptr
	addr_QTapAndHoldGesture_timerEvent_default uintptr
	addr_QTapAndHoldGesture_childEvent_default uintptr
	addr_QTapAndHoldGesture_customEvent_default uintptr
	addr_QTapAndHoldGesture_connectNotify_default uintptr
	addr_QTapAndHoldGesture_disconnectNotify_default uintptr
	pg_QTapAndHoldGesture_connect_destroyed func(uintptr, int64)
	pg_QTapAndHoldGesture_connect_destroyed2 func(uintptr, int64)
	pg_QTapAndHoldGesture_connect_objectNameChanged func(uintptr, int64)
	pg_QTapAndHoldGesture__disconnect func(int64)
	pg_QTapAndHoldGesture__set_override func(uintptr, int32)
	pg_QTapAndHoldGesture_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTapAndHoldGesture_chain, "QTapAndHoldGesture_chain")
	symaddr(&addr_QTapAndHoldGesture_new, "QTapAndHoldGesture_new")
	symaddr(&addr_QTapAndHoldGesture_new2, "QTapAndHoldGesture_new2")
	symaddr(&addr_QTapAndHoldGesture_tr, "QTapAndHoldGesture_tr")
	symaddr(&addr_QTapAndHoldGesture_position, "QTapAndHoldGesture_position")
	symaddr(&addr_QTapAndHoldGesture_setPosition, "QTapAndHoldGesture_setPosition")
	symaddr(&addr_QTapAndHoldGesture_setTimeout, "QTapAndHoldGesture_setTimeout")
	symaddr(&addr_QTapAndHoldGesture_timeout, "QTapAndHoldGesture_timeout")
	symaddr(&addr_QTapAndHoldGesture_connect, "QTapAndHoldGesture_connect")
	symaddr(&addr_QTapAndHoldGesture_connect2, "QTapAndHoldGesture_connect2")
	symaddr(&addr_QTapAndHoldGesture_sender, "QTapAndHoldGesture_sender")
	symaddr(&addr_QTapAndHoldGesture_senderSignalIndex, "QTapAndHoldGesture_senderSignalIndex")
	symaddr(&addr_QTapAndHoldGesture_receivers, "QTapAndHoldGesture_receivers")
	symaddr(&addr_QTapAndHoldGesture_isSignalConnected, "QTapAndHoldGesture_isSignalConnected")
	symaddr(&addr_QTapAndHoldGesture_tr2, "QTapAndHoldGesture_tr2")
	symaddr(&addr_QTapAndHoldGesture_tr3, "QTapAndHoldGesture_tr3")
	symaddr(&addr_QTapAndHoldGesture_connect3, "QTapAndHoldGesture_connect3")
	symaddr(&addr_QTapAndHoldGesture_connect4, "QTapAndHoldGesture_connect4")
	symaddr(&addr_QTapAndHoldGesture_metaObject_default, "QTapAndHoldGesture_metaObject_default")
	symaddr(&addr_QTapAndHoldGesture_event_default, "QTapAndHoldGesture_event_default")
	symaddr(&addr_QTapAndHoldGesture_eventFilter_default, "QTapAndHoldGesture_eventFilter_default")
	symaddr(&addr_QTapAndHoldGesture_timerEvent_default, "QTapAndHoldGesture_timerEvent_default")
	symaddr(&addr_QTapAndHoldGesture_childEvent_default, "QTapAndHoldGesture_childEvent_default")
	symaddr(&addr_QTapAndHoldGesture_customEvent_default, "QTapAndHoldGesture_customEvent_default")
	symaddr(&addr_QTapAndHoldGesture_connectNotify_default, "QTapAndHoldGesture_connectNotify_default")
	symaddr(&addr_QTapAndHoldGesture_disconnectNotify_default, "QTapAndHoldGesture_disconnectNotify_default")
	reg(&pg_QTapAndHoldGesture_connect_destroyed, "QTapAndHoldGesture_connect_destroyed")
	reg(&pg_QTapAndHoldGesture_connect_destroyed2, "QTapAndHoldGesture_connect_destroyed2")
	reg(&pg_QTapAndHoldGesture_connect_objectNameChanged, "QTapAndHoldGesture_connect_objectNameChanged")
	reg(&pg_QTapAndHoldGesture__disconnect, "QTapAndHoldGesture__disconnect")
	reg(&pg_QTapAndHoldGesture__set_override, "QTapAndHoldGesture__set_override")
	reg(&pg_QTapAndHoldGesture_delete, "QTapAndHoldGesture_delete")
}

type QTapAndHoldGesture struct {
	h uintptr
	*QGesture
}

func bindQTapAndHoldGesture(h uintptr) *QTapAndHoldGesture {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTapAndHoldGesture_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTapAndHoldGesture{h: h, QGesture: &QGesture{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQTapAndHoldGesture(h uintptr) *QTapAndHoldGesture { return bindQTapAndHoldGesture(h) }

func (this *QTapAndHoldGesture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTapAndHoldGesture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTapAndHoldGesture(h unsafe.Pointer) *QTapAndHoldGesture { return bindQTapAndHoldGesture(uintptr(h)) }
func UnsafeNewQTapAndHoldGestureOwned(h unsafe.Pointer) *QTapAndHoldGesture { return newownedQTapAndHoldGesture(uintptr(h)) }

func (this *QTapAndHoldGesture) UnsafeRelease() {
	this.h = 0
	if this.QGesture != nil {
		this.QGesture.UnsafeRelease()
	}
}

func NewQTapAndHoldGesture() *QTapAndHoldGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_new, uintptr(__goid))
	__o := newownedQTapAndHoldGesture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTapAndHoldGesture_virtIdx, QTapAndHoldGesture_enable)
	return __o
}

func NewQTapAndHoldGesture2(parent *QObject) *QTapAndHoldGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTapAndHoldGesture(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTapAndHoldGesture_virtIdx, QTapAndHoldGesture_enable)
	return __o
}

func QTapAndHoldGesture_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTapAndHoldGesture) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QTapAndHoldGesture) SetPosition(pos *QPointF) {
	purego.SyscallN(addr_QTapAndHoldGesture_setPosition, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func QTapAndHoldGesture_SetTimeout(msecs int) {
	purego.SyscallN(addr_QTapAndHoldGesture_setTimeout, uintptr(msecs))
}

func QTapAndHoldGesture_Timeout() int {
	__r0, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_timeout)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QTapAndHoldGesture_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTapAndHoldGesture_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTapAndHoldGesture) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTapAndHoldGesture) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTapAndHoldGesture) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTapAndHoldGesture) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTapAndHoldGesture_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTapAndHoldGesture_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTapAndHoldGesture_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTapAndHoldGesture_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTapAndHoldGesture) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTapAndHoldGesture) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTapAndHoldGesture) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTapAndHoldGesture_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTapAndHoldGesture) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTapAndHoldGesture_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTapAndHoldGesture) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTapAndHoldGesture_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTapAndHoldGesture) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTapAndHoldGesture_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTapAndHoldGesture) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTapAndHoldGesture_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTapAndHoldGesture) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTapAndHoldGesture_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTapAndHoldGesture) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTapAndHoldGesture_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTapAndHoldGesture) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTapAndHoldGesture_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTapAndHoldGesture) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTapAndHoldGesture_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTapAndHoldGesture) Disconnect(handle int64) {
	pg_QTapAndHoldGesture__disconnect(handle)
	unregisterCallback(handle)
}

func QTapAndHoldGesture_virtIdx(slot int64) int32 {
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

func QTapAndHoldGesture_enable(h uintptr, idx int32) { pg_QTapAndHoldGesture__set_override(h, idx) }

func (this *QTapAndHoldGesture) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTapAndHoldGesture) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTapAndHoldGesture) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTapAndHoldGesture) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTapAndHoldGesture) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTapAndHoldGesture) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTapAndHoldGesture) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTapAndHoldGesture) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTapAndHoldGesture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTapAndHoldGesture_delete(this.h)
	}
	this.UnsafeRelease()
}
