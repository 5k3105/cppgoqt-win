package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsTransform_base func(uintptr) uintptr
	addr_QGraphicsTransform_tr uintptr
	addr_QGraphicsTransform_applyTo uintptr
	addr_QGraphicsTransform_update uintptr
	addr_QGraphicsTransform_connect uintptr
	addr_QGraphicsTransform_connect2 uintptr
	addr_QGraphicsTransform_sender uintptr
	addr_QGraphicsTransform_senderSignalIndex uintptr
	addr_QGraphicsTransform_receivers uintptr
	addr_QGraphicsTransform_isSignalConnected uintptr
	addr_QGraphicsTransform_tr2 uintptr
	addr_QGraphicsTransform_tr3 uintptr
	addr_QGraphicsTransform_connect3 uintptr
	addr_QGraphicsTransform_connect4 uintptr
	addr_QGraphicsTransform_metaObject_default uintptr
	addr_QGraphicsTransform_event_default uintptr
	addr_QGraphicsTransform_eventFilter_default uintptr
	addr_QGraphicsTransform_timerEvent_default uintptr
	addr_QGraphicsTransform_childEvent_default uintptr
	addr_QGraphicsTransform_customEvent_default uintptr
	addr_QGraphicsTransform_connectNotify_default uintptr
	addr_QGraphicsTransform_disconnectNotify_default uintptr
	pg_QGraphicsTransform_connect_destroyed func(uintptr, int64)
	pg_QGraphicsTransform_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsTransform_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsTransform__disconnect func(int64)
	pg_QGraphicsTransform__set_override func(uintptr, int32)
	pg_QGraphicsTransform__Go_new func(int64, uintptr) uintptr
	pg_QGraphicsTransform_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsTransform_base, "QGraphicsTransform_base")
	symaddr(&addr_QGraphicsTransform_tr, "QGraphicsTransform_tr")
	symaddr(&addr_QGraphicsTransform_applyTo, "QGraphicsTransform_applyTo")
	symaddr(&addr_QGraphicsTransform_update, "QGraphicsTransform_update")
	symaddr(&addr_QGraphicsTransform_connect, "QGraphicsTransform_connect")
	symaddr(&addr_QGraphicsTransform_connect2, "QGraphicsTransform_connect2")
	symaddr(&addr_QGraphicsTransform_sender, "QGraphicsTransform_sender")
	symaddr(&addr_QGraphicsTransform_senderSignalIndex, "QGraphicsTransform_senderSignalIndex")
	symaddr(&addr_QGraphicsTransform_receivers, "QGraphicsTransform_receivers")
	symaddr(&addr_QGraphicsTransform_isSignalConnected, "QGraphicsTransform_isSignalConnected")
	symaddr(&addr_QGraphicsTransform_tr2, "QGraphicsTransform_tr2")
	symaddr(&addr_QGraphicsTransform_tr3, "QGraphicsTransform_tr3")
	symaddr(&addr_QGraphicsTransform_connect3, "QGraphicsTransform_connect3")
	symaddr(&addr_QGraphicsTransform_connect4, "QGraphicsTransform_connect4")
	symaddr(&addr_QGraphicsTransform_metaObject_default, "QGraphicsTransform_metaObject_default")
	symaddr(&addr_QGraphicsTransform_event_default, "QGraphicsTransform_event_default")
	symaddr(&addr_QGraphicsTransform_eventFilter_default, "QGraphicsTransform_eventFilter_default")
	symaddr(&addr_QGraphicsTransform_timerEvent_default, "QGraphicsTransform_timerEvent_default")
	symaddr(&addr_QGraphicsTransform_childEvent_default, "QGraphicsTransform_childEvent_default")
	symaddr(&addr_QGraphicsTransform_customEvent_default, "QGraphicsTransform_customEvent_default")
	symaddr(&addr_QGraphicsTransform_connectNotify_default, "QGraphicsTransform_connectNotify_default")
	symaddr(&addr_QGraphicsTransform_disconnectNotify_default, "QGraphicsTransform_disconnectNotify_default")
	reg(&pg_QGraphicsTransform_connect_destroyed, "QGraphicsTransform_connect_destroyed")
	reg(&pg_QGraphicsTransform_connect_destroyed2, "QGraphicsTransform_connect_destroyed2")
	reg(&pg_QGraphicsTransform_connect_objectNameChanged, "QGraphicsTransform_connect_objectNameChanged")
	reg(&pg_QGraphicsTransform__disconnect, "QGraphicsTransform__disconnect")
	reg(&pg_QGraphicsTransform__set_override, "QGraphicsTransform__set_override")
	reg(&pg_QGraphicsTransform__Go_new, "QGraphicsTransform__Go_new")
	reg(&pg_QGraphicsTransform_delete, "QGraphicsTransform_delete")
}

type QGraphicsTransform struct {
	h uintptr
	*QObject
}

func bindQGraphicsTransform(h uintptr) *QGraphicsTransform {
	if h == 0 {
		return nil
	}
	return &QGraphicsTransform{h: h, QObject: bindQObject(pg_QGraphicsTransform_base(h))}
}

func newownedQGraphicsTransform(h uintptr) *QGraphicsTransform { return bindQGraphicsTransform(h) }

func (this *QGraphicsTransform) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsTransform) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsTransform(h unsafe.Pointer) *QGraphicsTransform { return bindQGraphicsTransform(uintptr(h)) }
func UnsafeNewQGraphicsTransformOwned(h unsafe.Pointer) *QGraphicsTransform { return newownedQGraphicsTransform(uintptr(h)) }

func (this *QGraphicsTransform) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QGraphicsTransform_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsTransform) ApplyTo(matrix *QMatrix4x4) {
	purego.SyscallN(addr_QGraphicsTransform_applyTo, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsTransform) Update() {
	purego.SyscallN(addr_QGraphicsTransform_update, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsTransform_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsTransform_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsTransform) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsTransform) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTransform_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsTransform) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTransform_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsTransform) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTransform_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsTransform_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsTransform_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsTransform_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsTransform_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsTransform) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTransform_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsTransform) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTransform_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsTransform) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTransform_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsTransform) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsTransform_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTransform) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsTransform_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTransform) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsTransform_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTransform) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsTransform_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsTransform) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsTransform_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsTransform) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTransform_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTransform) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsTransform_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTransform) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsTransform_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTransform) Disconnect(handle int64) {
	pg_QGraphicsTransform__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsTransform_virtIdx(slot int64) int32 {
	switch slot {
	case 1580758529722784438:
		return 0
	case 1607084412591819503:
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

func QGraphicsTransform_enable(h uintptr, idx int32) { pg_QGraphicsTransform__set_override(h, idx) }

func NewQGraphicsTransform() *QGraphicsTransform {
	id := nextVirtID()
	o := newownedQGraphicsTransform(pg_QGraphicsTransform__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsTransform_virtIdx, QGraphicsTransform_enable)
	return o
}

func NewQGraphicsTransform2(parent *QObject) *QGraphicsTransform {
	id := nextVirtID()
	o := newownedQGraphicsTransform(pg_QGraphicsTransform__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsTransform_virtIdx, QGraphicsTransform_enable)
	return o
}

func (this *QGraphicsTransform) OnApplyTo(fn func(*QMatrix4x4)) {
	registerVirt(this.UnsafeGoID(), 1580758529722784438, func(matrix uintptr) { fn(UnsafeNewQMatrix4x4(unsafe.Pointer(matrix))) }) // slot "applyTo"
}

func (this *QGraphicsTransform) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsTransform) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsTransform) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsTransform) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsTransform) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsTransform) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsTransform) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsTransform) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsTransform) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsTransform_delete(this.h)
	}
	this.UnsafeRelease()
}
