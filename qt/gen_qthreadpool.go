package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QThreadPool_base func(uintptr) uintptr
	addr_QThreadPool_new uintptr
	addr_QThreadPool_new2 uintptr
	addr_QThreadPool_tr uintptr
	addr_QThreadPool_globalInstance uintptr
	addr_QThreadPool_start uintptr
	addr_QThreadPool_tryStart uintptr
	addr_QThreadPool_startOnReservedThread uintptr
	addr_QThreadPool_expiryTimeout uintptr
	addr_QThreadPool_setExpiryTimeout uintptr
	addr_QThreadPool_maxThreadCount uintptr
	addr_QThreadPool_setMaxThreadCount uintptr
	addr_QThreadPool_activeThreadCount uintptr
	addr_QThreadPool_setStackSize uintptr
	addr_QThreadPool_stackSize uintptr
	addr_QThreadPool_setThreadPriority uintptr
	addr_QThreadPool_threadPriority uintptr
	addr_QThreadPool_reserveThread uintptr
	addr_QThreadPool_releaseThread uintptr
	addr_QThreadPool_setServiceLevel uintptr
	addr_QThreadPool_serviceLevel uintptr
	addr_QThreadPool_waitForDone uintptr
	addr_QThreadPool_waitForDone2 uintptr
	addr_QThreadPool_clear uintptr
	addr_QThreadPool_contains uintptr
	addr_QThreadPool_tryTake uintptr
	addr_QThreadPool_connect uintptr
	addr_QThreadPool_connect2 uintptr
	addr_QThreadPool_sender uintptr
	addr_QThreadPool_senderSignalIndex uintptr
	addr_QThreadPool_receivers uintptr
	addr_QThreadPool_isSignalConnected uintptr
	addr_QThreadPool_tr2 uintptr
	addr_QThreadPool_tr3 uintptr
	addr_QThreadPool_start2 uintptr
	addr_QThreadPool_waitForDone3 uintptr
	addr_QThreadPool_connect3 uintptr
	addr_QThreadPool_connect4 uintptr
	addr_QThreadPool_metaObject_default uintptr
	addr_QThreadPool_event_default uintptr
	addr_QThreadPool_eventFilter_default uintptr
	addr_QThreadPool_timerEvent_default uintptr
	addr_QThreadPool_childEvent_default uintptr
	addr_QThreadPool_customEvent_default uintptr
	addr_QThreadPool_connectNotify_default uintptr
	addr_QThreadPool_disconnectNotify_default uintptr
	pg_QThreadPool_connect_destroyed func(uintptr, int64)
	pg_QThreadPool_connect_destroyed2 func(uintptr, int64)
	pg_QThreadPool_connect_objectNameChanged func(uintptr, int64)
	pg_QThreadPool__disconnect func(int64)
	pg_QThreadPool__set_override func(uintptr, int32)
	pg_QThreadPool_delete func(uintptr)
)

func init() {
	reg(&pg_QThreadPool_base, "QThreadPool_base")
	symaddr(&addr_QThreadPool_new, "QThreadPool_new")
	symaddr(&addr_QThreadPool_new2, "QThreadPool_new2")
	symaddr(&addr_QThreadPool_tr, "QThreadPool_tr")
	symaddr(&addr_QThreadPool_globalInstance, "QThreadPool_globalInstance")
	symaddr(&addr_QThreadPool_start, "QThreadPool_start")
	symaddr(&addr_QThreadPool_tryStart, "QThreadPool_tryStart")
	symaddr(&addr_QThreadPool_startOnReservedThread, "QThreadPool_startOnReservedThread")
	symaddr(&addr_QThreadPool_expiryTimeout, "QThreadPool_expiryTimeout")
	symaddr(&addr_QThreadPool_setExpiryTimeout, "QThreadPool_setExpiryTimeout")
	symaddr(&addr_QThreadPool_maxThreadCount, "QThreadPool_maxThreadCount")
	symaddr(&addr_QThreadPool_setMaxThreadCount, "QThreadPool_setMaxThreadCount")
	symaddr(&addr_QThreadPool_activeThreadCount, "QThreadPool_activeThreadCount")
	symaddr(&addr_QThreadPool_setStackSize, "QThreadPool_setStackSize")
	symaddr(&addr_QThreadPool_stackSize, "QThreadPool_stackSize")
	symaddr(&addr_QThreadPool_setThreadPriority, "QThreadPool_setThreadPriority")
	symaddr(&addr_QThreadPool_threadPriority, "QThreadPool_threadPriority")
	symaddr(&addr_QThreadPool_reserveThread, "QThreadPool_reserveThread")
	symaddr(&addr_QThreadPool_releaseThread, "QThreadPool_releaseThread")
	symaddr(&addr_QThreadPool_setServiceLevel, "QThreadPool_setServiceLevel")
	symaddr(&addr_QThreadPool_serviceLevel, "QThreadPool_serviceLevel")
	symaddr(&addr_QThreadPool_waitForDone, "QThreadPool_waitForDone")
	symaddr(&addr_QThreadPool_waitForDone2, "QThreadPool_waitForDone2")
	symaddr(&addr_QThreadPool_clear, "QThreadPool_clear")
	symaddr(&addr_QThreadPool_contains, "QThreadPool_contains")
	symaddr(&addr_QThreadPool_tryTake, "QThreadPool_tryTake")
	symaddr(&addr_QThreadPool_connect, "QThreadPool_connect")
	symaddr(&addr_QThreadPool_connect2, "QThreadPool_connect2")
	symaddr(&addr_QThreadPool_sender, "QThreadPool_sender")
	symaddr(&addr_QThreadPool_senderSignalIndex, "QThreadPool_senderSignalIndex")
	symaddr(&addr_QThreadPool_receivers, "QThreadPool_receivers")
	symaddr(&addr_QThreadPool_isSignalConnected, "QThreadPool_isSignalConnected")
	symaddr(&addr_QThreadPool_tr2, "QThreadPool_tr2")
	symaddr(&addr_QThreadPool_tr3, "QThreadPool_tr3")
	symaddr(&addr_QThreadPool_start2, "QThreadPool_start2")
	symaddr(&addr_QThreadPool_waitForDone3, "QThreadPool_waitForDone3")
	symaddr(&addr_QThreadPool_connect3, "QThreadPool_connect3")
	symaddr(&addr_QThreadPool_connect4, "QThreadPool_connect4")
	symaddr(&addr_QThreadPool_metaObject_default, "QThreadPool_metaObject_default")
	symaddr(&addr_QThreadPool_event_default, "QThreadPool_event_default")
	symaddr(&addr_QThreadPool_eventFilter_default, "QThreadPool_eventFilter_default")
	symaddr(&addr_QThreadPool_timerEvent_default, "QThreadPool_timerEvent_default")
	symaddr(&addr_QThreadPool_childEvent_default, "QThreadPool_childEvent_default")
	symaddr(&addr_QThreadPool_customEvent_default, "QThreadPool_customEvent_default")
	symaddr(&addr_QThreadPool_connectNotify_default, "QThreadPool_connectNotify_default")
	symaddr(&addr_QThreadPool_disconnectNotify_default, "QThreadPool_disconnectNotify_default")
	reg(&pg_QThreadPool_connect_destroyed, "QThreadPool_connect_destroyed")
	reg(&pg_QThreadPool_connect_destroyed2, "QThreadPool_connect_destroyed2")
	reg(&pg_QThreadPool_connect_objectNameChanged, "QThreadPool_connect_objectNameChanged")
	reg(&pg_QThreadPool__disconnect, "QThreadPool__disconnect")
	reg(&pg_QThreadPool__set_override, "QThreadPool__set_override")
	reg(&pg_QThreadPool_delete, "QThreadPool_delete")
}

type QThreadPool struct {
	h uintptr
	*QObject
}

func bindQThreadPool(h uintptr) *QThreadPool {
	if h == 0 {
		return nil
	}
	return &QThreadPool{h: h, QObject: bindQObject(pg_QThreadPool_base(h))}
}

func newownedQThreadPool(h uintptr) *QThreadPool { return bindQThreadPool(h) }

func (this *QThreadPool) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QThreadPool) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQThreadPool(h unsafe.Pointer) *QThreadPool { return bindQThreadPool(uintptr(h)) }
func UnsafeNewQThreadPoolOwned(h unsafe.Pointer) *QThreadPool { return newownedQThreadPool(uintptr(h)) }

func (this *QThreadPool) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQThreadPool() *QThreadPool {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QThreadPool_new, uintptr(__goid))
	__o := newownedQThreadPool(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QThreadPool_virtIdx, QThreadPool_enable)
	return __o
}

func NewQThreadPool2(parent *QObject) *QThreadPool {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QThreadPool_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQThreadPool(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QThreadPool_virtIdx, QThreadPool_enable)
	return __o
}

func QThreadPool_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QThreadPool_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QThreadPool_GlobalInstance() *QThreadPool {
	__r, _, _ := purego.SyscallN(addr_QThreadPool_globalInstance)
	return bindQThreadPool(__r)
}

func (this *QThreadPool) Start(runnable *QRunnable) {
	purego.SyscallN(addr_QThreadPool_start, this.h, uintptr(runnable.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(runnable)
}

func (this *QThreadPool) TryStart(runnable *QRunnable) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_tryStart, this.h, uintptr(runnable.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(runnable)
	return __r != 0
}

func (this *QThreadPool) StartOnReservedThread(runnable *QRunnable) {
	purego.SyscallN(addr_QThreadPool_startOnReservedThread, this.h, uintptr(runnable.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(runnable)
}

func (this *QThreadPool) ExpiryTimeout() int {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_expiryTimeout, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	purego.SyscallN(addr_QThreadPool_setExpiryTimeout, this.h, uintptr(expiryTimeout))
	runtime.KeepAlive(this)
}

func (this *QThreadPool) MaxThreadCount() int {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_maxThreadCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	purego.SyscallN(addr_QThreadPool_setMaxThreadCount, this.h, uintptr(maxThreadCount))
	runtime.KeepAlive(this)
}

func (this *QThreadPool) ActiveThreadCount() int {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_activeThreadCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThreadPool) SetStackSize(stackSize uint) {
	purego.SyscallN(addr_QThreadPool_setStackSize, this.h, uintptr(stackSize))
	runtime.KeepAlive(this)
}

func (this *QThreadPool) StackSize() uint {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_stackSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QThreadPool) SetThreadPriority(priority QThread__Priority) {
	purego.SyscallN(addr_QThreadPool_setThreadPriority, this.h, uintptr(priority))
	runtime.KeepAlive(this)
}

func (this *QThreadPool) ThreadPriority() QThread__Priority {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_threadPriority, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QThread__Priority(__r)
}

func (this *QThreadPool) ReserveThread() {
	purego.SyscallN(addr_QThreadPool_reserveThread, this.h)
	runtime.KeepAlive(this)
}

func (this *QThreadPool) ReleaseThread() {
	purego.SyscallN(addr_QThreadPool_releaseThread, this.h)
	runtime.KeepAlive(this)
}

func (this *QThreadPool) SetServiceLevel(serviceLevel QThread__QualityOfService) {
	purego.SyscallN(addr_QThreadPool_setServiceLevel, this.h, uintptr(serviceLevel))
	runtime.KeepAlive(this)
}

func (this *QThreadPool) ServiceLevel() QThread__QualityOfService {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_serviceLevel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QThread__QualityOfService(__r)
}

func (this *QThreadPool) WaitForDone(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_waitForDone, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThreadPool) WaitForDone2() bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_waitForDone2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThreadPool) Clear() {
	purego.SyscallN(addr_QThreadPool_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QThreadPool) Contains(thread *QThread) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_contains, this.h, uintptr(thread.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(thread)
	return __r != 0
}

func (this *QThreadPool) TryTake(runnable *QRunnable) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_tryTake, this.h, uintptr(runnable.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(runnable)
	return __r != 0
}

func QThreadPool_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QThreadPool_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QThreadPool_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QThreadPool_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QThreadPool) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QThreadPool_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QThreadPool) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThreadPool) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QThreadPool) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QThreadPool_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QThreadPool_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QThreadPool_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QThreadPool_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QThreadPool) Start2(runnable *QRunnable, priority int) {
	purego.SyscallN(addr_QThreadPool_start2, this.h, uintptr(runnable.UnsafePointer()), uintptr(priority))
	runtime.KeepAlive(this)
	runtime.KeepAlive(runnable)
}

func (this *QThreadPool) WaitForDoneWithDeadline(deadline *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_waitForDone3, this.h, uintptr(deadline.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(deadline)
	return __r != 0
}

func QThreadPool_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QThreadPool_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QThreadPool_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QThreadPool_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QThreadPool) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QThreadPool_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QThreadPool) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QThreadPool) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QThreadPool_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QThreadPool) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QThreadPool_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QThreadPool) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QThreadPool_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QThreadPool) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QThreadPool_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QThreadPool) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QThreadPool_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QThreadPool) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QThreadPool_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QThreadPool) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QThreadPool_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThreadPool) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QThreadPool_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThreadPool) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QThreadPool_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThreadPool) Disconnect(handle int64) {
	pg_QThreadPool__disconnect(handle)
	unregisterCallback(handle)
}

func QThreadPool_virtIdx(slot int64) int32 {
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

func QThreadPool_enable(h uintptr, idx int32) { pg_QThreadPool__set_override(h, idx) }

func (this *QThreadPool) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QThreadPool) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QThreadPool) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QThreadPool) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QThreadPool) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QThreadPool) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QThreadPool) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QThreadPool) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QThreadPool) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QThreadPool_delete(this.h)
	}
	this.UnsafeRelease()
}
