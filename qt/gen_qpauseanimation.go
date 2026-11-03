package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPauseAnimation_chain uintptr
	addr_QPauseAnimation_new uintptr
	addr_QPauseAnimation_new2 uintptr
	addr_QPauseAnimation_new3 uintptr
	addr_QPauseAnimation_new4 uintptr
	addr_QPauseAnimation_tr uintptr
	addr_QPauseAnimation_setDuration uintptr
	addr_QPauseAnimation_connect uintptr
	addr_QPauseAnimation_connect2 uintptr
	addr_QPauseAnimation_sender uintptr
	addr_QPauseAnimation_senderSignalIndex uintptr
	addr_QPauseAnimation_receivers uintptr
	addr_QPauseAnimation_isSignalConnected uintptr
	addr_QPauseAnimation_tr2 uintptr
	addr_QPauseAnimation_tr3 uintptr
	addr_QPauseAnimation_connect3 uintptr
	addr_QPauseAnimation_connect4 uintptr
	addr_QPauseAnimation_metaObject_default uintptr
	addr_QPauseAnimation_duration_default uintptr
	addr_QPauseAnimation_event_default uintptr
	addr_QPauseAnimation_updateCurrentTime_default uintptr
	addr_QPauseAnimation_updateState_default uintptr
	addr_QPauseAnimation_updateDirection_default uintptr
	addr_QPauseAnimation_eventFilter_default uintptr
	addr_QPauseAnimation_timerEvent_default uintptr
	addr_QPauseAnimation_childEvent_default uintptr
	addr_QPauseAnimation_customEvent_default uintptr
	addr_QPauseAnimation_connectNotify_default uintptr
	addr_QPauseAnimation_disconnectNotify_default uintptr
	pg_QPauseAnimation_connect_finished func(uintptr, int64)
	pg_QPauseAnimation_connect_stateChanged func(uintptr, int64)
	pg_QPauseAnimation_connect_currentLoopChanged func(uintptr, int64)
	pg_QPauseAnimation_connect_directionChanged func(uintptr, int64)
	pg_QPauseAnimation_connect_destroyed func(uintptr, int64)
	pg_QPauseAnimation_connect_destroyed2 func(uintptr, int64)
	pg_QPauseAnimation_connect_objectNameChanged func(uintptr, int64)
	pg_QPauseAnimation__disconnect func(int64)
	pg_QPauseAnimation__set_override func(uintptr, int32)
	pg_QPauseAnimation_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPauseAnimation_chain, "QPauseAnimation_chain")
	symaddr(&addr_QPauseAnimation_new, "QPauseAnimation_new")
	symaddr(&addr_QPauseAnimation_new2, "QPauseAnimation_new2")
	symaddr(&addr_QPauseAnimation_new3, "QPauseAnimation_new3")
	symaddr(&addr_QPauseAnimation_new4, "QPauseAnimation_new4")
	symaddr(&addr_QPauseAnimation_tr, "QPauseAnimation_tr")
	symaddr(&addr_QPauseAnimation_setDuration, "QPauseAnimation_setDuration")
	symaddr(&addr_QPauseAnimation_connect, "QPauseAnimation_connect")
	symaddr(&addr_QPauseAnimation_connect2, "QPauseAnimation_connect2")
	symaddr(&addr_QPauseAnimation_sender, "QPauseAnimation_sender")
	symaddr(&addr_QPauseAnimation_senderSignalIndex, "QPauseAnimation_senderSignalIndex")
	symaddr(&addr_QPauseAnimation_receivers, "QPauseAnimation_receivers")
	symaddr(&addr_QPauseAnimation_isSignalConnected, "QPauseAnimation_isSignalConnected")
	symaddr(&addr_QPauseAnimation_tr2, "QPauseAnimation_tr2")
	symaddr(&addr_QPauseAnimation_tr3, "QPauseAnimation_tr3")
	symaddr(&addr_QPauseAnimation_connect3, "QPauseAnimation_connect3")
	symaddr(&addr_QPauseAnimation_connect4, "QPauseAnimation_connect4")
	symaddr(&addr_QPauseAnimation_metaObject_default, "QPauseAnimation_metaObject_default")
	symaddr(&addr_QPauseAnimation_duration_default, "QPauseAnimation_duration_default")
	symaddr(&addr_QPauseAnimation_event_default, "QPauseAnimation_event_default")
	symaddr(&addr_QPauseAnimation_updateCurrentTime_default, "QPauseAnimation_updateCurrentTime_default")
	symaddr(&addr_QPauseAnimation_updateState_default, "QPauseAnimation_updateState_default")
	symaddr(&addr_QPauseAnimation_updateDirection_default, "QPauseAnimation_updateDirection_default")
	symaddr(&addr_QPauseAnimation_eventFilter_default, "QPauseAnimation_eventFilter_default")
	symaddr(&addr_QPauseAnimation_timerEvent_default, "QPauseAnimation_timerEvent_default")
	symaddr(&addr_QPauseAnimation_childEvent_default, "QPauseAnimation_childEvent_default")
	symaddr(&addr_QPauseAnimation_customEvent_default, "QPauseAnimation_customEvent_default")
	symaddr(&addr_QPauseAnimation_connectNotify_default, "QPauseAnimation_connectNotify_default")
	symaddr(&addr_QPauseAnimation_disconnectNotify_default, "QPauseAnimation_disconnectNotify_default")
	reg(&pg_QPauseAnimation_connect_finished, "QPauseAnimation_connect_finished")
	reg(&pg_QPauseAnimation_connect_stateChanged, "QPauseAnimation_connect_stateChanged")
	reg(&pg_QPauseAnimation_connect_currentLoopChanged, "QPauseAnimation_connect_currentLoopChanged")
	reg(&pg_QPauseAnimation_connect_directionChanged, "QPauseAnimation_connect_directionChanged")
	reg(&pg_QPauseAnimation_connect_destroyed, "QPauseAnimation_connect_destroyed")
	reg(&pg_QPauseAnimation_connect_destroyed2, "QPauseAnimation_connect_destroyed2")
	reg(&pg_QPauseAnimation_connect_objectNameChanged, "QPauseAnimation_connect_objectNameChanged")
	reg(&pg_QPauseAnimation__disconnect, "QPauseAnimation__disconnect")
	reg(&pg_QPauseAnimation__set_override, "QPauseAnimation__set_override")
	reg(&pg_QPauseAnimation_delete, "QPauseAnimation_delete")
}

type QPauseAnimation struct {
	h uintptr
	*QAbstractAnimation
}

func bindQPauseAnimation(h uintptr) *QPauseAnimation {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPauseAnimation_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPauseAnimation{h: h, QAbstractAnimation: &QAbstractAnimation{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQPauseAnimation(h uintptr) *QPauseAnimation { return bindQPauseAnimation(h) }

func (this *QPauseAnimation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPauseAnimation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPauseAnimation(h unsafe.Pointer) *QPauseAnimation { return bindQPauseAnimation(uintptr(h)) }
func UnsafeNewQPauseAnimationOwned(h unsafe.Pointer) *QPauseAnimation { return newownedQPauseAnimation(uintptr(h)) }

func (this *QPauseAnimation) UnsafeRelease() {
	this.h = 0
	if this.QAbstractAnimation != nil {
		this.QAbstractAnimation.UnsafeRelease()
	}
}

func NewQPauseAnimation() *QPauseAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_new, uintptr(__goid))
	__o := newownedQPauseAnimation(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPauseAnimation_virtIdx, QPauseAnimation_enable)
	return __o
}

func NewQPauseAnimation2(msecs int) *QPauseAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_new2, uintptr(__goid), uintptr(msecs))
	__o := newownedQPauseAnimation(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPauseAnimation_virtIdx, QPauseAnimation_enable)
	return __o
}

func NewQPauseAnimation3(parent *QObject) *QPauseAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPauseAnimation(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPauseAnimation_virtIdx, QPauseAnimation_enable)
	return __o
}

func NewQPauseAnimation4(msecs int, parent *QObject) *QPauseAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_new4, uintptr(__goid), uintptr(msecs), uintptr(parent.UnsafePointer()))
	__o := newownedQPauseAnimation(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPauseAnimation_virtIdx, QPauseAnimation_enable)
	return __o
}

func QPauseAnimation_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPauseAnimation) SetDuration(msecs int) {
	purego.SyscallN(addr_QPauseAnimation_setDuration, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
}

func QPauseAnimation_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPauseAnimation_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPauseAnimation) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPauseAnimation) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPauseAnimation_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPauseAnimation) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPauseAnimation_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPauseAnimation) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPauseAnimation_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPauseAnimation_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPauseAnimation_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPauseAnimation_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPauseAnimation_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPauseAnimation) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPauseAnimation_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPauseAnimation) DurationDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QPauseAnimation_duration_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPauseAnimation) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPauseAnimation_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QPauseAnimation) UpdateCurrentTimeDefault(arg1 int) {
	purego.SyscallN(addr_QPauseAnimation_updateCurrentTime_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QPauseAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QPauseAnimation_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QPauseAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QPauseAnimation_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QPauseAnimation) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPauseAnimation_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPauseAnimation) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QPauseAnimation_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPauseAnimation) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPauseAnimation_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPauseAnimation) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPauseAnimation_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPauseAnimation) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPauseAnimation_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPauseAnimation) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPauseAnimation_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPauseAnimation) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPauseAnimation_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QPauseAnimation_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QPauseAnimation_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QPauseAnimation_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPauseAnimation_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPauseAnimation_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPauseAnimation_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPauseAnimation) Disconnect(handle int64) {
	pg_QPauseAnimation__disconnect(handle)
	unregisterCallback(handle)
}

func QPauseAnimation_virtIdx(slot int64) int32 {
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

func QPauseAnimation_enable(h uintptr, idx int32) { pg_QPauseAnimation__set_override(h, idx) }

func (this *QPauseAnimation) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPauseAnimation) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QPauseAnimation) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QPauseAnimation) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(arg1 int32) { fn(int(arg1)) }) // slot "updateCurrentTime"
}

func (this *QPauseAnimation) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QPauseAnimation) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QPauseAnimation) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPauseAnimation) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QPauseAnimation) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPauseAnimation) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPauseAnimation) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPauseAnimation) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPauseAnimation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPauseAnimation_delete(this.h)
	}
	this.UnsafeRelease()
}
