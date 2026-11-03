package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibilityHints_base func(uintptr) uintptr
	addr_QAccessibilityHints_new uintptr
	addr_QAccessibilityHints_new2 uintptr
	addr_QAccessibilityHints_tr uintptr
	addr_QAccessibilityHints_contrastPreference uintptr
	addr_QAccessibilityHints_contrastPreferenceChanged uintptr
	addr_QAccessibilityHints_connect uintptr
	addr_QAccessibilityHints_connect2 uintptr
	addr_QAccessibilityHints_sender uintptr
	addr_QAccessibilityHints_senderSignalIndex uintptr
	addr_QAccessibilityHints_receivers uintptr
	addr_QAccessibilityHints_isSignalConnected uintptr
	addr_QAccessibilityHints_tr2 uintptr
	addr_QAccessibilityHints_tr3 uintptr
	addr_QAccessibilityHints_connect3 uintptr
	addr_QAccessibilityHints_connect4 uintptr
	addr_QAccessibilityHints_metaObject_default uintptr
	addr_QAccessibilityHints_event_default uintptr
	addr_QAccessibilityHints_eventFilter_default uintptr
	addr_QAccessibilityHints_timerEvent_default uintptr
	addr_QAccessibilityHints_childEvent_default uintptr
	addr_QAccessibilityHints_customEvent_default uintptr
	addr_QAccessibilityHints_connectNotify_default uintptr
	addr_QAccessibilityHints_disconnectNotify_default uintptr
	pg_QAccessibilityHints_connect_contrastPreferenceChanged func(uintptr, int64)
	pg_QAccessibilityHints_connect_destroyed func(uintptr, int64)
	pg_QAccessibilityHints_connect_destroyed2 func(uintptr, int64)
	pg_QAccessibilityHints_connect_objectNameChanged func(uintptr, int64)
	pg_QAccessibilityHints__disconnect func(int64)
	pg_QAccessibilityHints__set_override func(uintptr, int32)
	pg_QAccessibilityHints_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibilityHints_base, "QAccessibilityHints_base")
	symaddr(&addr_QAccessibilityHints_new, "QAccessibilityHints_new")
	symaddr(&addr_QAccessibilityHints_new2, "QAccessibilityHints_new2")
	symaddr(&addr_QAccessibilityHints_tr, "QAccessibilityHints_tr")
	symaddr(&addr_QAccessibilityHints_contrastPreference, "QAccessibilityHints_contrastPreference")
	symaddr(&addr_QAccessibilityHints_contrastPreferenceChanged, "QAccessibilityHints_contrastPreferenceChanged")
	symaddr(&addr_QAccessibilityHints_connect, "QAccessibilityHints_connect")
	symaddr(&addr_QAccessibilityHints_connect2, "QAccessibilityHints_connect2")
	symaddr(&addr_QAccessibilityHints_sender, "QAccessibilityHints_sender")
	symaddr(&addr_QAccessibilityHints_senderSignalIndex, "QAccessibilityHints_senderSignalIndex")
	symaddr(&addr_QAccessibilityHints_receivers, "QAccessibilityHints_receivers")
	symaddr(&addr_QAccessibilityHints_isSignalConnected, "QAccessibilityHints_isSignalConnected")
	symaddr(&addr_QAccessibilityHints_tr2, "QAccessibilityHints_tr2")
	symaddr(&addr_QAccessibilityHints_tr3, "QAccessibilityHints_tr3")
	symaddr(&addr_QAccessibilityHints_connect3, "QAccessibilityHints_connect3")
	symaddr(&addr_QAccessibilityHints_connect4, "QAccessibilityHints_connect4")
	symaddr(&addr_QAccessibilityHints_metaObject_default, "QAccessibilityHints_metaObject_default")
	symaddr(&addr_QAccessibilityHints_event_default, "QAccessibilityHints_event_default")
	symaddr(&addr_QAccessibilityHints_eventFilter_default, "QAccessibilityHints_eventFilter_default")
	symaddr(&addr_QAccessibilityHints_timerEvent_default, "QAccessibilityHints_timerEvent_default")
	symaddr(&addr_QAccessibilityHints_childEvent_default, "QAccessibilityHints_childEvent_default")
	symaddr(&addr_QAccessibilityHints_customEvent_default, "QAccessibilityHints_customEvent_default")
	symaddr(&addr_QAccessibilityHints_connectNotify_default, "QAccessibilityHints_connectNotify_default")
	symaddr(&addr_QAccessibilityHints_disconnectNotify_default, "QAccessibilityHints_disconnectNotify_default")
	reg(&pg_QAccessibilityHints_connect_contrastPreferenceChanged, "QAccessibilityHints_connect_contrastPreferenceChanged")
	reg(&pg_QAccessibilityHints_connect_destroyed, "QAccessibilityHints_connect_destroyed")
	reg(&pg_QAccessibilityHints_connect_destroyed2, "QAccessibilityHints_connect_destroyed2")
	reg(&pg_QAccessibilityHints_connect_objectNameChanged, "QAccessibilityHints_connect_objectNameChanged")
	reg(&pg_QAccessibilityHints__disconnect, "QAccessibilityHints__disconnect")
	reg(&pg_QAccessibilityHints__set_override, "QAccessibilityHints__set_override")
	reg(&pg_QAccessibilityHints_delete, "QAccessibilityHints_delete")
}

type QAccessibilityHints struct {
	h uintptr
	*QObject
}

func bindQAccessibilityHints(h uintptr) *QAccessibilityHints {
	if h == 0 {
		return nil
	}
	return &QAccessibilityHints{h: h, QObject: bindQObject(pg_QAccessibilityHints_base(h))}
}

func newownedQAccessibilityHints(h uintptr) *QAccessibilityHints { return bindQAccessibilityHints(h) }

func (this *QAccessibilityHints) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibilityHints) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibilityHints(h unsafe.Pointer) *QAccessibilityHints { return bindQAccessibilityHints(uintptr(h)) }
func UnsafeNewQAccessibilityHintsOwned(h unsafe.Pointer) *QAccessibilityHints { return newownedQAccessibilityHints(uintptr(h)) }

func (this *QAccessibilityHints) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQAccessibilityHints() *QAccessibilityHints {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_new, uintptr(__goid))
	__o := newownedQAccessibilityHints(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAccessibilityHints_virtIdx, QAccessibilityHints_enable)
	return __o
}

func NewQAccessibilityHints2(parent *QObject) *QAccessibilityHints {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQAccessibilityHints(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAccessibilityHints_virtIdx, QAccessibilityHints_enable)
	return __o
}

func QAccessibilityHints_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAccessibilityHints) ContrastPreference() ContrastPreference {
	__r0, _, _ := purego.SyscallN(addr_QAccessibilityHints_contrastPreference, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ContrastPreference(__r)
}

func (this *QAccessibilityHints) ContrastPreferenceChanged(contrastPreference ContrastPreference) {
	purego.SyscallN(addr_QAccessibilityHints_contrastPreferenceChanged, this.h, uintptr(contrastPreference))
	runtime.KeepAlive(this)
}

func QAccessibilityHints_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAccessibilityHints_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAccessibilityHints) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAccessibilityHints) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibilityHints_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibilityHints) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAccessibilityHints_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAccessibilityHints) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibilityHints_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAccessibilityHints_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAccessibilityHints_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAccessibilityHints_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAccessibilityHints_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAccessibilityHints) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibilityHints_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAccessibilityHints) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibilityHints_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAccessibilityHints) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibilityHints_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAccessibilityHints) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAccessibilityHints_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibilityHints) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAccessibilityHints_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibilityHints) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAccessibilityHints_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibilityHints) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAccessibilityHints_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAccessibilityHints) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAccessibilityHints_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAccessibilityHints) OnContrastPreferenceChanged(fn func(ContrastPreference)) int64 {
	id := registerCallback(func(contrastPreference int32) { fn(ContrastPreference(contrastPreference)) })
	pg_QAccessibilityHints_connect_contrastPreferenceChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibilityHints) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAccessibilityHints_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibilityHints) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAccessibilityHints_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibilityHints) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAccessibilityHints_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibilityHints) Disconnect(handle int64) {
	pg_QAccessibilityHints__disconnect(handle)
	unregisterCallback(handle)
}

func QAccessibilityHints_virtIdx(slot int64) int32 {
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

func QAccessibilityHints_enable(h uintptr, idx int32) { pg_QAccessibilityHints__set_override(h, idx) }

func (this *QAccessibilityHints) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAccessibilityHints) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAccessibilityHints) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAccessibilityHints) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAccessibilityHints) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAccessibilityHints) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAccessibilityHints) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAccessibilityHints) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAccessibilityHints) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibilityHints_delete(this.h)
	}
	this.UnsafeRelease()
}
