package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsScale_chain uintptr
	addr_QGraphicsScale_new uintptr
	addr_QGraphicsScale_new2 uintptr
	addr_QGraphicsScale_tr uintptr
	addr_QGraphicsScale_origin uintptr
	addr_QGraphicsScale_setOrigin uintptr
	pg_QGraphicsScale_xScale func(uintptr) float64
	pg_QGraphicsScale_setXScale func(uintptr, float64)
	pg_QGraphicsScale_yScale func(uintptr) float64
	pg_QGraphicsScale_setYScale func(uintptr, float64)
	pg_QGraphicsScale_zScale func(uintptr) float64
	pg_QGraphicsScale_setZScale func(uintptr, float64)
	addr_QGraphicsScale_originChanged uintptr
	addr_QGraphicsScale_xScaleChanged uintptr
	addr_QGraphicsScale_yScaleChanged uintptr
	addr_QGraphicsScale_zScaleChanged uintptr
	addr_QGraphicsScale_scaleChanged uintptr
	addr_QGraphicsScale_update uintptr
	addr_QGraphicsScale_connect uintptr
	addr_QGraphicsScale_connect2 uintptr
	addr_QGraphicsScale_sender uintptr
	addr_QGraphicsScale_senderSignalIndex uintptr
	addr_QGraphicsScale_receivers uintptr
	addr_QGraphicsScale_isSignalConnected uintptr
	addr_QGraphicsScale_tr2 uintptr
	addr_QGraphicsScale_tr3 uintptr
	addr_QGraphicsScale_connect3 uintptr
	addr_QGraphicsScale_connect4 uintptr
	addr_QGraphicsScale_metaObject_default uintptr
	addr_QGraphicsScale_applyTo_default uintptr
	addr_QGraphicsScale_event_default uintptr
	addr_QGraphicsScale_eventFilter_default uintptr
	addr_QGraphicsScale_timerEvent_default uintptr
	addr_QGraphicsScale_childEvent_default uintptr
	addr_QGraphicsScale_customEvent_default uintptr
	addr_QGraphicsScale_connectNotify_default uintptr
	addr_QGraphicsScale_disconnectNotify_default uintptr
	pg_QGraphicsScale_connect_originChanged func(uintptr, int64)
	pg_QGraphicsScale_connect_xScaleChanged func(uintptr, int64)
	pg_QGraphicsScale_connect_yScaleChanged func(uintptr, int64)
	pg_QGraphicsScale_connect_zScaleChanged func(uintptr, int64)
	pg_QGraphicsScale_connect_scaleChanged func(uintptr, int64)
	pg_QGraphicsScale_connect_destroyed func(uintptr, int64)
	pg_QGraphicsScale_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsScale_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsScale__disconnect func(int64)
	pg_QGraphicsScale__set_override func(uintptr, int32)
	pg_QGraphicsScale_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsScale_chain, "QGraphicsScale_chain")
	symaddr(&addr_QGraphicsScale_new, "QGraphicsScale_new")
	symaddr(&addr_QGraphicsScale_new2, "QGraphicsScale_new2")
	symaddr(&addr_QGraphicsScale_tr, "QGraphicsScale_tr")
	symaddr(&addr_QGraphicsScale_origin, "QGraphicsScale_origin")
	symaddr(&addr_QGraphicsScale_setOrigin, "QGraphicsScale_setOrigin")
	reg(&pg_QGraphicsScale_xScale, "QGraphicsScale_xScale")
	reg(&pg_QGraphicsScale_setXScale, "QGraphicsScale_setXScale")
	reg(&pg_QGraphicsScale_yScale, "QGraphicsScale_yScale")
	reg(&pg_QGraphicsScale_setYScale, "QGraphicsScale_setYScale")
	reg(&pg_QGraphicsScale_zScale, "QGraphicsScale_zScale")
	reg(&pg_QGraphicsScale_setZScale, "QGraphicsScale_setZScale")
	symaddr(&addr_QGraphicsScale_originChanged, "QGraphicsScale_originChanged")
	symaddr(&addr_QGraphicsScale_xScaleChanged, "QGraphicsScale_xScaleChanged")
	symaddr(&addr_QGraphicsScale_yScaleChanged, "QGraphicsScale_yScaleChanged")
	symaddr(&addr_QGraphicsScale_zScaleChanged, "QGraphicsScale_zScaleChanged")
	symaddr(&addr_QGraphicsScale_scaleChanged, "QGraphicsScale_scaleChanged")
	symaddr(&addr_QGraphicsScale_update, "QGraphicsScale_update")
	symaddr(&addr_QGraphicsScale_connect, "QGraphicsScale_connect")
	symaddr(&addr_QGraphicsScale_connect2, "QGraphicsScale_connect2")
	symaddr(&addr_QGraphicsScale_sender, "QGraphicsScale_sender")
	symaddr(&addr_QGraphicsScale_senderSignalIndex, "QGraphicsScale_senderSignalIndex")
	symaddr(&addr_QGraphicsScale_receivers, "QGraphicsScale_receivers")
	symaddr(&addr_QGraphicsScale_isSignalConnected, "QGraphicsScale_isSignalConnected")
	symaddr(&addr_QGraphicsScale_tr2, "QGraphicsScale_tr2")
	symaddr(&addr_QGraphicsScale_tr3, "QGraphicsScale_tr3")
	symaddr(&addr_QGraphicsScale_connect3, "QGraphicsScale_connect3")
	symaddr(&addr_QGraphicsScale_connect4, "QGraphicsScale_connect4")
	symaddr(&addr_QGraphicsScale_metaObject_default, "QGraphicsScale_metaObject_default")
	symaddr(&addr_QGraphicsScale_applyTo_default, "QGraphicsScale_applyTo_default")
	symaddr(&addr_QGraphicsScale_event_default, "QGraphicsScale_event_default")
	symaddr(&addr_QGraphicsScale_eventFilter_default, "QGraphicsScale_eventFilter_default")
	symaddr(&addr_QGraphicsScale_timerEvent_default, "QGraphicsScale_timerEvent_default")
	symaddr(&addr_QGraphicsScale_childEvent_default, "QGraphicsScale_childEvent_default")
	symaddr(&addr_QGraphicsScale_customEvent_default, "QGraphicsScale_customEvent_default")
	symaddr(&addr_QGraphicsScale_connectNotify_default, "QGraphicsScale_connectNotify_default")
	symaddr(&addr_QGraphicsScale_disconnectNotify_default, "QGraphicsScale_disconnectNotify_default")
	reg(&pg_QGraphicsScale_connect_originChanged, "QGraphicsScale_connect_originChanged")
	reg(&pg_QGraphicsScale_connect_xScaleChanged, "QGraphicsScale_connect_xScaleChanged")
	reg(&pg_QGraphicsScale_connect_yScaleChanged, "QGraphicsScale_connect_yScaleChanged")
	reg(&pg_QGraphicsScale_connect_zScaleChanged, "QGraphicsScale_connect_zScaleChanged")
	reg(&pg_QGraphicsScale_connect_scaleChanged, "QGraphicsScale_connect_scaleChanged")
	reg(&pg_QGraphicsScale_connect_destroyed, "QGraphicsScale_connect_destroyed")
	reg(&pg_QGraphicsScale_connect_destroyed2, "QGraphicsScale_connect_destroyed2")
	reg(&pg_QGraphicsScale_connect_objectNameChanged, "QGraphicsScale_connect_objectNameChanged")
	reg(&pg_QGraphicsScale__disconnect, "QGraphicsScale__disconnect")
	reg(&pg_QGraphicsScale__set_override, "QGraphicsScale__set_override")
	reg(&pg_QGraphicsScale_delete, "QGraphicsScale_delete")
}

type QGraphicsScale struct {
	h uintptr
	*QGraphicsTransform
}

func bindQGraphicsScale(h uintptr) *QGraphicsScale {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsScale_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsScale{h: h, QGraphicsTransform: &QGraphicsTransform{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsScale(h uintptr) *QGraphicsScale { return bindQGraphicsScale(h) }

func (this *QGraphicsScale) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsScale) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsScale(h unsafe.Pointer) *QGraphicsScale { return bindQGraphicsScale(uintptr(h)) }
func UnsafeNewQGraphicsScaleOwned(h unsafe.Pointer) *QGraphicsScale { return newownedQGraphicsScale(uintptr(h)) }

func (this *QGraphicsScale) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsTransform != nil {
		this.QGraphicsTransform.UnsafeRelease()
	}
}

func NewQGraphicsScale() *QGraphicsScale {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_new, uintptr(__goid))
	__o := newownedQGraphicsScale(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScale_virtIdx, QGraphicsScale_enable)
	return __o
}

func NewQGraphicsScale2(parent *QObject) *QGraphicsScale {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsScale(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScale_virtIdx, QGraphicsScale_enable)
	return __o
}

func QGraphicsScale_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsScale) Origin() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_origin, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QGraphicsScale) SetOrigin(point *QVector3D) {
	purego.SyscallN(addr_QGraphicsScale_setOrigin, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
}

func (this *QGraphicsScale) XScale() float64 {
	__r := pg_QGraphicsScale_xScale(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsScale) SetXScale(arg1 float64) {
	pg_QGraphicsScale_setXScale(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) YScale() float64 {
	__r := pg_QGraphicsScale_yScale(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsScale) SetYScale(arg1 float64) {
	pg_QGraphicsScale_setYScale(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) ZScale() float64 {
	__r := pg_QGraphicsScale_zScale(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsScale) SetZScale(arg1 float64) {
	pg_QGraphicsScale_setZScale(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) OriginChanged() {
	purego.SyscallN(addr_QGraphicsScale_originChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) XScaleChanged() {
	purego.SyscallN(addr_QGraphicsScale_xScaleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) YScaleChanged() {
	purego.SyscallN(addr_QGraphicsScale_yScaleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) ZScaleChanged() {
	purego.SyscallN(addr_QGraphicsScale_zScaleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) ScaleChanged() {
	purego.SyscallN(addr_QGraphicsScale_scaleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScale) Update() {
	purego.SyscallN(addr_QGraphicsScale_update, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsScale_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsScale_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsScale) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsScale) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScale_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsScale) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScale_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsScale) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScale_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsScale_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsScale_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsScale_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsScale_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsScale) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScale_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsScale) ApplyToDefault(matrix *QMatrix4x4) {
	purego.SyscallN(addr_QGraphicsScale_applyTo_default, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsScale) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScale_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsScale) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScale_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsScale) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsScale_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScale) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsScale_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScale) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsScale_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScale) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsScale_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsScale) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsScale_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsScale) OnOriginChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScale_connect_originChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnXScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScale_connect_xScaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnYScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScale_connect_yScaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnZScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScale_connect_zScaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScale_connect_scaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScale_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsScale_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsScale_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScale) Disconnect(handle int64) {
	pg_QGraphicsScale__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsScale_virtIdx(slot int64) int32 {
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

func QGraphicsScale_enable(h uintptr, idx int32) { pg_QGraphicsScale__set_override(h, idx) }

func (this *QGraphicsScale) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsScale) OnApplyTo(fn func(*QMatrix4x4)) {
	registerVirt(this.UnsafeGoID(), 1580758529722784438, func(matrix uintptr) { fn(UnsafeNewQMatrix4x4(unsafe.Pointer(matrix))) }) // slot "applyTo"
}

func (this *QGraphicsScale) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsScale) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsScale) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsScale) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsScale) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsScale) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsScale) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsScale) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsScale_delete(this.h)
	}
	this.UnsafeRelease()
}
