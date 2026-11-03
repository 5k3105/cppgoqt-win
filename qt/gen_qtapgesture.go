package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTapGesture_chain uintptr
	addr_QTapGesture_new uintptr
	addr_QTapGesture_new2 uintptr
	addr_QTapGesture_tr uintptr
	addr_QTapGesture_position uintptr
	addr_QTapGesture_setPosition uintptr
	addr_QTapGesture_connect uintptr
	addr_QTapGesture_connect2 uintptr
	addr_QTapGesture_sender uintptr
	addr_QTapGesture_senderSignalIndex uintptr
	addr_QTapGesture_receivers uintptr
	addr_QTapGesture_isSignalConnected uintptr
	addr_QTapGesture_tr2 uintptr
	addr_QTapGesture_tr3 uintptr
	addr_QTapGesture_connect3 uintptr
	addr_QTapGesture_connect4 uintptr
	addr_QTapGesture_metaObject_default uintptr
	addr_QTapGesture_event_default uintptr
	addr_QTapGesture_eventFilter_default uintptr
	addr_QTapGesture_timerEvent_default uintptr
	addr_QTapGesture_childEvent_default uintptr
	addr_QTapGesture_customEvent_default uintptr
	addr_QTapGesture_connectNotify_default uintptr
	addr_QTapGesture_disconnectNotify_default uintptr
	pg_QTapGesture_connect_destroyed func(uintptr, int64)
	pg_QTapGesture_connect_destroyed2 func(uintptr, int64)
	pg_QTapGesture_connect_objectNameChanged func(uintptr, int64)
	pg_QTapGesture__disconnect func(int64)
	pg_QTapGesture__set_override func(uintptr, int32)
	pg_QTapGesture_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTapGesture_chain, "QTapGesture_chain")
	symaddr(&addr_QTapGesture_new, "QTapGesture_new")
	symaddr(&addr_QTapGesture_new2, "QTapGesture_new2")
	symaddr(&addr_QTapGesture_tr, "QTapGesture_tr")
	symaddr(&addr_QTapGesture_position, "QTapGesture_position")
	symaddr(&addr_QTapGesture_setPosition, "QTapGesture_setPosition")
	symaddr(&addr_QTapGesture_connect, "QTapGesture_connect")
	symaddr(&addr_QTapGesture_connect2, "QTapGesture_connect2")
	symaddr(&addr_QTapGesture_sender, "QTapGesture_sender")
	symaddr(&addr_QTapGesture_senderSignalIndex, "QTapGesture_senderSignalIndex")
	symaddr(&addr_QTapGesture_receivers, "QTapGesture_receivers")
	symaddr(&addr_QTapGesture_isSignalConnected, "QTapGesture_isSignalConnected")
	symaddr(&addr_QTapGesture_tr2, "QTapGesture_tr2")
	symaddr(&addr_QTapGesture_tr3, "QTapGesture_tr3")
	symaddr(&addr_QTapGesture_connect3, "QTapGesture_connect3")
	symaddr(&addr_QTapGesture_connect4, "QTapGesture_connect4")
	symaddr(&addr_QTapGesture_metaObject_default, "QTapGesture_metaObject_default")
	symaddr(&addr_QTapGesture_event_default, "QTapGesture_event_default")
	symaddr(&addr_QTapGesture_eventFilter_default, "QTapGesture_eventFilter_default")
	symaddr(&addr_QTapGesture_timerEvent_default, "QTapGesture_timerEvent_default")
	symaddr(&addr_QTapGesture_childEvent_default, "QTapGesture_childEvent_default")
	symaddr(&addr_QTapGesture_customEvent_default, "QTapGesture_customEvent_default")
	symaddr(&addr_QTapGesture_connectNotify_default, "QTapGesture_connectNotify_default")
	symaddr(&addr_QTapGesture_disconnectNotify_default, "QTapGesture_disconnectNotify_default")
	reg(&pg_QTapGesture_connect_destroyed, "QTapGesture_connect_destroyed")
	reg(&pg_QTapGesture_connect_destroyed2, "QTapGesture_connect_destroyed2")
	reg(&pg_QTapGesture_connect_objectNameChanged, "QTapGesture_connect_objectNameChanged")
	reg(&pg_QTapGesture__disconnect, "QTapGesture__disconnect")
	reg(&pg_QTapGesture__set_override, "QTapGesture__set_override")
	reg(&pg_QTapGesture_delete, "QTapGesture_delete")
}

type QTapGesture struct {
	h uintptr
	*QGesture
}

func bindQTapGesture(h uintptr) *QTapGesture {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTapGesture_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTapGesture{h: h, QGesture: &QGesture{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQTapGesture(h uintptr) *QTapGesture { return bindQTapGesture(h) }

func (this *QTapGesture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTapGesture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTapGesture(h unsafe.Pointer) *QTapGesture { return bindQTapGesture(uintptr(h)) }
func UnsafeNewQTapGestureOwned(h unsafe.Pointer) *QTapGesture { return newownedQTapGesture(uintptr(h)) }

func (this *QTapGesture) UnsafeRelease() {
	this.h = 0
	if this.QGesture != nil {
		this.QGesture.UnsafeRelease()
	}
}

func NewQTapGesture() *QTapGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTapGesture_new, uintptr(__goid))
	__o := newownedQTapGesture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTapGesture_virtIdx, QTapGesture_enable)
	return __o
}

func NewQTapGesture2(parent *QObject) *QTapGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTapGesture_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTapGesture(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTapGesture_virtIdx, QTapGesture_enable)
	return __o
}

func QTapGesture_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTapGesture_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTapGesture) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTapGesture_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QTapGesture) SetPosition(pos *QPointF) {
	purego.SyscallN(addr_QTapGesture_setPosition, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func QTapGesture_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTapGesture_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTapGesture_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTapGesture_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTapGesture) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTapGesture_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTapGesture) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTapGesture_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTapGesture) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTapGesture_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTapGesture) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTapGesture_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTapGesture_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTapGesture_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTapGesture_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTapGesture_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTapGesture_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTapGesture_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTapGesture_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTapGesture_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTapGesture) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTapGesture_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTapGesture) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTapGesture_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTapGesture) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTapGesture_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTapGesture) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTapGesture_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTapGesture) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTapGesture_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTapGesture) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTapGesture_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTapGesture) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTapGesture_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTapGesture) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTapGesture_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTapGesture) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTapGesture_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTapGesture) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTapGesture_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTapGesture) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTapGesture_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTapGesture) Disconnect(handle int64) {
	pg_QTapGesture__disconnect(handle)
	unregisterCallback(handle)
}

func QTapGesture_virtIdx(slot int64) int32 {
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

func QTapGesture_enable(h uintptr, idx int32) { pg_QTapGesture__set_override(h, idx) }

func (this *QTapGesture) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTapGesture) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTapGesture) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTapGesture) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTapGesture) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTapGesture) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTapGesture) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTapGesture) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTapGesture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTapGesture_delete(this.h)
	}
	this.UnsafeRelease()
}
