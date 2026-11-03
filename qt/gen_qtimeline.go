package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTimeLine_base func(uintptr) uintptr
	addr_QTimeLine_new uintptr
	addr_QTimeLine_new2 uintptr
	addr_QTimeLine_new3 uintptr
	addr_QTimeLine_tr uintptr
	addr_QTimeLine_state uintptr
	addr_QTimeLine_loopCount uintptr
	addr_QTimeLine_setLoopCount uintptr
	addr_QTimeLine_direction uintptr
	addr_QTimeLine_setDirection uintptr
	addr_QTimeLine_duration uintptr
	addr_QTimeLine_setDuration uintptr
	addr_QTimeLine_startFrame uintptr
	addr_QTimeLine_setStartFrame uintptr
	addr_QTimeLine_endFrame uintptr
	addr_QTimeLine_setEndFrame uintptr
	addr_QTimeLine_setFrameRange uintptr
	addr_QTimeLine_updateInterval uintptr
	addr_QTimeLine_setUpdateInterval uintptr
	addr_QTimeLine_easingCurve uintptr
	addr_QTimeLine_setEasingCurve uintptr
	addr_QTimeLine_currentTime uintptr
	addr_QTimeLine_currentFrame uintptr
	pg_QTimeLine_currentValue func(uintptr) float64
	addr_QTimeLine_frameForTime uintptr
	pg_QTimeLine_valueForTime func(uintptr, int32) float64
	addr_QTimeLine_start uintptr
	addr_QTimeLine_resume uintptr
	addr_QTimeLine_stop uintptr
	addr_QTimeLine_setPaused uintptr
	addr_QTimeLine_setCurrentTime uintptr
	addr_QTimeLine_toggleDirection uintptr
	addr_QTimeLine_connect uintptr
	addr_QTimeLine_connect2 uintptr
	addr_QTimeLine_sender uintptr
	addr_QTimeLine_senderSignalIndex uintptr
	addr_QTimeLine_receivers uintptr
	addr_QTimeLine_isSignalConnected uintptr
	addr_QTimeLine_tr2 uintptr
	addr_QTimeLine_tr3 uintptr
	addr_QTimeLine_connect3 uintptr
	addr_QTimeLine_connect4 uintptr
	pg_QTimeLine_valueForTime_default func(uintptr, int32) float64
	addr_QTimeLine_metaObject_default uintptr
	addr_QTimeLine_timerEvent_default uintptr
	addr_QTimeLine_event_default uintptr
	addr_QTimeLine_eventFilter_default uintptr
	addr_QTimeLine_childEvent_default uintptr
	addr_QTimeLine_customEvent_default uintptr
	addr_QTimeLine_connectNotify_default uintptr
	addr_QTimeLine_disconnectNotify_default uintptr
	pg_QTimeLine_connect_valueChanged func(uintptr, int64)
	pg_QTimeLine_connect_frameChanged func(uintptr, int64)
	pg_QTimeLine_connect_stateChanged func(uintptr, int64)
	pg_QTimeLine_connect_finished func(uintptr, int64)
	pg_QTimeLine_connect_destroyed func(uintptr, int64)
	pg_QTimeLine_connect_destroyed2 func(uintptr, int64)
	pg_QTimeLine_connect_objectNameChanged func(uintptr, int64)
	pg_QTimeLine__disconnect func(int64)
	pg_QTimeLine__set_override func(uintptr, int32)
	pg_QTimeLine_delete func(uintptr)
)

func init() {
	reg(&pg_QTimeLine_base, "QTimeLine_base")
	symaddr(&addr_QTimeLine_new, "QTimeLine_new")
	symaddr(&addr_QTimeLine_new2, "QTimeLine_new2")
	symaddr(&addr_QTimeLine_new3, "QTimeLine_new3")
	symaddr(&addr_QTimeLine_tr, "QTimeLine_tr")
	symaddr(&addr_QTimeLine_state, "QTimeLine_state")
	symaddr(&addr_QTimeLine_loopCount, "QTimeLine_loopCount")
	symaddr(&addr_QTimeLine_setLoopCount, "QTimeLine_setLoopCount")
	symaddr(&addr_QTimeLine_direction, "QTimeLine_direction")
	symaddr(&addr_QTimeLine_setDirection, "QTimeLine_setDirection")
	symaddr(&addr_QTimeLine_duration, "QTimeLine_duration")
	symaddr(&addr_QTimeLine_setDuration, "QTimeLine_setDuration")
	symaddr(&addr_QTimeLine_startFrame, "QTimeLine_startFrame")
	symaddr(&addr_QTimeLine_setStartFrame, "QTimeLine_setStartFrame")
	symaddr(&addr_QTimeLine_endFrame, "QTimeLine_endFrame")
	symaddr(&addr_QTimeLine_setEndFrame, "QTimeLine_setEndFrame")
	symaddr(&addr_QTimeLine_setFrameRange, "QTimeLine_setFrameRange")
	symaddr(&addr_QTimeLine_updateInterval, "QTimeLine_updateInterval")
	symaddr(&addr_QTimeLine_setUpdateInterval, "QTimeLine_setUpdateInterval")
	symaddr(&addr_QTimeLine_easingCurve, "QTimeLine_easingCurve")
	symaddr(&addr_QTimeLine_setEasingCurve, "QTimeLine_setEasingCurve")
	symaddr(&addr_QTimeLine_currentTime, "QTimeLine_currentTime")
	symaddr(&addr_QTimeLine_currentFrame, "QTimeLine_currentFrame")
	reg(&pg_QTimeLine_currentValue, "QTimeLine_currentValue")
	symaddr(&addr_QTimeLine_frameForTime, "QTimeLine_frameForTime")
	reg(&pg_QTimeLine_valueForTime, "QTimeLine_valueForTime")
	symaddr(&addr_QTimeLine_start, "QTimeLine_start")
	symaddr(&addr_QTimeLine_resume, "QTimeLine_resume")
	symaddr(&addr_QTimeLine_stop, "QTimeLine_stop")
	symaddr(&addr_QTimeLine_setPaused, "QTimeLine_setPaused")
	symaddr(&addr_QTimeLine_setCurrentTime, "QTimeLine_setCurrentTime")
	symaddr(&addr_QTimeLine_toggleDirection, "QTimeLine_toggleDirection")
	symaddr(&addr_QTimeLine_connect, "QTimeLine_connect")
	symaddr(&addr_QTimeLine_connect2, "QTimeLine_connect2")
	symaddr(&addr_QTimeLine_sender, "QTimeLine_sender")
	symaddr(&addr_QTimeLine_senderSignalIndex, "QTimeLine_senderSignalIndex")
	symaddr(&addr_QTimeLine_receivers, "QTimeLine_receivers")
	symaddr(&addr_QTimeLine_isSignalConnected, "QTimeLine_isSignalConnected")
	symaddr(&addr_QTimeLine_tr2, "QTimeLine_tr2")
	symaddr(&addr_QTimeLine_tr3, "QTimeLine_tr3")
	symaddr(&addr_QTimeLine_connect3, "QTimeLine_connect3")
	symaddr(&addr_QTimeLine_connect4, "QTimeLine_connect4")
	reg(&pg_QTimeLine_valueForTime_default, "QTimeLine_valueForTime_default")
	symaddr(&addr_QTimeLine_metaObject_default, "QTimeLine_metaObject_default")
	symaddr(&addr_QTimeLine_timerEvent_default, "QTimeLine_timerEvent_default")
	symaddr(&addr_QTimeLine_event_default, "QTimeLine_event_default")
	symaddr(&addr_QTimeLine_eventFilter_default, "QTimeLine_eventFilter_default")
	symaddr(&addr_QTimeLine_childEvent_default, "QTimeLine_childEvent_default")
	symaddr(&addr_QTimeLine_customEvent_default, "QTimeLine_customEvent_default")
	symaddr(&addr_QTimeLine_connectNotify_default, "QTimeLine_connectNotify_default")
	symaddr(&addr_QTimeLine_disconnectNotify_default, "QTimeLine_disconnectNotify_default")
	reg(&pg_QTimeLine_connect_valueChanged, "QTimeLine_connect_valueChanged")
	reg(&pg_QTimeLine_connect_frameChanged, "QTimeLine_connect_frameChanged")
	reg(&pg_QTimeLine_connect_stateChanged, "QTimeLine_connect_stateChanged")
	reg(&pg_QTimeLine_connect_finished, "QTimeLine_connect_finished")
	reg(&pg_QTimeLine_connect_destroyed, "QTimeLine_connect_destroyed")
	reg(&pg_QTimeLine_connect_destroyed2, "QTimeLine_connect_destroyed2")
	reg(&pg_QTimeLine_connect_objectNameChanged, "QTimeLine_connect_objectNameChanged")
	reg(&pg_QTimeLine__disconnect, "QTimeLine__disconnect")
	reg(&pg_QTimeLine__set_override, "QTimeLine__set_override")
	reg(&pg_QTimeLine_delete, "QTimeLine_delete")
}

type QTimeLine struct {
	h uintptr
	*QObject
}

func bindQTimeLine(h uintptr) *QTimeLine {
	if h == 0 {
		return nil
	}
	return &QTimeLine{h: h, QObject: bindQObject(pg_QTimeLine_base(h))}
}

func newownedQTimeLine(h uintptr) *QTimeLine { return bindQTimeLine(h) }

func (this *QTimeLine) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTimeLine) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTimeLine(h unsafe.Pointer) *QTimeLine { return bindQTimeLine(uintptr(h)) }
func UnsafeNewQTimeLineOwned(h unsafe.Pointer) *QTimeLine { return newownedQTimeLine(uintptr(h)) }

func (this *QTimeLine) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQTimeLine() *QTimeLine {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTimeLine_new, uintptr(__goid))
	__o := newownedQTimeLine(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTimeLine_virtIdx, QTimeLine_enable)
	return __o
}

func NewQTimeLine2(duration int) *QTimeLine {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTimeLine_new2, uintptr(__goid), uintptr(duration))
	__o := newownedQTimeLine(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTimeLine_virtIdx, QTimeLine_enable)
	return __o
}

func NewQTimeLine3(duration int, parent *QObject) *QTimeLine {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTimeLine_new3, uintptr(__goid), uintptr(duration), uintptr(parent.UnsafePointer()))
	__o := newownedQTimeLine(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTimeLine_virtIdx, QTimeLine_enable)
	return __o
}

func QTimeLine_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeLine_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTimeLine) State() QTimeLine__State {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTimeLine__State(__r)
}

func (this *QTimeLine) LoopCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_loopCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) SetLoopCount(count int) {
	purego.SyscallN(addr_QTimeLine_setLoopCount, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) Direction() QTimeLine__Direction {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_direction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTimeLine__Direction(__r)
}

func (this *QTimeLine) SetDirection(direction QTimeLine__Direction) {
	purego.SyscallN(addr_QTimeLine_setDirection, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) Duration() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_duration, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) SetDuration(duration int) {
	purego.SyscallN(addr_QTimeLine_setDuration, this.h, uintptr(duration))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) StartFrame() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_startFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) SetStartFrame(frame int) {
	purego.SyscallN(addr_QTimeLine_setStartFrame, this.h, uintptr(frame))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) EndFrame() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_endFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) SetEndFrame(frame int) {
	purego.SyscallN(addr_QTimeLine_setEndFrame, this.h, uintptr(frame))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	purego.SyscallN(addr_QTimeLine_setFrameRange, this.h, uintptr(startFrame), uintptr(endFrame))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) UpdateInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_updateInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) SetUpdateInterval(interval int) {
	purego.SyscallN(addr_QTimeLine_setUpdateInterval, this.h, uintptr(interval))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) EasingCurve() *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QTimeLine_easingCurve, this.h)
	runtime.KeepAlive(this)
	return newownedQEasingCurve(__r)
}

func (this *QTimeLine) SetEasingCurve(curve *QEasingCurve) {
	purego.SyscallN(addr_QTimeLine_setEasingCurve, this.h, uintptr(curve.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(curve)
}

func (this *QTimeLine) CurrentTime() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_currentTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) CurrentFrame() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_currentFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) CurrentValue() float64 {
	__r := pg_QTimeLine_currentValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTimeLine) FrameForTime(msec int) int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_frameForTime, this.h, uintptr(msec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) ValueForTime(msec int) float64 {
	__r := pg_QTimeLine_valueForTime(this.h, int32(msec))
	runtime.KeepAlive(this)
	return __r
}

func (this *QTimeLine) Start() {
	purego.SyscallN(addr_QTimeLine_start, this.h)
	runtime.KeepAlive(this)
}

func (this *QTimeLine) Resume() {
	purego.SyscallN(addr_QTimeLine_resume, this.h)
	runtime.KeepAlive(this)
}

func (this *QTimeLine) Stop() {
	purego.SyscallN(addr_QTimeLine_stop, this.h)
	runtime.KeepAlive(this)
}

func (this *QTimeLine) SetPaused(paused bool) {
	purego.SyscallN(addr_QTimeLine_setPaused, this.h, uintptr(b2i(paused)))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) SetCurrentTime(msec int) {
	purego.SyscallN(addr_QTimeLine_setCurrentTime, this.h, uintptr(msec))
	runtime.KeepAlive(this)
}

func (this *QTimeLine) ToggleDirection() {
	purego.SyscallN(addr_QTimeLine_toggleDirection, this.h)
	runtime.KeepAlive(this)
}

func QTimeLine_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTimeLine_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTimeLine_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTimeLine_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTimeLine) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTimeLine_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTimeLine) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeLine) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTimeLine) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTimeLine_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeLine_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTimeLine_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeLine_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTimeLine_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTimeLine_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTimeLine_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTimeLine_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTimeLine) ValueForTimeDefault(msec int) float64 {
	__r := pg_QTimeLine_valueForTime_default(this.h, int32(msec))
	runtime.KeepAlive(this)
	return __r
}

func (this *QTimeLine) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTimeLine_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTimeLine) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTimeLine_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTimeLine) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTimeLine) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeLine_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTimeLine) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTimeLine_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTimeLine) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTimeLine_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTimeLine) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTimeLine_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTimeLine) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTimeLine_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTimeLine) OnValueChanged(fn func(float64)) int64 {
	id := registerCallback(func(x uintptr) { fn(math.Float64frombits(uint64(x))) })
	pg_QTimeLine_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) OnFrameChanged(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QTimeLine_connect_frameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) OnStateChanged(fn func(QTimeLine__State)) int64 {
	id := registerCallback(func(newState int32) { fn(QTimeLine__State(newState)) })
	pg_QTimeLine_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTimeLine_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTimeLine_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTimeLine_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTimeLine_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTimeLine) Disconnect(handle int64) {
	pg_QTimeLine__disconnect(handle)
	unregisterCallback(handle)
}

func QTimeLine_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 3548770768539434384:
		return 1
	case -5409927581812374498:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
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

func QTimeLine_enable(h uintptr, idx int32) { pg_QTimeLine__set_override(h, idx) }

func (this *QTimeLine) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTimeLine) OnValueForTime(fn func(int) float64) {
	registerVirt(this.UnsafeGoID(), 3548770768539434384, func(msec int32) uintptr { return uintptr(math.Float64bits(fn(int(msec)))) }) // slot "valueForTime"
}

func (this *QTimeLine) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTimeLine) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTimeLine) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTimeLine) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTimeLine) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTimeLine) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTimeLine) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTimeLine) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTimeLine_delete(this.h)
	}
	this.UnsafeRelease()
}
