package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsRotation_chain uintptr
	addr_QGraphicsRotation_new uintptr
	addr_QGraphicsRotation_new2 uintptr
	addr_QGraphicsRotation_tr uintptr
	addr_QGraphicsRotation_origin uintptr
	addr_QGraphicsRotation_setOrigin uintptr
	pg_QGraphicsRotation_angle func(uintptr) float64
	pg_QGraphicsRotation_setAngle func(uintptr, float64)
	addr_QGraphicsRotation_axis uintptr
	addr_QGraphicsRotation_setAxis uintptr
	addr_QGraphicsRotation_setAxis2 uintptr
	addr_QGraphicsRotation_originChanged uintptr
	addr_QGraphicsRotation_angleChanged uintptr
	addr_QGraphicsRotation_axisChanged uintptr
	addr_QGraphicsRotation_update uintptr
	addr_QGraphicsRotation_connect uintptr
	addr_QGraphicsRotation_connect2 uintptr
	addr_QGraphicsRotation_sender uintptr
	addr_QGraphicsRotation_senderSignalIndex uintptr
	addr_QGraphicsRotation_receivers uintptr
	addr_QGraphicsRotation_isSignalConnected uintptr
	addr_QGraphicsRotation_tr2 uintptr
	addr_QGraphicsRotation_tr3 uintptr
	addr_QGraphicsRotation_connect3 uintptr
	addr_QGraphicsRotation_connect4 uintptr
	addr_QGraphicsRotation_metaObject_default uintptr
	addr_QGraphicsRotation_applyTo_default uintptr
	addr_QGraphicsRotation_event_default uintptr
	addr_QGraphicsRotation_eventFilter_default uintptr
	addr_QGraphicsRotation_timerEvent_default uintptr
	addr_QGraphicsRotation_childEvent_default uintptr
	addr_QGraphicsRotation_customEvent_default uintptr
	addr_QGraphicsRotation_connectNotify_default uintptr
	addr_QGraphicsRotation_disconnectNotify_default uintptr
	pg_QGraphicsRotation_connect_originChanged func(uintptr, int64)
	pg_QGraphicsRotation_connect_angleChanged func(uintptr, int64)
	pg_QGraphicsRotation_connect_axisChanged func(uintptr, int64)
	pg_QGraphicsRotation_connect_destroyed func(uintptr, int64)
	pg_QGraphicsRotation_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsRotation_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsRotation__disconnect func(int64)
	pg_QGraphicsRotation__set_override func(uintptr, int32)
	pg_QGraphicsRotation_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsRotation_chain, "QGraphicsRotation_chain")
	symaddr(&addr_QGraphicsRotation_new, "QGraphicsRotation_new")
	symaddr(&addr_QGraphicsRotation_new2, "QGraphicsRotation_new2")
	symaddr(&addr_QGraphicsRotation_tr, "QGraphicsRotation_tr")
	symaddr(&addr_QGraphicsRotation_origin, "QGraphicsRotation_origin")
	symaddr(&addr_QGraphicsRotation_setOrigin, "QGraphicsRotation_setOrigin")
	reg(&pg_QGraphicsRotation_angle, "QGraphicsRotation_angle")
	reg(&pg_QGraphicsRotation_setAngle, "QGraphicsRotation_setAngle")
	symaddr(&addr_QGraphicsRotation_axis, "QGraphicsRotation_axis")
	symaddr(&addr_QGraphicsRotation_setAxis, "QGraphicsRotation_setAxis")
	symaddr(&addr_QGraphicsRotation_setAxis2, "QGraphicsRotation_setAxis2")
	symaddr(&addr_QGraphicsRotation_originChanged, "QGraphicsRotation_originChanged")
	symaddr(&addr_QGraphicsRotation_angleChanged, "QGraphicsRotation_angleChanged")
	symaddr(&addr_QGraphicsRotation_axisChanged, "QGraphicsRotation_axisChanged")
	symaddr(&addr_QGraphicsRotation_update, "QGraphicsRotation_update")
	symaddr(&addr_QGraphicsRotation_connect, "QGraphicsRotation_connect")
	symaddr(&addr_QGraphicsRotation_connect2, "QGraphicsRotation_connect2")
	symaddr(&addr_QGraphicsRotation_sender, "QGraphicsRotation_sender")
	symaddr(&addr_QGraphicsRotation_senderSignalIndex, "QGraphicsRotation_senderSignalIndex")
	symaddr(&addr_QGraphicsRotation_receivers, "QGraphicsRotation_receivers")
	symaddr(&addr_QGraphicsRotation_isSignalConnected, "QGraphicsRotation_isSignalConnected")
	symaddr(&addr_QGraphicsRotation_tr2, "QGraphicsRotation_tr2")
	symaddr(&addr_QGraphicsRotation_tr3, "QGraphicsRotation_tr3")
	symaddr(&addr_QGraphicsRotation_connect3, "QGraphicsRotation_connect3")
	symaddr(&addr_QGraphicsRotation_connect4, "QGraphicsRotation_connect4")
	symaddr(&addr_QGraphicsRotation_metaObject_default, "QGraphicsRotation_metaObject_default")
	symaddr(&addr_QGraphicsRotation_applyTo_default, "QGraphicsRotation_applyTo_default")
	symaddr(&addr_QGraphicsRotation_event_default, "QGraphicsRotation_event_default")
	symaddr(&addr_QGraphicsRotation_eventFilter_default, "QGraphicsRotation_eventFilter_default")
	symaddr(&addr_QGraphicsRotation_timerEvent_default, "QGraphicsRotation_timerEvent_default")
	symaddr(&addr_QGraphicsRotation_childEvent_default, "QGraphicsRotation_childEvent_default")
	symaddr(&addr_QGraphicsRotation_customEvent_default, "QGraphicsRotation_customEvent_default")
	symaddr(&addr_QGraphicsRotation_connectNotify_default, "QGraphicsRotation_connectNotify_default")
	symaddr(&addr_QGraphicsRotation_disconnectNotify_default, "QGraphicsRotation_disconnectNotify_default")
	reg(&pg_QGraphicsRotation_connect_originChanged, "QGraphicsRotation_connect_originChanged")
	reg(&pg_QGraphicsRotation_connect_angleChanged, "QGraphicsRotation_connect_angleChanged")
	reg(&pg_QGraphicsRotation_connect_axisChanged, "QGraphicsRotation_connect_axisChanged")
	reg(&pg_QGraphicsRotation_connect_destroyed, "QGraphicsRotation_connect_destroyed")
	reg(&pg_QGraphicsRotation_connect_destroyed2, "QGraphicsRotation_connect_destroyed2")
	reg(&pg_QGraphicsRotation_connect_objectNameChanged, "QGraphicsRotation_connect_objectNameChanged")
	reg(&pg_QGraphicsRotation__disconnect, "QGraphicsRotation__disconnect")
	reg(&pg_QGraphicsRotation__set_override, "QGraphicsRotation__set_override")
	reg(&pg_QGraphicsRotation_delete, "QGraphicsRotation_delete")
}

type QGraphicsRotation struct {
	h uintptr
	*QGraphicsTransform
}

func bindQGraphicsRotation(h uintptr) *QGraphicsRotation {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsRotation_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsRotation{h: h, QGraphicsTransform: &QGraphicsTransform{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsRotation(h uintptr) *QGraphicsRotation { return bindQGraphicsRotation(h) }

func (this *QGraphicsRotation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsRotation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsRotation(h unsafe.Pointer) *QGraphicsRotation { return bindQGraphicsRotation(uintptr(h)) }
func UnsafeNewQGraphicsRotationOwned(h unsafe.Pointer) *QGraphicsRotation { return newownedQGraphicsRotation(uintptr(h)) }

func (this *QGraphicsRotation) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsTransform != nil {
		this.QGraphicsTransform.UnsafeRelease()
	}
}

func NewQGraphicsRotation() *QGraphicsRotation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_new, uintptr(__goid))
	__o := newownedQGraphicsRotation(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRotation_virtIdx, QGraphicsRotation_enable)
	return __o
}

func NewQGraphicsRotation2(parent *QObject) *QGraphicsRotation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsRotation(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRotation_virtIdx, QGraphicsRotation_enable)
	return __o
}

func QGraphicsRotation_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsRotation) Origin() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_origin, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QGraphicsRotation) SetOrigin(point *QVector3D) {
	purego.SyscallN(addr_QGraphicsRotation_setOrigin, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
}

func (this *QGraphicsRotation) Angle() float64 {
	__r := pg_QGraphicsRotation_angle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsRotation) SetAngle(arg1 float64) {
	pg_QGraphicsRotation_setAngle(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRotation) Axis() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_axis, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QGraphicsRotation) SetAxis(axis *QVector3D) {
	purego.SyscallN(addr_QGraphicsRotation_setAxis, this.h, uintptr(axis.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(axis)
}

func (this *QGraphicsRotation) SetAxisWithAxis(axis Axis) {
	purego.SyscallN(addr_QGraphicsRotation_setAxis2, this.h, uintptr(axis))
	runtime.KeepAlive(this)
}

func (this *QGraphicsRotation) OriginChanged() {
	purego.SyscallN(addr_QGraphicsRotation_originChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRotation) AngleChanged() {
	purego.SyscallN(addr_QGraphicsRotation_angleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRotation) AxisChanged() {
	purego.SyscallN(addr_QGraphicsRotation_axisChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRotation) Update() {
	purego.SyscallN(addr_QGraphicsRotation_update, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsRotation_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsRotation_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsRotation) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsRotation) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRotation_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsRotation) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRotation_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsRotation) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRotation_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsRotation_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsRotation_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsRotation_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsRotation_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsRotation) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRotation_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsRotation) ApplyToDefault(matrix *QMatrix4x4) {
	purego.SyscallN(addr_QGraphicsRotation_applyTo_default, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsRotation) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRotation_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsRotation) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRotation_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsRotation) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsRotation_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRotation) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsRotation_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRotation) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsRotation_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRotation) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsRotation_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsRotation) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsRotation_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsRotation) OnOriginChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsRotation_connect_originChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsRotation) OnAngleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsRotation_connect_angleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsRotation) OnAxisChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsRotation_connect_axisChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsRotation) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsRotation_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsRotation) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsRotation_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsRotation) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsRotation_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsRotation) Disconnect(handle int64) {
	pg_QGraphicsRotation__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsRotation_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1580758529722784438:
		return 1
	case 1237236456312824991:
		return 2
	case 5217623514630858297:
		return 3
	case -5409927581812374498:
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

func QGraphicsRotation_enable(h uintptr, idx int32) { pg_QGraphicsRotation__set_override(h, idx) }

func (this *QGraphicsRotation) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsRotation) OnApplyTo(fn func(*QMatrix4x4)) {
	registerVirt(this.UnsafeGoID(), 1580758529722784438, func(matrix uintptr) { fn(UnsafeNewQMatrix4x4(unsafe.Pointer(matrix))) }) // slot "applyTo"
}

func (this *QGraphicsRotation) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsRotation) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsRotation) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsRotation) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsRotation) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsRotation) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsRotation) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsRotation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsRotation_delete(this.h)
	}
	this.UnsafeRelease()
}
