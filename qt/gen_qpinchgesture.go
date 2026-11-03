package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPinchGesture_chain uintptr
	addr_QPinchGesture_new uintptr
	addr_QPinchGesture_new2 uintptr
	addr_QPinchGesture_tr uintptr
	addr_QPinchGesture_totalChangeFlags uintptr
	addr_QPinchGesture_setTotalChangeFlags uintptr
	addr_QPinchGesture_changeFlags uintptr
	addr_QPinchGesture_setChangeFlags uintptr
	addr_QPinchGesture_startCenterPoint uintptr
	addr_QPinchGesture_lastCenterPoint uintptr
	addr_QPinchGesture_centerPoint uintptr
	addr_QPinchGesture_setStartCenterPoint uintptr
	addr_QPinchGesture_setLastCenterPoint uintptr
	addr_QPinchGesture_setCenterPoint uintptr
	pg_QPinchGesture_totalScaleFactor func(uintptr) float64
	pg_QPinchGesture_lastScaleFactor func(uintptr) float64
	pg_QPinchGesture_scaleFactor func(uintptr) float64
	pg_QPinchGesture_setTotalScaleFactor func(uintptr, float64)
	pg_QPinchGesture_setLastScaleFactor func(uintptr, float64)
	pg_QPinchGesture_setScaleFactor func(uintptr, float64)
	pg_QPinchGesture_totalRotationAngle func(uintptr) float64
	pg_QPinchGesture_lastRotationAngle func(uintptr) float64
	pg_QPinchGesture_rotationAngle func(uintptr) float64
	pg_QPinchGesture_setTotalRotationAngle func(uintptr, float64)
	pg_QPinchGesture_setLastRotationAngle func(uintptr, float64)
	pg_QPinchGesture_setRotationAngle func(uintptr, float64)
	addr_QPinchGesture_connect uintptr
	addr_QPinchGesture_connect2 uintptr
	addr_QPinchGesture_sender uintptr
	addr_QPinchGesture_senderSignalIndex uintptr
	addr_QPinchGesture_receivers uintptr
	addr_QPinchGesture_isSignalConnected uintptr
	addr_QPinchGesture_tr2 uintptr
	addr_QPinchGesture_tr3 uintptr
	addr_QPinchGesture_connect3 uintptr
	addr_QPinchGesture_connect4 uintptr
	addr_QPinchGesture_metaObject_default uintptr
	addr_QPinchGesture_event_default uintptr
	addr_QPinchGesture_eventFilter_default uintptr
	addr_QPinchGesture_timerEvent_default uintptr
	addr_QPinchGesture_childEvent_default uintptr
	addr_QPinchGesture_customEvent_default uintptr
	addr_QPinchGesture_connectNotify_default uintptr
	addr_QPinchGesture_disconnectNotify_default uintptr
	pg_QPinchGesture_connect_destroyed func(uintptr, int64)
	pg_QPinchGesture_connect_destroyed2 func(uintptr, int64)
	pg_QPinchGesture_connect_objectNameChanged func(uintptr, int64)
	pg_QPinchGesture__disconnect func(int64)
	pg_QPinchGesture__set_override func(uintptr, int32)
	pg_QPinchGesture_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPinchGesture_chain, "QPinchGesture_chain")
	symaddr(&addr_QPinchGesture_new, "QPinchGesture_new")
	symaddr(&addr_QPinchGesture_new2, "QPinchGesture_new2")
	symaddr(&addr_QPinchGesture_tr, "QPinchGesture_tr")
	symaddr(&addr_QPinchGesture_totalChangeFlags, "QPinchGesture_totalChangeFlags")
	symaddr(&addr_QPinchGesture_setTotalChangeFlags, "QPinchGesture_setTotalChangeFlags")
	symaddr(&addr_QPinchGesture_changeFlags, "QPinchGesture_changeFlags")
	symaddr(&addr_QPinchGesture_setChangeFlags, "QPinchGesture_setChangeFlags")
	symaddr(&addr_QPinchGesture_startCenterPoint, "QPinchGesture_startCenterPoint")
	symaddr(&addr_QPinchGesture_lastCenterPoint, "QPinchGesture_lastCenterPoint")
	symaddr(&addr_QPinchGesture_centerPoint, "QPinchGesture_centerPoint")
	symaddr(&addr_QPinchGesture_setStartCenterPoint, "QPinchGesture_setStartCenterPoint")
	symaddr(&addr_QPinchGesture_setLastCenterPoint, "QPinchGesture_setLastCenterPoint")
	symaddr(&addr_QPinchGesture_setCenterPoint, "QPinchGesture_setCenterPoint")
	reg(&pg_QPinchGesture_totalScaleFactor, "QPinchGesture_totalScaleFactor")
	reg(&pg_QPinchGesture_lastScaleFactor, "QPinchGesture_lastScaleFactor")
	reg(&pg_QPinchGesture_scaleFactor, "QPinchGesture_scaleFactor")
	reg(&pg_QPinchGesture_setTotalScaleFactor, "QPinchGesture_setTotalScaleFactor")
	reg(&pg_QPinchGesture_setLastScaleFactor, "QPinchGesture_setLastScaleFactor")
	reg(&pg_QPinchGesture_setScaleFactor, "QPinchGesture_setScaleFactor")
	reg(&pg_QPinchGesture_totalRotationAngle, "QPinchGesture_totalRotationAngle")
	reg(&pg_QPinchGesture_lastRotationAngle, "QPinchGesture_lastRotationAngle")
	reg(&pg_QPinchGesture_rotationAngle, "QPinchGesture_rotationAngle")
	reg(&pg_QPinchGesture_setTotalRotationAngle, "QPinchGesture_setTotalRotationAngle")
	reg(&pg_QPinchGesture_setLastRotationAngle, "QPinchGesture_setLastRotationAngle")
	reg(&pg_QPinchGesture_setRotationAngle, "QPinchGesture_setRotationAngle")
	symaddr(&addr_QPinchGesture_connect, "QPinchGesture_connect")
	symaddr(&addr_QPinchGesture_connect2, "QPinchGesture_connect2")
	symaddr(&addr_QPinchGesture_sender, "QPinchGesture_sender")
	symaddr(&addr_QPinchGesture_senderSignalIndex, "QPinchGesture_senderSignalIndex")
	symaddr(&addr_QPinchGesture_receivers, "QPinchGesture_receivers")
	symaddr(&addr_QPinchGesture_isSignalConnected, "QPinchGesture_isSignalConnected")
	symaddr(&addr_QPinchGesture_tr2, "QPinchGesture_tr2")
	symaddr(&addr_QPinchGesture_tr3, "QPinchGesture_tr3")
	symaddr(&addr_QPinchGesture_connect3, "QPinchGesture_connect3")
	symaddr(&addr_QPinchGesture_connect4, "QPinchGesture_connect4")
	symaddr(&addr_QPinchGesture_metaObject_default, "QPinchGesture_metaObject_default")
	symaddr(&addr_QPinchGesture_event_default, "QPinchGesture_event_default")
	symaddr(&addr_QPinchGesture_eventFilter_default, "QPinchGesture_eventFilter_default")
	symaddr(&addr_QPinchGesture_timerEvent_default, "QPinchGesture_timerEvent_default")
	symaddr(&addr_QPinchGesture_childEvent_default, "QPinchGesture_childEvent_default")
	symaddr(&addr_QPinchGesture_customEvent_default, "QPinchGesture_customEvent_default")
	symaddr(&addr_QPinchGesture_connectNotify_default, "QPinchGesture_connectNotify_default")
	symaddr(&addr_QPinchGesture_disconnectNotify_default, "QPinchGesture_disconnectNotify_default")
	reg(&pg_QPinchGesture_connect_destroyed, "QPinchGesture_connect_destroyed")
	reg(&pg_QPinchGesture_connect_destroyed2, "QPinchGesture_connect_destroyed2")
	reg(&pg_QPinchGesture_connect_objectNameChanged, "QPinchGesture_connect_objectNameChanged")
	reg(&pg_QPinchGesture__disconnect, "QPinchGesture__disconnect")
	reg(&pg_QPinchGesture__set_override, "QPinchGesture__set_override")
	reg(&pg_QPinchGesture_delete, "QPinchGesture_delete")
}

type QPinchGesture struct {
	h uintptr
	*QGesture
}

func bindQPinchGesture(h uintptr) *QPinchGesture {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPinchGesture_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPinchGesture{h: h, QGesture: &QGesture{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQPinchGesture(h uintptr) *QPinchGesture { return bindQPinchGesture(h) }

func (this *QPinchGesture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPinchGesture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPinchGesture(h unsafe.Pointer) *QPinchGesture { return bindQPinchGesture(uintptr(h)) }
func UnsafeNewQPinchGestureOwned(h unsafe.Pointer) *QPinchGesture { return newownedQPinchGesture(uintptr(h)) }

func (this *QPinchGesture) UnsafeRelease() {
	this.h = 0
	if this.QGesture != nil {
		this.QGesture.UnsafeRelease()
	}
}

func NewQPinchGesture() *QPinchGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_new, uintptr(__goid))
	__o := newownedQPinchGesture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPinchGesture_virtIdx, QPinchGesture_enable)
	return __o
}

func NewQPinchGesture2(parent *QObject) *QPinchGesture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPinchGesture(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPinchGesture_virtIdx, QPinchGesture_enable)
	return __o
}

func QPinchGesture_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPinchGesture) TotalChangeFlags() QPinchGesture__ChangeFlag {
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_totalChangeFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPinchGesture__ChangeFlag(__r)
}

func (this *QPinchGesture) SetTotalChangeFlags(value QPinchGesture__ChangeFlag) {
	purego.SyscallN(addr_QPinchGesture_setTotalChangeFlags, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) ChangeFlags() QPinchGesture__ChangeFlag {
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_changeFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPinchGesture__ChangeFlag(__r)
}

func (this *QPinchGesture) SetChangeFlags(value QPinchGesture__ChangeFlag) {
	purego.SyscallN(addr_QPinchGesture_setChangeFlags, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) StartCenterPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_startCenterPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPinchGesture) LastCenterPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_lastCenterPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPinchGesture) CenterPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_centerPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPinchGesture) SetStartCenterPoint(value *QPointF) {
	purego.SyscallN(addr_QPinchGesture_setStartCenterPoint, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QPinchGesture) SetLastCenterPoint(value *QPointF) {
	purego.SyscallN(addr_QPinchGesture_setLastCenterPoint, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QPinchGesture) SetCenterPoint(value *QPointF) {
	purego.SyscallN(addr_QPinchGesture_setCenterPoint, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QPinchGesture) TotalScaleFactor() float64 {
	__r := pg_QPinchGesture_totalScaleFactor(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPinchGesture) LastScaleFactor() float64 {
	__r := pg_QPinchGesture_lastScaleFactor(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPinchGesture) ScaleFactor() float64 {
	__r := pg_QPinchGesture_scaleFactor(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	pg_QPinchGesture_setTotalScaleFactor(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	pg_QPinchGesture_setLastScaleFactor(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) SetScaleFactor(value float64) {
	pg_QPinchGesture_setScaleFactor(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) TotalRotationAngle() float64 {
	__r := pg_QPinchGesture_totalRotationAngle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPinchGesture) LastRotationAngle() float64 {
	__r := pg_QPinchGesture_lastRotationAngle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPinchGesture) RotationAngle() float64 {
	__r := pg_QPinchGesture_rotationAngle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	pg_QPinchGesture_setTotalRotationAngle(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	pg_QPinchGesture_setLastRotationAngle(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QPinchGesture) SetRotationAngle(value float64) {
	pg_QPinchGesture_setRotationAngle(this.h, value)
	runtime.KeepAlive(this)
}

func QPinchGesture_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPinchGesture_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPinchGesture) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPinchGesture) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPinchGesture) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPinchGesture) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPinchGesture_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPinchGesture_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPinchGesture_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPinchGesture_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPinchGesture) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPinchGesture_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPinchGesture) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPinchGesture) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPinchGesture_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPinchGesture) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QPinchGesture_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPinchGesture) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPinchGesture_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPinchGesture) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPinchGesture_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPinchGesture) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPinchGesture_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPinchGesture) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPinchGesture_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPinchGesture) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPinchGesture_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPinchGesture) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPinchGesture_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPinchGesture) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPinchGesture_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPinchGesture) Disconnect(handle int64) {
	pg_QPinchGesture__disconnect(handle)
	unregisterCallback(handle)
}

func QPinchGesture_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QPinchGesture_enable(h uintptr, idx int32) { pg_QPinchGesture__set_override(h, idx) }

func (this *QPinchGesture) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPinchGesture) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QPinchGesture) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPinchGesture) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QPinchGesture) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPinchGesture) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPinchGesture) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPinchGesture) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPinchGesture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPinchGesture_delete(this.h)
	}
	this.UnsafeRelease()
}
