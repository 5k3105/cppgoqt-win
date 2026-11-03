package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QThread_base func(uintptr) uintptr
	addr_QThread_new uintptr
	addr_QThread_new2 uintptr
	addr_QThread_tr uintptr
	addr_QThread_currentThread uintptr
	addr_QThread_isMainThread uintptr
	addr_QThread_idealThreadCount uintptr
	addr_QThread_yieldCurrentThread uintptr
	addr_QThread_setPriority uintptr
	addr_QThread_priority uintptr
	addr_QThread_isFinished uintptr
	addr_QThread_isRunning uintptr
	addr_QThread_requestInterruption uintptr
	addr_QThread_isInterruptionRequested uintptr
	addr_QThread_setStackSize uintptr
	addr_QThread_stackSize uintptr
	addr_QThread_eventDispatcher uintptr
	addr_QThread_setEventDispatcher uintptr
	addr_QThread_loopLevel uintptr
	addr_QThread_isCurrentThread uintptr
	addr_QThread_setServiceLevel uintptr
	addr_QThread_serviceLevel uintptr
	addr_QThread_start uintptr
	addr_QThread_terminate uintptr
	addr_QThread_exit uintptr
	addr_QThread_quit uintptr
	addr_QThread_wait uintptr
	addr_QThread_wait2 uintptr
	addr_QThread_sleep uintptr
	addr_QThread_msleep uintptr
	addr_QThread_usleep uintptr
	addr_QThread_sleep2 uintptr
	addr_QThread_exec uintptr
	addr_QThread_connect uintptr
	addr_QThread_connect2 uintptr
	addr_QThread_sender uintptr
	addr_QThread_senderSignalIndex uintptr
	addr_QThread_receivers uintptr
	addr_QThread_isSignalConnected uintptr
	addr_QThread_tr2 uintptr
	addr_QThread_tr3 uintptr
	addr_QThread_start2 uintptr
	addr_QThread_exit2 uintptr
	addr_QThread_wait3 uintptr
	addr_QThread_connect3 uintptr
	addr_QThread_connect4 uintptr
	addr_QThread_metaObject_default uintptr
	addr_QThread_event_default uintptr
	addr_QThread_run_default uintptr
	addr_QThread_eventFilter_default uintptr
	addr_QThread_timerEvent_default uintptr
	addr_QThread_childEvent_default uintptr
	addr_QThread_customEvent_default uintptr
	addr_QThread_connectNotify_default uintptr
	addr_QThread_disconnectNotify_default uintptr
	pg_QThread_connect_started func(uintptr, int64)
	pg_QThread_connect_finished func(uintptr, int64)
	pg_QThread_connect_destroyed func(uintptr, int64)
	pg_QThread_connect_destroyed2 func(uintptr, int64)
	pg_QThread_connect_objectNameChanged func(uintptr, int64)
	pg_QThread__disconnect func(int64)
	pg_QThread__set_override func(uintptr, int32)
	pg_QThread_delete func(uintptr)
)

func init() {
	reg(&pg_QThread_base, "QThread_base")
	symaddr(&addr_QThread_new, "QThread_new")
	symaddr(&addr_QThread_new2, "QThread_new2")
	symaddr(&addr_QThread_tr, "QThread_tr")
	symaddr(&addr_QThread_currentThread, "QThread_currentThread")
	symaddr(&addr_QThread_isMainThread, "QThread_isMainThread")
	symaddr(&addr_QThread_idealThreadCount, "QThread_idealThreadCount")
	symaddr(&addr_QThread_yieldCurrentThread, "QThread_yieldCurrentThread")
	symaddr(&addr_QThread_setPriority, "QThread_setPriority")
	symaddr(&addr_QThread_priority, "QThread_priority")
	symaddr(&addr_QThread_isFinished, "QThread_isFinished")
	symaddr(&addr_QThread_isRunning, "QThread_isRunning")
	symaddr(&addr_QThread_requestInterruption, "QThread_requestInterruption")
	symaddr(&addr_QThread_isInterruptionRequested, "QThread_isInterruptionRequested")
	symaddr(&addr_QThread_setStackSize, "QThread_setStackSize")
	symaddr(&addr_QThread_stackSize, "QThread_stackSize")
	symaddr(&addr_QThread_eventDispatcher, "QThread_eventDispatcher")
	symaddr(&addr_QThread_setEventDispatcher, "QThread_setEventDispatcher")
	symaddr(&addr_QThread_loopLevel, "QThread_loopLevel")
	symaddr(&addr_QThread_isCurrentThread, "QThread_isCurrentThread")
	symaddr(&addr_QThread_setServiceLevel, "QThread_setServiceLevel")
	symaddr(&addr_QThread_serviceLevel, "QThread_serviceLevel")
	symaddr(&addr_QThread_start, "QThread_start")
	symaddr(&addr_QThread_terminate, "QThread_terminate")
	symaddr(&addr_QThread_exit, "QThread_exit")
	symaddr(&addr_QThread_quit, "QThread_quit")
	symaddr(&addr_QThread_wait, "QThread_wait")
	symaddr(&addr_QThread_wait2, "QThread_wait2")
	symaddr(&addr_QThread_sleep, "QThread_sleep")
	symaddr(&addr_QThread_msleep, "QThread_msleep")
	symaddr(&addr_QThread_usleep, "QThread_usleep")
	symaddr(&addr_QThread_sleep2, "QThread_sleep2")
	symaddr(&addr_QThread_exec, "QThread_exec")
	symaddr(&addr_QThread_connect, "QThread_connect")
	symaddr(&addr_QThread_connect2, "QThread_connect2")
	symaddr(&addr_QThread_sender, "QThread_sender")
	symaddr(&addr_QThread_senderSignalIndex, "QThread_senderSignalIndex")
	symaddr(&addr_QThread_receivers, "QThread_receivers")
	symaddr(&addr_QThread_isSignalConnected, "QThread_isSignalConnected")
	symaddr(&addr_QThread_tr2, "QThread_tr2")
	symaddr(&addr_QThread_tr3, "QThread_tr3")
	symaddr(&addr_QThread_start2, "QThread_start2")
	symaddr(&addr_QThread_exit2, "QThread_exit2")
	symaddr(&addr_QThread_wait3, "QThread_wait3")
	symaddr(&addr_QThread_connect3, "QThread_connect3")
	symaddr(&addr_QThread_connect4, "QThread_connect4")
	symaddr(&addr_QThread_metaObject_default, "QThread_metaObject_default")
	symaddr(&addr_QThread_event_default, "QThread_event_default")
	symaddr(&addr_QThread_run_default, "QThread_run_default")
	symaddr(&addr_QThread_eventFilter_default, "QThread_eventFilter_default")
	symaddr(&addr_QThread_timerEvent_default, "QThread_timerEvent_default")
	symaddr(&addr_QThread_childEvent_default, "QThread_childEvent_default")
	symaddr(&addr_QThread_customEvent_default, "QThread_customEvent_default")
	symaddr(&addr_QThread_connectNotify_default, "QThread_connectNotify_default")
	symaddr(&addr_QThread_disconnectNotify_default, "QThread_disconnectNotify_default")
	reg(&pg_QThread_connect_started, "QThread_connect_started")
	reg(&pg_QThread_connect_finished, "QThread_connect_finished")
	reg(&pg_QThread_connect_destroyed, "QThread_connect_destroyed")
	reg(&pg_QThread_connect_destroyed2, "QThread_connect_destroyed2")
	reg(&pg_QThread_connect_objectNameChanged, "QThread_connect_objectNameChanged")
	reg(&pg_QThread__disconnect, "QThread__disconnect")
	reg(&pg_QThread__set_override, "QThread__set_override")
	reg(&pg_QThread_delete, "QThread_delete")
}

type QThread struct {
	h uintptr
	*QObject
}

func bindQThread(h uintptr) *QThread {
	if h == 0 {
		return nil
	}
	return &QThread{h: h, QObject: bindQObject(pg_QThread_base(h))}
}

func newownedQThread(h uintptr) *QThread { return bindQThread(h) }

func (this *QThread) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QThread) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQThread(h unsafe.Pointer) *QThread { return bindQThread(uintptr(h)) }
func UnsafeNewQThreadOwned(h unsafe.Pointer) *QThread { return newownedQThread(uintptr(h)) }

func (this *QThread) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQThread() *QThread {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QThread_new, uintptr(__goid))
	__o := newownedQThread(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QThread_virtIdx, QThread_enable)
	return __o
}

func NewQThread2(parent *QObject) *QThread {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QThread_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQThread(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QThread_virtIdx, QThread_enable)
	return __o
}

func QThread_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QThread_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QThread_CurrentThread() *QThread {
	__r, _, _ := purego.SyscallN(addr_QThread_currentThread)
	return bindQThread(__r)
}

func QThread_IsMainThread() bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_isMainThread)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QThread_IdealThreadCount() int {
	__r0, _, _ := purego.SyscallN(addr_QThread_idealThreadCount)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QThread_YieldCurrentThread() {
	purego.SyscallN(addr_QThread_yieldCurrentThread)
}

func (this *QThread) SetPriority(priority QThread__Priority) {
	purego.SyscallN(addr_QThread_setPriority, this.h, uintptr(priority))
	runtime.KeepAlive(this)
}

func (this *QThread) Priority() QThread__Priority {
	__r0, _, _ := purego.SyscallN(addr_QThread_priority, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QThread__Priority(__r)
}

func (this *QThread) IsFinished() bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_isFinished, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThread) IsRunning() bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_isRunning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThread) RequestInterruption() {
	purego.SyscallN(addr_QThread_requestInterruption, this.h)
	runtime.KeepAlive(this)
}

func (this *QThread) IsInterruptionRequested() bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_isInterruptionRequested, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThread) SetStackSize(stackSize uint) {
	purego.SyscallN(addr_QThread_setStackSize, this.h, uintptr(stackSize))
	runtime.KeepAlive(this)
}

func (this *QThread) StackSize() uint {
	__r0, _, _ := purego.SyscallN(addr_QThread_stackSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QThread) EventDispatcher() *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QThread_eventDispatcher, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractEventDispatcher(__r)
}

func (this *QThread) SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	purego.SyscallN(addr_QThread_setEventDispatcher, this.h, uintptr(eventDispatcher.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(eventDispatcher)
}

func (this *QThread) LoopLevel() int {
	__r0, _, _ := purego.SyscallN(addr_QThread_loopLevel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThread) IsCurrentThread() bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_isCurrentThread, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThread) SetServiceLevel(serviceLevel QThread__QualityOfService) {
	purego.SyscallN(addr_QThread_setServiceLevel, this.h, uintptr(serviceLevel))
	runtime.KeepAlive(this)
}

func (this *QThread) ServiceLevel() QThread__QualityOfService {
	__r0, _, _ := purego.SyscallN(addr_QThread_serviceLevel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QThread__QualityOfService(__r)
}

func (this *QThread) Start() {
	purego.SyscallN(addr_QThread_start, this.h)
	runtime.KeepAlive(this)
}

func (this *QThread) Terminate() {
	purego.SyscallN(addr_QThread_terminate, this.h)
	runtime.KeepAlive(this)
}

func (this *QThread) Exit() {
	purego.SyscallN(addr_QThread_exit, this.h)
	runtime.KeepAlive(this)
}

func (this *QThread) Quit() {
	purego.SyscallN(addr_QThread_quit, this.h)
	runtime.KeepAlive(this)
}

func (this *QThread) Wait() bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_wait, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QThread) WaitWithTime(time uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_wait2, this.h, uintptr(time))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QThread_Sleep(arg1 uint32) {
	purego.SyscallN(addr_QThread_sleep, uintptr(arg1))
}

func QThread_Msleep(arg1 uint32) {
	purego.SyscallN(addr_QThread_msleep, uintptr(arg1))
}

func QThread_Usleep(arg1 uint32) {
	purego.SyscallN(addr_QThread_usleep, uintptr(arg1))
}

func QThread_SleepWithNsec(nsec int64) {
	purego.SyscallN(addr_QThread_sleep2, uintptr(nsec))
}

func (this *QThread) Exec() int {
	__r0, _, _ := purego.SyscallN(addr_QThread_exec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QThread_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QThread_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QThread_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QThread_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QThread) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QThread_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QThread) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QThread_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThread) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QThread_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QThread) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QThread_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QThread_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QThread_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QThread_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QThread) StartWithQThreadPriority(arg1 QThread__Priority) {
	purego.SyscallN(addr_QThread_start2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QThread) ExitWithRetcode(retcode int) {
	purego.SyscallN(addr_QThread_exit2, this.h, uintptr(retcode))
	runtime.KeepAlive(this)
}

func (this *QThread) WaitWithDeadline(deadline *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_wait3, this.h, uintptr(deadline.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(deadline)
	return __r != 0
}

func QThread_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QThread_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QThread_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QThread_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QThread) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QThread_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QThread) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QThread) RunDefault() {
	purego.SyscallN(addr_QThread_run_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QThread) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QThread_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QThread) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QThread_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QThread) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QThread_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QThread) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QThread_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QThread) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QThread_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QThread) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QThread_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QThread) OnStarted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QThread_connect_started(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThread) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QThread_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThread) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QThread_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThread) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QThread_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThread) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QThread_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QThread) Disconnect(handle int64) {
	pg_QThread__disconnect(handle)
	unregisterCallback(handle)
}

func QThread_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case -8524566317707732406:
		return 2
	case 5217623514630858297:
		return 3
	case -5409927581812374498:
		return 4
	case -1185315367526857833:
		return 5
	case 3769868229123868474:
		return 6
	case 7379029982170503008:
		return 7
	case -568369772982215962:
		return 8
	}
	return -1
}

func QThread_enable(h uintptr, idx int32) { pg_QThread__set_override(h, idx) }

func (this *QThread) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QThread) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QThread) OnRun(fn func()) {
	registerVirt(this.UnsafeGoID(), -8524566317707732406, func() { fn() }) // slot "run"
}

func (this *QThread) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QThread) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QThread) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QThread) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QThread) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QThread) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QThread) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QThread_delete(this.h)
	}
	this.UnsafeRelease()
}
