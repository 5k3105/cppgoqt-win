package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibleBridgePlugin_base func(uintptr) uintptr
	addr_QAccessibleBridgePlugin_tr uintptr
	addr_QAccessibleBridgePlugin_create uintptr
	addr_QAccessibleBridgePlugin_connect uintptr
	addr_QAccessibleBridgePlugin_connect2 uintptr
	addr_QAccessibleBridgePlugin_sender uintptr
	addr_QAccessibleBridgePlugin_senderSignalIndex uintptr
	addr_QAccessibleBridgePlugin_receivers uintptr
	addr_QAccessibleBridgePlugin_isSignalConnected uintptr
	addr_QAccessibleBridgePlugin_tr2 uintptr
	addr_QAccessibleBridgePlugin_tr3 uintptr
	addr_QAccessibleBridgePlugin_connect3 uintptr
	addr_QAccessibleBridgePlugin_connect4 uintptr
	addr_QAccessibleBridgePlugin_metaObject_default uintptr
	addr_QAccessibleBridgePlugin_event_default uintptr
	addr_QAccessibleBridgePlugin_eventFilter_default uintptr
	addr_QAccessibleBridgePlugin_timerEvent_default uintptr
	addr_QAccessibleBridgePlugin_childEvent_default uintptr
	addr_QAccessibleBridgePlugin_customEvent_default uintptr
	addr_QAccessibleBridgePlugin_connectNotify_default uintptr
	addr_QAccessibleBridgePlugin_disconnectNotify_default uintptr
	pg_QAccessibleBridgePlugin_connect_destroyed func(uintptr, int64)
	pg_QAccessibleBridgePlugin_connect_destroyed2 func(uintptr, int64)
	pg_QAccessibleBridgePlugin_connect_objectNameChanged func(uintptr, int64)
	pg_QAccessibleBridgePlugin__disconnect func(int64)
	pg_QAccessibleBridgePlugin__set_override func(uintptr, int32)
	pg_QAccessibleBridgePlugin__Go_new func(int64, uintptr) uintptr
	pg_QAccessibleBridgePlugin_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibleBridgePlugin_base, "QAccessibleBridgePlugin_base")
	symaddr(&addr_QAccessibleBridgePlugin_tr, "QAccessibleBridgePlugin_tr")
	symaddr(&addr_QAccessibleBridgePlugin_create, "QAccessibleBridgePlugin_create")
	symaddr(&addr_QAccessibleBridgePlugin_connect, "QAccessibleBridgePlugin_connect")
	symaddr(&addr_QAccessibleBridgePlugin_connect2, "QAccessibleBridgePlugin_connect2")
	symaddr(&addr_QAccessibleBridgePlugin_sender, "QAccessibleBridgePlugin_sender")
	symaddr(&addr_QAccessibleBridgePlugin_senderSignalIndex, "QAccessibleBridgePlugin_senderSignalIndex")
	symaddr(&addr_QAccessibleBridgePlugin_receivers, "QAccessibleBridgePlugin_receivers")
	symaddr(&addr_QAccessibleBridgePlugin_isSignalConnected, "QAccessibleBridgePlugin_isSignalConnected")
	symaddr(&addr_QAccessibleBridgePlugin_tr2, "QAccessibleBridgePlugin_tr2")
	symaddr(&addr_QAccessibleBridgePlugin_tr3, "QAccessibleBridgePlugin_tr3")
	symaddr(&addr_QAccessibleBridgePlugin_connect3, "QAccessibleBridgePlugin_connect3")
	symaddr(&addr_QAccessibleBridgePlugin_connect4, "QAccessibleBridgePlugin_connect4")
	symaddr(&addr_QAccessibleBridgePlugin_metaObject_default, "QAccessibleBridgePlugin_metaObject_default")
	symaddr(&addr_QAccessibleBridgePlugin_event_default, "QAccessibleBridgePlugin_event_default")
	symaddr(&addr_QAccessibleBridgePlugin_eventFilter_default, "QAccessibleBridgePlugin_eventFilter_default")
	symaddr(&addr_QAccessibleBridgePlugin_timerEvent_default, "QAccessibleBridgePlugin_timerEvent_default")
	symaddr(&addr_QAccessibleBridgePlugin_childEvent_default, "QAccessibleBridgePlugin_childEvent_default")
	symaddr(&addr_QAccessibleBridgePlugin_customEvent_default, "QAccessibleBridgePlugin_customEvent_default")
	symaddr(&addr_QAccessibleBridgePlugin_connectNotify_default, "QAccessibleBridgePlugin_connectNotify_default")
	symaddr(&addr_QAccessibleBridgePlugin_disconnectNotify_default, "QAccessibleBridgePlugin_disconnectNotify_default")
	reg(&pg_QAccessibleBridgePlugin_connect_destroyed, "QAccessibleBridgePlugin_connect_destroyed")
	reg(&pg_QAccessibleBridgePlugin_connect_destroyed2, "QAccessibleBridgePlugin_connect_destroyed2")
	reg(&pg_QAccessibleBridgePlugin_connect_objectNameChanged, "QAccessibleBridgePlugin_connect_objectNameChanged")
	reg(&pg_QAccessibleBridgePlugin__disconnect, "QAccessibleBridgePlugin__disconnect")
	reg(&pg_QAccessibleBridgePlugin__set_override, "QAccessibleBridgePlugin__set_override")
	reg(&pg_QAccessibleBridgePlugin__Go_new, "QAccessibleBridgePlugin__Go_new")
	reg(&pg_QAccessibleBridgePlugin_delete, "QAccessibleBridgePlugin_delete")
}

type QAccessibleBridgePlugin struct {
	h uintptr
	*QObject
}

func bindQAccessibleBridgePlugin(h uintptr) *QAccessibleBridgePlugin {
	if h == 0 {
		return nil
	}
	return &QAccessibleBridgePlugin{h: h, QObject: bindQObject(pg_QAccessibleBridgePlugin_base(h))}
}

func newownedQAccessibleBridgePlugin(h uintptr) *QAccessibleBridgePlugin { return bindQAccessibleBridgePlugin(h) }

func (this *QAccessibleBridgePlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleBridgePlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleBridgePlugin(h unsafe.Pointer) *QAccessibleBridgePlugin { return bindQAccessibleBridgePlugin(uintptr(h)) }
func UnsafeNewQAccessibleBridgePluginOwned(h unsafe.Pointer) *QAccessibleBridgePlugin { return newownedQAccessibleBridgePlugin(uintptr(h)) }

func (this *QAccessibleBridgePlugin) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAccessibleBridgePlugin_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAccessibleBridgePlugin) Create(key string) *QAccessibleBridge {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_create, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return bindQAccessibleBridge(__r)
}

func QAccessibleBridgePlugin_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAccessibleBridgePlugin_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAccessibleBridgePlugin) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAccessibleBridgePlugin) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleBridgePlugin) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAccessibleBridgePlugin) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAccessibleBridgePlugin_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAccessibleBridgePlugin_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAccessibleBridgePlugin_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAccessibleBridgePlugin_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAccessibleBridgePlugin) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAccessibleBridgePlugin) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAccessibleBridgePlugin) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleBridgePlugin_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAccessibleBridgePlugin) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAccessibleBridgePlugin_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibleBridgePlugin) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAccessibleBridgePlugin_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibleBridgePlugin) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAccessibleBridgePlugin_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibleBridgePlugin) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAccessibleBridgePlugin_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAccessibleBridgePlugin) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAccessibleBridgePlugin_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAccessibleBridgePlugin) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAccessibleBridgePlugin_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibleBridgePlugin) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAccessibleBridgePlugin_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibleBridgePlugin) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAccessibleBridgePlugin_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessibleBridgePlugin) Disconnect(handle int64) {
	pg_QAccessibleBridgePlugin__disconnect(handle)
	unregisterCallback(handle)
}

func QAccessibleBridgePlugin_virtIdx(slot int64) int32 {
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

func QAccessibleBridgePlugin_enable(h uintptr, idx int32) { pg_QAccessibleBridgePlugin__set_override(h, idx) }

func NewQAccessibleBridgePlugin() *QAccessibleBridgePlugin {
	id := nextVirtID()
	o := newownedQAccessibleBridgePlugin(pg_QAccessibleBridgePlugin__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleBridgePlugin_virtIdx, QAccessibleBridgePlugin_enable)
	return o
}

func NewQAccessibleBridgePlugin2(parent *QObject) *QAccessibleBridgePlugin {
	id := nextVirtID()
	o := newownedQAccessibleBridgePlugin(pg_QAccessibleBridgePlugin__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleBridgePlugin_virtIdx, QAccessibleBridgePlugin_enable)
	return o
}

func (this *QAccessibleBridgePlugin) OnCreate(fn func(string) *QAccessibleBridge) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(key uintptr, key_len int64) uintptr { return uintptr(fn(pgStrNoFree(key, key_len)).UnsafePointer()) }) // slot "create"
}

func (this *QAccessibleBridgePlugin) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAccessibleBridgePlugin) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAccessibleBridgePlugin) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAccessibleBridgePlugin) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAccessibleBridgePlugin) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAccessibleBridgePlugin) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAccessibleBridgePlugin) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAccessibleBridgePlugin) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAccessibleBridgePlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleBridgePlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
