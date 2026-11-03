package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPropertyAnimation_chain uintptr
	addr_QPropertyAnimation_new uintptr
	addr_QPropertyAnimation_new2 uintptr
	addr_QPropertyAnimation_new3 uintptr
	addr_QPropertyAnimation_new4 uintptr
	addr_QPropertyAnimation_tr uintptr
	addr_QPropertyAnimation_targetObject uintptr
	addr_QPropertyAnimation_setTargetObject uintptr
	addr_QPropertyAnimation_propertyName uintptr
	addr_QPropertyAnimation_setPropertyName uintptr
	addr_QPropertyAnimation_connect uintptr
	addr_QPropertyAnimation_connect2 uintptr
	addr_QPropertyAnimation_sender uintptr
	addr_QPropertyAnimation_senderSignalIndex uintptr
	addr_QPropertyAnimation_receivers uintptr
	addr_QPropertyAnimation_isSignalConnected uintptr
	addr_QPropertyAnimation_tr2 uintptr
	addr_QPropertyAnimation_tr3 uintptr
	addr_QPropertyAnimation_connect3 uintptr
	addr_QPropertyAnimation_connect4 uintptr
	addr_QPropertyAnimation_metaObject_default uintptr
	addr_QPropertyAnimation_event_default uintptr
	addr_QPropertyAnimation_updateCurrentValue_default uintptr
	addr_QPropertyAnimation_updateState_default uintptr
	addr_QPropertyAnimation_duration_default uintptr
	addr_QPropertyAnimation_updateCurrentTime_default uintptr
	pg_QPropertyAnimation_interpolated_default func(uintptr, uintptr, uintptr, float64) uintptr
	addr_QPropertyAnimation_updateDirection_default uintptr
	addr_QPropertyAnimation_eventFilter_default uintptr
	addr_QPropertyAnimation_timerEvent_default uintptr
	addr_QPropertyAnimation_childEvent_default uintptr
	addr_QPropertyAnimation_customEvent_default uintptr
	addr_QPropertyAnimation_connectNotify_default uintptr
	addr_QPropertyAnimation_disconnectNotify_default uintptr
	pg_QPropertyAnimation_connect_valueChanged func(uintptr, int64)
	pg_QPropertyAnimation_connect_finished func(uintptr, int64)
	pg_QPropertyAnimation_connect_stateChanged func(uintptr, int64)
	pg_QPropertyAnimation_connect_currentLoopChanged func(uintptr, int64)
	pg_QPropertyAnimation_connect_directionChanged func(uintptr, int64)
	pg_QPropertyAnimation_connect_destroyed func(uintptr, int64)
	pg_QPropertyAnimation_connect_destroyed2 func(uintptr, int64)
	pg_QPropertyAnimation_connect_objectNameChanged func(uintptr, int64)
	pg_QPropertyAnimation__disconnect func(int64)
	pg_QPropertyAnimation__set_override func(uintptr, int32)
	pg_QPropertyAnimation_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPropertyAnimation_chain, "QPropertyAnimation_chain")
	symaddr(&addr_QPropertyAnimation_new, "QPropertyAnimation_new")
	symaddr(&addr_QPropertyAnimation_new2, "QPropertyAnimation_new2")
	symaddr(&addr_QPropertyAnimation_new3, "QPropertyAnimation_new3")
	symaddr(&addr_QPropertyAnimation_new4, "QPropertyAnimation_new4")
	symaddr(&addr_QPropertyAnimation_tr, "QPropertyAnimation_tr")
	symaddr(&addr_QPropertyAnimation_targetObject, "QPropertyAnimation_targetObject")
	symaddr(&addr_QPropertyAnimation_setTargetObject, "QPropertyAnimation_setTargetObject")
	symaddr(&addr_QPropertyAnimation_propertyName, "QPropertyAnimation_propertyName")
	symaddr(&addr_QPropertyAnimation_setPropertyName, "QPropertyAnimation_setPropertyName")
	symaddr(&addr_QPropertyAnimation_connect, "QPropertyAnimation_connect")
	symaddr(&addr_QPropertyAnimation_connect2, "QPropertyAnimation_connect2")
	symaddr(&addr_QPropertyAnimation_sender, "QPropertyAnimation_sender")
	symaddr(&addr_QPropertyAnimation_senderSignalIndex, "QPropertyAnimation_senderSignalIndex")
	symaddr(&addr_QPropertyAnimation_receivers, "QPropertyAnimation_receivers")
	symaddr(&addr_QPropertyAnimation_isSignalConnected, "QPropertyAnimation_isSignalConnected")
	symaddr(&addr_QPropertyAnimation_tr2, "QPropertyAnimation_tr2")
	symaddr(&addr_QPropertyAnimation_tr3, "QPropertyAnimation_tr3")
	symaddr(&addr_QPropertyAnimation_connect3, "QPropertyAnimation_connect3")
	symaddr(&addr_QPropertyAnimation_connect4, "QPropertyAnimation_connect4")
	symaddr(&addr_QPropertyAnimation_metaObject_default, "QPropertyAnimation_metaObject_default")
	symaddr(&addr_QPropertyAnimation_event_default, "QPropertyAnimation_event_default")
	symaddr(&addr_QPropertyAnimation_updateCurrentValue_default, "QPropertyAnimation_updateCurrentValue_default")
	symaddr(&addr_QPropertyAnimation_updateState_default, "QPropertyAnimation_updateState_default")
	symaddr(&addr_QPropertyAnimation_duration_default, "QPropertyAnimation_duration_default")
	symaddr(&addr_QPropertyAnimation_updateCurrentTime_default, "QPropertyAnimation_updateCurrentTime_default")
	reg(&pg_QPropertyAnimation_interpolated_default, "QPropertyAnimation_interpolated_default")
	symaddr(&addr_QPropertyAnimation_updateDirection_default, "QPropertyAnimation_updateDirection_default")
	symaddr(&addr_QPropertyAnimation_eventFilter_default, "QPropertyAnimation_eventFilter_default")
	symaddr(&addr_QPropertyAnimation_timerEvent_default, "QPropertyAnimation_timerEvent_default")
	symaddr(&addr_QPropertyAnimation_childEvent_default, "QPropertyAnimation_childEvent_default")
	symaddr(&addr_QPropertyAnimation_customEvent_default, "QPropertyAnimation_customEvent_default")
	symaddr(&addr_QPropertyAnimation_connectNotify_default, "QPropertyAnimation_connectNotify_default")
	symaddr(&addr_QPropertyAnimation_disconnectNotify_default, "QPropertyAnimation_disconnectNotify_default")
	reg(&pg_QPropertyAnimation_connect_valueChanged, "QPropertyAnimation_connect_valueChanged")
	reg(&pg_QPropertyAnimation_connect_finished, "QPropertyAnimation_connect_finished")
	reg(&pg_QPropertyAnimation_connect_stateChanged, "QPropertyAnimation_connect_stateChanged")
	reg(&pg_QPropertyAnimation_connect_currentLoopChanged, "QPropertyAnimation_connect_currentLoopChanged")
	reg(&pg_QPropertyAnimation_connect_directionChanged, "QPropertyAnimation_connect_directionChanged")
	reg(&pg_QPropertyAnimation_connect_destroyed, "QPropertyAnimation_connect_destroyed")
	reg(&pg_QPropertyAnimation_connect_destroyed2, "QPropertyAnimation_connect_destroyed2")
	reg(&pg_QPropertyAnimation_connect_objectNameChanged, "QPropertyAnimation_connect_objectNameChanged")
	reg(&pg_QPropertyAnimation__disconnect, "QPropertyAnimation__disconnect")
	reg(&pg_QPropertyAnimation__set_override, "QPropertyAnimation__set_override")
	reg(&pg_QPropertyAnimation_delete, "QPropertyAnimation_delete")
}

type QPropertyAnimation struct {
	h uintptr
	*QVariantAnimation
}

func bindQPropertyAnimation(h uintptr) *QPropertyAnimation {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QPropertyAnimation_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPropertyAnimation{h: h, QVariantAnimation: &QVariantAnimation{h: __c[0], QAbstractAnimation: &QAbstractAnimation{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQPropertyAnimation(h uintptr) *QPropertyAnimation { return bindQPropertyAnimation(h) }

func (this *QPropertyAnimation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyAnimation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyAnimation(h unsafe.Pointer) *QPropertyAnimation { return bindQPropertyAnimation(uintptr(h)) }
func UnsafeNewQPropertyAnimationOwned(h unsafe.Pointer) *QPropertyAnimation { return newownedQPropertyAnimation(uintptr(h)) }

func (this *QPropertyAnimation) UnsafeRelease() {
	this.h = 0
	if this.QVariantAnimation != nil {
		this.QVariantAnimation.UnsafeRelease()
	}
}

func NewQPropertyAnimation() *QPropertyAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_new, uintptr(__goid))
	__o := newownedQPropertyAnimation(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPropertyAnimation_virtIdx, QPropertyAnimation_enable)
	return __o
}

func NewQPropertyAnimation2(target *QObject, propertyName *QByteArray) *QPropertyAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_new2, uintptr(__goid), uintptr(target.UnsafePointer()), uintptr(propertyName.UnsafePointer()))
	__o := newownedQPropertyAnimation(__r)
	runtime.KeepAlive(target)
	runtime.KeepAlive(propertyName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPropertyAnimation_virtIdx, QPropertyAnimation_enable)
	return __o
}

func NewQPropertyAnimation3(parent *QObject) *QPropertyAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPropertyAnimation(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPropertyAnimation_virtIdx, QPropertyAnimation_enable)
	return __o
}

func NewQPropertyAnimation4(target *QObject, propertyName *QByteArray, parent *QObject) *QPropertyAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_new4, uintptr(__goid), uintptr(target.UnsafePointer()), uintptr(propertyName.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQPropertyAnimation(__r)
	runtime.KeepAlive(target)
	runtime.KeepAlive(propertyName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPropertyAnimation_virtIdx, QPropertyAnimation_enable)
	return __o
}

func QPropertyAnimation_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPropertyAnimation) TargetObject() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_targetObject, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPropertyAnimation) SetTargetObject(target *QObject) {
	purego.SyscallN(addr_QPropertyAnimation_setTargetObject, this.h, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
}

func (this *QPropertyAnimation) PropertyName() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_propertyName, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QPropertyAnimation) SetPropertyName(propertyName *QByteArray) {
	purego.SyscallN(addr_QPropertyAnimation_setPropertyName, this.h, uintptr(propertyName.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(propertyName)
}

func QPropertyAnimation_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPropertyAnimation_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPropertyAnimation) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPropertyAnimation) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPropertyAnimation_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPropertyAnimation) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPropertyAnimation_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPropertyAnimation) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPropertyAnimation_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPropertyAnimation_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPropertyAnimation_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPropertyAnimation_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPropertyAnimation_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPropertyAnimation) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPropertyAnimation_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPropertyAnimation) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPropertyAnimation_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPropertyAnimation) UpdateCurrentValueDefault(value *QVariant) {
	purego.SyscallN(addr_QPropertyAnimation_updateCurrentValue_default, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QPropertyAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QPropertyAnimation_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QPropertyAnimation) DurationDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QPropertyAnimation_duration_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPropertyAnimation) UpdateCurrentTimeDefault(arg1 int) {
	purego.SyscallN(addr_QPropertyAnimation_updateCurrentTime_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QPropertyAnimation) InterpolatedDefault(from *QVariant, to *QVariant, progress float64) *QVariant {
	__r := pg_QPropertyAnimation_interpolated_default(this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), progress)
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return newownedQVariant(__r)
}

func (this *QPropertyAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QPropertyAnimation_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QPropertyAnimation) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPropertyAnimation_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPropertyAnimation) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QPropertyAnimation_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPropertyAnimation) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPropertyAnimation_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPropertyAnimation) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPropertyAnimation_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPropertyAnimation) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPropertyAnimation_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPropertyAnimation) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPropertyAnimation_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPropertyAnimation) OnValueChanged(fn func(*QVariant)) int64 {
	id := registerCallback(func(value uintptr) { fn(UnsafeNewQVariant(unsafe.Pointer(value))) })
	pg_QPropertyAnimation_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPropertyAnimation_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QPropertyAnimation_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QPropertyAnimation_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QPropertyAnimation_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPropertyAnimation_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPropertyAnimation_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPropertyAnimation_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPropertyAnimation) Disconnect(handle int64) {
	pg_QPropertyAnimation__disconnect(handle)
	unregisterCallback(handle)
}

func QPropertyAnimation_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case -4717240701590708418:
		return 2
	case 7692482089134926569:
		return 3
	case -8434675112194400115:
		return 4
	case -1342418961870150968:
		return 5
	case 4185842742461511100:
		return 6
	case -3231888725407197643:
		return 7
	case 5217623514630858297:
		return 8
	case -5409927581812374498:
		return 9
	case -1185315367526857833:
		return 10
	case 3769868229123868474:
		return 11
	case 7379029982170503008:
		return 12
	case -568369772982215962:
		return 13
	}
	return -1
}

func QPropertyAnimation_enable(h uintptr, idx int32) { pg_QPropertyAnimation__set_override(h, idx) }

func (this *QPropertyAnimation) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPropertyAnimation) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QPropertyAnimation) OnUpdateCurrentValue(fn func(*QVariant)) {
	registerVirt(this.UnsafeGoID(), -4717240701590708418, func(value uintptr) { fn(UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "updateCurrentValue"
}

func (this *QPropertyAnimation) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QPropertyAnimation) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QPropertyAnimation) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(arg1 int32) { fn(int(arg1)) }) // slot "updateCurrentTime"
}

func (this *QPropertyAnimation) OnInterpolated(fn func(*QVariant, *QVariant, float64) *QVariant) {
	registerVirt(this.UnsafeGoID(), 4185842742461511100, func(from uintptr, to uintptr, progress uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(from)), UnsafeNewQVariant(unsafe.Pointer(to)), math.Float64frombits(uint64(progress))).UnsafePointer()) }) // slot "interpolated"
}

func (this *QPropertyAnimation) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QPropertyAnimation) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPropertyAnimation) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QPropertyAnimation) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPropertyAnimation) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPropertyAnimation) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPropertyAnimation) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPropertyAnimation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyAnimation_delete(this.h)
	}
	this.UnsafeRelease()
}
