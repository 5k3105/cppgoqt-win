package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QWinEventNotifier_base func(uintptr) uintptr
	addr_QWinEventNotifier_new uintptr
	addr_QWinEventNotifier_new2 uintptr
	addr_QWinEventNotifier_tr uintptr
	addr_QWinEventNotifier_isEnabled uintptr
	addr_QWinEventNotifier_setEnabled uintptr
	addr_QWinEventNotifier_connect uintptr
	addr_QWinEventNotifier_connect2 uintptr
	addr_QWinEventNotifier_sender uintptr
	addr_QWinEventNotifier_senderSignalIndex uintptr
	addr_QWinEventNotifier_receivers uintptr
	addr_QWinEventNotifier_isSignalConnected uintptr
	addr_QWinEventNotifier_tr2 uintptr
	addr_QWinEventNotifier_tr3 uintptr
	addr_QWinEventNotifier_connect3 uintptr
	addr_QWinEventNotifier_connect4 uintptr
	addr_QWinEventNotifier_metaObject_default uintptr
	addr_QWinEventNotifier_event_default uintptr
	addr_QWinEventNotifier_eventFilter_default uintptr
	addr_QWinEventNotifier_timerEvent_default uintptr
	addr_QWinEventNotifier_childEvent_default uintptr
	addr_QWinEventNotifier_customEvent_default uintptr
	addr_QWinEventNotifier_connectNotify_default uintptr
	addr_QWinEventNotifier_disconnectNotify_default uintptr
	pg_QWinEventNotifier_connect_destroyed func(uintptr, int64)
	pg_QWinEventNotifier_connect_destroyed2 func(uintptr, int64)
	pg_QWinEventNotifier_connect_objectNameChanged func(uintptr, int64)
	pg_QWinEventNotifier__disconnect func(int64)
	pg_QWinEventNotifier__set_override func(uintptr, int32)
	pg_QWinEventNotifier_delete func(uintptr)
)

func init() {
	reg(&pg_QWinEventNotifier_base, "QWinEventNotifier_base")
	symaddr(&addr_QWinEventNotifier_new, "QWinEventNotifier_new")
	symaddr(&addr_QWinEventNotifier_new2, "QWinEventNotifier_new2")
	symaddr(&addr_QWinEventNotifier_tr, "QWinEventNotifier_tr")
	symaddr(&addr_QWinEventNotifier_isEnabled, "QWinEventNotifier_isEnabled")
	symaddr(&addr_QWinEventNotifier_setEnabled, "QWinEventNotifier_setEnabled")
	symaddr(&addr_QWinEventNotifier_connect, "QWinEventNotifier_connect")
	symaddr(&addr_QWinEventNotifier_connect2, "QWinEventNotifier_connect2")
	symaddr(&addr_QWinEventNotifier_sender, "QWinEventNotifier_sender")
	symaddr(&addr_QWinEventNotifier_senderSignalIndex, "QWinEventNotifier_senderSignalIndex")
	symaddr(&addr_QWinEventNotifier_receivers, "QWinEventNotifier_receivers")
	symaddr(&addr_QWinEventNotifier_isSignalConnected, "QWinEventNotifier_isSignalConnected")
	symaddr(&addr_QWinEventNotifier_tr2, "QWinEventNotifier_tr2")
	symaddr(&addr_QWinEventNotifier_tr3, "QWinEventNotifier_tr3")
	symaddr(&addr_QWinEventNotifier_connect3, "QWinEventNotifier_connect3")
	symaddr(&addr_QWinEventNotifier_connect4, "QWinEventNotifier_connect4")
	symaddr(&addr_QWinEventNotifier_metaObject_default, "QWinEventNotifier_metaObject_default")
	symaddr(&addr_QWinEventNotifier_event_default, "QWinEventNotifier_event_default")
	symaddr(&addr_QWinEventNotifier_eventFilter_default, "QWinEventNotifier_eventFilter_default")
	symaddr(&addr_QWinEventNotifier_timerEvent_default, "QWinEventNotifier_timerEvent_default")
	symaddr(&addr_QWinEventNotifier_childEvent_default, "QWinEventNotifier_childEvent_default")
	symaddr(&addr_QWinEventNotifier_customEvent_default, "QWinEventNotifier_customEvent_default")
	symaddr(&addr_QWinEventNotifier_connectNotify_default, "QWinEventNotifier_connectNotify_default")
	symaddr(&addr_QWinEventNotifier_disconnectNotify_default, "QWinEventNotifier_disconnectNotify_default")
	reg(&pg_QWinEventNotifier_connect_destroyed, "QWinEventNotifier_connect_destroyed")
	reg(&pg_QWinEventNotifier_connect_destroyed2, "QWinEventNotifier_connect_destroyed2")
	reg(&pg_QWinEventNotifier_connect_objectNameChanged, "QWinEventNotifier_connect_objectNameChanged")
	reg(&pg_QWinEventNotifier__disconnect, "QWinEventNotifier__disconnect")
	reg(&pg_QWinEventNotifier__set_override, "QWinEventNotifier__set_override")
	reg(&pg_QWinEventNotifier_delete, "QWinEventNotifier_delete")
}

type QWinEventNotifier struct {
	h uintptr
	*QObject
}

func bindQWinEventNotifier(h uintptr) *QWinEventNotifier {
	if h == 0 {
		return nil
	}
	return &QWinEventNotifier{h: h, QObject: bindQObject(pg_QWinEventNotifier_base(h))}
}

func newownedQWinEventNotifier(h uintptr) *QWinEventNotifier { return bindQWinEventNotifier(h) }

func (this *QWinEventNotifier) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWinEventNotifier) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWinEventNotifier(h unsafe.Pointer) *QWinEventNotifier { return bindQWinEventNotifier(uintptr(h)) }
func UnsafeNewQWinEventNotifierOwned(h unsafe.Pointer) *QWinEventNotifier { return newownedQWinEventNotifier(uintptr(h)) }

func (this *QWinEventNotifier) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQWinEventNotifier() *QWinEventNotifier {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_new, uintptr(__goid))
	__o := newownedQWinEventNotifier(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWinEventNotifier_virtIdx, QWinEventNotifier_enable)
	return __o
}

func NewQWinEventNotifier2(parent *QObject) *QWinEventNotifier {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQWinEventNotifier(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWinEventNotifier_virtIdx, QWinEventNotifier_enable)
	return __o
}

func QWinEventNotifier_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QWinEventNotifier) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QWinEventNotifier_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWinEventNotifier) SetEnabled(enable bool) {
	purego.SyscallN(addr_QWinEventNotifier_setEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func QWinEventNotifier_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWinEventNotifier_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWinEventNotifier) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWinEventNotifier) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QWinEventNotifier_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWinEventNotifier) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QWinEventNotifier_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QWinEventNotifier) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QWinEventNotifier_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QWinEventNotifier_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWinEventNotifier_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWinEventNotifier_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWinEventNotifier_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWinEventNotifier) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QWinEventNotifier_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QWinEventNotifier) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWinEventNotifier_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QWinEventNotifier) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWinEventNotifier_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWinEventNotifier) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QWinEventNotifier_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWinEventNotifier) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QWinEventNotifier_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWinEventNotifier) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWinEventNotifier_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWinEventNotifier) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWinEventNotifier_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWinEventNotifier) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWinEventNotifier_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWinEventNotifier) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWinEventNotifier_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWinEventNotifier) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QWinEventNotifier_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWinEventNotifier) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QWinEventNotifier_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWinEventNotifier) Disconnect(handle int64) {
	pg_QWinEventNotifier__disconnect(handle)
	unregisterCallback(handle)
}

func QWinEventNotifier_virtIdx(slot int64) int32 {
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

func QWinEventNotifier_enable(h uintptr, idx int32) { pg_QWinEventNotifier__set_override(h, idx) }

func (this *QWinEventNotifier) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QWinEventNotifier) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QWinEventNotifier) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QWinEventNotifier) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QWinEventNotifier) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QWinEventNotifier) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QWinEventNotifier) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QWinEventNotifier) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QWinEventNotifier) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWinEventNotifier_delete(this.h)
	}
	this.UnsafeRelease()
}
