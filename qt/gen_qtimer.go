package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTimer_base func(uintptr) uintptr
	addr_QTimer_new uintptr
	addr_QTimer_new2 uintptr
	addr_QTimer_tr uintptr
	addr_QTimer_isActive uintptr
	addr_QTimer_timerId uintptr
	addr_QTimer_id uintptr
	addr_QTimer_setInterval uintptr
	addr_QTimer_interval uintptr
	addr_QTimer_remainingTime uintptr
	addr_QTimer_setTimerType uintptr
	addr_QTimer_timerType uintptr
	addr_QTimer_setSingleShot uintptr
	addr_QTimer_isSingleShot uintptr
	addr_QTimer_singleShot uintptr
	addr_QTimer_singleShot2 uintptr
	addr_QTimer_start uintptr
	addr_QTimer_start2 uintptr
	addr_QTimer_stop uintptr
	addr_QTimer_setInterval2 uintptr
	addr_QTimer_intervalAsDuration uintptr
	addr_QTimer_remainingTimeAsDuration uintptr
	addr_QTimer_singleShot3 uintptr
	addr_QTimer_singleShot4 uintptr
	addr_QTimer_start3 uintptr
	addr_QTimer_connect uintptr
	addr_QTimer_connect2 uintptr
	addr_QTimer_sender uintptr
	addr_QTimer_senderSignalIndex uintptr
	addr_QTimer_receivers uintptr
	addr_QTimer_isSignalConnected uintptr
	addr_QTimer_tr2 uintptr
	addr_QTimer_tr3 uintptr
	addr_QTimer_connect3 uintptr
	addr_QTimer_connect4 uintptr
	addr_QTimer_metaObject_default uintptr
	addr_QTimer_timerEvent_default uintptr
	addr_QTimer_event_default uintptr
	addr_QTimer_eventFilter_default uintptr
	addr_QTimer_childEvent_default uintptr
	addr_QTimer_customEvent_default uintptr
	addr_QTimer_connectNotify_default uintptr
	addr_QTimer_disconnectNotify_default uintptr
	pg_QTimer_connect_timeout func(uintptr, int64)
	pg_QTimer_connect_destroyed func(uintptr, int64)
	pg_QTimer_connect_destroyed2 func(uintptr, int64)
	pg_QTimer_connect_objectNameChanged func(uintptr, int64)
	pg_QTimer__disconnect func(int64)
	pg_QTimer__set_override func(uintptr, int32)
	pg_QTimer_delete func(uintptr)
)

func init() {
	reg(&pg_QTimer_base, "QTimer_base")
	symaddr(&addr_QTimer_new, "QTimer_new")
	symaddr(&addr_QTimer_new2, "QTimer_new2")
	symaddr(&addr_QTimer_tr, "QTimer_tr")
	symaddr(&addr_QTimer_isActive, "QTimer_isActive")
	symaddr(&addr_QTimer_timerId, "QTimer_timerId")
	symaddr(&addr_QTimer_id, "QTimer_id")
	symaddr(&addr_QTimer_setInterval, "QTimer_setInterval")
	symaddr(&addr_QTimer_interval, "QTimer_interval")
	symaddr(&addr_QTimer_remainingTime, "QTimer_remainingTime")
	symaddr(&addr_QTimer_setTimerType, "QTimer_setTimerType")
	symaddr(&addr_QTimer_timerType, "QTimer_timerType")
	symaddr(&addr_QTimer_setSingleShot, "QTimer_setSingleShot")
	symaddr(&addr_QTimer_isSingleShot, "QTimer_isSingleShot")
	symaddr(&addr_QTimer_singleShot, "QTimer_singleShot")
	symaddr(&addr_QTimer_singleShot2, "QTimer_singleShot2")
	symaddr(&addr_QTimer_start, "QTimer_start")
	symaddr(&addr_QTimer_start2, "QTimer_start2")
	symaddr(&addr_QTimer_stop, "QTimer_stop")
	symaddr(&addr_QTimer_setInterval2, "QTimer_setInterval2")
	symaddr(&addr_QTimer_intervalAsDuration, "QTimer_intervalAsDuration")
	symaddr(&addr_QTimer_remainingTimeAsDuration, "QTimer_remainingTimeAsDuration")
	symaddr(&addr_QTimer_singleShot3, "QTimer_singleShot3")
	symaddr(&addr_QTimer_singleShot4, "QTimer_singleShot4")
	symaddr(&addr_QTimer_start3, "QTimer_start3")
	symaddr(&addr_QTimer_connect, "QTimer_connect")
	symaddr(&addr_QTimer_connect2, "QTimer_connect2")
	symaddr(&addr_QTimer_sender, "QTimer_sender")
	symaddr(&addr_QTimer_senderSignalIndex, "QTimer_senderSignalIndex")
	symaddr(&addr_QTimer_receivers, "QTimer_receivers")
	symaddr(&addr_QTimer_isSignalConnected, "QTimer_isSignalConnected")
	symaddr(&addr_QTimer_tr2, "QTimer_tr2")
	symaddr(&addr_QTimer_tr3, "QTimer_tr3")
	symaddr(&addr_QTimer_connect3, "QTimer_connect3")
	symaddr(&addr_QTimer_connect4, "QTimer_connect4")
	symaddr(&addr_QTimer_metaObject_default, "QTimer_metaObject_default")
	symaddr(&addr_QTimer_timerEvent_default, "QTimer_timerEvent_default")
	symaddr(&addr_QTimer_event_default, "QTimer_event_default")
	symaddr(&addr_QTimer_eventFilter_default, "QTimer_eventFilter_default")
	symaddr(&addr_QTimer_childEvent_default, "QTimer_childEvent_default")
	symaddr(&addr_QTimer_customEvent_default, "QTimer_customEvent_default")
	symaddr(&addr_QTimer_connectNotify_default, "QTimer_connectNotify_default")
	symaddr(&addr_QTimer_disconnectNotify_default, "QTimer_disconnectNotify_default")
	reg(&pg_QTimer_connect_timeout, "QTimer_connect_timeout")
	reg(&pg_QTimer_connect_destroyed, "QTimer_connect_destroyed")
	reg(&pg_QTimer_connect_destroyed2, "QTimer_connect_destroyed2")
	reg(&pg_QTimer_connect_objectNameChanged, "QTimer_connect_objectNameChanged")
	reg(&pg_QTimer__disconnect, "QTimer__disconnect")
	reg(&pg_QTimer__set_override, "QTimer__set_override")
	reg(&pg_QTimer_delete, "QTimer_delete")
}

type QTimer struct {
	h uintptr
	*QObject
}

func bindQTimer(h uintptr) *QTimer {
	if h == 0 {
		return nil
	}
	return &QTimer{h: h, QObject: bindQObject(pg_QTimer_base(h))}
}

func newownedQTimer(h uintptr) *QTimer { return bindQTimer(h) }

func (this *QTimer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTimer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTimer(h unsafe.Pointer) *QTimer { return bindQTimer(uintptr(h)) }
func UnsafeNewQTimerOwned(h unsafe.Pointer) *QTimer { return newownedQTimer(uintptr(h)) }

func (this *QTimer) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQTimer() *QTimer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTimer_new, uintptr(__goid))
	__o := newownedQTimer(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTimer_virtIdx, QTimer_enable)
	return __o
}

func NewQTimer2(parent *QObject) *QTimer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTimer_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTimer(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTimer_virtIdx, QTimer_enable)
	return __o
}

func QTimer_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimer_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTimer) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QTimer_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTimer) TimerId() int {
	__r0, _, _ := purego.SyscallN(addr_QTimer_timerId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimer) Id() TimerId {
	__r0, _, _ := purego.SyscallN(addr_QTimer_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerId(__r)
}

func (this *QTimer) SetInterval(msec int) {
	purego.SyscallN(addr_QTimer_setInterval, this.h, uintptr(msec))
	runtime.KeepAlive(this)
}

func (this *QTimer) Interval() int {
	__r0, _, _ := purego.SyscallN(addr_QTimer_interval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimer) RemainingTime() int {
	__r0, _, _ := purego.SyscallN(addr_QTimer_remainingTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimer) SetTimerType(atype TimerType) {
	purego.SyscallN(addr_QTimer_setTimerType, this.h, uintptr(atype))
	runtime.KeepAlive(this)
}

func (this *QTimer) TimerType() TimerType {
	__r0, _, _ := purego.SyscallN(addr_QTimer_timerType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerType(__r)
}

func (this *QTimer) SetSingleShot(singleShot bool) {
	purego.SyscallN(addr_QTimer_setSingleShot, this.h, uintptr(b2i(singleShot)))
	runtime.KeepAlive(this)
}

func (this *QTimer) IsSingleShot() bool {
	__r0, _, _ := purego.SyscallN(addr_QTimer_isSingleShot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTimer_SingleShot(msec int, receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QTimer_singleShot, uintptr(msec), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func QTimer_SingleShot2(msec int, timerType TimerType, receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QTimer_singleShot2, uintptr(msec), uintptr(timerType), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func (this *QTimer) Start(msec int) {
	purego.SyscallN(addr_QTimer_start, this.h, uintptr(msec))
	runtime.KeepAlive(this)
}

func (this *QTimer) Start2() {
	purego.SyscallN(addr_QTimer_start2, this.h)
	runtime.KeepAlive(this)
}

func (this *QTimer) Stop() {
	purego.SyscallN(addr_QTimer_stop, this.h)
	runtime.KeepAlive(this)
}

func (this *QTimer) SetIntervalWithValue(value int64) {
	purego.SyscallN(addr_QTimer_setInterval2, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTimer) IntervalAsDuration() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTimer_intervalAsDuration, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTimer) RemainingTimeAsDuration() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTimer_remainingTimeAsDuration, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func QTimer_SingleShot3(value int64, receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QTimer_singleShot3, uintptr(value), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func QTimer_SingleShot4(interval int64, timerType TimerType, receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QTimer_singleShot4, uintptr(interval), uintptr(timerType), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func (this *QTimer) StartWithValue(value int64) {
	purego.SyscallN(addr_QTimer_start3, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func QTimer_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTimer_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTimer_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTimer_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTimer) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTimer_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTimer) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTimer_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimer) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTimer_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTimer) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimer_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTimer_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimer_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTimer_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimer_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTimer_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTimer_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTimer_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTimer_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTimer) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTimer_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTimer) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QTimer_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTimer) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimer_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTimer) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimer_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTimer) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTimer_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTimer) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTimer_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTimer) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTimer_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTimer) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTimer_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTimer) OnTimeout(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTimer_connect_timeout(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimer) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTimer_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimer) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTimer_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimer) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTimer_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimer) Disconnect(handle int64) {
	pg_QTimer__disconnect(handle)
	unregisterCallback(handle)
}

func QTimer_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -5409927581812374498:
		return 1
	case 1237236456312824991:
		return 2
	case 5217623514630858297:
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

func QTimer_enable(h uintptr, idx int32) { pg_QTimer__set_override(h, idx) }

func (this *QTimer) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTimer) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QTimer) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTimer) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTimer) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTimer) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTimer) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTimer) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTimer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTimer_delete(this.h)
	}
	this.UnsafeRelease()
}
