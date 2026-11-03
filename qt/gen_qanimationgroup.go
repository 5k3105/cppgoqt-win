package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAnimationGroup_chain uintptr
	addr_QAnimationGroup_tr uintptr
	addr_QAnimationGroup_animationAt uintptr
	addr_QAnimationGroup_animationCount uintptr
	addr_QAnimationGroup_indexOfAnimation uintptr
	addr_QAnimationGroup_addAnimation uintptr
	addr_QAnimationGroup_insertAnimation uintptr
	addr_QAnimationGroup_removeAnimation uintptr
	addr_QAnimationGroup_takeAnimation uintptr
	addr_QAnimationGroup_clear uintptr
	addr_QAnimationGroup_connect uintptr
	addr_QAnimationGroup_connect2 uintptr
	addr_QAnimationGroup_sender uintptr
	addr_QAnimationGroup_senderSignalIndex uintptr
	addr_QAnimationGroup_receivers uintptr
	addr_QAnimationGroup_isSignalConnected uintptr
	addr_QAnimationGroup_tr2 uintptr
	addr_QAnimationGroup_tr3 uintptr
	addr_QAnimationGroup_connect3 uintptr
	addr_QAnimationGroup_connect4 uintptr
	addr_QAnimationGroup_metaObject_default uintptr
	addr_QAnimationGroup_event_default uintptr
	addr_QAnimationGroup_updateState_default uintptr
	addr_QAnimationGroup_updateDirection_default uintptr
	addr_QAnimationGroup_eventFilter_default uintptr
	addr_QAnimationGroup_timerEvent_default uintptr
	addr_QAnimationGroup_childEvent_default uintptr
	addr_QAnimationGroup_customEvent_default uintptr
	addr_QAnimationGroup_connectNotify_default uintptr
	addr_QAnimationGroup_disconnectNotify_default uintptr
	pg_QAnimationGroup_connect_finished func(uintptr, int64)
	pg_QAnimationGroup_connect_stateChanged func(uintptr, int64)
	pg_QAnimationGroup_connect_currentLoopChanged func(uintptr, int64)
	pg_QAnimationGroup_connect_directionChanged func(uintptr, int64)
	pg_QAnimationGroup_connect_destroyed func(uintptr, int64)
	pg_QAnimationGroup_connect_destroyed2 func(uintptr, int64)
	pg_QAnimationGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QAnimationGroup__disconnect func(int64)
	pg_QAnimationGroup__set_override func(uintptr, int32)
	pg_QAnimationGroup__Go_new func(int64, uintptr) uintptr
	pg_QAnimationGroup_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAnimationGroup_chain, "QAnimationGroup_chain")
	symaddr(&addr_QAnimationGroup_tr, "QAnimationGroup_tr")
	symaddr(&addr_QAnimationGroup_animationAt, "QAnimationGroup_animationAt")
	symaddr(&addr_QAnimationGroup_animationCount, "QAnimationGroup_animationCount")
	symaddr(&addr_QAnimationGroup_indexOfAnimation, "QAnimationGroup_indexOfAnimation")
	symaddr(&addr_QAnimationGroup_addAnimation, "QAnimationGroup_addAnimation")
	symaddr(&addr_QAnimationGroup_insertAnimation, "QAnimationGroup_insertAnimation")
	symaddr(&addr_QAnimationGroup_removeAnimation, "QAnimationGroup_removeAnimation")
	symaddr(&addr_QAnimationGroup_takeAnimation, "QAnimationGroup_takeAnimation")
	symaddr(&addr_QAnimationGroup_clear, "QAnimationGroup_clear")
	symaddr(&addr_QAnimationGroup_connect, "QAnimationGroup_connect")
	symaddr(&addr_QAnimationGroup_connect2, "QAnimationGroup_connect2")
	symaddr(&addr_QAnimationGroup_sender, "QAnimationGroup_sender")
	symaddr(&addr_QAnimationGroup_senderSignalIndex, "QAnimationGroup_senderSignalIndex")
	symaddr(&addr_QAnimationGroup_receivers, "QAnimationGroup_receivers")
	symaddr(&addr_QAnimationGroup_isSignalConnected, "QAnimationGroup_isSignalConnected")
	symaddr(&addr_QAnimationGroup_tr2, "QAnimationGroup_tr2")
	symaddr(&addr_QAnimationGroup_tr3, "QAnimationGroup_tr3")
	symaddr(&addr_QAnimationGroup_connect3, "QAnimationGroup_connect3")
	symaddr(&addr_QAnimationGroup_connect4, "QAnimationGroup_connect4")
	symaddr(&addr_QAnimationGroup_metaObject_default, "QAnimationGroup_metaObject_default")
	symaddr(&addr_QAnimationGroup_event_default, "QAnimationGroup_event_default")
	symaddr(&addr_QAnimationGroup_updateState_default, "QAnimationGroup_updateState_default")
	symaddr(&addr_QAnimationGroup_updateDirection_default, "QAnimationGroup_updateDirection_default")
	symaddr(&addr_QAnimationGroup_eventFilter_default, "QAnimationGroup_eventFilter_default")
	symaddr(&addr_QAnimationGroup_timerEvent_default, "QAnimationGroup_timerEvent_default")
	symaddr(&addr_QAnimationGroup_childEvent_default, "QAnimationGroup_childEvent_default")
	symaddr(&addr_QAnimationGroup_customEvent_default, "QAnimationGroup_customEvent_default")
	symaddr(&addr_QAnimationGroup_connectNotify_default, "QAnimationGroup_connectNotify_default")
	symaddr(&addr_QAnimationGroup_disconnectNotify_default, "QAnimationGroup_disconnectNotify_default")
	reg(&pg_QAnimationGroup_connect_finished, "QAnimationGroup_connect_finished")
	reg(&pg_QAnimationGroup_connect_stateChanged, "QAnimationGroup_connect_stateChanged")
	reg(&pg_QAnimationGroup_connect_currentLoopChanged, "QAnimationGroup_connect_currentLoopChanged")
	reg(&pg_QAnimationGroup_connect_directionChanged, "QAnimationGroup_connect_directionChanged")
	reg(&pg_QAnimationGroup_connect_destroyed, "QAnimationGroup_connect_destroyed")
	reg(&pg_QAnimationGroup_connect_destroyed2, "QAnimationGroup_connect_destroyed2")
	reg(&pg_QAnimationGroup_connect_objectNameChanged, "QAnimationGroup_connect_objectNameChanged")
	reg(&pg_QAnimationGroup__disconnect, "QAnimationGroup__disconnect")
	reg(&pg_QAnimationGroup__set_override, "QAnimationGroup__set_override")
	reg(&pg_QAnimationGroup__Go_new, "QAnimationGroup__Go_new")
	reg(&pg_QAnimationGroup_delete, "QAnimationGroup_delete")
}

type QAnimationGroup struct {
	h uintptr
	*QAbstractAnimation
}

func bindQAnimationGroup(h uintptr) *QAnimationGroup {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAnimationGroup_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAnimationGroup{h: h, QAbstractAnimation: &QAbstractAnimation{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAnimationGroup(h uintptr) *QAnimationGroup { return bindQAnimationGroup(h) }

func (this *QAnimationGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAnimationGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAnimationGroup(h unsafe.Pointer) *QAnimationGroup { return bindQAnimationGroup(uintptr(h)) }
func UnsafeNewQAnimationGroupOwned(h unsafe.Pointer) *QAnimationGroup { return newownedQAnimationGroup(uintptr(h)) }

func (this *QAnimationGroup) UnsafeRelease() {
	this.h = 0
	if this.QAbstractAnimation != nil {
		this.QAbstractAnimation.UnsafeRelease()
	}
}

func QAnimationGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAnimationGroup) AnimationAt(index int) *QAbstractAnimation {
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_animationAt, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQAbstractAnimation(__r)
}

func (this *QAnimationGroup) AnimationCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_animationCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAnimationGroup) IndexOfAnimation(animation *QAbstractAnimation) int {
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_indexOfAnimation, this.h, uintptr(animation.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(animation)
	return int(__r)
}

func (this *QAnimationGroup) AddAnimation(animation *QAbstractAnimation) {
	purego.SyscallN(addr_QAnimationGroup_addAnimation, this.h, uintptr(animation.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(animation)
}

func (this *QAnimationGroup) InsertAnimation(index int, animation *QAbstractAnimation) {
	purego.SyscallN(addr_QAnimationGroup_insertAnimation, this.h, uintptr(index), uintptr(animation.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(animation)
}

func (this *QAnimationGroup) RemoveAnimation(animation *QAbstractAnimation) {
	purego.SyscallN(addr_QAnimationGroup_removeAnimation, this.h, uintptr(animation.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(animation)
}

func (this *QAnimationGroup) TakeAnimation(index int) *QAbstractAnimation {
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_takeAnimation, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQAbstractAnimation(__r)
}

func (this *QAnimationGroup) Clear() {
	purego.SyscallN(addr_QAnimationGroup_clear, this.h)
	runtime.KeepAlive(this)
}

func QAnimationGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAnimationGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAnimationGroup) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAnimationGroup) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAnimationGroup) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAnimationGroup) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAnimationGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAnimationGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAnimationGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAnimationGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAnimationGroup) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAnimationGroup_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAnimationGroup) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAnimationGroup) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QAnimationGroup_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QAnimationGroup) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QAnimationGroup_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QAnimationGroup) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAnimationGroup_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAnimationGroup) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAnimationGroup_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAnimationGroup) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAnimationGroup_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAnimationGroup) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAnimationGroup_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAnimationGroup) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAnimationGroup_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAnimationGroup) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAnimationGroup_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAnimationGroup) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAnimationGroup_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QAnimationGroup_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QAnimationGroup_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QAnimationGroup_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAnimationGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAnimationGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAnimationGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAnimationGroup) Disconnect(handle int64) {
	pg_QAnimationGroup__disconnect(handle)
	unregisterCallback(handle)
}

func QAnimationGroup_virtIdx(slot int64) int32 {
	switch slot {
	case -8434675112194400115:
		return 0
	case -1342418961870150968:
		return 1
	case 1607084412591819503:
		return 2
	case 1237236456312824991:
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

func QAnimationGroup_enable(h uintptr, idx int32) { pg_QAnimationGroup__set_override(h, idx) }

func NewQAnimationGroup() *QAnimationGroup {
	id := nextVirtID()
	o := newownedQAnimationGroup(pg_QAnimationGroup__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAnimationGroup_virtIdx, QAnimationGroup_enable)
	return o
}

func NewQAnimationGroup2(parent *QObject) *QAnimationGroup {
	id := nextVirtID()
	o := newownedQAnimationGroup(pg_QAnimationGroup__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAnimationGroup_virtIdx, QAnimationGroup_enable)
	return o
}

func (this *QAnimationGroup) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QAnimationGroup) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(currentTime int32) { fn(int(currentTime)) }) // slot "updateCurrentTime"
}

func (this *QAnimationGroup) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAnimationGroup) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAnimationGroup) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QAnimationGroup) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QAnimationGroup) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAnimationGroup) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAnimationGroup) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAnimationGroup) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAnimationGroup) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAnimationGroup) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAnimationGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAnimationGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
