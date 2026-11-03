package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QEventLoop_base func(uintptr) uintptr
	addr_QEventLoop_new uintptr
	addr_QEventLoop_new2 uintptr
	addr_QEventLoop_tr uintptr
	addr_QEventLoop_processEvents uintptr
	addr_QEventLoop_processEvents2 uintptr
	addr_QEventLoop_processEvents3 uintptr
	addr_QEventLoop_exec uintptr
	addr_QEventLoop_isRunning uintptr
	addr_QEventLoop_wakeUp uintptr
	addr_QEventLoop_exit uintptr
	addr_QEventLoop_quit uintptr
	addr_QEventLoop_connect uintptr
	addr_QEventLoop_connect2 uintptr
	addr_QEventLoop_sender uintptr
	addr_QEventLoop_senderSignalIndex uintptr
	addr_QEventLoop_receivers uintptr
	addr_QEventLoop_isSignalConnected uintptr
	addr_QEventLoop_tr2 uintptr
	addr_QEventLoop_tr3 uintptr
	addr_QEventLoop_processEvents4 uintptr
	addr_QEventLoop_exec2 uintptr
	addr_QEventLoop_exit2 uintptr
	addr_QEventLoop_connect3 uintptr
	addr_QEventLoop_connect4 uintptr
	addr_QEventLoop_metaObject_default uintptr
	addr_QEventLoop_event_default uintptr
	addr_QEventLoop_eventFilter_default uintptr
	addr_QEventLoop_timerEvent_default uintptr
	addr_QEventLoop_childEvent_default uintptr
	addr_QEventLoop_customEvent_default uintptr
	addr_QEventLoop_connectNotify_default uintptr
	addr_QEventLoop_disconnectNotify_default uintptr
	pg_QEventLoop_connect_destroyed func(uintptr, int64)
	pg_QEventLoop_connect_destroyed2 func(uintptr, int64)
	pg_QEventLoop_connect_objectNameChanged func(uintptr, int64)
	pg_QEventLoop__disconnect func(int64)
	pg_QEventLoop__set_override func(uintptr, int32)
	pg_QEventLoop_delete func(uintptr)
)

func init() {
	reg(&pg_QEventLoop_base, "QEventLoop_base")
	symaddr(&addr_QEventLoop_new, "QEventLoop_new")
	symaddr(&addr_QEventLoop_new2, "QEventLoop_new2")
	symaddr(&addr_QEventLoop_tr, "QEventLoop_tr")
	symaddr(&addr_QEventLoop_processEvents, "QEventLoop_processEvents")
	symaddr(&addr_QEventLoop_processEvents2, "QEventLoop_processEvents2")
	symaddr(&addr_QEventLoop_processEvents3, "QEventLoop_processEvents3")
	symaddr(&addr_QEventLoop_exec, "QEventLoop_exec")
	symaddr(&addr_QEventLoop_isRunning, "QEventLoop_isRunning")
	symaddr(&addr_QEventLoop_wakeUp, "QEventLoop_wakeUp")
	symaddr(&addr_QEventLoop_exit, "QEventLoop_exit")
	symaddr(&addr_QEventLoop_quit, "QEventLoop_quit")
	symaddr(&addr_QEventLoop_connect, "QEventLoop_connect")
	symaddr(&addr_QEventLoop_connect2, "QEventLoop_connect2")
	symaddr(&addr_QEventLoop_sender, "QEventLoop_sender")
	symaddr(&addr_QEventLoop_senderSignalIndex, "QEventLoop_senderSignalIndex")
	symaddr(&addr_QEventLoop_receivers, "QEventLoop_receivers")
	symaddr(&addr_QEventLoop_isSignalConnected, "QEventLoop_isSignalConnected")
	symaddr(&addr_QEventLoop_tr2, "QEventLoop_tr2")
	symaddr(&addr_QEventLoop_tr3, "QEventLoop_tr3")
	symaddr(&addr_QEventLoop_processEvents4, "QEventLoop_processEvents4")
	symaddr(&addr_QEventLoop_exec2, "QEventLoop_exec2")
	symaddr(&addr_QEventLoop_exit2, "QEventLoop_exit2")
	symaddr(&addr_QEventLoop_connect3, "QEventLoop_connect3")
	symaddr(&addr_QEventLoop_connect4, "QEventLoop_connect4")
	symaddr(&addr_QEventLoop_metaObject_default, "QEventLoop_metaObject_default")
	symaddr(&addr_QEventLoop_event_default, "QEventLoop_event_default")
	symaddr(&addr_QEventLoop_eventFilter_default, "QEventLoop_eventFilter_default")
	symaddr(&addr_QEventLoop_timerEvent_default, "QEventLoop_timerEvent_default")
	symaddr(&addr_QEventLoop_childEvent_default, "QEventLoop_childEvent_default")
	symaddr(&addr_QEventLoop_customEvent_default, "QEventLoop_customEvent_default")
	symaddr(&addr_QEventLoop_connectNotify_default, "QEventLoop_connectNotify_default")
	symaddr(&addr_QEventLoop_disconnectNotify_default, "QEventLoop_disconnectNotify_default")
	reg(&pg_QEventLoop_connect_destroyed, "QEventLoop_connect_destroyed")
	reg(&pg_QEventLoop_connect_destroyed2, "QEventLoop_connect_destroyed2")
	reg(&pg_QEventLoop_connect_objectNameChanged, "QEventLoop_connect_objectNameChanged")
	reg(&pg_QEventLoop__disconnect, "QEventLoop__disconnect")
	reg(&pg_QEventLoop__set_override, "QEventLoop__set_override")
	reg(&pg_QEventLoop_delete, "QEventLoop_delete")
}

type QEventLoop struct {
	h uintptr
	*QObject
}

func bindQEventLoop(h uintptr) *QEventLoop {
	if h == 0 {
		return nil
	}
	return &QEventLoop{h: h, QObject: bindQObject(pg_QEventLoop_base(h))}
}

func newownedQEventLoop(h uintptr) *QEventLoop { return bindQEventLoop(h) }

func (this *QEventLoop) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QEventLoop) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQEventLoop(h unsafe.Pointer) *QEventLoop { return bindQEventLoop(uintptr(h)) }
func UnsafeNewQEventLoopOwned(h unsafe.Pointer) *QEventLoop { return newownedQEventLoop(uintptr(h)) }

func (this *QEventLoop) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQEventLoop() *QEventLoop {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QEventLoop_new, uintptr(__goid))
	__o := newownedQEventLoop(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QEventLoop_virtIdx, QEventLoop_enable)
	return __o
}

func NewQEventLoop2(parent *QObject) *QEventLoop {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QEventLoop_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQEventLoop(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QEventLoop_virtIdx, QEventLoop_enable)
	return __o
}

func QEventLoop_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QEventLoop_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QEventLoop) ProcessEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_processEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEventLoop) ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maximumTime int) {
	purego.SyscallN(addr_QEventLoop_processEvents2, this.h, uintptr(flags), uintptr(maximumTime))
	runtime.KeepAlive(this)
}

func (this *QEventLoop) ProcessEvents3(flags QEventLoop__ProcessEventsFlag, deadline *QDeadlineTimer) {
	purego.SyscallN(addr_QEventLoop_processEvents3, this.h, uintptr(flags), uintptr(deadline.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(deadline)
}

func (this *QEventLoop) Exec() int {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_exec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEventLoop) IsRunning() bool {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_isRunning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEventLoop) WakeUp() {
	purego.SyscallN(addr_QEventLoop_wakeUp, this.h)
	runtime.KeepAlive(this)
}

func (this *QEventLoop) Exit() {
	purego.SyscallN(addr_QEventLoop_exit, this.h)
	runtime.KeepAlive(this)
}

func (this *QEventLoop) Quit() {
	purego.SyscallN(addr_QEventLoop_quit, this.h)
	runtime.KeepAlive(this)
}

func QEventLoop_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QEventLoop_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QEventLoop_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QEventLoop_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QEventLoop) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QEventLoop_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QEventLoop) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEventLoop) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QEventLoop) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QEventLoop_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QEventLoop_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QEventLoop_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QEventLoop_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QEventLoop) ProcessEventsWithFlags(flags QEventLoop__ProcessEventsFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_processEvents4, this.h, uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEventLoop) ExecWithFlags(flags QEventLoop__ProcessEventsFlag) int {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_exec2, this.h, uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEventLoop) ExitWithReturnCode(returnCode int) {
	purego.SyscallN(addr_QEventLoop_exit2, this.h, uintptr(returnCode))
	runtime.KeepAlive(this)
}

func QEventLoop_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QEventLoop_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QEventLoop_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QEventLoop_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QEventLoop) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QEventLoop_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QEventLoop) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QEventLoop) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QEventLoop_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QEventLoop) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QEventLoop_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QEventLoop) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QEventLoop_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QEventLoop) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QEventLoop_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QEventLoop) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QEventLoop_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QEventLoop) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QEventLoop_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QEventLoop) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QEventLoop_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QEventLoop) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QEventLoop_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QEventLoop) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QEventLoop_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QEventLoop) Disconnect(handle int64) {
	pg_QEventLoop__disconnect(handle)
	unregisterCallback(handle)
}

func QEventLoop_virtIdx(slot int64) int32 {
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

func QEventLoop_enable(h uintptr, idx int32) { pg_QEventLoop__set_override(h, idx) }

func (this *QEventLoop) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QEventLoop) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QEventLoop) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QEventLoop) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QEventLoop) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QEventLoop) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QEventLoop) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QEventLoop) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QEventLoop) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QEventLoop_delete(this.h)
	}
	this.UnsafeRelease()
}
