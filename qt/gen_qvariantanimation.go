package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVariantAnimation_chain uintptr
	addr_QVariantAnimation_new uintptr
	addr_QVariantAnimation_new2 uintptr
	addr_QVariantAnimation_tr uintptr
	addr_QVariantAnimation_startValue uintptr
	addr_QVariantAnimation_setStartValue uintptr
	addr_QVariantAnimation_endValue uintptr
	addr_QVariantAnimation_setEndValue uintptr
	pg_QVariantAnimation_keyValueAt func(uintptr, float64) uintptr
	pg_QVariantAnimation_setKeyValueAt func(uintptr, float64, uintptr)
	addr_QVariantAnimation_currentValue uintptr
	addr_QVariantAnimation_setDuration uintptr
	addr_QVariantAnimation_easingCurve uintptr
	addr_QVariantAnimation_setEasingCurve uintptr
	addr_QVariantAnimation_valueChanged uintptr
	addr_QVariantAnimation_connect uintptr
	addr_QVariantAnimation_connect2 uintptr
	addr_QVariantAnimation_sender uintptr
	addr_QVariantAnimation_senderSignalIndex uintptr
	addr_QVariantAnimation_receivers uintptr
	addr_QVariantAnimation_isSignalConnected uintptr
	addr_QVariantAnimation_tr2 uintptr
	addr_QVariantAnimation_tr3 uintptr
	addr_QVariantAnimation_connect3 uintptr
	addr_QVariantAnimation_connect4 uintptr
	addr_QVariantAnimation_metaObject_default uintptr
	addr_QVariantAnimation_duration_default uintptr
	addr_QVariantAnimation_event_default uintptr
	addr_QVariantAnimation_updateCurrentTime_default uintptr
	addr_QVariantAnimation_updateState_default uintptr
	addr_QVariantAnimation_updateCurrentValue_default uintptr
	pg_QVariantAnimation_interpolated_default func(uintptr, uintptr, uintptr, float64) uintptr
	addr_QVariantAnimation_updateDirection_default uintptr
	addr_QVariantAnimation_eventFilter_default uintptr
	addr_QVariantAnimation_timerEvent_default uintptr
	addr_QVariantAnimation_childEvent_default uintptr
	addr_QVariantAnimation_customEvent_default uintptr
	addr_QVariantAnimation_connectNotify_default uintptr
	addr_QVariantAnimation_disconnectNotify_default uintptr
	pg_QVariantAnimation_connect_valueChanged func(uintptr, int64)
	pg_QVariantAnimation_connect_finished func(uintptr, int64)
	pg_QVariantAnimation_connect_stateChanged func(uintptr, int64)
	pg_QVariantAnimation_connect_currentLoopChanged func(uintptr, int64)
	pg_QVariantAnimation_connect_directionChanged func(uintptr, int64)
	pg_QVariantAnimation_connect_destroyed func(uintptr, int64)
	pg_QVariantAnimation_connect_destroyed2 func(uintptr, int64)
	pg_QVariantAnimation_connect_objectNameChanged func(uintptr, int64)
	pg_QVariantAnimation__disconnect func(int64)
	pg_QVariantAnimation__set_override func(uintptr, int32)
	pg_QVariantAnimation_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVariantAnimation_chain, "QVariantAnimation_chain")
	symaddr(&addr_QVariantAnimation_new, "QVariantAnimation_new")
	symaddr(&addr_QVariantAnimation_new2, "QVariantAnimation_new2")
	symaddr(&addr_QVariantAnimation_tr, "QVariantAnimation_tr")
	symaddr(&addr_QVariantAnimation_startValue, "QVariantAnimation_startValue")
	symaddr(&addr_QVariantAnimation_setStartValue, "QVariantAnimation_setStartValue")
	symaddr(&addr_QVariantAnimation_endValue, "QVariantAnimation_endValue")
	symaddr(&addr_QVariantAnimation_setEndValue, "QVariantAnimation_setEndValue")
	reg(&pg_QVariantAnimation_keyValueAt, "QVariantAnimation_keyValueAt")
	reg(&pg_QVariantAnimation_setKeyValueAt, "QVariantAnimation_setKeyValueAt")
	symaddr(&addr_QVariantAnimation_currentValue, "QVariantAnimation_currentValue")
	symaddr(&addr_QVariantAnimation_setDuration, "QVariantAnimation_setDuration")
	symaddr(&addr_QVariantAnimation_easingCurve, "QVariantAnimation_easingCurve")
	symaddr(&addr_QVariantAnimation_setEasingCurve, "QVariantAnimation_setEasingCurve")
	symaddr(&addr_QVariantAnimation_valueChanged, "QVariantAnimation_valueChanged")
	symaddr(&addr_QVariantAnimation_connect, "QVariantAnimation_connect")
	symaddr(&addr_QVariantAnimation_connect2, "QVariantAnimation_connect2")
	symaddr(&addr_QVariantAnimation_sender, "QVariantAnimation_sender")
	symaddr(&addr_QVariantAnimation_senderSignalIndex, "QVariantAnimation_senderSignalIndex")
	symaddr(&addr_QVariantAnimation_receivers, "QVariantAnimation_receivers")
	symaddr(&addr_QVariantAnimation_isSignalConnected, "QVariantAnimation_isSignalConnected")
	symaddr(&addr_QVariantAnimation_tr2, "QVariantAnimation_tr2")
	symaddr(&addr_QVariantAnimation_tr3, "QVariantAnimation_tr3")
	symaddr(&addr_QVariantAnimation_connect3, "QVariantAnimation_connect3")
	symaddr(&addr_QVariantAnimation_connect4, "QVariantAnimation_connect4")
	symaddr(&addr_QVariantAnimation_metaObject_default, "QVariantAnimation_metaObject_default")
	symaddr(&addr_QVariantAnimation_duration_default, "QVariantAnimation_duration_default")
	symaddr(&addr_QVariantAnimation_event_default, "QVariantAnimation_event_default")
	symaddr(&addr_QVariantAnimation_updateCurrentTime_default, "QVariantAnimation_updateCurrentTime_default")
	symaddr(&addr_QVariantAnimation_updateState_default, "QVariantAnimation_updateState_default")
	symaddr(&addr_QVariantAnimation_updateCurrentValue_default, "QVariantAnimation_updateCurrentValue_default")
	reg(&pg_QVariantAnimation_interpolated_default, "QVariantAnimation_interpolated_default")
	symaddr(&addr_QVariantAnimation_updateDirection_default, "QVariantAnimation_updateDirection_default")
	symaddr(&addr_QVariantAnimation_eventFilter_default, "QVariantAnimation_eventFilter_default")
	symaddr(&addr_QVariantAnimation_timerEvent_default, "QVariantAnimation_timerEvent_default")
	symaddr(&addr_QVariantAnimation_childEvent_default, "QVariantAnimation_childEvent_default")
	symaddr(&addr_QVariantAnimation_customEvent_default, "QVariantAnimation_customEvent_default")
	symaddr(&addr_QVariantAnimation_connectNotify_default, "QVariantAnimation_connectNotify_default")
	symaddr(&addr_QVariantAnimation_disconnectNotify_default, "QVariantAnimation_disconnectNotify_default")
	reg(&pg_QVariantAnimation_connect_valueChanged, "QVariantAnimation_connect_valueChanged")
	reg(&pg_QVariantAnimation_connect_finished, "QVariantAnimation_connect_finished")
	reg(&pg_QVariantAnimation_connect_stateChanged, "QVariantAnimation_connect_stateChanged")
	reg(&pg_QVariantAnimation_connect_currentLoopChanged, "QVariantAnimation_connect_currentLoopChanged")
	reg(&pg_QVariantAnimation_connect_directionChanged, "QVariantAnimation_connect_directionChanged")
	reg(&pg_QVariantAnimation_connect_destroyed, "QVariantAnimation_connect_destroyed")
	reg(&pg_QVariantAnimation_connect_destroyed2, "QVariantAnimation_connect_destroyed2")
	reg(&pg_QVariantAnimation_connect_objectNameChanged, "QVariantAnimation_connect_objectNameChanged")
	reg(&pg_QVariantAnimation__disconnect, "QVariantAnimation__disconnect")
	reg(&pg_QVariantAnimation__set_override, "QVariantAnimation__set_override")
	reg(&pg_QVariantAnimation_delete, "QVariantAnimation_delete")
}

type QVariantAnimation struct {
	h uintptr
	*QAbstractAnimation
}

func bindQVariantAnimation(h uintptr) *QVariantAnimation {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QVariantAnimation_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QVariantAnimation{h: h, QAbstractAnimation: &QAbstractAnimation{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQVariantAnimation(h uintptr) *QVariantAnimation { return bindQVariantAnimation(h) }

func (this *QVariantAnimation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVariantAnimation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVariantAnimation(h unsafe.Pointer) *QVariantAnimation { return bindQVariantAnimation(uintptr(h)) }
func UnsafeNewQVariantAnimationOwned(h unsafe.Pointer) *QVariantAnimation { return newownedQVariantAnimation(uintptr(h)) }

func (this *QVariantAnimation) UnsafeRelease() {
	this.h = 0
	if this.QAbstractAnimation != nil {
		this.QAbstractAnimation.UnsafeRelease()
	}
}

func NewQVariantAnimation() *QVariantAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_new, uintptr(__goid))
	__o := newownedQVariantAnimation(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QVariantAnimation_virtIdx, QVariantAnimation_enable)
	return __o
}

func NewQVariantAnimation2(parent *QObject) *QVariantAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQVariantAnimation(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QVariantAnimation_virtIdx, QVariantAnimation_enable)
	return __o
}

func QVariantAnimation_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QVariantAnimation) StartValue() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_startValue, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVariantAnimation) SetStartValue(value *QVariant) {
	purego.SyscallN(addr_QVariantAnimation_setStartValue, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QVariantAnimation) EndValue() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_endValue, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVariantAnimation) SetEndValue(value *QVariant) {
	purego.SyscallN(addr_QVariantAnimation_setEndValue, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QVariantAnimation) KeyValueAt(step float64) *QVariant {
	__r := pg_QVariantAnimation_keyValueAt(this.h, step)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVariantAnimation) SetKeyValueAt(step float64, value *QVariant) {
	pg_QVariantAnimation_setKeyValueAt(this.h, step, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QVariantAnimation) CurrentValue() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_currentValue, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVariantAnimation) SetDuration(msecs int) {
	purego.SyscallN(addr_QVariantAnimation_setDuration, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
}

func (this *QVariantAnimation) EasingCurve() *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_easingCurve, this.h)
	runtime.KeepAlive(this)
	return newownedQEasingCurve(__r)
}

func (this *QVariantAnimation) SetEasingCurve(easing *QEasingCurve) {
	purego.SyscallN(addr_QVariantAnimation_setEasingCurve, this.h, uintptr(easing.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(easing)
}

func (this *QVariantAnimation) ValueChanged(value *QVariant) {
	purego.SyscallN(addr_QVariantAnimation_valueChanged, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func QVariantAnimation_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QVariantAnimation_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QVariantAnimation) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QVariantAnimation) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QVariantAnimation_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVariantAnimation) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QVariantAnimation_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QVariantAnimation) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariantAnimation_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QVariantAnimation_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QVariantAnimation_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QVariantAnimation_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QVariantAnimation_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QVariantAnimation) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QVariantAnimation_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QVariantAnimation) DurationDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QVariantAnimation_duration_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVariantAnimation) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariantAnimation_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QVariantAnimation) UpdateCurrentTimeDefault(arg1 int) {
	purego.SyscallN(addr_QVariantAnimation_updateCurrentTime_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QVariantAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	purego.SyscallN(addr_QVariantAnimation_updateState_default, this.h, uintptr(newState), uintptr(oldState))
	runtime.KeepAlive(this)
}

func (this *QVariantAnimation) UpdateCurrentValueDefault(value *QVariant) {
	purego.SyscallN(addr_QVariantAnimation_updateCurrentValue_default, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QVariantAnimation) InterpolatedDefault(from *QVariant, to *QVariant, progress float64) *QVariant {
	__r := pg_QVariantAnimation_interpolated_default(this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), progress)
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return newownedQVariant(__r)
}

func (this *QVariantAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	purego.SyscallN(addr_QVariantAnimation_updateDirection_default, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QVariantAnimation) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariantAnimation_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QVariantAnimation) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QVariantAnimation_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QVariantAnimation) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QVariantAnimation_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QVariantAnimation) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QVariantAnimation_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QVariantAnimation) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QVariantAnimation_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QVariantAnimation) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QVariantAnimation_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QVariantAnimation) OnValueChanged(fn func(*QVariant)) int64 {
	id := registerCallback(func(value uintptr) { fn(UnsafeNewQVariant(unsafe.Pointer(value))) })
	pg_QVariantAnimation_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QVariantAnimation_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnStateChanged(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) int64 {
	id := registerCallback(func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) })
	pg_QVariantAnimation_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnCurrentLoopChanged(fn func(int)) int64 {
	id := registerCallback(func(currentLoop int32) { fn(int(currentLoop)) })
	pg_QVariantAnimation_connect_currentLoopChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnDirectionChanged(fn func(QAbstractAnimation__Direction)) int64 {
	id := registerCallback(func(arg1 int32) { fn(QAbstractAnimation__Direction(arg1)) })
	pg_QVariantAnimation_connect_directionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QVariantAnimation_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QVariantAnimation_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QVariantAnimation_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVariantAnimation) Disconnect(handle int64) {
	pg_QVariantAnimation__disconnect(handle)
	unregisterCallback(handle)
}

func QVariantAnimation_virtIdx(slot int64) int32 {
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
	case -4717240701590708418:
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

func QVariantAnimation_enable(h uintptr, idx int32) { pg_QVariantAnimation__set_override(h, idx) }

func (this *QVariantAnimation) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QVariantAnimation) OnDuration(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8434675112194400115, func() int32 { return int32(fn()) }) // slot "duration"
}

func (this *QVariantAnimation) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QVariantAnimation) OnUpdateCurrentTime(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1342418961870150968, func(arg1 int32) { fn(int(arg1)) }) // slot "updateCurrentTime"
}

func (this *QVariantAnimation) OnUpdateState(fn func(QAbstractAnimation__State, QAbstractAnimation__State)) {
	registerVirt(this.UnsafeGoID(), 7692482089134926569, func(newState int32, oldState int32) { fn(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState)) }) // slot "updateState"
}

func (this *QVariantAnimation) OnUpdateCurrentValue(fn func(*QVariant)) {
	registerVirt(this.UnsafeGoID(), -4717240701590708418, func(value uintptr) { fn(UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "updateCurrentValue"
}

func (this *QVariantAnimation) OnInterpolated(fn func(*QVariant, *QVariant, float64) *QVariant) {
	registerVirt(this.UnsafeGoID(), 4185842742461511100, func(from uintptr, to uintptr, progress uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(from)), UnsafeNewQVariant(unsafe.Pointer(to)), math.Float64frombits(uint64(progress))).UnsafePointer()) }) // slot "interpolated"
}

func (this *QVariantAnimation) OnUpdateDirection(fn func(QAbstractAnimation__Direction)) {
	registerVirt(this.UnsafeGoID(), -3231888725407197643, func(direction int32) { fn(QAbstractAnimation__Direction(direction)) }) // slot "updateDirection"
}

func (this *QVariantAnimation) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QVariantAnimation) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QVariantAnimation) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QVariantAnimation) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QVariantAnimation) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QVariantAnimation) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QVariantAnimation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVariantAnimation_delete(this.h)
	}
	this.UnsafeRelease()
}
