package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QParallelAnimationGroup_chain uintptr
	addr_QParallelAnimationGroup_new uintptr
	addr_QParallelAnimationGroup_new2 uintptr
	addr_QParallelAnimationGroup_tr uintptr
	addr_QParallelAnimationGroup_connect uintptr
	addr_QParallelAnimationGroup_connect2 uintptr
	addr_QParallelAnimationGroup_sender uintptr
	addr_QParallelAnimationGroup_senderSignalIndex uintptr
	addr_QParallelAnimationGroup_receivers uintptr
	addr_QParallelAnimationGroup_isSignalConnected uintptr
	addr_QParallelAnimationGroup_tr2 uintptr
	addr_QParallelAnimationGroup_tr3 uintptr
	addr_QParallelAnimationGroup_connect3 uintptr
	addr_QParallelAnimationGroup_connect4 uintptr
	addr_QParallelAnimationGroup_metaObject_default uintptr
	addr_QParallelAnimationGroup_duration_default uintptr
	addr_QParallelAnimationGroup_event_default uintptr
	addr_QParallelAnimationGroup_updateCurrentTime_default uintptr
	addr_QParallelAnimationGroup_updateState_default uintptr
	addr_QParallelAnimationGroup_updateDirection_default uintptr
	addr_QParallelAnimationGroup_eventFilter_default uintptr
	addr_QParallelAnimationGroup_timerEvent_default uintptr
	addr_QParallelAnimationGroup_childEvent_default uintptr
	addr_QParallelAnimationGroup_customEvent_default uintptr
	addr_QParallelAnimationGroup_connectNotify_default uintptr
	addr_QParallelAnimationGroup_disconnectNotify_default uintptr
	pg_QParallelAnimationGroup_connect_finished func(uintptr, int64)
	pg_QParallelAnimationGroup_connect_stateChanged func(uintptr, int64)
	pg_QParallelAnimationGroup_connect_currentLoopChanged func(uintptr, int64)
	pg_QParallelAnimationGroup_connect_directionChanged func(uintptr, int64)
	pg_QParallelAnimationGroup_connect_destroyed func(uintptr, int64)
	pg_QParallelAnimationGroup_connect_destroyed2 func(uintptr, int64)
	pg_QParallelAnimationGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QParallelAnimationGroup__disconnect func(int64)
	pg_QParallelAnimationGroup__set_override func(uintptr, int32)
	pg_QParallelAnimationGroup_delete func(uintptr)
)

func init() {
	symaddr(&addr_QParallelAnimationGroup_chain, "QParallelAnimationGroup_chain")
	symaddr(&addr_QParallelAnimationGroup_new, "QParallelAnimationGroup_new")
	symaddr(&addr_QParallelAnimationGroup_new2, "QParallelAnimationGroup_new2")
	symaddr(&addr_QParallelAnimationGroup_tr, "QParallelAnimationGroup_tr")
	symaddr(&addr_QParallelAnimationGroup_connect, "QParallelAnimationGroup_connect")
	symaddr(&addr_QParallelAnimationGroup_connect2, "QParallelAnimationGroup_connect2")
	symaddr(&addr_QParallelAnimationGroup_sender, "QParallelAnimationGroup_sender")
	symaddr(&addr_QParallelAnimationGroup_senderSignalIndex, "QParallelAnimationGroup_senderSignalIndex")
	symaddr(&addr_QParallelAnimationGroup_receivers, "QParallelAnimationGroup_receivers")
	symaddr(&addr_QParallelAnimationGroup_isSignalConnected, "QParallelAnimationGroup_isSignalConnected")
	symaddr(&addr_QParallelAnimationGroup_tr2, "QParallelAnimationGroup_tr2")
	symaddr(&addr_QParallelAnimationGroup_tr3, "QParallelAnimationGroup_tr3")
	symaddr(&addr_QParallelAnimationGroup_connect3, "QParallelAnimationGroup_connect3")
	symaddr(&addr_QParallelAnimationGroup_connect4, "QParallelAnimationGroup_connect4")
	symaddr(&addr_QParallelAnimationGroup_metaObject_default, "QParallelAnimationGroup_metaObject_default")
	symaddr(&addr_QParallelAnimationGroup_duration_default, "QParallelAnimationGroup_duration_default")
	symaddr(&addr_QParallelAnimationGroup_event_default, "QParallelAnimationGroup_event_default")
	symaddr(&addr_QParallelAnimationGroup_updateCurrentTime_default, "QParallelAnimationGroup_updateCurrentTime_default")
	symaddr(&addr_QParallelAnimationGroup_updateState_default, "QParallelAnimationGroup_updateState_default")
	symaddr(&addr_QParallelAnimationGroup_updateDirection_default, "QParallelAnimationGroup_updateDirection_default")
	symaddr(&addr_QParallelAnimationGroup_eventFilter_default, "QParallelAnimationGroup_eventFilter_default")
	symaddr(&addr_QParallelAnimationGroup_timerEvent_default, "QParallelAnimationGroup_timerEvent_default")
	symaddr(&addr_QParallelAnimationGroup_childEvent_default, "QParallelAnimationGroup_childEvent_default")
	symaddr(&addr_QParallelAnimationGroup_customEvent_default, "QParallelAnimationGroup_customEvent_default")
	symaddr(&addr_QParallelAnimationGroup_connectNotify_default, "QParallelAnimationGroup_connectNotify_default")
	symaddr(&addr_QParallelAnimationGroup_disconnectNotify_default, "QParallelAnimationGroup_disconnectNotify_default")
	reg(&pg_QParallelAnimationGroup_connect_finished, "QParallelAnimationGroup_connect_finished")
	reg(&pg_QParallelAnimationGroup_connect_stateChanged, "QParallelAnimationGroup_connect_stateChanged")
	reg(&pg_QParallelAnimationGroup_connect_currentLoopChanged, "QParallelAnimationGroup_connect_currentLoopChanged")
	reg(&pg_QParallelAnimationGroup_connect_directionChanged, "QParallelAnimationGroup_connect_directionChanged")
	reg(&pg_QParallelAnimationGroup_connect_destroyed, "QParallelAnimationGroup_connect_destroyed")
	reg(&pg_QParallelAnimationGroup_connect_destroyed2, "QParallelAnimationGroup_connect_destroyed2")
	reg(&pg_QParallelAnimationGroup_connect_objectNameChanged, "QParallelAnimationGroup_connect_objectNameChanged")
	reg(&pg_QParallelAnimationGroup__disconnect, "QParallelAnimationGroup__disconnect")
	reg(&pg_QParallelAnimationGroup__set_override, "QParallelAnimationGroup__set_override")
	reg(&pg_QParallelAnimationGroup_delete, "QParallelAnimationGroup_delete")
}

type QParallelAnimationGroup struct {
	h uintptr
	*QAnimationGroup
}

func bindQParallelAnimationGroup(h uintptr) *QParallelAnimationGroup {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QParallelAnimationGroup_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QParallelAnimationGroup{h: h, QAnimationGroup: &QAnimationGroup{h: __c[0], QAbstractAnimation: &QAbstractAnimation{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQParallelAnimationGroup(h uintptr) *QParallelAnimationGroup { return bindQParallelAnimationGroup(h) }

func (this *QParallelAnimationGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QParallelAnimationGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQParallelAnimationGroup(h unsafe.Pointer) *QParallelAnimationGroup { return bindQParallelAnimationGroup(uintptr(h)) }
func UnsafeNewQParallelAnimationGroupOwned(h unsafe.Pointer) *QParallelAnimationGroup { return newownedQParallelAnimationGroup(uintptr(h)) }

func (this *QParallelAnimationGroup) UnsafeRelease() {
	this.h = 0
	if this.QAnimationGroup != nil {
		this.QAnimationGroup.UnsafeRelease()
	}
}

func NewQParallelAnimationGroup() *QParallelAnimationGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_new, uintptr(__goid))
	__o := newownedQParallelAnimationGroup(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QParallelAnimationGroup_virtIdx, QParallelAnimationGroup_enable)
	return __o
}

func NewQParallelAnimationGroup2(parent *QObject) *QParallelAnimationGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQParallelAnimationGroup(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QParallelAnimationGroup_virtIdx, QParallelAnimationGroup_enable)
	return __o
}

func QParallelAnimationGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QParallelAnimationGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QParallelAnimationGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QParallelAnimationGroup) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QParallelAnimationGroup) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QParallelAnimationGroup) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QParallelAnimationGroup) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QParallelAnimationGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QParallelAnimationGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QParallelAnimationGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QParallelAnimationGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QParallelAnimationGroup) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QParallelAnimationGroup) DurationDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_duration_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QParallelAnimationGroup) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QParallelAnimationGroup) UpdateCurrentTimeDefault(currentTime int) {
	purego.SyscallN(addr_QParallelAnimationGroup_updateCurrentTime_default, this.h, uintptr(currentTime))
	runtime.KeepAlive(this)
}

func (this *QParallelAnimationGroup) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QParallelAnimationGroup_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QParallelAnimationGroup) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QParallelAnimationGroup_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QParallelAnimationGroup) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QParallelAnimationGroup_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QParallelAnimationGroup) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QParallelAnimationGroup_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QParallelAnimationGroup) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QParallelAnimationGroup_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QParallelAnimationGroup) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QParallelAnimationGroup_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QParallelAnimationGroup) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QParallelAnimationGroup_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QParallelAnimationGroup) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QParallelAnimationGroup_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QParallelAnimationGroup) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QParallelAnimationGroup_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QParallelAnimationGroup_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QParallelAnimationGroup_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QParallelAnimationGroup_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QParallelAnimationGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QParallelAnimationGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QParallelAnimationGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QParallelAnimationGroup) Disconnect(handle int64) {
	pg_QParallelAnimationGroup__disconnect(handle)
	unregisterCallback(handle)
}

func QParallelAnimationGroup_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -8434675112194400115:
		return 1
	case 1237236456312824991:
		return 2
	case -1342418961870150968:
		return 3
	case 7692482089134926569:
		return 4
	case -3231888725407197643:
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

func QParallelAnimationGroup_enable(h uintptr, idx int32) { pg_QParallelAnimationGroup__set_override(h, idx) }

func (this *QParallelAnimationGroup) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QParallelAnimationGroup) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QParallelAnimationGroup) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QParallelAnimationGroup) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(currentTime int32) { fn(int(currentTime)) }) // slot "updateCurrentTime"
}

func (this *QParallelAnimationGroup) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QParallelAnimationGroup) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QParallelAnimationGroup) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QParallelAnimationGroup) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QParallelAnimationGroup) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QParallelAnimationGroup) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QParallelAnimationGroup) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QParallelAnimationGroup) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QParallelAnimationGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QParallelAnimationGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
