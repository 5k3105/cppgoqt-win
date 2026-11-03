package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAnimationDriver_base func(uintptr) uintptr
	addr_QAnimationDriver_new uintptr
	addr_QAnimationDriver_new2 uintptr
	addr_QAnimationDriver_tr uintptr
	addr_QAnimationDriver_advance uintptr
	addr_QAnimationDriver_install uintptr
	addr_QAnimationDriver_uninstall uintptr
	addr_QAnimationDriver_isRunning uintptr
	addr_QAnimationDriver_elapsed uintptr
	addr_QAnimationDriver_started uintptr
	addr_QAnimationDriver_stopped uintptr
	addr_QAnimationDriver_advanceAnimation uintptr
	addr_QAnimationDriver_connect uintptr
	addr_QAnimationDriver_connect2 uintptr
	addr_QAnimationDriver_sender uintptr
	addr_QAnimationDriver_senderSignalIndex uintptr
	addr_QAnimationDriver_receivers uintptr
	addr_QAnimationDriver_isSignalConnected uintptr
	addr_QAnimationDriver_tr2 uintptr
	addr_QAnimationDriver_tr3 uintptr
	addr_QAnimationDriver_connect3 uintptr
	addr_QAnimationDriver_connect4 uintptr
	addr_QAnimationDriver_advance_default uintptr
	addr_QAnimationDriver_elapsed_default uintptr
	addr_QAnimationDriver_metaObject_default uintptr
	addr_QAnimationDriver_start_default uintptr
	addr_QAnimationDriver_stop_default uintptr
	addr_QAnimationDriver_event_default uintptr
	addr_QAnimationDriver_eventFilter_default uintptr
	addr_QAnimationDriver_timerEvent_default uintptr
	addr_QAnimationDriver_childEvent_default uintptr
	addr_QAnimationDriver_customEvent_default uintptr
	addr_QAnimationDriver_connectNotify_default uintptr
	addr_QAnimationDriver_disconnectNotify_default uintptr
	pg_QAnimationDriver_connect_started func(uintptr, int64)
	pg_QAnimationDriver_connect_stopped func(uintptr, int64)
	pg_QAnimationDriver_connect_destroyed func(uintptr, int64)
	pg_QAnimationDriver_connect_destroyed2 func(uintptr, int64)
	pg_QAnimationDriver_connect_objectNameChanged func(uintptr, int64)
	pg_QAnimationDriver__disconnect func(int64)
	pg_QAnimationDriver__set_override func(uintptr, int32)
	pg_QAnimationDriver_delete func(uintptr)
)

func init() {
	reg(&pg_QAnimationDriver_base, "QAnimationDriver_base")
	symaddr(&addr_QAnimationDriver_new, "QAnimationDriver_new")
	symaddr(&addr_QAnimationDriver_new2, "QAnimationDriver_new2")
	symaddr(&addr_QAnimationDriver_tr, "QAnimationDriver_tr")
	symaddr(&addr_QAnimationDriver_advance, "QAnimationDriver_advance")
	symaddr(&addr_QAnimationDriver_install, "QAnimationDriver_install")
	symaddr(&addr_QAnimationDriver_uninstall, "QAnimationDriver_uninstall")
	symaddr(&addr_QAnimationDriver_isRunning, "QAnimationDriver_isRunning")
	symaddr(&addr_QAnimationDriver_elapsed, "QAnimationDriver_elapsed")
	symaddr(&addr_QAnimationDriver_started, "QAnimationDriver_started")
	symaddr(&addr_QAnimationDriver_stopped, "QAnimationDriver_stopped")
	symaddr(&addr_QAnimationDriver_advanceAnimation, "QAnimationDriver_advanceAnimation")
	symaddr(&addr_QAnimationDriver_connect, "QAnimationDriver_connect")
	symaddr(&addr_QAnimationDriver_connect2, "QAnimationDriver_connect2")
	symaddr(&addr_QAnimationDriver_sender, "QAnimationDriver_sender")
	symaddr(&addr_QAnimationDriver_senderSignalIndex, "QAnimationDriver_senderSignalIndex")
	symaddr(&addr_QAnimationDriver_receivers, "QAnimationDriver_receivers")
	symaddr(&addr_QAnimationDriver_isSignalConnected, "QAnimationDriver_isSignalConnected")
	symaddr(&addr_QAnimationDriver_tr2, "QAnimationDriver_tr2")
	symaddr(&addr_QAnimationDriver_tr3, "QAnimationDriver_tr3")
	symaddr(&addr_QAnimationDriver_connect3, "QAnimationDriver_connect3")
	symaddr(&addr_QAnimationDriver_connect4, "QAnimationDriver_connect4")
	symaddr(&addr_QAnimationDriver_advance_default, "QAnimationDriver_advance_default")
	symaddr(&addr_QAnimationDriver_elapsed_default, "QAnimationDriver_elapsed_default")
	symaddr(&addr_QAnimationDriver_metaObject_default, "QAnimationDriver_metaObject_default")
	symaddr(&addr_QAnimationDriver_start_default, "QAnimationDriver_start_default")
	symaddr(&addr_QAnimationDriver_stop_default, "QAnimationDriver_stop_default")
	symaddr(&addr_QAnimationDriver_event_default, "QAnimationDriver_event_default")
	symaddr(&addr_QAnimationDriver_eventFilter_default, "QAnimationDriver_eventFilter_default")
	symaddr(&addr_QAnimationDriver_timerEvent_default, "QAnimationDriver_timerEvent_default")
	symaddr(&addr_QAnimationDriver_childEvent_default, "QAnimationDriver_childEvent_default")
	symaddr(&addr_QAnimationDriver_customEvent_default, "QAnimationDriver_customEvent_default")
	symaddr(&addr_QAnimationDriver_connectNotify_default, "QAnimationDriver_connectNotify_default")
	symaddr(&addr_QAnimationDriver_disconnectNotify_default, "QAnimationDriver_disconnectNotify_default")
	reg(&pg_QAnimationDriver_connect_started, "QAnimationDriver_connect_started")
	reg(&pg_QAnimationDriver_connect_stopped, "QAnimationDriver_connect_stopped")
	reg(&pg_QAnimationDriver_connect_destroyed, "QAnimationDriver_connect_destroyed")
	reg(&pg_QAnimationDriver_connect_destroyed2, "QAnimationDriver_connect_destroyed2")
	reg(&pg_QAnimationDriver_connect_objectNameChanged, "QAnimationDriver_connect_objectNameChanged")
	reg(&pg_QAnimationDriver__disconnect, "QAnimationDriver__disconnect")
	reg(&pg_QAnimationDriver__set_override, "QAnimationDriver__set_override")
	reg(&pg_QAnimationDriver_delete, "QAnimationDriver_delete")
}

type QAnimationDriver struct {
	h uintptr
	*QObject
}

func bindQAnimationDriver(h uintptr) *QAnimationDriver {
	if h == 0 {
		return nil
	}
	return &QAnimationDriver{h: h, QObject: bindQObject(pg_QAnimationDriver_base(h))}
}

func newownedQAnimationDriver(h uintptr) *QAnimationDriver { return bindQAnimationDriver(h) }

func (this *QAnimationDriver) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAnimationDriver) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAnimationDriver(h unsafe.Pointer) *QAnimationDriver { return bindQAnimationDriver(uintptr(h)) }
func UnsafeNewQAnimationDriverOwned(h unsafe.Pointer) *QAnimationDriver { return newownedQAnimationDriver(uintptr(h)) }

func (this *QAnimationDriver) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQAnimationDriver() *QAnimationDriver {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_new, uintptr(__goid))
	__o := newownedQAnimationDriver(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAnimationDriver_virtIdx, QAnimationDriver_enable)
	return __o
}

func NewQAnimationDriver2(parent *QObject) *QAnimationDriver {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQAnimationDriver(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAnimationDriver_virtIdx, QAnimationDriver_enable)
	return __o
}

func QAnimationDriver_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAnimationDriver) Advance() {
	purego.SyscallN(addr_QAnimationDriver_advance, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) Install() {
	purego.SyscallN(addr_QAnimationDriver_install, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) Uninstall() {
	purego.SyscallN(addr_QAnimationDriver_uninstall, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) IsRunning() bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_isRunning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAnimationDriver) Elapsed() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_elapsed, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QAnimationDriver) Started() {
	purego.SyscallN(addr_QAnimationDriver_started, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) Stopped() {
	purego.SyscallN(addr_QAnimationDriver_stopped, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) AdvanceAnimation() {
	purego.SyscallN(addr_QAnimationDriver_advanceAnimation, this.h)
	runtime.KeepAlive(this)
}

func QAnimationDriver_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAnimationDriver_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAnimationDriver) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAnimationDriver) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAnimationDriver) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAnimationDriver) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAnimationDriver_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAnimationDriver_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAnimationDriver_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAnimationDriver_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAnimationDriver) AdvanceDefault() {
	purego.SyscallN(addr_QAnimationDriver_advance_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) ElapsedDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_elapsed_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QAnimationDriver) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAnimationDriver_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAnimationDriver) StartDefault() {
	purego.SyscallN(addr_QAnimationDriver_start_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) StopDefault() {
	purego.SyscallN(addr_QAnimationDriver_stop_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAnimationDriver) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAnimationDriver) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationDriver_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAnimationDriver) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAnimationDriver_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAnimationDriver) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAnimationDriver_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAnimationDriver) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAnimationDriver_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAnimationDriver) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAnimationDriver_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAnimationDriver) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAnimationDriver_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAnimationDriver) OnStarted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAnimationDriver_connect_started(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationDriver) OnStopped(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAnimationDriver_connect_stopped(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationDriver) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAnimationDriver_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationDriver) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAnimationDriver_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationDriver) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAnimationDriver_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationDriver) Disconnect(handle int64) {
	pg_QAnimationDriver__disconnect(handle)
	unregisterCallback(handle)
}

func QAnimationDriver_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -4830836341480073475:
		return 1
	case -1391384709503831749:
		return 2
	case -1270692749394500321:
		return 3
	case -5827849124968219451:
		return 4
	case 1237236456312824991:
		return 5
	case 5217623514630858297:
		return 6
	case -5409927581812374498:
		return 7
	case -1185315367526857833:
		return 8
	case 3769868229123868474:
		return 9
	case 7379029982170503008:
		return 10
	case -568369772982215962:
		return 11
	}
	return -1
}

func QAnimationDriver_enable(h uintptr, idx int32) { pg_QAnimationDriver__set_override(h, idx) }

func (this *QAnimationDriver) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAnimationDriver) OnAdvance(fn func()) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func() { fn() }) // slot "advance"
}

func (this *QAnimationDriver) OnElapsed(fn func() int64) {
	registerVirt(this.UnsafeGoID(), -1391384709503831749, func() int64 { return fn() }) // slot "elapsed"
}

func (this *QAnimationDriver) OnStart(fn func()) {
	registerVirt(this.UnsafeGoID(), -1270692749394500321, func() { fn() }) // slot "start"
}

func (this *QAnimationDriver) OnStop(fn func()) {
	registerVirt(this.UnsafeGoID(), -5827849124968219451, func() { fn() }) // slot "stop"
}

func (this *QAnimationDriver) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAnimationDriver) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAnimationDriver) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAnimationDriver) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAnimationDriver) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAnimationDriver) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAnimationDriver) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAnimationDriver) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAnimationDriver_delete(this.h)
	}
	this.UnsafeRelease()
}
