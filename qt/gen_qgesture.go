package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGesture_base func(uintptr) uintptr
	addr_QGesture_new uintptr
	addr_QGesture_new2 uintptr
	addr_QGesture_tr uintptr
	addr_QGesture_gestureType uintptr
	addr_QGesture_state uintptr
	addr_QGesture_hotSpot uintptr
	addr_QGesture_setHotSpot uintptr
	addr_QGesture_hasHotSpot uintptr
	addr_QGesture_unsetHotSpot uintptr
	addr_QGesture_setGestureCancelPolicy uintptr
	addr_QGesture_gestureCancelPolicy uintptr
	addr_QGesture_connect uintptr
	addr_QGesture_connect2 uintptr
	addr_QGesture_sender uintptr
	addr_QGesture_senderSignalIndex uintptr
	addr_QGesture_receivers uintptr
	addr_QGesture_isSignalConnected uintptr
	addr_QGesture_tr2 uintptr
	addr_QGesture_tr3 uintptr
	addr_QGesture_connect3 uintptr
	addr_QGesture_connect4 uintptr
	addr_QGesture_metaObject_default uintptr
	addr_QGesture_event_default uintptr
	addr_QGesture_eventFilter_default uintptr
	addr_QGesture_timerEvent_default uintptr
	addr_QGesture_childEvent_default uintptr
	addr_QGesture_customEvent_default uintptr
	addr_QGesture_connectNotify_default uintptr
	addr_QGesture_disconnectNotify_default uintptr
	pg_QGesture_connect_destroyed func(uintptr, int64)
	pg_QGesture_connect_destroyed2 func(uintptr, int64)
	pg_QGesture_connect_objectNameChanged func(uintptr, int64)
	pg_QGesture__disconnect func(int64)
	pg_QGesture__set_override func(uintptr, int32)
	pg_QGesture_delete func(uintptr)
)

func init() {
	reg(&pg_QGesture_base, "QGesture_base")
	symaddr(&addr_QGesture_new, "QGesture_new")
	symaddr(&addr_QGesture_new2, "QGesture_new2")
	symaddr(&addr_QGesture_tr, "QGesture_tr")
	symaddr(&addr_QGesture_gestureType, "QGesture_gestureType")
	symaddr(&addr_QGesture_state, "QGesture_state")
	symaddr(&addr_QGesture_hotSpot, "QGesture_hotSpot")
	symaddr(&addr_QGesture_setHotSpot, "QGesture_setHotSpot")
	symaddr(&addr_QGesture_hasHotSpot, "QGesture_hasHotSpot")
	symaddr(&addr_QGesture_unsetHotSpot, "QGesture_unsetHotSpot")
	symaddr(&addr_QGesture_setGestureCancelPolicy, "QGesture_setGestureCancelPolicy")
	symaddr(&addr_QGesture_gestureCancelPolicy, "QGesture_gestureCancelPolicy")
	symaddr(&addr_QGesture_connect, "QGesture_connect")
	symaddr(&addr_QGesture_connect2, "QGesture_connect2")
	symaddr(&addr_QGesture_sender, "QGesture_sender")
	symaddr(&addr_QGesture_senderSignalIndex, "QGesture_senderSignalIndex")
	symaddr(&addr_QGesture_receivers, "QGesture_receivers")
	symaddr(&addr_QGesture_isSignalConnected, "QGesture_isSignalConnected")
	symaddr(&addr_QGesture_tr2, "QGesture_tr2")
	symaddr(&addr_QGesture_tr3, "QGesture_tr3")
	symaddr(&addr_QGesture_connect3, "QGesture_connect3")
	symaddr(&addr_QGesture_connect4, "QGesture_connect4")
	symaddr(&addr_QGesture_metaObject_default, "QGesture_metaObject_default")
	symaddr(&addr_QGesture_event_default, "QGesture_event_default")
	symaddr(&addr_QGesture_eventFilter_default, "QGesture_eventFilter_default")
	symaddr(&addr_QGesture_timerEvent_default, "QGesture_timerEvent_default")
	symaddr(&addr_QGesture_childEvent_default, "QGesture_childEvent_default")
	symaddr(&addr_QGesture_customEvent_default, "QGesture_customEvent_default")
	symaddr(&addr_QGesture_connectNotify_default, "QGesture_connectNotify_default")
	symaddr(&addr_QGesture_disconnectNotify_default, "QGesture_disconnectNotify_default")
	reg(&pg_QGesture_connect_destroyed, "QGesture_connect_destroyed")
	reg(&pg_QGesture_connect_destroyed2, "QGesture_connect_destroyed2")
	reg(&pg_QGesture_connect_objectNameChanged, "QGesture_connect_objectNameChanged")
	reg(&pg_QGesture__disconnect, "QGesture__disconnect")
	reg(&pg_QGesture__set_override, "QGesture__set_override")
	reg(&pg_QGesture_delete, "QGesture_delete")
}

type QGesture struct {
	h uintptr
	*QObject
}

func bindQGesture(h uintptr) *QGesture {
	if h == 0 {
		return nil
	}
	return &QGesture{h: h, QObject: bindQObject(pg_QGesture_base(h))}
}

func newownedQGesture(h uintptr) *QGesture { return bindQGesture(h) }

func (this *QGesture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGesture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGesture(h unsafe.Pointer) *QGesture { return bindQGesture(uintptr(h)) }
func UnsafeNewQGestureOwned(h unsafe.Pointer) *QGesture { return newownedQGesture(uintptr(h)) }

func (this *QGesture) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQGesture() *QGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGesture_new, uintptr(__goid))
	__o := newownedQGesture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGesture_virtIdx, QGesture_enable)
	return __o
}

func NewQGesture2(parent *QObject) *QGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGesture_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGesture(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGesture_virtIdx, QGesture_enable)
	return __o
}

func QGesture_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGesture_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGesture) GestureType() GestureType {
	__r0, _, _ := purego.SyscallN(addr_QGesture_gestureType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return GestureType(__r)
}

func (this *QGesture) State() GestureState {
	__r0, _, _ := purego.SyscallN(addr_QGesture_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return GestureState(__r)
}

func (this *QGesture) HotSpot() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGesture_hotSpot, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGesture) SetHotSpot(value *QPointF) {
	purego.SyscallN(addr_QGesture_setHotSpot, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QGesture) HasHotSpot() bool {
	__r0, _, _ := purego.SyscallN(addr_QGesture_hasHotSpot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGesture) UnsetHotSpot() {
	purego.SyscallN(addr_QGesture_unsetHotSpot, this.h)
	runtime.KeepAlive(this)
}

func (this *QGesture) SetGestureCancelPolicy(policy QGesture__GestureCancelPolicy) {
	purego.SyscallN(addr_QGesture_setGestureCancelPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QGesture) GestureCancelPolicy() QGesture__GestureCancelPolicy {
	__r0, _, _ := purego.SyscallN(addr_QGesture_gestureCancelPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGesture__GestureCancelPolicy(__r)
}

func QGesture_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGesture_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGesture_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGesture_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGesture) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGesture_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGesture) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGesture_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGesture) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGesture_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGesture) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGesture_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGesture_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGesture_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGesture_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGesture_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGesture_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGesture_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGesture_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGesture_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGesture) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGesture_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGesture) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGesture_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGesture) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGesture_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGesture) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGesture_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGesture) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGesture_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGesture) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGesture_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGesture) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGesture_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGesture) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGesture_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGesture) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGesture_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGesture) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGesture_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGesture) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGesture_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGesture) Disconnect(handle int64) {
	pg_QGesture__disconnect(handle)
	unregisterCallback(handle)
}

func QGesture_virtIdx(slot int64) int32 {
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

func QGesture_enable(h uintptr, idx int32) { pg_QGesture__set_override(h, idx) }

func (this *QGesture) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGesture) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGesture) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGesture) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGesture) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGesture) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGesture) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGesture) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGesture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGesture_delete(this.h)
	}
	this.UnsafeRelease()
}
