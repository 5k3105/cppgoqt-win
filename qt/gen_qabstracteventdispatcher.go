package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAbstractEventDispatcher_base func(uintptr) uintptr
	addr_QAbstractEventDispatcher_tr uintptr
	addr_QAbstractEventDispatcher_instance uintptr
	addr_QAbstractEventDispatcher_processEvents uintptr
	addr_QAbstractEventDispatcher_registerSocketNotifier uintptr
	addr_QAbstractEventDispatcher_unregisterSocketNotifier uintptr
	addr_QAbstractEventDispatcher_registerTimer uintptr
	addr_QAbstractEventDispatcher_registerTimer2 uintptr
	addr_QAbstractEventDispatcher_unregisterTimer uintptr
	addr_QAbstractEventDispatcher_unregisterTimers uintptr
	addr_QAbstractEventDispatcher_registeredTimers uintptr
	addr_QAbstractEventDispatcher_remainingTime uintptr
	addr_QAbstractEventDispatcher_unregisterTimer2 uintptr
	addr_QAbstractEventDispatcher_timersForObject uintptr
	addr_QAbstractEventDispatcher_wakeUp uintptr
	addr_QAbstractEventDispatcher_interrupt uintptr
	addr_QAbstractEventDispatcher_startingUp uintptr
	addr_QAbstractEventDispatcher_closingDown uintptr
	addr_QAbstractEventDispatcher_installNativeEventFilter uintptr
	addr_QAbstractEventDispatcher_removeNativeEventFilter uintptr
	addr_QAbstractEventDispatcher_aboutToBlock uintptr
	addr_QAbstractEventDispatcher_awake uintptr
	addr_QAbstractEventDispatcher_connect uintptr
	addr_QAbstractEventDispatcher_connect2 uintptr
	addr_QAbstractEventDispatcher_sender uintptr
	addr_QAbstractEventDispatcher_senderSignalIndex uintptr
	addr_QAbstractEventDispatcher_receivers uintptr
	addr_QAbstractEventDispatcher_isSignalConnected uintptr
	addr_QAbstractEventDispatcher_tr2 uintptr
	addr_QAbstractEventDispatcher_tr3 uintptr
	addr_QAbstractEventDispatcher_instance2 uintptr
	addr_QAbstractEventDispatcher_connect3 uintptr
	addr_QAbstractEventDispatcher_connect4 uintptr
	addr_QAbstractEventDispatcher_startingUp_default uintptr
	addr_QAbstractEventDispatcher_closingDown_default uintptr
	addr_QAbstractEventDispatcher_metaObject_default uintptr
	addr_QAbstractEventDispatcher_event_default uintptr
	addr_QAbstractEventDispatcher_eventFilter_default uintptr
	addr_QAbstractEventDispatcher_timerEvent_default uintptr
	addr_QAbstractEventDispatcher_childEvent_default uintptr
	addr_QAbstractEventDispatcher_customEvent_default uintptr
	addr_QAbstractEventDispatcher_connectNotify_default uintptr
	addr_QAbstractEventDispatcher_disconnectNotify_default uintptr
	pg_QAbstractEventDispatcher_connect_aboutToBlock func(uintptr, int64)
	pg_QAbstractEventDispatcher_connect_awake func(uintptr, int64)
	pg_QAbstractEventDispatcher_connect_destroyed func(uintptr, int64)
	pg_QAbstractEventDispatcher_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractEventDispatcher_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractEventDispatcher__disconnect func(int64)
	pg_QAbstractEventDispatcher__set_override func(uintptr, int32)
	pg_QAbstractEventDispatcher__Go_new func(int64, uintptr) uintptr
	pg_QAbstractEventDispatcher_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractEventDispatcher_base, "QAbstractEventDispatcher_base")
	symaddr(&addr_QAbstractEventDispatcher_tr, "QAbstractEventDispatcher_tr")
	symaddr(&addr_QAbstractEventDispatcher_instance, "QAbstractEventDispatcher_instance")
	symaddr(&addr_QAbstractEventDispatcher_processEvents, "QAbstractEventDispatcher_processEvents")
	symaddr(&addr_QAbstractEventDispatcher_registerSocketNotifier, "QAbstractEventDispatcher_registerSocketNotifier")
	symaddr(&addr_QAbstractEventDispatcher_unregisterSocketNotifier, "QAbstractEventDispatcher_unregisterSocketNotifier")
	symaddr(&addr_QAbstractEventDispatcher_registerTimer, "QAbstractEventDispatcher_registerTimer")
	symaddr(&addr_QAbstractEventDispatcher_registerTimer2, "QAbstractEventDispatcher_registerTimer2")
	symaddr(&addr_QAbstractEventDispatcher_unregisterTimer, "QAbstractEventDispatcher_unregisterTimer")
	symaddr(&addr_QAbstractEventDispatcher_unregisterTimers, "QAbstractEventDispatcher_unregisterTimers")
	symaddr(&addr_QAbstractEventDispatcher_registeredTimers, "QAbstractEventDispatcher_registeredTimers")
	symaddr(&addr_QAbstractEventDispatcher_remainingTime, "QAbstractEventDispatcher_remainingTime")
	symaddr(&addr_QAbstractEventDispatcher_unregisterTimer2, "QAbstractEventDispatcher_unregisterTimer2")
	symaddr(&addr_QAbstractEventDispatcher_timersForObject, "QAbstractEventDispatcher_timersForObject")
	symaddr(&addr_QAbstractEventDispatcher_wakeUp, "QAbstractEventDispatcher_wakeUp")
	symaddr(&addr_QAbstractEventDispatcher_interrupt, "QAbstractEventDispatcher_interrupt")
	symaddr(&addr_QAbstractEventDispatcher_startingUp, "QAbstractEventDispatcher_startingUp")
	symaddr(&addr_QAbstractEventDispatcher_closingDown, "QAbstractEventDispatcher_closingDown")
	symaddr(&addr_QAbstractEventDispatcher_installNativeEventFilter, "QAbstractEventDispatcher_installNativeEventFilter")
	symaddr(&addr_QAbstractEventDispatcher_removeNativeEventFilter, "QAbstractEventDispatcher_removeNativeEventFilter")
	symaddr(&addr_QAbstractEventDispatcher_aboutToBlock, "QAbstractEventDispatcher_aboutToBlock")
	symaddr(&addr_QAbstractEventDispatcher_awake, "QAbstractEventDispatcher_awake")
	symaddr(&addr_QAbstractEventDispatcher_connect, "QAbstractEventDispatcher_connect")
	symaddr(&addr_QAbstractEventDispatcher_connect2, "QAbstractEventDispatcher_connect2")
	symaddr(&addr_QAbstractEventDispatcher_sender, "QAbstractEventDispatcher_sender")
	symaddr(&addr_QAbstractEventDispatcher_senderSignalIndex, "QAbstractEventDispatcher_senderSignalIndex")
	symaddr(&addr_QAbstractEventDispatcher_receivers, "QAbstractEventDispatcher_receivers")
	symaddr(&addr_QAbstractEventDispatcher_isSignalConnected, "QAbstractEventDispatcher_isSignalConnected")
	symaddr(&addr_QAbstractEventDispatcher_tr2, "QAbstractEventDispatcher_tr2")
	symaddr(&addr_QAbstractEventDispatcher_tr3, "QAbstractEventDispatcher_tr3")
	symaddr(&addr_QAbstractEventDispatcher_instance2, "QAbstractEventDispatcher_instance2")
	symaddr(&addr_QAbstractEventDispatcher_connect3, "QAbstractEventDispatcher_connect3")
	symaddr(&addr_QAbstractEventDispatcher_connect4, "QAbstractEventDispatcher_connect4")
	symaddr(&addr_QAbstractEventDispatcher_startingUp_default, "QAbstractEventDispatcher_startingUp_default")
	symaddr(&addr_QAbstractEventDispatcher_closingDown_default, "QAbstractEventDispatcher_closingDown_default")
	symaddr(&addr_QAbstractEventDispatcher_metaObject_default, "QAbstractEventDispatcher_metaObject_default")
	symaddr(&addr_QAbstractEventDispatcher_event_default, "QAbstractEventDispatcher_event_default")
	symaddr(&addr_QAbstractEventDispatcher_eventFilter_default, "QAbstractEventDispatcher_eventFilter_default")
	symaddr(&addr_QAbstractEventDispatcher_timerEvent_default, "QAbstractEventDispatcher_timerEvent_default")
	symaddr(&addr_QAbstractEventDispatcher_childEvent_default, "QAbstractEventDispatcher_childEvent_default")
	symaddr(&addr_QAbstractEventDispatcher_customEvent_default, "QAbstractEventDispatcher_customEvent_default")
	symaddr(&addr_QAbstractEventDispatcher_connectNotify_default, "QAbstractEventDispatcher_connectNotify_default")
	symaddr(&addr_QAbstractEventDispatcher_disconnectNotify_default, "QAbstractEventDispatcher_disconnectNotify_default")
	reg(&pg_QAbstractEventDispatcher_connect_aboutToBlock, "QAbstractEventDispatcher_connect_aboutToBlock")
	reg(&pg_QAbstractEventDispatcher_connect_awake, "QAbstractEventDispatcher_connect_awake")
	reg(&pg_QAbstractEventDispatcher_connect_destroyed, "QAbstractEventDispatcher_connect_destroyed")
	reg(&pg_QAbstractEventDispatcher_connect_destroyed2, "QAbstractEventDispatcher_connect_destroyed2")
	reg(&pg_QAbstractEventDispatcher_connect_objectNameChanged, "QAbstractEventDispatcher_connect_objectNameChanged")
	reg(&pg_QAbstractEventDispatcher__disconnect, "QAbstractEventDispatcher__disconnect")
	reg(&pg_QAbstractEventDispatcher__set_override, "QAbstractEventDispatcher__set_override")
	reg(&pg_QAbstractEventDispatcher__Go_new, "QAbstractEventDispatcher__Go_new")
	reg(&pg_QAbstractEventDispatcher_delete, "QAbstractEventDispatcher_delete")
}

type QAbstractEventDispatcher struct {
	h uintptr
	*QObject
}

func bindQAbstractEventDispatcher(h uintptr) *QAbstractEventDispatcher {
	if h == 0 {
		return nil
	}
	return &QAbstractEventDispatcher{h: h, QObject: bindQObject(pg_QAbstractEventDispatcher_base(h))}
}

func newownedQAbstractEventDispatcher(h uintptr) *QAbstractEventDispatcher { return bindQAbstractEventDispatcher(h) }

func (this *QAbstractEventDispatcher) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractEventDispatcher) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractEventDispatcher(h unsafe.Pointer) *QAbstractEventDispatcher { return bindQAbstractEventDispatcher(uintptr(h)) }
func UnsafeNewQAbstractEventDispatcherOwned(h unsafe.Pointer) *QAbstractEventDispatcher { return newownedQAbstractEventDispatcher(uintptr(h)) }

func (this *QAbstractEventDispatcher) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAbstractEventDispatcher_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QAbstractEventDispatcher_Instance() *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_instance)
	return bindQAbstractEventDispatcher(__r)
}

func (this *QAbstractEventDispatcher) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_processEvents, this.h, uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractEventDispatcher) RegisterSocketNotifier(notifier *QSocketNotifier) {
	purego.SyscallN(addr_QAbstractEventDispatcher_registerSocketNotifier, this.h, uintptr(notifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(notifier)
}

func (this *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier *QSocketNotifier) {
	purego.SyscallN(addr_QAbstractEventDispatcher_unregisterSocketNotifier, this.h, uintptr(notifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(notifier)
}

func (this *QAbstractEventDispatcher) RegisterTimer(interval int64, timerType TimerType, object *QObject) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_registerTimer, this.h, uintptr(interval), uintptr(timerType), uintptr(object.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	return int(__r)
}

func (this *QAbstractEventDispatcher) RegisterTimer2(timerId int, interval int64, timerType TimerType, object *QObject) {
	purego.SyscallN(addr_QAbstractEventDispatcher_registerTimer2, this.h, uintptr(timerId), uintptr(interval), uintptr(timerType), uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
}

func (this *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_unregisterTimer, this.h, uintptr(timerId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractEventDispatcher) UnregisterTimers(object *QObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_unregisterTimers, this.h, uintptr(object.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	return __r != 0
}

func (this *QAbstractEventDispatcher) RegisteredTimers(object *QObject) []*QAbstractEventDispatcher__TimerInfo {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_registeredTimers, this.h, uintptr(object.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	__out := make([]*QAbstractEventDispatcher__TimerInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQAbstractEventDispatcher__TimerInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_remainingTime, this.h, uintptr(timerId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractEventDispatcher) UnregisterTimerWithTimerId(timerId TimerId) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_unregisterTimer2, this.h, uintptr(timerId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractEventDispatcher) TimersForObject(object *QObject) []*QAbstractEventDispatcher__TimerInfoV2 {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_timersForObject, this.h, uintptr(object.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	__out := make([]*QAbstractEventDispatcher__TimerInfoV2, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQAbstractEventDispatcher__TimerInfoV2(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractEventDispatcher) WakeUp() {
	purego.SyscallN(addr_QAbstractEventDispatcher_wakeUp, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) Interrupt() {
	purego.SyscallN(addr_QAbstractEventDispatcher_interrupt, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) StartingUp() {
	purego.SyscallN(addr_QAbstractEventDispatcher_startingUp, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) ClosingDown() {
	purego.SyscallN(addr_QAbstractEventDispatcher_closingDown, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	purego.SyscallN(addr_QAbstractEventDispatcher_installNativeEventFilter, this.h, uintptr(filterObj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterObj)
}

func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	purego.SyscallN(addr_QAbstractEventDispatcher_removeNativeEventFilter, this.h, uintptr(filterObj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterObj)
}

func (this *QAbstractEventDispatcher) AboutToBlock() {
	purego.SyscallN(addr_QAbstractEventDispatcher_aboutToBlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) Awake() {
	purego.SyscallN(addr_QAbstractEventDispatcher_awake, this.h)
	runtime.KeepAlive(this)
}

func QAbstractEventDispatcher_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractEventDispatcher_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractEventDispatcher) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractEventDispatcher) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractEventDispatcher) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractEventDispatcher) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractEventDispatcher_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractEventDispatcher_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractEventDispatcher_InstanceWithThread(thread *QThread) *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_instance2, uintptr(thread.UnsafePointer()))
	runtime.KeepAlive(thread)
	return bindQAbstractEventDispatcher(__r)
}

func QAbstractEventDispatcher_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractEventDispatcher_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractEventDispatcher) StartingUpDefault() {
	purego.SyscallN(addr_QAbstractEventDispatcher_startingUp_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) ClosingDownDefault() {
	purego.SyscallN(addr_QAbstractEventDispatcher_closingDown_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractEventDispatcher) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractEventDispatcher) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractEventDispatcher) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractEventDispatcher_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractEventDispatcher) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractEventDispatcher_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractEventDispatcher) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractEventDispatcher_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractEventDispatcher) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractEventDispatcher_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractEventDispatcher) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractEventDispatcher_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractEventDispatcher) OnAboutToBlock(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractEventDispatcher_connect_aboutToBlock(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcher) OnAwake(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractEventDispatcher_connect_awake(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcher) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractEventDispatcher_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcher) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractEventDispatcher_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcher) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractEventDispatcher_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcher) Disconnect(handle int64) {
	pg_QAbstractEventDispatcher__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractEventDispatcher_virtIdx(slot int64) int32 {
	switch slot {
	case -2075143971860206337:
		return 0
	case 849885779796930167:
		return 1
	case 5438220604201114580:
		return 2
	case 5345355196729094449:
		return 3
	case -85727209796773320:
		return 4
	case 8418593569725967729:
		return 5
	case 776599310866639255:
		return 6
	case -7776574712936801734:
		return 7
	case -6649993720439713720:
		return 8
	case -3579312661142054576:
		return 9
	case 1607084412591819503:
		return 10
	case -4418351852319543908:
		return 11
	case -8213485174827668208:
		return 12
	case 1237236456312824991:
		return 13
	case 5217623514630858297:
		return 14
	case -5409927581812374498:
		return 15
	case -1185315367526857833:
		return 16
	case 3769868229123868474:
		return 17
	case 7379029982170503008:
		return 18
	case -568369772982215962:
		return 19
	}
	return -1
}

func QAbstractEventDispatcher_enable(h uintptr, idx int32) { pg_QAbstractEventDispatcher__set_override(h, idx) }

func NewQAbstractEventDispatcher() *QAbstractEventDispatcher {
	id := nextVirtID()
	o := newownedQAbstractEventDispatcher(pg_QAbstractEventDispatcher__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractEventDispatcher_virtIdx, QAbstractEventDispatcher_enable)
	return o
}

func NewQAbstractEventDispatcher2(parent *QObject) *QAbstractEventDispatcher {
	id := nextVirtID()
	o := newownedQAbstractEventDispatcher(pg_QAbstractEventDispatcher__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractEventDispatcher_virtIdx, QAbstractEventDispatcher_enable)
	return o
}

func (this *QAbstractEventDispatcher) OnProcessEvents(fn func(QEventLoop__ProcessEventsFlag) bool) {
	registerVirt(this.UnsafeGoID(), -2075143971860206337, func(flags int32) int32 { return b2i(fn(QEventLoop__ProcessEventsFlag(flags))) }) // slot "processEvents"
}

func (this *QAbstractEventDispatcher) OnRegisterSocketNotifier(fn func(*QSocketNotifier)) {
	registerVirt(this.UnsafeGoID(), 849885779796930167, func(notifier uintptr) { fn(UnsafeNewQSocketNotifier(unsafe.Pointer(notifier))) }) // slot "registerSocketNotifier"
}

func (this *QAbstractEventDispatcher) OnUnregisterSocketNotifier(fn func(*QSocketNotifier)) {
	registerVirt(this.UnsafeGoID(), 5438220604201114580, func(notifier uintptr) { fn(UnsafeNewQSocketNotifier(unsafe.Pointer(notifier))) }) // slot "unregisterSocketNotifier"
}

func (this *QAbstractEventDispatcher) OnRegisterTimer(fn func(int, int64, TimerType, *QObject)) {
	registerVirt(this.UnsafeGoID(), 5345355196729094449, func(timerId int32, interval int64, timerType int32, object uintptr) { fn(int(timerId), interval, TimerType(timerType), UnsafeNewQObject(unsafe.Pointer(object))) }) // slot "registerTimer"
}

func (this *QAbstractEventDispatcher) OnUnregisterTimer(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), -85727209796773320, func(timerId int32) int32 { return b2i(fn(int(timerId))) }) // slot "unregisterTimer"
}

func (this *QAbstractEventDispatcher) OnUnregisterTimers(fn func(*QObject) bool) {
	registerVirt(this.UnsafeGoID(), 8418593569725967729, func(object uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)))) }) // slot "unregisterTimers"
}

func (this *QAbstractEventDispatcher) OnRegisteredTimers(fn func(*QObject) []*QAbstractEventDispatcher__TimerInfo) {
	registerVirt(this.UnsafeGoID(), 776599310866639255, func(object uintptr) (uintptr, int64) {
		__rs := fn(UnsafeNewQObject(unsafe.Pointer(object)))
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "registeredTimers"
}

func (this *QAbstractEventDispatcher) OnRemainingTime(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -7776574712936801734, func(timerId int32) int32 { return int32(fn(int(timerId))) }) // slot "remainingTime"
}

func (this *QAbstractEventDispatcher) OnWakeUp(fn func()) {
	registerVirt(this.UnsafeGoID(), -6649993720439713720, func() { fn() }) // slot "wakeUp"
}

func (this *QAbstractEventDispatcher) OnInterrupt(fn func()) {
	registerVirt(this.UnsafeGoID(), -3579312661142054576, func() { fn() }) // slot "interrupt"
}

func (this *QAbstractEventDispatcher) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractEventDispatcher) OnStartingUp(fn func()) {
	registerVirt(this.UnsafeGoID(), -4418351852319543908, func() { fn() }) // slot "startingUp"
}

func (this *QAbstractEventDispatcher) OnClosingDown(fn func()) {
	registerVirt(this.UnsafeGoID(), -8213485174827668208, func() { fn() }) // slot "closingDown"
}

func (this *QAbstractEventDispatcher) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractEventDispatcher) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractEventDispatcher) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractEventDispatcher) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractEventDispatcher) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractEventDispatcher) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractEventDispatcher) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractEventDispatcher) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractEventDispatcher_delete(this.h)
	}
	this.UnsafeRelease()
}
