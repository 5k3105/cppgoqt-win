package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QObjectCleanupHandler_base func(uintptr) uintptr
	addr_QObjectCleanupHandler_new uintptr
	addr_QObjectCleanupHandler_tr uintptr
	addr_QObjectCleanupHandler_add uintptr
	addr_QObjectCleanupHandler_remove uintptr
	addr_QObjectCleanupHandler_isEmpty uintptr
	addr_QObjectCleanupHandler_clear uintptr
	addr_QObjectCleanupHandler_connect uintptr
	addr_QObjectCleanupHandler_connect2 uintptr
	addr_QObjectCleanupHandler_sender uintptr
	addr_QObjectCleanupHandler_senderSignalIndex uintptr
	addr_QObjectCleanupHandler_receivers uintptr
	addr_QObjectCleanupHandler_isSignalConnected uintptr
	addr_QObjectCleanupHandler_tr2 uintptr
	addr_QObjectCleanupHandler_tr3 uintptr
	addr_QObjectCleanupHandler_connect3 uintptr
	addr_QObjectCleanupHandler_connect4 uintptr
	addr_QObjectCleanupHandler_metaObject_default uintptr
	addr_QObjectCleanupHandler_event_default uintptr
	addr_QObjectCleanupHandler_eventFilter_default uintptr
	addr_QObjectCleanupHandler_timerEvent_default uintptr
	addr_QObjectCleanupHandler_childEvent_default uintptr
	addr_QObjectCleanupHandler_customEvent_default uintptr
	addr_QObjectCleanupHandler_connectNotify_default uintptr
	addr_QObjectCleanupHandler_disconnectNotify_default uintptr
	pg_QObjectCleanupHandler_connect_destroyed func(uintptr, int64)
	pg_QObjectCleanupHandler_connect_destroyed2 func(uintptr, int64)
	pg_QObjectCleanupHandler_connect_objectNameChanged func(uintptr, int64)
	pg_QObjectCleanupHandler__disconnect func(int64)
	pg_QObjectCleanupHandler__set_override func(uintptr, int32)
	pg_QObjectCleanupHandler_delete func(uintptr)
)

func init() {
	reg(&pg_QObjectCleanupHandler_base, "QObjectCleanupHandler_base")
	symaddr(&addr_QObjectCleanupHandler_new, "QObjectCleanupHandler_new")
	symaddr(&addr_QObjectCleanupHandler_tr, "QObjectCleanupHandler_tr")
	symaddr(&addr_QObjectCleanupHandler_add, "QObjectCleanupHandler_add")
	symaddr(&addr_QObjectCleanupHandler_remove, "QObjectCleanupHandler_remove")
	symaddr(&addr_QObjectCleanupHandler_isEmpty, "QObjectCleanupHandler_isEmpty")
	symaddr(&addr_QObjectCleanupHandler_clear, "QObjectCleanupHandler_clear")
	symaddr(&addr_QObjectCleanupHandler_connect, "QObjectCleanupHandler_connect")
	symaddr(&addr_QObjectCleanupHandler_connect2, "QObjectCleanupHandler_connect2")
	symaddr(&addr_QObjectCleanupHandler_sender, "QObjectCleanupHandler_sender")
	symaddr(&addr_QObjectCleanupHandler_senderSignalIndex, "QObjectCleanupHandler_senderSignalIndex")
	symaddr(&addr_QObjectCleanupHandler_receivers, "QObjectCleanupHandler_receivers")
	symaddr(&addr_QObjectCleanupHandler_isSignalConnected, "QObjectCleanupHandler_isSignalConnected")
	symaddr(&addr_QObjectCleanupHandler_tr2, "QObjectCleanupHandler_tr2")
	symaddr(&addr_QObjectCleanupHandler_tr3, "QObjectCleanupHandler_tr3")
	symaddr(&addr_QObjectCleanupHandler_connect3, "QObjectCleanupHandler_connect3")
	symaddr(&addr_QObjectCleanupHandler_connect4, "QObjectCleanupHandler_connect4")
	symaddr(&addr_QObjectCleanupHandler_metaObject_default, "QObjectCleanupHandler_metaObject_default")
	symaddr(&addr_QObjectCleanupHandler_event_default, "QObjectCleanupHandler_event_default")
	symaddr(&addr_QObjectCleanupHandler_eventFilter_default, "QObjectCleanupHandler_eventFilter_default")
	symaddr(&addr_QObjectCleanupHandler_timerEvent_default, "QObjectCleanupHandler_timerEvent_default")
	symaddr(&addr_QObjectCleanupHandler_childEvent_default, "QObjectCleanupHandler_childEvent_default")
	symaddr(&addr_QObjectCleanupHandler_customEvent_default, "QObjectCleanupHandler_customEvent_default")
	symaddr(&addr_QObjectCleanupHandler_connectNotify_default, "QObjectCleanupHandler_connectNotify_default")
	symaddr(&addr_QObjectCleanupHandler_disconnectNotify_default, "QObjectCleanupHandler_disconnectNotify_default")
	reg(&pg_QObjectCleanupHandler_connect_destroyed, "QObjectCleanupHandler_connect_destroyed")
	reg(&pg_QObjectCleanupHandler_connect_destroyed2, "QObjectCleanupHandler_connect_destroyed2")
	reg(&pg_QObjectCleanupHandler_connect_objectNameChanged, "QObjectCleanupHandler_connect_objectNameChanged")
	reg(&pg_QObjectCleanupHandler__disconnect, "QObjectCleanupHandler__disconnect")
	reg(&pg_QObjectCleanupHandler__set_override, "QObjectCleanupHandler__set_override")
	reg(&pg_QObjectCleanupHandler_delete, "QObjectCleanupHandler_delete")
}

type QObjectCleanupHandler struct {
	h uintptr
	*QObject
}

func bindQObjectCleanupHandler(h uintptr) *QObjectCleanupHandler {
	if h == 0 {
		return nil
	}
	return &QObjectCleanupHandler{h: h, QObject: bindQObject(pg_QObjectCleanupHandler_base(h))}
}

func newownedQObjectCleanupHandler(h uintptr) *QObjectCleanupHandler { return bindQObjectCleanupHandler(h) }

func (this *QObjectCleanupHandler) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QObjectCleanupHandler) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQObjectCleanupHandler(h unsafe.Pointer) *QObjectCleanupHandler { return bindQObjectCleanupHandler(uintptr(h)) }
func UnsafeNewQObjectCleanupHandlerOwned(h unsafe.Pointer) *QObjectCleanupHandler { return newownedQObjectCleanupHandler(uintptr(h)) }

func (this *QObjectCleanupHandler) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_new, uintptr(__goid))
	__o := newownedQObjectCleanupHandler(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QObjectCleanupHandler_virtIdx, QObjectCleanupHandler_enable)
	return __o
}

func QObjectCleanupHandler_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QObjectCleanupHandler) Add(object *QObject) *QObject {
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_add, this.h, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	return bindQObject(__r)
}

func (this *QObjectCleanupHandler) Remove(object *QObject) {
	purego.SyscallN(addr_QObjectCleanupHandler_remove, this.h, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
}

func (this *QObjectCleanupHandler) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObjectCleanupHandler) Clear() {
	purego.SyscallN(addr_QObjectCleanupHandler_clear, this.h)
	runtime.KeepAlive(this)
}

func QObjectCleanupHandler_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QObjectCleanupHandler_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QObjectCleanupHandler) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QObjectCleanupHandler) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QObjectCleanupHandler) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QObjectCleanupHandler) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QObjectCleanupHandler_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QObjectCleanupHandler_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QObjectCleanupHandler_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QObjectCleanupHandler_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QObjectCleanupHandler) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QObjectCleanupHandler) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QObjectCleanupHandler) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QObjectCleanupHandler_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QObjectCleanupHandler) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QObjectCleanupHandler_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QObjectCleanupHandler) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QObjectCleanupHandler_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QObjectCleanupHandler) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QObjectCleanupHandler_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QObjectCleanupHandler) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QObjectCleanupHandler_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QObjectCleanupHandler) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QObjectCleanupHandler_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QObjectCleanupHandler) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QObjectCleanupHandler_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QObjectCleanupHandler) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QObjectCleanupHandler_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QObjectCleanupHandler) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QObjectCleanupHandler_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QObjectCleanupHandler) Disconnect(handle int64) {
	pg_QObjectCleanupHandler__disconnect(handle)
	unregisterCallback(handle)
}

func QObjectCleanupHandler_virtIdx(slot int64) int32 {
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

func QObjectCleanupHandler_enable(h uintptr, idx int32) { pg_QObjectCleanupHandler__set_override(h, idx) }

func (this *QObjectCleanupHandler) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QObjectCleanupHandler) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QObjectCleanupHandler) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QObjectCleanupHandler) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QObjectCleanupHandler) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QObjectCleanupHandler) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QObjectCleanupHandler) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QObjectCleanupHandler) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QObjectCleanupHandler) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QObjectCleanupHandler_delete(this.h)
	}
	this.UnsafeRelease()
}
