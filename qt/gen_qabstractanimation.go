package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAbstractAnimation_base func(uintptr) uintptr
	addr_QAbstractAnimation_tr uintptr
	addr_QAbstractAnimation_state uintptr
	addr_QAbstractAnimation_group uintptr
	addr_QAbstractAnimation_direction uintptr
	addr_QAbstractAnimation_setDirection uintptr
	addr_QAbstractAnimation_currentTime uintptr
	addr_QAbstractAnimation_currentLoopTime uintptr
	addr_QAbstractAnimation_loopCount uintptr
	addr_QAbstractAnimation_setLoopCount uintptr
	addr_QAbstractAnimation_currentLoop uintptr
	addr_QAbstractAnimation_duration uintptr
	addr_QAbstractAnimation_totalDuration uintptr
	addr_QAbstractAnimation_finished uintptr
	addr_QAbstractAnimation_stateChanged uintptr
	addr_QAbstractAnimation_currentLoopChanged uintptr
	addr_QAbstractAnimation_directionChanged uintptr
	addr_QAbstractAnimation_start uintptr
	addr_QAbstractAnimation_pause uintptr
	addr_QAbstractAnimation_resume uintptr
	addr_QAbstractAnimation_setPaused uintptr
	addr_QAbstractAnimation_stop uintptr
	addr_QAbstractAnimation_setCurrentTime uintptr
	addr_QAbstractAnimation_connect uintptr
	addr_QAbstractAnimation_connect2 uintptr
	addr_QAbstractAnimation_sender uintptr
	addr_QAbstractAnimation_senderSignalIndex uintptr
	addr_QAbstractAnimation_receivers uintptr
	addr_QAbstractAnimation_isSignalConnected uintptr
	addr_QAbstractAnimation_tr2 uintptr
	addr_QAbstractAnimation_tr3 uintptr
	addr_QAbstractAnimation_start2 uintptr
	addr_QAbstractAnimation_connect3 uintptr
	addr_QAbstractAnimation_connect4 uintptr
	addr_QAbstractAnimation_metaObject_default uintptr
	addr_QAbstractAnimation_event_default uintptr
	addr_QAbstractAnimation_updateState_default uintptr
	addr_QAbstractAnimation_updateDirection_default uintptr
	addr_QAbstractAnimation_eventFilter_default uintptr
	addr_QAbstractAnimation_timerEvent_default uintptr
	addr_QAbstractAnimation_childEvent_default uintptr
	addr_QAbstractAnimation_customEvent_default uintptr
	addr_QAbstractAnimation_connectNotify_default uintptr
	addr_QAbstractAnimation_disconnectNotify_default uintptr
	pg_QAbstractAnimation_connect_finished func(uintptr, int64)
	pg_QAbstractAnimation_connect_stateChanged func(uintptr, int64)
	pg_QAbstractAnimation_connect_currentLoopChanged func(uintptr, int64)
	pg_QAbstractAnimation_connect_directionChanged func(uintptr, int64)
	pg_QAbstractAnimation_connect_destroyed func(uintptr, int64)
	pg_QAbstractAnimation_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractAnimation_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractAnimation__disconnect func(int64)
	pg_QAbstractAnimation__set_override func(uintptr, int32)
	pg_QAbstractAnimation__Go_new func(int64, uintptr) uintptr
	pg_QAbstractAnimation_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractAnimation_base, "QAbstractAnimation_base")
	symaddr(&addr_QAbstractAnimation_tr, "QAbstractAnimation_tr")
	symaddr(&addr_QAbstractAnimation_state, "QAbstractAnimation_state")
	symaddr(&addr_QAbstractAnimation_group, "QAbstractAnimation_group")
	symaddr(&addr_QAbstractAnimation_direction, "QAbstractAnimation_direction")
	symaddr(&addr_QAbstractAnimation_setDirection, "QAbstractAnimation_setDirection")
	symaddr(&addr_QAbstractAnimation_currentTime, "QAbstractAnimation_currentTime")
	symaddr(&addr_QAbstractAnimation_currentLoopTime, "QAbstractAnimation_currentLoopTime")
	symaddr(&addr_QAbstractAnimation_loopCount, "QAbstractAnimation_loopCount")
	symaddr(&addr_QAbstractAnimation_setLoopCount, "QAbstractAnimation_setLoopCount")
	symaddr(&addr_QAbstractAnimation_currentLoop, "QAbstractAnimation_currentLoop")
	symaddr(&addr_QAbstractAnimation_duration, "QAbstractAnimation_duration")
	symaddr(&addr_QAbstractAnimation_totalDuration, "QAbstractAnimation_totalDuration")
	symaddr(&addr_QAbstractAnimation_finished, "QAbstractAnimation_finished")
	symaddr(&addr_QAbstractAnimation_stateChanged, "QAbstractAnimation_stateChanged")
	symaddr(&addr_QAbstractAnimation_currentLoopChanged, "QAbstractAnimation_currentLoopChanged")
	symaddr(&addr_QAbstractAnimation_directionChanged, "QAbstractAnimation_directionChanged")
	symaddr(&addr_QAbstractAnimation_start, "QAbstractAnimation_start")
	symaddr(&addr_QAbstractAnimation_pause, "QAbstractAnimation_pause")
	symaddr(&addr_QAbstractAnimation_resume, "QAbstractAnimation_resume")
	symaddr(&addr_QAbstractAnimation_setPaused, "QAbstractAnimation_setPaused")
	symaddr(&addr_QAbstractAnimation_stop, "QAbstractAnimation_stop")
	symaddr(&addr_QAbstractAnimation_setCurrentTime, "QAbstractAnimation_setCurrentTime")
	symaddr(&addr_QAbstractAnimation_connect, "QAbstractAnimation_connect")
	symaddr(&addr_QAbstractAnimation_connect2, "QAbstractAnimation_connect2")
	symaddr(&addr_QAbstractAnimation_sender, "QAbstractAnimation_sender")
	symaddr(&addr_QAbstractAnimation_senderSignalIndex, "QAbstractAnimation_senderSignalIndex")
	symaddr(&addr_QAbstractAnimation_receivers, "QAbstractAnimation_receivers")
	symaddr(&addr_QAbstractAnimation_isSignalConnected, "QAbstractAnimation_isSignalConnected")
	symaddr(&addr_QAbstractAnimation_tr2, "QAbstractAnimation_tr2")
	symaddr(&addr_QAbstractAnimation_tr3, "QAbstractAnimation_tr3")
	symaddr(&addr_QAbstractAnimation_start2, "QAbstractAnimation_start2")
	symaddr(&addr_QAbstractAnimation_connect3, "QAbstractAnimation_connect3")
	symaddr(&addr_QAbstractAnimation_connect4, "QAbstractAnimation_connect4")
	symaddr(&addr_QAbstractAnimation_metaObject_default, "QAbstractAnimation_metaObject_default")
	symaddr(&addr_QAbstractAnimation_event_default, "QAbstractAnimation_event_default")
	symaddr(&addr_QAbstractAnimation_updateState_default, "QAbstractAnimation_updateState_default")
	symaddr(&addr_QAbstractAnimation_updateDirection_default, "QAbstractAnimation_updateDirection_default")
	symaddr(&addr_QAbstractAnimation_eventFilter_default, "QAbstractAnimation_eventFilter_default")
	symaddr(&addr_QAbstractAnimation_timerEvent_default, "QAbstractAnimation_timerEvent_default")
	symaddr(&addr_QAbstractAnimation_childEvent_default, "QAbstractAnimation_childEvent_default")
	symaddr(&addr_QAbstractAnimation_customEvent_default, "QAbstractAnimation_customEvent_default")
	symaddr(&addr_QAbstractAnimation_connectNotify_default, "QAbstractAnimation_connectNotify_default")
	symaddr(&addr_QAbstractAnimation_disconnectNotify_default, "QAbstractAnimation_disconnectNotify_default")
	reg(&pg_QAbstractAnimation_connect_finished, "QAbstractAnimation_connect_finished")
	reg(&pg_QAbstractAnimation_connect_stateChanged, "QAbstractAnimation_connect_stateChanged")
	reg(&pg_QAbstractAnimation_connect_currentLoopChanged, "QAbstractAnimation_connect_currentLoopChanged")
	reg(&pg_QAbstractAnimation_connect_directionChanged, "QAbstractAnimation_connect_directionChanged")
	reg(&pg_QAbstractAnimation_connect_destroyed, "QAbstractAnimation_connect_destroyed")
	reg(&pg_QAbstractAnimation_connect_destroyed2, "QAbstractAnimation_connect_destroyed2")
	reg(&pg_QAbstractAnimation_connect_objectNameChanged, "QAbstractAnimation_connect_objectNameChanged")
	reg(&pg_QAbstractAnimation__disconnect, "QAbstractAnimation__disconnect")
	reg(&pg_QAbstractAnimation__set_override, "QAbstractAnimation__set_override")
	reg(&pg_QAbstractAnimation__Go_new, "QAbstractAnimation__Go_new")
	reg(&pg_QAbstractAnimation_delete, "QAbstractAnimation_delete")
}

type QAbstractAnimation struct {
	h uintptr
	*QObject
}

func bindQAbstractAnimation(h uintptr) *QAbstractAnimation {
	if h == 0 {
		return nil
	}
	return &QAbstractAnimation{h: h, QObject: bindQObject(pg_QAbstractAnimation_base(h))}
}

func newownedQAbstractAnimation(h uintptr) *QAbstractAnimation { return bindQAbstractAnimation(h) }

func (this *QAbstractAnimation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractAnimation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractAnimation(h unsafe.Pointer) *QAbstractAnimation { return bindQAbstractAnimation(uintptr(h)) }
func UnsafeNewQAbstractAnimationOwned(h unsafe.Pointer) *QAbstractAnimation { return newownedQAbstractAnimation(uintptr(h)) }

func (this *QAbstractAnimation) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAbstractAnimation_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractAnimation) State() QAbstractAnimation__State {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractAnimation__State(__r)
}

func (this *QAbstractAnimation) Group() *QAnimationGroup {
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_group, this.h)
	runtime.KeepAlive(this)
	return bindQAnimationGroup(__r)
}

func (this *QAbstractAnimation) Direction() QAbstractAnimation__Direction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_direction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractAnimation__Direction(__r)
}

func (this *QAbstractAnimation) SetDirection(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QAbstractAnimation_setDirection, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) CurrentTime() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_currentTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) CurrentLoopTime() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_currentLoopTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) LoopCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_loopCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) SetLoopCount(loopCount int) {
	purego.SyscallN(addr_QAbstractAnimation_setLoopCount, this.h, uintptr(loopCount))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) CurrentLoop() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_currentLoop, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) Duration() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_duration, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) TotalDuration() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_totalDuration, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) Finished() {
	purego.SyscallN(addr_QAbstractAnimation_finished, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) StateChanged(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QAbstractAnimation_stateChanged, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) CurrentLoopChanged(currentLoop int) {
	purego.SyscallN(addr_QAbstractAnimation_currentLoopChanged, this.h, uintptr(currentLoop))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) DirectionChanged(arg1 QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QAbstractAnimation_directionChanged, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) Start() {
	purego.SyscallN(addr_QAbstractAnimation_start, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) Pause() {
	purego.SyscallN(addr_QAbstractAnimation_pause, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) Resume() {
	purego.SyscallN(addr_QAbstractAnimation_resume, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) SetPaused(arg1 bool) {
	purego.SyscallN(addr_QAbstractAnimation_setPaused, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) Stop() {
	purego.SyscallN(addr_QAbstractAnimation_stop, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) SetCurrentTime(msecs int) {
	purego.SyscallN(addr_QAbstractAnimation_setCurrentTime, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
}

func QAbstractAnimation_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractAnimation_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractAnimation) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractAnimation) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractAnimation) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractAnimation) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractAnimation_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractAnimation_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAbstractAnimation) StartWithPolicy(policy QAbstractAnimation__DeletionPolicy) {
	purego.SyscallN(addr_QAbstractAnimation_start2, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func QAbstractAnimation_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractAnimation_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractAnimation) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractAnimation_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractAnimation) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QAbstractAnimation_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QAbstractAnimation_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QAbstractAnimation) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractAnimation_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractAnimation) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractAnimation_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractAnimation) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractAnimation_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractAnimation) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractAnimation_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractAnimation) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractAnimation_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractAnimation) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractAnimation_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractAnimation) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractAnimation_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QAbstractAnimation_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QAbstractAnimation_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QAbstractAnimation_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractAnimation_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractAnimation_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractAnimation_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractAnimation) Disconnect(handle int64) {
	pg_QAbstractAnimation__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractAnimation_virtIdx(slot int64) int32 {
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

func QAbstractAnimation_enable(h uintptr, idx int32) { pg_QAbstractAnimation__set_override(h, idx) }

func NewQAbstractAnimation() *QAbstractAnimation {
	id := nextVirtID()
	o := newownedQAbstractAnimation(pg_QAbstractAnimation__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractAnimation_virtIdx, QAbstractAnimation_enable)
	return o
}

func NewQAbstractAnimation2(parent *QObject) *QAbstractAnimation {
	id := nextVirtID()
	o := newownedQAbstractAnimation(pg_QAbstractAnimation__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractAnimation_virtIdx, QAbstractAnimation_enable)
	return o
}

func (this *QAbstractAnimation) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QAbstractAnimation) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(currentTime int32) { fn(int(currentTime)) }) // slot "updateCurrentTime"
}

func (this *QAbstractAnimation) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractAnimation) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractAnimation) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QAbstractAnimation) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QAbstractAnimation) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractAnimation) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractAnimation) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractAnimation) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractAnimation) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractAnimation) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractAnimation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractAnimation_delete(this.h)
	}
	this.UnsafeRelease()
}
