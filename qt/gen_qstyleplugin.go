package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStylePlugin_base func(uintptr) uintptr
	addr_QStylePlugin_tr uintptr
	addr_QStylePlugin_create uintptr
	addr_QStylePlugin_connect uintptr
	addr_QStylePlugin_connect2 uintptr
	addr_QStylePlugin_sender uintptr
	addr_QStylePlugin_senderSignalIndex uintptr
	addr_QStylePlugin_receivers uintptr
	addr_QStylePlugin_isSignalConnected uintptr
	addr_QStylePlugin_tr2 uintptr
	addr_QStylePlugin_tr3 uintptr
	addr_QStylePlugin_connect3 uintptr
	addr_QStylePlugin_connect4 uintptr
	addr_QStylePlugin_metaObject_default uintptr
	addr_QStylePlugin_event_default uintptr
	addr_QStylePlugin_eventFilter_default uintptr
	addr_QStylePlugin_timerEvent_default uintptr
	addr_QStylePlugin_childEvent_default uintptr
	addr_QStylePlugin_customEvent_default uintptr
	addr_QStylePlugin_connectNotify_default uintptr
	addr_QStylePlugin_disconnectNotify_default uintptr
	pg_QStylePlugin_connect_destroyed func(uintptr, int64)
	pg_QStylePlugin_connect_destroyed2 func(uintptr, int64)
	pg_QStylePlugin_connect_objectNameChanged func(uintptr, int64)
	pg_QStylePlugin__disconnect func(int64)
	pg_QStylePlugin__set_override func(uintptr, int32)
	pg_QStylePlugin__Go_new func(int64, uintptr) uintptr
	pg_QStylePlugin_delete func(uintptr)
)

func init() {
	reg(&pg_QStylePlugin_base, "QStylePlugin_base")
	symaddr(&addr_QStylePlugin_tr, "QStylePlugin_tr")
	symaddr(&addr_QStylePlugin_create, "QStylePlugin_create")
	symaddr(&addr_QStylePlugin_connect, "QStylePlugin_connect")
	symaddr(&addr_QStylePlugin_connect2, "QStylePlugin_connect2")
	symaddr(&addr_QStylePlugin_sender, "QStylePlugin_sender")
	symaddr(&addr_QStylePlugin_senderSignalIndex, "QStylePlugin_senderSignalIndex")
	symaddr(&addr_QStylePlugin_receivers, "QStylePlugin_receivers")
	symaddr(&addr_QStylePlugin_isSignalConnected, "QStylePlugin_isSignalConnected")
	symaddr(&addr_QStylePlugin_tr2, "QStylePlugin_tr2")
	symaddr(&addr_QStylePlugin_tr3, "QStylePlugin_tr3")
	symaddr(&addr_QStylePlugin_connect3, "QStylePlugin_connect3")
	symaddr(&addr_QStylePlugin_connect4, "QStylePlugin_connect4")
	symaddr(&addr_QStylePlugin_metaObject_default, "QStylePlugin_metaObject_default")
	symaddr(&addr_QStylePlugin_event_default, "QStylePlugin_event_default")
	symaddr(&addr_QStylePlugin_eventFilter_default, "QStylePlugin_eventFilter_default")
	symaddr(&addr_QStylePlugin_timerEvent_default, "QStylePlugin_timerEvent_default")
	symaddr(&addr_QStylePlugin_childEvent_default, "QStylePlugin_childEvent_default")
	symaddr(&addr_QStylePlugin_customEvent_default, "QStylePlugin_customEvent_default")
	symaddr(&addr_QStylePlugin_connectNotify_default, "QStylePlugin_connectNotify_default")
	symaddr(&addr_QStylePlugin_disconnectNotify_default, "QStylePlugin_disconnectNotify_default")
	reg(&pg_QStylePlugin_connect_destroyed, "QStylePlugin_connect_destroyed")
	reg(&pg_QStylePlugin_connect_destroyed2, "QStylePlugin_connect_destroyed2")
	reg(&pg_QStylePlugin_connect_objectNameChanged, "QStylePlugin_connect_objectNameChanged")
	reg(&pg_QStylePlugin__disconnect, "QStylePlugin__disconnect")
	reg(&pg_QStylePlugin__set_override, "QStylePlugin__set_override")
	reg(&pg_QStylePlugin__Go_new, "QStylePlugin__Go_new")
	reg(&pg_QStylePlugin_delete, "QStylePlugin_delete")
}

type QStylePlugin struct {
	h uintptr
	*QObject
}

func bindQStylePlugin(h uintptr) *QStylePlugin {
	if h == 0 {
		return nil
	}
	return &QStylePlugin{h: h, QObject: bindQObject(pg_QStylePlugin_base(h))}
}

func newownedQStylePlugin(h uintptr) *QStylePlugin { return bindQStylePlugin(h) }

func (this *QStylePlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStylePlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStylePlugin(h unsafe.Pointer) *QStylePlugin { return bindQStylePlugin(uintptr(h)) }
func UnsafeNewQStylePluginOwned(h unsafe.Pointer) *QStylePlugin { return newownedQStylePlugin(uintptr(h)) }

func (this *QStylePlugin) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QStylePlugin_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStylePlugin) Create(key string) *QStyle {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_create, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return bindQStyle(__r)
}

func QStylePlugin_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStylePlugin_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStylePlugin) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStylePlugin) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStylePlugin_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStylePlugin) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStylePlugin_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStylePlugin) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStylePlugin_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStylePlugin_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStylePlugin_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStylePlugin_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStylePlugin_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStylePlugin) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStylePlugin_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStylePlugin) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStylePlugin_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStylePlugin) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStylePlugin_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStylePlugin) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStylePlugin_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStylePlugin) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStylePlugin_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStylePlugin) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStylePlugin_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStylePlugin) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStylePlugin_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStylePlugin) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStylePlugin_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStylePlugin) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStylePlugin_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStylePlugin) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStylePlugin_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStylePlugin) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStylePlugin_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStylePlugin) Disconnect(handle int64) {
	pg_QStylePlugin__disconnect(handle)
	unregisterCallback(handle)
}

func QStylePlugin_virtIdx(slot int64) int32 {
	switch slot {
	case -626457835795418115:
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

func QStylePlugin_enable(h uintptr, idx int32) { pg_QStylePlugin__set_override(h, idx) }

func NewQStylePlugin() *QStylePlugin {
	id := nextVirtID()
	o := newownedQStylePlugin(pg_QStylePlugin__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QStylePlugin_virtIdx, QStylePlugin_enable)
	return o
}

func NewQStylePlugin2(parent *QObject) *QStylePlugin {
	id := nextVirtID()
	o := newownedQStylePlugin(pg_QStylePlugin__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QStylePlugin_virtIdx, QStylePlugin_enable)
	return o
}

func (this *QStylePlugin) OnCreate(fn func(string) *QStyle) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(key uintptr, key_len int64) uintptr { return uintptr(fn(pgStrNoFree(key, key_len)).UnsafePointer()) }) // slot "create"
}

func (this *QStylePlugin) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStylePlugin) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QStylePlugin) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStylePlugin) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStylePlugin) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStylePlugin) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStylePlugin) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStylePlugin) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStylePlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStylePlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
