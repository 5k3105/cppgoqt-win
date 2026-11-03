package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSequentialAnimationGroup_chain uintptr
	addr_QSequentialAnimationGroup_new uintptr
	addr_QSequentialAnimationGroup_new2 uintptr
	addr_QSequentialAnimationGroup_tr uintptr
	addr_QSequentialAnimationGroup_addPause uintptr
	addr_QSequentialAnimationGroup_insertPause uintptr
	addr_QSequentialAnimationGroup_currentAnimation uintptr
	addr_QSequentialAnimationGroup_currentAnimationChanged uintptr
	addr_QSequentialAnimationGroup_connect uintptr
	addr_QSequentialAnimationGroup_connect2 uintptr
	addr_QSequentialAnimationGroup_sender uintptr
	addr_QSequentialAnimationGroup_senderSignalIndex uintptr
	addr_QSequentialAnimationGroup_receivers uintptr
	addr_QSequentialAnimationGroup_isSignalConnected uintptr
	addr_QSequentialAnimationGroup_tr2 uintptr
	addr_QSequentialAnimationGroup_tr3 uintptr
	addr_QSequentialAnimationGroup_connect3 uintptr
	addr_QSequentialAnimationGroup_connect4 uintptr
	addr_QSequentialAnimationGroup_metaObject_default uintptr
	addr_QSequentialAnimationGroup_duration_default uintptr
	addr_QSequentialAnimationGroup_event_default uintptr
	addr_QSequentialAnimationGroup_updateCurrentTime_default uintptr
	addr_QSequentialAnimationGroup_updateState_default uintptr
	addr_QSequentialAnimationGroup_updateDirection_default uintptr
	addr_QSequentialAnimationGroup_eventFilter_default uintptr
	addr_QSequentialAnimationGroup_timerEvent_default uintptr
	addr_QSequentialAnimationGroup_childEvent_default uintptr
	addr_QSequentialAnimationGroup_customEvent_default uintptr
	addr_QSequentialAnimationGroup_connectNotify_default uintptr
	addr_QSequentialAnimationGroup_disconnectNotify_default uintptr
	pg_QSequentialAnimationGroup_connect_currentAnimationChanged func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_finished func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_stateChanged func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_currentLoopChanged func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_directionChanged func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_destroyed func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_destroyed2 func(uintptr, int64)
	pg_QSequentialAnimationGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QSequentialAnimationGroup__disconnect func(int64)
	pg_QSequentialAnimationGroup__set_override func(uintptr, int32)
	pg_QSequentialAnimationGroup_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSequentialAnimationGroup_chain, "QSequentialAnimationGroup_chain")
	symaddr(&addr_QSequentialAnimationGroup_new, "QSequentialAnimationGroup_new")
	symaddr(&addr_QSequentialAnimationGroup_new2, "QSequentialAnimationGroup_new2")
	symaddr(&addr_QSequentialAnimationGroup_tr, "QSequentialAnimationGroup_tr")
	symaddr(&addr_QSequentialAnimationGroup_addPause, "QSequentialAnimationGroup_addPause")
	symaddr(&addr_QSequentialAnimationGroup_insertPause, "QSequentialAnimationGroup_insertPause")
	symaddr(&addr_QSequentialAnimationGroup_currentAnimation, "QSequentialAnimationGroup_currentAnimation")
	symaddr(&addr_QSequentialAnimationGroup_currentAnimationChanged, "QSequentialAnimationGroup_currentAnimationChanged")
	symaddr(&addr_QSequentialAnimationGroup_connect, "QSequentialAnimationGroup_connect")
	symaddr(&addr_QSequentialAnimationGroup_connect2, "QSequentialAnimationGroup_connect2")
	symaddr(&addr_QSequentialAnimationGroup_sender, "QSequentialAnimationGroup_sender")
	symaddr(&addr_QSequentialAnimationGroup_senderSignalIndex, "QSequentialAnimationGroup_senderSignalIndex")
	symaddr(&addr_QSequentialAnimationGroup_receivers, "QSequentialAnimationGroup_receivers")
	symaddr(&addr_QSequentialAnimationGroup_isSignalConnected, "QSequentialAnimationGroup_isSignalConnected")
	symaddr(&addr_QSequentialAnimationGroup_tr2, "QSequentialAnimationGroup_tr2")
	symaddr(&addr_QSequentialAnimationGroup_tr3, "QSequentialAnimationGroup_tr3")
	symaddr(&addr_QSequentialAnimationGroup_connect3, "QSequentialAnimationGroup_connect3")
	symaddr(&addr_QSequentialAnimationGroup_connect4, "QSequentialAnimationGroup_connect4")
	symaddr(&addr_QSequentialAnimationGroup_metaObject_default, "QSequentialAnimationGroup_metaObject_default")
	symaddr(&addr_QSequentialAnimationGroup_duration_default, "QSequentialAnimationGroup_duration_default")
	symaddr(&addr_QSequentialAnimationGroup_event_default, "QSequentialAnimationGroup_event_default")
	symaddr(&addr_QSequentialAnimationGroup_updateCurrentTime_default, "QSequentialAnimationGroup_updateCurrentTime_default")
	symaddr(&addr_QSequentialAnimationGroup_updateState_default, "QSequentialAnimationGroup_updateState_default")
	symaddr(&addr_QSequentialAnimationGroup_updateDirection_default, "QSequentialAnimationGroup_updateDirection_default")
	symaddr(&addr_QSequentialAnimationGroup_eventFilter_default, "QSequentialAnimationGroup_eventFilter_default")
	symaddr(&addr_QSequentialAnimationGroup_timerEvent_default, "QSequentialAnimationGroup_timerEvent_default")
	symaddr(&addr_QSequentialAnimationGroup_childEvent_default, "QSequentialAnimationGroup_childEvent_default")
	symaddr(&addr_QSequentialAnimationGroup_customEvent_default, "QSequentialAnimationGroup_customEvent_default")
	symaddr(&addr_QSequentialAnimationGroup_connectNotify_default, "QSequentialAnimationGroup_connectNotify_default")
	symaddr(&addr_QSequentialAnimationGroup_disconnectNotify_default, "QSequentialAnimationGroup_disconnectNotify_default")
	reg(&pg_QSequentialAnimationGroup_connect_currentAnimationChanged, "QSequentialAnimationGroup_connect_currentAnimationChanged")
	reg(&pg_QSequentialAnimationGroup_connect_finished, "QSequentialAnimationGroup_connect_finished")
	reg(&pg_QSequentialAnimationGroup_connect_stateChanged, "QSequentialAnimationGroup_connect_stateChanged")
	reg(&pg_QSequentialAnimationGroup_connect_currentLoopChanged, "QSequentialAnimationGroup_connect_currentLoopChanged")
	reg(&pg_QSequentialAnimationGroup_connect_directionChanged, "QSequentialAnimationGroup_connect_directionChanged")
	reg(&pg_QSequentialAnimationGroup_connect_destroyed, "QSequentialAnimationGroup_connect_destroyed")
	reg(&pg_QSequentialAnimationGroup_connect_destroyed2, "QSequentialAnimationGroup_connect_destroyed2")
	reg(&pg_QSequentialAnimationGroup_connect_objectNameChanged, "QSequentialAnimationGroup_connect_objectNameChanged")
	reg(&pg_QSequentialAnimationGroup__disconnect, "QSequentialAnimationGroup__disconnect")
	reg(&pg_QSequentialAnimationGroup__set_override, "QSequentialAnimationGroup__set_override")
	reg(&pg_QSequentialAnimationGroup_delete, "QSequentialAnimationGroup_delete")
}

type QSequentialAnimationGroup struct {
	h uintptr
	*QAnimationGroup
}

func bindQSequentialAnimationGroup(h uintptr) *QSequentialAnimationGroup {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSequentialAnimationGroup_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSequentialAnimationGroup{h: h, QAnimationGroup: &QAnimationGroup{h: __c[0], QAbstractAnimation: &QAbstractAnimation{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQSequentialAnimationGroup(h uintptr) *QSequentialAnimationGroup { return bindQSequentialAnimationGroup(h) }

func (this *QSequentialAnimationGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSequentialAnimationGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSequentialAnimationGroup(h unsafe.Pointer) *QSequentialAnimationGroup { return bindQSequentialAnimationGroup(uintptr(h)) }
func UnsafeNewQSequentialAnimationGroupOwned(h unsafe.Pointer) *QSequentialAnimationGroup { return newownedQSequentialAnimationGroup(uintptr(h)) }

func (this *QSequentialAnimationGroup) UnsafeRelease() {
	this.h = 0
	if this.QAnimationGroup != nil {
		this.QAnimationGroup.UnsafeRelease()
	}
}

func NewQSequentialAnimationGroup() *QSequentialAnimationGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_new, uintptr(__goid))
	__o := newownedQSequentialAnimationGroup(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSequentialAnimationGroup_virtIdx, QSequentialAnimationGroup_enable)
	return __o
}

func NewQSequentialAnimationGroup2(parent *QObject) *QSequentialAnimationGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSequentialAnimationGroup(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSequentialAnimationGroup_virtIdx, QSequentialAnimationGroup_enable)
	return __o
}

func QSequentialAnimationGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_addPause, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return bindQPauseAnimation(__r)
}

func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_insertPause, this.h, uintptr(index), uintptr(msecs))
	runtime.KeepAlive(this)
	return bindQPauseAnimation(__r)
}

func (this *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_currentAnimation, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractAnimation(__r)
}

func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current *QAbstractAnimation) {
	purego.SyscallN(addr_QSequentialAnimationGroup_currentAnimationChanged, this.h, uintptr(current.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
}

func QSequentialAnimationGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSequentialAnimationGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSequentialAnimationGroup) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSequentialAnimationGroup) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSequentialAnimationGroup) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSequentialAnimationGroup) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSequentialAnimationGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSequentialAnimationGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSequentialAnimationGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSequentialAnimationGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSequentialAnimationGroup) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSequentialAnimationGroup) DurationDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_duration_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSequentialAnimationGroup) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSequentialAnimationGroup) UpdateCurrentTimeDefault(arg1 int) {
	purego.SyscallN(addr_QSequentialAnimationGroup_updateCurrentTime_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSequentialAnimationGroup) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QSequentialAnimationGroup_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QSequentialAnimationGroup) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QSequentialAnimationGroup_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QSequentialAnimationGroup) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialAnimationGroup_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSequentialAnimationGroup) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSequentialAnimationGroup_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSequentialAnimationGroup) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSequentialAnimationGroup_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSequentialAnimationGroup) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSequentialAnimationGroup_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSequentialAnimationGroup) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSequentialAnimationGroup_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSequentialAnimationGroup) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSequentialAnimationGroup_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSequentialAnimationGroup) OnCurrentAnimationChanged(fn func(*QAbstractAnimation)) int64 {
	id := registerCallback(func(current uintptr) { fn(UnsafeNewQAbstractAnimation(unsafe.Pointer(current))) })
	pg_QSequentialAnimationGroup_connect_currentAnimationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSequentialAnimationGroup_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QSequentialAnimationGroup_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QSequentialAnimationGroup_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QSequentialAnimationGroup_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSequentialAnimationGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSequentialAnimationGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSequentialAnimationGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSequentialAnimationGroup) Disconnect(handle int64) {
	pg_QSequentialAnimationGroup__disconnect(handle)
	unregisterCallback(handle)
}

func QSequentialAnimationGroup_virtIdx(slot int64) int32 {
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

func QSequentialAnimationGroup_enable(h uintptr, idx int32) { pg_QSequentialAnimationGroup__set_override(h, idx) }

func (this *QSequentialAnimationGroup) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSequentialAnimationGroup) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QSequentialAnimationGroup) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSequentialAnimationGroup) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(arg1 int32) { fn(int(arg1)) }) // slot "updateCurrentTime"
}

func (this *QSequentialAnimationGroup) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QSequentialAnimationGroup) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QSequentialAnimationGroup) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSequentialAnimationGroup) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSequentialAnimationGroup) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSequentialAnimationGroup) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSequentialAnimationGroup) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSequentialAnimationGroup) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSequentialAnimationGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSequentialAnimationGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
