package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QChronoTimer_base func(uintptr) uintptr
	addr_QChronoTimer_new uintptr
	addr_QChronoTimer_new2 uintptr
	addr_QChronoTimer_new3 uintptr
	addr_QChronoTimer_new4 uintptr
	addr_QChronoTimer_tr uintptr
	addr_QChronoTimer_isActive uintptr
	addr_QChronoTimer_id uintptr
	addr_QChronoTimer_setInterval uintptr
	addr_QChronoTimer_interval uintptr
	addr_QChronoTimer_remainingTime uintptr
	addr_QChronoTimer_setTimerType uintptr
	addr_QChronoTimer_timerType uintptr
	addr_QChronoTimer_setSingleShot uintptr
	addr_QChronoTimer_isSingleShot uintptr
	addr_QChronoTimer_start uintptr
	addr_QChronoTimer_stop uintptr
	addr_QChronoTimer_connect uintptr
	addr_QChronoTimer_connect2 uintptr
	addr_QChronoTimer_sender uintptr
	addr_QChronoTimer_senderSignalIndex uintptr
	addr_QChronoTimer_receivers uintptr
	addr_QChronoTimer_isSignalConnected uintptr
	addr_QChronoTimer_tr2 uintptr
	addr_QChronoTimer_tr3 uintptr
	addr_QChronoTimer_connect3 uintptr
	addr_QChronoTimer_connect4 uintptr
	addr_QChronoTimer_metaObject_default uintptr
	addr_QChronoTimer_timerEvent_default uintptr
	addr_QChronoTimer_event_default uintptr
	addr_QChronoTimer_eventFilter_default uintptr
	addr_QChronoTimer_childEvent_default uintptr
	addr_QChronoTimer_customEvent_default uintptr
	addr_QChronoTimer_connectNotify_default uintptr
	addr_QChronoTimer_disconnectNotify_default uintptr
	pg_QChronoTimer_connect_timeout func(uintptr, int64)
	pg_QChronoTimer_connect_destroyed func(uintptr, int64)
	pg_QChronoTimer_connect_destroyed2 func(uintptr, int64)
	pg_QChronoTimer_connect_objectNameChanged func(uintptr, int64)
	pg_QChronoTimer__disconnect func(int64)
	pg_QChronoTimer__set_override func(uintptr, int32)
	pg_QChronoTimer_delete func(uintptr)
)

func init() {
	reg(&pg_QChronoTimer_base, "QChronoTimer_base")
	symaddr(&addr_QChronoTimer_new, "QChronoTimer_new")
	symaddr(&addr_QChronoTimer_new2, "QChronoTimer_new2")
	symaddr(&addr_QChronoTimer_new3, "QChronoTimer_new3")
	symaddr(&addr_QChronoTimer_new4, "QChronoTimer_new4")
	symaddr(&addr_QChronoTimer_tr, "QChronoTimer_tr")
	symaddr(&addr_QChronoTimer_isActive, "QChronoTimer_isActive")
	symaddr(&addr_QChronoTimer_id, "QChronoTimer_id")
	symaddr(&addr_QChronoTimer_setInterval, "QChronoTimer_setInterval")
	symaddr(&addr_QChronoTimer_interval, "QChronoTimer_interval")
	symaddr(&addr_QChronoTimer_remainingTime, "QChronoTimer_remainingTime")
	symaddr(&addr_QChronoTimer_setTimerType, "QChronoTimer_setTimerType")
	symaddr(&addr_QChronoTimer_timerType, "QChronoTimer_timerType")
	symaddr(&addr_QChronoTimer_setSingleShot, "QChronoTimer_setSingleShot")
	symaddr(&addr_QChronoTimer_isSingleShot, "QChronoTimer_isSingleShot")
	symaddr(&addr_QChronoTimer_start, "QChronoTimer_start")
	symaddr(&addr_QChronoTimer_stop, "QChronoTimer_stop")
	symaddr(&addr_QChronoTimer_connect, "QChronoTimer_connect")
	symaddr(&addr_QChronoTimer_connect2, "QChronoTimer_connect2")
	symaddr(&addr_QChronoTimer_sender, "QChronoTimer_sender")
	symaddr(&addr_QChronoTimer_senderSignalIndex, "QChronoTimer_senderSignalIndex")
	symaddr(&addr_QChronoTimer_receivers, "QChronoTimer_receivers")
	symaddr(&addr_QChronoTimer_isSignalConnected, "QChronoTimer_isSignalConnected")
	symaddr(&addr_QChronoTimer_tr2, "QChronoTimer_tr2")
	symaddr(&addr_QChronoTimer_tr3, "QChronoTimer_tr3")
	symaddr(&addr_QChronoTimer_connect3, "QChronoTimer_connect3")
	symaddr(&addr_QChronoTimer_connect4, "QChronoTimer_connect4")
	symaddr(&addr_QChronoTimer_metaObject_default, "QChronoTimer_metaObject_default")
	symaddr(&addr_QChronoTimer_timerEvent_default, "QChronoTimer_timerEvent_default")
	symaddr(&addr_QChronoTimer_event_default, "QChronoTimer_event_default")
	symaddr(&addr_QChronoTimer_eventFilter_default, "QChronoTimer_eventFilter_default")
	symaddr(&addr_QChronoTimer_childEvent_default, "QChronoTimer_childEvent_default")
	symaddr(&addr_QChronoTimer_customEvent_default, "QChronoTimer_customEvent_default")
	symaddr(&addr_QChronoTimer_connectNotify_default, "QChronoTimer_connectNotify_default")
	symaddr(&addr_QChronoTimer_disconnectNotify_default, "QChronoTimer_disconnectNotify_default")
	reg(&pg_QChronoTimer_connect_timeout, "QChronoTimer_connect_timeout")
	reg(&pg_QChronoTimer_connect_destroyed, "QChronoTimer_connect_destroyed")
	reg(&pg_QChronoTimer_connect_destroyed2, "QChronoTimer_connect_destroyed2")
	reg(&pg_QChronoTimer_connect_objectNameChanged, "QChronoTimer_connect_objectNameChanged")
	reg(&pg_QChronoTimer__disconnect, "QChronoTimer__disconnect")
	reg(&pg_QChronoTimer__set_override, "QChronoTimer__set_override")
	reg(&pg_QChronoTimer_delete, "QChronoTimer_delete")
}

type QChronoTimer struct {
	h uintptr
	*QObject
}

func bindQChronoTimer(h uintptr) *QChronoTimer {
	if h == 0 {
		return nil
	}
	return &QChronoTimer{h: h, QObject: bindQObject(pg_QChronoTimer_base(h))}
}

func newownedQChronoTimer(h uintptr) *QChronoTimer { return bindQChronoTimer(h) }

func (this *QChronoTimer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QChronoTimer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQChronoTimer(h unsafe.Pointer) *QChronoTimer { return bindQChronoTimer(uintptr(h)) }
func UnsafeNewQChronoTimerOwned(h unsafe.Pointer) *QChronoTimer { return newownedQChronoTimer(uintptr(h)) }

func (this *QChronoTimer) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQChronoTimer(nsec int64) *QChronoTimer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_new, uintptr(__goid), uintptr(nsec))
	__o := newownedQChronoTimer(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QChronoTimer_virtIdx, QChronoTimer_enable)
	return __o
}

func NewQChronoTimer2() *QChronoTimer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_new2, uintptr(__goid))
	__o := newownedQChronoTimer(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QChronoTimer_virtIdx, QChronoTimer_enable)
	return __o
}

func NewQChronoTimer3(nsec int64, parent *QObject) *QChronoTimer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_new3, uintptr(__goid), uintptr(nsec), uintptr(parent.UnsafePointer()))
	__o := newownedQChronoTimer(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QChronoTimer_virtIdx, QChronoTimer_enable)
	return __o
}

func NewQChronoTimer4(parent *QObject) *QChronoTimer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQChronoTimer(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QChronoTimer_virtIdx, QChronoTimer_enable)
	return __o
}

func QChronoTimer_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QChronoTimer) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChronoTimer) Id() TimerId {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerId(__r)
}

func (this *QChronoTimer) SetInterval(nsec int64) {
	purego.SyscallN(addr_QChronoTimer_setInterval, this.h, uintptr(nsec))
	runtime.KeepAlive(this)
}

func (this *QChronoTimer) Interval() int64 {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_interval, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QChronoTimer) RemainingTime() int64 {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_remainingTime, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QChronoTimer) SetTimerType(atype TimerType) {
	purego.SyscallN(addr_QChronoTimer_setTimerType, this.h, uintptr(atype))
	runtime.KeepAlive(this)
}

func (this *QChronoTimer) TimerType() TimerType {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_timerType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerType(__r)
}

func (this *QChronoTimer) SetSingleShot(singleShot bool) {
	purego.SyscallN(addr_QChronoTimer_setSingleShot, this.h, uintptr(b2i(singleShot)))
	runtime.KeepAlive(this)
}

func (this *QChronoTimer) IsSingleShot() bool {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_isSingleShot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChronoTimer) Start() {
	purego.SyscallN(addr_QChronoTimer_start, this.h)
	runtime.KeepAlive(this)
}

func (this *QChronoTimer) Stop() {
	purego.SyscallN(addr_QChronoTimer_stop, this.h)
	runtime.KeepAlive(this)
}

func QChronoTimer_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QChronoTimer_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QChronoTimer) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QChronoTimer) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QChronoTimer) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QChronoTimer) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QChronoTimer_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QChronoTimer_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QChronoTimer_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QChronoTimer_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QChronoTimer) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QChronoTimer_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QChronoTimer) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QChronoTimer_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QChronoTimer) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QChronoTimer) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QChronoTimer_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QChronoTimer) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QChronoTimer_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QChronoTimer) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QChronoTimer_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QChronoTimer) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QChronoTimer_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QChronoTimer) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QChronoTimer_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QChronoTimer) OnTimeout(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QChronoTimer_connect_timeout(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QChronoTimer) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QChronoTimer_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QChronoTimer) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QChronoTimer_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QChronoTimer) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QChronoTimer_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QChronoTimer) Disconnect(handle int64) {
	pg_QChronoTimer__disconnect(handle)
	unregisterCallback(handle)
}

func QChronoTimer_virtIdx(slot int64) int32 {
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

func QChronoTimer_enable(h uintptr, idx int32) { pg_QChronoTimer__set_override(h, idx) }

func (this *QChronoTimer) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QChronoTimer) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QChronoTimer) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QChronoTimer) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QChronoTimer) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QChronoTimer) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QChronoTimer) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QChronoTimer) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QChronoTimer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QChronoTimer_delete(this.h)
	}
	this.UnsafeRelease()
}
