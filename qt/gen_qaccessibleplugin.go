package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessiblePlugin_base func(uintptr) uintptr
	addr_QAccessiblePlugin_tr uintptr
	addr_QAccessiblePlugin_create uintptr
	addr_QAccessiblePlugin_connect uintptr
	addr_QAccessiblePlugin_connect2 uintptr
	addr_QAccessiblePlugin_sender uintptr
	addr_QAccessiblePlugin_senderSignalIndex uintptr
	addr_QAccessiblePlugin_receivers uintptr
	addr_QAccessiblePlugin_isSignalConnected uintptr
	addr_QAccessiblePlugin_tr2 uintptr
	addr_QAccessiblePlugin_tr3 uintptr
	addr_QAccessiblePlugin_connect3 uintptr
	addr_QAccessiblePlugin_connect4 uintptr
	addr_QAccessiblePlugin_metaObject_default uintptr
	addr_QAccessiblePlugin_event_default uintptr
	addr_QAccessiblePlugin_eventFilter_default uintptr
	addr_QAccessiblePlugin_timerEvent_default uintptr
	addr_QAccessiblePlugin_childEvent_default uintptr
	addr_QAccessiblePlugin_customEvent_default uintptr
	addr_QAccessiblePlugin_connectNotify_default uintptr
	addr_QAccessiblePlugin_disconnectNotify_default uintptr
	pg_QAccessiblePlugin_connect_destroyed func(uintptr, int64)
	pg_QAccessiblePlugin_connect_destroyed2 func(uintptr, int64)
	pg_QAccessiblePlugin_connect_objectNameChanged func(uintptr, int64)
	pg_QAccessiblePlugin__disconnect func(int64)
	pg_QAccessiblePlugin__set_override func(uintptr, int32)
	pg_QAccessiblePlugin__Go_new func(int64, uintptr) uintptr
	pg_QAccessiblePlugin_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessiblePlugin_base, "QAccessiblePlugin_base")
	symaddr(&addr_QAccessiblePlugin_tr, "QAccessiblePlugin_tr")
	symaddr(&addr_QAccessiblePlugin_create, "QAccessiblePlugin_create")
	symaddr(&addr_QAccessiblePlugin_connect, "QAccessiblePlugin_connect")
	symaddr(&addr_QAccessiblePlugin_connect2, "QAccessiblePlugin_connect2")
	symaddr(&addr_QAccessiblePlugin_sender, "QAccessiblePlugin_sender")
	symaddr(&addr_QAccessiblePlugin_senderSignalIndex, "QAccessiblePlugin_senderSignalIndex")
	symaddr(&addr_QAccessiblePlugin_receivers, "QAccessiblePlugin_receivers")
	symaddr(&addr_QAccessiblePlugin_isSignalConnected, "QAccessiblePlugin_isSignalConnected")
	symaddr(&addr_QAccessiblePlugin_tr2, "QAccessiblePlugin_tr2")
	symaddr(&addr_QAccessiblePlugin_tr3, "QAccessiblePlugin_tr3")
	symaddr(&addr_QAccessiblePlugin_connect3, "QAccessiblePlugin_connect3")
	symaddr(&addr_QAccessiblePlugin_connect4, "QAccessiblePlugin_connect4")
	symaddr(&addr_QAccessiblePlugin_metaObject_default, "QAccessiblePlugin_metaObject_default")
	symaddr(&addr_QAccessiblePlugin_event_default, "QAccessiblePlugin_event_default")
	symaddr(&addr_QAccessiblePlugin_eventFilter_default, "QAccessiblePlugin_eventFilter_default")
	symaddr(&addr_QAccessiblePlugin_timerEvent_default, "QAccessiblePlugin_timerEvent_default")
	symaddr(&addr_QAccessiblePlugin_childEvent_default, "QAccessiblePlugin_childEvent_default")
	symaddr(&addr_QAccessiblePlugin_customEvent_default, "QAccessiblePlugin_customEvent_default")
	symaddr(&addr_QAccessiblePlugin_connectNotify_default, "QAccessiblePlugin_connectNotify_default")
	symaddr(&addr_QAccessiblePlugin_disconnectNotify_default, "QAccessiblePlugin_disconnectNotify_default")
	reg(&pg_QAccessiblePlugin_connect_destroyed, "QAccessiblePlugin_connect_destroyed")
	reg(&pg_QAccessiblePlugin_connect_destroyed2, "QAccessiblePlugin_connect_destroyed2")
	reg(&pg_QAccessiblePlugin_connect_objectNameChanged, "QAccessiblePlugin_connect_objectNameChanged")
	reg(&pg_QAccessiblePlugin__disconnect, "QAccessiblePlugin__disconnect")
	reg(&pg_QAccessiblePlugin__set_override, "QAccessiblePlugin__set_override")
	reg(&pg_QAccessiblePlugin__Go_new, "QAccessiblePlugin__Go_new")
	reg(&pg_QAccessiblePlugin_delete, "QAccessiblePlugin_delete")
}

type QAccessiblePlugin struct {
	h uintptr
	*QObject
}

func bindQAccessiblePlugin(h uintptr) *QAccessiblePlugin {
	if h == 0 {
		return nil
	}
	return &QAccessiblePlugin{h: h, QObject: bindQObject(pg_QAccessiblePlugin_base(h))}
}

func newownedQAccessiblePlugin(h uintptr) *QAccessiblePlugin { return bindQAccessiblePlugin(h) }

func (this *QAccessiblePlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessiblePlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessiblePlugin(h unsafe.Pointer) *QAccessiblePlugin { return bindQAccessiblePlugin(uintptr(h)) }
func UnsafeNewQAccessiblePluginOwned(h unsafe.Pointer) *QAccessiblePlugin { return newownedQAccessiblePlugin(uintptr(h)) }

func (this *QAccessiblePlugin) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAccessiblePlugin_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAccessiblePlugin) Create(key string, object *QObject) *QAccessibleInterface {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_create, this.h, key_p, uintptr(len(key)), uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(object)
	return bindQAccessibleInterface(__r)
}

func QAccessiblePlugin_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAccessiblePlugin_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAccessiblePlugin) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAccessiblePlugin) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessiblePlugin_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessiblePlugin) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAccessiblePlugin_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAccessiblePlugin) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessiblePlugin_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAccessiblePlugin_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAccessiblePlugin_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAccessiblePlugin_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAccessiblePlugin_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAccessiblePlugin) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAccessiblePlugin_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAccessiblePlugin) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessiblePlugin_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAccessiblePlugin) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessiblePlugin_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAccessiblePlugin) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAccessiblePlugin_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessiblePlugin) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAccessiblePlugin_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessiblePlugin) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAccessiblePlugin_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessiblePlugin) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAccessiblePlugin_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAccessiblePlugin) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAccessiblePlugin_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAccessiblePlugin) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAccessiblePlugin_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessiblePlugin) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAccessiblePlugin_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessiblePlugin) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAccessiblePlugin_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAccessiblePlugin) Disconnect(handle int64) {
	pg_QAccessiblePlugin__disconnect(handle)
	unregisterCallback(handle)
}

func QAccessiblePlugin_virtIdx(slot int64) int32 {
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

func QAccessiblePlugin_enable(h uintptr, idx int32) { pg_QAccessiblePlugin__set_override(h, idx) }

func NewQAccessiblePlugin() *QAccessiblePlugin {
	id := nextVirtID()
	o := newownedQAccessiblePlugin(pg_QAccessiblePlugin__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessiblePlugin_virtIdx, QAccessiblePlugin_enable)
	return o
}

func NewQAccessiblePlugin2(parent *QObject) *QAccessiblePlugin {
	id := nextVirtID()
	o := newownedQAccessiblePlugin(pg_QAccessiblePlugin__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessiblePlugin_virtIdx, QAccessiblePlugin_enable)
	return o
}

func (this *QAccessiblePlugin) OnCreate(fn func(string, *QObject) *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(key uintptr, key_len int64, object uintptr) uintptr { return uintptr(fn(pgStrNoFree(key, key_len), UnsafeNewQObject(unsafe.Pointer(object))).UnsafePointer()) }) // slot "create"
}

func (this *QAccessiblePlugin) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAccessiblePlugin) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAccessiblePlugin) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAccessiblePlugin) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAccessiblePlugin) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAccessiblePlugin) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAccessiblePlugin) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAccessiblePlugin) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAccessiblePlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessiblePlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
